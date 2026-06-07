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
		"'inicio'", "'fin'", "'vars'", "'entero'", "'boolean'", "'string'",
		"'flotante'", "'escribe'", "'mientras'", "'haz'", "'si'", "'sino'",
		"'nula'", "'regresa'",
	}
	staticData.SymbolicNames = []string{
		"", "SEMICOLON", "COMMA", "DPUNTOS", "LCORCHETE", "RCORCHETE", "LPAR",
		"RPAR", "LBRACKET", "RBRACKET", "ASIG", "MENOR", "MAYOR", "MAS", "MENOS",
		"MULT", "DIV", "NEQ", "EQ", "CTE_FLOAT", "CTE_ENT", "PROGRAMA", "INICIO",
		"FIN", "VARS", "ENTERO", "BOOLEAN", "STRING", "FLOTANTE", "ESCRIBE",
		"MIENTRAS", "HAZ", "SI", "SINO", "NULA", "REGRESA", "LETRERO", "ID",
		"WS", "COMMENT_LINE", "COMMENT_BLOCK",
	}
	staticData.RuleNames = []string{
		"SEMICOLON", "COMMA", "DPUNTOS", "LCORCHETE", "RCORCHETE", "LPAR", "RPAR",
		"LBRACKET", "RBRACKET", "ASIG", "MENOR", "MAYOR", "MAS", "MENOS", "MULT",
		"DIV", "NEQ", "EQ", "CTE_FLOAT", "CTE_ENT", "PROGRAMA", "INICIO", "FIN",
		"VARS", "ENTERO", "BOOLEAN", "STRING", "FLOTANTE", "ESCRIBE", "MIENTRAS",
		"HAZ", "SI", "SINO", "NULA", "REGRESA", "LETRERO", "ID", "WS", "COMMENT_LINE",
		"COMMENT_BLOCK",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 40, 283, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 18, 4, 18, 121, 8, 18, 11, 18, 12, 18, 122, 1, 18, 1, 18, 4, 18,
		127, 8, 18, 11, 18, 12, 18, 128, 1, 19, 4, 19, 132, 8, 19, 11, 19, 12,
		19, 133, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 236, 8, 35, 10,
		35, 12, 35, 239, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36, 5, 36, 245, 8, 36,
		10, 36, 12, 36, 248, 9, 36, 1, 37, 4, 37, 251, 8, 37, 11, 37, 12, 37, 252,
		1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 261, 8, 38, 10, 38, 12,
		38, 264, 9, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39,
		5, 39, 274, 8, 39, 10, 39, 12, 39, 277, 9, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 275, 0, 40, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 1, 0, 6, 1, 0, 48, 57, 1, 0, 34,
		34, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122,
		3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 94, 94, 290, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1,
		0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33,
		1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0,
		41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0,
		0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0,
		0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
		0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1,
		0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79,
		1, 0, 0, 0, 1, 81, 1, 0, 0, 0, 3, 83, 1, 0, 0, 0, 5, 85, 1, 0, 0, 0, 7,
		87, 1, 0, 0, 0, 9, 89, 1, 0, 0, 0, 11, 91, 1, 0, 0, 0, 13, 93, 1, 0, 0,
		0, 15, 95, 1, 0, 0, 0, 17, 97, 1, 0, 0, 0, 19, 99, 1, 0, 0, 0, 21, 101,
		1, 0, 0, 0, 23, 103, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0, 27, 107, 1, 0, 0,
		0, 29, 109, 1, 0, 0, 0, 31, 111, 1, 0, 0, 0, 33, 113, 1, 0, 0, 0, 35, 116,
		1, 0, 0, 0, 37, 120, 1, 0, 0, 0, 39, 131, 1, 0, 0, 0, 41, 135, 1, 0, 0,
		0, 43, 144, 1, 0, 0, 0, 45, 151, 1, 0, 0, 0, 47, 155, 1, 0, 0, 0, 49, 160,
		1, 0, 0, 0, 51, 167, 1, 0, 0, 0, 53, 175, 1, 0, 0, 0, 55, 182, 1, 0, 0,
		0, 57, 191, 1, 0, 0, 0, 59, 199, 1, 0, 0, 0, 61, 208, 1, 0, 0, 0, 63, 212,
		1, 0, 0, 0, 65, 215, 1, 0, 0, 0, 67, 220, 1, 0, 0, 0, 69, 225, 1, 0, 0,
		0, 71, 233, 1, 0, 0, 0, 73, 242, 1, 0, 0, 0, 75, 250, 1, 0, 0, 0, 77, 256,
		1, 0, 0, 0, 79, 269, 1, 0, 0, 0, 81, 82, 5, 59, 0, 0, 82, 2, 1, 0, 0, 0,
		83, 84, 5, 44, 0, 0, 84, 4, 1, 0, 0, 0, 85, 86, 5, 58, 0, 0, 86, 6, 1,
		0, 0, 0, 87, 88, 5, 123, 0, 0, 88, 8, 1, 0, 0, 0, 89, 90, 5, 125, 0, 0,
		90, 10, 1, 0, 0, 0, 91, 92, 5, 40, 0, 0, 92, 12, 1, 0, 0, 0, 93, 94, 5,
		41, 0, 0, 94, 14, 1, 0, 0, 0, 95, 96, 5, 91, 0, 0, 96, 16, 1, 0, 0, 0,
		97, 98, 5, 93, 0, 0, 98, 18, 1, 0, 0, 0, 99, 100, 5, 61, 0, 0, 100, 20,
		1, 0, 0, 0, 101, 102, 5, 60, 0, 0, 102, 22, 1, 0, 0, 0, 103, 104, 5, 62,
		0, 0, 104, 24, 1, 0, 0, 0, 105, 106, 5, 43, 0, 0, 106, 26, 1, 0, 0, 0,
		107, 108, 5, 45, 0, 0, 108, 28, 1, 0, 0, 0, 109, 110, 5, 42, 0, 0, 110,
		30, 1, 0, 0, 0, 111, 112, 5, 47, 0, 0, 112, 32, 1, 0, 0, 0, 113, 114, 5,
		33, 0, 0, 114, 115, 5, 61, 0, 0, 115, 34, 1, 0, 0, 0, 116, 117, 5, 61,
		0, 0, 117, 118, 5, 61, 0, 0, 118, 36, 1, 0, 0, 0, 119, 121, 7, 0, 0, 0,
		120, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 122,
		123, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 5, 46, 0, 0, 125, 127,
		7, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 126, 1, 0,
		0, 0, 128, 129, 1, 0, 0, 0, 129, 38, 1, 0, 0, 0, 130, 132, 7, 0, 0, 0,
		131, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133,
		134, 1, 0, 0, 0, 134, 40, 1, 0, 0, 0, 135, 136, 5, 112, 0, 0, 136, 137,
		5, 114, 0, 0, 137, 138, 5, 111, 0, 0, 138, 139, 5, 103, 0, 0, 139, 140,
		5, 114, 0, 0, 140, 141, 5, 97, 0, 0, 141, 142, 5, 109, 0, 0, 142, 143,
		5, 97, 0, 0, 143, 42, 1, 0, 0, 0, 144, 145, 5, 105, 0, 0, 145, 146, 5,
		110, 0, 0, 146, 147, 5, 105, 0, 0, 147, 148, 5, 99, 0, 0, 148, 149, 5,
		105, 0, 0, 149, 150, 5, 111, 0, 0, 150, 44, 1, 0, 0, 0, 151, 152, 5, 102,
		0, 0, 152, 153, 5, 105, 0, 0, 153, 154, 5, 110, 0, 0, 154, 46, 1, 0, 0,
		0, 155, 156, 5, 118, 0, 0, 156, 157, 5, 97, 0, 0, 157, 158, 5, 114, 0,
		0, 158, 159, 5, 115, 0, 0, 159, 48, 1, 0, 0, 0, 160, 161, 5, 101, 0, 0,
		161, 162, 5, 110, 0, 0, 162, 163, 5, 116, 0, 0, 163, 164, 5, 101, 0, 0,
		164, 165, 5, 114, 0, 0, 165, 166, 5, 111, 0, 0, 166, 50, 1, 0, 0, 0, 167,
		168, 5, 98, 0, 0, 168, 169, 5, 111, 0, 0, 169, 170, 5, 111, 0, 0, 170,
		171, 5, 108, 0, 0, 171, 172, 5, 101, 0, 0, 172, 173, 5, 97, 0, 0, 173,
		174, 5, 110, 0, 0, 174, 52, 1, 0, 0, 0, 175, 176, 5, 115, 0, 0, 176, 177,
		5, 116, 0, 0, 177, 178, 5, 114, 0, 0, 178, 179, 5, 105, 0, 0, 179, 180,
		5, 110, 0, 0, 180, 181, 5, 103, 0, 0, 181, 54, 1, 0, 0, 0, 182, 183, 5,
		102, 0, 0, 183, 184, 5, 108, 0, 0, 184, 185, 5, 111, 0, 0, 185, 186, 5,
		116, 0, 0, 186, 187, 5, 97, 0, 0, 187, 188, 5, 110, 0, 0, 188, 189, 5,
		116, 0, 0, 189, 190, 5, 101, 0, 0, 190, 56, 1, 0, 0, 0, 191, 192, 5, 101,
		0, 0, 192, 193, 5, 115, 0, 0, 193, 194, 5, 99, 0, 0, 194, 195, 5, 114,
		0, 0, 195, 196, 5, 105, 0, 0, 196, 197, 5, 98, 0, 0, 197, 198, 5, 101,
		0, 0, 198, 58, 1, 0, 0, 0, 199, 200, 5, 109, 0, 0, 200, 201, 5, 105, 0,
		0, 201, 202, 5, 101, 0, 0, 202, 203, 5, 110, 0, 0, 203, 204, 5, 116, 0,
		0, 204, 205, 5, 114, 0, 0, 205, 206, 5, 97, 0, 0, 206, 207, 5, 115, 0,
		0, 207, 60, 1, 0, 0, 0, 208, 209, 5, 104, 0, 0, 209, 210, 5, 97, 0, 0,
		210, 211, 5, 122, 0, 0, 211, 62, 1, 0, 0, 0, 212, 213, 5, 115, 0, 0, 213,
		214, 5, 105, 0, 0, 214, 64, 1, 0, 0, 0, 215, 216, 5, 115, 0, 0, 216, 217,
		5, 105, 0, 0, 217, 218, 5, 110, 0, 0, 218, 219, 5, 111, 0, 0, 219, 66,
		1, 0, 0, 0, 220, 221, 5, 110, 0, 0, 221, 222, 5, 117, 0, 0, 222, 223, 5,
		108, 0, 0, 223, 224, 5, 97, 0, 0, 224, 68, 1, 0, 0, 0, 225, 226, 5, 114,
		0, 0, 226, 227, 5, 101, 0, 0, 227, 228, 5, 103, 0, 0, 228, 229, 5, 114,
		0, 0, 229, 230, 5, 101, 0, 0, 230, 231, 5, 115, 0, 0, 231, 232, 5, 97,
		0, 0, 232, 70, 1, 0, 0, 0, 233, 237, 5, 34, 0, 0, 234, 236, 8, 1, 0, 0,
		235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0, 0, 0, 237,
		238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 240, 241,
		5, 34, 0, 0, 241, 72, 1, 0, 0, 0, 242, 246, 7, 2, 0, 0, 243, 245, 7, 3,
		0, 0, 244, 243, 1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0,
		246, 247, 1, 0, 0, 0, 247, 74, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 251,
		7, 4, 0, 0, 250, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 250, 1, 0,
		0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 6, 37, 0, 0,
		255, 76, 1, 0, 0, 0, 256, 257, 5, 47, 0, 0, 257, 258, 5, 47, 0, 0, 258,
		262, 1, 0, 0, 0, 259, 261, 7, 5, 0, 0, 260, 259, 1, 0, 0, 0, 261, 264,
		1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 1, 0,
		0, 0, 264, 262, 1, 0, 0, 0, 265, 266, 5, 10, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 268, 6, 38, 0, 0, 268, 78, 1, 0, 0, 0, 269, 270, 5, 47, 0, 0, 270,
		271, 5, 42, 0, 0, 271, 275, 1, 0, 0, 0, 272, 274, 9, 0, 0, 0, 273, 272,
		1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 275, 273, 1, 0,
		0, 0, 276, 278, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 279, 5, 42, 0, 0,
		279, 280, 5, 47, 0, 0, 280, 281, 1, 0, 0, 0, 281, 282, 6, 39, 0, 0, 282,
		80, 1, 0, 0, 0, 9, 0, 122, 128, 133, 237, 246, 252, 262, 275, 1, 6, 0,
		0,
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
	gramaticaLexerBOOLEAN       = 26
	gramaticaLexerSTRING        = 27
	gramaticaLexerFLOTANTE      = 28
	gramaticaLexerESCRIBE       = 29
	gramaticaLexerMIENTRAS      = 30
	gramaticaLexerHAZ           = 31
	gramaticaLexerSI            = 32
	gramaticaLexerSINO          = 33
	gramaticaLexerNULA          = 34
	gramaticaLexerREGRESA       = 35
	gramaticaLexerLETRERO       = 36
	gramaticaLexerID            = 37
	gramaticaLexerWS            = 38
	gramaticaLexerCOMMENT_LINE  = 39
	gramaticaLexerCOMMENT_BLOCK = 40
)
