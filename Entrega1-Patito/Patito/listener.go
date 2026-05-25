package main

import (
	"fmt"
	"os"
	"patito/parser"
)

type PatitoListener struct {
	*parser.BasegramaticaListener

	pilaOperandos  Pila
	pilaTipos      Pila
	pilaOperadores Pila
	pilaSaltos     Pila // índices de cuádruplos pendientes de backpatch
	pilaRetornos   Pila // índice de inicio de ciclo (para GOTO de regreso)

	cuadruplos Cuadruplos

	cuboSemantico      *CuboSemantico
	tablaVariables     map[string]string
	contadorTemporales int
}

func NuevoPatitoListener() *PatitoListener {
	return &PatitoListener{
		cuboSemantico:  NuevoCuboSemantico(),
		tablaVariables: make(map[string]string),
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
		l.tablaVariables[nombre] = tipo
	}
}

// Registra parámetros de función en la tabla de variables
func (l *PatitoListener) ExitFuncr(ctx *parser.FuncrContext) {
	if ctx.ID() == nil {
		return
	}
	tipo := tipoDesdeCtx(ctx.Tipo().(*parser.TipoContext))
	if tipo != "" {
		l.tablaVariables[ctx.ID().GetText()] = tipo
	}
}

// --- Factores (operandos hoja) ---

func (l *PatitoListener) ExitFactor(ctx *parser.FactorContext) {
	switch {
	case ctx.ID() != nil && ctx.Llamada() == nil:
		nombre := ctx.ID().GetText()
		tipo, ok := l.tablaVariables[nombre]
		if !ok {
			fmt.Printf("Variable no declarada: %s\n", nombre)
			os.Exit(1)
		}
		Push(&l.pilaOperandos, nombre)
		Push(&l.pilaTipos, tipo)

	case ctx.Cte() != nil:
		valor := ctx.Cte().GetText()
		cteCtx := ctx.Cte().(*parser.CteContext)
		tipo := "entero"
		if cteCtx.CTE_FLOAT() != nil {
			tipo = "flotante"
		}
		Push(&l.pilaOperandos, valor)
		Push(&l.pilaTipos, tipo)

	case ctx.LETRERO() != nil:
		Push(&l.pilaOperandos, ctx.LETRERO().GetText())
		Push(&l.pilaTipos, "string")

	case ctx.MENOS() != nil:
		// Unario negativo: genera 0 - operando
		operando := Pop(&l.pilaOperandos).(string)
		tipoOp := Pop(&l.pilaTipos).(string)
		t := fmt.Sprintf("t%d", l.contadorTemporales)
		l.contadorTemporales++
		l.cuadruplos.AgregarCuadruplo("0", operando, "-", t)
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

// generarCuadruplo es un helper que pop dos operandos+tipos, consulta el cubo y emite el cuádruplo.
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

	t := fmt.Sprintf("t%d", l.contadorTemporales)
	l.contadorTemporales++

	l.cuadruplos.AgregarCuadruplo(opIzq, opDer, operador, t)
	Push(&l.pilaOperandos, t)
	Push(&l.pilaTipos, tipoRes)
}

// --- Asignación ---

func (l *PatitoListener) ExitAsigna(ctx *parser.AsignaContext) {
	nombre := ctx.ID().GetText()

	valor := Pop(&l.pilaOperandos).(string)
	tipoValor := Pop(&l.pilaTipos).(string)

	tipoVar, ok := l.tablaVariables[nombre]
	if !ok {
		fmt.Printf("Variable no declarada: %s\n", nombre)
		os.Exit(1)
	}

	// Compatibilidad de tipos: mismo tipo, flotante←entero, boolean←entero (resultado relacional)
	if tipoVar != tipoValor &&
		!(tipoVar == "flotante" && tipoValor == "entero") &&
		!(tipoVar == "boolean" && tipoValor == "entero") {
		fmt.Printf("Error semántico en asignación: no se puede asignar %s a variable de tipo %s\n", tipoValor, tipoVar)
		os.Exit(1)
	}

	l.cuadruplos.AgregarCuadruplo(valor, "_", "=", nombre)
}

// --- Expresión: punto neurálgico para condicion, ciclo e imprime ---

func (l *PatitoListener) ExitExpresion(ctx *parser.ExpresionContext) {
	switch ctx.GetParent().(type) {

	case *parser.CondicionContext:
		// Generar GotoF con destino pendiente (backpatch)
		cond := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		idx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo(cond, "_", "GotoF", "_")
		Push(&l.pilaSaltos, idx)

	case *parser.CicloContext:
		// Generar GotoF con destino pendiente (backpatch)
		cond := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		idx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo(cond, "_", "GotoF", "_")
		Push(&l.pilaSaltos, idx)

	case *parser.ExpresionesContext:
		// Cada expresión en escribe(...) genera un PRINT
		val := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		l.cuadruplos.AgregarCuadruplo(val, "_", "PRINT", "_")

	case *parser.LlamadaexpContext:
		// Cada argumento de llamada genera un PARAM
		val := Pop(&l.pilaOperandos).(string)
		Pop(&l.pilaTipos)
		l.cuadruplos.AgregarCuadruplo(val, "_", "PARAM", "_")
	}
}

// --- Ciclo (mientras) ---

func (l *PatitoListener) EnterCiclo(ctx *parser.CicloContext) {
	// Guardar índice de inicio de condición para el salto de regreso
	Push(&l.pilaRetornos, l.cuadruplos.Len())
}

func (l *PatitoListener) ExitCiclo(ctx *parser.CicloContext) {
	inicio := Pop(&l.pilaRetornos).(int)
	// GOTO regresa al inicio de la condición
	l.cuadruplos.AgregarCuadruplo("_", "_", "GOTO", fmt.Sprintf("%d", inicio))
	// Backpatch del GotoF: salta al cuádruplo siguiente (fuera del ciclo)
	gotoFIdx := Pop(&l.pilaSaltos).(int)
	l.cuadruplos.Backpatch(gotoFIdx, fmt.Sprintf("%d", l.cuadruplos.Len()))
}

// --- Condición (si/sino) ---

func (l *PatitoListener) EnterSinoop(ctx *parser.SinoopContext) {
	if ctx.SINO() != nil {
		// Emitir GOTO para saltar el bloque sino (destino pendiente)
		gotoIdx := l.cuadruplos.Len()
		l.cuadruplos.AgregarCuadruplo("_", "_", "GOTO", "_")
		// Backpatch del GotoF: el bloque sino empieza justo aquí (después del GOTO)
		gotoFIdx := Pop(&l.pilaSaltos).(int)
		l.cuadruplos.Backpatch(gotoFIdx, fmt.Sprintf("%d", l.cuadruplos.Len()))
		// El GOTO queda pendiente para ExitSinoop
		Push(&l.pilaSaltos, gotoIdx)
	}
}

func (l *PatitoListener) ExitSinoop(ctx *parser.SinoopContext) {
	// Backpatch del salto pendiente (GotoF sin sino, o GOTO con sino)
	idx := Pop(&l.pilaSaltos).(int)
	l.cuadruplos.Backpatch(idx, fmt.Sprintf("%d", l.cuadruplos.Len()))
}

// --- Imprime: letreros (literales string) ---

func (l *PatitoListener) EnterLetreros(ctx *parser.LetrerosContext) {
	// Emitir PRINT en Enter para respetar el orden izquierda→derecha
	// (la gramática es recursiva a la derecha)
	if ctx.LETRERO() != nil {
		l.cuadruplos.AgregarCuadruplo(ctx.LETRERO().GetText(), "_", "PRINT", "_")
	}
}

// --- Llamada a función ---

func (l *PatitoListener) ExitLlamada(ctx *parser.LlamadaContext) {
	nombre := ctx.ID().GetText()
	l.cuadruplos.AgregarCuadruplo(nombre, "_", "CALL", "_")
}
