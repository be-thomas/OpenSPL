// Code generated from pkg/parser/spl2/grammar/SPL2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spl2gen // SPL2
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

type SPL2Parser struct {
	*antlr.BaseParser
}

var SPL2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func spl2ParserInit() {
	staticData := &SPL2ParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "'=='", "'!='",
		"'<='", "'>='", "'='", "'<'", "'>'", "'!'", "'|'", "','", "'('", "')'",
		"'['", "']'", "'+'", "'-'", "'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "FROM", "WHERE", "STATS", "BY", "AS", "AND", "OR", "NOT", "DURATION",
		"NUMBER", "STRING", "IDENT", "EQ2", "NE", "LE", "GE", "EQ", "LT", "GT",
		"BANG", "PIPE", "COMMA", "LPAREN", "RPAREN", "LBRACKET", "RBRACKET",
		"PLUS", "MINUS", "STAR", "SLASH", "PERCENT", "WS", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"stageFrom", "stageWhere", "stageStats", "stageEnd", "indexName", "indexTail",
		"expr", "orExpr", "orTail", "andExpr", "andTail", "notExpr", "comparison",
		"comparisonTail", "additive", "additiveTail", "multiplicative", "multiplicativeTail",
		"unary", "primary", "subsearch", "absentOperand", "aggList", "agg",
		"aliasClause", "byClause", "fieldList", "field", "compOp", "addOp",
		"mulOp", "sign", "identLike",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 33, 265, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 78, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 5, 4, 86,
		8, 4, 10, 4, 12, 4, 89, 9, 4, 1, 4, 1, 4, 3, 4, 93, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 100, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 5, 7, 106, 8, 7,
		10, 7, 12, 7, 109, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 115, 8, 8, 1, 9,
		1, 9, 5, 9, 119, 8, 9, 10, 9, 12, 9, 122, 9, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 128, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 133, 8, 11, 1, 11, 3,
		11, 136, 8, 11, 1, 12, 1, 12, 5, 12, 140, 8, 12, 10, 12, 12, 12, 143, 9,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 149, 8, 13, 1, 14, 1, 14, 5, 14,
		153, 8, 14, 10, 14, 12, 14, 156, 9, 14, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 162, 8, 15, 1, 16, 1, 16, 5, 16, 166, 8, 16, 10, 16, 12, 16, 169, 9,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 175, 8, 17, 1, 18, 1, 18, 1, 18,
		3, 18, 180, 8, 18, 1, 18, 3, 18, 183, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 195, 8, 19, 10, 19, 12,
		19, 198, 9, 19, 3, 19, 200, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 3, 19, 209, 8, 19, 1, 20, 1, 20, 1, 20, 5, 20, 214, 8, 20, 10,
		20, 12, 20, 217, 9, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22,
		225, 8, 22, 10, 22, 12, 22, 228, 9, 22, 1, 23, 1, 23, 1, 23, 3, 23, 233,
		8, 23, 1, 23, 1, 23, 3, 23, 237, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 5, 26, 248, 8, 26, 10, 26, 12, 26, 251,
		9, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 0, 0, 33, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
		58, 60, 62, 64, 0, 7, 1, 1, 21, 21, 2, 0, 8, 8, 20, 20, 4, 1, 13, 19, 21,
		22, 24, 24, 26, 26, 1, 0, 13, 19, 1, 0, 27, 28, 1, 0, 29, 31, 2, 0, 1,
		8, 12, 12, 263, 0, 66, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 4, 74, 1, 0, 0, 0,
		6, 81, 1, 0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 94, 1, 0, 0, 0, 12, 101, 1, 0,
		0, 0, 14, 103, 1, 0, 0, 0, 16, 110, 1, 0, 0, 0, 18, 116, 1, 0, 0, 0, 20,
		123, 1, 0, 0, 0, 22, 135, 1, 0, 0, 0, 24, 137, 1, 0, 0, 0, 26, 144, 1,
		0, 0, 0, 28, 150, 1, 0, 0, 0, 30, 157, 1, 0, 0, 0, 32, 163, 1, 0, 0, 0,
		34, 170, 1, 0, 0, 0, 36, 182, 1, 0, 0, 0, 38, 208, 1, 0, 0, 0, 40, 210,
		1, 0, 0, 0, 42, 218, 1, 0, 0, 0, 44, 221, 1, 0, 0, 0, 46, 229, 1, 0, 0,
		0, 48, 238, 1, 0, 0, 0, 50, 241, 1, 0, 0, 0, 52, 244, 1, 0, 0, 0, 54, 252,
		1, 0, 0, 0, 56, 254, 1, 0, 0, 0, 58, 256, 1, 0, 0, 0, 60, 258, 1, 0, 0,
		0, 62, 260, 1, 0, 0, 0, 64, 262, 1, 0, 0, 0, 66, 67, 5, 1, 0, 0, 67, 68,
		3, 8, 4, 0, 68, 69, 3, 6, 3, 0, 69, 1, 1, 0, 0, 0, 70, 71, 5, 2, 0, 0,
		71, 72, 3, 12, 6, 0, 72, 73, 3, 6, 3, 0, 73, 3, 1, 0, 0, 0, 74, 75, 5,
		3, 0, 0, 75, 77, 3, 44, 22, 0, 76, 78, 3, 50, 25, 0, 77, 76, 1, 0, 0, 0,
		77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 3, 6, 3, 0, 80, 5, 1, 0,
		0, 0, 81, 82, 7, 0, 0, 0, 82, 7, 1, 0, 0, 0, 83, 87, 3, 64, 32, 0, 84,
		86, 3, 10, 5, 0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0,
		0, 0, 87, 88, 1, 0, 0, 0, 88, 92, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91,
		4, 4, 0, 0, 91, 93, 5, 29, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 9, 1, 0, 0, 0, 94, 95, 4, 5, 1, 0, 95, 99, 5, 28, 0, 0, 96, 100, 3,
		64, 32, 0, 97, 100, 5, 10, 0, 0, 98, 100, 5, 29, 0, 0, 99, 96, 1, 0, 0,
		0, 99, 97, 1, 0, 0, 0, 99, 98, 1, 0, 0, 0, 100, 11, 1, 0, 0, 0, 101, 102,
		3, 14, 7, 0, 102, 13, 1, 0, 0, 0, 103, 107, 3, 18, 9, 0, 104, 106, 3, 16,
		8, 0, 105, 104, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0,
		107, 108, 1, 0, 0, 0, 108, 15, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 111,
		4, 8, 2, 0, 111, 114, 5, 7, 0, 0, 112, 115, 3, 18, 9, 0, 113, 115, 3, 42,
		21, 0, 114, 112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 17, 1, 0, 0, 0,
		116, 120, 3, 22, 11, 0, 117, 119, 3, 20, 10, 0, 118, 117, 1, 0, 0, 0, 119,
		122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 19, 1,
		0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 4, 10, 3, 0, 124, 127, 5, 6, 0,
		0, 125, 128, 3, 22, 11, 0, 126, 128, 3, 42, 21, 0, 127, 125, 1, 0, 0, 0,
		127, 126, 1, 0, 0, 0, 128, 21, 1, 0, 0, 0, 129, 132, 7, 1, 0, 0, 130, 133,
		3, 22, 11, 0, 131, 133, 3, 42, 21, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1,
		0, 0, 0, 133, 136, 1, 0, 0, 0, 134, 136, 3, 24, 12, 0, 135, 129, 1, 0,
		0, 0, 135, 134, 1, 0, 0, 0, 136, 23, 1, 0, 0, 0, 137, 141, 3, 28, 14, 0,
		138, 140, 3, 26, 13, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 25, 1, 0, 0, 0, 143, 141, 1,
		0, 0, 0, 144, 145, 4, 13, 4, 0, 145, 148, 3, 56, 28, 0, 146, 149, 3, 28,
		14, 0, 147, 149, 3, 42, 21, 0, 148, 146, 1, 0, 0, 0, 148, 147, 1, 0, 0,
		0, 149, 27, 1, 0, 0, 0, 150, 154, 3, 32, 16, 0, 151, 153, 3, 30, 15, 0,
		152, 151, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 29, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158, 4,
		15, 5, 0, 158, 161, 3, 58, 29, 0, 159, 162, 3, 32, 16, 0, 160, 162, 3,
		42, 21, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162, 31, 1, 0, 0,
		0, 163, 167, 3, 36, 18, 0, 164, 166, 3, 34, 17, 0, 165, 164, 1, 0, 0, 0,
		166, 169, 1, 0, 0, 0, 167, 165, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168,
		33, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 170, 171, 4, 17, 6, 0, 171, 174,
		3, 60, 30, 0, 172, 175, 3, 36, 18, 0, 173, 175, 3, 42, 21, 0, 174, 172,
		1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175, 35, 1, 0, 0, 0, 176, 179, 3, 62,
		31, 0, 177, 180, 3, 36, 18, 0, 178, 180, 3, 42, 21, 0, 179, 177, 1, 0,
		0, 0, 179, 178, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 183, 3, 38, 19,
		0, 182, 176, 1, 0, 0, 0, 182, 181, 1, 0, 0, 0, 183, 37, 1, 0, 0, 0, 184,
		185, 5, 23, 0, 0, 185, 186, 3, 12, 6, 0, 186, 187, 5, 24, 0, 0, 187, 209,
		1, 0, 0, 0, 188, 189, 4, 19, 7, 0, 189, 190, 3, 64, 32, 0, 190, 199, 5,
		23, 0, 0, 191, 196, 3, 12, 6, 0, 192, 193, 5, 22, 0, 0, 193, 195, 3, 12,
		6, 0, 194, 192, 1, 0, 0, 0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0,
		196, 197, 1, 0, 0, 0, 197, 200, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199,
		191, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202,
		5, 24, 0, 0, 202, 209, 1, 0, 0, 0, 203, 209, 3, 64, 32, 0, 204, 209, 5,
		11, 0, 0, 205, 209, 5, 10, 0, 0, 206, 209, 5, 9, 0, 0, 207, 209, 3, 40,
		20, 0, 208, 184, 1, 0, 0, 0, 208, 188, 1, 0, 0, 0, 208, 203, 1, 0, 0, 0,
		208, 204, 1, 0, 0, 0, 208, 205, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208,
		207, 1, 0, 0, 0, 209, 39, 1, 0, 0, 0, 210, 211, 5, 25, 0, 0, 211, 215,
		6, 20, -1, 0, 212, 214, 8, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1,
		0, 0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 41, 1, 0, 0,
		0, 217, 215, 1, 0, 0, 0, 218, 219, 6, 21, -1, 0, 219, 220, 7, 2, 0, 0,
		220, 43, 1, 0, 0, 0, 221, 226, 3, 46, 23, 0, 222, 223, 5, 22, 0, 0, 223,
		225, 3, 46, 23, 0, 224, 222, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224,
		1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 45, 1, 0, 0, 0, 228, 226, 1, 0,
		0, 0, 229, 230, 3, 64, 32, 0, 230, 232, 5, 23, 0, 0, 231, 233, 3, 12, 6,
		0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		236, 5, 24, 0, 0, 235, 237, 3, 48, 24, 0, 236, 235, 1, 0, 0, 0, 236, 237,
		1, 0, 0, 0, 237, 47, 1, 0, 0, 0, 238, 239, 5, 5, 0, 0, 239, 240, 3, 64,
		32, 0, 240, 49, 1, 0, 0, 0, 241, 242, 5, 4, 0, 0, 242, 243, 3, 52, 26,
		0, 243, 51, 1, 0, 0, 0, 244, 249, 3, 54, 27, 0, 245, 246, 5, 22, 0, 0,
		246, 248, 3, 54, 27, 0, 247, 245, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249,
		247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 53, 1, 0, 0, 0, 251, 249, 1,
		0, 0, 0, 252, 253, 3, 64, 32, 0, 253, 55, 1, 0, 0, 0, 254, 255, 7, 3, 0,
		0, 255, 57, 1, 0, 0, 0, 256, 257, 7, 4, 0, 0, 257, 59, 1, 0, 0, 0, 258,
		259, 7, 5, 0, 0, 259, 61, 1, 0, 0, 0, 260, 261, 7, 4, 0, 0, 261, 63, 1,
		0, 0, 0, 262, 263, 7, 6, 0, 0, 263, 65, 1, 0, 0, 0, 26, 77, 87, 92, 99,
		107, 114, 120, 127, 132, 135, 141, 148, 154, 161, 167, 174, 179, 182, 196,
		199, 208, 215, 226, 232, 236, 249,
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

// SPL2ParserInit initializes any static state used to implement SPL2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSPL2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SPL2ParserInit() {
	staticData := &SPL2ParserStaticData
	staticData.once.Do(spl2ParserInit)
}

// NewSPL2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewSPL2Parser(input antlr.TokenStream) *SPL2Parser {
	SPL2ParserInit()
	this := new(SPL2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SPL2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SPL2.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

// la returns the token type at lookahead k.
func (p *SPL2Parser) la(k int) int { return p.GetTokenStream().LA(k) }

// tok returns the token at lookahead k; k may be negative.
func (p *SPL2Parser) tok(k int) antlr.Token { return p.GetTokenStream().LT(k) }

// at reports whether the current token has type t.
func (p *SPL2Parser) at(t int) bool { return p.la(1) == t }

// isIdentLike reports whether token type t is accepted by identLike.
func (p *SPL2Parser) isIdentLike(t int) bool {
	switch t {
	case SPL2ParserIDENT, SPL2ParserFROM, SPL2ParserWHERE, SPL2ParserSTATS,
		SPL2ParserBY, SPL2ParserAS, SPL2ParserAND, SPL2ParserOR, SPL2ParserNOT:
		return true
	}
	return false
}

// touches reports whether the tokens at k and k+1 are adjacent in the source.
// Token offsets are rune indexes, so this is exact.
func (p *SPL2Parser) touches(k int) bool {
	a, b := p.tok(k), p.tok(k+1)
	if a == nil || b == nil {
		return false
	}
	return a.GetStop()+1 == b.GetStart()
}

// touchingPrev reports whether the current token starts exactly where the
// last consumed token ended.
func (p *SPL2Parser) touchingPrev() bool {
	prev, cur := p.tok(-1), p.tok(1)
	if prev == nil || cur == nil {
		return false
	}
	return prev.GetStop()+1 == cur.GetStart()
}

// hyphenRun reports whether a '-' continues a dataset name: it must touch the
// previous component and be followed, without a space, by a word, number, or
// wildcard. Whitespace therefore ends the name, and `from logs- *` is a
// syntax error at the '-'.
func (p *SPL2Parser) hyphenRun() bool {
	return p.touchingPrev() && p.touches(1) &&
		(p.isIdentLike(p.la(2)) || p.la(2) == SPL2ParserNUMBER || p.la(2) == SPL2ParserSTAR)
}

// callAhead reports whether an identifier is followed by '('. Adjacency is
// not required: `where lower (msg)="x"` is a call.
func (p *SPL2Parser) callAhead() bool { return p.la(2) == SPL2ParserLPAREN }

// atNot, atCompare, atAdd, and atMul report whether the current token is an
// operator of that kind. The binary and unary rules gate on them so the parser
// commits to the operator and blames the missing operand after it.
func (p *SPL2Parser) atNot() bool { return p.at(SPL2ParserNOT) || p.at(SPL2ParserBANG) }

func (p *SPL2Parser) atCompare() bool {
	switch p.la(1) {
	case SPL2ParserEQ2, SPL2ParserEQ, SPL2ParserNE, SPL2ParserLT, SPL2ParserLE,
		SPL2ParserGT, SPL2ParserGE:
		return true
	}
	return false
}

func (p *SPL2Parser) atAdd() bool { return p.at(SPL2ParserPLUS) || p.at(SPL2ParserMINUS) }

func (p *SPL2Parser) atMul() bool {
	return p.at(SPL2ParserSTAR) || p.at(SPL2ParserSLASH) || p.at(SPL2ParserPERCENT)
}

// parseAbort is the panic value for a grammar action that ends the parse at a
// token. The front-end recovers it structurally and maps the token through its
// SourceMap; generated code cannot import the front-end, so the methods below
// are the whole contract.
type parseAbort struct {
	token   antlr.Token
	code    string
	message string
}

func (a parseAbort) AbortToken() antlr.Token { return a.token }
func (a parseAbort) AbortCode() string       { return a.code }
func (a parseAbort) AbortMessage() string    { return a.message }

// abortMissingOperand ends the parse at the token where an operand was
// required.
func (p *SPL2Parser) abortMissingOperand() {
	t := p.tok(1)
	message := "expected an expression, found end of query"
	if t != nil && t.GetTokenType() != antlr.TokenEOF {
		message = "expected an expression, found " + fmt.Sprintf("%q", t.GetText())
	}
	panic(parseAbort{token: t, code: "syntax", message: message})
}

// abortSubsearch refuses a subsearch at its opening bracket, before the rest of
// the stage is parsed.
func (p *SPL2Parser) abortSubsearch() {
	panic(parseAbort{
		token:   p.tok(-1),
		code:    "unsupported-subsearch",
		message: "subsearches are not supported",
	})
}

// SPL2Parser tokens.
const (
	SPL2ParserEOF      = antlr.TokenEOF
	SPL2ParserFROM     = 1
	SPL2ParserWHERE    = 2
	SPL2ParserSTATS    = 3
	SPL2ParserBY       = 4
	SPL2ParserAS       = 5
	SPL2ParserAND      = 6
	SPL2ParserOR       = 7
	SPL2ParserNOT      = 8
	SPL2ParserDURATION = 9
	SPL2ParserNUMBER   = 10
	SPL2ParserSTRING   = 11
	SPL2ParserIDENT    = 12
	SPL2ParserEQ2      = 13
	SPL2ParserNE       = 14
	SPL2ParserLE       = 15
	SPL2ParserGE       = 16
	SPL2ParserEQ       = 17
	SPL2ParserLT       = 18
	SPL2ParserGT       = 19
	SPL2ParserBANG     = 20
	SPL2ParserPIPE     = 21
	SPL2ParserCOMMA    = 22
	SPL2ParserLPAREN   = 23
	SPL2ParserRPAREN   = 24
	SPL2ParserLBRACKET = 25
	SPL2ParserRBRACKET = 26
	SPL2ParserPLUS     = 27
	SPL2ParserMINUS    = 28
	SPL2ParserSTAR     = 29
	SPL2ParserSLASH    = 30
	SPL2ParserPERCENT  = 31
	SPL2ParserWS       = 32
	SPL2ParserUNKNOWN  = 33
)

// SPL2Parser rules.
const (
	SPL2ParserRULE_stageFrom          = 0
	SPL2ParserRULE_stageWhere         = 1
	SPL2ParserRULE_stageStats         = 2
	SPL2ParserRULE_stageEnd           = 3
	SPL2ParserRULE_indexName          = 4
	SPL2ParserRULE_indexTail          = 5
	SPL2ParserRULE_expr               = 6
	SPL2ParserRULE_orExpr             = 7
	SPL2ParserRULE_orTail             = 8
	SPL2ParserRULE_andExpr            = 9
	SPL2ParserRULE_andTail            = 10
	SPL2ParserRULE_notExpr            = 11
	SPL2ParserRULE_comparison         = 12
	SPL2ParserRULE_comparisonTail     = 13
	SPL2ParserRULE_additive           = 14
	SPL2ParserRULE_additiveTail       = 15
	SPL2ParserRULE_multiplicative     = 16
	SPL2ParserRULE_multiplicativeTail = 17
	SPL2ParserRULE_unary              = 18
	SPL2ParserRULE_primary            = 19
	SPL2ParserRULE_subsearch          = 20
	SPL2ParserRULE_absentOperand      = 21
	SPL2ParserRULE_aggList            = 22
	SPL2ParserRULE_agg                = 23
	SPL2ParserRULE_aliasClause        = 24
	SPL2ParserRULE_byClause           = 25
	SPL2ParserRULE_fieldList          = 26
	SPL2ParserRULE_field              = 27
	SPL2ParserRULE_compOp             = 28
	SPL2ParserRULE_addOp              = 29
	SPL2ParserRULE_mulOp              = 30
	SPL2ParserRULE_sign               = 31
	SPL2ParserRULE_identLike          = 32
)

// IStageFromContext is an interface to support dynamic dispatch.
type IStageFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	IndexName() IIndexNameContext
	StageEnd() IStageEndContext

	// IsStageFromContext differentiates from other interfaces.
	IsStageFromContext()
}

type StageFromContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageFromContext() *StageFromContext {
	var p = new(StageFromContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageFrom
	return p
}

func InitEmptyStageFromContext(p *StageFromContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageFrom
}

func (*StageFromContext) IsStageFromContext() {}

func NewStageFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageFromContext {
	var p = new(StageFromContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_stageFrom

	return p
}

func (s *StageFromContext) GetParser() antlr.Parser { return s.parser }

func (s *StageFromContext) FROM() antlr.TerminalNode {
	return s.GetToken(SPL2ParserFROM, 0)
}

func (s *StageFromContext) IndexName() IIndexNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexNameContext)
}

func (s *StageFromContext) StageEnd() IStageEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStageEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStageEndContext)
}

func (s *StageFromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageFromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageFromContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitStageFrom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) StageFrom() (localctx IStageFromContext) {
	localctx = NewStageFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SPL2ParserRULE_stageFrom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(SPL2ParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.IndexName()
	}
	{
		p.SetState(68)
		p.StageEnd()
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

// IStageWhereContext is an interface to support dynamic dispatch.
type IStageWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	Expr() IExprContext
	StageEnd() IStageEndContext

	// IsStageWhereContext differentiates from other interfaces.
	IsStageWhereContext()
}

type StageWhereContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageWhereContext() *StageWhereContext {
	var p = new(StageWhereContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageWhere
	return p
}

func InitEmptyStageWhereContext(p *StageWhereContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageWhere
}

func (*StageWhereContext) IsStageWhereContext() {}

func NewStageWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageWhereContext {
	var p = new(StageWhereContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_stageWhere

	return p
}

func (s *StageWhereContext) GetParser() antlr.Parser { return s.parser }

func (s *StageWhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserWHERE, 0)
}

func (s *StageWhereContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StageWhereContext) StageEnd() IStageEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStageEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStageEndContext)
}

func (s *StageWhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageWhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageWhereContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitStageWhere(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) StageWhere() (localctx IStageWhereContext) {
	localctx = NewStageWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SPL2ParserRULE_stageWhere)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(SPL2ParserWHERE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Expr()
	}
	{
		p.SetState(72)
		p.StageEnd()
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

// IStageStatsContext is an interface to support dynamic dispatch.
type IStageStatsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STATS() antlr.TerminalNode
	AggList() IAggListContext
	StageEnd() IStageEndContext
	ByClause() IByClauseContext

	// IsStageStatsContext differentiates from other interfaces.
	IsStageStatsContext()
}

type StageStatsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageStatsContext() *StageStatsContext {
	var p = new(StageStatsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageStats
	return p
}

func InitEmptyStageStatsContext(p *StageStatsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageStats
}

func (*StageStatsContext) IsStageStatsContext() {}

func NewStageStatsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageStatsContext {
	var p = new(StageStatsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_stageStats

	return p
}

func (s *StageStatsContext) GetParser() antlr.Parser { return s.parser }

func (s *StageStatsContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTATS, 0)
}

func (s *StageStatsContext) AggList() IAggListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggListContext)
}

func (s *StageStatsContext) StageEnd() IStageEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStageEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStageEndContext)
}

func (s *StageStatsContext) ByClause() IByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IByClauseContext)
}

func (s *StageStatsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageStatsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageStatsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitStageStats(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) StageStats() (localctx IStageStatsContext) {
	localctx = NewStageStatsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SPL2ParserRULE_stageStats)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(SPL2ParserSTATS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)
		p.AggList()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL2ParserBY {
		{
			p.SetState(76)
			p.ByClause()
		}

	}
	{
		p.SetState(79)
		p.StageEnd()
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

// IStageEndContext is an interface to support dynamic dispatch.
type IStageEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PIPE() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsStageEndContext differentiates from other interfaces.
	IsStageEndContext()
}

type StageEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageEndContext() *StageEndContext {
	var p = new(StageEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageEnd
	return p
}

func InitEmptyStageEndContext(p *StageEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_stageEnd
}

func (*StageEndContext) IsStageEndContext() {}

func NewStageEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageEndContext {
	var p = new(StageEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_stageEnd

	return p
}

func (s *StageEndContext) GetParser() antlr.Parser { return s.parser }

func (s *StageEndContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserPIPE, 0)
}

func (s *StageEndContext) EOF() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEOF, 0)
}

func (s *StageEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitStageEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) StageEnd() (localctx IStageEndContext) {
	localctx = NewStageEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SPL2ParserRULE_stageEnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL2ParserEOF || _la == SPL2ParserPIPE) {
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

// IIndexNameContext is an interface to support dynamic dispatch.
type IIndexNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext
	AllIndexTail() []IIndexTailContext
	IndexTail(i int) IIndexTailContext
	STAR() antlr.TerminalNode

	// IsIndexNameContext differentiates from other interfaces.
	IsIndexNameContext()
}

type IndexNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexNameContext() *IndexNameContext {
	var p = new(IndexNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_indexName
	return p
}

func InitEmptyIndexNameContext(p *IndexNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_indexName
}

func (*IndexNameContext) IsIndexNameContext() {}

func NewIndexNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexNameContext {
	var p = new(IndexNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_indexName

	return p
}

func (s *IndexNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexNameContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *IndexNameContext) AllIndexTail() []IIndexTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexTailContext); ok {
			len++
		}
	}

	tst := make([]IIndexTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexTailContext); ok {
			tst[i] = t.(IIndexTailContext)
			i++
		}
	}

	return tst
}

func (s *IndexNameContext) IndexTail(i int) IIndexTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexTailContext)
}

func (s *IndexNameContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTAR, 0)
}

func (s *IndexNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitIndexName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) IndexName() (localctx IIndexNameContext) {
	localctx = NewIndexNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SPL2ParserRULE_indexName)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.IdentLike()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.IndexTail()
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		p.SetState(90)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(91)
			p.Match(SPL2ParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
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

// IIndexTailContext is an interface to support dynamic dispatch.
type IIndexTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	IdentLike() IIdentLikeContext
	NUMBER() antlr.TerminalNode
	STAR() antlr.TerminalNode

	// IsIndexTailContext differentiates from other interfaces.
	IsIndexTailContext()
}

type IndexTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexTailContext() *IndexTailContext {
	var p = new(IndexTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_indexTail
	return p
}

func InitEmptyIndexTailContext(p *IndexTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_indexTail
}

func (*IndexTailContext) IsIndexTailContext() {}

func NewIndexTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexTailContext {
	var p = new(IndexTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_indexTail

	return p
}

func (s *IndexTailContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexTailContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserMINUS, 0)
}

func (s *IndexTailContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *IndexTailContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNUMBER, 0)
}

func (s *IndexTailContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTAR, 0)
}

func (s *IndexTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitIndexTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) IndexTail() (localctx IIndexTailContext) {
	localctx = NewIndexTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SPL2ParserRULE_indexTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(94)

	if !(p.hyphenRun()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.hyphenRun()", ""))
		goto errorExit
	}
	{
		p.SetState(95)
		p.Match(SPL2ParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL2ParserFROM, SPL2ParserWHERE, SPL2ParserSTATS, SPL2ParserBY, SPL2ParserAS, SPL2ParserAND, SPL2ParserOR, SPL2ParserNOT, SPL2ParserIDENT:
		{
			p.SetState(96)
			p.IdentLike()
		}

	case SPL2ParserNUMBER:
		{
			p.SetState(97)
			p.Match(SPL2ParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL2ParserSTAR:
		{
			p.SetState(98)
			p.Match(SPL2ParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrExpr() IOrExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) OrExpr() IOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SPL2ParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.OrExpr()
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

// IOrExprContext is an interface to support dynamic dispatch.
type IOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AndExpr() IAndExprContext
	AllOrTail() []IOrTailContext
	OrTail(i int) IOrTailContext

	// IsOrExprContext differentiates from other interfaces.
	IsOrExprContext()
}

type OrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrExprContext() *OrExprContext {
	var p = new(OrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_orExpr
	return p
}

func InitEmptyOrExprContext(p *OrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_orExpr
}

func (*OrExprContext) IsOrExprContext() {}

func NewOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrExprContext {
	var p = new(OrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_orExpr

	return p
}

func (s *OrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrExprContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrExprContext) AllOrTail() []IOrTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrTailContext); ok {
			len++
		}
	}

	tst := make([]IOrTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrTailContext); ok {
			tst[i] = t.(IOrTailContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) OrTail(i int) IOrTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrTailContext)
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) OrExpr() (localctx IOrExprContext) {
	localctx = NewOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SPL2ParserRULE_orExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.AndExpr()
	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(104)
				p.OrTail()
			}

		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
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

// IOrTailContext is an interface to support dynamic dispatch.
type IOrTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode
	AndExpr() IAndExprContext
	AbsentOperand() IAbsentOperandContext

	// IsOrTailContext differentiates from other interfaces.
	IsOrTailContext()
}

type OrTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrTailContext() *OrTailContext {
	var p = new(OrTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_orTail
	return p
}

func InitEmptyOrTailContext(p *OrTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_orTail
}

func (*OrTailContext) IsOrTailContext() {}

func NewOrTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrTailContext {
	var p = new(OrTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_orTail

	return p
}

func (s *OrTailContext) GetParser() antlr.Parser { return s.parser }

func (s *OrTailContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL2ParserOR, 0)
}

func (s *OrTailContext) AndExpr() IAndExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndExprContext)
}

func (s *OrTailContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *OrTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitOrTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) OrTail() (localctx IOrTailContext) {
	localctx = NewOrTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SPL2ParserRULE_orTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(110)

	if !(p.at(SPL2ParserOR)) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL2ParserOR)", ""))
		goto errorExit
	}
	{
		p.SetState(111)
		p.Match(SPL2ParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(112)
			p.AndExpr()
		}

	case 2:
		{
			p.SetState(113)
			p.AbsentOperand()
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

// IAndExprContext is an interface to support dynamic dispatch.
type IAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NotExpr() INotExprContext
	AllAndTail() []IAndTailContext
	AndTail(i int) IAndTailContext

	// IsAndExprContext differentiates from other interfaces.
	IsAndExprContext()
}

type AndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndExprContext() *AndExprContext {
	var p = new(AndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_andExpr
	return p
}

func InitEmptyAndExprContext(p *AndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_andExpr
}

func (*AndExprContext) IsAndExprContext() {}

func NewAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndExprContext {
	var p = new(AndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_andExpr

	return p
}

func (s *AndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AndExprContext) NotExpr() INotExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotExprContext)
}

func (s *AndExprContext) AllAndTail() []IAndTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndTailContext); ok {
			len++
		}
	}

	tst := make([]IAndTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndTailContext); ok {
			tst[i] = t.(IAndTailContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) AndTail(i int) IAndTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndTailContext)
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AndExpr() (localctx IAndExprContext) {
	localctx = NewAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SPL2ParserRULE_andExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.NotExpr()
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(117)
				p.AndTail()
			}

		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
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

// IAndTailContext is an interface to support dynamic dispatch.
type IAndTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode
	NotExpr() INotExprContext
	AbsentOperand() IAbsentOperandContext

	// IsAndTailContext differentiates from other interfaces.
	IsAndTailContext()
}

type AndTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndTailContext() *AndTailContext {
	var p = new(AndTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_andTail
	return p
}

func InitEmptyAndTailContext(p *AndTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_andTail
}

func (*AndTailContext) IsAndTailContext() {}

func NewAndTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndTailContext {
	var p = new(AndTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_andTail

	return p
}

func (s *AndTailContext) GetParser() antlr.Parser { return s.parser }

func (s *AndTailContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL2ParserAND, 0)
}

func (s *AndTailContext) NotExpr() INotExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotExprContext)
}

func (s *AndTailContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *AndTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAndTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AndTail() (localctx IAndTailContext) {
	localctx = NewAndTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SPL2ParserRULE_andTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(123)

	if !(p.at(SPL2ParserAND)) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL2ParserAND)", ""))
		goto errorExit
	}
	{
		p.SetState(124)
		p.Match(SPL2ParserAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(125)
			p.NotExpr()
		}

	case 2:
		{
			p.SetState(126)
			p.AbsentOperand()
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

// INotExprContext is an interface to support dynamic dispatch.
type INotExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode
	BANG() antlr.TerminalNode
	NotExpr() INotExprContext
	AbsentOperand() IAbsentOperandContext
	Comparison() IComparisonContext

	// IsNotExprContext differentiates from other interfaces.
	IsNotExprContext()
}

type NotExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotExprContext() *NotExprContext {
	var p = new(NotExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_notExpr
	return p
}

func InitEmptyNotExprContext(p *NotExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_notExpr
}

func (*NotExprContext) IsNotExprContext() {}

func NewNotExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotExprContext {
	var p = new(NotExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_notExpr

	return p
}

func (s *NotExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NotExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNOT, 0)
}

func (s *NotExprContext) BANG() antlr.TerminalNode {
	return s.GetToken(SPL2ParserBANG, 0)
}

func (s *NotExprContext) NotExpr() INotExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotExprContext)
}

func (s *NotExprContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *NotExprContext) Comparison() IComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) NotExpr() (localctx INotExprContext) {
	localctx = NewNotExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SPL2ParserRULE_notExpr)
	var _la int

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SPL2ParserNOT || _la == SPL2ParserBANG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(130)
				p.NotExpr()
			}

		case 2:
			{
				p.SetState(131)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(134)
			p.Comparison()
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

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Additive() IAdditiveContext
	AllComparisonTail() []IComparisonTailContext
	ComparisonTail(i int) IComparisonTailContext

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *ComparisonContext) AllComparisonTail() []IComparisonTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonTailContext); ok {
			len++
		}
	}

	tst := make([]IComparisonTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonTailContext); ok {
			tst[i] = t.(IComparisonTailContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonContext) ComparisonTail(i int) IComparisonTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonTailContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SPL2ParserRULE_comparison)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		p.Additive()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(138)
				p.ComparisonTail()
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
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

// IComparisonTailContext is an interface to support dynamic dispatch.
type IComparisonTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CompOp() ICompOpContext
	Additive() IAdditiveContext
	AbsentOperand() IAbsentOperandContext

	// IsComparisonTailContext differentiates from other interfaces.
	IsComparisonTailContext()
}

type ComparisonTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonTailContext() *ComparisonTailContext {
	var p = new(ComparisonTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_comparisonTail
	return p
}

func InitEmptyComparisonTailContext(p *ComparisonTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_comparisonTail
}

func (*ComparisonTailContext) IsComparisonTailContext() {}

func NewComparisonTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonTailContext {
	var p = new(ComparisonTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_comparisonTail

	return p
}

func (s *ComparisonTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonTailContext) CompOp() ICompOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompOpContext)
}

func (s *ComparisonTailContext) Additive() IAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveContext)
}

func (s *ComparisonTailContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *ComparisonTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitComparisonTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) ComparisonTail() (localctx IComparisonTailContext) {
	localctx = NewComparisonTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SPL2ParserRULE_comparisonTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(144)

	if !(p.atCompare()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atCompare()", ""))
		goto errorExit
	}
	{
		p.SetState(145)
		p.CompOp()
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(146)
			p.Additive()
		}

	case 2:
		{
			p.SetState(147)
			p.AbsentOperand()
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

// IAdditiveContext is an interface to support dynamic dispatch.
type IAdditiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Multiplicative() IMultiplicativeContext
	AllAdditiveTail() []IAdditiveTailContext
	AdditiveTail(i int) IAdditiveTailContext

	// IsAdditiveContext differentiates from other interfaces.
	IsAdditiveContext()
}

type AdditiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveContext() *AdditiveContext {
	var p = new(AdditiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_additive
	return p
}

func InitEmptyAdditiveContext(p *AdditiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_additive
}

func (*AdditiveContext) IsAdditiveContext() {}

func NewAdditiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveContext {
	var p = new(AdditiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_additive

	return p
}

func (s *AdditiveContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveContext) Multiplicative() IMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeContext)
}

func (s *AdditiveContext) AllAdditiveTail() []IAdditiveTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveTailContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveTailContext); ok {
			tst[i] = t.(IAdditiveTailContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveContext) AdditiveTail(i int) IAdditiveTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveTailContext)
}

func (s *AdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAdditive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Additive() (localctx IAdditiveContext) {
	localctx = NewAdditiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SPL2ParserRULE_additive)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Multiplicative()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(151)
				p.AdditiveTail()
			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
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

// IAdditiveTailContext is an interface to support dynamic dispatch.
type IAdditiveTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddOp() IAddOpContext
	Multiplicative() IMultiplicativeContext
	AbsentOperand() IAbsentOperandContext

	// IsAdditiveTailContext differentiates from other interfaces.
	IsAdditiveTailContext()
}

type AdditiveTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveTailContext() *AdditiveTailContext {
	var p = new(AdditiveTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_additiveTail
	return p
}

func InitEmptyAdditiveTailContext(p *AdditiveTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_additiveTail
}

func (*AdditiveTailContext) IsAdditiveTailContext() {}

func NewAdditiveTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveTailContext {
	var p = new(AdditiveTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_additiveTail

	return p
}

func (s *AdditiveTailContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveTailContext) AddOp() IAddOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}

func (s *AdditiveTailContext) Multiplicative() IMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeContext)
}

func (s *AdditiveTailContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *AdditiveTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAdditiveTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AdditiveTail() (localctx IAdditiveTailContext) {
	localctx = NewAdditiveTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SPL2ParserRULE_additiveTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(157)

	if !(p.atAdd()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atAdd()", ""))
		goto errorExit
	}
	{
		p.SetState(158)
		p.AddOp()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(159)
			p.Multiplicative()
		}

	case 2:
		{
			p.SetState(160)
			p.AbsentOperand()
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

// IMultiplicativeContext is an interface to support dynamic dispatch.
type IMultiplicativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Unary() IUnaryContext
	AllMultiplicativeTail() []IMultiplicativeTailContext
	MultiplicativeTail(i int) IMultiplicativeTailContext

	// IsMultiplicativeContext differentiates from other interfaces.
	IsMultiplicativeContext()
}

type MultiplicativeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeContext() *MultiplicativeContext {
	var p = new(MultiplicativeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_multiplicative
	return p
}

func InitEmptyMultiplicativeContext(p *MultiplicativeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_multiplicative
}

func (*MultiplicativeContext) IsMultiplicativeContext() {}

func NewMultiplicativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeContext {
	var p = new(MultiplicativeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_multiplicative

	return p
}

func (s *MultiplicativeContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *MultiplicativeContext) AllMultiplicativeTail() []IMultiplicativeTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeTailContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeTailContext); ok {
			tst[i] = t.(IMultiplicativeTailContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeContext) MultiplicativeTail(i int) IMultiplicativeTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeTailContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeTailContext)
}

func (s *MultiplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitMultiplicative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Multiplicative() (localctx IMultiplicativeContext) {
	localctx = NewMultiplicativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SPL2ParserRULE_multiplicative)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Unary()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(164)
				p.MultiplicativeTail()
			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
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

// IMultiplicativeTailContext is an interface to support dynamic dispatch.
type IMultiplicativeTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MulOp() IMulOpContext
	Unary() IUnaryContext
	AbsentOperand() IAbsentOperandContext

	// IsMultiplicativeTailContext differentiates from other interfaces.
	IsMultiplicativeTailContext()
}

type MultiplicativeTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeTailContext() *MultiplicativeTailContext {
	var p = new(MultiplicativeTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_multiplicativeTail
	return p
}

func InitEmptyMultiplicativeTailContext(p *MultiplicativeTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_multiplicativeTail
}

func (*MultiplicativeTailContext) IsMultiplicativeTailContext() {}

func NewMultiplicativeTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeTailContext {
	var p = new(MultiplicativeTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_multiplicativeTail

	return p
}

func (s *MultiplicativeTailContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeTailContext) MulOp() IMulOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOpContext)
}

func (s *MultiplicativeTailContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *MultiplicativeTailContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *MultiplicativeTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitMultiplicativeTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) MultiplicativeTail() (localctx IMultiplicativeTailContext) {
	localctx = NewMultiplicativeTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SPL2ParserRULE_multiplicativeTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)

	if !(p.atMul()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atMul()", ""))
		goto errorExit
	}
	{
		p.SetState(171)
		p.MulOp()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(172)
			p.Unary()
		}

	case 2:
		{
			p.SetState(173)
			p.AbsentOperand()
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

// IUnaryContext is an interface to support dynamic dispatch.
type IUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Sign() ISignContext
	Unary() IUnaryContext
	AbsentOperand() IAbsentOperandContext
	Primary() IPrimaryContext

	// IsUnaryContext differentiates from other interfaces.
	IsUnaryContext()
}

type UnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryContext() *UnaryContext {
	var p = new(UnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_unary
	return p
}

func InitEmptyUnaryContext(p *UnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_unary
}

func (*UnaryContext) IsUnaryContext() {}

func NewUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryContext {
	var p = new(UnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_unary

	return p
}

func (s *UnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryContext) Sign() ISignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISignContext)
}

func (s *UnaryContext) Unary() IUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryContext)
}

func (s *UnaryContext) AbsentOperand() IAbsentOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsentOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsentOperandContext)
}

func (s *UnaryContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Unary() (localctx IUnaryContext) {
	localctx = NewUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SPL2ParserRULE_unary)
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Sign()
		}
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(177)
				p.Unary()
			}

		case 2:
			{
				p.SetState(178)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Primary()
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

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	RPAREN() antlr.TerminalNode
	IdentLike() IIdentLikeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	Subsearch() ISubsearchContext

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLPAREN, 0)
}

func (s *PrimaryContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrimaryContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL2ParserRPAREN, 0)
}

func (s *PrimaryContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *PrimaryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL2ParserCOMMA)
}

func (s *PrimaryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL2ParserCOMMA, i)
}

func (s *PrimaryContext) STRING() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTRING, 0)
}

func (s *PrimaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNUMBER, 0)
}

func (s *PrimaryContext) DURATION() antlr.TerminalNode {
	return s.GetToken(SPL2ParserDURATION, 0)
}

func (s *PrimaryContext) Subsearch() ISubsearchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubsearchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubsearchContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SPL2ParserRULE_primary)
	var _la int

	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Match(SPL2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Expr()
		}
		{
			p.SetState(186)
			p.Match(SPL2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(188)

		if !(p.callAhead()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.callAhead()", ""))
			goto errorExit
		}
		{
			p.SetState(189)
			p.IdentLike()
		}
		{
			p.SetState(190)
			p.Match(SPL2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(191)
				p.Expr()
			}
			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == SPL2ParserCOMMA {
				{
					p.SetState(192)
					p.Match(SPL2ParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(193)
					p.Expr()
				}

				p.SetState(198)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(201)
			p.Match(SPL2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(203)
			p.IdentLike()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(204)
			p.Match(SPL2ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)
			p.Match(SPL2ParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(206)
			p.Match(SPL2ParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(207)
			p.Subsearch()
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

// ISubsearchContext is an interface to support dynamic dispatch.
type ISubsearchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	AllPIPE() []antlr.TerminalNode
	PIPE(i int) antlr.TerminalNode
	AllEOF() []antlr.TerminalNode
	EOF(i int) antlr.TerminalNode

	// IsSubsearchContext differentiates from other interfaces.
	IsSubsearchContext()
}

type SubsearchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubsearchContext() *SubsearchContext {
	var p = new(SubsearchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_subsearch
	return p
}

func InitEmptySubsearchContext(p *SubsearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_subsearch
}

func (*SubsearchContext) IsSubsearchContext() {}

func NewSubsearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubsearchContext {
	var p = new(SubsearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_subsearch

	return p
}

func (s *SubsearchContext) GetParser() antlr.Parser { return s.parser }

func (s *SubsearchContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLBRACKET, 0)
}

func (s *SubsearchContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(SPL2ParserPIPE)
}

func (s *SubsearchContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(SPL2ParserPIPE, i)
}

func (s *SubsearchContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(SPL2ParserEOF)
}

func (s *SubsearchContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(SPL2ParserEOF, i)
}

func (s *SubsearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubsearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubsearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitSubsearch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Subsearch() (localctx ISubsearchContext) {
	localctx = NewSubsearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SPL2ParserRULE_subsearch)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(SPL2ParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.abortSubsearch()
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(212)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == SPL2ParserEOF || _la == SPL2ParserPIPE {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
		if p.HasError() {
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

// IAbsentOperandContext is an interface to support dynamic dispatch.
type IAbsentOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ2() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	PIPE() antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsAbsentOperandContext differentiates from other interfaces.
	IsAbsentOperandContext()
}

type AbsentOperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsentOperandContext() *AbsentOperandContext {
	var p = new(AbsentOperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_absentOperand
	return p
}

func InitEmptyAbsentOperandContext(p *AbsentOperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_absentOperand
}

func (*AbsentOperandContext) IsAbsentOperandContext() {}

func NewAbsentOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsentOperandContext {
	var p = new(AbsentOperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_absentOperand

	return p
}

func (s *AbsentOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsentOperandContext) EQ2() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEQ2, 0)
}

func (s *AbsentOperandContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEQ, 0)
}

func (s *AbsentOperandContext) NE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNE, 0)
}

func (s *AbsentOperandContext) LT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLT, 0)
}

func (s *AbsentOperandContext) LE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLE, 0)
}

func (s *AbsentOperandContext) GT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserGT, 0)
}

func (s *AbsentOperandContext) GE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserGE, 0)
}

func (s *AbsentOperandContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL2ParserCOMMA, 0)
}

func (s *AbsentOperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL2ParserRPAREN, 0)
}

func (s *AbsentOperandContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SPL2ParserRBRACKET, 0)
}

func (s *AbsentOperandContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserPIPE, 0)
}

func (s *AbsentOperandContext) EOF() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEOF, 0)
}

func (s *AbsentOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsentOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsentOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAbsentOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AbsentOperand() (localctx IAbsentOperandContext) {
	localctx = NewAbsentOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SPL2ParserRULE_absentOperand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.abortMissingOperand()
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&182435841) != 0) {
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

// IAggListContext is an interface to support dynamic dispatch.
type IAggListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAgg() []IAggContext
	Agg(i int) IAggContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAggListContext differentiates from other interfaces.
	IsAggListContext()
}

type AggListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggListContext() *AggListContext {
	var p = new(AggListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_aggList
	return p
}

func InitEmptyAggListContext(p *AggListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_aggList
}

func (*AggListContext) IsAggListContext() {}

func NewAggListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggListContext {
	var p = new(AggListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_aggList

	return p
}

func (s *AggListContext) GetParser() antlr.Parser { return s.parser }

func (s *AggListContext) AllAgg() []IAggContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggContext); ok {
			len++
		}
	}

	tst := make([]IAggContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggContext); ok {
			tst[i] = t.(IAggContext)
			i++
		}
	}

	return tst
}

func (s *AggListContext) Agg(i int) IAggContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggContext)
}

func (s *AggListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL2ParserCOMMA)
}

func (s *AggListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL2ParserCOMMA, i)
}

func (s *AggListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAggList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AggList() (localctx IAggListContext) {
	localctx = NewAggListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SPL2ParserRULE_aggList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Agg()
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SPL2ParserCOMMA {
		{
			p.SetState(222)
			p.Match(SPL2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Agg()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAggContext is an interface to support dynamic dispatch.
type IAggContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Expr() IExprContext
	AliasClause() IAliasClauseContext

	// IsAggContext differentiates from other interfaces.
	IsAggContext()
}

type AggContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggContext() *AggContext {
	var p = new(AggContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_agg
	return p
}

func InitEmptyAggContext(p *AggContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_agg
}

func (*AggContext) IsAggContext() {}

func NewAggContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggContext {
	var p = new(AggContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_agg

	return p
}

func (s *AggContext) GetParser() antlr.Parser { return s.parser }

func (s *AggContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *AggContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLPAREN, 0)
}

func (s *AggContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL2ParserRPAREN, 0)
}

func (s *AggContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AggContext) AliasClause() IAliasClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasClauseContext)
}

func (s *AggContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAgg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Agg() (localctx IAggContext) {
	localctx = NewAggContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SPL2ParserRULE_agg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.IdentLike()
	}
	{
		p.SetState(230)
		p.Match(SPL2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(231)
			p.Expr()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(234)
		p.Match(SPL2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL2ParserAS {
		{
			p.SetState(235)
			p.AliasClause()
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

// IAliasClauseContext is an interface to support dynamic dispatch.
type IAliasClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AS() antlr.TerminalNode
	IdentLike() IIdentLikeContext

	// IsAliasClauseContext differentiates from other interfaces.
	IsAliasClauseContext()
}

type AliasClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasClauseContext() *AliasClauseContext {
	var p = new(AliasClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_aliasClause
	return p
}

func InitEmptyAliasClauseContext(p *AliasClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_aliasClause
}

func (*AliasClauseContext) IsAliasClauseContext() {}

func NewAliasClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasClauseContext {
	var p = new(AliasClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_aliasClause

	return p
}

func (s *AliasClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasClauseContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserAS, 0)
}

func (s *AliasClauseContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *AliasClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAliasClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AliasClause() (localctx IAliasClauseContext) {
	localctx = NewAliasClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SPL2ParserRULE_aliasClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(SPL2ParserAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.IdentLike()
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

// IByClauseContext is an interface to support dynamic dispatch.
type IByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BY() antlr.TerminalNode
	FieldList() IFieldListContext

	// IsByClauseContext differentiates from other interfaces.
	IsByClauseContext()
}

type ByClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyByClauseContext() *ByClauseContext {
	var p = new(ByClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_byClause
	return p
}

func InitEmptyByClauseContext(p *ByClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_byClause
}

func (*ByClauseContext) IsByClauseContext() {}

func NewByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByClauseContext {
	var p = new(ByClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_byClause

	return p
}

func (s *ByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL2ParserBY, 0)
}

func (s *ByClauseContext) FieldList() IFieldListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *ByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) ByClause() (localctx IByClauseContext) {
	localctx = NewByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SPL2ParserRULE_byClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(SPL2ParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.FieldList()
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

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllField() []IFieldContext
	Field(i int) IFieldContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_fieldList
	return p
}

func InitEmptyFieldListContext(p *FieldListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_fieldList
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllField() []IFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldContext); ok {
			len++
		}
	}

	tst := make([]IFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldContext); ok {
			tst[i] = t.(IFieldContext)
			i++
		}
	}

	return tst
}

func (s *FieldListContext) Field(i int) IFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL2ParserCOMMA)
}

func (s *FieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL2ParserCOMMA, i)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SPL2ParserRULE_fieldList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Field()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SPL2ParserCOMMA {
		{
			p.SetState(245)
			p.Match(SPL2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(246)
			p.Field()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) IdentLike() IIdentLikeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SPL2ParserRULE_field)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.IdentLike()
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

// ICompOpContext is an interface to support dynamic dispatch.
type ICompOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ2() antlr.TerminalNode
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	LT() antlr.TerminalNode
	LE() antlr.TerminalNode
	GT() antlr.TerminalNode
	GE() antlr.TerminalNode

	// IsCompOpContext differentiates from other interfaces.
	IsCompOpContext()
}

type CompOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompOpContext() *CompOpContext {
	var p = new(CompOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_compOp
	return p
}

func InitEmptyCompOpContext(p *CompOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_compOp
}

func (*CompOpContext) IsCompOpContext() {}

func NewCompOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompOpContext {
	var p = new(CompOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_compOp

	return p
}

func (s *CompOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompOpContext) EQ2() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEQ2, 0)
}

func (s *CompOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL2ParserEQ, 0)
}

func (s *CompOpContext) NE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNE, 0)
}

func (s *CompOpContext) LT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLT, 0)
}

func (s *CompOpContext) LE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserLE, 0)
}

func (s *CompOpContext) GT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserGT, 0)
}

func (s *CompOpContext) GE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserGE, 0)
}

func (s *CompOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitCompOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) CompOp() (localctx ICompOpContext) {
	localctx = NewCompOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SPL2ParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1040384) != 0) {
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

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserPLUS, 0)
}

func (s *AddOpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserMINUS, 0)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SPL2ParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL2ParserPLUS || _la == SPL2ParserMINUS) {
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

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PERCENT() antlr.TerminalNode

	// IsMulOpContext differentiates from other interfaces.
	IsMulOpContext()
}

type MulOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOpContext() *MulOpContext {
	var p = new(MulOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_mulOp
	return p
}

func InitEmptyMulOpContext(p *MulOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_mulOp
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MulOpContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTAR, 0)
}

func (s *MulOpContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSLASH, 0)
}

func (s *MulOpContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserPERCENT, 0)
}

func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) MulOp() (localctx IMulOpContext) {
	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SPL2ParserRULE_mulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3758096384) != 0) {
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

// ISignContext is an interface to support dynamic dispatch.
type ISignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsSignContext differentiates from other interfaces.
	IsSignContext()
}

type SignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySignContext() *SignContext {
	var p = new(SignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_sign
	return p
}

func InitEmptySignContext(p *SignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_sign
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserMINUS, 0)
}

func (s *SignContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserPLUS, 0)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) Sign() (localctx ISignContext) {
	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SPL2ParserRULE_sign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL2ParserPLUS || _la == SPL2ParserMINUS) {
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

// IIdentLikeContext is an interface to support dynamic dispatch.
type IIdentLikeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	FROM() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	STATS() antlr.TerminalNode
	BY() antlr.TerminalNode
	AS() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsIdentLikeContext differentiates from other interfaces.
	IsIdentLikeContext()
}

type IdentLikeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentLikeContext() *IdentLikeContext {
	var p = new(IdentLikeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_identLike
	return p
}

func InitEmptyIdentLikeContext(p *IdentLikeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL2ParserRULE_identLike
}

func (*IdentLikeContext) IsIdentLikeContext() {}

func NewIdentLikeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentLikeContext {
	var p = new(IdentLikeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL2ParserRULE_identLike

	return p
}

func (s *IdentLikeContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentLikeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserIDENT, 0)
}

func (s *IdentLikeContext) FROM() antlr.TerminalNode {
	return s.GetToken(SPL2ParserFROM, 0)
}

func (s *IdentLikeContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL2ParserWHERE, 0)
}

func (s *IdentLikeContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserSTATS, 0)
}

func (s *IdentLikeContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL2ParserBY, 0)
}

func (s *IdentLikeContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL2ParserAS, 0)
}

func (s *IdentLikeContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL2ParserAND, 0)
}

func (s *IdentLikeContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL2ParserOR, 0)
}

func (s *IdentLikeContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL2ParserNOT, 0)
}

func (s *IdentLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentLikeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentLikeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL2Visitor:
		return t.VisitIdentLike(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL2Parser) IdentLike() (localctx IIdentLikeContext) {
	localctx = NewIdentLikeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SPL2ParserRULE_identLike)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4606) != 0) {
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

func (p *SPL2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
		var t *IndexNameContext = nil
		if localctx != nil {
			t = localctx.(*IndexNameContext)
		}
		return p.IndexName_Sempred(t, predIndex)

	case 5:
		var t *IndexTailContext = nil
		if localctx != nil {
			t = localctx.(*IndexTailContext)
		}
		return p.IndexTail_Sempred(t, predIndex)

	case 8:
		var t *OrTailContext = nil
		if localctx != nil {
			t = localctx.(*OrTailContext)
		}
		return p.OrTail_Sempred(t, predIndex)

	case 10:
		var t *AndTailContext = nil
		if localctx != nil {
			t = localctx.(*AndTailContext)
		}
		return p.AndTail_Sempred(t, predIndex)

	case 13:
		var t *ComparisonTailContext = nil
		if localctx != nil {
			t = localctx.(*ComparisonTailContext)
		}
		return p.ComparisonTail_Sempred(t, predIndex)

	case 15:
		var t *AdditiveTailContext = nil
		if localctx != nil {
			t = localctx.(*AdditiveTailContext)
		}
		return p.AdditiveTail_Sempred(t, predIndex)

	case 17:
		var t *MultiplicativeTailContext = nil
		if localctx != nil {
			t = localctx.(*MultiplicativeTailContext)
		}
		return p.MultiplicativeTail_Sempred(t, predIndex)

	case 19:
		var t *PrimaryContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryContext)
		}
		return p.Primary_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SPL2Parser) IndexName_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.touchingPrev()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) IndexTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.hyphenRun()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) OrTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.at(SPL2ParserOR)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) AndTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.at(SPL2ParserAND)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) ComparisonTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.atCompare()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) AdditiveTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.atAdd()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) MultiplicativeTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.atMul()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL2Parser) Primary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.callAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
