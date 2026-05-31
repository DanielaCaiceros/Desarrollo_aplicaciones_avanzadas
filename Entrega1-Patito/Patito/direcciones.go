package main

import (
	"fmt"
	"os"
	"sort"
)

// --- Distribución de Direcciones Virtuales ---
//
// Cada combinación (ámbito, tipo) ocupa un rango de 1000 direcciones.
//
//   Ámbito       entero   flotante  boolean   string
//   ---------    -------  --------  -------   -------
//   global        1000     2000      3000      4000
//   local         5000     6000      7000      8000
//   temporal      9000    10000     11000     12000
//   constante    13000    14000       -       15000
//
// El rango en que cae una dirección identifica el tipo de dato y su alcance,
// lo cual hace eficiente la ejecución (no hay que consultar tablas).

const (
	BaseGlobalEnt = 1000
	BaseGlobalFlo = 2000
	BaseGlobalBoo = 3000
	BaseGlobalStr = 4000

	BaseLocalEnt = 5000
	BaseLocalFlo = 6000
	BaseLocalBoo = 7000
	BaseLocalStr = 8000

	BaseTempEnt = 9000
	BaseTempFlo = 10000
	BaseTempBoo = 11000
	BaseTempStr = 12000

	BaseCteEnt = 13000
	BaseCteFlo = 14000
	BaseCteStr = 15000

	TamanoSegmento = 1000
)

// AdministradorMemoria reparte direcciones virtuales por ámbito y tipo.
type AdministradorMemoria struct {
	contadores map[string]int // clave "ambito_tipo" -> siguiente desplazamiento
	constantes map[string]int // texto del literal -> dirección (evita duplicados)
}

func NuevoAdministradorMemoria() *AdministradorMemoria {
	return &AdministradorMemoria{
		contadores: make(map[string]int),
		constantes: make(map[string]int),
	}
}

func baseDireccion(ambito, tipo string) int {
	switch ambito {
	case "global":
		switch tipo {
		case "entero":
			return BaseGlobalEnt
		case "flotante":
			return BaseGlobalFlo
		case "boolean":
			return BaseGlobalBoo
		case "string":
			return BaseGlobalStr
		}
	case "local":
		switch tipo {
		case "entero":
			return BaseLocalEnt
		case "flotante":
			return BaseLocalFlo
		case "boolean":
			return BaseLocalBoo
		case "string":
			return BaseLocalStr
		}
	case "temporal":
		switch tipo {
		case "entero":
			return BaseTempEnt
		case "flotante":
			return BaseTempFlo
		case "boolean":
			return BaseTempBoo
		case "string":
			return BaseTempStr
		}
	case "constante":
		switch tipo {
		case "entero":
			return BaseCteEnt
		case "flotante":
			return BaseCteFlo
		case "string":
			return BaseCteStr
		}
	}
	return -1
}

// Asignar entrega la siguiente dirección disponible para (ambito, tipo).
func (a *AdministradorMemoria) Asignar(ambito, tipo string) int {
	base := baseDireccion(ambito, tipo)
	if base == -1 {
		fmt.Printf("Error: no existe segmento para %s %s\n", ambito, tipo)
		os.Exit(1)
	}
	clave := ambito + "_" + tipo
	off := a.contadores[clave]
	if off >= TamanoSegmento {
		fmt.Printf("Error: memoria agotada para %s %s\n", ambito, tipo)
		os.Exit(1)
	}
	a.contadores[clave]++
	return base + off
}

// Constante asigna (o reutiliza) la dirección de un literal.
func (a *AdministradorMemoria) Constante(valor, tipo string) int {
	if dir, ok := a.constantes[valor]; ok {
		return dir
	}
	dir := a.Asignar("constante", tipo)
	a.constantes[valor] = dir
	return dir
}

// ResetLocal reinicia los contadores locales y temporales al terminar una función.
func (a *AdministradorMemoria) ResetLocal() {
	for _, tipo := range []string{"entero", "flotante", "boolean", "string"} {
		delete(a.contadores, "local_"+tipo)
		delete(a.contadores, "temporal_"+tipo)
	}
}

// ImprimirConstantes muestra la tabla de constantes con sus direcciones.
func (a *AdministradorMemoria) ImprimirConstantes() {
	fmt.Println("\n=== Tabla de Constantes ===")
	type cte struct {
		dir   int
		valor string
	}
	lista := make([]cte, 0, len(a.constantes))
	for valor, dir := range a.constantes {
		lista = append(lista, cte{dir, valor})
	}
	sort.Slice(lista, func(i, j int) bool { return lista[i].dir < lista[j].dir })
	for _, c := range lista {
		fmt.Printf("  %d -> %s\n", c.dir, c.valor)
	}
}

// ImprimirDistribucion muestra la distribución de direcciones virtuales diseñada.
func (a *AdministradorMemoria) ImprimirDistribucion() {
	fmt.Println("\n=== Distribución de Direcciones Virtuales ===")
	fmt.Printf("  %-12s %-8s %-9s %-8s %-8s\n", "Ambito", "entero", "flotante", "boolean", "string")
	fmt.Printf("  %-12s %-8d %-9d %-8d %-8d\n", "global", BaseGlobalEnt, BaseGlobalFlo, BaseGlobalBoo, BaseGlobalStr)
	fmt.Printf("  %-12s %-8d %-9d %-8d %-8d\n", "local", BaseLocalEnt, BaseLocalFlo, BaseLocalBoo, BaseLocalStr)
	fmt.Printf("  %-12s %-8d %-9d %-8d %-8d\n", "temporal", BaseTempEnt, BaseTempFlo, BaseTempBoo, BaseTempStr)
	fmt.Printf("  %-12s %-8d %-9d %-8s %-8d\n", "constante", BaseCteEnt, BaseCteFlo, "-", BaseCteStr)
}
