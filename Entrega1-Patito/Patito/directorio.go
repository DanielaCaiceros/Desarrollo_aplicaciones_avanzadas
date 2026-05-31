package main

import "fmt"

// Variable es una entrada de la tabla de variables (con su dirección virtual).
type Variable struct {
	nombre    string
	tipo      string
	direccion int
}

// FuncionInfo es una entrada del directorio de funciones.
type FuncionInfo struct {
	nombre      string
	tipoRetorno string
	dirInicio   int // cuádruplo donde inicia la función
	dirRetorno  int // dirección global del valor de retorno (-1 si nula)
	parametros  []*Variable
	variables   map[string]*Variable
}

func NuevaFuncionInfo(nombre, tipoRetorno string) *FuncionInfo {
	return &FuncionInfo{
		nombre:      nombre,
		tipoRetorno: tipoRetorno,
		dirRetorno:  -1,
		variables:   make(map[string]*Variable),
	}
}

// ImprimirDirectorioFunciones muestra el directorio de funciones generado.
func ImprimirDirectorioFunciones(dir map[string]*FuncionInfo) {
	fmt.Println("\n=== Directorio de Funciones ===")
	for _, f := range dir {
		params := ""
		for i, p := range f.parametros {
			if i > 0 {
				params += ", "
			}
			params += fmt.Sprintf("%s:%s@%d", p.nombre, p.tipo, p.direccion)
		}
		fmt.Printf("  %s (ret=%s, inicio=%d) params=[%s]\n", f.nombre, f.tipoRetorno, f.dirInicio, params)
	}
}
