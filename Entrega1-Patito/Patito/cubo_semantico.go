package main

import "fmt"

type CuboSemantico struct {
	tabla map[string]string
}

func clave(tipoIzq, tipoDer, operador string) string {
	return tipoIzq + "," + tipoDer + "," + operador
}

func NuevoCuboSemantico() *CuboSemantico {
	cs := &CuboSemantico{tabla: make(map[string]string)}
	cs.inicializar()
	return cs
}

func (cs *CuboSemantico) set(tipoIzq, tipoDer, operador, resultado string) {
	cs.tabla[clave(tipoIzq, tipoDer, operador)] = resultado
}

func (cs *CuboSemantico) inicializar() {
	// Aritmética: entero y flotante
	for _, op := range []string{"+", "-", "*", "/"} {
		cs.set("entero", "entero", op, "entero")
		cs.set("entero", "flotante", op, "flotante")
		cs.set("flotante", "entero", op, "flotante")
		cs.set("flotante", "flotante", op, "flotante")
	}

	// Comparación numérica: resultado es entero (1 = verdadero, 0 = falso)
	for _, op := range []string{">", "<", "!=", "=="} {
		for _, t1 := range []string{"entero", "flotante"} {
			for _, t2 := range []string{"entero", "flotante"} {
				cs.set(t1, t2, op, "entero")
			}
		}
	}

	// Strings: concatenación con '+', igualdad/desigualdad con '==' y '!='.
	cs.set("string", "string", "+", "string")
	cs.set("string", "string", "==", "entero")
	cs.set("string", "string", "!=", "entero")
}

func (cs *CuboSemantico) Consultar(tipoIzq, tipoDer, operador string) (string, error) {
	resultado, ok := cs.tabla[clave(tipoIzq, tipoDer, operador)]
	if !ok {
		return "error", fmt.Errorf("tipo inválido: %s %s %s", tipoIzq, operador, tipoDer)
	}
	return resultado, nil
}
