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

// tightTokenTypes reports the retyping for a filled token stream.
//
// A STAR that touches a preceding word is STAR_TERM: the wildcard is part of
// that word's search term (`error*`, `*error*`, `a**`). This is decided before
// the value-run rules and takes precedence, because such a star touches a word
// on the left even when it also touches one on the right (`a*b`, where the hand
// parser reads one term, not a product).
//
// The remaining tokens start a value run only when they touch what follows:
//
//	NUMBER  followed by MINUS or STAR
//	STAR    followed by a word or number, or by MINUS and then a word, number,
//	        or another STAR
//	MINUS   followed by a word or STAR
//
// A duration is a word for neither a star nor a hyphen: `*1h` and `-1h` do not
// start a value run. A hyphen never merges a number, so `host=-1` stays unary
// minus, exactly as startsSearchValue has it. A number is not a word for
// STAR_TERM either, so `1*2` and `5m*2` stay multiplication, like `(a)*b`.
//
// A NOT that follows an arithmetic or comparison operator is NOT_WORD: there
// it is a field name (`eval x=1+not`), while at the start of an expression it
// is the negation operator and must commit (`not%x` is an error). An `eval`
// assignment's `=` starts an expression, so `eval x=not(2)` stays a negation.
// The grammar's expression primaries accept NOT_WORD and reject a bare NOT, so
// the choice does not depend on prediction.
func tightTokenTypes(toks []antlr.Token) map[int]int {
	types := make(map[int]int)
	assignments := evalAssignmentEquals(toks)
	for i, t := range toks {
		if t.GetTokenType() != spl1gen.SPL1ParserNOT {
			continue
		}
		prev := tokenAt(toks, i-1)
		if prev == nil || !notFollowsOperator(toks, i-1, assignments) {
			continue
		}
		// An assignment's '=' starts an expression, so `eval x=not` is the
		// negation operator; only an operator inside an expression makes a
		// following `not` a field name.
		if prev.GetTokenType() == spl1gen.SPL1ParserEQ && assignments[prev.GetTokenIndex()] {
			continue
		}
		types[t.GetTokenIndex()] = spl1gen.SPL1ParserNOT_WORD
	}

	for i, t := range toks {
		if t.GetTokenType() != spl1gen.SPL1ParserSTAR {
			continue
		}
		prev := tokenAt(toks, i-1)
		if prev == nil || prev.GetStop()+1 != t.GetStart() {
			continue
		}
		prevType := prev.GetTokenType()
		if isIdentLikeToken(prevType) || types[prev.GetTokenIndex()] == spl1gen.SPL1ParserSTAR_TERM {
			types[t.GetTokenIndex()] = spl1gen.SPL1ParserSTAR_TERM
		}
	}

	// A star with an operand on its left can only multiply, never begin a
	// term, and the grammar says so with its own token type. A predicate
	// cannot: SLL prediction simulates the implicit-AND alternative whenever
	// both start with a star, and ignores the predicate. Two stars in a row
	// alternate roles, and the alternation differs between term position and
	// value position, so a star that follows a star keeps its plain type and
	// the grammar's own loops decide (`a * *` is `a * TRUE`).
	for i, t := range toks {
		if t.GetTokenType() != spl1gen.SPL1ParserSTAR {
			continue
		}
		if _, ok := types[t.GetTokenIndex()]; ok {
			continue
		}
		if !starHasLeftOperand(toks, i, assignments) {
			continue
		}
		if prev := tokenAt(toks, i-1); prev != nil &&
			prev.GetTokenType() == spl1gen.SPL1ParserSTAR {
			continue
		}
		types[t.GetTokenIndex()] = spl1gen.SPL1ParserSTAR_OP
	}

	for i, t := range toks {
		if _, ok := types[t.GetTokenIndex()]; ok {
			continue
		}
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

// evalAssignmentEquals reports which EQ tokens are `eval` assignment
// separators rather than comparisons. Only those start an expression, where a
// following `not` is the negation operator.
func evalAssignmentEquals(toks []antlr.Token) map[int]bool {
	eqs := make(map[int]bool)
	// A pipeline may also begin with `eval`, with no pipe before it.
	inEval := false
	expectLHS := false
	if first := tokenAt(toks, 0); first != nil &&
		first.GetTokenType() == spl1gen.SPL1ParserEVAL {
		inEval = true
		expectLHS = true
	}
	depth := 0
	for i, t := range toks {
		switch t.GetTokenType() {
		case spl1gen.SPL1ParserPIPE:
			inEval = false
			expectLHS = false
			depth = 0
			if cmd := tokenAt(toks, i+1); cmd != nil &&
				cmd.GetTokenType() == spl1gen.SPL1ParserEVAL {
				inEval = true
				expectLHS = true
			}
		case spl1gen.SPL1ParserLPAREN:
			depth++
		case spl1gen.SPL1ParserRPAREN:
			if depth > 0 {
				depth--
			}
		case spl1gen.SPL1ParserCOMMA:
			if inEval && depth == 0 {
				expectLHS = true
			}
		case spl1gen.SPL1ParserEQ:
			if inEval && depth == 0 && expectLHS {
				eqs[t.GetTokenIndex()] = true
				expectLHS = false
			}
		}
	}
	return eqs
}

// notFollowsOperator reports whether the token at index j puts a following
// `not` in field position, as opposed to the start of an expression where it is
// the negation keyword.
func notFollowsOperator(toks []antlr.Token, j int, assignments map[int]bool) bool {
	switch t := toks[j].GetTokenType(); t {
	case spl1gen.SPL1ParserSTAR:
		// A star only multiplies when it has a left operand. `where a * not`
		// reads `not` as a field, while `where * not` is a wildcard term
		// followed by the negation keyword.
		return starHasLeftOperand(toks, j, assignments)
	case spl1gen.SPL1ParserPLUS, spl1gen.SPL1ParserMINUS,
		spl1gen.SPL1ParserSLASH, spl1gen.SPL1ParserPERCENT,
		spl1gen.SPL1ParserEQ2, spl1gen.SPL1ParserEQ, spl1gen.SPL1ParserNE,
		spl1gen.SPL1ParserLT, spl1gen.SPL1ParserLE, spl1gen.SPL1ParserGT,
		spl1gen.SPL1ParserGE:
		return true
	}
	return false
}

// starHasLeftOperand reports whether the star at index j multiplies an operand
// to its left rather than starting an expression. A star after a completed
// operand is multiplication; a star after a command keyword that starts the
// stage, or after an operator or `(`, begins a wildcard term.
func starHasLeftOperand(toks []antlr.Token, j int, assignments map[int]bool) bool {
	prev := tokenAt(toks, j-1)
	if prev == nil {
		return false
	}
	t := prev.GetTokenType()
	switch {
	case t == spl1gen.SPL1ParserNOT:
		// `not` is the negation keyword where an expression starts and a
		// field name only after an arithmetic or comparison operator. An
		// assignment's '=' starts an expression, so it does not make the
		// `not` a field: `eval y=not * b` negates the wildcard.
		before := tokenAt(toks, j-2)
		if before == nil || !isArithmeticOrComparison(before.GetTokenType()) {
			return false
		}
		return !(before.GetTokenType() == spl1gen.SPL1ParserEQ &&
			assignments[before.GetTokenIndex()])
	case isIdentLikeToken(t):
		// A command keyword at the start of a stage is the command itself,
		// not a field: `where * not` starts with a wildcard term.
		return !(isCommandKeyword(t) && isStageStart(toks, j-1))
	case t == spl1gen.SPL1ParserSTRING || t == spl1gen.SPL1ParserNUMBER ||
		t == spl1gen.SPL1ParserDURATION || t == spl1gen.SPL1ParserRPAREN ||
		t == spl1gen.SPL1ParserRBRACKET || t == spl1gen.SPL1ParserSTAR:
		return true
	}
	return false
}

// isArithmeticOrComparison reports whether t is an operator whose right operand
// is a primary, so a following `not` is a field name there.
func isArithmeticOrComparison(t int) bool {
	switch t {
	case spl1gen.SPL1ParserPLUS, spl1gen.SPL1ParserMINUS,
		spl1gen.SPL1ParserSLASH, spl1gen.SPL1ParserPERCENT,
		spl1gen.SPL1ParserEQ2, spl1gen.SPL1ParserEQ, spl1gen.SPL1ParserNE,
		spl1gen.SPL1ParserLT, spl1gen.SPL1ParserLE, spl1gen.SPL1ParserGT,
		spl1gen.SPL1ParserGE:
		return true
	}
	return false
}

// isCommandKeyword reports whether t names a command the grammar implements.
func isCommandKeyword(t int) bool {
	switch t {
	case spl1gen.SPL1ParserSEARCH, spl1gen.SPL1ParserWHERE,
		spl1gen.SPL1ParserSTATS, spl1gen.SPL1ParserTIMECHART,
		spl1gen.SPL1ParserEVAL, spl1gen.SPL1ParserREX,
		spl1gen.SPL1ParserTABLE, spl1gen.SPL1ParserFIELDS,
		spl1gen.SPL1ParserSORT, spl1gen.SPL1ParserHEAD,
		spl1gen.SPL1ParserBIN, spl1gen.SPL1ParserRENAME:
		return true
	}
	return false
}

// isStageStart reports whether index j begins a stage: the first token of the
// query, or the first token after a pipe.
func isStageStart(toks []antlr.Token, j int) bool {
	if j == 0 {
		return true
	}
	prev := tokenAt(toks, j-1)
	return prev != nil && prev.GetTokenType() == spl1gen.SPL1ParserPIPE
}

func isStarOrMinus(t int) bool {
	return t == spl1gen.SPL1ParserSTAR || t == spl1gen.SPL1ParserMINUS
}

func tokenAt(toks []antlr.Token, i int) antlr.Token {
	if i >= 0 && i < len(toks) {
		return toks[i]
	}
	return nil
}
