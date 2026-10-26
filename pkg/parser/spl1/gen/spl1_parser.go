// Code generated from pkg/parser/spl1/grammar/SPL1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package spl1gen // SPL1
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

type SPL1Parser struct {
	*antlr.BaseParser
}

var SPL1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func spl1ParserInit() {
	staticData := &SPL1ParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "'=='", "'!='", "'<='", "'>='", "'='", "'<'",
		"'>'", "'!'", "'|'", "','", "'('", "')'", "'['", "']'", "'+'", "'-'",
		"'*'", "'/'", "'%'",
	}
	staticData.SymbolicNames = []string{
		"", "SEARCH", "WHERE", "STATS", "TIMECHART", "EVAL", "REX", "TABLE",
		"FIELDS", "SORT", "HEAD", "BIN", "BY", "AS", "SPAN", "LIMIT", "FIELD",
		"AND", "OR", "NOT", "DURATION", "NUMBER", "STRING", "IDENT", "EQ2",
		"NE", "LE", "GE", "EQ", "LT", "GT", "BANG", "PIPE", "COMMA", "LPAREN",
		"RPAREN", "LBRACKET", "RBRACKET", "PLUS", "MINUS", "STAR", "SLASH",
		"PERCENT", "WS", "UNKNOWN", "NUMBER_TIGHT", "STAR_TIGHT", "MINUS_TIGHT",
		"STAR_TERM", "NOT_WORD",
	}
	staticData.RuleNames = []string{
		"stageCommand", "stageSearch", "stageEnd", "command", "searchExpr",
		"searchOrTail", "searchAnd", "searchAndPart", "searchNot", "searchComparison",
		"searchComparisonTail", "searchOperand", "searchOperandValue", "rhsTail",
		"searchOperandExpr", "searchOperandUnary", "searchOperandPrimary", "searchAdditive",
		"searchAddTail", "searchMultiplicative", "searchMulTail", "searchUnary",
		"searchPrimary", "searchTermTail", "valueExpr", "valueOrTail", "valueAnd",
		"valueAndPart", "valueNot", "valueComparison", "valueComparisonTail",
		"valueAdd", "valueAddTail", "valueMul", "valueMulTail", "valueUnary",
		"valuePrimary", "subsearch", "absentOperand", "comparisonOp", "star",
		"minus", "number", "addOp", "mulOp", "searchMulOp", "sign", "identLike",
		"exprWord", "aggList", "agg", "aggListNotSpan", "aggNotSpan", "aliasClause",
		"byClause", "timechartBody", "timechartSpan", "spanAssign", "rexField",
		"fieldExprList", "fieldExpr", "assignment", "spanValue", "tableItem",
		"fieldsSign", "fieldsItem", "sortBody", "sortItem", "sortFirstItem",
		"sortField", "limitArg", "limitTail", "binItem", "binField", "identLikeNotSpan",
		"identLikeNotLimit",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 49, 740, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 3, 3, 163, 8, 3, 1, 3, 1, 3, 3, 3, 167, 8, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 172, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		180, 8, 3, 10, 3, 12, 3, 183, 9, 3, 1, 3, 1, 3, 3, 3, 187, 8, 3, 1, 3,
		1, 3, 1, 3, 5, 3, 192, 8, 3, 10, 3, 12, 3, 195, 9, 3, 1, 3, 1, 3, 3, 3,
		199, 8, 3, 1, 3, 5, 3, 202, 8, 3, 10, 3, 12, 3, 205, 9, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 3, 3, 211, 8, 3, 1, 3, 1, 3, 5, 3, 215, 8, 3, 10, 3, 12, 3,
		218, 9, 3, 3, 3, 220, 8, 3, 1, 4, 1, 4, 5, 4, 224, 8, 4, 10, 4, 12, 4,
		227, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 233, 8, 5, 1, 6, 1, 6, 5, 6, 237,
		8, 6, 10, 6, 12, 6, 240, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 246, 8, 7,
		1, 7, 1, 7, 3, 7, 250, 8, 7, 1, 8, 1, 8, 1, 8, 3, 8, 255, 8, 8, 1, 8, 3,
		8, 258, 8, 8, 1, 9, 1, 9, 5, 9, 262, 8, 9, 10, 9, 12, 9, 265, 9, 9, 1,
		10, 1, 10, 1, 10, 1, 10, 3, 10, 271, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11,
		276, 8, 11, 1, 12, 1, 12, 5, 12, 280, 8, 12, 10, 12, 12, 12, 283, 9, 12,
		1, 12, 1, 12, 5, 12, 287, 8, 12, 10, 12, 12, 12, 290, 9, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 299, 8, 12, 3, 12, 301, 8,
		12, 1, 12, 5, 12, 304, 8, 12, 10, 12, 12, 12, 307, 9, 12, 1, 12, 1, 12,
		1, 12, 3, 12, 312, 8, 12, 1, 12, 5, 12, 315, 8, 12, 10, 12, 12, 12, 318,
		9, 12, 3, 12, 320, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 333, 8, 13, 3, 13, 335, 8, 13, 1,
		14, 1, 14, 5, 14, 339, 8, 14, 10, 14, 12, 14, 342, 9, 14, 1, 14, 5, 14,
		345, 8, 14, 10, 14, 12, 14, 348, 9, 14, 1, 15, 1, 15, 1, 15, 3, 15, 353,
		8, 15, 1, 15, 3, 15, 356, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 3, 16, 366, 8, 16, 1, 17, 1, 17, 5, 17, 370, 8, 17, 10,
		17, 12, 17, 373, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 379, 8, 18,
		1, 19, 1, 19, 5, 19, 383, 8, 19, 10, 19, 12, 19, 386, 9, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 392, 8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 397, 8,
		21, 1, 21, 3, 21, 400, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 5, 22, 409, 8, 22, 10, 22, 12, 22, 412, 9, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 4, 22, 419, 8, 22, 11, 22, 12, 22, 420, 1, 22, 1, 22,
		3, 22, 425, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 433,
		8, 23, 1, 24, 1, 24, 5, 24, 437, 8, 24, 10, 24, 12, 24, 440, 9, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 3, 25, 446, 8, 25, 1, 26, 1, 26, 5, 26, 450, 8,
		26, 10, 26, 12, 26, 453, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 459,
		8, 27, 1, 27, 1, 27, 3, 27, 463, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28, 468,
		8, 28, 1, 28, 3, 28, 471, 8, 28, 1, 29, 1, 29, 5, 29, 475, 8, 29, 10, 29,
		12, 29, 478, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 484, 8, 30, 1, 31,
		1, 31, 5, 31, 488, 8, 31, 10, 31, 12, 31, 491, 9, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 497, 8, 32, 1, 33, 1, 33, 5, 33, 501, 8, 33, 10, 33,
		12, 33, 504, 9, 33, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 510, 8, 34, 1, 35,
		1, 35, 1, 35, 3, 35, 515, 8, 35, 1, 35, 3, 35, 518, 8, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 530, 8,
		36, 10, 36, 12, 36, 533, 9, 36, 3, 36, 535, 8, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 545, 8, 36, 1, 37, 1, 37, 1,
		37, 5, 37, 550, 8, 37, 10, 37, 12, 37, 553, 9, 37, 1, 38, 1, 38, 1, 38,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 3,
		43, 568, 8, 43, 1, 44, 1, 44, 1, 44, 3, 44, 573, 8, 44, 1, 45, 1, 45, 1,
		46, 1, 46, 3, 46, 579, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 5, 49, 588, 8, 49, 10, 49, 12, 49, 591, 9, 49, 1, 50, 1, 50, 1,
		50, 3, 50, 596, 8, 50, 1, 50, 3, 50, 599, 8, 50, 1, 50, 3, 50, 602, 8,
		50, 1, 51, 1, 51, 1, 51, 5, 51, 607, 8, 51, 10, 51, 12, 51, 610, 9, 51,
		1, 52, 1, 52, 1, 52, 3, 52, 615, 8, 52, 1, 52, 3, 52, 618, 8, 52, 1, 52,
		3, 52, 621, 8, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 55, 1, 55, 3, 55, 633, 8, 55, 1, 55, 1, 55, 3, 55, 637, 8, 55, 3,
		55, 639, 8, 55, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58,
		1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 656, 8, 59, 10,
		59, 12, 59, 659, 9, 59, 1, 60, 1, 60, 3, 60, 663, 8, 60, 1, 61, 1, 61,
		1, 61, 1, 61, 1, 62, 1, 62, 3, 62, 671, 8, 62, 1, 63, 1, 63, 1, 63, 3,
		63, 676, 8, 63, 1, 64, 1, 64, 3, 64, 680, 8, 64, 1, 65, 1, 65, 3, 65, 684,
		8, 65, 1, 66, 1, 66, 5, 66, 688, 8, 66, 10, 66, 12, 66, 691, 9, 66, 1,
		66, 1, 66, 5, 66, 695, 8, 66, 10, 66, 12, 66, 698, 9, 66, 1, 66, 3, 66,
		701, 8, 66, 1, 67, 1, 67, 3, 67, 705, 8, 67, 1, 67, 3, 67, 708, 8, 67,
		1, 68, 1, 68, 3, 68, 712, 8, 68, 1, 68, 3, 68, 715, 8, 68, 1, 69, 1, 69,
		1, 70, 1, 70, 1, 70, 1, 70, 3, 70, 723, 8, 70, 1, 71, 1, 71, 1, 71, 1,
		72, 1, 72, 1, 72, 1, 72, 3, 72, 732, 8, 72, 1, 73, 1, 73, 1, 74, 1, 74,
		1, 75, 1, 75, 1, 75, 0, 0, 76, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
		96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124,
		126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 0, 13,
		1, 1, 32, 32, 2, 0, 19, 19, 31, 31, 1, 0, 38, 39, 4, 1, 24, 30, 32, 33,
		35, 35, 37, 37, 1, 0, 24, 30, 3, 0, 40, 40, 46, 46, 48, 48, 2, 0, 39, 39,
		47, 47, 2, 0, 21, 21, 45, 45, 2, 0, 40, 42, 46, 46, 3, 0, 1, 19, 23, 23,
		49, 49, 3, 0, 1, 18, 23, 23, 49, 49, 3, 0, 1, 13, 15, 19, 23, 23, 3, 0,
		1, 14, 16, 19, 23, 23, 794, 0, 152, 1, 0, 0, 0, 2, 155, 1, 0, 0, 0, 4,
		158, 1, 0, 0, 0, 6, 219, 1, 0, 0, 0, 8, 221, 1, 0, 0, 0, 10, 228, 1, 0,
		0, 0, 12, 234, 1, 0, 0, 0, 14, 249, 1, 0, 0, 0, 16, 257, 1, 0, 0, 0, 18,
		259, 1, 0, 0, 0, 20, 266, 1, 0, 0, 0, 22, 275, 1, 0, 0, 0, 24, 319, 1,
		0, 0, 0, 26, 334, 1, 0, 0, 0, 28, 336, 1, 0, 0, 0, 30, 355, 1, 0, 0, 0,
		32, 365, 1, 0, 0, 0, 34, 367, 1, 0, 0, 0, 36, 374, 1, 0, 0, 0, 38, 380,
		1, 0, 0, 0, 40, 387, 1, 0, 0, 0, 42, 399, 1, 0, 0, 0, 44, 424, 1, 0, 0,
		0, 46, 432, 1, 0, 0, 0, 48, 434, 1, 0, 0, 0, 50, 441, 1, 0, 0, 0, 52, 447,
		1, 0, 0, 0, 54, 462, 1, 0, 0, 0, 56, 470, 1, 0, 0, 0, 58, 472, 1, 0, 0,
		0, 60, 479, 1, 0, 0, 0, 62, 485, 1, 0, 0, 0, 64, 492, 1, 0, 0, 0, 66, 498,
		1, 0, 0, 0, 68, 505, 1, 0, 0, 0, 70, 517, 1, 0, 0, 0, 72, 544, 1, 0, 0,
		0, 74, 546, 1, 0, 0, 0, 76, 554, 1, 0, 0, 0, 78, 557, 1, 0, 0, 0, 80, 559,
		1, 0, 0, 0, 82, 561, 1, 0, 0, 0, 84, 563, 1, 0, 0, 0, 86, 567, 1, 0, 0,
		0, 88, 572, 1, 0, 0, 0, 90, 574, 1, 0, 0, 0, 92, 578, 1, 0, 0, 0, 94, 580,
		1, 0, 0, 0, 96, 582, 1, 0, 0, 0, 98, 584, 1, 0, 0, 0, 100, 592, 1, 0, 0,
		0, 102, 603, 1, 0, 0, 0, 104, 611, 1, 0, 0, 0, 106, 622, 1, 0, 0, 0, 108,
		625, 1, 0, 0, 0, 110, 638, 1, 0, 0, 0, 112, 640, 1, 0, 0, 0, 114, 643,
		1, 0, 0, 0, 116, 646, 1, 0, 0, 0, 118, 651, 1, 0, 0, 0, 120, 662, 1, 0,
		0, 0, 122, 664, 1, 0, 0, 0, 124, 670, 1, 0, 0, 0, 126, 675, 1, 0, 0, 0,
		128, 679, 1, 0, 0, 0, 130, 683, 1, 0, 0, 0, 132, 700, 1, 0, 0, 0, 134,
		707, 1, 0, 0, 0, 136, 714, 1, 0, 0, 0, 138, 716, 1, 0, 0, 0, 140, 722,
		1, 0, 0, 0, 142, 724, 1, 0, 0, 0, 144, 731, 1, 0, 0, 0, 146, 733, 1, 0,
		0, 0, 148, 735, 1, 0, 0, 0, 150, 737, 1, 0, 0, 0, 152, 153, 3, 6, 3, 0,
		153, 154, 3, 4, 2, 0, 154, 1, 1, 0, 0, 0, 155, 156, 3, 8, 4, 0, 156, 157,
		3, 4, 2, 0, 157, 3, 1, 0, 0, 0, 158, 159, 7, 0, 0, 0, 159, 5, 1, 0, 0,
		0, 160, 162, 5, 1, 0, 0, 161, 163, 3, 8, 4, 0, 162, 161, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 220, 1, 0, 0, 0, 164, 166, 5, 2, 0, 0, 165, 167,
		3, 48, 24, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 220, 1,
		0, 0, 0, 168, 169, 5, 3, 0, 0, 169, 171, 3, 98, 49, 0, 170, 172, 3, 108,
		54, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 220, 1, 0, 0, 0,
		173, 174, 5, 4, 0, 0, 174, 220, 3, 110, 55, 0, 175, 176, 5, 5, 0, 0, 176,
		181, 3, 122, 61, 0, 177, 178, 5, 33, 0, 0, 178, 180, 3, 122, 61, 0, 179,
		177, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 181, 182,
		1, 0, 0, 0, 182, 220, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184, 186, 5, 6,
		0, 0, 185, 187, 3, 116, 58, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0,
		0, 187, 188, 1, 0, 0, 0, 188, 220, 5, 22, 0, 0, 189, 193, 5, 7, 0, 0, 190,
		192, 3, 126, 63, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191,
		1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 220, 1, 0, 0, 0, 195, 193, 1, 0,
		0, 0, 196, 198, 5, 8, 0, 0, 197, 199, 3, 128, 64, 0, 198, 197, 1, 0, 0,
		0, 198, 199, 1, 0, 0, 0, 199, 203, 1, 0, 0, 0, 200, 202, 3, 130, 65, 0,
		201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 220, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207,
		5, 9, 0, 0, 207, 220, 3, 132, 66, 0, 208, 210, 5, 10, 0, 0, 209, 211, 3,
		140, 70, 0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 220, 1, 0,
		0, 0, 212, 216, 5, 11, 0, 0, 213, 215, 3, 144, 72, 0, 214, 213, 1, 0, 0,
		0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217,
		220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 160, 1, 0, 0, 0, 219, 164,
		1, 0, 0, 0, 219, 168, 1, 0, 0, 0, 219, 173, 1, 0, 0, 0, 219, 175, 1, 0,
		0, 0, 219, 184, 1, 0, 0, 0, 219, 189, 1, 0, 0, 0, 219, 196, 1, 0, 0, 0,
		219, 206, 1, 0, 0, 0, 219, 208, 1, 0, 0, 0, 219, 212, 1, 0, 0, 0, 220,
		7, 1, 0, 0, 0, 221, 225, 3, 12, 6, 0, 222, 224, 3, 10, 5, 0, 223, 222,
		1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0,
		0, 0, 226, 9, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 229, 4, 5, 0, 0, 229,
		232, 5, 18, 0, 0, 230, 233, 3, 12, 6, 0, 231, 233, 3, 76, 38, 0, 232, 230,
		1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 11, 1, 0, 0, 0, 234, 238, 3, 16,
		8, 0, 235, 237, 3, 14, 7, 0, 236, 235, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0,
		238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 13, 1, 0, 0, 0, 240, 238,
		1, 0, 0, 0, 241, 242, 4, 7, 1, 0, 242, 245, 5, 17, 0, 0, 243, 246, 3, 16,
		8, 0, 244, 246, 3, 76, 38, 0, 245, 243, 1, 0, 0, 0, 245, 244, 1, 0, 0,
		0, 246, 250, 1, 0, 0, 0, 247, 248, 4, 7, 2, 0, 248, 250, 3, 16, 8, 0, 249,
		241, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250, 15, 1, 0, 0, 0, 251, 254, 7,
		1, 0, 0, 252, 255, 3, 16, 8, 0, 253, 255, 3, 76, 38, 0, 254, 252, 1, 0,
		0, 0, 254, 253, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 258, 3, 18, 9, 0,
		257, 251, 1, 0, 0, 0, 257, 256, 1, 0, 0, 0, 258, 17, 1, 0, 0, 0, 259, 263,
		3, 34, 17, 0, 260, 262, 3, 20, 10, 0, 261, 260, 1, 0, 0, 0, 262, 265, 1,
		0, 0, 0, 263, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 19, 1, 0, 0,
		0, 265, 263, 1, 0, 0, 0, 266, 267, 4, 10, 3, 0, 267, 270, 3, 78, 39, 0,
		268, 271, 3, 22, 11, 0, 269, 271, 3, 76, 38, 0, 270, 268, 1, 0, 0, 0, 270,
		269, 1, 0, 0, 0, 271, 21, 1, 0, 0, 0, 272, 276, 5, 22, 0, 0, 273, 276,
		3, 24, 12, 0, 274, 276, 3, 28, 14, 0, 275, 272, 1, 0, 0, 0, 275, 273, 1,
		0, 0, 0, 275, 274, 1, 0, 0, 0, 276, 23, 1, 0, 0, 0, 277, 281, 3, 94, 47,
		0, 278, 280, 3, 26, 13, 0, 279, 278, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0,
		281, 279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 320, 1, 0, 0, 0, 283,
		281, 1, 0, 0, 0, 284, 288, 5, 45, 0, 0, 285, 287, 3, 26, 13, 0, 286, 285,
		1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0,
		0, 0, 289, 320, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 300, 5, 46, 0, 0,
		292, 301, 3, 94, 47, 0, 293, 301, 3, 84, 42, 0, 294, 298, 3, 82, 41, 0,
		295, 299, 3, 94, 47, 0, 296, 299, 3, 84, 42, 0, 297, 299, 3, 80, 40, 0,
		298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 297, 1, 0, 0, 0, 299,
		301, 1, 0, 0, 0, 300, 292, 1, 0, 0, 0, 300, 293, 1, 0, 0, 0, 300, 294,
		1, 0, 0, 0, 301, 305, 1, 0, 0, 0, 302, 304, 3, 26, 13, 0, 303, 302, 1,
		0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306, 1, 0, 0,
		0, 306, 320, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 308, 311, 5, 47, 0, 0, 309,
		312, 3, 94, 47, 0, 310, 312, 3, 80, 40, 0, 311, 309, 1, 0, 0, 0, 311, 310,
		1, 0, 0, 0, 312, 316, 1, 0, 0, 0, 313, 315, 3, 26, 13, 0, 314, 313, 1,
		0, 0, 0, 315, 318, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0,
		0, 317, 320, 1, 0, 0, 0, 318, 316, 1, 0, 0, 0, 319, 277, 1, 0, 0, 0, 319,
		284, 1, 0, 0, 0, 319, 291, 1, 0, 0, 0, 319, 308, 1, 0, 0, 0, 320, 25, 1,
		0, 0, 0, 321, 322, 4, 13, 4, 0, 322, 335, 3, 80, 40, 0, 323, 324, 4, 13,
		5, 0, 324, 335, 3, 94, 47, 0, 325, 326, 4, 13, 6, 0, 326, 335, 3, 84, 42,
		0, 327, 328, 4, 13, 7, 0, 328, 332, 3, 82, 41, 0, 329, 333, 3, 94, 47,
		0, 330, 333, 3, 84, 42, 0, 331, 333, 3, 80, 40, 0, 332, 329, 1, 0, 0, 0,
		332, 330, 1, 0, 0, 0, 332, 331, 1, 0, 0, 0, 333, 335, 1, 0, 0, 0, 334,
		321, 1, 0, 0, 0, 334, 323, 1, 0, 0, 0, 334, 325, 1, 0, 0, 0, 334, 327,
		1, 0, 0, 0, 335, 27, 1, 0, 0, 0, 336, 340, 3, 30, 15, 0, 337, 339, 3, 40,
		20, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0,
		340, 341, 1, 0, 0, 0, 341, 346, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 343,
		345, 3, 36, 18, 0, 344, 343, 1, 0, 0, 0, 345, 348, 1, 0, 0, 0, 346, 344,
		1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 29, 1, 0, 0, 0, 348, 346, 1, 0,
		0, 0, 349, 352, 7, 2, 0, 0, 350, 353, 3, 42, 21, 0, 351, 353, 3, 76, 38,
		0, 352, 350, 1, 0, 0, 0, 352, 351, 1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354,
		356, 3, 32, 16, 0, 355, 349, 1, 0, 0, 0, 355, 354, 1, 0, 0, 0, 356, 31,
		1, 0, 0, 0, 357, 358, 5, 34, 0, 0, 358, 359, 3, 8, 4, 0, 359, 360, 5, 35,
		0, 0, 360, 366, 1, 0, 0, 0, 361, 366, 5, 20, 0, 0, 362, 366, 5, 21, 0,
		0, 363, 366, 5, 40, 0, 0, 364, 366, 3, 74, 37, 0, 365, 357, 1, 0, 0, 0,
		365, 361, 1, 0, 0, 0, 365, 362, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365,
		364, 1, 0, 0, 0, 366, 33, 1, 0, 0, 0, 367, 371, 3, 38, 19, 0, 368, 370,
		3, 36, 18, 0, 369, 368, 1, 0, 0, 0, 370, 373, 1, 0, 0, 0, 371, 369, 1,
		0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 35, 1, 0, 0, 0, 373, 371, 1, 0, 0,
		0, 374, 375, 4, 18, 8, 0, 375, 378, 3, 86, 43, 0, 376, 379, 3, 38, 19,
		0, 377, 379, 3, 76, 38, 0, 378, 376, 1, 0, 0, 0, 378, 377, 1, 0, 0, 0,
		379, 37, 1, 0, 0, 0, 380, 384, 3, 42, 21, 0, 381, 383, 3, 40, 20, 0, 382,
		381, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385,
		1, 0, 0, 0, 385, 39, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 388, 4, 20,
		9, 0, 388, 391, 3, 90, 45, 0, 389, 392, 3, 42, 21, 0, 390, 392, 3, 76,
		38, 0, 391, 389, 1, 0, 0, 0, 391, 390, 1, 0, 0, 0, 392, 41, 1, 0, 0, 0,
		393, 396, 3, 92, 46, 0, 394, 397, 3, 42, 21, 0, 395, 397, 3, 76, 38, 0,
		396, 394, 1, 0, 0, 0, 396, 395, 1, 0, 0, 0, 397, 400, 1, 0, 0, 0, 398,
		400, 3, 44, 22, 0, 399, 393, 1, 0, 0, 0, 399, 398, 1, 0, 0, 0, 400, 43,
		1, 0, 0, 0, 401, 402, 5, 34, 0, 0, 402, 403, 3, 8, 4, 0, 403, 404, 5, 35,
		0, 0, 404, 425, 1, 0, 0, 0, 405, 425, 5, 22, 0, 0, 406, 410, 3, 96, 48,
		0, 407, 409, 3, 46, 23, 0, 408, 407, 1, 0, 0, 0, 409, 412, 1, 0, 0, 0,
		410, 408, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 425, 1, 0, 0, 0, 412,
		410, 1, 0, 0, 0, 413, 425, 3, 84, 42, 0, 414, 425, 5, 20, 0, 0, 415, 416,
		4, 22, 10, 0, 416, 418, 3, 80, 40, 0, 417, 419, 3, 46, 23, 0, 418, 417,
		1, 0, 0, 0, 419, 420, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 420, 421, 1, 0,
		0, 0, 421, 425, 1, 0, 0, 0, 422, 425, 3, 80, 40, 0, 423, 425, 3, 74, 37,
		0, 424, 401, 1, 0, 0, 0, 424, 405, 1, 0, 0, 0, 424, 406, 1, 0, 0, 0, 424,
		413, 1, 0, 0, 0, 424, 414, 1, 0, 0, 0, 424, 415, 1, 0, 0, 0, 424, 422,
		1, 0, 0, 0, 424, 423, 1, 0, 0, 0, 425, 45, 1, 0, 0, 0, 426, 427, 4, 23,
		11, 0, 427, 433, 5, 48, 0, 0, 428, 429, 4, 23, 12, 0, 429, 433, 3, 94,
		47, 0, 430, 431, 4, 23, 13, 0, 431, 433, 3, 84, 42, 0, 432, 426, 1, 0,
		0, 0, 432, 428, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 47, 1, 0, 0, 0,
		434, 438, 3, 52, 26, 0, 435, 437, 3, 50, 25, 0, 436, 435, 1, 0, 0, 0, 437,
		440, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 49, 1,
		0, 0, 0, 440, 438, 1, 0, 0, 0, 441, 442, 4, 25, 14, 0, 442, 445, 5, 18,
		0, 0, 443, 446, 3, 52, 26, 0, 444, 446, 3, 76, 38, 0, 445, 443, 1, 0, 0,
		0, 445, 444, 1, 0, 0, 0, 446, 51, 1, 0, 0, 0, 447, 451, 3, 56, 28, 0, 448,
		450, 3, 54, 27, 0, 449, 448, 1, 0, 0, 0, 450, 453, 1, 0, 0, 0, 451, 449,
		1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452, 53, 1, 0, 0, 0, 453, 451, 1, 0,
		0, 0, 454, 455, 4, 27, 15, 0, 455, 458, 5, 17, 0, 0, 456, 459, 3, 56, 28,
		0, 457, 459, 3, 76, 38, 0, 458, 456, 1, 0, 0, 0, 458, 457, 1, 0, 0, 0,
		459, 463, 1, 0, 0, 0, 460, 461, 4, 27, 16, 0, 461, 463, 3, 56, 28, 0, 462,
		454, 1, 0, 0, 0, 462, 460, 1, 0, 0, 0, 463, 55, 1, 0, 0, 0, 464, 467, 7,
		1, 0, 0, 465, 468, 3, 56, 28, 0, 466, 468, 3, 76, 38, 0, 467, 465, 1, 0,
		0, 0, 467, 466, 1, 0, 0, 0, 468, 471, 1, 0, 0, 0, 469, 471, 3, 58, 29,
		0, 470, 464, 1, 0, 0, 0, 470, 469, 1, 0, 0, 0, 471, 57, 1, 0, 0, 0, 472,
		476, 3, 62, 31, 0, 473, 475, 3, 60, 30, 0, 474, 473, 1, 0, 0, 0, 475, 478,
		1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476, 477, 1, 0, 0, 0, 477, 59, 1, 0,
		0, 0, 478, 476, 1, 0, 0, 0, 479, 480, 4, 30, 17, 0, 480, 483, 3, 78, 39,
		0, 481, 484, 3, 62, 31, 0, 482, 484, 3, 76, 38, 0, 483, 481, 1, 0, 0, 0,
		483, 482, 1, 0, 0, 0, 484, 61, 1, 0, 0, 0, 485, 489, 3, 66, 33, 0, 486,
		488, 3, 64, 32, 0, 487, 486, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487,
		1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 63, 1, 0, 0, 0, 491, 489, 1, 0,
		0, 0, 492, 493, 4, 32, 18, 0, 493, 496, 3, 86, 43, 0, 494, 497, 3, 66,
		33, 0, 495, 497, 3, 76, 38, 0, 496, 494, 1, 0, 0, 0, 496, 495, 1, 0, 0,
		0, 497, 65, 1, 0, 0, 0, 498, 502, 3, 70, 35, 0, 499, 501, 3, 68, 34, 0,
		500, 499, 1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 502,
		503, 1, 0, 0, 0, 503, 67, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 505, 506, 4,
		34, 19, 0, 506, 509, 3, 88, 44, 0, 507, 510, 3, 70, 35, 0, 508, 510, 3,
		76, 38, 0, 509, 507, 1, 0, 0, 0, 509, 508, 1, 0, 0, 0, 510, 69, 1, 0, 0,
		0, 511, 514, 3, 92, 46, 0, 512, 515, 3, 70, 35, 0, 513, 515, 3, 76, 38,
		0, 514, 512, 1, 0, 0, 0, 514, 513, 1, 0, 0, 0, 515, 518, 1, 0, 0, 0, 516,
		518, 3, 72, 36, 0, 517, 511, 1, 0, 0, 0, 517, 516, 1, 0, 0, 0, 518, 71,
		1, 0, 0, 0, 519, 520, 5, 34, 0, 0, 520, 521, 3, 48, 24, 0, 521, 522, 5,
		35, 0, 0, 522, 545, 1, 0, 0, 0, 523, 524, 4, 36, 20, 0, 524, 525, 3, 96,
		48, 0, 525, 534, 5, 34, 0, 0, 526, 531, 3, 48, 24, 0, 527, 528, 5, 33,
		0, 0, 528, 530, 3, 48, 24, 0, 529, 527, 1, 0, 0, 0, 530, 533, 1, 0, 0,
		0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 535, 1, 0, 0, 0, 533,
		531, 1, 0, 0, 0, 534, 526, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 536,
		1, 0, 0, 0, 536, 537, 5, 35, 0, 0, 537, 545, 1, 0, 0, 0, 538, 545, 3, 96,
		48, 0, 539, 545, 5, 22, 0, 0, 540, 545, 3, 84, 42, 0, 541, 545, 5, 20,
		0, 0, 542, 545, 3, 80, 40, 0, 543, 545, 3, 74, 37, 0, 544, 519, 1, 0, 0,
		0, 544, 523, 1, 0, 0, 0, 544, 538, 1, 0, 0, 0, 544, 539, 1, 0, 0, 0, 544,
		540, 1, 0, 0, 0, 544, 541, 1, 0, 0, 0, 544, 542, 1, 0, 0, 0, 544, 543,
		1, 0, 0, 0, 545, 73, 1, 0, 0, 0, 546, 547, 5, 36, 0, 0, 547, 551, 6, 37,
		-1, 0, 548, 550, 8, 0, 0, 0, 549, 548, 1, 0, 0, 0, 550, 553, 1, 0, 0, 0,
		551, 549, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 75, 1, 0, 0, 0, 553, 551,
		1, 0, 0, 0, 554, 555, 6, 38, -1, 0, 555, 556, 7, 3, 0, 0, 556, 77, 1, 0,
		0, 0, 557, 558, 7, 4, 0, 0, 558, 79, 1, 0, 0, 0, 559, 560, 7, 5, 0, 0,
		560, 81, 1, 0, 0, 0, 561, 562, 7, 6, 0, 0, 562, 83, 1, 0, 0, 0, 563, 564,
		7, 7, 0, 0, 564, 85, 1, 0, 0, 0, 565, 568, 5, 38, 0, 0, 566, 568, 3, 82,
		41, 0, 567, 565, 1, 0, 0, 0, 567, 566, 1, 0, 0, 0, 568, 87, 1, 0, 0, 0,
		569, 573, 3, 80, 40, 0, 570, 573, 5, 41, 0, 0, 571, 573, 5, 42, 0, 0, 572,
		569, 1, 0, 0, 0, 572, 570, 1, 0, 0, 0, 572, 571, 1, 0, 0, 0, 573, 89, 1,
		0, 0, 0, 574, 575, 7, 8, 0, 0, 575, 91, 1, 0, 0, 0, 576, 579, 3, 82, 41,
		0, 577, 579, 5, 38, 0, 0, 578, 576, 1, 0, 0, 0, 578, 577, 1, 0, 0, 0, 579,
		93, 1, 0, 0, 0, 580, 581, 7, 9, 0, 0, 581, 95, 1, 0, 0, 0, 582, 583, 7,
		10, 0, 0, 583, 97, 1, 0, 0, 0, 584, 589, 3, 100, 50, 0, 585, 586, 5, 33,
		0, 0, 586, 588, 3, 100, 50, 0, 587, 585, 1, 0, 0, 0, 588, 591, 1, 0, 0,
		0, 589, 587, 1, 0, 0, 0, 589, 590, 1, 0, 0, 0, 590, 99, 1, 0, 0, 0, 591,
		589, 1, 0, 0, 0, 592, 598, 3, 94, 47, 0, 593, 595, 5, 34, 0, 0, 594, 596,
		3, 48, 24, 0, 595, 594, 1, 0, 0, 0, 595, 596, 1, 0, 0, 0, 596, 597, 1,
		0, 0, 0, 597, 599, 5, 35, 0, 0, 598, 593, 1, 0, 0, 0, 598, 599, 1, 0, 0,
		0, 599, 601, 1, 0, 0, 0, 600, 602, 3, 106, 53, 0, 601, 600, 1, 0, 0, 0,
		601, 602, 1, 0, 0, 0, 602, 101, 1, 0, 0, 0, 603, 608, 3, 104, 52, 0, 604,
		605, 5, 33, 0, 0, 605, 607, 3, 100, 50, 0, 606, 604, 1, 0, 0, 0, 607, 610,
		1, 0, 0, 0, 608, 606, 1, 0, 0, 0, 608, 609, 1, 0, 0, 0, 609, 103, 1, 0,
		0, 0, 610, 608, 1, 0, 0, 0, 611, 617, 3, 148, 74, 0, 612, 614, 5, 34, 0,
		0, 613, 615, 3, 48, 24, 0, 614, 613, 1, 0, 0, 0, 614, 615, 1, 0, 0, 0,
		615, 616, 1, 0, 0, 0, 616, 618, 5, 35, 0, 0, 617, 612, 1, 0, 0, 0, 617,
		618, 1, 0, 0, 0, 618, 620, 1, 0, 0, 0, 619, 621, 3, 106, 53, 0, 620, 619,
		1, 0, 0, 0, 620, 621, 1, 0, 0, 0, 621, 105, 1, 0, 0, 0, 622, 623, 5, 13,
		0, 0, 623, 624, 3, 94, 47, 0, 624, 107, 1, 0, 0, 0, 625, 626, 5, 12, 0,
		0, 626, 627, 3, 118, 59, 0, 627, 109, 1, 0, 0, 0, 628, 629, 4, 55, 21,
		0, 629, 630, 3, 112, 56, 0, 630, 632, 3, 98, 49, 0, 631, 633, 3, 108, 54,
		0, 632, 631, 1, 0, 0, 0, 632, 633, 1, 0, 0, 0, 633, 639, 1, 0, 0, 0, 634,
		636, 3, 102, 51, 0, 635, 637, 3, 108, 54, 0, 636, 635, 1, 0, 0, 0, 636,
		637, 1, 0, 0, 0, 637, 639, 1, 0, 0, 0, 638, 628, 1, 0, 0, 0, 638, 634,
		1, 0, 0, 0, 639, 111, 1, 0, 0, 0, 640, 641, 5, 14, 0, 0, 641, 642, 3, 114,
		57, 0, 642, 113, 1, 0, 0, 0, 643, 644, 5, 28, 0, 0, 644, 645, 3, 124, 62,
		0, 645, 115, 1, 0, 0, 0, 646, 647, 4, 58, 22, 0, 647, 648, 5, 16, 0, 0,
		648, 649, 5, 28, 0, 0, 649, 650, 3, 94, 47, 0, 650, 117, 1, 0, 0, 0, 651,
		657, 3, 120, 60, 0, 652, 653, 5, 33, 0, 0, 653, 656, 3, 120, 60, 0, 654,
		656, 3, 120, 60, 0, 655, 652, 1, 0, 0, 0, 655, 654, 1, 0, 0, 0, 656, 659,
		1, 0, 0, 0, 657, 655, 1, 0, 0, 0, 657, 658, 1, 0, 0, 0, 658, 119, 1, 0,
		0, 0, 659, 657, 1, 0, 0, 0, 660, 663, 3, 94, 47, 0, 661, 663, 3, 80, 40,
		0, 662, 660, 1, 0, 0, 0, 662, 661, 1, 0, 0, 0, 663, 121, 1, 0, 0, 0, 664,
		665, 3, 94, 47, 0, 665, 666, 5, 28, 0, 0, 666, 667, 3, 48, 24, 0, 667,
		123, 1, 0, 0, 0, 668, 671, 5, 20, 0, 0, 669, 671, 3, 94, 47, 0, 670, 668,
		1, 0, 0, 0, 670, 669, 1, 0, 0, 0, 671, 125, 1, 0, 0, 0, 672, 676, 5, 33,
		0, 0, 673, 676, 3, 80, 40, 0, 674, 676, 3, 94, 47, 0, 675, 672, 1, 0, 0,
		0, 675, 673, 1, 0, 0, 0, 675, 674, 1, 0, 0, 0, 676, 127, 1, 0, 0, 0, 677,
		680, 3, 82, 41, 0, 678, 680, 5, 38, 0, 0, 679, 677, 1, 0, 0, 0, 679, 678,
		1, 0, 0, 0, 680, 129, 1, 0, 0, 0, 681, 684, 5, 33, 0, 0, 682, 684, 3, 94,
		47, 0, 683, 681, 1, 0, 0, 0, 683, 682, 1, 0, 0, 0, 684, 131, 1, 0, 0, 0,
		685, 689, 3, 140, 70, 0, 686, 688, 3, 134, 67, 0, 687, 686, 1, 0, 0, 0,
		688, 691, 1, 0, 0, 0, 689, 687, 1, 0, 0, 0, 689, 690, 1, 0, 0, 0, 690,
		701, 1, 0, 0, 0, 691, 689, 1, 0, 0, 0, 692, 696, 3, 136, 68, 0, 693, 695,
		3, 134, 67, 0, 694, 693, 1, 0, 0, 0, 695, 698, 1, 0, 0, 0, 696, 694, 1,
		0, 0, 0, 696, 697, 1, 0, 0, 0, 697, 701, 1, 0, 0, 0, 698, 696, 1, 0, 0,
		0, 699, 701, 1, 0, 0, 0, 700, 685, 1, 0, 0, 0, 700, 692, 1, 0, 0, 0, 700,
		699, 1, 0, 0, 0, 701, 133, 1, 0, 0, 0, 702, 708, 5, 33, 0, 0, 703, 705,
		3, 92, 46, 0, 704, 703, 1, 0, 0, 0, 704, 705, 1, 0, 0, 0, 705, 706, 1,
		0, 0, 0, 706, 708, 3, 94, 47, 0, 707, 702, 1, 0, 0, 0, 707, 704, 1, 0,
		0, 0, 708, 135, 1, 0, 0, 0, 709, 715, 5, 33, 0, 0, 710, 712, 3, 92, 46,
		0, 711, 710, 1, 0, 0, 0, 711, 712, 1, 0, 0, 0, 712, 713, 1, 0, 0, 0, 713,
		715, 3, 138, 69, 0, 714, 709, 1, 0, 0, 0, 714, 711, 1, 0, 0, 0, 715, 137,
		1, 0, 0, 0, 716, 717, 3, 150, 75, 0, 717, 139, 1, 0, 0, 0, 718, 719, 4,
		70, 23, 0, 719, 720, 5, 15, 0, 0, 720, 723, 3, 142, 71, 0, 721, 723, 3,
		84, 42, 0, 722, 718, 1, 0, 0, 0, 722, 721, 1, 0, 0, 0, 723, 141, 1, 0,
		0, 0, 724, 725, 5, 28, 0, 0, 725, 726, 3, 84, 42, 0, 726, 143, 1, 0, 0,
		0, 727, 728, 4, 72, 24, 0, 728, 729, 5, 14, 0, 0, 729, 732, 3, 114, 57,
		0, 730, 732, 3, 146, 73, 0, 731, 727, 1, 0, 0, 0, 731, 730, 1, 0, 0, 0,
		732, 145, 1, 0, 0, 0, 733, 734, 3, 148, 74, 0, 734, 147, 1, 0, 0, 0, 735,
		736, 7, 11, 0, 0, 736, 149, 1, 0, 0, 0, 737, 738, 7, 12, 0, 0, 738, 151,
		1, 0, 0, 0, 95, 162, 166, 171, 181, 186, 193, 198, 203, 210, 216, 219,
		225, 232, 238, 245, 249, 254, 257, 263, 270, 275, 281, 288, 298, 300, 305,
		311, 316, 319, 332, 334, 340, 346, 352, 355, 365, 371, 378, 384, 391, 396,
		399, 410, 420, 424, 432, 438, 445, 451, 458, 462, 467, 470, 476, 483, 489,
		496, 502, 509, 514, 517, 531, 534, 544, 551, 567, 572, 578, 589, 595, 598,
		601, 608, 614, 617, 620, 632, 636, 638, 655, 657, 662, 670, 675, 679, 683,
		689, 696, 700, 704, 707, 711, 714, 722, 731,
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

// SPL1ParserInit initializes any static state used to implement SPL1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSPL1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SPL1ParserInit() {
	staticData := &SPL1ParserStaticData
	staticData.once.Do(spl1ParserInit)
}

// NewSPL1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewSPL1Parser(input antlr.TokenStream) *SPL1Parser {
	SPL1ParserInit()
	this := new(SPL1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SPL1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SPL1.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

// la returns the token type at lookahead k.
func (p *SPL1Parser) la(k int) int { return p.GetTokenStream().LA(k) }

// tok returns the token at lookahead k; k may be negative.
func (p *SPL1Parser) tok(k int) antlr.Token { return p.GetTokenStream().LT(k) }

// at reports whether the current token has type t.
func (p *SPL1Parser) at(t int) bool { return p.la(1) == t }

// isIdentLike reports whether token type t is accepted by identLike.
func (p *SPL1Parser) isIdentLike(t int) bool {
	switch t {
	case SPL1ParserIDENT, SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS,
		SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE,
		SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN,
		SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT,
		SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT,
		SPL1ParserNOT_WORD:
		return true
	}
	return false
}

// touches reports whether the tokens at k and k+1 are adjacent in the source.
// Token offsets are rune indexes, so this is exact.
func (p *SPL1Parser) touches(k int) bool {
	a, b := p.tok(k), p.tok(k+1)
	if a == nil || b == nil {
		return false
	}
	return a.GetStop()+1 == b.GetStart()
}

// touchingPrev reports whether the current token starts exactly where the
// last consumed token ended.
func (p *SPL1Parser) touchingPrev() bool {
	prev, cur := p.tok(-1), p.tok(1)
	if prev == nil || cur == nil {
		return false
	}
	return prev.GetStop()+1 == cur.GetStart()
}

// startsTerm is the implicit-AND boundary. Every word starts a term except
// the conjunction and clause keywords, matching the hand parser. Tight tokens
// are the same spellings as their plain forms, so they start a term too.
func (p *SPL1Parser) startsTerm() bool {
	t := p.la(1)
	switch t {
	case SPL1ParserAND, SPL1ParserOR, SPL1ParserBY, SPL1ParserAS:
		return false
	case SPL1ParserSTRING, SPL1ParserNUMBER, SPL1ParserNUMBER_TIGHT,
		SPL1ParserDURATION, SPL1ParserLPAREN, SPL1ParserLBRACKET,
		SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
		return true
	}
	return p.isIdentLike(t)
}

// isStar, isNumber, and isMinus accept every spelling of a token, wherever the
// term, value-run, and operator distinctions do not matter.
func (p *SPL1Parser) isStar(t int) bool {
	return t == SPL1ParserSTAR || t == SPL1ParserSTAR_TIGHT ||
		t == SPL1ParserSTAR_TERM
}

func (p *SPL1Parser) isNumber(t int) bool {
	return t == SPL1ParserNUMBER || t == SPL1ParserNUMBER_TIGHT
}

func (p *SPL1Parser) isMinus(t int) bool {
	return t == SPL1ParserMINUS || t == SPL1ParserMINUS_TIGHT
}

// callAhead reports whether an identifier is followed by '('. Adjacency is
// not required: `eval x=lower (msg)` is a call.
func (p *SPL1Parser) callAhead() bool { return p.la(2) == SPL1ParserLPAREN }

// atNot, atCompare, atAdd, and atMul report whether the current token is an
// operator of that kind. The binary and unary rules gate on them so the
// parser commits to the operator and blames the missing operand after it,
// which is where the hand parser reports.
func (p *SPL1Parser) atNot() bool { return p.at(SPL1ParserNOT) || p.at(SPL1ParserBANG) }

func (p *SPL1Parser) atCompare() bool {
	switch p.la(1) {
	case SPL1ParserEQ2, SPL1ParserEQ, SPL1ParserNE, SPL1ParserLT, SPL1ParserLE,
		SPL1ParserGT, SPL1ParserGE:
		return true
	}
	return false
}

func (p *SPL1Parser) atAdd() bool {
	return p.at(SPL1ParserPLUS) || p.at(SPL1ParserMINUS) || p.at(SPL1ParserMINUS_TIGHT)
}

func (p *SPL1Parser) atMul() bool {
	t := p.la(1)
	return p.isStar(t) || t == SPL1ParserSLASH || t == SPL1ParserPERCENT
}

// atSearchMul is atMul for term position, where an adjacent STAR is part of
// the preceding term (`error*`) rather than a multiplication operator.
func (p *SPL1Parser) atSearchMul() bool {
	t := p.la(1)
	return t == SPL1ParserSTAR || t == SPL1ParserSTAR_TIGHT ||
		t == SPL1ParserSLASH || t == SPL1ParserPERCENT
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
// required, which is where the hand parser fails.
func (p *SPL1Parser) abortMissingOperand() {
	t := p.tok(1)
	message := "expected an expression, found end of query"
	if t != nil && t.GetTokenType() != antlr.TokenEOF {
		message = "expected an expression, found " + fmt.Sprintf("%q", t.GetText())
	}
	panic(parseAbort{token: t, code: "syntax", message: message})
}

// abortSubsearch refuses a subsearch at its opening bracket, before the rest of
// the stage is parsed, matching the hand parser's fail-fast priority.
func (p *SPL1Parser) abortSubsearch() {
	panic(parseAbort{
		token:   p.tok(-1),
		code:    "unsupported-subsearch",
		message: "subsearches are not supported",
	})
}

// termStarRun reports whether a '*' begins a tight free-text wildcard run
// such as `*error*`, as opposed to a bare wildcard.
func (p *SPL1Parser) termStarRun() bool {
	if !p.isStar(p.la(1)) || !p.touches(1) {
		return false
	}
	return p.isIdentLike(p.la(2)) || p.isNumber(p.la(2))
}

// hyphenTail reports whether a '-' continues a value run: it must touch the
// previous component and be followed, without a space, by a word, number, or
// wildcard.
func (p *SPL1Parser) hyphenTail() bool {
	return p.touchingPrev() && p.touches(1) &&
		(p.isIdentLike(p.la(2)) || p.isNumber(p.la(2)) || p.isStar(p.la(2)))
}

// SPL1Parser tokens.
const (
	SPL1ParserEOF          = antlr.TokenEOF
	SPL1ParserSEARCH       = 1
	SPL1ParserWHERE        = 2
	SPL1ParserSTATS        = 3
	SPL1ParserTIMECHART    = 4
	SPL1ParserEVAL         = 5
	SPL1ParserREX          = 6
	SPL1ParserTABLE        = 7
	SPL1ParserFIELDS       = 8
	SPL1ParserSORT         = 9
	SPL1ParserHEAD         = 10
	SPL1ParserBIN          = 11
	SPL1ParserBY           = 12
	SPL1ParserAS           = 13
	SPL1ParserSPAN         = 14
	SPL1ParserLIMIT        = 15
	SPL1ParserFIELD        = 16
	SPL1ParserAND          = 17
	SPL1ParserOR           = 18
	SPL1ParserNOT          = 19
	SPL1ParserDURATION     = 20
	SPL1ParserNUMBER       = 21
	SPL1ParserSTRING       = 22
	SPL1ParserIDENT        = 23
	SPL1ParserEQ2          = 24
	SPL1ParserNE           = 25
	SPL1ParserLE           = 26
	SPL1ParserGE           = 27
	SPL1ParserEQ           = 28
	SPL1ParserLT           = 29
	SPL1ParserGT           = 30
	SPL1ParserBANG         = 31
	SPL1ParserPIPE         = 32
	SPL1ParserCOMMA        = 33
	SPL1ParserLPAREN       = 34
	SPL1ParserRPAREN       = 35
	SPL1ParserLBRACKET     = 36
	SPL1ParserRBRACKET     = 37
	SPL1ParserPLUS         = 38
	SPL1ParserMINUS        = 39
	SPL1ParserSTAR         = 40
	SPL1ParserSLASH        = 41
	SPL1ParserPERCENT      = 42
	SPL1ParserWS           = 43
	SPL1ParserUNKNOWN      = 44
	SPL1ParserNUMBER_TIGHT = 45
	SPL1ParserSTAR_TIGHT   = 46
	SPL1ParserMINUS_TIGHT  = 47
	SPL1ParserSTAR_TERM    = 48
	SPL1ParserNOT_WORD     = 49
)

// SPL1Parser rules.
const (
	SPL1ParserRULE_stageCommand         = 0
	SPL1ParserRULE_stageSearch          = 1
	SPL1ParserRULE_stageEnd             = 2
	SPL1ParserRULE_command              = 3
	SPL1ParserRULE_searchExpr           = 4
	SPL1ParserRULE_searchOrTail         = 5
	SPL1ParserRULE_searchAnd            = 6
	SPL1ParserRULE_searchAndPart        = 7
	SPL1ParserRULE_searchNot            = 8
	SPL1ParserRULE_searchComparison     = 9
	SPL1ParserRULE_searchComparisonTail = 10
	SPL1ParserRULE_searchOperand        = 11
	SPL1ParserRULE_searchOperandValue   = 12
	SPL1ParserRULE_rhsTail              = 13
	SPL1ParserRULE_searchOperandExpr    = 14
	SPL1ParserRULE_searchOperandUnary   = 15
	SPL1ParserRULE_searchOperandPrimary = 16
	SPL1ParserRULE_searchAdditive       = 17
	SPL1ParserRULE_searchAddTail        = 18
	SPL1ParserRULE_searchMultiplicative = 19
	SPL1ParserRULE_searchMulTail        = 20
	SPL1ParserRULE_searchUnary          = 21
	SPL1ParserRULE_searchPrimary        = 22
	SPL1ParserRULE_searchTermTail       = 23
	SPL1ParserRULE_valueExpr            = 24
	SPL1ParserRULE_valueOrTail          = 25
	SPL1ParserRULE_valueAnd             = 26
	SPL1ParserRULE_valueAndPart         = 27
	SPL1ParserRULE_valueNot             = 28
	SPL1ParserRULE_valueComparison      = 29
	SPL1ParserRULE_valueComparisonTail  = 30
	SPL1ParserRULE_valueAdd             = 31
	SPL1ParserRULE_valueAddTail         = 32
	SPL1ParserRULE_valueMul             = 33
	SPL1ParserRULE_valueMulTail         = 34
	SPL1ParserRULE_valueUnary           = 35
	SPL1ParserRULE_valuePrimary         = 36
	SPL1ParserRULE_subsearch            = 37
	SPL1ParserRULE_absentOperand        = 38
	SPL1ParserRULE_comparisonOp         = 39
	SPL1ParserRULE_star                 = 40
	SPL1ParserRULE_minus                = 41
	SPL1ParserRULE_number               = 42
	SPL1ParserRULE_addOp                = 43
	SPL1ParserRULE_mulOp                = 44
	SPL1ParserRULE_searchMulOp          = 45
	SPL1ParserRULE_sign                 = 46
	SPL1ParserRULE_identLike            = 47
	SPL1ParserRULE_exprWord             = 48
	SPL1ParserRULE_aggList              = 49
	SPL1ParserRULE_agg                  = 50
	SPL1ParserRULE_aggListNotSpan       = 51
	SPL1ParserRULE_aggNotSpan           = 52
	SPL1ParserRULE_aliasClause          = 53
	SPL1ParserRULE_byClause             = 54
	SPL1ParserRULE_timechartBody        = 55
	SPL1ParserRULE_timechartSpan        = 56
	SPL1ParserRULE_spanAssign           = 57
	SPL1ParserRULE_rexField             = 58
	SPL1ParserRULE_fieldExprList        = 59
	SPL1ParserRULE_fieldExpr            = 60
	SPL1ParserRULE_assignment           = 61
	SPL1ParserRULE_spanValue            = 62
	SPL1ParserRULE_tableItem            = 63
	SPL1ParserRULE_fieldsSign           = 64
	SPL1ParserRULE_fieldsItem           = 65
	SPL1ParserRULE_sortBody             = 66
	SPL1ParserRULE_sortItem             = 67
	SPL1ParserRULE_sortFirstItem        = 68
	SPL1ParserRULE_sortField            = 69
	SPL1ParserRULE_limitArg             = 70
	SPL1ParserRULE_limitTail            = 71
	SPL1ParserRULE_binItem              = 72
	SPL1ParserRULE_binField             = 73
	SPL1ParserRULE_identLikeNotSpan     = 74
	SPL1ParserRULE_identLikeNotLimit    = 75
)

// IStageCommandContext is an interface to support dynamic dispatch.
type IStageCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Command() ICommandContext
	StageEnd() IStageEndContext

	// IsStageCommandContext differentiates from other interfaces.
	IsStageCommandContext()
}

type StageCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageCommandContext() *StageCommandContext {
	var p = new(StageCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_stageCommand
	return p
}

func InitEmptyStageCommandContext(p *StageCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_stageCommand
}

func (*StageCommandContext) IsStageCommandContext() {}

func NewStageCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageCommandContext {
	var p = new(StageCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_stageCommand

	return p
}

func (s *StageCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *StageCommandContext) Command() ICommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *StageCommandContext) StageEnd() IStageEndContext {
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

func (s *StageCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageCommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitStageCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) StageCommand() (localctx IStageCommandContext) {
	localctx = NewStageCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SPL1ParserRULE_stageCommand)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Command()
	}
	{
		p.SetState(153)
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

// IStageSearchContext is an interface to support dynamic dispatch.
type IStageSearchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchExpr() ISearchExprContext
	StageEnd() IStageEndContext

	// IsStageSearchContext differentiates from other interfaces.
	IsStageSearchContext()
}

type StageSearchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStageSearchContext() *StageSearchContext {
	var p = new(StageSearchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_stageSearch
	return p
}

func InitEmptyStageSearchContext(p *StageSearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_stageSearch
}

func (*StageSearchContext) IsStageSearchContext() {}

func NewStageSearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageSearchContext {
	var p = new(StageSearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_stageSearch

	return p
}

func (s *StageSearchContext) GetParser() antlr.Parser { return s.parser }

func (s *StageSearchContext) SearchExpr() ISearchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchExprContext)
}

func (s *StageSearchContext) StageEnd() IStageEndContext {
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

func (s *StageSearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageSearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageSearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitStageSearch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) StageSearch() (localctx IStageSearchContext) {
	localctx = NewStageSearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SPL1ParserRULE_stageSearch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.SearchExpr()
	}
	{
		p.SetState(156)
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
	p.RuleIndex = SPL1ParserRULE_stageEnd
	return p
}

func InitEmptyStageEndContext(p *StageEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_stageEnd
}

func (*StageEndContext) IsStageEndContext() {}

func NewStageEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StageEndContext {
	var p = new(StageEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_stageEnd

	return p
}

func (s *StageEndContext) GetParser() antlr.Parser { return s.parser }

func (s *StageEndContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPIPE, 0)
}

func (s *StageEndContext) EOF() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEOF, 0)
}

func (s *StageEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StageEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitStageEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) StageEnd() (localctx IStageEndContext) {
	localctx = NewStageEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SPL1ParserRULE_stageEnd)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL1ParserEOF || _la == SPL1ParserPIPE) {
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

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) CopyAll(ctx *CommandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CmdFieldsContext struct {
	CommandContext
}

func NewCmdFieldsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdFieldsContext {
	var p = new(CmdFieldsContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdFieldsContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELDS, 0)
}

func (s *CmdFieldsContext) FieldsSign() IFieldsSignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsSignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldsSignContext)
}

func (s *CmdFieldsContext) AllFieldsItem() []IFieldsItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldsItemContext); ok {
			len++
		}
	}

	tst := make([]IFieldsItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldsItemContext); ok {
			tst[i] = t.(IFieldsItemContext)
			i++
		}
	}

	return tst
}

func (s *CmdFieldsContext) FieldsItem(i int) IFieldsItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldsItemContext); ok {
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

	return t.(IFieldsItemContext)
}

func (s *CmdFieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdFields(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdEvalContext struct {
	CommandContext
}

func NewCmdEvalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdEvalContext {
	var p = new(CmdEvalContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdEvalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdEvalContext) EVAL() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEVAL, 0)
}

func (s *CmdEvalContext) AllAssignment() []IAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssignmentContext); ok {
			tst[i] = t.(IAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *CmdEvalContext) Assignment(i int) IAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
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

	return t.(IAssignmentContext)
}

func (s *CmdEvalContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserCOMMA)
}

func (s *CmdEvalContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, i)
}

func (s *CmdEvalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdEval(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdSortContext struct {
	CommandContext
}

func NewCmdSortContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdSortContext {
	var p = new(CmdSortContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdSortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdSortContext) SORT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSORT, 0)
}

func (s *CmdSortContext) SortBody() ISortBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortBodyContext)
}

func (s *CmdSortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdSort(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdTableContext struct {
	CommandContext
}

func NewCmdTableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdTableContext {
	var p = new(CmdTableContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdTableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTABLE, 0)
}

func (s *CmdTableContext) AllTableItem() []ITableItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableItemContext); ok {
			len++
		}
	}

	tst := make([]ITableItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableItemContext); ok {
			tst[i] = t.(ITableItemContext)
			i++
		}
	}

	return tst
}

func (s *CmdTableContext) TableItem(i int) ITableItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableItemContext); ok {
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

	return t.(ITableItemContext)
}

func (s *CmdTableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdTable(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdSearchContext struct {
	CommandContext
}

func NewCmdSearchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdSearchContext {
	var p = new(CmdSearchContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdSearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdSearchContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSEARCH, 0)
}

func (s *CmdSearchContext) SearchExpr() ISearchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchExprContext)
}

func (s *CmdSearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdSearch(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdTimechartContext struct {
	CommandContext
}

func NewCmdTimechartContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdTimechartContext {
	var p = new(CmdTimechartContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdTimechartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdTimechartContext) TIMECHART() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTIMECHART, 0)
}

func (s *CmdTimechartContext) TimechartBody() ITimechartBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimechartBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimechartBodyContext)
}

func (s *CmdTimechartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdTimechart(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdStatsContext struct {
	CommandContext
}

func NewCmdStatsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdStatsContext {
	var p = new(CmdStatsContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdStatsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdStatsContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTATS, 0)
}

func (s *CmdStatsContext) AggList() IAggListContext {
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

func (s *CmdStatsContext) ByClause() IByClauseContext {
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

func (s *CmdStatsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdStats(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdWhereContext struct {
	CommandContext
}

func NewCmdWhereContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdWhereContext {
	var p = new(CmdWhereContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdWhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdWhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserWHERE, 0)
}

func (s *CmdWhereContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *CmdWhereContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdWhere(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdRexContext struct {
	CommandContext
}

func NewCmdRexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdRexContext {
	var p = new(CmdRexContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdRexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdRexContext) REX() antlr.TerminalNode {
	return s.GetToken(SPL1ParserREX, 0)
}

func (s *CmdRexContext) STRING() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTRING, 0)
}

func (s *CmdRexContext) RexField() IRexFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRexFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRexFieldContext)
}

func (s *CmdRexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdRex(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdBinContext struct {
	CommandContext
}

func NewCmdBinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdBinContext {
	var p = new(CmdBinContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdBinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdBinContext) BIN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBIN, 0)
}

func (s *CmdBinContext) AllBinItem() []IBinItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinItemContext); ok {
			len++
		}
	}

	tst := make([]IBinItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinItemContext); ok {
			tst[i] = t.(IBinItemContext)
			i++
		}
	}

	return tst
}

func (s *CmdBinContext) BinItem(i int) IBinItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinItemContext); ok {
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

	return t.(IBinItemContext)
}

func (s *CmdBinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdBin(s)

	default:
		return t.VisitChildren(s)
	}
}

type CmdHeadContext struct {
	CommandContext
}

func NewCmdHeadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CmdHeadContext {
	var p = new(CmdHeadContext)

	InitEmptyCommandContext(&p.CommandContext)
	p.parser = parser
	p.CopyAll(ctx.(*CommandContext))

	return p
}

func (s *CmdHeadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdHeadContext) HEAD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserHEAD, 0)
}

func (s *CmdHeadContext) LimitArg() ILimitArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitArgContext)
}

func (s *CmdHeadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitCmdHead(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SPL1ParserRULE_command)
	var _la int

	var _alt int

	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSEARCH:
		localctx = NewCmdSearchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Match(SPL1ParserSEARCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.SearchExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SPL1ParserWHERE:
		localctx = NewCmdWhereContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(SPL1ParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(165)
				p.ValueExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SPL1ParserSTATS:
		localctx = NewCmdStatsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Match(SPL1ParserSTATS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.AggList()
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SPL1ParserBY {
			{
				p.SetState(170)
				p.ByClause()
			}

		}

	case SPL1ParserTIMECHART:
		localctx = NewCmdTimechartContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(173)
			p.Match(SPL1ParserTIMECHART)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.TimechartBody()
		}

	case SPL1ParserEVAL:
		localctx = NewCmdEvalContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.Match(SPL1ParserEVAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Assignment()
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SPL1ParserCOMMA {
			{
				p.SetState(177)
				p.Match(SPL1ParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(178)
				p.Assignment()
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SPL1ParserREX:
		localctx = NewCmdRexContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(184)
			p.Match(SPL1ParserREX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(185)
				p.RexField()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(188)
			p.Match(SPL1ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserTABLE:
		localctx = NewCmdTableContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(189)
			p.Match(SPL1ParserTABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&915901785309182) != 0 {
			{
				p.SetState(190)
				p.TableItem()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SPL1ParserFIELDS:
		localctx = NewCmdFieldsContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(196)
			p.Match(SPL1ParserFIELDS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141562122076160) != 0 {
			{
				p.SetState(197)
				p.FieldsSign()
			}

		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562958552793086) != 0 {
			{
				p.SetState(200)
				p.FieldsItem()
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SPL1ParserSORT:
		localctx = NewCmdSortContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(206)
			p.Match(SPL1ParserSORT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.SortBody()
		}

	case SPL1ParserHEAD:
		localctx = NewCmdHeadContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(208)
			p.Match(SPL1ParserHEAD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(209)
				p.LimitArg()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case SPL1ParserBIN:
		localctx = NewCmdBinContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(212)
			p.Match(SPL1ParserBIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(213)
					p.BinItem()
				}

			}
			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
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

// ISearchExprContext is an interface to support dynamic dispatch.
type ISearchExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchAnd() ISearchAndContext
	AllSearchOrTail() []ISearchOrTailContext
	SearchOrTail(i int) ISearchOrTailContext

	// IsSearchExprContext differentiates from other interfaces.
	IsSearchExprContext()
}

type SearchExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchExprContext() *SearchExprContext {
	var p = new(SearchExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchExpr
	return p
}

func InitEmptySearchExprContext(p *SearchExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchExpr
}

func (*SearchExprContext) IsSearchExprContext() {}

func NewSearchExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchExprContext {
	var p = new(SearchExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchExpr

	return p
}

func (s *SearchExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchExprContext) SearchAnd() ISearchAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchAndContext)
}

func (s *SearchExprContext) AllSearchOrTail() []ISearchOrTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchOrTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchOrTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchOrTailContext); ok {
			tst[i] = t.(ISearchOrTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchExprContext) SearchOrTail(i int) ISearchOrTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOrTailContext); ok {
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

	return t.(ISearchOrTailContext)
}

func (s *SearchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchExpr() (localctx ISearchExprContext) {
	localctx = NewSearchExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SPL1ParserRULE_searchExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.SearchAnd()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				p.SearchOrTail()
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

// ISearchOrTailContext is an interface to support dynamic dispatch.
type ISearchOrTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode
	SearchAnd() ISearchAndContext
	AbsentOperand() IAbsentOperandContext

	// IsSearchOrTailContext differentiates from other interfaces.
	IsSearchOrTailContext()
}

type SearchOrTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOrTailContext() *SearchOrTailContext {
	var p = new(SearchOrTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOrTail
	return p
}

func InitEmptySearchOrTailContext(p *SearchOrTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOrTail
}

func (*SearchOrTailContext) IsSearchOrTailContext() {}

func NewSearchOrTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOrTailContext {
	var p = new(SearchOrTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOrTail

	return p
}

func (s *SearchOrTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOrTailContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *SearchOrTailContext) SearchAnd() ISearchAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchAndContext)
}

func (s *SearchOrTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SearchOrTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOrTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchOrTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchOrTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOrTail() (localctx ISearchOrTailContext) {
	localctx = NewSearchOrTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SPL1ParserRULE_searchOrTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(228)

	if !(p.at(SPL1ParserOR)) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserOR)", ""))
		goto errorExit
	}
	{
		p.SetState(229)
		p.Match(SPL1ParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(230)
			p.SearchAnd()
		}

	case 2:
		{
			p.SetState(231)
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

// ISearchAndContext is an interface to support dynamic dispatch.
type ISearchAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchNot() ISearchNotContext
	AllSearchAndPart() []ISearchAndPartContext
	SearchAndPart(i int) ISearchAndPartContext

	// IsSearchAndContext differentiates from other interfaces.
	IsSearchAndContext()
}

type SearchAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchAndContext() *SearchAndContext {
	var p = new(SearchAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAnd
	return p
}

func InitEmptySearchAndContext(p *SearchAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAnd
}

func (*SearchAndContext) IsSearchAndContext() {}

func NewSearchAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchAndContext {
	var p = new(SearchAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchAnd

	return p
}

func (s *SearchAndContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchAndContext) SearchNot() ISearchNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchNotContext)
}

func (s *SearchAndContext) AllSearchAndPart() []ISearchAndPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchAndPartContext); ok {
			len++
		}
	}

	tst := make([]ISearchAndPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchAndPartContext); ok {
			tst[i] = t.(ISearchAndPartContext)
			i++
		}
	}

	return tst
}

func (s *SearchAndContext) SearchAndPart(i int) ISearchAndPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAndPartContext); ok {
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

	return t.(ISearchAndPartContext)
}

func (s *SearchAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchAnd() (localctx ISearchAndContext) {
	localctx = NewSearchAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SPL1ParserRULE_searchAnd)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.SearchNot()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(235)
				p.SearchAndPart()
			}

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
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

// ISearchAndPartContext is an interface to support dynamic dispatch.
type ISearchAndPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchAndPartContext differentiates from other interfaces.
	IsSearchAndPartContext()
}

type SearchAndPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchAndPartContext() *SearchAndPartContext {
	var p = new(SearchAndPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAndPart
	return p
}

func InitEmptySearchAndPartContext(p *SearchAndPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAndPart
}

func (*SearchAndPartContext) IsSearchAndPartContext() {}

func NewSearchAndPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchAndPartContext {
	var p = new(SearchAndPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchAndPart

	return p
}

func (s *SearchAndPartContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchAndPartContext) CopyAll(ctx *SearchAndPartContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchAndPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchAndPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SAndExplicitContext struct {
	SearchAndPartContext
}

func NewSAndExplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SAndExplicitContext {
	var p = new(SAndExplicitContext)

	InitEmptySearchAndPartContext(&p.SearchAndPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchAndPartContext))

	return p
}

func (s *SAndExplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SAndExplicitContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *SAndExplicitContext) SearchNot() ISearchNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchNotContext)
}

func (s *SAndExplicitContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SAndExplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSAndExplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

type SAndImplicitContext struct {
	SearchAndPartContext
}

func NewSAndImplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SAndImplicitContext {
	var p = new(SAndImplicitContext)

	InitEmptySearchAndPartContext(&p.SearchAndPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchAndPartContext))

	return p
}

func (s *SAndImplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SAndImplicitContext) SearchNot() ISearchNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchNotContext)
}

func (s *SAndImplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSAndImplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchAndPart() (localctx ISearchAndPartContext) {
	localctx = NewSearchAndPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SPL1ParserRULE_searchAndPart)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSAndExplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(241)

		if !(p.at(SPL1ParserAND)) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserAND)", ""))
			goto errorExit
		}
		{
			p.SetState(242)
			p.Match(SPL1ParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(243)
				p.SearchNot()
			}

		case 2:
			{
				p.SetState(244)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewSAndImplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(247)

		if !(p.startsTerm()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.startsTerm()", ""))
			goto errorExit
		}
		{
			p.SetState(248)
			p.SearchNot()
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

// ISearchNotContext is an interface to support dynamic dispatch.
type ISearchNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchNotContext differentiates from other interfaces.
	IsSearchNotContext()
}

type SearchNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchNotContext() *SearchNotContext {
	var p = new(SearchNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchNot
	return p
}

func InitEmptySearchNotContext(p *SearchNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchNot
}

func (*SearchNotContext) IsSearchNotContext() {}

func NewSearchNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchNotContext {
	var p = new(SearchNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchNot

	return p
}

func (s *SearchNotContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchNotContext) CopyAll(ctx *SearchNotContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SNotPrimaryContext struct {
	SearchNotContext
}

func NewSNotPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SNotPrimaryContext {
	var p = new(SNotPrimaryContext)

	InitEmptySearchNotContext(&p.SearchNotContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchNotContext))

	return p
}

func (s *SNotPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SNotPrimaryContext) SearchComparison() ISearchComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchComparisonContext)
}

func (s *SNotPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSNotPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

type SNotUnaryContext struct {
	SearchNotContext
}

func NewSNotUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SNotUnaryContext {
	var p = new(SNotUnaryContext)

	InitEmptySearchNotContext(&p.SearchNotContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchNotContext))

	return p
}

func (s *SNotUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SNotUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT, 0)
}

func (s *SNotUnaryContext) BANG() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBANG, 0)
}

func (s *SNotUnaryContext) SearchNot() ISearchNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchNotContext)
}

func (s *SNotUnaryContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SNotUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSNotUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchNot() (localctx ISearchNotContext) {
	localctx = NewSearchNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SPL1ParserRULE_searchNot)
	var _la int

	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSNotUnaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SPL1ParserNOT || _la == SPL1ParserBANG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(252)
				p.SearchNot()
			}

		case 2:
			{
				p.SetState(253)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewSNotPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.SearchComparison()
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

// ISearchComparisonContext is an interface to support dynamic dispatch.
type ISearchComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchAdditive() ISearchAdditiveContext
	AllSearchComparisonTail() []ISearchComparisonTailContext
	SearchComparisonTail(i int) ISearchComparisonTailContext

	// IsSearchComparisonContext differentiates from other interfaces.
	IsSearchComparisonContext()
}

type SearchComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchComparisonContext() *SearchComparisonContext {
	var p = new(SearchComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchComparison
	return p
}

func InitEmptySearchComparisonContext(p *SearchComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchComparison
}

func (*SearchComparisonContext) IsSearchComparisonContext() {}

func NewSearchComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchComparisonContext {
	var p = new(SearchComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchComparison

	return p
}

func (s *SearchComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchComparisonContext) SearchAdditive() ISearchAdditiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAdditiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchAdditiveContext)
}

func (s *SearchComparisonContext) AllSearchComparisonTail() []ISearchComparisonTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchComparisonTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchComparisonTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchComparisonTailContext); ok {
			tst[i] = t.(ISearchComparisonTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchComparisonContext) SearchComparisonTail(i int) ISearchComparisonTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchComparisonTailContext); ok {
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

	return t.(ISearchComparisonTailContext)
}

func (s *SearchComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchComparison() (localctx ISearchComparisonContext) {
	localctx = NewSearchComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SPL1ParserRULE_searchComparison)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.SearchAdditive()
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(260)
				p.SearchComparisonTail()
			}

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

// ISearchComparisonTailContext is an interface to support dynamic dispatch.
type ISearchComparisonTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComparisonOp() IComparisonOpContext
	SearchOperand() ISearchOperandContext
	AbsentOperand() IAbsentOperandContext

	// IsSearchComparisonTailContext differentiates from other interfaces.
	IsSearchComparisonTailContext()
}

type SearchComparisonTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchComparisonTailContext() *SearchComparisonTailContext {
	var p = new(SearchComparisonTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchComparisonTail
	return p
}

func InitEmptySearchComparisonTailContext(p *SearchComparisonTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchComparisonTail
}

func (*SearchComparisonTailContext) IsSearchComparisonTailContext() {}

func NewSearchComparisonTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchComparisonTailContext {
	var p = new(SearchComparisonTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchComparisonTail

	return p
}

func (s *SearchComparisonTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchComparisonTailContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *SearchComparisonTailContext) SearchOperand() ISearchOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchOperandContext)
}

func (s *SearchComparisonTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SearchComparisonTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchComparisonTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchComparisonTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchComparisonTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchComparisonTail() (localctx ISearchComparisonTailContext) {
	localctx = NewSearchComparisonTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SPL1ParserRULE_searchComparisonTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(266)

	if !(p.atCompare()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atCompare()", ""))
		goto errorExit
	}
	{
		p.SetState(267)
		p.ComparisonOp()
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserDURATION, SPL1ParserNUMBER, SPL1ParserSTRING, SPL1ParserIDENT, SPL1ParserLPAREN, SPL1ParserLBRACKET, SPL1ParserPLUS, SPL1ParserMINUS, SPL1ParserSTAR, SPL1ParserNUMBER_TIGHT, SPL1ParserSTAR_TIGHT, SPL1ParserMINUS_TIGHT, SPL1ParserNOT_WORD:
		{
			p.SetState(268)
			p.SearchOperand()
		}

	case SPL1ParserEOF, SPL1ParserEQ2, SPL1ParserNE, SPL1ParserLE, SPL1ParserGE, SPL1ParserEQ, SPL1ParserLT, SPL1ParserGT, SPL1ParserPIPE, SPL1ParserCOMMA, SPL1ParserRPAREN, SPL1ParserRBRACKET:
		{
			p.SetState(269)
			p.AbsentOperand()
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

// ISearchOperandContext is an interface to support dynamic dispatch.
type ISearchOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchOperandContext differentiates from other interfaces.
	IsSearchOperandContext()
}

type SearchOperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOperandContext() *SearchOperandContext {
	var p = new(SearchOperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperand
	return p
}

func InitEmptySearchOperandContext(p *SearchOperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperand
}

func (*SearchOperandContext) IsSearchOperandContext() {}

func NewSearchOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOperandContext {
	var p = new(SearchOperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOperand

	return p
}

func (s *SearchOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOperandContext) CopyAll(ctx *SearchOperandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SOpStringContext struct {
	SearchOperandContext
}

func NewSOpStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SOpStringContext {
	var p = new(SOpStringContext)

	InitEmptySearchOperandContext(&p.SearchOperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchOperandContext))

	return p
}

func (s *SOpStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SOpStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTRING, 0)
}

func (s *SOpStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSOpString(s)

	default:
		return t.VisitChildren(s)
	}
}

type SOpExprContext struct {
	SearchOperandContext
}

func NewSOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SOpExprContext {
	var p = new(SOpExprContext)

	InitEmptySearchOperandContext(&p.SearchOperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchOperandContext))

	return p
}

func (s *SOpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SOpExprContext) SearchOperandExpr() ISearchOperandExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOperandExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchOperandExprContext)
}

func (s *SOpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SOpValueContext struct {
	SearchOperandContext
}

func NewSOpValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SOpValueContext {
	var p = new(SOpValueContext)

	InitEmptySearchOperandContext(&p.SearchOperandContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchOperandContext))

	return p
}

func (s *SOpValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SOpValueContext) SearchOperandValue() ISearchOperandValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOperandValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchOperandValueContext)
}

func (s *SOpValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSOpValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOperand() (localctx ISearchOperandContext) {
	localctx = NewSearchOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SPL1ParserRULE_searchOperand)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSTRING:
		localctx = NewSOpStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(SPL1ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNUMBER_TIGHT, SPL1ParserSTAR_TIGHT, SPL1ParserMINUS_TIGHT, SPL1ParserNOT_WORD:
		localctx = NewSOpValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.SearchOperandValue()
		}

	case SPL1ParserDURATION, SPL1ParserNUMBER, SPL1ParserLPAREN, SPL1ParserLBRACKET, SPL1ParserPLUS, SPL1ParserMINUS, SPL1ParserSTAR:
		localctx = NewSOpExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(274)
			p.SearchOperandExpr()
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

// ISearchOperandValueContext is an interface to support dynamic dispatch.
type ISearchOperandValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext
	AllRhsTail() []IRhsTailContext
	RhsTail(i int) IRhsTailContext
	NUMBER_TIGHT() antlr.TerminalNode
	STAR_TIGHT() antlr.TerminalNode
	Number() INumberContext
	Minus() IMinusContext
	Star() IStarContext
	MINUS_TIGHT() antlr.TerminalNode

	// IsSearchOperandValueContext differentiates from other interfaces.
	IsSearchOperandValueContext()
}

type SearchOperandValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOperandValueContext() *SearchOperandValueContext {
	var p = new(SearchOperandValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandValue
	return p
}

func InitEmptySearchOperandValueContext(p *SearchOperandValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandValue
}

func (*SearchOperandValueContext) IsSearchOperandValueContext() {}

func NewSearchOperandValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOperandValueContext {
	var p = new(SearchOperandValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOperandValue

	return p
}

func (s *SearchOperandValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOperandValueContext) IdentLike() IIdentLikeContext {
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

func (s *SearchOperandValueContext) AllRhsTail() []IRhsTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRhsTailContext); ok {
			len++
		}
	}

	tst := make([]IRhsTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRhsTailContext); ok {
			tst[i] = t.(IRhsTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchOperandValueContext) RhsTail(i int) IRhsTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRhsTailContext); ok {
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

	return t.(IRhsTailContext)
}

func (s *SearchOperandValueContext) NUMBER_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNUMBER_TIGHT, 0)
}

func (s *SearchOperandValueContext) STAR_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR_TIGHT, 0)
}

func (s *SearchOperandValueContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *SearchOperandValueContext) Minus() IMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinusContext)
}

func (s *SearchOperandValueContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *SearchOperandValueContext) MINUS_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserMINUS_TIGHT, 0)
}

func (s *SearchOperandValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOperandValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchOperandValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchOperandValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOperandValue() (localctx ISearchOperandValueContext) {
	localctx = NewSearchOperandValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SPL1ParserRULE_searchOperandValue)
	var _alt int

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.IdentLike()
		}
		p.SetState(281)
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
					p.SetState(278)
					p.RhsTail()
				}

			}
			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case SPL1ParserNUMBER_TIGHT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)
			p.Match(SPL1ParserNUMBER_TIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(285)
					p.RhsTail()
				}

			}
			p.SetState(290)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case SPL1ParserSTAR_TIGHT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Match(SPL1ParserSTAR_TIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
			{
				p.SetState(292)
				p.IdentLike()
			}

		case SPL1ParserNUMBER, SPL1ParserNUMBER_TIGHT:
			{
				p.SetState(293)
				p.Number()
			}

		case SPL1ParserMINUS, SPL1ParserMINUS_TIGHT:
			{
				p.SetState(294)
				p.Minus()
			}
			p.SetState(298)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
				{
					p.SetState(295)
					p.IdentLike()
				}

			case SPL1ParserNUMBER, SPL1ParserNUMBER_TIGHT:
				{
					p.SetState(296)
					p.Number()
				}

			case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
				{
					p.SetState(297)
					p.Star()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(302)
					p.RhsTail()
				}

			}
			p.SetState(307)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case SPL1ParserMINUS_TIGHT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(308)
			p.Match(SPL1ParserMINUS_TIGHT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
			{
				p.SetState(309)
				p.IdentLike()
			}

		case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
			{
				p.SetState(310)
				p.Star()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(313)
					p.RhsTail()
				}

			}
			p.SetState(318)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
			if p.HasError() {
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

// IRhsTailContext is an interface to support dynamic dispatch.
type IRhsTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Star() IStarContext
	IdentLike() IIdentLikeContext
	Number() INumberContext
	Minus() IMinusContext

	// IsRhsTailContext differentiates from other interfaces.
	IsRhsTailContext()
}

type RhsTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsTailContext() *RhsTailContext {
	var p = new(RhsTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_rhsTail
	return p
}

func InitEmptyRhsTailContext(p *RhsTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_rhsTail
}

func (*RhsTailContext) IsRhsTailContext() {}

func NewRhsTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsTailContext {
	var p = new(RhsTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_rhsTail

	return p
}

func (s *RhsTailContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsTailContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *RhsTailContext) IdentLike() IIdentLikeContext {
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

func (s *RhsTailContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *RhsTailContext) Minus() IMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinusContext)
}

func (s *RhsTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitRhsTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) RhsTail() (localctx IRhsTailContext) {
	localctx = NewRhsTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SPL1ParserRULE_rhsTail)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(321)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(322)
			p.Star()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(323)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(324)
			p.IdentLike()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(325)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(326)
			p.Number()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(327)

		if !(p.hyphenTail()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.hyphenTail()", ""))
			goto errorExit
		}
		{
			p.SetState(328)
			p.Minus()
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
			{
				p.SetState(329)
				p.IdentLike()
			}

		case SPL1ParserNUMBER, SPL1ParserNUMBER_TIGHT:
			{
				p.SetState(330)
				p.Number()
			}

		case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
			{
				p.SetState(331)
				p.Star()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

// ISearchOperandExprContext is an interface to support dynamic dispatch.
type ISearchOperandExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchOperandUnary() ISearchOperandUnaryContext
	AllSearchMulTail() []ISearchMulTailContext
	SearchMulTail(i int) ISearchMulTailContext
	AllSearchAddTail() []ISearchAddTailContext
	SearchAddTail(i int) ISearchAddTailContext

	// IsSearchOperandExprContext differentiates from other interfaces.
	IsSearchOperandExprContext()
}

type SearchOperandExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOperandExprContext() *SearchOperandExprContext {
	var p = new(SearchOperandExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandExpr
	return p
}

func InitEmptySearchOperandExprContext(p *SearchOperandExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandExpr
}

func (*SearchOperandExprContext) IsSearchOperandExprContext() {}

func NewSearchOperandExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOperandExprContext {
	var p = new(SearchOperandExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOperandExpr

	return p
}

func (s *SearchOperandExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOperandExprContext) SearchOperandUnary() ISearchOperandUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOperandUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchOperandUnaryContext)
}

func (s *SearchOperandExprContext) AllSearchMulTail() []ISearchMulTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchMulTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchMulTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchMulTailContext); ok {
			tst[i] = t.(ISearchMulTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchOperandExprContext) SearchMulTail(i int) ISearchMulTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchMulTailContext); ok {
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

	return t.(ISearchMulTailContext)
}

func (s *SearchOperandExprContext) AllSearchAddTail() []ISearchAddTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchAddTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchAddTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchAddTailContext); ok {
			tst[i] = t.(ISearchAddTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchOperandExprContext) SearchAddTail(i int) ISearchAddTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAddTailContext); ok {
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

	return t.(ISearchAddTailContext)
}

func (s *SearchOperandExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOperandExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchOperandExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchOperandExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOperandExpr() (localctx ISearchOperandExprContext) {
	localctx = NewSearchOperandExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SPL1ParserRULE_searchOperandExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.SearchOperandUnary()
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(337)
				p.SearchMulTail()
			}

		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(343)
				p.SearchAddTail()
			}

		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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

// ISearchOperandUnaryContext is an interface to support dynamic dispatch.
type ISearchOperandUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	SearchUnary() ISearchUnaryContext
	AbsentOperand() IAbsentOperandContext
	SearchOperandPrimary() ISearchOperandPrimaryContext

	// IsSearchOperandUnaryContext differentiates from other interfaces.
	IsSearchOperandUnaryContext()
}

type SearchOperandUnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOperandUnaryContext() *SearchOperandUnaryContext {
	var p = new(SearchOperandUnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandUnary
	return p
}

func InitEmptySearchOperandUnaryContext(p *SearchOperandUnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandUnary
}

func (*SearchOperandUnaryContext) IsSearchOperandUnaryContext() {}

func NewSearchOperandUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOperandUnaryContext {
	var p = new(SearchOperandUnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOperandUnary

	return p
}

func (s *SearchOperandUnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOperandUnaryContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserMINUS, 0)
}

func (s *SearchOperandUnaryContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPLUS, 0)
}

func (s *SearchOperandUnaryContext) SearchUnary() ISearchUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchUnaryContext)
}

func (s *SearchOperandUnaryContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SearchOperandUnaryContext) SearchOperandPrimary() ISearchOperandPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchOperandPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchOperandPrimaryContext)
}

func (s *SearchOperandUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOperandUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchOperandUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchOperandUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOperandUnary() (localctx ISearchOperandUnaryContext) {
	localctx = NewSearchOperandUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SPL1ParserRULE_searchOperandUnary)
	var _la int

	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserPLUS, SPL1ParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SPL1ParserPLUS || _la == SPL1ParserMINUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(350)
				p.SearchUnary()
			}

		case 2:
			{
				p.SetState(351)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case SPL1ParserDURATION, SPL1ParserNUMBER, SPL1ParserLPAREN, SPL1ParserLBRACKET, SPL1ParserSTAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.SearchOperandPrimary()
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

// ISearchOperandPrimaryContext is an interface to support dynamic dispatch.
type ISearchOperandPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	SearchExpr() ISearchExprContext
	RPAREN() antlr.TerminalNode
	DURATION() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	STAR() antlr.TerminalNode
	Subsearch() ISubsearchContext

	// IsSearchOperandPrimaryContext differentiates from other interfaces.
	IsSearchOperandPrimaryContext()
}

type SearchOperandPrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchOperandPrimaryContext() *SearchOperandPrimaryContext {
	var p = new(SearchOperandPrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandPrimary
	return p
}

func InitEmptySearchOperandPrimaryContext(p *SearchOperandPrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchOperandPrimary
}

func (*SearchOperandPrimaryContext) IsSearchOperandPrimaryContext() {}

func NewSearchOperandPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchOperandPrimaryContext {
	var p = new(SearchOperandPrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchOperandPrimary

	return p
}

func (s *SearchOperandPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchOperandPrimaryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *SearchOperandPrimaryContext) SearchExpr() ISearchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchExprContext)
}

func (s *SearchOperandPrimaryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *SearchOperandPrimaryContext) DURATION() antlr.TerminalNode {
	return s.GetToken(SPL1ParserDURATION, 0)
}

func (s *SearchOperandPrimaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNUMBER, 0)
}

func (s *SearchOperandPrimaryContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR, 0)
}

func (s *SearchOperandPrimaryContext) Subsearch() ISubsearchContext {
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

func (s *SearchOperandPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchOperandPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchOperandPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchOperandPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchOperandPrimary() (localctx ISearchOperandPrimaryContext) {
	localctx = NewSearchOperandPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SPL1ParserRULE_searchOperandPrimary)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(358)
			p.SearchExpr()
		}
		{
			p.SetState(359)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserDURATION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(361)
			p.Match(SPL1ParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(362)
			p.Match(SPL1ParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSTAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(363)
			p.Match(SPL1ParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserLBRACKET:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(364)
			p.Subsearch()
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

// ISearchAdditiveContext is an interface to support dynamic dispatch.
type ISearchAdditiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchMultiplicative() ISearchMultiplicativeContext
	AllSearchAddTail() []ISearchAddTailContext
	SearchAddTail(i int) ISearchAddTailContext

	// IsSearchAdditiveContext differentiates from other interfaces.
	IsSearchAdditiveContext()
}

type SearchAdditiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchAdditiveContext() *SearchAdditiveContext {
	var p = new(SearchAdditiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAdditive
	return p
}

func InitEmptySearchAdditiveContext(p *SearchAdditiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAdditive
}

func (*SearchAdditiveContext) IsSearchAdditiveContext() {}

func NewSearchAdditiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchAdditiveContext {
	var p = new(SearchAdditiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchAdditive

	return p
}

func (s *SearchAdditiveContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchAdditiveContext) SearchMultiplicative() ISearchMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchMultiplicativeContext)
}

func (s *SearchAdditiveContext) AllSearchAddTail() []ISearchAddTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchAddTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchAddTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchAddTailContext); ok {
			tst[i] = t.(ISearchAddTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchAdditiveContext) SearchAddTail(i int) ISearchAddTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchAddTailContext); ok {
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

	return t.(ISearchAddTailContext)
}

func (s *SearchAdditiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchAdditiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchAdditiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchAdditive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchAdditive() (localctx ISearchAdditiveContext) {
	localctx = NewSearchAdditiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SPL1ParserRULE_searchAdditive)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.SearchMultiplicative()
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(368)
				p.SearchAddTail()
			}

		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
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

// ISearchAddTailContext is an interface to support dynamic dispatch.
type ISearchAddTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddOp() IAddOpContext
	SearchMultiplicative() ISearchMultiplicativeContext
	AbsentOperand() IAbsentOperandContext

	// IsSearchAddTailContext differentiates from other interfaces.
	IsSearchAddTailContext()
}

type SearchAddTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchAddTailContext() *SearchAddTailContext {
	var p = new(SearchAddTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAddTail
	return p
}

func InitEmptySearchAddTailContext(p *SearchAddTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchAddTail
}

func (*SearchAddTailContext) IsSearchAddTailContext() {}

func NewSearchAddTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchAddTailContext {
	var p = new(SearchAddTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchAddTail

	return p
}

func (s *SearchAddTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchAddTailContext) AddOp() IAddOpContext {
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

func (s *SearchAddTailContext) SearchMultiplicative() ISearchMultiplicativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchMultiplicativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchMultiplicativeContext)
}

func (s *SearchAddTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SearchAddTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchAddTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchAddTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchAddTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchAddTail() (localctx ISearchAddTailContext) {
	localctx = NewSearchAddTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SPL1ParserRULE_searchAddTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(374)

	if !(p.atAdd()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atAdd()", ""))
		goto errorExit
	}
	{
		p.SetState(375)
		p.AddOp()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(376)
			p.SearchMultiplicative()
		}

	case 2:
		{
			p.SetState(377)
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

// ISearchMultiplicativeContext is an interface to support dynamic dispatch.
type ISearchMultiplicativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchUnary() ISearchUnaryContext
	AllSearchMulTail() []ISearchMulTailContext
	SearchMulTail(i int) ISearchMulTailContext

	// IsSearchMultiplicativeContext differentiates from other interfaces.
	IsSearchMultiplicativeContext()
}

type SearchMultiplicativeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchMultiplicativeContext() *SearchMultiplicativeContext {
	var p = new(SearchMultiplicativeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMultiplicative
	return p
}

func InitEmptySearchMultiplicativeContext(p *SearchMultiplicativeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMultiplicative
}

func (*SearchMultiplicativeContext) IsSearchMultiplicativeContext() {}

func NewSearchMultiplicativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchMultiplicativeContext {
	var p = new(SearchMultiplicativeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchMultiplicative

	return p
}

func (s *SearchMultiplicativeContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchMultiplicativeContext) SearchUnary() ISearchUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchUnaryContext)
}

func (s *SearchMultiplicativeContext) AllSearchMulTail() []ISearchMulTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchMulTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchMulTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchMulTailContext); ok {
			tst[i] = t.(ISearchMulTailContext)
			i++
		}
	}

	return tst
}

func (s *SearchMultiplicativeContext) SearchMulTail(i int) ISearchMulTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchMulTailContext); ok {
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

	return t.(ISearchMulTailContext)
}

func (s *SearchMultiplicativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchMultiplicativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchMultiplicativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchMultiplicative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchMultiplicative() (localctx ISearchMultiplicativeContext) {
	localctx = NewSearchMultiplicativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SPL1ParserRULE_searchMultiplicative)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.SearchUnary()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.SearchMulTail()
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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

// ISearchMulTailContext is an interface to support dynamic dispatch.
type ISearchMulTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SearchMulOp() ISearchMulOpContext
	SearchUnary() ISearchUnaryContext
	AbsentOperand() IAbsentOperandContext

	// IsSearchMulTailContext differentiates from other interfaces.
	IsSearchMulTailContext()
}

type SearchMulTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchMulTailContext() *SearchMulTailContext {
	var p = new(SearchMulTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMulTail
	return p
}

func InitEmptySearchMulTailContext(p *SearchMulTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMulTail
}

func (*SearchMulTailContext) IsSearchMulTailContext() {}

func NewSearchMulTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchMulTailContext {
	var p = new(SearchMulTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchMulTail

	return p
}

func (s *SearchMulTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchMulTailContext) SearchMulOp() ISearchMulOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchMulOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchMulOpContext)
}

func (s *SearchMulTailContext) SearchUnary() ISearchUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchUnaryContext)
}

func (s *SearchMulTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SearchMulTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchMulTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchMulTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchMulTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchMulTail() (localctx ISearchMulTailContext) {
	localctx = NewSearchMulTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SPL1ParserRULE_searchMulTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(387)

	if !(p.atSearchMul()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atSearchMul()", ""))
		goto errorExit
	}
	{
		p.SetState(388)
		p.SearchMulOp()
	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(389)
			p.SearchUnary()
		}

	case 2:
		{
			p.SetState(390)
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

// ISearchUnaryContext is an interface to support dynamic dispatch.
type ISearchUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchUnaryContext differentiates from other interfaces.
	IsSearchUnaryContext()
}

type SearchUnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchUnaryContext() *SearchUnaryContext {
	var p = new(SearchUnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchUnary
	return p
}

func InitEmptySearchUnaryContext(p *SearchUnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchUnary
}

func (*SearchUnaryContext) IsSearchUnaryContext() {}

func NewSearchUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchUnaryContext {
	var p = new(SearchUnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchUnary

	return p
}

func (s *SearchUnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchUnaryContext) CopyAll(ctx *SearchUnaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SUnarySignContext struct {
	SearchUnaryContext
}

func NewSUnarySignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SUnarySignContext {
	var p = new(SUnarySignContext)

	InitEmptySearchUnaryContext(&p.SearchUnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchUnaryContext))

	return p
}

func (s *SUnarySignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SUnarySignContext) Sign() ISignContext {
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

func (s *SUnarySignContext) SearchUnary() ISearchUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchUnaryContext)
}

func (s *SUnarySignContext) AbsentOperand() IAbsentOperandContext {
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

func (s *SUnarySignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSUnarySign(s)

	default:
		return t.VisitChildren(s)
	}
}

type SUnaryPrimaryContext struct {
	SearchUnaryContext
}

func NewSUnaryPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SUnaryPrimaryContext {
	var p = new(SUnaryPrimaryContext)

	InitEmptySearchUnaryContext(&p.SearchUnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchUnaryContext))

	return p
}

func (s *SUnaryPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SUnaryPrimaryContext) SearchPrimary() ISearchPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchPrimaryContext)
}

func (s *SUnaryPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSUnaryPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchUnary() (localctx ISearchUnaryContext) {
	localctx = NewSearchUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SPL1ParserRULE_searchUnary)
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSUnarySignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(393)
			p.Sign()
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(394)
				p.SearchUnary()
			}

		case 2:
			{
				p.SetState(395)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewSUnaryPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.SearchPrimary()
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

// ISearchPrimaryContext is an interface to support dynamic dispatch.
type ISearchPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchPrimaryContext differentiates from other interfaces.
	IsSearchPrimaryContext()
}

type SearchPrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchPrimaryContext() *SearchPrimaryContext {
	var p = new(SearchPrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchPrimary
	return p
}

func InitEmptySearchPrimaryContext(p *SearchPrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchPrimary
}

func (*SearchPrimaryContext) IsSearchPrimaryContext() {}

func NewSearchPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchPrimaryContext {
	var p = new(SearchPrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchPrimary

	return p
}

func (s *SearchPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchPrimaryContext) CopyAll(ctx *SearchPrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SPrimSubsearchContext struct {
	SearchPrimaryContext
}

func NewSPrimSubsearchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimSubsearchContext {
	var p = new(SPrimSubsearchContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimSubsearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimSubsearchContext) Subsearch() ISubsearchContext {
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

func (s *SPrimSubsearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimSubsearch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimNumberContext struct {
	SearchPrimaryContext
}

func NewSPrimNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimNumberContext {
	var p = new(SPrimNumberContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimNumberContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *SPrimNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimStarContext struct {
	SearchPrimaryContext
}

func NewSPrimStarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimStarContext {
	var p = new(SPrimStarContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimStarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimStarContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *SPrimStarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimStar(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimStringContext struct {
	SearchPrimaryContext
}

func NewSPrimStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimStringContext {
	var p = new(SPrimStringContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTRING, 0)
}

func (s *SPrimStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimString(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimWordContext struct {
	SearchPrimaryContext
}

func NewSPrimWordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimWordContext {
	var p = new(SPrimWordContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimWordContext) ExprWord() IExprWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprWordContext)
}

func (s *SPrimWordContext) AllSearchTermTail() []ISearchTermTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchTermTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchTermTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchTermTailContext); ok {
			tst[i] = t.(ISearchTermTailContext)
			i++
		}
	}

	return tst
}

func (s *SPrimWordContext) SearchTermTail(i int) ISearchTermTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchTermTailContext); ok {
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

	return t.(ISearchTermTailContext)
}

func (s *SPrimWordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimWord(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimStarRunContext struct {
	SearchPrimaryContext
}

func NewSPrimStarRunContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimStarRunContext {
	var p = new(SPrimStarRunContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimStarRunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimStarRunContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *SPrimStarRunContext) AllSearchTermTail() []ISearchTermTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchTermTailContext); ok {
			len++
		}
	}

	tst := make([]ISearchTermTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchTermTailContext); ok {
			tst[i] = t.(ISearchTermTailContext)
			i++
		}
	}

	return tst
}

func (s *SPrimStarRunContext) SearchTermTail(i int) ISearchTermTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchTermTailContext); ok {
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

	return t.(ISearchTermTailContext)
}

func (s *SPrimStarRunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimStarRun(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimDurationContext struct {
	SearchPrimaryContext
}

func NewSPrimDurationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimDurationContext {
	var p = new(SPrimDurationContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimDurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimDurationContext) DURATION() antlr.TerminalNode {
	return s.GetToken(SPL1ParserDURATION, 0)
}

func (s *SPrimDurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimDuration(s)

	default:
		return t.VisitChildren(s)
	}
}

type SPrimParenContext struct {
	SearchPrimaryContext
}

func NewSPrimParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SPrimParenContext {
	var p = new(SPrimParenContext)

	InitEmptySearchPrimaryContext(&p.SearchPrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchPrimaryContext))

	return p
}

func (s *SPrimParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SPrimParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *SPrimParenContext) SearchExpr() ISearchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchExprContext)
}

func (s *SPrimParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *SPrimParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSPrimParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchPrimary() (localctx ISearchPrimaryContext) {
	localctx = NewSearchPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SPL1ParserRULE_searchPrimary)
	var _alt int

	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSPrimParenContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)
			p.SearchExpr()
		}
		{
			p.SetState(403)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewSPrimStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Match(SPL1ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSPrimWordContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(406)
			p.ExprWord()
		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(407)
					p.SearchTermTail()
				}

			}
			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 4:
		localctx = NewSPrimNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(413)
			p.Number()
		}

	case 5:
		localctx = NewSPrimDurationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(414)
			p.Match(SPL1ParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewSPrimStarRunContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(415)

		if !(p.termStarRun()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.termStarRun()", ""))
			goto errorExit
		}
		{
			p.SetState(416)
			p.Star()
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(417)
					p.SearchTermTail()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 7:
		localctx = NewSPrimStarContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(422)
			p.Star()
		}

	case 8:
		localctx = NewSPrimSubsearchContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(423)
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

// ISearchTermTailContext is an interface to support dynamic dispatch.
type ISearchTermTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR_TERM() antlr.TerminalNode
	IdentLike() IIdentLikeContext
	Number() INumberContext

	// IsSearchTermTailContext differentiates from other interfaces.
	IsSearchTermTailContext()
}

type SearchTermTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchTermTailContext() *SearchTermTailContext {
	var p = new(SearchTermTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchTermTail
	return p
}

func InitEmptySearchTermTailContext(p *SearchTermTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchTermTail
}

func (*SearchTermTailContext) IsSearchTermTailContext() {}

func NewSearchTermTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchTermTailContext {
	var p = new(SearchTermTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchTermTail

	return p
}

func (s *SearchTermTailContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchTermTailContext) STAR_TERM() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR_TERM, 0)
}

func (s *SearchTermTailContext) IdentLike() IIdentLikeContext {
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

func (s *SearchTermTailContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *SearchTermTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchTermTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchTermTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchTermTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchTermTail() (localctx ISearchTermTailContext) {
	localctx = NewSearchTermTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SPL1ParserRULE_searchTermTail)
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(426)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(427)
			p.Match(SPL1ParserSTAR_TERM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(428)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(429)
			p.IdentLike()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(430)

		if !(p.touchingPrev()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.touchingPrev()", ""))
			goto errorExit
		}
		{
			p.SetState(431)
			p.Number()
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

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueAnd() IValueAndContext
	AllValueOrTail() []IValueOrTailContext
	ValueOrTail(i int) IValueOrTailContext

	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueExpr
	return p
}

func InitEmptyValueExprContext(p *ValueExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueExpr
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) ValueAnd() IValueAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueAndContext)
}

func (s *ValueExprContext) AllValueOrTail() []IValueOrTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueOrTailContext); ok {
			len++
		}
	}

	tst := make([]IValueOrTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueOrTailContext); ok {
			tst[i] = t.(IValueOrTailContext)
			i++
		}
	}

	return tst
}

func (s *ValueExprContext) ValueOrTail(i int) IValueOrTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueOrTailContext); ok {
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

	return t.(IValueOrTailContext)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueExpr() (localctx IValueExprContext) {
	localctx = NewValueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SPL1ParserRULE_valueExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		p.ValueAnd()
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(435)
				p.ValueOrTail()
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
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

// IValueOrTailContext is an interface to support dynamic dispatch.
type IValueOrTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode
	ValueAnd() IValueAndContext
	AbsentOperand() IAbsentOperandContext

	// IsValueOrTailContext differentiates from other interfaces.
	IsValueOrTailContext()
}

type ValueOrTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueOrTailContext() *ValueOrTailContext {
	var p = new(ValueOrTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueOrTail
	return p
}

func InitEmptyValueOrTailContext(p *ValueOrTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueOrTail
}

func (*ValueOrTailContext) IsValueOrTailContext() {}

func NewValueOrTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueOrTailContext {
	var p = new(ValueOrTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueOrTail

	return p
}

func (s *ValueOrTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueOrTailContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *ValueOrTailContext) ValueAnd() IValueAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueAndContext)
}

func (s *ValueOrTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *ValueOrTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueOrTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueOrTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueOrTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueOrTail() (localctx IValueOrTailContext) {
	localctx = NewValueOrTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SPL1ParserRULE_valueOrTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(441)

	if !(p.at(SPL1ParserOR)) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserOR)", ""))
		goto errorExit
	}
	{
		p.SetState(442)
		p.Match(SPL1ParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(443)
			p.ValueAnd()
		}

	case 2:
		{
			p.SetState(444)
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

// IValueAndContext is an interface to support dynamic dispatch.
type IValueAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueNot() IValueNotContext
	AllValueAndPart() []IValueAndPartContext
	ValueAndPart(i int) IValueAndPartContext

	// IsValueAndContext differentiates from other interfaces.
	IsValueAndContext()
}

type ValueAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAndContext() *ValueAndContext {
	var p = new(ValueAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAnd
	return p
}

func InitEmptyValueAndContext(p *ValueAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAnd
}

func (*ValueAndContext) IsValueAndContext() {}

func NewValueAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAndContext {
	var p = new(ValueAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueAnd

	return p
}

func (s *ValueAndContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAndContext) ValueNot() IValueNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueNotContext)
}

func (s *ValueAndContext) AllValueAndPart() []IValueAndPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueAndPartContext); ok {
			len++
		}
	}

	tst := make([]IValueAndPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueAndPartContext); ok {
			tst[i] = t.(IValueAndPartContext)
			i++
		}
	}

	return tst
}

func (s *ValueAndContext) ValueAndPart(i int) IValueAndPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAndPartContext); ok {
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

	return t.(IValueAndPartContext)
}

func (s *ValueAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueAndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueAnd() (localctx IValueAndContext) {
	localctx = NewValueAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SPL1ParserRULE_valueAnd)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.ValueNot()
	}
	p.SetState(451)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(448)
				p.ValueAndPart()
			}

		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
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

// IValueAndPartContext is an interface to support dynamic dispatch.
type IValueAndPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueAndPartContext differentiates from other interfaces.
	IsValueAndPartContext()
}

type ValueAndPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAndPartContext() *ValueAndPartContext {
	var p = new(ValueAndPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAndPart
	return p
}

func InitEmptyValueAndPartContext(p *ValueAndPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAndPart
}

func (*ValueAndPartContext) IsValueAndPartContext() {}

func NewValueAndPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAndPartContext {
	var p = new(ValueAndPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueAndPart

	return p
}

func (s *ValueAndPartContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAndPartContext) CopyAll(ctx *ValueAndPartContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueAndPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAndPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VAndImplicitContext struct {
	ValueAndPartContext
}

func NewVAndImplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VAndImplicitContext {
	var p = new(VAndImplicitContext)

	InitEmptyValueAndPartContext(&p.ValueAndPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueAndPartContext))

	return p
}

func (s *VAndImplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VAndImplicitContext) ValueNot() IValueNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueNotContext)
}

func (s *VAndImplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVAndImplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

type VAndExplicitContext struct {
	ValueAndPartContext
}

func NewVAndExplicitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VAndExplicitContext {
	var p = new(VAndExplicitContext)

	InitEmptyValueAndPartContext(&p.ValueAndPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueAndPartContext))

	return p
}

func (s *VAndExplicitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VAndExplicitContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *VAndExplicitContext) ValueNot() IValueNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueNotContext)
}

func (s *VAndExplicitContext) AbsentOperand() IAbsentOperandContext {
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

func (s *VAndExplicitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVAndExplicit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueAndPart() (localctx IValueAndPartContext) {
	localctx = NewValueAndPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SPL1ParserRULE_valueAndPart)
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVAndExplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(454)

		if !(p.at(SPL1ParserAND)) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserAND)", ""))
			goto errorExit
		}
		{
			p.SetState(455)
			p.Match(SPL1ParserAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(456)
				p.ValueNot()
			}

		case 2:
			{
				p.SetState(457)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewVAndImplicitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(460)

		if !(p.startsTerm()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.startsTerm()", ""))
			goto errorExit
		}
		{
			p.SetState(461)
			p.ValueNot()
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

// IValueNotContext is an interface to support dynamic dispatch.
type IValueNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueNotContext differentiates from other interfaces.
	IsValueNotContext()
}

type ValueNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueNotContext() *ValueNotContext {
	var p = new(ValueNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueNot
	return p
}

func InitEmptyValueNotContext(p *ValueNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueNot
}

func (*ValueNotContext) IsValueNotContext() {}

func NewValueNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueNotContext {
	var p = new(ValueNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueNot

	return p
}

func (s *ValueNotContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueNotContext) CopyAll(ctx *ValueNotContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VNotUnaryContext struct {
	ValueNotContext
}

func NewVNotUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VNotUnaryContext {
	var p = new(VNotUnaryContext)

	InitEmptyValueNotContext(&p.ValueNotContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueNotContext))

	return p
}

func (s *VNotUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VNotUnaryContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT, 0)
}

func (s *VNotUnaryContext) BANG() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBANG, 0)
}

func (s *VNotUnaryContext) ValueNot() IValueNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueNotContext)
}

func (s *VNotUnaryContext) AbsentOperand() IAbsentOperandContext {
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

func (s *VNotUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVNotUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type VNotPrimaryContext struct {
	ValueNotContext
}

func NewVNotPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VNotPrimaryContext {
	var p = new(VNotPrimaryContext)

	InitEmptyValueNotContext(&p.ValueNotContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueNotContext))

	return p
}

func (s *VNotPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VNotPrimaryContext) ValueComparison() IValueComparisonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueComparisonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueComparisonContext)
}

func (s *VNotPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVNotPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueNot() (localctx IValueNotContext) {
	localctx = NewValueNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SPL1ParserRULE_valueNot)
	var _la int

	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVNotUnaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SPL1ParserNOT || _la == SPL1ParserBANG) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(465)
				p.ValueNot()
			}

		case 2:
			{
				p.SetState(466)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewVNotPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.ValueComparison()
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

// IValueComparisonContext is an interface to support dynamic dispatch.
type IValueComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueAdd() IValueAddContext
	AllValueComparisonTail() []IValueComparisonTailContext
	ValueComparisonTail(i int) IValueComparisonTailContext

	// IsValueComparisonContext differentiates from other interfaces.
	IsValueComparisonContext()
}

type ValueComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueComparisonContext() *ValueComparisonContext {
	var p = new(ValueComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueComparison
	return p
}

func InitEmptyValueComparisonContext(p *ValueComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueComparison
}

func (*ValueComparisonContext) IsValueComparisonContext() {}

func NewValueComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueComparisonContext {
	var p = new(ValueComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueComparison

	return p
}

func (s *ValueComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueComparisonContext) ValueAdd() IValueAddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueAddContext)
}

func (s *ValueComparisonContext) AllValueComparisonTail() []IValueComparisonTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueComparisonTailContext); ok {
			len++
		}
	}

	tst := make([]IValueComparisonTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueComparisonTailContext); ok {
			tst[i] = t.(IValueComparisonTailContext)
			i++
		}
	}

	return tst
}

func (s *ValueComparisonContext) ValueComparisonTail(i int) IValueComparisonTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueComparisonTailContext); ok {
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

	return t.(IValueComparisonTailContext)
}

func (s *ValueComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueComparison() (localctx IValueComparisonContext) {
	localctx = NewValueComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SPL1ParserRULE_valueComparison)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		p.ValueAdd()
	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(473)
				p.ValueComparisonTail()
			}

		}
		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext())
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

// IValueComparisonTailContext is an interface to support dynamic dispatch.
type IValueComparisonTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ComparisonOp() IComparisonOpContext
	ValueAdd() IValueAddContext
	AbsentOperand() IAbsentOperandContext

	// IsValueComparisonTailContext differentiates from other interfaces.
	IsValueComparisonTailContext()
}

type ValueComparisonTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueComparisonTailContext() *ValueComparisonTailContext {
	var p = new(ValueComparisonTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueComparisonTail
	return p
}

func InitEmptyValueComparisonTailContext(p *ValueComparisonTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueComparisonTail
}

func (*ValueComparisonTailContext) IsValueComparisonTailContext() {}

func NewValueComparisonTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueComparisonTailContext {
	var p = new(ValueComparisonTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueComparisonTail

	return p
}

func (s *ValueComparisonTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueComparisonTailContext) ComparisonOp() IComparisonOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparisonOpContext)
}

func (s *ValueComparisonTailContext) ValueAdd() IValueAddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueAddContext)
}

func (s *ValueComparisonTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *ValueComparisonTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueComparisonTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueComparisonTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueComparisonTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueComparisonTail() (localctx IValueComparisonTailContext) {
	localctx = NewValueComparisonTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SPL1ParserRULE_valueComparisonTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(479)

	if !(p.atCompare()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atCompare()", ""))
		goto errorExit
	}
	{
		p.SetState(480)
		p.ComparisonOp()
	}
	p.SetState(483)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(481)
			p.ValueAdd()
		}

	case 2:
		{
			p.SetState(482)
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

// IValueAddContext is an interface to support dynamic dispatch.
type IValueAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueMul() IValueMulContext
	AllValueAddTail() []IValueAddTailContext
	ValueAddTail(i int) IValueAddTailContext

	// IsValueAddContext differentiates from other interfaces.
	IsValueAddContext()
}

type ValueAddContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAddContext() *ValueAddContext {
	var p = new(ValueAddContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAdd
	return p
}

func InitEmptyValueAddContext(p *ValueAddContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAdd
}

func (*ValueAddContext) IsValueAddContext() {}

func NewValueAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAddContext {
	var p = new(ValueAddContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueAdd

	return p
}

func (s *ValueAddContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAddContext) ValueMul() IValueMulContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueMulContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueMulContext)
}

func (s *ValueAddContext) AllValueAddTail() []IValueAddTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueAddTailContext); ok {
			len++
		}
	}

	tst := make([]IValueAddTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueAddTailContext); ok {
			tst[i] = t.(IValueAddTailContext)
			i++
		}
	}

	return tst
}

func (s *ValueAddContext) ValueAddTail(i int) IValueAddTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueAddTailContext); ok {
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

	return t.(IValueAddTailContext)
}

func (s *ValueAddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueAddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueAdd() (localctx IValueAddContext) {
	localctx = NewValueAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SPL1ParserRULE_valueAdd)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		p.ValueMul()
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(486)
				p.ValueAddTail()
			}

		}
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
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

// IValueAddTailContext is an interface to support dynamic dispatch.
type IValueAddTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AddOp() IAddOpContext
	ValueMul() IValueMulContext
	AbsentOperand() IAbsentOperandContext

	// IsValueAddTailContext differentiates from other interfaces.
	IsValueAddTailContext()
}

type ValueAddTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAddTailContext() *ValueAddTailContext {
	var p = new(ValueAddTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAddTail
	return p
}

func InitEmptyValueAddTailContext(p *ValueAddTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueAddTail
}

func (*ValueAddTailContext) IsValueAddTailContext() {}

func NewValueAddTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAddTailContext {
	var p = new(ValueAddTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueAddTail

	return p
}

func (s *ValueAddTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAddTailContext) AddOp() IAddOpContext {
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

func (s *ValueAddTailContext) ValueMul() IValueMulContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueMulContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueMulContext)
}

func (s *ValueAddTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *ValueAddTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAddTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueAddTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueAddTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueAddTail() (localctx IValueAddTailContext) {
	localctx = NewValueAddTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SPL1ParserRULE_valueAddTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(492)

	if !(p.atAdd()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atAdd()", ""))
		goto errorExit
	}
	{
		p.SetState(493)
		p.AddOp()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(494)
			p.ValueMul()
		}

	case 2:
		{
			p.SetState(495)
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

// IValueMulContext is an interface to support dynamic dispatch.
type IValueMulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValueUnary() IValueUnaryContext
	AllValueMulTail() []IValueMulTailContext
	ValueMulTail(i int) IValueMulTailContext

	// IsValueMulContext differentiates from other interfaces.
	IsValueMulContext()
}

type ValueMulContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueMulContext() *ValueMulContext {
	var p = new(ValueMulContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueMul
	return p
}

func InitEmptyValueMulContext(p *ValueMulContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueMul
}

func (*ValueMulContext) IsValueMulContext() {}

func NewValueMulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueMulContext {
	var p = new(ValueMulContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueMul

	return p
}

func (s *ValueMulContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueMulContext) ValueUnary() IValueUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueUnaryContext)
}

func (s *ValueMulContext) AllValueMulTail() []IValueMulTailContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueMulTailContext); ok {
			len++
		}
	}

	tst := make([]IValueMulTailContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueMulTailContext); ok {
			tst[i] = t.(IValueMulTailContext)
			i++
		}
	}

	return tst
}

func (s *ValueMulContext) ValueMulTail(i int) IValueMulTailContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueMulTailContext); ok {
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

	return t.(IValueMulTailContext)
}

func (s *ValueMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueMulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueMulContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueMul(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueMul() (localctx IValueMulContext) {
	localctx = NewValueMulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SPL1ParserRULE_valueMul)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.ValueUnary()
	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(499)
				p.ValueMulTail()
			}

		}
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
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

// IValueMulTailContext is an interface to support dynamic dispatch.
type IValueMulTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MulOp() IMulOpContext
	ValueUnary() IValueUnaryContext
	AbsentOperand() IAbsentOperandContext

	// IsValueMulTailContext differentiates from other interfaces.
	IsValueMulTailContext()
}

type ValueMulTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueMulTailContext() *ValueMulTailContext {
	var p = new(ValueMulTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueMulTail
	return p
}

func InitEmptyValueMulTailContext(p *ValueMulTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueMulTail
}

func (*ValueMulTailContext) IsValueMulTailContext() {}

func NewValueMulTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueMulTailContext {
	var p = new(ValueMulTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueMulTail

	return p
}

func (s *ValueMulTailContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueMulTailContext) MulOp() IMulOpContext {
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

func (s *ValueMulTailContext) ValueUnary() IValueUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueUnaryContext)
}

func (s *ValueMulTailContext) AbsentOperand() IAbsentOperandContext {
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

func (s *ValueMulTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueMulTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueMulTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitValueMulTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueMulTail() (localctx IValueMulTailContext) {
	localctx = NewValueMulTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SPL1ParserRULE_valueMulTail)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(505)

	if !(p.atMul()) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.atMul()", ""))
		goto errorExit
	}
	{
		p.SetState(506)
		p.MulOp()
	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(507)
			p.ValueUnary()
		}

	case 2:
		{
			p.SetState(508)
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

// IValueUnaryContext is an interface to support dynamic dispatch.
type IValueUnaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueUnaryContext differentiates from other interfaces.
	IsValueUnaryContext()
}

type ValueUnaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueUnaryContext() *ValueUnaryContext {
	var p = new(ValueUnaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueUnary
	return p
}

func InitEmptyValueUnaryContext(p *ValueUnaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valueUnary
}

func (*ValueUnaryContext) IsValueUnaryContext() {}

func NewValueUnaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueUnaryContext {
	var p = new(ValueUnaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valueUnary

	return p
}

func (s *ValueUnaryContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueUnaryContext) CopyAll(ctx *ValueUnaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueUnaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VUnarySignContext struct {
	ValueUnaryContext
}

func NewVUnarySignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VUnarySignContext {
	var p = new(VUnarySignContext)

	InitEmptyValueUnaryContext(&p.ValueUnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueUnaryContext))

	return p
}

func (s *VUnarySignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VUnarySignContext) Sign() ISignContext {
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

func (s *VUnarySignContext) ValueUnary() IValueUnaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueUnaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueUnaryContext)
}

func (s *VUnarySignContext) AbsentOperand() IAbsentOperandContext {
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

func (s *VUnarySignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVUnarySign(s)

	default:
		return t.VisitChildren(s)
	}
}

type VUnaryPrimaryContext struct {
	ValueUnaryContext
}

func NewVUnaryPrimaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VUnaryPrimaryContext {
	var p = new(VUnaryPrimaryContext)

	InitEmptyValueUnaryContext(&p.ValueUnaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueUnaryContext))

	return p
}

func (s *VUnaryPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VUnaryPrimaryContext) ValuePrimary() IValuePrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuePrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuePrimaryContext)
}

func (s *VUnaryPrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVUnaryPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValueUnary() (localctx IValueUnaryContext) {
	localctx = NewValueUnaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SPL1ParserRULE_valueUnary)
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVUnarySignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(511)
			p.Sign()
		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(512)
				p.ValueUnary()
			}

		case 2:
			{
				p.SetState(513)
				p.AbsentOperand()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	case 2:
		localctx = NewVUnaryPrimaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(516)
			p.ValuePrimary()
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

// IValuePrimaryContext is an interface to support dynamic dispatch.
type IValuePrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValuePrimaryContext differentiates from other interfaces.
	IsValuePrimaryContext()
}

type ValuePrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuePrimaryContext() *ValuePrimaryContext {
	var p = new(ValuePrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valuePrimary
	return p
}

func InitEmptyValuePrimaryContext(p *ValuePrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_valuePrimary
}

func (*ValuePrimaryContext) IsValuePrimaryContext() {}

func NewValuePrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuePrimaryContext {
	var p = new(ValuePrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_valuePrimary

	return p
}

func (s *ValuePrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuePrimaryContext) CopyAll(ctx *ValuePrimaryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValuePrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuePrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VPrimStringContext struct {
	ValuePrimaryContext
}

func NewVPrimStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimStringContext {
	var p = new(VPrimStringContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTRING, 0)
}

func (s *VPrimStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimString(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimDurationContext struct {
	ValuePrimaryContext
}

func NewVPrimDurationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimDurationContext {
	var p = new(VPrimDurationContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimDurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimDurationContext) DURATION() antlr.TerminalNode {
	return s.GetToken(SPL1ParserDURATION, 0)
}

func (s *VPrimDurationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimDuration(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimWordContext struct {
	ValuePrimaryContext
}

func NewVPrimWordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimWordContext {
	var p = new(VPrimWordContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimWordContext) ExprWord() IExprWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprWordContext)
}

func (s *VPrimWordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimWord(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimCallContext struct {
	ValuePrimaryContext
}

func NewVPrimCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimCallContext {
	var p = new(VPrimCallContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimCallContext) ExprWord() IExprWordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprWordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprWordContext)
}

func (s *VPrimCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *VPrimCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *VPrimCallContext) AllValueExpr() []IValueExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueExprContext); ok {
			len++
		}
	}

	tst := make([]IValueExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueExprContext); ok {
			tst[i] = t.(IValueExprContext)
			i++
		}
	}

	return tst
}

func (s *VPrimCallContext) ValueExpr(i int) IValueExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
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

	return t.(IValueExprContext)
}

func (s *VPrimCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserCOMMA)
}

func (s *VPrimCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, i)
}

func (s *VPrimCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimNumberContext struct {
	ValuePrimaryContext
}

func NewVPrimNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimNumberContext {
	var p = new(VPrimNumberContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimNumberContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *VPrimNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimSubsearchContext struct {
	ValuePrimaryContext
}

func NewVPrimSubsearchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimSubsearchContext {
	var p = new(VPrimSubsearchContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimSubsearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimSubsearchContext) Subsearch() ISubsearchContext {
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

func (s *VPrimSubsearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimSubsearch(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimParenContext struct {
	ValuePrimaryContext
}

func NewVPrimParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimParenContext {
	var p = new(VPrimParenContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimParenContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *VPrimParenContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *VPrimParenContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *VPrimParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type VPrimStarContext struct {
	ValuePrimaryContext
}

func NewVPrimStarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VPrimStarContext {
	var p = new(VPrimStarContext)

	InitEmptyValuePrimaryContext(&p.ValuePrimaryContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValuePrimaryContext))

	return p
}

func (s *VPrimStarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VPrimStarContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *VPrimStarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitVPrimStar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ValuePrimary() (localctx IValuePrimaryContext) {
	localctx = NewValuePrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SPL1ParserRULE_valuePrimary)
	var _la int

	p.SetState(544)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVPrimParenContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(520)
			p.ValueExpr()
		}
		{
			p.SetState(521)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewVPrimCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(523)

		if !(p.callAhead()) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.callAhead()", ""))
			goto errorExit
		}
		{
			p.SetState(524)
			p.ExprWord()
		}
		{
			p.SetState(525)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(534)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(526)
				p.ValueExpr()
			}
			p.SetState(531)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == SPL1ParserCOMMA {
				{
					p.SetState(527)
					p.Match(SPL1ParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(528)
					p.ValueExpr()
				}

				p.SetState(533)
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
			p.SetState(536)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVPrimWordContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(538)
			p.ExprWord()
		}

	case 4:
		localctx = NewVPrimStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(539)
			p.Match(SPL1ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewVPrimNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(540)
			p.Number()
		}

	case 6:
		localctx = NewVPrimDurationContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(541)
			p.Match(SPL1ParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewVPrimStarContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(542)
			p.Star()
		}

	case 8:
		localctx = NewVPrimSubsearchContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(543)
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
	p.RuleIndex = SPL1ParserRULE_subsearch
	return p
}

func InitEmptySubsearchContext(p *SubsearchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_subsearch
}

func (*SubsearchContext) IsSubsearchContext() {}

func NewSubsearchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubsearchContext {
	var p = new(SubsearchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_subsearch

	return p
}

func (s *SubsearchContext) GetParser() antlr.Parser { return s.parser }

func (s *SubsearchContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLBRACKET, 0)
}

func (s *SubsearchContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserPIPE)
}

func (s *SubsearchContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserPIPE, i)
}

func (s *SubsearchContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserEOF)
}

func (s *SubsearchContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserEOF, i)
}

func (s *SubsearchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubsearchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubsearchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSubsearch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Subsearch() (localctx ISubsearchContext) {
	localctx = NewSubsearchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SPL1ParserRULE_subsearch)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(SPL1ParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.abortSubsearch()
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(548)
				_la = p.GetTokenStream().LA(1)

				if _la <= 0 || _la == SPL1ParserEOF || _la == SPL1ParserPIPE {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext())
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
	p.RuleIndex = SPL1ParserRULE_absentOperand
	return p
}

func InitEmptyAbsentOperandContext(p *AbsentOperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_absentOperand
}

func (*AbsentOperandContext) IsAbsentOperandContext() {}

func NewAbsentOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsentOperandContext {
	var p = new(AbsentOperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_absentOperand

	return p
}

func (s *AbsentOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsentOperandContext) EQ2() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ2, 0)
}

func (s *AbsentOperandContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *AbsentOperandContext) NE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNE, 0)
}

func (s *AbsentOperandContext) LT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLT, 0)
}

func (s *AbsentOperandContext) LE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLE, 0)
}

func (s *AbsentOperandContext) GT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserGT, 0)
}

func (s *AbsentOperandContext) GE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserGE, 0)
}

func (s *AbsentOperandContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, 0)
}

func (s *AbsentOperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *AbsentOperandContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRBRACKET, 0)
}

func (s *AbsentOperandContext) PIPE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPIPE, 0)
}

func (s *AbsentOperandContext) EOF() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEOF, 0)
}

func (s *AbsentOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsentOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsentOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAbsentOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AbsentOperand() (localctx IAbsentOperandContext) {
	localctx = NewAbsentOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SPL1ParserRULE_absentOperand)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.abortMissingOperand()
	{
		p.SetState(555)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&373628600321) != 0) {
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

// IComparisonOpContext is an interface to support dynamic dispatch.
type IComparisonOpContext interface {
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

	// IsComparisonOpContext differentiates from other interfaces.
	IsComparisonOpContext()
}

type ComparisonOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOpContext() *ComparisonOpContext {
	var p = new(ComparisonOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_comparisonOp
	return p
}

func InitEmptyComparisonOpContext(p *ComparisonOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_comparisonOp
}

func (*ComparisonOpContext) IsComparisonOpContext() {}

func NewComparisonOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOpContext {
	var p = new(ComparisonOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_comparisonOp

	return p
}

func (s *ComparisonOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOpContext) EQ2() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ2, 0)
}

func (s *ComparisonOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *ComparisonOpContext) NE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNE, 0)
}

func (s *ComparisonOpContext) LT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLT, 0)
}

func (s *ComparisonOpContext) LE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLE, 0)
}

func (s *ComparisonOpContext) GT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserGT, 0)
}

func (s *ComparisonOpContext) GE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserGE, 0)
}

func (s *ComparisonOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitComparisonOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ComparisonOp() (localctx IComparisonOpContext) {
	localctx = NewComparisonOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SPL1ParserRULE_comparisonOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2130706432) != 0) {
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

// IStarContext is an interface to support dynamic dispatch.
type IStarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	STAR_TIGHT() antlr.TerminalNode
	STAR_TERM() antlr.TerminalNode

	// IsStarContext differentiates from other interfaces.
	IsStarContext()
}

type StarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStarContext() *StarContext {
	var p = new(StarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_star
	return p
}

func InitEmptyStarContext(p *StarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_star
}

func (*StarContext) IsStarContext() {}

func NewStarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StarContext {
	var p = new(StarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_star

	return p
}

func (s *StarContext) GetParser() antlr.Parser { return s.parser }

func (s *StarContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR, 0)
}

func (s *StarContext) STAR_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR_TIGHT, 0)
}

func (s *StarContext) STAR_TERM() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR_TERM, 0)
}

func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitStar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Star() (localctx IStarContext) {
	localctx = NewStarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SPL1ParserRULE_star)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&352943232516096) != 0) {
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

// IMinusContext is an interface to support dynamic dispatch.
type IMinusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	MINUS_TIGHT() antlr.TerminalNode

	// IsMinusContext differentiates from other interfaces.
	IsMinusContext()
}

type MinusContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinusContext() *MinusContext {
	var p = new(MinusContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_minus
	return p
}

func InitEmptyMinusContext(p *MinusContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_minus
}

func (*MinusContext) IsMinusContext() {}

func NewMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinusContext {
	var p = new(MinusContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_minus

	return p
}

func (s *MinusContext) GetParser() antlr.Parser { return s.parser }

func (s *MinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserMINUS, 0)
}

func (s *MinusContext) MINUS_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserMINUS_TIGHT, 0)
}

func (s *MinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Minus() (localctx IMinusContext) {
	localctx = NewMinusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SPL1ParserRULE_minus)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL1ParserMINUS || _la == SPL1ParserMINUS_TIGHT) {
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	NUMBER_TIGHT() antlr.TerminalNode

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNUMBER, 0)
}

func (s *NumberContext) NUMBER_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNUMBER_TIGHT, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SPL1ParserRULE_number)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(563)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SPL1ParserNUMBER || _la == SPL1ParserNUMBER_TIGHT) {
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
	Minus() IMinusContext

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
	p.RuleIndex = SPL1ParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPLUS, 0)
}

func (s *AddOpContext) Minus() IMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinusContext)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SPL1ParserRULE_addOp)
	p.SetState(567)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(565)
			p.Match(SPL1ParserPLUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserMINUS, SPL1ParserMINUS_TIGHT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(566)
			p.Minus()
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

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Star() IStarContext
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
	p.RuleIndex = SPL1ParserRULE_mulOp
	return p
}

func InitEmptyMulOpContext(p *MulOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_mulOp
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MulOpContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *MulOpContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSLASH, 0)
}

func (s *MulOpContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPERCENT, 0)
}

func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) MulOp() (localctx IMulOpContext) {
	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SPL1ParserRULE_mulOp)
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(569)
			p.Star()
		}

	case SPL1ParserSLASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(570)
			p.Match(SPL1ParserSLASH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserPERCENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(571)
			p.Match(SPL1ParserPERCENT)
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

// ISearchMulOpContext is an interface to support dynamic dispatch.
type ISearchMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STAR() antlr.TerminalNode
	STAR_TIGHT() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PERCENT() antlr.TerminalNode

	// IsSearchMulOpContext differentiates from other interfaces.
	IsSearchMulOpContext()
}

type SearchMulOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchMulOpContext() *SearchMulOpContext {
	var p = new(SearchMulOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMulOp
	return p
}

func InitEmptySearchMulOpContext(p *SearchMulOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_searchMulOp
}

func (*SearchMulOpContext) IsSearchMulOpContext() {}

func NewSearchMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchMulOpContext {
	var p = new(SearchMulOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_searchMulOp

	return p
}

func (s *SearchMulOpContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchMulOpContext) STAR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR, 0)
}

func (s *SearchMulOpContext) STAR_TIGHT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTAR_TIGHT, 0)
}

func (s *SearchMulOpContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSLASH, 0)
}

func (s *SearchMulOpContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPERCENT, 0)
}

func (s *SearchMulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchMulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchMulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSearchMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SearchMulOp() (localctx ISearchMulOpContext) {
	localctx = NewSearchMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SPL1ParserRULE_searchMulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(574)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&78065325572096) != 0) {
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
	Minus() IMinusContext
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
	p.RuleIndex = SPL1ParserRULE_sign
	return p
}

func InitEmptySignContext(p *SignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sign
}

func (*SignContext) IsSignContext() {}

func NewSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SignContext {
	var p = new(SignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_sign

	return p
}

func (s *SignContext) GetParser() antlr.Parser { return s.parser }

func (s *SignContext) Minus() IMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinusContext)
}

func (s *SignContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPLUS, 0)
}

func (s *SignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Sign() (localctx ISignContext) {
	localctx = NewSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SPL1ParserRULE_sign)
	p.SetState(578)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserMINUS, SPL1ParserMINUS_TIGHT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(576)
			p.Minus()
		}

	case SPL1ParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(577)
			p.Match(SPL1ParserPLUS)
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

// IIdentLikeContext is an interface to support dynamic dispatch.
type IIdentLikeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	SEARCH() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	STATS() antlr.TerminalNode
	TIMECHART() antlr.TerminalNode
	EVAL() antlr.TerminalNode
	REX() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	FIELDS() antlr.TerminalNode
	SORT() antlr.TerminalNode
	HEAD() antlr.TerminalNode
	BIN() antlr.TerminalNode
	BY() antlr.TerminalNode
	AS() antlr.TerminalNode
	SPAN() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	FIELD() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode
	NOT_WORD() antlr.TerminalNode

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
	p.RuleIndex = SPL1ParserRULE_identLike
	return p
}

func InitEmptyIdentLikeContext(p *IdentLikeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_identLike
}

func (*IdentLikeContext) IsIdentLikeContext() {}

func NewIdentLikeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentLikeContext {
	var p = new(IdentLikeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_identLike

	return p
}

func (s *IdentLikeContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentLikeContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserIDENT, 0)
}

func (s *IdentLikeContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSEARCH, 0)
}

func (s *IdentLikeContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserWHERE, 0)
}

func (s *IdentLikeContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTATS, 0)
}

func (s *IdentLikeContext) TIMECHART() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTIMECHART, 0)
}

func (s *IdentLikeContext) EVAL() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEVAL, 0)
}

func (s *IdentLikeContext) REX() antlr.TerminalNode {
	return s.GetToken(SPL1ParserREX, 0)
}

func (s *IdentLikeContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTABLE, 0)
}

func (s *IdentLikeContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELDS, 0)
}

func (s *IdentLikeContext) SORT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSORT, 0)
}

func (s *IdentLikeContext) HEAD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserHEAD, 0)
}

func (s *IdentLikeContext) BIN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBIN, 0)
}

func (s *IdentLikeContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBY, 0)
}

func (s *IdentLikeContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAS, 0)
}

func (s *IdentLikeContext) SPAN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSPAN, 0)
}

func (s *IdentLikeContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLIMIT, 0)
}

func (s *IdentLikeContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELD, 0)
}

func (s *IdentLikeContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *IdentLikeContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *IdentLikeContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT, 0)
}

func (s *IdentLikeContext) NOT_WORD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT_WORD, 0)
}

func (s *IdentLikeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentLikeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentLikeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitIdentLike(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) IdentLike() (localctx IIdentLikeContext) {
	localctx = NewIdentLikeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SPL1ParserRULE_identLike)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562949962858494) != 0) {
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

// IExprWordContext is an interface to support dynamic dispatch.
type IExprWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	SEARCH() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	STATS() antlr.TerminalNode
	TIMECHART() antlr.TerminalNode
	EVAL() antlr.TerminalNode
	REX() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	FIELDS() antlr.TerminalNode
	SORT() antlr.TerminalNode
	HEAD() antlr.TerminalNode
	BIN() antlr.TerminalNode
	BY() antlr.TerminalNode
	AS() antlr.TerminalNode
	SPAN() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	FIELD() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT_WORD() antlr.TerminalNode

	// IsExprWordContext differentiates from other interfaces.
	IsExprWordContext()
}

type ExprWordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprWordContext() *ExprWordContext {
	var p = new(ExprWordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_exprWord
	return p
}

func InitEmptyExprWordContext(p *ExprWordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_exprWord
}

func (*ExprWordContext) IsExprWordContext() {}

func NewExprWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprWordContext {
	var p = new(ExprWordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_exprWord

	return p
}

func (s *ExprWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprWordContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserIDENT, 0)
}

func (s *ExprWordContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSEARCH, 0)
}

func (s *ExprWordContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserWHERE, 0)
}

func (s *ExprWordContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTATS, 0)
}

func (s *ExprWordContext) TIMECHART() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTIMECHART, 0)
}

func (s *ExprWordContext) EVAL() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEVAL, 0)
}

func (s *ExprWordContext) REX() antlr.TerminalNode {
	return s.GetToken(SPL1ParserREX, 0)
}

func (s *ExprWordContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTABLE, 0)
}

func (s *ExprWordContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELDS, 0)
}

func (s *ExprWordContext) SORT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSORT, 0)
}

func (s *ExprWordContext) HEAD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserHEAD, 0)
}

func (s *ExprWordContext) BIN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBIN, 0)
}

func (s *ExprWordContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBY, 0)
}

func (s *ExprWordContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAS, 0)
}

func (s *ExprWordContext) SPAN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSPAN, 0)
}

func (s *ExprWordContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLIMIT, 0)
}

func (s *ExprWordContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELD, 0)
}

func (s *ExprWordContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *ExprWordContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *ExprWordContext) NOT_WORD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT_WORD, 0)
}

func (s *ExprWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprWordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitExprWord(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ExprWord() (localctx IExprWordContext) {
	localctx = NewExprWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SPL1ParserRULE_exprWord)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(582)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562949962334206) != 0) {
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
	p.RuleIndex = SPL1ParserRULE_aggList
	return p
}

func InitEmptyAggListContext(p *AggListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aggList
}

func (*AggListContext) IsAggListContext() {}

func NewAggListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggListContext {
	var p = new(AggListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_aggList

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
	return s.GetTokens(SPL1ParserCOMMA)
}

func (s *AggListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, i)
}

func (s *AggListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAggList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AggList() (localctx IAggListContext) {
	localctx = NewAggListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SPL1ParserRULE_aggList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.Agg()
	}
	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SPL1ParserCOMMA {
		{
			p.SetState(585)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(586)
			p.Agg()
		}

		p.SetState(591)
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
	AliasClause() IAliasClauseContext
	ValueExpr() IValueExprContext

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
	p.RuleIndex = SPL1ParserRULE_agg
	return p
}

func InitEmptyAggContext(p *AggContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_agg
}

func (*AggContext) IsAggContext() {}

func NewAggContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggContext {
	var p = new(AggContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_agg

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
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *AggContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
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

func (s *AggContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *AggContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAgg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Agg() (localctx IAggContext) {
	localctx = NewAggContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SPL1ParserRULE_agg)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(592)
		p.IdentLike()
	}
	p.SetState(598)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL1ParserLPAREN {
		{
			p.SetState(593)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(595)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(594)
				p.ValueExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(597)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL1ParserAS {
		{
			p.SetState(600)
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

// IAggListNotSpanContext is an interface to support dynamic dispatch.
type IAggListNotSpanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AggNotSpan() IAggNotSpanContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllAgg() []IAggContext
	Agg(i int) IAggContext

	// IsAggListNotSpanContext differentiates from other interfaces.
	IsAggListNotSpanContext()
}

type AggListNotSpanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggListNotSpanContext() *AggListNotSpanContext {
	var p = new(AggListNotSpanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aggListNotSpan
	return p
}

func InitEmptyAggListNotSpanContext(p *AggListNotSpanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aggListNotSpan
}

func (*AggListNotSpanContext) IsAggListNotSpanContext() {}

func NewAggListNotSpanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggListNotSpanContext {
	var p = new(AggListNotSpanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_aggListNotSpan

	return p
}

func (s *AggListNotSpanContext) GetParser() antlr.Parser { return s.parser }

func (s *AggListNotSpanContext) AggNotSpan() IAggNotSpanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggNotSpanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggNotSpanContext)
}

func (s *AggListNotSpanContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserCOMMA)
}

func (s *AggListNotSpanContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, i)
}

func (s *AggListNotSpanContext) AllAgg() []IAggContext {
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

func (s *AggListNotSpanContext) Agg(i int) IAggContext {
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

func (s *AggListNotSpanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggListNotSpanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggListNotSpanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAggListNotSpan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AggListNotSpan() (localctx IAggListNotSpanContext) {
	localctx = NewAggListNotSpanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SPL1ParserRULE_aggListNotSpan)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(603)
		p.AggNotSpan()
	}
	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SPL1ParserCOMMA {
		{
			p.SetState(604)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(605)
			p.Agg()
		}

		p.SetState(610)
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

// IAggNotSpanContext is an interface to support dynamic dispatch.
type IAggNotSpanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLikeNotSpan() IIdentLikeNotSpanContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AliasClause() IAliasClauseContext
	ValueExpr() IValueExprContext

	// IsAggNotSpanContext differentiates from other interfaces.
	IsAggNotSpanContext()
}

type AggNotSpanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggNotSpanContext() *AggNotSpanContext {
	var p = new(AggNotSpanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aggNotSpan
	return p
}

func InitEmptyAggNotSpanContext(p *AggNotSpanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aggNotSpan
}

func (*AggNotSpanContext) IsAggNotSpanContext() {}

func NewAggNotSpanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggNotSpanContext {
	var p = new(AggNotSpanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_aggNotSpan

	return p
}

func (s *AggNotSpanContext) GetParser() antlr.Parser { return s.parser }

func (s *AggNotSpanContext) IdentLikeNotSpan() IIdentLikeNotSpanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeNotSpanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeNotSpanContext)
}

func (s *AggNotSpanContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLPAREN, 0)
}

func (s *AggNotSpanContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserRPAREN, 0)
}

func (s *AggNotSpanContext) AliasClause() IAliasClauseContext {
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

func (s *AggNotSpanContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *AggNotSpanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggNotSpanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggNotSpanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAggNotSpan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AggNotSpan() (localctx IAggNotSpanContext) {
	localctx = NewAggNotSpanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SPL1ParserRULE_aggNotSpan)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(611)
		p.IdentLikeNotSpan()
	}
	p.SetState(617)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL1ParserLPAREN {
		{
			p.SetState(612)
			p.Match(SPL1ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(613)
				p.ValueExpr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(616)
			p.Match(SPL1ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(620)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SPL1ParserAS {
		{
			p.SetState(619)
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
	p.RuleIndex = SPL1ParserRULE_aliasClause
	return p
}

func InitEmptyAliasClauseContext(p *AliasClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_aliasClause
}

func (*AliasClauseContext) IsAliasClauseContext() {}

func NewAliasClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasClauseContext {
	var p = new(AliasClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_aliasClause

	return p
}

func (s *AliasClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasClauseContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAS, 0)
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
	case SPL1Visitor:
		return t.VisitAliasClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) AliasClause() (localctx IAliasClauseContext) {
	localctx = NewAliasClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SPL1ParserRULE_aliasClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(622)
		p.Match(SPL1ParserAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(623)
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
	FieldExprList() IFieldExprListContext

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
	p.RuleIndex = SPL1ParserRULE_byClause
	return p
}

func InitEmptyByClauseContext(p *ByClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_byClause
}

func (*ByClauseContext) IsByClauseContext() {}

func NewByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ByClauseContext {
	var p = new(ByClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_byClause

	return p
}

func (s *ByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBY, 0)
}

func (s *ByClauseContext) FieldExprList() IFieldExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldExprListContext)
}

func (s *ByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) ByClause() (localctx IByClauseContext) {
	localctx = NewByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SPL1ParserRULE_byClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(625)
		p.Match(SPL1ParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(626)
		p.FieldExprList()
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

// ITimechartBodyContext is an interface to support dynamic dispatch.
type ITimechartBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTimechartBodyContext differentiates from other interfaces.
	IsTimechartBodyContext()
}

type TimechartBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimechartBodyContext() *TimechartBodyContext {
	var p = new(TimechartBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_timechartBody
	return p
}

func InitEmptyTimechartBodyContext(p *TimechartBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_timechartBody
}

func (*TimechartBodyContext) IsTimechartBodyContext() {}

func NewTimechartBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimechartBodyContext {
	var p = new(TimechartBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_timechartBody

	return p
}

func (s *TimechartBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TimechartBodyContext) CopyAll(ctx *TimechartBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TimechartBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimechartBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TcAggFirstContext struct {
	TimechartBodyContext
}

func NewTcAggFirstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TcAggFirstContext {
	var p = new(TcAggFirstContext)

	InitEmptyTimechartBodyContext(&p.TimechartBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*TimechartBodyContext))

	return p
}

func (s *TcAggFirstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TcAggFirstContext) AggListNotSpan() IAggListNotSpanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggListNotSpanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggListNotSpanContext)
}

func (s *TcAggFirstContext) ByClause() IByClauseContext {
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

func (s *TcAggFirstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitTcAggFirst(s)

	default:
		return t.VisitChildren(s)
	}
}

type TcSpanFirstContext struct {
	TimechartBodyContext
}

func NewTcSpanFirstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TcSpanFirstContext {
	var p = new(TcSpanFirstContext)

	InitEmptyTimechartBodyContext(&p.TimechartBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*TimechartBodyContext))

	return p
}

func (s *TcSpanFirstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TcSpanFirstContext) TimechartSpan() ITimechartSpanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimechartSpanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimechartSpanContext)
}

func (s *TcSpanFirstContext) AggList() IAggListContext {
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

func (s *TcSpanFirstContext) ByClause() IByClauseContext {
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

func (s *TcSpanFirstContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitTcSpanFirst(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) TimechartBody() (localctx ITimechartBodyContext) {
	localctx = NewTimechartBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SPL1ParserRULE_timechartBody)
	var _la int

	p.SetState(638)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTcSpanFirstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(628)

		if !(p.at(SPL1ParserSPAN)) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserSPAN)", ""))
			goto errorExit
		}
		{
			p.SetState(629)
			p.TimechartSpan()
		}
		{
			p.SetState(630)
			p.AggList()
		}
		p.SetState(632)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SPL1ParserBY {
			{
				p.SetState(631)
				p.ByClause()
			}

		}

	case 2:
		localctx = NewTcAggFirstContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.AggListNotSpan()
		}
		p.SetState(636)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SPL1ParserBY {
			{
				p.SetState(635)
				p.ByClause()
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

// ITimechartSpanContext is an interface to support dynamic dispatch.
type ITimechartSpanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SPAN() antlr.TerminalNode
	SpanAssign() ISpanAssignContext

	// IsTimechartSpanContext differentiates from other interfaces.
	IsTimechartSpanContext()
}

type TimechartSpanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimechartSpanContext() *TimechartSpanContext {
	var p = new(TimechartSpanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_timechartSpan
	return p
}

func InitEmptyTimechartSpanContext(p *TimechartSpanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_timechartSpan
}

func (*TimechartSpanContext) IsTimechartSpanContext() {}

func NewTimechartSpanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimechartSpanContext {
	var p = new(TimechartSpanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_timechartSpan

	return p
}

func (s *TimechartSpanContext) GetParser() antlr.Parser { return s.parser }

func (s *TimechartSpanContext) SPAN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSPAN, 0)
}

func (s *TimechartSpanContext) SpanAssign() ISpanAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpanAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpanAssignContext)
}

func (s *TimechartSpanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimechartSpanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimechartSpanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitTimechartSpan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) TimechartSpan() (localctx ITimechartSpanContext) {
	localctx = NewTimechartSpanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SPL1ParserRULE_timechartSpan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(640)
		p.Match(SPL1ParserSPAN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(641)
		p.SpanAssign()
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

// ISpanAssignContext is an interface to support dynamic dispatch.
type ISpanAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	SpanValue() ISpanValueContext

	// IsSpanAssignContext differentiates from other interfaces.
	IsSpanAssignContext()
}

type SpanAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpanAssignContext() *SpanAssignContext {
	var p = new(SpanAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_spanAssign
	return p
}

func InitEmptySpanAssignContext(p *SpanAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_spanAssign
}

func (*SpanAssignContext) IsSpanAssignContext() {}

func NewSpanAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpanAssignContext {
	var p = new(SpanAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_spanAssign

	return p
}

func (s *SpanAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *SpanAssignContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *SpanAssignContext) SpanValue() ISpanValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpanValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpanValueContext)
}

func (s *SpanAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpanAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSpanAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SpanAssign() (localctx ISpanAssignContext) {
	localctx = NewSpanAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SPL1ParserRULE_spanAssign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(643)
		p.Match(SPL1ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(644)
		p.SpanValue()
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

// IRexFieldContext is an interface to support dynamic dispatch.
type IRexFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FIELD() antlr.TerminalNode
	EQ() antlr.TerminalNode
	IdentLike() IIdentLikeContext

	// IsRexFieldContext differentiates from other interfaces.
	IsRexFieldContext()
}

type RexFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRexFieldContext() *RexFieldContext {
	var p = new(RexFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_rexField
	return p
}

func InitEmptyRexFieldContext(p *RexFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_rexField
}

func (*RexFieldContext) IsRexFieldContext() {}

func NewRexFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RexFieldContext {
	var p = new(RexFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_rexField

	return p
}

func (s *RexFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *RexFieldContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELD, 0)
}

func (s *RexFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *RexFieldContext) IdentLike() IIdentLikeContext {
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

func (s *RexFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RexFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RexFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitRexField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) RexField() (localctx IRexFieldContext) {
	localctx = NewRexFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SPL1ParserRULE_rexField)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(646)

	if !(p.at(SPL1ParserFIELD)) {
		p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserFIELD)", ""))
		goto errorExit
	}
	{
		p.SetState(647)
		p.Match(SPL1ParserFIELD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(648)
		p.Match(SPL1ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(649)
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

// IFieldExprListContext is an interface to support dynamic dispatch.
type IFieldExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldExpr() []IFieldExprContext
	FieldExpr(i int) IFieldExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldExprListContext differentiates from other interfaces.
	IsFieldExprListContext()
}

type FieldExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExprListContext() *FieldExprListContext {
	var p = new(FieldExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldExprList
	return p
}

func InitEmptyFieldExprListContext(p *FieldExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldExprList
}

func (*FieldExprListContext) IsFieldExprListContext() {}

func NewFieldExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExprListContext {
	var p = new(FieldExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_fieldExprList

	return p
}

func (s *FieldExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExprListContext) AllFieldExpr() []IFieldExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldExprContext); ok {
			len++
		}
	}

	tst := make([]IFieldExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldExprContext); ok {
			tst[i] = t.(IFieldExprContext)
			i++
		}
	}

	return tst
}

func (s *FieldExprListContext) FieldExpr(i int) IFieldExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExprContext); ok {
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

	return t.(IFieldExprContext)
}

func (s *FieldExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SPL1ParserCOMMA)
}

func (s *FieldExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, i)
}

func (s *FieldExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitFieldExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) FieldExprList() (localctx IFieldExprListContext) {
	localctx = NewFieldExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SPL1ParserRULE_fieldExprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(651)
		p.FieldExpr()
	}
	p.SetState(657)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&915901785309182) != 0 {
		p.SetState(655)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SPL1ParserCOMMA:
			{
				p.SetState(652)
				p.Match(SPL1ParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(653)
				p.FieldExpr()
			}

		case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM, SPL1ParserNOT_WORD:
			{
				p.SetState(654)
				p.FieldExpr()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(659)
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

// IFieldExprContext is an interface to support dynamic dispatch.
type IFieldExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext
	Star() IStarContext

	// IsFieldExprContext differentiates from other interfaces.
	IsFieldExprContext()
}

type FieldExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExprContext() *FieldExprContext {
	var p = new(FieldExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldExpr
	return p
}

func InitEmptyFieldExprContext(p *FieldExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldExpr
}

func (*FieldExprContext) IsFieldExprContext() {}

func NewFieldExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExprContext {
	var p = new(FieldExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_fieldExpr

	return p
}

func (s *FieldExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExprContext) IdentLike() IIdentLikeContext {
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

func (s *FieldExprContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitFieldExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) FieldExpr() (localctx IFieldExprContext) {
	localctx = NewFieldExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SPL1ParserRULE_fieldExpr)
	p.SetState(662)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(660)
			p.IdentLike()
		}

	case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(661)
			p.Star()
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

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLike() IIdentLikeContext
	EQ() antlr.TerminalNode
	ValueExpr() IValueExprContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) IdentLike() IIdentLikeContext {
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

func (s *AssignmentContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *AssignmentContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SPL1ParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(664)
		p.IdentLike()
	}
	{
		p.SetState(665)
		p.Match(SPL1ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(666)
		p.ValueExpr()
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

// ISpanValueContext is an interface to support dynamic dispatch.
type ISpanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DURATION() antlr.TerminalNode
	IdentLike() IIdentLikeContext

	// IsSpanValueContext differentiates from other interfaces.
	IsSpanValueContext()
}

type SpanValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpanValueContext() *SpanValueContext {
	var p = new(SpanValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_spanValue
	return p
}

func InitEmptySpanValueContext(p *SpanValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_spanValue
}

func (*SpanValueContext) IsSpanValueContext() {}

func NewSpanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpanValueContext {
	var p = new(SpanValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_spanValue

	return p
}

func (s *SpanValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SpanValueContext) DURATION() antlr.TerminalNode {
	return s.GetToken(SPL1ParserDURATION, 0)
}

func (s *SpanValueContext) IdentLike() IIdentLikeContext {
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

func (s *SpanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpanValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSpanValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SpanValue() (localctx ISpanValueContext) {
	localctx = NewSpanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SPL1ParserRULE_spanValue)
	p.SetState(670)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserDURATION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(668)
			p.Match(SPL1ParserDURATION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(669)
			p.IdentLike()
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

// ITableItemContext is an interface to support dynamic dispatch.
type ITableItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	Star() IStarContext
	IdentLike() IIdentLikeContext

	// IsTableItemContext differentiates from other interfaces.
	IsTableItemContext()
}

type TableItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableItemContext() *TableItemContext {
	var p = new(TableItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_tableItem
	return p
}

func InitEmptyTableItemContext(p *TableItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_tableItem
}

func (*TableItemContext) IsTableItemContext() {}

func NewTableItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableItemContext {
	var p = new(TableItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_tableItem

	return p
}

func (s *TableItemContext) GetParser() antlr.Parser { return s.parser }

func (s *TableItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, 0)
}

func (s *TableItemContext) Star() IStarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStarContext)
}

func (s *TableItemContext) IdentLike() IIdentLikeContext {
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

func (s *TableItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitTableItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) TableItem() (localctx ITableItemContext) {
	localctx = NewTableItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SPL1ParserRULE_tableItem)
	p.SetState(675)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(672)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSTAR, SPL1ParserSTAR_TIGHT, SPL1ParserSTAR_TERM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(673)
			p.Star()
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(674)
			p.IdentLike()
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

// IFieldsSignContext is an interface to support dynamic dispatch.
type IFieldsSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Minus() IMinusContext
	PLUS() antlr.TerminalNode

	// IsFieldsSignContext differentiates from other interfaces.
	IsFieldsSignContext()
}

type FieldsSignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsSignContext() *FieldsSignContext {
	var p = new(FieldsSignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldsSign
	return p
}

func InitEmptyFieldsSignContext(p *FieldsSignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldsSign
}

func (*FieldsSignContext) IsFieldsSignContext() {}

func NewFieldsSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsSignContext {
	var p = new(FieldsSignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_fieldsSign

	return p
}

func (s *FieldsSignContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsSignContext) Minus() IMinusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMinusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMinusContext)
}

func (s *FieldsSignContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserPLUS, 0)
}

func (s *FieldsSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitFieldsSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) FieldsSign() (localctx IFieldsSignContext) {
	localctx = NewFieldsSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, SPL1ParserRULE_fieldsSign)
	p.SetState(679)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserMINUS, SPL1ParserMINUS_TIGHT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(677)
			p.Minus()
		}

	case SPL1ParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(678)
			p.Match(SPL1ParserPLUS)
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

// IFieldsItemContext is an interface to support dynamic dispatch.
type IFieldsItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	IdentLike() IIdentLikeContext

	// IsFieldsItemContext differentiates from other interfaces.
	IsFieldsItemContext()
}

type FieldsItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsItemContext() *FieldsItemContext {
	var p = new(FieldsItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldsItem
	return p
}

func InitEmptyFieldsItemContext(p *FieldsItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_fieldsItem
}

func (*FieldsItemContext) IsFieldsItemContext() {}

func NewFieldsItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsItemContext {
	var p = new(FieldsItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_fieldsItem

	return p
}

func (s *FieldsItemContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, 0)
}

func (s *FieldsItemContext) IdentLike() IIdentLikeContext {
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

func (s *FieldsItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitFieldsItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) FieldsItem() (localctx IFieldsItemContext) {
	localctx = NewFieldsItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, SPL1ParserRULE_fieldsItem)
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(681)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(682)
			p.IdentLike()
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

// ISortBodyContext is an interface to support dynamic dispatch.
type ISortBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSortBodyContext differentiates from other interfaces.
	IsSortBodyContext()
}

type SortBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortBodyContext() *SortBodyContext {
	var p = new(SortBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortBody
	return p
}

func InitEmptySortBodyContext(p *SortBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortBody
}

func (*SortBodyContext) IsSortBodyContext() {}

func NewSortBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortBodyContext {
	var p = new(SortBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_sortBody

	return p
}

func (s *SortBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SortBodyContext) CopyAll(ctx *SortBodyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SortBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SortFirstFieldContext struct {
	SortBodyContext
}

func NewSortFirstFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortFirstFieldContext {
	var p = new(SortFirstFieldContext)

	InitEmptySortBodyContext(&p.SortBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortBodyContext))

	return p
}

func (s *SortFirstFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortFirstFieldContext) SortFirstItem() ISortFirstItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortFirstItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortFirstItemContext)
}

func (s *SortFirstFieldContext) AllSortItem() []ISortItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortItemContext); ok {
			len++
		}
	}

	tst := make([]ISortItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortItemContext); ok {
			tst[i] = t.(ISortItemContext)
			i++
		}
	}

	return tst
}

func (s *SortFirstFieldContext) SortItem(i int) ISortItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortItemContext); ok {
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

	return t.(ISortItemContext)
}

func (s *SortFirstFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortFirstField(s)

	default:
		return t.VisitChildren(s)
	}
}

type SortEmptyContext struct {
	SortBodyContext
}

func NewSortEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortEmptyContext {
	var p = new(SortEmptyContext)

	InitEmptySortBodyContext(&p.SortBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortBodyContext))

	return p
}

func (s *SortEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type SortLimitContext struct {
	SortBodyContext
}

func NewSortLimitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SortLimitContext {
	var p = new(SortLimitContext)

	InitEmptySortBodyContext(&p.SortBodyContext)
	p.parser = parser
	p.CopyAll(ctx.(*SortBodyContext))

	return p
}

func (s *SortLimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortLimitContext) LimitArg() ILimitArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitArgContext)
}

func (s *SortLimitContext) AllSortItem() []ISortItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortItemContext); ok {
			len++
		}
	}

	tst := make([]ISortItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortItemContext); ok {
			tst[i] = t.(ISortItemContext)
			i++
		}
	}

	return tst
}

func (s *SortLimitContext) SortItem(i int) ISortItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortItemContext); ok {
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

	return t.(ISortItemContext)
}

func (s *SortLimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SortBody() (localctx ISortBodyContext) {
	localctx = NewSortBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, SPL1ParserRULE_sortBody)
	var _la int

	p.SetState(700)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 88, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSortLimitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(685)
			p.LimitArg()
		}
		p.SetState(689)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&704520674869246) != 0 {
			{
				p.SetState(686)
				p.SortItem()
			}

			p.SetState(691)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		localctx = NewSortFirstFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(692)
			p.SortFirstItem()
		}
		p.SetState(696)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&704520674869246) != 0 {
			{
				p.SetState(693)
				p.SortItem()
			}

			p.SetState(698)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		localctx = NewSortEmptyContext(p, localctx)
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

// ISortItemContext is an interface to support dynamic dispatch.
type ISortItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	IdentLike() IIdentLikeContext
	Sign() ISignContext

	// IsSortItemContext differentiates from other interfaces.
	IsSortItemContext()
}

type SortItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortItemContext() *SortItemContext {
	var p = new(SortItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortItem
	return p
}

func InitEmptySortItemContext(p *SortItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortItem
}

func (*SortItemContext) IsSortItemContext() {}

func NewSortItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortItemContext {
	var p = new(SortItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_sortItem

	return p
}

func (s *SortItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SortItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, 0)
}

func (s *SortItemContext) IdentLike() IIdentLikeContext {
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

func (s *SortItemContext) Sign() ISignContext {
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

func (s *SortItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SortItem() (localctx ISortItemContext) {
	localctx = NewSortItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, SPL1ParserRULE_sortItem)
	var _la int

	p.SetState(707)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(702)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserLIMIT, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserPLUS, SPL1ParserMINUS, SPL1ParserMINUS_TIGHT, SPL1ParserNOT_WORD:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(704)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141562122076160) != 0 {
			{
				p.SetState(703)
				p.Sign()
			}

		}
		{
			p.SetState(706)
			p.IdentLike()
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

// ISortFirstItemContext is an interface to support dynamic dispatch.
type ISortFirstItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMA() antlr.TerminalNode
	SortField() ISortFieldContext
	Sign() ISignContext

	// IsSortFirstItemContext differentiates from other interfaces.
	IsSortFirstItemContext()
}

type SortFirstItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortFirstItemContext() *SortFirstItemContext {
	var p = new(SortFirstItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortFirstItem
	return p
}

func InitEmptySortFirstItemContext(p *SortFirstItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortFirstItem
}

func (*SortFirstItemContext) IsSortFirstItemContext() {}

func NewSortFirstItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortFirstItemContext {
	var p = new(SortFirstItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_sortFirstItem

	return p
}

func (s *SortFirstItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SortFirstItemContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SPL1ParserCOMMA, 0)
}

func (s *SortFirstItemContext) SortField() ISortFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortFieldContext)
}

func (s *SortFirstItemContext) Sign() ISignContext {
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

func (s *SortFirstItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortFirstItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortFirstItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortFirstItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SortFirstItem() (localctx ISortFirstItemContext) {
	localctx = NewSortFirstItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, SPL1ParserRULE_sortFirstItem)
	var _la int

	p.SetState(714)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SPL1ParserCOMMA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(709)
			p.Match(SPL1ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SPL1ParserSEARCH, SPL1ParserWHERE, SPL1ParserSTATS, SPL1ParserTIMECHART, SPL1ParserEVAL, SPL1ParserREX, SPL1ParserTABLE, SPL1ParserFIELDS, SPL1ParserSORT, SPL1ParserHEAD, SPL1ParserBIN, SPL1ParserBY, SPL1ParserAS, SPL1ParserSPAN, SPL1ParserFIELD, SPL1ParserAND, SPL1ParserOR, SPL1ParserNOT, SPL1ParserIDENT, SPL1ParserPLUS, SPL1ParserMINUS, SPL1ParserMINUS_TIGHT:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(711)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&141562122076160) != 0 {
			{
				p.SetState(710)
				p.Sign()
			}

		}
		{
			p.SetState(713)
			p.SortField()
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

// ISortFieldContext is an interface to support dynamic dispatch.
type ISortFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLikeNotLimit() IIdentLikeNotLimitContext

	// IsSortFieldContext differentiates from other interfaces.
	IsSortFieldContext()
}

type SortFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortFieldContext() *SortFieldContext {
	var p = new(SortFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortField
	return p
}

func InitEmptySortFieldContext(p *SortFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_sortField
}

func (*SortFieldContext) IsSortFieldContext() {}

func NewSortFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortFieldContext {
	var p = new(SortFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_sortField

	return p
}

func (s *SortFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SortFieldContext) IdentLikeNotLimit() IIdentLikeNotLimitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeNotLimitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeNotLimitContext)
}

func (s *SortFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitSortField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) SortField() (localctx ISortFieldContext) {
	localctx = NewSortFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, SPL1ParserRULE_sortField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(716)
		p.IdentLikeNotLimit()
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

// ILimitArgContext is an interface to support dynamic dispatch.
type ILimitArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLimitArgContext differentiates from other interfaces.
	IsLimitArgContext()
}

type LimitArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitArgContext() *LimitArgContext {
	var p = new(LimitArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_limitArg
	return p
}

func InitEmptyLimitArgContext(p *LimitArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_limitArg
}

func (*LimitArgContext) IsLimitArgContext() {}

func NewLimitArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitArgContext {
	var p = new(LimitArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_limitArg

	return p
}

func (s *LimitArgContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitArgContext) CopyAll(ctx *LimitArgContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LimitArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LimitNumberContext struct {
	LimitArgContext
}

func NewLimitNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LimitNumberContext {
	var p = new(LimitNumberContext)

	InitEmptyLimitArgContext(&p.LimitArgContext)
	p.parser = parser
	p.CopyAll(ctx.(*LimitArgContext))

	return p
}

func (s *LimitNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitNumberContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LimitNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitLimitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type LimitKeywordContext struct {
	LimitArgContext
}

func NewLimitKeywordContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LimitKeywordContext {
	var p = new(LimitKeywordContext)

	InitEmptyLimitArgContext(&p.LimitArgContext)
	p.parser = parser
	p.CopyAll(ctx.(*LimitArgContext))

	return p
}

func (s *LimitKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitKeywordContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLIMIT, 0)
}

func (s *LimitKeywordContext) LimitTail() ILimitTailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitTailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitTailContext)
}

func (s *LimitKeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitLimitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) LimitArg() (localctx ILimitArgContext) {
	localctx = NewLimitArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, SPL1ParserRULE_limitArg)
	p.SetState(722)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 93, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLimitKeywordContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(718)

		if !(p.at(SPL1ParserLIMIT)) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserLIMIT)", ""))
			goto errorExit
		}
		{
			p.SetState(719)
			p.Match(SPL1ParserLIMIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(720)
			p.LimitTail()
		}

	case 2:
		localctx = NewLimitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(721)
			p.Number()
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

// ILimitTailContext is an interface to support dynamic dispatch.
type ILimitTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	Number() INumberContext

	// IsLimitTailContext differentiates from other interfaces.
	IsLimitTailContext()
}

type LimitTailContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitTailContext() *LimitTailContext {
	var p = new(LimitTailContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_limitTail
	return p
}

func InitEmptyLimitTailContext(p *LimitTailContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_limitTail
}

func (*LimitTailContext) IsLimitTailContext() {}

func NewLimitTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitTailContext {
	var p = new(LimitTailContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_limitTail

	return p
}

func (s *LimitTailContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitTailContext) EQ() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEQ, 0)
}

func (s *LimitTailContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *LimitTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitTailContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitLimitTail(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) LimitTail() (localctx ILimitTailContext) {
	localctx = NewLimitTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, SPL1ParserRULE_limitTail)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(724)
		p.Match(SPL1ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(725)
		p.Number()
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

// IBinItemContext is an interface to support dynamic dispatch.
type IBinItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBinItemContext differentiates from other interfaces.
	IsBinItemContext()
}

type BinItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinItemContext() *BinItemContext {
	var p = new(BinItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_binItem
	return p
}

func InitEmptyBinItemContext(p *BinItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_binItem
}

func (*BinItemContext) IsBinItemContext() {}

func NewBinItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinItemContext {
	var p = new(BinItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_binItem

	return p
}

func (s *BinItemContext) GetParser() antlr.Parser { return s.parser }

func (s *BinItemContext) CopyAll(ctx *BinItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BinItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinFieldItemContext struct {
	BinItemContext
}

func NewBinFieldItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinFieldItemContext {
	var p = new(BinFieldItemContext)

	InitEmptyBinItemContext(&p.BinItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*BinItemContext))

	return p
}

func (s *BinFieldItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinFieldItemContext) BinField() IBinFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinFieldContext)
}

func (s *BinFieldItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitBinFieldItem(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinSpanItemContext struct {
	BinItemContext
}

func NewBinSpanItemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinSpanItemContext {
	var p = new(BinSpanItemContext)

	InitEmptyBinItemContext(&p.BinItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*BinItemContext))

	return p
}

func (s *BinSpanItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinSpanItemContext) SPAN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSPAN, 0)
}

func (s *BinSpanItemContext) SpanAssign() ISpanAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpanAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpanAssignContext)
}

func (s *BinSpanItemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitBinSpanItem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) BinItem() (localctx IBinItemContext) {
	localctx = NewBinItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, SPL1ParserRULE_binItem)
	p.SetState(731)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 94, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBinSpanItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(727)

		if !(p.at(SPL1ParserSPAN)) {
			p.SetError(antlr.NewFailedPredicateException(p, "p.at(SPL1ParserSPAN)", ""))
			goto errorExit
		}
		{
			p.SetState(728)
			p.Match(SPL1ParserSPAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(729)
			p.SpanAssign()
		}

	case 2:
		localctx = NewBinFieldItemContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(730)
			p.BinField()
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

// IBinFieldContext is an interface to support dynamic dispatch.
type IBinFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentLikeNotSpan() IIdentLikeNotSpanContext

	// IsBinFieldContext differentiates from other interfaces.
	IsBinFieldContext()
}

type BinFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinFieldContext() *BinFieldContext {
	var p = new(BinFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_binField
	return p
}

func InitEmptyBinFieldContext(p *BinFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_binField
}

func (*BinFieldContext) IsBinFieldContext() {}

func NewBinFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinFieldContext {
	var p = new(BinFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_binField

	return p
}

func (s *BinFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *BinFieldContext) IdentLikeNotSpan() IIdentLikeNotSpanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentLikeNotSpanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentLikeNotSpanContext)
}

func (s *BinFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitBinField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) BinField() (localctx IBinFieldContext) {
	localctx = NewBinFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, SPL1ParserRULE_binField)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(733)
		p.IdentLikeNotSpan()
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

// IIdentLikeNotSpanContext is an interface to support dynamic dispatch.
type IIdentLikeNotSpanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	SEARCH() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	STATS() antlr.TerminalNode
	TIMECHART() antlr.TerminalNode
	EVAL() antlr.TerminalNode
	REX() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	FIELDS() antlr.TerminalNode
	SORT() antlr.TerminalNode
	HEAD() antlr.TerminalNode
	BIN() antlr.TerminalNode
	BY() antlr.TerminalNode
	AS() antlr.TerminalNode
	LIMIT() antlr.TerminalNode
	FIELD() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsIdentLikeNotSpanContext differentiates from other interfaces.
	IsIdentLikeNotSpanContext()
}

type IdentLikeNotSpanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentLikeNotSpanContext() *IdentLikeNotSpanContext {
	var p = new(IdentLikeNotSpanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_identLikeNotSpan
	return p
}

func InitEmptyIdentLikeNotSpanContext(p *IdentLikeNotSpanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_identLikeNotSpan
}

func (*IdentLikeNotSpanContext) IsIdentLikeNotSpanContext() {}

func NewIdentLikeNotSpanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentLikeNotSpanContext {
	var p = new(IdentLikeNotSpanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_identLikeNotSpan

	return p
}

func (s *IdentLikeNotSpanContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentLikeNotSpanContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserIDENT, 0)
}

func (s *IdentLikeNotSpanContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSEARCH, 0)
}

func (s *IdentLikeNotSpanContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserWHERE, 0)
}

func (s *IdentLikeNotSpanContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTATS, 0)
}

func (s *IdentLikeNotSpanContext) TIMECHART() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTIMECHART, 0)
}

func (s *IdentLikeNotSpanContext) EVAL() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEVAL, 0)
}

func (s *IdentLikeNotSpanContext) REX() antlr.TerminalNode {
	return s.GetToken(SPL1ParserREX, 0)
}

func (s *IdentLikeNotSpanContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTABLE, 0)
}

func (s *IdentLikeNotSpanContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELDS, 0)
}

func (s *IdentLikeNotSpanContext) SORT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSORT, 0)
}

func (s *IdentLikeNotSpanContext) HEAD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserHEAD, 0)
}

func (s *IdentLikeNotSpanContext) BIN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBIN, 0)
}

func (s *IdentLikeNotSpanContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBY, 0)
}

func (s *IdentLikeNotSpanContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAS, 0)
}

func (s *IdentLikeNotSpanContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserLIMIT, 0)
}

func (s *IdentLikeNotSpanContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELD, 0)
}

func (s *IdentLikeNotSpanContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *IdentLikeNotSpanContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *IdentLikeNotSpanContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT, 0)
}

func (s *IdentLikeNotSpanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentLikeNotSpanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentLikeNotSpanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitIdentLikeNotSpan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) IdentLikeNotSpan() (localctx IIdentLikeNotSpanContext) {
	localctx = NewIdentLikeNotSpanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, SPL1ParserRULE_identLikeNotSpan)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(735)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9420798) != 0) {
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

// IIdentLikeNotLimitContext is an interface to support dynamic dispatch.
type IIdentLikeNotLimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENT() antlr.TerminalNode
	SEARCH() antlr.TerminalNode
	WHERE() antlr.TerminalNode
	STATS() antlr.TerminalNode
	TIMECHART() antlr.TerminalNode
	EVAL() antlr.TerminalNode
	REX() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	FIELDS() antlr.TerminalNode
	SORT() antlr.TerminalNode
	HEAD() antlr.TerminalNode
	BIN() antlr.TerminalNode
	BY() antlr.TerminalNode
	AS() antlr.TerminalNode
	SPAN() antlr.TerminalNode
	FIELD() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsIdentLikeNotLimitContext differentiates from other interfaces.
	IsIdentLikeNotLimitContext()
}

type IdentLikeNotLimitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentLikeNotLimitContext() *IdentLikeNotLimitContext {
	var p = new(IdentLikeNotLimitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_identLikeNotLimit
	return p
}

func InitEmptyIdentLikeNotLimitContext(p *IdentLikeNotLimitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SPL1ParserRULE_identLikeNotLimit
}

func (*IdentLikeNotLimitContext) IsIdentLikeNotLimitContext() {}

func NewIdentLikeNotLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentLikeNotLimitContext {
	var p = new(IdentLikeNotLimitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SPL1ParserRULE_identLikeNotLimit

	return p
}

func (s *IdentLikeNotLimitContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentLikeNotLimitContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserIDENT, 0)
}

func (s *IdentLikeNotLimitContext) SEARCH() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSEARCH, 0)
}

func (s *IdentLikeNotLimitContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserWHERE, 0)
}

func (s *IdentLikeNotLimitContext) STATS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSTATS, 0)
}

func (s *IdentLikeNotLimitContext) TIMECHART() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTIMECHART, 0)
}

func (s *IdentLikeNotLimitContext) EVAL() antlr.TerminalNode {
	return s.GetToken(SPL1ParserEVAL, 0)
}

func (s *IdentLikeNotLimitContext) REX() antlr.TerminalNode {
	return s.GetToken(SPL1ParserREX, 0)
}

func (s *IdentLikeNotLimitContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SPL1ParserTABLE, 0)
}

func (s *IdentLikeNotLimitContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELDS, 0)
}

func (s *IdentLikeNotLimitContext) SORT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSORT, 0)
}

func (s *IdentLikeNotLimitContext) HEAD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserHEAD, 0)
}

func (s *IdentLikeNotLimitContext) BIN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBIN, 0)
}

func (s *IdentLikeNotLimitContext) BY() antlr.TerminalNode {
	return s.GetToken(SPL1ParserBY, 0)
}

func (s *IdentLikeNotLimitContext) AS() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAS, 0)
}

func (s *IdentLikeNotLimitContext) SPAN() antlr.TerminalNode {
	return s.GetToken(SPL1ParserSPAN, 0)
}

func (s *IdentLikeNotLimitContext) FIELD() antlr.TerminalNode {
	return s.GetToken(SPL1ParserFIELD, 0)
}

func (s *IdentLikeNotLimitContext) AND() antlr.TerminalNode {
	return s.GetToken(SPL1ParserAND, 0)
}

func (s *IdentLikeNotLimitContext) OR() antlr.TerminalNode {
	return s.GetToken(SPL1ParserOR, 0)
}

func (s *IdentLikeNotLimitContext) NOT() antlr.TerminalNode {
	return s.GetToken(SPL1ParserNOT, 0)
}

func (s *IdentLikeNotLimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentLikeNotLimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentLikeNotLimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SPL1Visitor:
		return t.VisitIdentLikeNotLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SPL1Parser) IdentLikeNotLimit() (localctx IIdentLikeNotLimitContext) {
	localctx = NewIdentLikeNotLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, SPL1ParserRULE_identLikeNotLimit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(737)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9404414) != 0) {
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

func (p *SPL1Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *SearchOrTailContext = nil
		if localctx != nil {
			t = localctx.(*SearchOrTailContext)
		}
		return p.SearchOrTail_Sempred(t, predIndex)

	case 7:
		var t *SearchAndPartContext = nil
		if localctx != nil {
			t = localctx.(*SearchAndPartContext)
		}
		return p.SearchAndPart_Sempred(t, predIndex)

	case 10:
		var t *SearchComparisonTailContext = nil
		if localctx != nil {
			t = localctx.(*SearchComparisonTailContext)
		}
		return p.SearchComparisonTail_Sempred(t, predIndex)

	case 13:
		var t *RhsTailContext = nil
		if localctx != nil {
			t = localctx.(*RhsTailContext)
		}
		return p.RhsTail_Sempred(t, predIndex)

	case 18:
		var t *SearchAddTailContext = nil
		if localctx != nil {
			t = localctx.(*SearchAddTailContext)
		}
		return p.SearchAddTail_Sempred(t, predIndex)

	case 20:
		var t *SearchMulTailContext = nil
		if localctx != nil {
			t = localctx.(*SearchMulTailContext)
		}
		return p.SearchMulTail_Sempred(t, predIndex)

	case 22:
		var t *SearchPrimaryContext = nil
		if localctx != nil {
			t = localctx.(*SearchPrimaryContext)
		}
		return p.SearchPrimary_Sempred(t, predIndex)

	case 23:
		var t *SearchTermTailContext = nil
		if localctx != nil {
			t = localctx.(*SearchTermTailContext)
		}
		return p.SearchTermTail_Sempred(t, predIndex)

	case 25:
		var t *ValueOrTailContext = nil
		if localctx != nil {
			t = localctx.(*ValueOrTailContext)
		}
		return p.ValueOrTail_Sempred(t, predIndex)

	case 27:
		var t *ValueAndPartContext = nil
		if localctx != nil {
			t = localctx.(*ValueAndPartContext)
		}
		return p.ValueAndPart_Sempred(t, predIndex)

	case 30:
		var t *ValueComparisonTailContext = nil
		if localctx != nil {
			t = localctx.(*ValueComparisonTailContext)
		}
		return p.ValueComparisonTail_Sempred(t, predIndex)

	case 32:
		var t *ValueAddTailContext = nil
		if localctx != nil {
			t = localctx.(*ValueAddTailContext)
		}
		return p.ValueAddTail_Sempred(t, predIndex)

	case 34:
		var t *ValueMulTailContext = nil
		if localctx != nil {
			t = localctx.(*ValueMulTailContext)
		}
		return p.ValueMulTail_Sempred(t, predIndex)

	case 36:
		var t *ValuePrimaryContext = nil
		if localctx != nil {
			t = localctx.(*ValuePrimaryContext)
		}
		return p.ValuePrimary_Sempred(t, predIndex)

	case 55:
		var t *TimechartBodyContext = nil
		if localctx != nil {
			t = localctx.(*TimechartBodyContext)
		}
		return p.TimechartBody_Sempred(t, predIndex)

	case 58:
		var t *RexFieldContext = nil
		if localctx != nil {
			t = localctx.(*RexFieldContext)
		}
		return p.RexField_Sempred(t, predIndex)

	case 70:
		var t *LimitArgContext = nil
		if localctx != nil {
			t = localctx.(*LimitArgContext)
		}
		return p.LimitArg_Sempred(t, predIndex)

	case 72:
		var t *BinItemContext = nil
		if localctx != nil {
			t = localctx.(*BinItemContext)
		}
		return p.BinItem_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SPL1Parser) SearchOrTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.at(SPL1ParserOR)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchAndPart_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.at(SPL1ParserAND)

	case 2:
		return p.startsTerm()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchComparisonTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.atCompare()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) RhsTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.touchingPrev()

	case 5:
		return p.touchingPrev()

	case 6:
		return p.touchingPrev()

	case 7:
		return p.hyphenTail()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchAddTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.atAdd()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchMulTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.atSearchMul()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchPrimary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.termStarRun()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) SearchTermTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.touchingPrev()

	case 12:
		return p.touchingPrev()

	case 13:
		return p.touchingPrev()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValueOrTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 14:
		return p.at(SPL1ParserOR)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValueAndPart_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 15:
		return p.at(SPL1ParserAND)

	case 16:
		return p.startsTerm()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValueComparisonTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 17:
		return p.atCompare()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValueAddTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 18:
		return p.atAdd()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValueMulTail_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 19:
		return p.atMul()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) ValuePrimary_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 20:
		return p.callAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) TimechartBody_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 21:
		return p.at(SPL1ParserSPAN)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) RexField_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 22:
		return p.at(SPL1ParserFIELD)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) LimitArg_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 23:
		return p.at(SPL1ParserLIMIT)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SPL1Parser) BinItem_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 24:
		return p.at(SPL1ParserSPAN)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
