package main

import (
	"fmt"
	"os"
	"patito/parser"
)

type LlamadaCtx struct {
	info     *FuncionInfo
	numParam int
}

type PatitoListener struct {
	*parser.BasegramaticaListener

	pilaOperandos  Pila // direcciones virtuales (como string)
	pilaTipos      Pila
	pilaOperadores Pila
	pilaSaltos     Pila // índices de cuádruplos pendientes de backpatch
	pilaRetornos   Pila // índice de inicio de ciclo (para GOTO de regreso)

	cuadruplos Cuadruplos

	cuboSemantico *CuboSemantico
	adminMem      *AdministradorMemoria

	varsGlobales  map[string]*Variable
	dirFunciones  map[string]*FuncionInfo
	funcionActual *FuncionInfo // nil cuando el ámbito es global
	ambitoActual  string

	mainGotoIdx  int
	pilaLlamadas []*LlamadaCtx
}

func NuevoPatitoListener() *PatitoListener {
	return &PatitoListener{
		cuboSemantico: NuevoCuboSemantico(),
		adminMem:      NuevoAdministradorMemoria(),
		varsGlobales:  make(map[string]*Variable),
		dirFunciones:  make(map[string]*FuncionInfo),
		ambitoActual:  "global",
	}
}

// --- Helpers de ámbito y memoria ---

func dirStr(dir int) string {
	return fmt.Sprintf("%d", dir)
}

func (l *PatitoListener) buscarVariable(nombre string) (*Variable, bool) {
	if l.funcionActual != nil {
		if v, ok := l.funcionActual.variables[nombre]; ok {
			return v, true
		}
	}
	if v, ok := l.varsGlobales[nombre]; ok {
		return v, true
	}
	return nil, false
}

func (l *PatitoListener) declararVariable(nombre, tipo string) *Variable {
	if l.funcionActual != nil {
		if _, ok := l.funcionActual.variables[nombre]; ok {
			fmt.Printf("Variable ya declarada en la función: %s\n", nombre)
			os.Exit(1)
		}
		v := &Variable{nombre: nombre, tipo: tipo, direccion: l.adminMem.Asignar("local", tipo)}
		l.funcionActual.variables[nombre] = v
		return v
	}
	if _, ok := l.varsGlobales[nombre]; ok {
		fmt.Printf("Variable global ya declarada: %s\n", nombre)
		os.Exit(1)
	}
	v := &Variable{nombre: nombre, tipo: tipo, direccion: l.adminMem.Asignar("global", tipo)}
	l.varsGlobales[nombre] = v
	return v
}

func (l *PatitoListener) nuevoTemporal(tipo string) string {
	return dirStr(l.adminMem.Asignar("temporal", tipo))
}

func compatibles(destino, origen string) bool {
	return destino == origen ||
		(destino == "flotante" && origen == "entero") ||
		(destino == "boolean" && origen == "entero")
}

// --- Programa: GOTO al main ---

func (l *PatitoListener) EnterPrograma(ctx *parser.ProgramaContext) {
	l.mainGotoIdx = l.cuadruplos.Len()
	l.cuadruplos.AgregarCuadruplo("_", "_", "GOTO", "_")
}

// EnterCuerpo: backpatch del GOTO principal cuando inicia el cuerpo del programa
func (l *PatitoListener) EnterCuerpo(ctx *parser.CuerpoContext) {
	if _, ok := ctx.GetParent().(*parser.ProgramaContext); ok {
		l.cuadruplos.Backpatch(l.mainGotoIdx, fmt.Sprintf("%d", l.cuadruplos.Len()))
	}
}

// --- Declaración de variables ---

func extraerIDs(ctx parser.IIdopContext) []string {
	ids := []string{ctx.(*parser.IdopContext).ID().GetText()}
	if ctx.(*parser.IdopContext).Idop() != nil {
		ids = append(ids, extraerIDs(ctx.(*parser.IdopContext).Idop())...)
	}
	return ids
}

func tipoDesdeCtx(tipoCtx *parser.TipoContext) string {
	switch {
	case tipoCtx.ENTERO() != nil:
		return "entero"
	case tipoCtx.FLOTANTE() != nil:
		return "flotante"
	case tipoCtx.BOOLEAN() != nil:
		return "boolean"
	case tipoCtx.STRING() != nil:
		return "string"
	}
	return ""
}

func (l *PatitoListener) ExitVars(ctx *parser.VarsContext) {
	tipo := tipoDesdeCtx(ctx.Tipo().(*parser.TipoContext))
	if tipo == "" {
		fmt.Printf("Tipo no soportado: %s\n", ctx.Tipo().GetText())
		os.Exit(1)
	}
	for _, nombre := range extraerIDs(ctx.Idop()) {
		l.declararVariable(nombre, tipo)
	}
}

// --- Declaración de funciones ---

func (l *PatitoListener) EnterFuncs(ctx *parser.FuncsContext) {
	nombre := ctx.ID().GetText()
	if _, ok := l.dirFunciones[nombre]; ok {
		fmt.Printf("Función ya declarada: %s\n", nombre)
		os.Exit(1)
	}

	tipoRetorno := "nula"
	fop := ctx.Funcsopc().(*parser.FuncsopcContext)
	if fop.NULA() == nil && fop.Tipo() != nil {
		tipoRetorno = tipoDesdeCtx(fop.Tipo().(*parser.TipoContext))
	}

	info := NuevaFuncionInfo(nombre, tipoRetorno)
	// Las declaraciones (params y vars locales) no generan cuádruplos,
	// por lo que el cuerpo de la función inicia en el cuádruplo actual.
	info.dirInicio = l.cuadruplos.Len()

	// Para funciones con retorno, reservar una dirección global para el valor.
	if tipoRetorno != "nula" {
		retVar := &Variable{nombre: nombre, tipo: tipoRetorno, direccion: l.adminMem.Asignar("global", tipoRetorno)}
		l.varsGlobales[nombre] = retVar
		info.dirRetorno = retVar.direccion
	}

	l.dirFunciones[nombre] = info
	l.funcionActual = info
	l.ambitoActual = nombre
}

func (l *PatitoListener) ExitFuncs(ctx *parser.FuncsContext) {
	l.cuadruplos.AgregarCuadruplo("_", "_", "ENDPROC", "_")
	l.adminMem.ResetLocal()
	l.funcionActual = nil
	l.ambitoActual = "global"
}

// EnterFuncr registra parámetros (orden izquierda→derecha) con dirección local.
func (l *PatitoListener) EnterFuncr(ctx *parser.FuncrContext) {
	if ctx.ID() == nil || l.funcionActual == nil {
		return
	}
	tipo := tipoDesdeCtx(ctx.Tipo().(*parser.TipoContext))
	if tipo == "" {
		return
	}
	nombre := ctx.ID().GetText()
	v := l.declararVariable(nombre, tipo)
	l.funcionActual.parametros = append(l.funcionActual.parametros, v)
}

// --- Factores (operandos hoja) ---

func (l *PatitoListener) ExitFactor(ctx *parser.FactorContext) {
	switch {
	case ctx.ID() != nil && ctx.Llamada() == nil:
		nombre := ctx.ID().GetText()
		v, ok := l.buscarVariable(nombre)
		if !ok {
			fmt.Printf("Variable no declarada: %s\n", nombre)
			os.Exit(1)
		}
		Push(&l.pilaOperandos, dirStr(v.direccion))
		Push(&l.pilaTipos, v.tipo)

	case ctx.Cte() != nil:
		valor := ctx.Cte().GetText()
		cteCtx := ctx.Cte().(*parser.CteContext)
		tipo := "entero"
		if cteCtx.CTE_FLOAT() != nil {
			tipo = "flotante"
		}
		dir := l.adminMem.Constante(valor, tipo)
		Push(&l.pilaOperandos, dirStr(dir))
		Push(&l.pilaTipos, tipo)

	case ctx.LETRERO() != nil:
		dir := l.adminMem.Constante(ctx.LETRERO().GetText(), "string")
		Push(&l.pilaOperandos, dirStr(dir))
		Push(&l.pilaTipos, "string")

	case ctx.MENOS() != nil:
		// Unario negativo: genera 0 - operando
		operando := Pop(&l.pilaOperandos).(string)
		tipoOp := Pop(&l.pilaTipos).(string)
		cero := dirStr(l.adminMem.Constante("0", "entero"))
		t := l.nuevoTemporal(tipoOp)
		l.cuadruplos.AgregarCuadruplo(cero, operando, "-", t)
		Push(&l.pilaOperandos, t)
		Push(&l.pilaTipos, tipoOp)

		// MAS (unario +): dejar el resultado en la pila sin cambios
		// LPAR expresion RPAR: resultado ya está en la pila
		// llamada: manejado en ExitLlamada
	}
}

// --- Operadores aritméticos: + y - ---

func (l *PatitoListener) EnterExopc(ctx *parser.ExopcContext) {
	if ctx.MAS() != nil {
		Push(&l.pilaOperadores, "+")
	} else if ctx.MENOS() != nil {
		Push(&l.pilaOperadores, "-")
	}
}

func (l *PatitoListener) ExitExopc(ctx *parser.ExopcContext) {
	if ctx.MAS() == nil && ctx.MENOS() == nil {
		return
	}
	l.generarCuadruplo()
}

// --- Operadores aritméticos: * y / ---

func (l *PatitoListener) EnterTeropc(ctx *parser.TeropcContext) {
	if ctx.MULT() != nil {
		Push(&l.pilaOperadores, "*")
	} else if ctx.DIV() != nil {
		Push(&l.pilaOperadores, "/")
	}
}

func (l *PatitoListener) ExitTeropc(ctx *parser.TeropcContext) {
	if ctx.MULT() == nil && ctx.DIV() == nil {
		return
	}
	l.generarCuadruplo()
}

// --- Operadores relacionales ---

func (l *PatitoListener) EnterOpc(ctx *parser.OpcContext) {
	switch {
	case ctx.MAYOR() != nil:
		Push(&l.pilaOperadores, ">")
	case ctx.MENOR() != nil:
		Push(&l.pilaOperadores, "<")
	case ctx.NEQ() != nil:
		Push(&l.pilaOperadores, "!=")
	case ctx.EQ() != nil:
		Push(&l.pilaOperadores, "==")
	}
}

func (l *PatitoListener) ExitOpc(ctx *parser.OpcContext) {
	if ctx.MAYOR() == nil && ctx.MENOR() == nil && ctx.NEQ() == nil && ctx.EQ() == nil {
		return
	}
	l.generarCuadruplo()
}

// generarCuadruplo pop dos operandos+tipos, consulta el cubo y emite el cuádruplo.
func (l *PatitoListener) generarCuadruplo() {
	operador := Pop(&l.pilaOperadores).(string)

	opDer := Pop(&l.pilaOperandos).(string)
	tipoDer := Pop(&l.pilaTipos).(string)

	opIzq := Pop(&l.pilaOperandos).(string)
	tipoIzq := Pop(&l.pilaTipos).(string)

	tipoRes, err := l.cuboSemantico.Consultar(tipoIzq, tipoDer, operador)
	if err != nil {
		fmt.Printf("Error semántico: %v\n", err)
		os.Exit(1)
	}

	t := l.nuevoTemporal(tipoRes)
	l.cuadruplos.AgregarCuadruplo(opIzq, opDer, operador, t)
	Push(&l.pilaOperandos, t)
	Push(&l.pilaTipos, tipoRes)
}

// --- Asignación ---

func (l *PatitoListener) ExitAsigna(ctx *parser.AsignaContext) {
	nombre := ctx.ID().GetText()

	valor := Pop(&l.pilaOperandos).(string)
	tipoValor := Pop(&l.pilaTipos).(string)

	v, ok := l.buscarVariable(nombre)
	if !ok {
		fmt.Printf("Variable no declarada: %s\n", nombre)
		os.Exit(1)
	}

	if !compatibles(v.tipo, tipoValor) {
		fmt.Printf("Error semántico en asignación: no se puede asignar %s a variable de tipo %s\n", tipoValor, v.tipo)
		os.Exit(1)
	}

	l.cuadruplos.AgregarCuadruplo(valor, "_", "=", dirStr(v.direccion))
}

// --- Expresión: punto neurálgico para condicion, ciclo, imprime y argumentos ---

func (l *PatitoListener) ExitExpresion(ctx *parser.ExpresionContext) {
	switch ctx.GetParent().(type) {

	case *parser.CondicionContext:
		cond := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		idx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo(cond, "_", "GotoF", "_")
		Push(&l.pilaSaltos, idx)

	case *parser.CicloContext:
		cond := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		idx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo(cond, "_", "GotoF", "_")
		Push(&l.pilaSaltos, idx)

	case *parser.ExpresionesContext:
		val := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		l.cuadruplos.AgregarCuadruplo(val, "_", "PRINT", "_")

	case *parser.LlamadaexpContext:
		// Argumento de llamada: validar tipo y emitir PARAM con su índice
		val := Pop(&l.pilaOperandos).(string)
		tipoArg := Pop(&l.pilaTipos).(string)
		if len(l.pilaLlamadas) == 0 {
			return
		}
		llam := l.pilaLlamadas[len(l.pilaLlamadas)-1]
		if llam.numParam >= len(llam.info.parametros) {
			fmt.Printf("Error: demasiados argumentos en llamada a %s\n", llam.info.nombre)
			os.Exit(1)
		}
		param := llam.info.parametros[llam.numParam]
		if !compatibles(param.tipo, tipoArg) {
			fmt.Printf("Error: argumento %d de %s espera %s, recibió %s\n",
				llam.numParam+1, llam.info.nombre, param.tipo, tipoArg)
			os.Exit(1)
		}
		l.cuadruplos.AgregarCuadruplo(val, "_", "PARAM", fmt.Sprintf("param%d", llam.numParam+1))
		llam.numParam++
	}
}

// --- Ciclo (mientras) ---

func (l *PatitoListener) EnterCiclo(ctx *parser.CicloContext) {
	Push(&l.pilaRetornos, l.cuadruplos.Len())
}

func (l *PatitoListener) ExitCiclo(ctx *parser.CicloContext) {
	inicio := Pop(&l.pilaRetornos).(int)
	l.cuadruplos.AgregarCuadruplo("_", "_", "GOTO", fmt.Sprintf("%d", inicio))
	gotoFIdx := Pop(&l.pilaSaltos).(int)
	l.cuadruplos.Backpatch(gotoFIdx, fmt.Sprintf("%d", l.cuadruplos.Len()))
}

// --- Condición (si/sino) ---

func (l *PatitoListener) EnterSinoop(ctx *parser.SinoopContext) {
	if ctx.SINO() != nil {
		gotoIdx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo("_", "_", "GOTO", "_")
		gotoFIdx := Pop(&l.pilaSaltos).(int)
		l.cuadruplos.Backpatch(gotoFIdx, fmt.Sprintf("%d", l.cuadruplos.Len()))
		Push(&l.pilaSaltos, gotoIdx)
	}
}

func (l *PatitoListener) ExitSinoop(ctx *parser.SinoopContext) {
	idx := Pop(&l.pilaSaltos).(int)
	l.cuadruplos.Backpatch(idx, fmt.Sprintf("%d", l.cuadruplos.Len()))
}

// --- Imprime: letreros (literales string) ---

func (l *PatitoListener) EnterLetreros(ctx *parser.LetrerosContext) {
	// PRINT en Enter para respetar el orden izquierda→derecha (gramática recursiva a la derecha)
	if ctx.LETRERO() != nil {
		dir := l.adminMem.Constante(ctx.LETRERO().GetText(), "string")
		l.cuadruplos.AgregarCuadruplo(dirStr(dir), "_", "PRINT", "_")
	}
}

// --- Llamada a función ---

func (l *PatitoListener) EnterLlamada(ctx *parser.LlamadaContext) {
	nombre := ctx.ID().GetText()
	info, ok := l.dirFunciones[nombre]
	if !ok {
		fmt.Printf("Función no declarada: %s\n", nombre)
		os.Exit(1)
	}
	l.cuadruplos.AgregarCuadruplo(nombre, "_", "ERA", "_")
	l.pilaLlamadas = append(l.pilaLlamadas, &LlamadaCtx{info: info, numParam: 0})
}

func (l *PatitoListener) ExitLlamada(ctx *parser.LlamadaContext) {
	llam := l.pilaLlamadas[len(l.pilaLlamadas)-1]
	l.pilaLlamadas = l.pilaLlamadas[:len(l.pilaLlamadas)-1]

	if llam.numParam != len(llam.info.parametros) {
		fmt.Printf("Error: %s espera %d argumentos, recibió %d\n",
			llam.info.nombre, len(llam.info.parametros), llam.numParam)
		os.Exit(1)
	}

	l.cuadruplos.AgregarCuadruplo(llam.info.nombre, "_", "GOSUB", fmt.Sprintf("%d", llam.info.dirInicio))

	// Si la función regresa valor, copiarlo a un temporal y dejarlo en la pila.
	if llam.info.tipoRetorno != "nula" && llam.info.dirRetorno != -1 {
		t := l.nuevoTemporal(llam.info.tipoRetorno)
		l.cuadruplos.AgregarCuadruplo(dirStr(llam.info.dirRetorno), "_", "=", t)
		Push(&l.pilaOperandos, t)
		Push(&l.pilaTipos, llam.info.tipoRetorno)
	}
}
