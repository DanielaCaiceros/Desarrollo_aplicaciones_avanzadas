package main

import (
    "fmt"
    "os"
    "github.com/antlr4-go/antlr/v4"
    "patito/parser"
)

func main() {
    // Leer archivo de entrada
    input, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error leyendo archivo:", err)
        os.Exit(1)
    }

    // Crear el stream de entrada
    is := antlr.NewInputStream(string(input))

    // Crear el lexer
    lexer := parser.NewgramaticaLexer(is)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

    // Crear el parser
    p := parser.NewgramaticaParser(stream)

    // Parsear desde la regla inicial
    tree := p.Programa()

    // Imprimir el árbol
    fmt.Println(tree.ToStringTree(nil, p))
}