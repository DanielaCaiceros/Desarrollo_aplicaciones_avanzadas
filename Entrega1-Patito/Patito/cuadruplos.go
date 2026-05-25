package main
import "fmt"

type Pila struct {
	items []any
}

func Push(p *Pila, elemento any) {
	p.items = append(p.items, elemento)
}

func Pop(p *Pila) any {
	if EstaVacia(p) {
		return nil
	}
	popItem := p.items[len(p.items)-1]
	p.items = p.items[:len(p.items)-1] 
	return popItem
}

func Peek(p *Pila)any {
	if EstaVacia(p) {
		return nil
	}
	return p.items[len(p.items)-1]
}

func EstaVacia(p *Pila) bool {
	return len(p.items) == 0
}

// Cuadruplos

type Cuadruplo struct {
	opizq string
	opder string
	operador string
	res string
}

type Cuadruplos struct {
	cuadruplos []Cuadruplo
}

func (c *Cuadruplos) AgregarCuadruplo(opizq, opder, operador, res string) {
	cuadruplo := Cuadruplo{
		opizq: opizq,
		opder: opder,
		operador: operador,
		res: res,
	}
	c.cuadruplos = append(c.cuadruplos, cuadruplo)
}

func (c *Cuadruplos) ImprimirCuadruplos() {
	for i, cuadruplo := range c.cuadruplos {
		fmt.Printf("%d: (%s, %s, %s, %s)\n", i, cuadruplo.opizq, cuadruplo.opder, cuadruplo.operador, cuadruplo.res)
	}
}