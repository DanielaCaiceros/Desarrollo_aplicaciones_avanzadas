package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================================
//  MÁQUINA VIRTUAL DE PATITO
//
//  Interpreta la fila de cuádruplos resolviendo las DIRECCIONES VIRTUALES en
//  tiempo de ejecución. La memoria de ejecución está segmentada igual que el
//  mapa de direcciones virtuales del compilador:
//
//     global / constante  -> memoria compartida (una sola copia)
//     local  / temporal   -> memoria del Registro de Activación en turno
//
//  El rango de una dirección (ver segmentoDe / tipoDeDireccion en direcciones.go)
//  basta para saber EN QUÉ estructura y de QUÉ tipo es el dato; por eso la
//  ejecución no necesita volver a consultar el directorio de funciones ni las
//  tablas de variables.
// ============================================================================

// --- Memoria: bloque de celdas indexado por dirección virtual ---

type Memoria struct {
	celdas map[int]interface{}
}

func NuevaMemoria() *Memoria {
	return &Memoria{celdas: make(map[int]interface{})}
}

func (m *Memoria) Guardar(dir int, valor interface{}) {
	m.celdas[dir] = valor
}

func (m *Memoria) Leer(dir int) interface{} {
	v, ok := m.celdas[dir]
	if !ok {
		fmt.Printf("Error de ejecución: se leyó la dirección %d antes de asignarla\n", dir)
		os.Exit(1)
	}
	return v
}

// --- Registro de Activación: memoria local+temporal de una llamada ---

type RegistroActivacion struct {
	memoria *Memoria // segmentos local y temporal de esta invocación
	regreso int      // índice de cuádruplo al que volver tras ENDPROC
}

func NuevoRegistroActivacion() *RegistroActivacion {
	return &RegistroActivacion{memoria: NuevaMemoria()}
}

// --- Máquina Virtual ---

type MaquinaVirtual struct {
	cuads      []Cuadruplo
	global     *Memoria              // segmento global (compartido)
	constantes *Memoria              // segmento constante (compartido, precargado)
	pila       []*RegistroActivacion // pila de registros de activación
	pendiente  *RegistroActivacion   // RA en construcción entre ERA y GOSUB
	ip         int                   // apuntador de instrucción (instruction pointer)
}

func NuevaMaquinaVirtual(cuads []Cuadruplo, constantes map[string]int) *MaquinaVirtual {
	vm := &MaquinaVirtual{
		cuads:      cuads,
		global:     NuevaMemoria(),
		constantes: NuevaMemoria(),
		// Registro de activación base: aloja los temporales del programa principal.
		pila: []*RegistroActivacion{NuevoRegistroActivacion()},
	}
	vm.cargarConstantes(constantes)
	return vm
}

// cargarConstantes precarga la memoria de constantes parseando cada literal
// al tipo que su dirección virtual indica.
func (vm *MaquinaVirtual) cargarConstantes(constantes map[string]int) {
	for literal, dir := range constantes {
		switch tipoDeDireccion(dir) {
		case "entero":
			n, _ := strconv.Atoi(literal)
			vm.constantes.Guardar(dir, n)
		case "flotante":
			f, _ := strconv.ParseFloat(literal, 64)
			vm.constantes.Guardar(dir, f)
		case "string":
			vm.constantes.Guardar(dir, strings.Trim(literal, "\""))
		}
	}
}

// memoriaDe escoge la estructura de memoria que corresponde a una dirección.
// Aquí es donde la dirección virtual "indexa" la memoria de ejecución.
func (vm *MaquinaVirtual) memoriaDe(dir int) *Memoria {
	switch segmentoDe(dir) {
	case "global":
		return vm.global
	case "constante":
		return vm.constantes
	case "local", "temporal":
		return vm.pila[len(vm.pila)-1].memoria
	}
	fmt.Printf("Error de ejecución: dirección virtual fuera de rango (%d)\n", dir)
	os.Exit(1)
	return nil
}

func (vm *MaquinaVirtual) leer(dirStr string) interface{} {
	dir, _ := strconv.Atoi(dirStr)
	return vm.memoriaDe(dir).Leer(dir)
}

func (vm *MaquinaVirtual) escribir(dirStr string, valor interface{}) {
	dir, _ := strconv.Atoi(dirStr)
	// Coerción entero->flotante al asignar a una celda flotante (widening).
	if tipoDeDireccion(dir) == "flotante" {
		if n, ok := valor.(int); ok {
			valor = float64(n)
		}
	}
	vm.memoriaDe(dir).Guardar(dir, valor)
}

// --- Bucle principal de interpretación ---

func (vm *MaquinaVirtual) Ejecutar() {
	fmt.Println("\n=== Salida del Programa (Máquina Virtual) ===")
	for vm.ip < len(vm.cuads) {
		c := vm.cuads[vm.ip]
		switch c.operador {

		case "GOTO":
			vm.ip = atoiCuad(c.res)
			continue

		case "GotoF":
			if !esVerdadero(vm.leer(c.opizq)) {
				vm.ip = atoiCuad(c.res)
				continue
			}

		case "=":
			vm.escribir(c.res, vm.leer(c.opizq))

		case "+", "-", "*", "/":
			vm.escribir(c.res, vm.aritmetica(c.operador, vm.leer(c.opizq), vm.leer(c.opder)))

		case ">", "<", "==", "!=":
			vm.escribir(c.res, vm.relacional(c.operador, vm.leer(c.opizq), vm.leer(c.opder)))

		case "PRINT":
			fmt.Println(formatear(vm.leer(c.opizq)))

		case "ERA":
			// Reserva el espacio de memoria de la próxima llamada.
			vm.pendiente = NuevoRegistroActivacion()

		case "PARAM":
			// Copia el argumento (contexto actual) al parámetro (RA pendiente).
			valor := vm.leer(c.opizq)
			destino, _ := strconv.Atoi(c.res)
			if tipoDeDireccion(destino) == "flotante" {
				if n, ok := valor.(int); ok {
					valor = float64(n)
				}
			}
			vm.pendiente.memoria.Guardar(destino, valor)

		case "GOSUB":
			vm.pendiente.regreso = vm.ip + 1
			vm.pila = append(vm.pila, vm.pendiente)
			vm.pendiente = nil
			vm.ip = atoiCuad(c.res)
			continue

		case "RETURN":
			// Copia el valor de retorno a la dirección global de la función
			// (leído ANTES de desapilar el registro de activación) y regresa.
			valor := vm.leer(c.opizq)
			destino, _ := strconv.Atoi(c.res)
			if tipoDeDireccion(destino) == "flotante" {
				if n, ok := valor.(int); ok {
					valor = float64(n)
				}
			}
			vm.global.Guardar(destino, valor)
			ra := vm.pila[len(vm.pila)-1]
			vm.pila = vm.pila[:len(vm.pila)-1]
			vm.ip = ra.regreso
			continue

		case "ENDPROC":
			ra := vm.pila[len(vm.pila)-1]
			vm.pila = vm.pila[:len(vm.pila)-1]
			vm.ip = ra.regreso
			continue

		default:
			fmt.Printf("Error de ejecución: operador desconocido %q\n", c.operador)
			os.Exit(1)
		}
		vm.ip++
	}
}

// --- Operaciones aritméticas y relacionales ---

func (vm *MaquinaVirtual) aritmetica(op string, a, b interface{}) interface{} {
	// Concatenación de strings con '+'.
	if as, aok := a.(string); aok {
		if bs, bok := b.(string); bok && op == "+" {
			return as + bs
		}
	}
	if esFlotante(a) || esFlotante(b) {
		x, y := aFlotante(a), aFlotante(b)
		switch op {
		case "+":
			return x + y
		case "-":
			return x - y
		case "*":
			return x * y
		case "/":
			return x / y
		}
	}
	x, y := aEntero(a), aEntero(b)
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		if y == 0 {
			fmt.Println("Error de ejecución: división entre cero")
			os.Exit(1)
		}
		return x / y
	}
	return 0
}

// relacional devuelve 1 (verdadero) o 0 (falso) como entero, igual que el cubo.
func (vm *MaquinaVirtual) relacional(op string, a, b interface{}) int {
	var res bool
	if as, aok := a.(string); aok {
		bs, _ := b.(string)
		switch op {
		case "==":
			res = as == bs
		case "!=":
			res = as != bs
		}
	} else {
		x, y := aFlotante(a), aFlotante(b)
		switch op {
		case ">":
			res = x > y
		case "<":
			res = x < y
		case "==":
			res = x == y
		case "!=":
			res = x != y
		}
	}
	if res {
		return 1
	}
	return 0
}

// --- Utilerías de tipo/valor ---

func atoiCuad(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func esFlotante(v interface{}) bool {
	_, ok := v.(float64)
	return ok
}

func aFlotante(v interface{}) float64 {
	switch x := v.(type) {
	case int:
		return float64(x)
	case float64:
		return x
	}
	return 0
}

func aEntero(v interface{}) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	}
	return 0
}

func esVerdadero(v interface{}) bool {
	switch x := v.(type) {
	case int:
		return x != 0
	case float64:
		return x != 0
	case bool:
		return x
	}
	return false
}

func formatear(v interface{}) string {
	switch x := v.(type) {
	case int:
		return strconv.Itoa(x)
	case float64:
		return strconv.FormatFloat(x, 'g', -1, 64)
	case string:
		return x
	case bool:
		if x {
			return "verdadero"
		}
		return "falso"
	}
	return fmt.Sprintf("%v", v)
}
