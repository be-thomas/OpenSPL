package spl1

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/parser/spl1/gen"
)

// tightStream retypes the tokens that begin a comparison-RHS value run.
//
// Classic SPL decides between a literal value run and an expression on the
// right of a comparison by what the value tokens look like: `host=web*` is the
// literal `web*`, while `host=1+1` is an expression. The grammar cannot make
// that choice with a predicate, because ANTLR's SLL prediction never uses a
// semantic predicate to prune a simulation: alternatives that share a first
// token and differ only by predicates are simulated together to the end of the
// entry rule, and when both die there the parser reports no viable alternative
// instead of asking the predicates. That is exactly what happened to
// `host=web* | head`, where the value and expression alternatives both reach
// the stage pipe and then run out of input.
//
// So the driver retypes the STAR, MINUS, and NUMBER tokens that begin a value
// run — the same test as the hand parser's startsSearchValue — and the grammar
// decides on the token type alone.
//
// Only the parser's lookahead is retyped. The token stream, the parse tree,
// and the Lex adapter keep the tokens the lexer produced, so the lowering code
// still reads STAR, MINUS, and NUMBER for operators and the public token API
// is unchanged; the tight spelling exists only for grammar decisions.
type tightStream struct {
	*antlr.CommonTokenStream
	types map[int]int
}

// tightToken is a token with a rewritten type. Every other method is the
// lexer's token, so text, spans, and positions are unchanged.
type tightToken struct {
	antlr.Token
	typ int
}

func (t tightToken) GetTokenType() int { return t.typ }

func (s *tightStream) LT(k int) antlr.Token {
	t := s.CommonTokenStream.LT(k)
	if t == nil {
		return nil
	}
	if typ, ok := s.types[t.GetTokenIndex()]; ok {
		return tightToken{Token: t, typ: typ}
	}
	return t
}

func (s *tightStream) LA(k int) int {
	t := s.LT(k)
	if t == nil {
		return antlr.TokenInvalidType
	}
	return t.GetTokenType()
}

// tightTokenTypes reports the value-run retyping for a filled token stream.
// A token starts a run only when it touches what follows it:
//
//	NUMBER  followed by MINUS or STAR
//	STAR    followed by a word or number, or by MINUS and then a word, number,
//	        or another STAR
//	MINUS   followed by a word or STAR
//
// A duration is a word for neither a star nor a hyphen: `*1h` and `-1h` do not
// start a value run. A hyphen never merges a number, so `host=-1` stays unary
// minus, exactly as startsSearchValue has it.
func tightTokenTypes(toks []antlr.Token) map[int]int {
	types := make(map[int]int)
	for i, t := range toks {
		next := tokenAt(toks, i+1)
		if next == nil || t.GetStop()+1 != next.GetStart() {
			continue
		}
		switch t.GetTokenType() {
		case spl1gen.SPL1ParserNUMBER:
			if isStarOrMinus(next.GetTokenType()) {
				types[t.GetTokenIndex()] = spl1gen.SPL1ParserNUMBER_TIGHT
			}

		case spl1gen.SPL1ParserSTAR:
			if isValueWord(next.GetTokenType()) {
				types[t.GetTokenIndex()] = spl1gen.SPL1ParserSTAR_TIGHT
				continue
			}
			if next.GetTokenType() != spl1gen.SPL1ParserMINUS {
				continue
			}
			after := tokenAt(toks, i+2)
			if after != nil && next.GetStop()+1 == after.GetStart() &&
				(isValueWord(after.GetTokenType()) || after.GetTokenType() == spl1gen.SPL1ParserSTAR) {
				types[t.GetTokenIndex()] = spl1gen.SPL1ParserSTAR_TIGHT
			}

		case spl1gen.SPL1ParserMINUS:
			if isIdentLikeToken(next.GetTokenType()) || next.GetTokenType() == spl1gen.SPL1ParserSTAR {
				types[t.GetTokenIndex()] = spl1gen.SPL1ParserMINUS_TIGHT
			}
		}
	}
	return types
}

// isValueWord reports whether t is a word or number as a value run sees it.
// Keywords are words too; a duration is not.
func isValueWord(t int) bool {
	return isIdentLikeToken(t) || t == spl1gen.SPL1ParserNUMBER
}

func isStarOrMinus(t int) bool {
	return t == spl1gen.SPL1ParserSTAR || t == spl1gen.SPL1ParserMINUS
}

func tokenAt(toks []antlr.Token, i int) antlr.Token {
	if i < len(toks) {
		return toks[i]
	}
	return nil
}
