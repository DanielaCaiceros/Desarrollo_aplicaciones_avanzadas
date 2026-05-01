// Code generated from gramatica.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type gramaticaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GramaticaLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gramaticalexerLexerInit() {
	staticData := &GramaticaLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"SEMICOLON", "COMMA", "DPUNTOS", "LCORCHETE", "RCORCHETE", "LPAR", "RPAR",
		"LBRACKET", "RBRACKET", "ASIG", "MENOR", "MAYOR", "MAS", "MENOS", "MULT",
		"DIV", "NEQ", "EQ", "CTE_FLOAT", "CTE_ENT", "PROGRAMA", "INICIO", "FIN",
		"VARS", "ENTERO", "FLOTANTE", "ESCRIBE", "MIENTRAS", "HAZ", "SI", "SINO",
		"NULA", "LETRERO", "ID", "WS", "COMMENT_LINE", "COMMENT_BLOCK",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 254, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 4, 18, 115, 8, 18, 11, 18, 12,
		18, 116, 1, 18, 1, 18, 4, 18, 121, 8, 18, 11, 18, 12, 18, 122, 1, 19, 4,
		19, 126, 8, 19, 11, 19, 12, 19, 127, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 5, 32, 207,
		8, 32, 10, 32, 12, 32, 210, 9, 32, 1, 32, 1, 32, 1, 33, 1, 33, 5, 33, 216,
		8, 33, 10, 33, 12, 33, 219, 9, 33, 1, 34, 4, 34, 222, 8, 34, 11, 34, 12,
		34, 223, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 232, 8, 35, 10,
		35, 12, 35, 235, 9, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 5, 36, 245, 8, 36, 10, 36, 12, 36, 248, 9, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 246, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25,
		51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34,
		69, 35, 71, 36, 73, 37, 1, 0, 6, 1, 0, 48, 57, 1, 0, 34, 34, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 10, 10, 94, 94, 261, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7,
		81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1, 0, 0,
		0, 15, 89, 1, 0, 0, 0, 17, 91, 1, 0, 0, 0, 19, 93, 1, 0, 0, 0, 21, 95,
		1, 0, 0, 0, 23, 97, 1, 0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 101, 1, 0, 0, 0,
		29, 103, 1, 0, 0, 0, 31, 105, 1, 0, 0, 0, 33, 107, 1, 0, 0, 0, 35, 110,
		1, 0, 0, 0, 37, 114, 1, 0, 0, 0, 39, 125, 1, 0, 0, 0, 41, 129, 1, 0, 0,
		0, 43, 138, 1, 0, 0, 0, 45, 145, 1, 0, 0, 0, 47, 149, 1, 0, 0, 0, 49, 154,
		1, 0, 0, 0, 51, 161, 1, 0, 0, 0, 53, 170, 1, 0, 0, 0, 55, 178, 1, 0, 0,
		0, 57, 187, 1, 0, 0, 0, 59, 191, 1, 0, 0, 0, 61, 194, 1, 0, 0, 0, 63, 199,
		1, 0, 0, 0, 65, 204, 1, 0, 0, 0, 67, 213, 1, 0, 0, 0, 69, 221, 1, 0, 0,
		0, 71, 227, 1, 0, 0, 0, 73, 240, 1, 0, 0, 0, 75, 76, 5, 59, 0, 0, 76, 2,
		1, 0, 0, 0, 77, 78, 5, 44, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 58, 0, 0,
		80, 6, 1, 0, 0, 0, 81, 82, 5, 123, 0, 0, 82, 8, 1, 0, 0, 0, 83, 84, 5,
		125, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 40, 0, 0, 86, 12, 1, 0, 0, 0,
		87, 88, 5, 41, 0, 0, 88, 14, 1, 0, 0, 0, 89, 90, 5, 91, 0, 0, 90, 16, 1,
		0, 0, 0, 91, 92, 5, 93, 0, 0, 92, 18, 1, 0, 0, 0, 93, 94, 5, 61, 0, 0,
		94, 20, 1, 0, 0, 0, 95, 96, 5, 60, 0, 0, 96, 22, 1, 0, 0, 0, 97, 98, 5,
		62, 0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 43, 0, 0, 100, 26, 1, 0, 0, 0,
		101, 102, 5, 45, 0, 0, 102, 28, 1, 0, 0, 0, 103, 104, 5, 42, 0, 0, 104,
		30, 1, 0, 0, 0, 105, 106, 5, 47, 0, 0, 106, 32, 1, 0, 0, 0, 107, 108, 5,
		33, 0, 0, 108, 109, 5, 61, 0, 0, 109, 34, 1, 0, 0, 0, 110, 111, 5, 61,
		0, 0, 111, 112, 5, 61, 0, 0, 112, 36, 1, 0, 0, 0, 113, 115, 7, 0, 0, 0,
		114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 120, 5, 46, 0, 0, 119, 121,
		7, 0, 0, 0, 120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0,
		0, 0, 122, 123, 1, 0, 0, 0, 123, 38, 1, 0, 0, 0, 124, 126, 7, 0, 0, 0,
		125, 124, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 40, 1, 0, 0, 0, 129, 130, 5, 112, 0, 0, 130, 131,
		5, 114, 0, 0, 131, 132, 5, 111, 0, 0, 132, 133, 5, 103, 0, 0, 133, 134,
		5, 114, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 109, 0, 0, 136, 137,
		5, 97, 0, 0, 137, 42, 1, 0, 0, 0, 138, 139, 5, 105, 0, 0, 139, 140, 5,
		110, 0, 0, 140, 141, 5, 105, 0, 0, 141, 142, 5, 99, 0, 0, 142, 143, 5,
		105, 0, 0, 143, 144, 5, 111, 0, 0, 144, 44, 1, 0, 0, 0, 145, 146, 5, 102,
		0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 110, 0, 0, 148, 46, 1, 0, 0,
		0, 149, 150, 5, 118, 0, 0, 150, 151, 5, 97, 0, 0, 151, 152, 5, 114, 0,
		0, 152, 153, 5, 115, 0, 0, 153, 48, 1, 0, 0, 0, 154, 155, 5, 101, 0, 0,
		155, 156, 5, 110, 0, 0, 156, 157, 5, 116, 0, 0, 157, 158, 5, 101, 0, 0,
		158, 159, 5, 114, 0, 0, 159, 160, 5, 111, 0, 0, 160, 50, 1, 0, 0, 0, 161,
		162, 5, 102, 0, 0, 162, 163, 5, 108, 0, 0, 163, 164, 5, 111, 0, 0, 164,
		165, 5, 116, 0, 0, 165, 166, 5, 97, 0, 0, 166, 167, 5, 110, 0, 0, 167,
		168, 5, 116, 0, 0, 168, 169, 5, 101, 0, 0, 169, 52, 1, 0, 0, 0, 170, 171,
		5, 101, 0, 0, 171, 172, 5, 115, 0, 0, 172, 173, 5, 99, 0, 0, 173, 174,
		5, 114, 0, 0, 174, 175, 5, 105, 0, 0, 175, 176, 5, 98, 0, 0, 176, 177,
		5, 101, 0, 0, 177, 54, 1, 0, 0, 0, 178, 179, 5, 109, 0, 0, 179, 180, 5,
		105, 0, 0, 180, 181, 5, 101, 0, 0, 181, 182, 5, 110, 0, 0, 182, 183, 5,
		116, 0, 0, 183, 184, 5, 114, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186, 5,
		115, 0, 0, 186, 56, 1, 0, 0, 0, 187, 188, 5, 104, 0, 0, 188, 189, 5, 97,
		0, 0, 189, 190, 5, 122, 0, 0, 190, 58, 1, 0, 0, 0, 191, 192, 5, 115, 0,
		0, 192, 193, 5, 105, 0, 0, 193, 60, 1, 0, 0, 0, 194, 195, 5, 115, 0, 0,
		195, 196, 5, 105, 0, 0, 196, 197, 5, 110, 0, 0, 197, 198, 5, 111, 0, 0,
		198, 62, 1, 0, 0, 0, 199, 200, 5, 110, 0, 0, 200, 201, 5, 117, 0, 0, 201,
		202, 5, 108, 0, 0, 202, 203, 5, 97, 0, 0, 203, 64, 1, 0, 0, 0, 204, 208,
		5, 34, 0, 0, 205, 207, 8, 1, 0, 0, 206, 205, 1, 0, 0, 0, 207, 210, 1, 0,
		0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 211, 1, 0, 0, 0,
		210, 208, 1, 0, 0, 0, 211, 212, 5, 34, 0, 0, 212, 66, 1, 0, 0, 0, 213,
		217, 7, 2, 0, 0, 214, 216, 7, 3, 0, 0, 215, 214, 1, 0, 0, 0, 216, 219,
		1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 68, 1, 0,
		0, 0, 219, 217, 1, 0, 0, 0, 220, 222, 7, 4, 0, 0, 221, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 224, 1, 0, 0, 0, 224,
		225, 1, 0, 0, 0, 225, 226, 6, 34, 0, 0, 226, 70, 1, 0, 0, 0, 227, 228,
		5, 47, 0, 0, 228, 229, 5, 47, 0, 0, 229, 233, 1, 0, 0, 0, 230, 232, 7,
		5, 0, 0, 231, 230, 1, 0, 0, 0, 232, 235, 1, 0, 0, 0, 233, 231, 1, 0, 0,
		0, 233, 234, 1, 0, 0, 0, 234, 236, 1, 0, 0, 0, 235, 233, 1, 0, 0, 0, 236,
		237, 5, 10, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 6, 35, 0, 0, 239, 72,
		1, 0, 0, 0, 240, 241, 5, 47, 0, 0, 241, 242, 5, 42, 0, 0, 242, 246, 1,
		0, 0, 0, 243, 245, 9, 0, 0, 0, 244, 243, 1, 0, 0, 0, 245, 248, 1, 0, 0,
		0, 246, 247, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 247, 249, 1, 0, 0, 0, 248,
		246, 1, 0, 0, 0, 249, 250, 5, 42, 0, 0, 250, 251, 5, 47, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 253, 6, 36, 0, 0, 253, 74, 1, 0, 0, 0, 9, 0, 116, 122,
		127, 208, 217, 223, 233, 246, 1, 6, 0, 0,
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

// gramaticaLexerInit initializes any static state used to implement gramaticaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewgramaticaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GramaticaLexerInit() {
	staticData := &GramaticaLexerLexerStaticData
	staticData.once.Do(gramaticalexerLexerInit)
}

// NewgramaticaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewgramaticaLexer(input antlr.CharStream) *gramaticaLexer {
	GramaticaLexerInit()
	l := new(gramaticaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GramaticaLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gramaticaLexer tokens.
const (
	gramaticaLexerSEMICOLON     = 1
	gramaticaLexerCOMMA         = 2
	gramaticaLexerDPUNTOS       = 3
	gramaticaLexerLCORCHETE     = 4
	gramaticaLexerRCORCHETE     = 5
	gramaticaLexerLPAR          = 6
	gramaticaLexerRPAR          = 7
	gramaticaLexerLBRACKET      = 8
	gramaticaLexerRBRACKET      = 9
	gramaticaLexerASIG          = 10
	gramaticaLexerMENOR         = 11
	gramaticaLexerMAYOR         = 12
	gramaticaLexerMAS           = 13
	gramaticaLexerMENOS         = 14
	gramaticaLexerMULT          = 15
	gramaticaLexerDIV           = 16
	gramaticaLexerNEQ           = 17
	gramaticaLexerEQ            = 18
	gramaticaLexerCTE_FLOAT     = 19
	gramaticaLexerCTE_ENT       = 20
	gramaticaLexerPROGRAMA      = 21
	gramaticaLexerINICIO        = 22
	gramaticaLexerFIN           = 23
	gramaticaLexerVARS          = 24
	gramaticaLexerENTERO        = 25
	gramaticaLexerFLOTANTE      = 26
	gramaticaLexerESCRIBE       = 27
	gramaticaLexerMIENTRAS      = 28
	gramaticaLexerHAZ           = 29
	gramaticaLexerSI            = 30
	gramaticaLexerSINO          = 31
	gramaticaLexerNULA          = 32
	gramaticaLexerLETRERO       = 33
	gramaticaLexerID            = 34
	gramaticaLexerWS            = 35
	gramaticaLexerCOMMENT_LINE  = 36
	gramaticaLexerCOMMENT_BLOCK = 37
)
