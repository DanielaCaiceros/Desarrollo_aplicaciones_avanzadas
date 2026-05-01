// Code generated from gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // gramatica
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type gramaticaParser struct {
	*antlr.BaseParser
}

var GramaticaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.LiteralNames = []string{
		"", "';'", "','", "':'", "'{'", "'}'", "'('", "')'", "'['", "']'", "'='",
		"'<'", "'>'", "'+'", "'-'", "'*'", "'/'", "'!='", "'=='", "", "", "'programa'",
		"'inicio'", "'fin'", "'vars'", "'entero'", "'flotante'", "'escribe'",
		"'mientras'", "'haz'", "'si'", "'sino'", "'nula'",
	}
	staticData.SymbolicNames = []string{
		"", "SEMICOLON", "COMMA", "DPUNTOS", "LCORCHETE", "RCORCHETE", "LPAR",
		"RPAR", "LBRACKET", "RBRACKET", "ASIG", "MENOR", "MAYOR", "MAS", "MENOS",
		"MULT", "DIV", "NEQ", "EQ", "CTE_FLOAT", "CTE_ENT", "PROGRAMA", "INICIO",
		"FIN", "VARS", "ENTERO", "FLOTANTE", "ESCRIBE", "MIENTRAS", "HAZ", "SI",
		"SINO", "NULA", "LETRERO", "ID", "WS", "COMMENT_LINE", "COMMENT_BLOCK",
	}
	staticData.RuleNames = []string{
		"programa", "varsop", "vars", "funcsop", "cuerpo", "estatutos", "idop",
		"tipo", "estatuto", "imprime", "explet", "letreros", "expresiones",
		"asigna", "ciclo", "condicion", "sinoop", "expresion", "opc", "cte",
		"exp", "exopc", "termino", "teropc", "factor", "funcs", "funcsopc",
		"funcr", "varsdec", "llamada", "llamadaexp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 273, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 78, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 89, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 99, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 105, 8, 6, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 120,
		8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 3, 10, 130, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 136, 8, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 143, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 169, 8,
		16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 3, 18, 183, 8, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 199,
		8, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 213, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 226, 8, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26,
		3, 26, 240, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 3, 27, 252, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28,
		258, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 3, 30, 271, 8, 30, 1, 30, 0, 0, 31, 0, 2, 4, 6, 8, 10,
		12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
		48, 50, 52, 54, 56, 58, 60, 0, 2, 1, 0, 25, 26, 1, 0, 19, 20, 273, 0, 62,
		1, 0, 0, 0, 2, 77, 1, 0, 0, 0, 4, 79, 1, 0, 0, 0, 6, 88, 1, 0, 0, 0, 8,
		90, 1, 0, 0, 0, 10, 98, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0, 14, 106, 1, 0,
		0, 0, 16, 119, 1, 0, 0, 0, 18, 121, 1, 0, 0, 0, 20, 129, 1, 0, 0, 0, 22,
		135, 1, 0, 0, 0, 24, 142, 1, 0, 0, 0, 26, 144, 1, 0, 0, 0, 28, 149, 1,
		0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 168, 1, 0, 0, 0, 34, 170, 1, 0, 0, 0,
		36, 182, 1, 0, 0, 0, 38, 184, 1, 0, 0, 0, 40, 186, 1, 0, 0, 0, 42, 198,
		1, 0, 0, 0, 44, 200, 1, 0, 0, 0, 46, 212, 1, 0, 0, 0, 48, 225, 1, 0, 0,
		0, 50, 227, 1, 0, 0, 0, 52, 239, 1, 0, 0, 0, 54, 251, 1, 0, 0, 0, 56, 257,
		1, 0, 0, 0, 58, 259, 1, 0, 0, 0, 60, 270, 1, 0, 0, 0, 62, 63, 5, 21, 0,
		0, 63, 64, 5, 34, 0, 0, 64, 65, 5, 1, 0, 0, 65, 66, 5, 24, 0, 0, 66, 67,
		5, 3, 0, 0, 67, 68, 3, 2, 1, 0, 68, 69, 3, 6, 3, 0, 69, 70, 5, 22, 0, 0,
		70, 71, 3, 8, 4, 0, 71, 72, 5, 23, 0, 0, 72, 1, 1, 0, 0, 0, 73, 74, 3,
		4, 2, 0, 74, 75, 3, 2, 1, 0, 75, 78, 1, 0, 0, 0, 76, 78, 1, 0, 0, 0, 77,
		73, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 3, 1, 0, 0, 0, 79, 80, 3, 12, 6,
		0, 80, 81, 5, 3, 0, 0, 81, 82, 3, 14, 7, 0, 82, 83, 5, 1, 0, 0, 83, 5,
		1, 0, 0, 0, 84, 85, 3, 50, 25, 0, 85, 86, 3, 6, 3, 0, 86, 89, 1, 0, 0,
		0, 87, 89, 1, 0, 0, 0, 88, 84, 1, 0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 7, 1,
		0, 0, 0, 90, 91, 5, 4, 0, 0, 91, 92, 3, 10, 5, 0, 92, 93, 5, 5, 0, 0, 93,
		9, 1, 0, 0, 0, 94, 95, 3, 16, 8, 0, 95, 96, 3, 10, 5, 0, 96, 99, 1, 0,
		0, 0, 97, 99, 1, 0, 0, 0, 98, 94, 1, 0, 0, 0, 98, 97, 1, 0, 0, 0, 99, 11,
		1, 0, 0, 0, 100, 101, 5, 34, 0, 0, 101, 102, 5, 2, 0, 0, 102, 105, 3, 12,
		6, 0, 103, 105, 5, 34, 0, 0, 104, 100, 1, 0, 0, 0, 104, 103, 1, 0, 0, 0,
		105, 13, 1, 0, 0, 0, 106, 107, 7, 0, 0, 0, 107, 15, 1, 0, 0, 0, 108, 120,
		3, 26, 13, 0, 109, 120, 3, 30, 15, 0, 110, 120, 3, 28, 14, 0, 111, 112,
		3, 58, 29, 0, 112, 113, 5, 1, 0, 0, 113, 120, 1, 0, 0, 0, 114, 120, 3,
		18, 9, 0, 115, 116, 5, 4, 0, 0, 116, 117, 3, 10, 5, 0, 117, 118, 5, 5,
		0, 0, 118, 120, 1, 0, 0, 0, 119, 108, 1, 0, 0, 0, 119, 109, 1, 0, 0, 0,
		119, 110, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 114, 1, 0, 0, 0, 119,
		115, 1, 0, 0, 0, 120, 17, 1, 0, 0, 0, 121, 122, 5, 27, 0, 0, 122, 123,
		5, 6, 0, 0, 123, 124, 3, 20, 10, 0, 124, 125, 5, 7, 0, 0, 125, 126, 5,
		1, 0, 0, 126, 19, 1, 0, 0, 0, 127, 130, 3, 24, 12, 0, 128, 130, 3, 22,
		11, 0, 129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 21, 1, 0, 0, 0,
		131, 132, 5, 33, 0, 0, 132, 133, 5, 2, 0, 0, 133, 136, 3, 22, 11, 0, 134,
		136, 5, 33, 0, 0, 135, 131, 1, 0, 0, 0, 135, 134, 1, 0, 0, 0, 136, 23,
		1, 0, 0, 0, 137, 138, 3, 34, 17, 0, 138, 139, 5, 2, 0, 0, 139, 140, 3,
		24, 12, 0, 140, 143, 1, 0, 0, 0, 141, 143, 3, 34, 17, 0, 142, 137, 1, 0,
		0, 0, 142, 141, 1, 0, 0, 0, 143, 25, 1, 0, 0, 0, 144, 145, 5, 34, 0, 0,
		145, 146, 5, 10, 0, 0, 146, 147, 3, 34, 17, 0, 147, 148, 5, 1, 0, 0, 148,
		27, 1, 0, 0, 0, 149, 150, 5, 28, 0, 0, 150, 151, 5, 6, 0, 0, 151, 152,
		3, 34, 17, 0, 152, 153, 5, 7, 0, 0, 153, 154, 5, 29, 0, 0, 154, 155, 3,
		8, 4, 0, 155, 156, 5, 1, 0, 0, 156, 29, 1, 0, 0, 0, 157, 158, 5, 30, 0,
		0, 158, 159, 5, 6, 0, 0, 159, 160, 3, 34, 17, 0, 160, 161, 5, 7, 0, 0,
		161, 162, 3, 8, 4, 0, 162, 163, 3, 32, 16, 0, 163, 164, 5, 1, 0, 0, 164,
		31, 1, 0, 0, 0, 165, 166, 5, 31, 0, 0, 166, 169, 3, 8, 4, 0, 167, 169,
		1, 0, 0, 0, 168, 165, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 33, 1, 0,
		0, 0, 170, 171, 3, 40, 20, 0, 171, 172, 3, 36, 18, 0, 172, 35, 1, 0, 0,
		0, 173, 174, 5, 12, 0, 0, 174, 183, 3, 40, 20, 0, 175, 176, 5, 11, 0, 0,
		176, 183, 3, 40, 20, 0, 177, 178, 5, 17, 0, 0, 178, 183, 3, 40, 20, 0,
		179, 180, 5, 18, 0, 0, 180, 183, 3, 40, 20, 0, 181, 183, 1, 0, 0, 0, 182,
		173, 1, 0, 0, 0, 182, 175, 1, 0, 0, 0, 182, 177, 1, 0, 0, 0, 182, 179,
		1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 37, 1, 0, 0, 0, 184, 185, 7, 1,
		0, 0, 185, 39, 1, 0, 0, 0, 186, 187, 3, 44, 22, 0, 187, 188, 3, 42, 21,
		0, 188, 41, 1, 0, 0, 0, 189, 190, 5, 13, 0, 0, 190, 191, 3, 44, 22, 0,
		191, 192, 3, 42, 21, 0, 192, 199, 1, 0, 0, 0, 193, 194, 5, 14, 0, 0, 194,
		195, 3, 44, 22, 0, 195, 196, 3, 42, 21, 0, 196, 199, 1, 0, 0, 0, 197, 199,
		1, 0, 0, 0, 198, 189, 1, 0, 0, 0, 198, 193, 1, 0, 0, 0, 198, 197, 1, 0,
		0, 0, 199, 43, 1, 0, 0, 0, 200, 201, 3, 48, 24, 0, 201, 202, 3, 46, 23,
		0, 202, 45, 1, 0, 0, 0, 203, 204, 5, 15, 0, 0, 204, 205, 3, 48, 24, 0,
		205, 206, 3, 46, 23, 0, 206, 213, 1, 0, 0, 0, 207, 208, 5, 16, 0, 0, 208,
		209, 3, 48, 24, 0, 209, 210, 3, 46, 23, 0, 210, 213, 1, 0, 0, 0, 211, 213,
		1, 0, 0, 0, 212, 203, 1, 0, 0, 0, 212, 207, 1, 0, 0, 0, 212, 211, 1, 0,
		0, 0, 213, 47, 1, 0, 0, 0, 214, 226, 5, 34, 0, 0, 215, 226, 3, 38, 19,
		0, 216, 217, 5, 6, 0, 0, 217, 218, 3, 34, 17, 0, 218, 219, 5, 7, 0, 0,
		219, 226, 1, 0, 0, 0, 220, 221, 5, 13, 0, 0, 221, 226, 3, 48, 24, 0, 222,
		223, 5, 14, 0, 0, 223, 226, 3, 48, 24, 0, 224, 226, 3, 58, 29, 0, 225,
		214, 1, 0, 0, 0, 225, 215, 1, 0, 0, 0, 225, 216, 1, 0, 0, 0, 225, 220,
		1, 0, 0, 0, 225, 222, 1, 0, 0, 0, 225, 224, 1, 0, 0, 0, 226, 49, 1, 0,
		0, 0, 227, 228, 3, 52, 26, 0, 228, 229, 5, 34, 0, 0, 229, 230, 5, 6, 0,
		0, 230, 231, 3, 54, 27, 0, 231, 232, 5, 7, 0, 0, 232, 233, 5, 4, 0, 0,
		233, 234, 3, 56, 28, 0, 234, 235, 3, 8, 4, 0, 235, 236, 5, 5, 0, 0, 236,
		51, 1, 0, 0, 0, 237, 240, 5, 32, 0, 0, 238, 240, 3, 14, 7, 0, 239, 237,
		1, 0, 0, 0, 239, 238, 1, 0, 0, 0, 240, 53, 1, 0, 0, 0, 241, 242, 5, 34,
		0, 0, 242, 243, 5, 3, 0, 0, 243, 244, 3, 14, 7, 0, 244, 245, 5, 2, 0, 0,
		245, 246, 3, 54, 27, 0, 246, 252, 1, 0, 0, 0, 247, 248, 5, 34, 0, 0, 248,
		249, 5, 3, 0, 0, 249, 252, 3, 14, 7, 0, 250, 252, 1, 0, 0, 0, 251, 241,
		1, 0, 0, 0, 251, 247, 1, 0, 0, 0, 251, 250, 1, 0, 0, 0, 252, 55, 1, 0,
		0, 0, 253, 254, 5, 24, 0, 0, 254, 255, 5, 3, 0, 0, 255, 258, 3, 2, 1, 0,
		256, 258, 1, 0, 0, 0, 257, 253, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258,
		57, 1, 0, 0, 0, 259, 260, 5, 34, 0, 0, 260, 261, 5, 6, 0, 0, 261, 262,
		3, 60, 30, 0, 262, 263, 5, 7, 0, 0, 263, 59, 1, 0, 0, 0, 264, 265, 3, 34,
		17, 0, 265, 266, 5, 2, 0, 0, 266, 267, 3, 60, 30, 0, 267, 271, 1, 0, 0,
		0, 268, 271, 3, 34, 17, 0, 269, 271, 1, 0, 0, 0, 270, 264, 1, 0, 0, 0,
		270, 268, 1, 0, 0, 0, 270, 269, 1, 0, 0, 0, 271, 61, 1, 0, 0, 0, 17, 77,
		88, 98, 104, 119, 129, 135, 142, 168, 182, 198, 212, 225, 239, 251, 257,
		270,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// gramaticaParserInit initializes any static state used to implement gramaticaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewgramaticaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaParserInit() {
	staticData := &GramaticaParserStaticData
	staticData.once.Do(gramaticaParserInit)
}

// NewgramaticaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewgramaticaParser(input antlr.TokenStream) *gramaticaParser {
	GramaticaParserInit()
	this := new(gramaticaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GramaticaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "gramatica.g4"

	return this
}

// gramaticaParser tokens.
const (
	gramaticaParserEOF           = antlr.TokenEOF
	gramaticaParserSEMICOLON     = 1
	gramaticaParserCOMMA         = 2
	gramaticaParserDPUNTOS       = 3
	gramaticaParserLCORCHETE     = 4
	gramaticaParserRCORCHETE     = 5
	gramaticaParserLPAR          = 6
	gramaticaParserRPAR          = 7
	gramaticaParserLBRACKET      = 8
	gramaticaParserRBRACKET      = 9
	gramaticaParserASIG          = 10
	gramaticaParserMENOR         = 11
	gramaticaParserMAYOR         = 12
	gramaticaParserMAS           = 13
	gramaticaParserMENOS         = 14
	gramaticaParserMULT          = 15
	gramaticaParserDIV           = 16
	gramaticaParserNEQ           = 17
	gramaticaParserEQ            = 18
	gramaticaParserCTE_FLOAT     = 19
	gramaticaParserCTE_ENT       = 20
	gramaticaParserPROGRAMA      = 21
	gramaticaParserINICIO        = 22
	gramaticaParserFIN           = 23
	gramaticaParserVARS          = 24
	gramaticaParserENTERO        = 25
	gramaticaParserFLOTANTE      = 26
	gramaticaParserESCRIBE       = 27
	gramaticaParserMIENTRAS      = 28
	gramaticaParserHAZ           = 29
	gramaticaParserSI            = 30
	gramaticaParserSINO          = 31
	gramaticaParserNULA          = 32
	gramaticaParserLETRERO       = 33
	gramaticaParserID            = 34
	gramaticaParserWS            = 35
	gramaticaParserCOMMENT_LINE  = 36
	gramaticaParserCOMMENT_BLOCK = 37
)

// gramaticaParser rules.
const (
	gramaticaParserRULE_programa    = 0
	gramaticaParserRULE_varsop      = 1
	gramaticaParserRULE_vars        = 2
	gramaticaParserRULE_funcsop     = 3
	gramaticaParserRULE_cuerpo      = 4
	gramaticaParserRULE_estatutos   = 5
	gramaticaParserRULE_idop        = 6
	gramaticaParserRULE_tipo        = 7
	gramaticaParserRULE_estatuto    = 8
	gramaticaParserRULE_imprime     = 9
	gramaticaParserRULE_explet      = 10
	gramaticaParserRULE_letreros    = 11
	gramaticaParserRULE_expresiones = 12
	gramaticaParserRULE_asigna      = 13
	gramaticaParserRULE_ciclo       = 14
	gramaticaParserRULE_condicion   = 15
	gramaticaParserRULE_sinoop      = 16
	gramaticaParserRULE_expresion   = 17
	gramaticaParserRULE_opc         = 18
	gramaticaParserRULE_cte         = 19
	gramaticaParserRULE_exp         = 20
	gramaticaParserRULE_exopc       = 21
	gramaticaParserRULE_termino     = 22
	gramaticaParserRULE_teropc      = 23
	gramaticaParserRULE_factor      = 24
	gramaticaParserRULE_funcs       = 25
	gramaticaParserRULE_funcsopc    = 26
	gramaticaParserRULE_funcr       = 27
	gramaticaParserRULE_varsdec     = 28
	gramaticaParserRULE_llamada     = 29
	gramaticaParserRULE_llamadaexp  = 30
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROGRAMA() antlr.TerminalNode
	ID() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	VARS() antlr.TerminalNode
	DPUNTOS() antlr.TerminalNode
	Varsop() IVarsopContext
	Funcsop() IFuncsopContext
	INICIO() antlr.TerminalNode
	Cuerpo() ICuerpoContext
	FIN() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) PROGRAMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserPROGRAMA, 0)
}

func (s *ProgramaContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *ProgramaContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *ProgramaContext) VARS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserVARS, 0)
}

func (s *ProgramaContext) DPUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDPUNTOS, 0)
}

func (s *ProgramaContext) Varsop() IVarsopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsopContext)
}

func (s *ProgramaContext) Funcsop() IFuncsopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsopContext)
}

func (s *ProgramaContext) INICIO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserINICIO, 0)
}

func (s *ProgramaContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *ProgramaContext) FIN() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFIN, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (p *gramaticaParser) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gramaticaParserRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(gramaticaParserPROGRAMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		p.Match(gramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(gramaticaParserVARS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(gramaticaParserDPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Varsop()
	}
	{
		p.SetState(68)
		p.Funcsop()
	}
	{
		p.SetState(69)
		p.Match(gramaticaParserINICIO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Cuerpo()
	}
	{
		p.SetState(71)
		p.Match(gramaticaParserFIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsopContext is an interface to support dynamic dispatch.
type IVarsopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Vars() IVarsContext
	Varsop() IVarsopContext

	// IsVarsopContext differentiates from other interfaces.
	IsVarsopContext()
}

type VarsopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsopContext() *VarsopContext {
	var p = new(VarsopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varsop
	return p
}

func InitEmptyVarsopContext(p *VarsopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varsop
}

func (*VarsopContext) IsVarsopContext() {}

func NewVarsopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsopContext {
	var p = new(VarsopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varsop

	return p
}

func (s *VarsopContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsopContext) Vars() IVarsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *VarsopContext) Varsop() IVarsopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsopContext)
}

func (s *VarsopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarsop(s)
	}
}

func (s *VarsopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarsop(s)
	}
}

func (p *gramaticaParser) Varsop() (localctx IVarsopContext) {
	localctx = NewVarsopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gramaticaParserRULE_varsop)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Vars()
		}
		{
			p.SetState(74)
			p.Varsop()
		}

	case gramaticaParserLCORCHETE, gramaticaParserINICIO, gramaticaParserENTERO, gramaticaParserFLOTANTE, gramaticaParserNULA:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Idop() IIdopContext
	DPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	SEMICOLON() antlr.TerminalNode

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_vars
	return p
}

func InitEmptyVarsContext(p *VarsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_vars
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) Idop() IIdopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdopContext)
}

func (s *VarsContext) DPUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDPUNTOS, 0)
}

func (s *VarsContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *VarsContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *gramaticaParser) Vars() (localctx IVarsContext) {
	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gramaticaParserRULE_vars)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Idop()
	}
	{
		p.SetState(80)
		p.Match(gramaticaParserDPUNTOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Tipo()
	}
	{
		p.SetState(82)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncsopContext is an interface to support dynamic dispatch.
type IFuncsopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcs() IFuncsContext
	Funcsop() IFuncsopContext

	// IsFuncsopContext differentiates from other interfaces.
	IsFuncsopContext()
}

type FuncsopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsopContext() *FuncsopContext {
	var p = new(FuncsopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcsop
	return p
}

func InitEmptyFuncsopContext(p *FuncsopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcsop
}

func (*FuncsopContext) IsFuncsopContext() {}

func NewFuncsopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsopContext {
	var p = new(FuncsopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcsop

	return p
}

func (s *FuncsopContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsopContext) Funcs() IFuncsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsContext)
}

func (s *FuncsopContext) Funcsop() IFuncsopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsopContext)
}

func (s *FuncsopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncsop(s)
	}
}

func (s *FuncsopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncsop(s)
	}
}

func (p *gramaticaParser) Funcsop() (localctx IFuncsopContext) {
	localctx = NewFuncsopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gramaticaParserRULE_funcsop)
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserENTERO, gramaticaParserFLOTANTE, gramaticaParserNULA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Funcs()
		}
		{
			p.SetState(85)
			p.Funcsop()
		}

	case gramaticaParserINICIO:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICuerpoContext is an interface to support dynamic dispatch.
type ICuerpoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LCORCHETE() antlr.TerminalNode
	Estatutos() IEstatutosContext
	RCORCHETE() antlr.TerminalNode

	// IsCuerpoContext differentiates from other interfaces.
	IsCuerpoContext()
}

type CuerpoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCuerpoContext() *CuerpoContext {
	var p = new(CuerpoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cuerpo
	return p
}

func InitEmptyCuerpoContext(p *CuerpoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cuerpo
}

func (*CuerpoContext) IsCuerpoContext() {}

func NewCuerpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CuerpoContext {
	var p = new(CuerpoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_cuerpo

	return p
}

func (s *CuerpoContext) GetParser() antlr.Parser { return s.parser }

func (s *CuerpoContext) LCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLCORCHETE, 0)
}

func (s *CuerpoContext) Estatutos() IEstatutosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstatutosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstatutosContext)
}

func (s *CuerpoContext) RCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRCORCHETE, 0)
}

func (s *CuerpoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CuerpoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CuerpoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCuerpo(s)
	}
}

func (s *CuerpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCuerpo(s)
	}
}

func (p *gramaticaParser) Cuerpo() (localctx ICuerpoContext) {
	localctx = NewCuerpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gramaticaParserRULE_cuerpo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(gramaticaParserLCORCHETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.Estatutos()
	}
	{
		p.SetState(92)
		p.Match(gramaticaParserRCORCHETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEstatutosContext is an interface to support dynamic dispatch.
type IEstatutosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Estatuto() IEstatutoContext
	Estatutos() IEstatutosContext

	// IsEstatutosContext differentiates from other interfaces.
	IsEstatutosContext()
}

type EstatutosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstatutosContext() *EstatutosContext {
	var p = new(EstatutosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_estatutos
	return p
}

func InitEmptyEstatutosContext(p *EstatutosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_estatutos
}

func (*EstatutosContext) IsEstatutosContext() {}

func NewEstatutosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstatutosContext {
	var p = new(EstatutosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_estatutos

	return p
}

func (s *EstatutosContext) GetParser() antlr.Parser { return s.parser }

func (s *EstatutosContext) Estatuto() IEstatutoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstatutoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstatutoContext)
}

func (s *EstatutosContext) Estatutos() IEstatutosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstatutosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstatutosContext)
}

func (s *EstatutosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstatutosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstatutosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterEstatutos(s)
	}
}

func (s *EstatutosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitEstatutos(s)
	}
}

func (p *gramaticaParser) Estatutos() (localctx IEstatutosContext) {
	localctx = NewEstatutosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gramaticaParserRULE_estatutos)
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserLCORCHETE, gramaticaParserESCRIBE, gramaticaParserMIENTRAS, gramaticaParserSI, gramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Estatuto()
		}
		{
			p.SetState(95)
			p.Estatutos()
		}

	case gramaticaParserRCORCHETE:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdopContext is an interface to support dynamic dispatch.
type IIdopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Idop() IIdopContext

	// IsIdopContext differentiates from other interfaces.
	IsIdopContext()
}

type IdopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdopContext() *IdopContext {
	var p = new(IdopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_idop
	return p
}

func InitEmptyIdopContext(p *IdopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_idop
}

func (*IdopContext) IsIdopContext() {}

func NewIdopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdopContext {
	var p = new(IdopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_idop

	return p
}

func (s *IdopContext) GetParser() antlr.Parser { return s.parser }

func (s *IdopContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *IdopContext) COMMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMMA, 0)
}

func (s *IdopContext) Idop() IIdopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdopContext)
}

func (s *IdopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterIdop(s)
	}
}

func (s *IdopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitIdop(s)
	}
}

func (p *gramaticaParser) Idop() (localctx IIdopContext) {
	localctx = NewIdopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gramaticaParserRULE_idop)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(gramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(gramaticaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Idop()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(gramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENTERO() antlr.TerminalNode
	FLOTANTE() antlr.TerminalNode

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserENTERO, 0)
}

func (s *TipoContext) FLOTANTE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserFLOTANTE, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *gramaticaParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gramaticaParserRULE_tipo)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gramaticaParserENTERO || _la == gramaticaParserFLOTANTE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEstatutoContext is an interface to support dynamic dispatch.
type IEstatutoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Asigna() IAsignaContext
	Condicion() ICondicionContext
	Ciclo() ICicloContext
	Llamada() ILlamadaContext
	SEMICOLON() antlr.TerminalNode
	Imprime() IImprimeContext
	LCORCHETE() antlr.TerminalNode
	Estatutos() IEstatutosContext
	RCORCHETE() antlr.TerminalNode

	// IsEstatutoContext differentiates from other interfaces.
	IsEstatutoContext()
}

type EstatutoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEstatutoContext() *EstatutoContext {
	var p = new(EstatutoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_estatuto
	return p
}

func InitEmptyEstatutoContext(p *EstatutoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_estatuto
}

func (*EstatutoContext) IsEstatutoContext() {}

func NewEstatutoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EstatutoContext {
	var p = new(EstatutoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_estatuto

	return p
}

func (s *EstatutoContext) GetParser() antlr.Parser { return s.parser }

func (s *EstatutoContext) Asigna() IAsignaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignaContext)
}

func (s *EstatutoContext) Condicion() ICondicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *EstatutoContext) Ciclo() ICicloContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICicloContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICicloContext)
}

func (s *EstatutoContext) Llamada() ILlamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
}

func (s *EstatutoContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *EstatutoContext) Imprime() IImprimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImprimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImprimeContext)
}

func (s *EstatutoContext) LCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLCORCHETE, 0)
}

func (s *EstatutoContext) Estatutos() IEstatutosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEstatutosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEstatutosContext)
}

func (s *EstatutoContext) RCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRCORCHETE, 0)
}

func (s *EstatutoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EstatutoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EstatutoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterEstatuto(s)
	}
}

func (s *EstatutoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitEstatuto(s)
	}
}

func (p *gramaticaParser) Estatuto() (localctx IEstatutoContext) {
	localctx = NewEstatutoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gramaticaParserRULE_estatuto)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Asigna()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Condicion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Ciclo()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Llamada()
		}
		{
			p.SetState(112)
			p.Match(gramaticaParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(114)
			p.Imprime()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.Match(gramaticaParserLCORCHETE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Estatutos()
		}
		{
			p.SetState(117)
			p.Match(gramaticaParserRCORCHETE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IImprimeContext is an interface to support dynamic dispatch.
type IImprimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ESCRIBE() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Explet() IExpletContext
	RPAR() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsImprimeContext differentiates from other interfaces.
	IsImprimeContext()
}

type ImprimeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImprimeContext() *ImprimeContext {
	var p = new(ImprimeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_imprime
	return p
}

func InitEmptyImprimeContext(p *ImprimeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_imprime
}

func (*ImprimeContext) IsImprimeContext() {}

func NewImprimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprimeContext {
	var p = new(ImprimeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_imprime

	return p
}

func (s *ImprimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprimeContext) ESCRIBE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserESCRIBE, 0)
}

func (s *ImprimeContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *ImprimeContext) Explet() IExpletContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpletContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpletContext)
}

func (s *ImprimeContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *ImprimeContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *ImprimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImprimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterImprime(s)
	}
}

func (s *ImprimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitImprime(s)
	}
}

func (p *gramaticaParser) Imprime() (localctx IImprimeContext) {
	localctx = NewImprimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, gramaticaParserRULE_imprime)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(121)
		p.Match(gramaticaParserESCRIBE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(122)
		p.Match(gramaticaParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.Explet()
	}
	{
		p.SetState(124)
		p.Match(gramaticaParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(125)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpletContext is an interface to support dynamic dispatch.
type IExpletContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresiones() IExpresionesContext
	Letreros() ILetrerosContext

	// IsExpletContext differentiates from other interfaces.
	IsExpletContext()
}

type ExpletContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpletContext() *ExpletContext {
	var p = new(ExpletContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_explet
	return p
}

func InitEmptyExpletContext(p *ExpletContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_explet
}

func (*ExpletContext) IsExpletContext() {}

func NewExpletContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpletContext {
	var p = new(ExpletContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_explet

	return p
}

func (s *ExpletContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpletContext) Expresiones() IExpresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionesContext)
}

func (s *ExpletContext) Letreros() ILetrerosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetrerosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetrerosContext)
}

func (s *ExpletContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpletContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpletContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExplet(s)
	}
}

func (s *ExpletContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExplet(s)
	}
}

func (p *gramaticaParser) Explet() (localctx IExpletContext) {
	localctx = NewExpletContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, gramaticaParserRULE_explet)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserLPAR, gramaticaParserMAS, gramaticaParserMENOS, gramaticaParserCTE_FLOAT, gramaticaParserCTE_ENT, gramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(127)
			p.Expresiones()
		}

	case gramaticaParserLETRERO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Letreros()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILetrerosContext is an interface to support dynamic dispatch.
type ILetrerosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LETRERO() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	Letreros() ILetrerosContext

	// IsLetrerosContext differentiates from other interfaces.
	IsLetrerosContext()
}

type LetrerosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetrerosContext() *LetrerosContext {
	var p = new(LetrerosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_letreros
	return p
}

func InitEmptyLetrerosContext(p *LetrerosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_letreros
}

func (*LetrerosContext) IsLetrerosContext() {}

func NewLetrerosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetrerosContext {
	var p = new(LetrerosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_letreros

	return p
}

func (s *LetrerosContext) GetParser() antlr.Parser { return s.parser }

func (s *LetrerosContext) LETRERO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLETRERO, 0)
}

func (s *LetrerosContext) COMMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMMA, 0)
}

func (s *LetrerosContext) Letreros() ILetrerosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILetrerosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILetrerosContext)
}

func (s *LetrerosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetrerosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetrerosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLetreros(s)
	}
}

func (s *LetrerosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLetreros(s)
	}
}

func (p *gramaticaParser) Letreros() (localctx ILetrerosContext) {
	localctx = NewLetrerosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gramaticaParserRULE_letreros)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(gramaticaParserLETRERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(gramaticaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Letreros()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Match(gramaticaParserLETRERO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionesContext is an interface to support dynamic dispatch.
type IExpresionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	COMMA() antlr.TerminalNode
	Expresiones() IExpresionesContext

	// IsExpresionesContext differentiates from other interfaces.
	IsExpresionesContext()
}

type ExpresionesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionesContext() *ExpresionesContext {
	var p = new(ExpresionesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresiones
	return p
}

func InitEmptyExpresionesContext(p *ExpresionesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresiones
}

func (*ExpresionesContext) IsExpresionesContext() {}

func NewExpresionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionesContext {
	var p = new(ExpresionesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresiones

	return p
}

func (s *ExpresionesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionesContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMMA, 0)
}

func (s *ExpresionesContext) Expresiones() IExpresionesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionesContext)
}

func (s *ExpresionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpresiones(s)
	}
}

func (s *ExpresionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpresiones(s)
	}
}

func (p *gramaticaParser) Expresiones() (localctx IExpresionesContext) {
	localctx = NewExpresionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gramaticaParserRULE_expresiones)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Expresion()
		}
		{
			p.SetState(138)
			p.Match(gramaticaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Expresiones()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Expresion()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignaContext is an interface to support dynamic dispatch.
type IAsignaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ASIG() antlr.TerminalNode
	Expresion() IExpresionContext
	SEMICOLON() antlr.TerminalNode

	// IsAsignaContext differentiates from other interfaces.
	IsAsignaContext()
}

type AsignaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignaContext() *AsignaContext {
	var p = new(AsignaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigna
	return p
}

func InitEmptyAsignaContext(p *AsignaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_asigna
}

func (*AsignaContext) IsAsignaContext() {}

func NewAsignaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignaContext {
	var p = new(AsignaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_asigna

	return p
}

func (s *AsignaContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignaContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *AsignaContext) ASIG() antlr.TerminalNode {
	return s.GetToken(gramaticaParserASIG, 0)
}

func (s *AsignaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignaContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *AsignaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterAsigna(s)
	}
}

func (s *AsignaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitAsigna(s)
	}
}

func (p *gramaticaParser) Asigna() (localctx IAsignaContext) {
	localctx = NewAsignaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gramaticaParserRULE_asigna)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(gramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(gramaticaParserASIG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Expresion()
	}
	{
		p.SetState(147)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICicloContext is an interface to support dynamic dispatch.
type ICicloContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MIENTRAS() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Expresion() IExpresionContext
	RPAR() antlr.TerminalNode
	HAZ() antlr.TerminalNode
	Cuerpo() ICuerpoContext
	SEMICOLON() antlr.TerminalNode

	// IsCicloContext differentiates from other interfaces.
	IsCicloContext()
}

type CicloContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCicloContext() *CicloContext {
	var p = new(CicloContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_ciclo
	return p
}

func InitEmptyCicloContext(p *CicloContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_ciclo
}

func (*CicloContext) IsCicloContext() {}

func NewCicloContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CicloContext {
	var p = new(CicloContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_ciclo

	return p
}

func (s *CicloContext) GetParser() antlr.Parser { return s.parser }

func (s *CicloContext) MIENTRAS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMIENTRAS, 0)
}

func (s *CicloContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *CicloContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CicloContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *CicloContext) HAZ() antlr.TerminalNode {
	return s.GetToken(gramaticaParserHAZ, 0)
}

func (s *CicloContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *CicloContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *CicloContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CicloContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CicloContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCiclo(s)
	}
}

func (s *CicloContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCiclo(s)
	}
}

func (p *gramaticaParser) Ciclo() (localctx ICicloContext) {
	localctx = NewCicloContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gramaticaParserRULE_ciclo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(gramaticaParserMIENTRAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(gramaticaParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Expresion()
	}
	{
		p.SetState(152)
		p.Match(gramaticaParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(gramaticaParserHAZ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Cuerpo()
	}
	{
		p.SetState(155)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICondicionContext is an interface to support dynamic dispatch.
type ICondicionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SI() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Expresion() IExpresionContext
	RPAR() antlr.TerminalNode
	Cuerpo() ICuerpoContext
	Sinoop() ISinoopContext
	SEMICOLON() antlr.TerminalNode

	// IsCondicionContext differentiates from other interfaces.
	IsCondicionContext()
}

type CondicionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondicionContext() *CondicionContext {
	var p = new(CondicionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_condicion
	return p
}

func InitEmptyCondicionContext(p *CondicionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_condicion
}

func (*CondicionContext) IsCondicionContext() {}

func NewCondicionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondicionContext {
	var p = new(CondicionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_condicion

	return p
}

func (s *CondicionContext) GetParser() antlr.Parser { return s.parser }

func (s *CondicionContext) SI() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSI, 0)
}

func (s *CondicionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *CondicionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *CondicionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *CondicionContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *CondicionContext) Sinoop() ISinoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISinoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISinoopContext)
}

func (s *CondicionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSEMICOLON, 0)
}

func (s *CondicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondicionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCondicion(s)
	}
}

func (s *CondicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCondicion(s)
	}
}

func (p *gramaticaParser) Condicion() (localctx ICondicionContext) {
	localctx = NewCondicionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gramaticaParserRULE_condicion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(gramaticaParserSI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Match(gramaticaParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Expresion()
	}
	{
		p.SetState(160)
		p.Match(gramaticaParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Cuerpo()
	}
	{
		p.SetState(162)
		p.Sinoop()
	}
	{
		p.SetState(163)
		p.Match(gramaticaParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISinoopContext is an interface to support dynamic dispatch.
type ISinoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SINO() antlr.TerminalNode
	Cuerpo() ICuerpoContext

	// IsSinoopContext differentiates from other interfaces.
	IsSinoopContext()
}

type SinoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinoopContext() *SinoopContext {
	var p = new(SinoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sinoop
	return p
}

func InitEmptySinoopContext(p *SinoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_sinoop
}

func (*SinoopContext) IsSinoopContext() {}

func NewSinoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinoopContext {
	var p = new(SinoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_sinoop

	return p
}

func (s *SinoopContext) GetParser() antlr.Parser { return s.parser }

func (s *SinoopContext) SINO() antlr.TerminalNode {
	return s.GetToken(gramaticaParserSINO, 0)
}

func (s *SinoopContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *SinoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterSinoop(s)
	}
}

func (s *SinoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitSinoop(s)
	}
}

func (p *gramaticaParser) Sinoop() (localctx ISinoopContext) {
	localctx = NewSinoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gramaticaParserRULE_sinoop)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserSINO:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(gramaticaParserSINO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Cuerpo()
		}

	case gramaticaParserSEMICOLON:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	Opc() IOpcContext

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpresionContext) Opc() IOpcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpcContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *gramaticaParser) Expresion() (localctx IExpresionContext) {
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gramaticaParserRULE_expresion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Exp()
	}
	{
		p.SetState(171)
		p.Opc()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpcContext is an interface to support dynamic dispatch.
type IOpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAYOR() antlr.TerminalNode
	Exp() IExpContext
	MENOR() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	EQ() antlr.TerminalNode

	// IsOpcContext differentiates from other interfaces.
	IsOpcContext()
}

type OpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpcContext() *OpcContext {
	var p = new(OpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_opc
	return p
}

func InitEmptyOpcContext(p *OpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_opc
}

func (*OpcContext) IsOpcContext() {}

func NewOpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpcContext {
	var p = new(OpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_opc

	return p
}

func (s *OpcContext) GetParser() antlr.Parser { return s.parser }

func (s *OpcContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAYOR, 0)
}

func (s *OpcContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *OpcContext) MENOR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENOR, 0)
}

func (s *OpcContext) NEQ() antlr.TerminalNode {
	return s.GetToken(gramaticaParserNEQ, 0)
}

func (s *OpcContext) EQ() antlr.TerminalNode {
	return s.GetToken(gramaticaParserEQ, 0)
}

func (s *OpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterOpc(s)
	}
}

func (s *OpcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitOpc(s)
	}
}

func (p *gramaticaParser) Opc() (localctx IOpcContext) {
	localctx = NewOpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gramaticaParserRULE_opc)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserMAYOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(gramaticaParserMAYOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Exp()
		}

	case gramaticaParserMENOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(gramaticaParserMENOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Exp()
		}

	case gramaticaParserNEQ:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Match(gramaticaParserNEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Exp()
		}

	case gramaticaParserEQ:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(179)
			p.Match(gramaticaParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Exp()
		}

	case gramaticaParserSEMICOLON, gramaticaParserCOMMA, gramaticaParserRPAR:
		p.EnterOuterAlt(localctx, 5)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICteContext is an interface to support dynamic dispatch.
type ICteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CTE_ENT() antlr.TerminalNode
	CTE_FLOAT() antlr.TerminalNode

	// IsCteContext differentiates from other interfaces.
	IsCteContext()
}

type CteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCteContext() *CteContext {
	var p = new(CteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cte
	return p
}

func InitEmptyCteContext(p *CteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_cte
}

func (*CteContext) IsCteContext() {}

func NewCteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CteContext {
	var p = new(CteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_cte

	return p
}

func (s *CteContext) GetParser() antlr.Parser { return s.parser }

func (s *CteContext) CTE_ENT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCTE_ENT, 0)
}

func (s *CteContext) CTE_FLOAT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCTE_FLOAT, 0)
}

func (s *CteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterCte(s)
	}
}

func (s *CteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitCte(s)
	}
}

func (p *gramaticaParser) Cte() (localctx ICteContext) {
	localctx = NewCteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gramaticaParserRULE_cte)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gramaticaParserCTE_FLOAT || _la == gramaticaParserCTE_ENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Termino() ITerminoContext
	Exopc() IExopcContext

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) Termino() ITerminoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminoContext)
}

func (s *ExpContext) Exopc() IExopcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExopcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExopcContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *gramaticaParser) Exp() (localctx IExpContext) {
	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gramaticaParserRULE_exp)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Termino()
	}
	{
		p.SetState(187)
		p.Exopc()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExopcContext is an interface to support dynamic dispatch.
type IExopcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAS() antlr.TerminalNode
	Termino() ITerminoContext
	Exopc() IExopcContext
	MENOS() antlr.TerminalNode

	// IsExopcContext differentiates from other interfaces.
	IsExopcContext()
}

type ExopcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExopcContext() *ExopcContext {
	var p = new(ExopcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_exopc
	return p
}

func InitEmptyExopcContext(p *ExopcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_exopc
}

func (*ExopcContext) IsExopcContext() {}

func NewExopcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExopcContext {
	var p = new(ExopcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_exopc

	return p
}

func (s *ExopcContext) GetParser() antlr.Parser { return s.parser }

func (s *ExopcContext) MAS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAS, 0)
}

func (s *ExopcContext) Termino() ITerminoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerminoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerminoContext)
}

func (s *ExopcContext) Exopc() IExopcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExopcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExopcContext)
}

func (s *ExopcContext) MENOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENOS, 0)
}

func (s *ExopcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExopcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExopcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterExopc(s)
	}
}

func (s *ExopcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitExopc(s)
	}
}

func (p *gramaticaParser) Exopc() (localctx IExopcContext) {
	localctx = NewExopcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gramaticaParserRULE_exopc)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserMAS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(gramaticaParserMAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Termino()
		}
		{
			p.SetState(191)
			p.Exopc()
		}

	case gramaticaParserMENOS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)
			p.Match(gramaticaParserMENOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Termino()
		}
		{
			p.SetState(195)
			p.Exopc()
		}

	case gramaticaParserSEMICOLON, gramaticaParserCOMMA, gramaticaParserRPAR, gramaticaParserMENOR, gramaticaParserMAYOR, gramaticaParserNEQ, gramaticaParserEQ:
		p.EnterOuterAlt(localctx, 3)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITerminoContext is an interface to support dynamic dispatch.
type ITerminoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Factor() IFactorContext
	Teropc() ITeropcContext

	// IsTerminoContext differentiates from other interfaces.
	IsTerminoContext()
}

type TerminoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminoContext() *TerminoContext {
	var p = new(TerminoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_termino
	return p
}

func InitEmptyTerminoContext(p *TerminoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_termino
}

func (*TerminoContext) IsTerminoContext() {}

func NewTerminoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminoContext {
	var p = new(TerminoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_termino

	return p
}

func (s *TerminoContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminoContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TerminoContext) Teropc() ITeropcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITeropcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITeropcContext)
}

func (s *TerminoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTermino(s)
	}
}

func (s *TerminoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTermino(s)
	}
}

func (p *gramaticaParser) Termino() (localctx ITerminoContext) {
	localctx = NewTerminoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gramaticaParserRULE_termino)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		p.Factor()
	}
	{
		p.SetState(201)
		p.Teropc()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITeropcContext is an interface to support dynamic dispatch.
type ITeropcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MULT() antlr.TerminalNode
	Factor() IFactorContext
	Teropc() ITeropcContext
	DIV() antlr.TerminalNode

	// IsTeropcContext differentiates from other interfaces.
	IsTeropcContext()
}

type TeropcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTeropcContext() *TeropcContext {
	var p = new(TeropcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_teropc
	return p
}

func InitEmptyTeropcContext(p *TeropcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_teropc
}

func (*TeropcContext) IsTeropcContext() {}

func NewTeropcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TeropcContext {
	var p = new(TeropcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_teropc

	return p
}

func (s *TeropcContext) GetParser() antlr.Parser { return s.parser }

func (s *TeropcContext) MULT() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMULT, 0)
}

func (s *TeropcContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TeropcContext) Teropc() ITeropcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITeropcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITeropcContext)
}

func (s *TeropcContext) DIV() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDIV, 0)
}

func (s *TeropcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TeropcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TeropcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterTeropc(s)
	}
}

func (s *TeropcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitTeropc(s)
	}
}

func (p *gramaticaParser) Teropc() (localctx ITeropcContext) {
	localctx = NewTeropcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gramaticaParserRULE_teropc)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserMULT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(gramaticaParserMULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Factor()
		}
		{
			p.SetState(205)
			p.Teropc()
		}

	case gramaticaParserDIV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Match(gramaticaParserDIV)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Factor()
		}
		{
			p.SetState(209)
			p.Teropc()
		}

	case gramaticaParserSEMICOLON, gramaticaParserCOMMA, gramaticaParserRPAR, gramaticaParserMENOR, gramaticaParserMAYOR, gramaticaParserMAS, gramaticaParserMENOS, gramaticaParserNEQ, gramaticaParserEQ:
		p.EnterOuterAlt(localctx, 3)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Cte() ICteContext
	LPAR() antlr.TerminalNode
	Expresion() IExpresionContext
	RPAR() antlr.TerminalNode
	MAS() antlr.TerminalNode
	Factor() IFactorContext
	MENOS() antlr.TerminalNode
	Llamada() ILlamadaContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *FactorContext) Cte() ICteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICteContext)
}

func (s *FactorContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *FactorContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *FactorContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *FactorContext) MAS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMAS, 0)
}

func (s *FactorContext) Factor() IFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *FactorContext) MENOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserMENOS, 0)
}

func (s *FactorContext) Llamada() ILlamadaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *gramaticaParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gramaticaParserRULE_factor)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Match(gramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Cte()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(gramaticaParserLPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Expresion()
		}
		{
			p.SetState(218)
			p.Match(gramaticaParserRPAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(gramaticaParserMAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.Factor()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)
			p.Match(gramaticaParserMENOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Factor()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)
			p.Llamada()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncsContext is an interface to support dynamic dispatch.
type IFuncsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Funcsopc() IFuncsopcContext
	ID() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Funcr() IFuncrContext
	RPAR() antlr.TerminalNode
	LCORCHETE() antlr.TerminalNode
	Varsdec() IVarsdecContext
	Cuerpo() ICuerpoContext
	RCORCHETE() antlr.TerminalNode

	// IsFuncsContext differentiates from other interfaces.
	IsFuncsContext()
}

type FuncsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsContext() *FuncsContext {
	var p = new(FuncsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcs
	return p
}

func InitEmptyFuncsContext(p *FuncsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcs
}

func (*FuncsContext) IsFuncsContext() {}

func NewFuncsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsContext {
	var p = new(FuncsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcs

	return p
}

func (s *FuncsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsContext) Funcsopc() IFuncsopcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncsopcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncsopcContext)
}

func (s *FuncsContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *FuncsContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *FuncsContext) Funcr() IFuncrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncrContext)
}

func (s *FuncsContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *FuncsContext) LCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLCORCHETE, 0)
}

func (s *FuncsContext) Varsdec() IVarsdecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsdecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsdecContext)
}

func (s *FuncsContext) Cuerpo() ICuerpoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICuerpoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICuerpoContext)
}

func (s *FuncsContext) RCORCHETE() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRCORCHETE, 0)
}

func (s *FuncsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncs(s)
	}
}

func (s *FuncsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncs(s)
	}
}

func (p *gramaticaParser) Funcs() (localctx IFuncsContext) {
	localctx = NewFuncsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gramaticaParserRULE_funcs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Funcsopc()
	}
	{
		p.SetState(228)
		p.Match(gramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(gramaticaParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Funcr()
	}
	{
		p.SetState(231)
		p.Match(gramaticaParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Match(gramaticaParserLCORCHETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(233)
		p.Varsdec()
	}
	{
		p.SetState(234)
		p.Cuerpo()
	}
	{
		p.SetState(235)
		p.Match(gramaticaParserRCORCHETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncsopcContext is an interface to support dynamic dispatch.
type IFuncsopcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NULA() antlr.TerminalNode
	Tipo() ITipoContext

	// IsFuncsopcContext differentiates from other interfaces.
	IsFuncsopcContext()
}

type FuncsopcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncsopcContext() *FuncsopcContext {
	var p = new(FuncsopcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcsopc
	return p
}

func InitEmptyFuncsopcContext(p *FuncsopcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcsopc
}

func (*FuncsopcContext) IsFuncsopcContext() {}

func NewFuncsopcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncsopcContext {
	var p = new(FuncsopcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcsopc

	return p
}

func (s *FuncsopcContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncsopcContext) NULA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserNULA, 0)
}

func (s *FuncsopcContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncsopcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncsopcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncsopcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncsopc(s)
	}
}

func (s *FuncsopcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncsopc(s)
	}
}

func (p *gramaticaParser) Funcsopc() (localctx IFuncsopcContext) {
	localctx = NewFuncsopcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gramaticaParserRULE_funcsopc)
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserNULA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(237)
			p.Match(gramaticaParserNULA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case gramaticaParserENTERO, gramaticaParserFLOTANTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Tipo()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncrContext is an interface to support dynamic dispatch.
type IFuncrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	DPUNTOS() antlr.TerminalNode
	Tipo() ITipoContext
	COMMA() antlr.TerminalNode
	Funcr() IFuncrContext

	// IsFuncrContext differentiates from other interfaces.
	IsFuncrContext()
}

type FuncrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncrContext() *FuncrContext {
	var p = new(FuncrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcr
	return p
}

func InitEmptyFuncrContext(p *FuncrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_funcr
}

func (*FuncrContext) IsFuncrContext() {}

func NewFuncrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncrContext {
	var p = new(FuncrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_funcr

	return p
}

func (s *FuncrContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncrContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *FuncrContext) DPUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDPUNTOS, 0)
}

func (s *FuncrContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *FuncrContext) COMMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMMA, 0)
}

func (s *FuncrContext) Funcr() IFuncrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncrContext)
}

func (s *FuncrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterFuncr(s)
	}
}

func (s *FuncrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitFuncr(s)
	}
}

func (p *gramaticaParser) Funcr() (localctx IFuncrContext) {
	localctx = NewFuncrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gramaticaParserRULE_funcr)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(241)
			p.Match(gramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(242)
			p.Match(gramaticaParserDPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Tipo()
		}
		{
			p.SetState(244)
			p.Match(gramaticaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Funcr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.Match(gramaticaParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Match(gramaticaParserDPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(249)
			p.Tipo()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsdecContext is an interface to support dynamic dispatch.
type IVarsdecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VARS() antlr.TerminalNode
	DPUNTOS() antlr.TerminalNode
	Varsop() IVarsopContext

	// IsVarsdecContext differentiates from other interfaces.
	IsVarsdecContext()
}

type VarsdecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsdecContext() *VarsdecContext {
	var p = new(VarsdecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varsdec
	return p
}

func InitEmptyVarsdecContext(p *VarsdecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_varsdec
}

func (*VarsdecContext) IsVarsdecContext() {}

func NewVarsdecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsdecContext {
	var p = new(VarsdecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_varsdec

	return p
}

func (s *VarsdecContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsdecContext) VARS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserVARS, 0)
}

func (s *VarsdecContext) DPUNTOS() antlr.TerminalNode {
	return s.GetToken(gramaticaParserDPUNTOS, 0)
}

func (s *VarsdecContext) Varsop() IVarsopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsopContext)
}

func (s *VarsdecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsdecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsdecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterVarsdec(s)
	}
}

func (s *VarsdecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitVarsdec(s)
	}
}

func (p *gramaticaParser) Varsdec() (localctx IVarsdecContext) {
	localctx = NewVarsdecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gramaticaParserRULE_varsdec)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case gramaticaParserVARS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(gramaticaParserVARS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(gramaticaParserDPUNTOS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Varsop()
		}

	case gramaticaParserLCORCHETE:
		p.EnterOuterAlt(localctx, 2)

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaContext is an interface to support dynamic dispatch.
type ILlamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAR() antlr.TerminalNode
	Llamadaexp() ILlamadaexpContext
	RPAR() antlr.TerminalNode

	// IsLlamadaContext differentiates from other interfaces.
	IsLlamadaContext()
}

type LlamadaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaContext() *LlamadaContext {
	var p = new(LlamadaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamada
	return p
}

func InitEmptyLlamadaContext(p *LlamadaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamada
}

func (*LlamadaContext) IsLlamadaContext() {}

func NewLlamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaContext {
	var p = new(LlamadaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamada

	return p
}

func (s *LlamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaContext) ID() antlr.TerminalNode {
	return s.GetToken(gramaticaParserID, 0)
}

func (s *LlamadaContext) LPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserLPAR, 0)
}

func (s *LlamadaContext) Llamadaexp() ILlamadaexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaexpContext)
}

func (s *LlamadaContext) RPAR() antlr.TerminalNode {
	return s.GetToken(gramaticaParserRPAR, 0)
}

func (s *LlamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamada(s)
	}
}

func (s *LlamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamada(s)
	}
}

func (p *gramaticaParser) Llamada() (localctx ILlamadaContext) {
	localctx = NewLlamadaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gramaticaParserRULE_llamada)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(gramaticaParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(gramaticaParserLPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Llamadaexp()
	}
	{
		p.SetState(262)
		p.Match(gramaticaParserRPAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILlamadaexpContext is an interface to support dynamic dispatch.
type ILlamadaexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expresion() IExpresionContext
	COMMA() antlr.TerminalNode
	Llamadaexp() ILlamadaexpContext

	// IsLlamadaexpContext differentiates from other interfaces.
	IsLlamadaexpContext()
}

type LlamadaexpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLlamadaexpContext() *LlamadaexpContext {
	var p = new(LlamadaexpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaexp
	return p
}

func InitEmptyLlamadaexpContext(p *LlamadaexpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = gramaticaParserRULE_llamadaexp
}

func (*LlamadaexpContext) IsLlamadaexpContext() {}

func NewLlamadaexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LlamadaexpContext {
	var p = new(LlamadaexpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = gramaticaParserRULE_llamadaexp

	return p
}

func (s *LlamadaexpContext) GetParser() antlr.Parser { return s.parser }

func (s *LlamadaexpContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *LlamadaexpContext) COMMA() antlr.TerminalNode {
	return s.GetToken(gramaticaParserCOMMA, 0)
}

func (s *LlamadaexpContext) Llamadaexp() ILlamadaexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILlamadaexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILlamadaexpContext)
}

func (s *LlamadaexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LlamadaexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LlamadaexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.EnterLlamadaexp(s)
	}
}

func (s *LlamadaexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gramaticaListener); ok {
		listenerT.ExitLlamadaexp(s)
	}
}

func (p *gramaticaParser) Llamadaexp() (localctx ILlamadaexpContext) {
	localctx = NewLlamadaexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gramaticaParserRULE_llamadaexp)
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Expresion()
		}
		{
			p.SetState(265)
			p.Match(gramaticaParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Llamadaexp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(268)
			p.Expresion()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
