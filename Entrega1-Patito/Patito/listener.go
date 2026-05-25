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

func extraerIDs(ctx parser.IIdopContext) []string {
	ids := []string{ctx.(*parser.IdopContext).ID().GetText()}
	if ctx.(*parser.IdopContext).Idop() != nil {
		ids = append(ids, extraerIDs(ctx.(*parser.IdopContext).Idop())...)
	}
	return ids
}

func (l *PatitoListener) ExitVars(ctx *parser.VarsContext) {
	tipoCtx := ctx.Tipo().(*parser.TipoContext)
	var tipo string
	switch {
	case tipoCtx.ENTERO() != nil:
		tipo = "entero"
	case tipoCtx.FLOTANTE() != nil:
		tipo = "flotante"
	default:
		fmt.Printf("Tipo no soportado: %s\n", tipoCtx.GetText())
		os.Exit(1)
	}

	for _, nombre := range extraerIDs(ctx.Idop()) {
		l.tablaVariables[nombre] = tipo
	}
}

func (l *PatitoListener) ExitFactor(ctx *parser.FactorContext) {
	if ctx.ID() != nil {
		nombre := ctx.ID().GetText()
		tipo := l.tablaVariables[nombre]
		Push(&l.pilaOperandos, nombre)
		Push(&l.pilaTipos, tipo)
	} else if ctx.Cte() != nil {
		valor := ctx.Cte().GetText()
		var tipo string
		cteCtx := ctx.Cte().(*parser.CteContext)
		if cteCtx.CTE_FLOAT() != nil {
			tipo = "flotante"
		} else {
			tipo = "entero"
		}
		Push(&l.pilaOperandos, valor)
		Push(&l.pilaTipos, tipo)
	}
}

func (l *PatitoListener) EnterExopc(ctx *parser.ExopcContext) {
	if ctx.MAS() != nil {
		Push(&l.pilaOperadores, "+")
	} else if ctx.MENOS() != nil {
		Push(&l.pilaOperadores, "-")
	}
}

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

	temporal := fmt.Sprintf("t%d", l.contadorTemporales)
	l.contadorTemporales++

	l.cuadruplos.AgregarCuadruplo(opIzq, opDer, operador, temporal)

	Push(&l.pilaOperandos, temporal)
	Push(&l.pilaTipos, tipoRes)
}

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

	temporal := fmt.Sprintf("t%d", l.contadorTemporales)
	l.contadorTemporales++

	l.cuadruplos.AgregarCuadruplo(opIzq, opDer, operador, temporal)

	Push(&l.pilaOperandos, temporal)
	Push(&l.pilaTipos, tipoRes)
}

func (l *PatitoListener) ExitAsigna(ctx *parser.AsignaContext) {
	nombre := ctx.ID().GetText()

	valor := Pop(&l.pilaOperandos).(string)
	tipoValor := Pop(&l.pilaTipos).(string)

	tipoVar := l.tablaVariables[nombre]
	// flotante acepta entero (promoción implícita); tipos iguales siempre son válidos
	if tipoVar != tipoValor && !(tipoVar == "flotante" && tipoValor == "entero") {
		fmt.Printf("Error semántico en asignación: no se puede asignar %s a variable de tipo %s\n", tipoValor, tipoVar)
		os.Exit(1)
	}

	l.cuadruplos.AgregarCuadruplo(valor, "", "=", nombre)
}

func (l *PatitoListener) ExitExopc(ctx *parser.ExopcContext) {
	if ctx.MAS() == nil && ctx.MENOS() == nil {
		return
	}

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

	temporal := fmt.Sprintf("t%d", l.contadorTemporales)
	l.contadorTemporales++

	l.cuadruplos.AgregarCuadruplo(opIzq, opDer, operador, temporal)

	Push(&l.pilaOperandos, temporal)
	Push(&l.pilaTipos, tipoRes)
}
