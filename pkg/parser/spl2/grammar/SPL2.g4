// SPL2.g4 — the supported subset of Splunk's SPL2, parsed by ANTLR.
//
// The subset is `from` followed by any number of `where` and `stats` stages.
// A separate grammar from SPL1.g4 on purpose: SPL2 has typed expressions, one
// meaning for a bare word (a field reference, even on the right of a
// comparison), no implicit AND, and a pipeline that always starts at `from`.
// Sharing SPL1's rules would pull classic-SPL quirks into the newer language.
//
// The lexical rules are duplicated rather than imported from SPL1.g4 because
// the Go target's grammar imports cross package boundaries awkwardly. The
// lexer-parity test tokenizes one corpus with both generated lexers and
// asserts identical text and positions, so the duplication cannot drift
// silently.
//
// As in SPL1, three things stay outside the grammar:
//
//   - Stage sequencing. `stageFrom`, `stageWhere`, and `stageStats` are entry
//     rules the Go driver calls one stage at a time, so a first stage that is
//     not `from` and a recognized-but-unimplemented command get their own
//     positioned diagnostics before anything is parsed.
//   - Dataset validation. The lowering refuses namespaced datasets, which the
//     grammar cannot know are unusable as Elasticsearch indexes.
//   - Index adjacency. `from logs-*` is one dataset name; the predicates below
//     compare token spans, because adjacency is not expressible in a token
//     rule.

grammar SPL2;

@parser::members {
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
}

// --- entry rules ---------------------------------------------------------
//
// The driver calls one of these per stage and inspects stageEnd to decide
// whether another stage follows.

stageFrom  : FROM indexName stageEnd ;
stageWhere : WHERE expr stageEnd ;
stageStats : STATS aggList byClause? stageEnd ;

stageEnd : PIPE | EOF ;

// --- datasets ------------------------------------------------------------

// A dataset is one unquoted name: words joined by adjacent hyphens, optionally
// ending in an adjacent wildcard. Every component must touch; whitespace ends
// the name.
indexName : identLike indexTail* ({p.touchingPrev()}? STAR)? ;

indexTail
  : {p.hyphenRun()}? MINUS (identLike | NUMBER | STAR)
  ;

// --- expressions ---------------------------------------------------------
//
// SPL2 is not context sensitive: a bare word is a field reference everywhere,
// including on the right of a comparison, and `AND` is required between
// conditions. There is no `*` primary, so `where status=*` and `stats count(*)`
// are syntax errors at the star rather than a guess about what a wildcard
// means in a typed expression.

expr
  : orExpr
  ;

orExpr
  : andExpr orTail*
  ;

orTail
  : {p.at(SPL2ParserOR)}? OR (andExpr | absentOperand)
  ;

andExpr
  : notExpr andTail*
  ;

andTail
  : {p.at(SPL2ParserAND)}? AND (notExpr | absentOperand)
  ;

notExpr
  : (NOT | BANG) (notExpr | absentOperand)
  | comparison
  ;

comparison
  : additive comparisonTail*
  ;

comparisonTail
  : {p.atCompare()}? compOp (additive | absentOperand)
  ;

additive
  : multiplicative additiveTail*
  ;

additiveTail
  : {p.atAdd()}? addOp (multiplicative | absentOperand)
  ;

multiplicative
  : unary multiplicativeTail*
  ;

multiplicativeTail
  : {p.atMul()}? mulOp (unary | absentOperand)
  ;

unary
  : sign (unary | absentOperand)
  | primary
  ;

primary
  : LPAREN expr RPAREN
  | {p.callAhead()}? identLike LPAREN (expr (COMMA expr)*)? RPAREN
  | identLike
  | STRING
  | NUMBER
  | DURATION
  | subsearch
  ;

// A subsearch consumes the rest of its stage so the lowering layer can refuse
// it at the bracket instead of reporting a grammar error further in.
subsearch
  : LBRACKET {p.abortSubsearch()} (~(PIPE | EOF))*
  ;

// absentOperand keeps every expression rule total: at a token that cannot
// begin a primary it ends the parse immediately, at the token where an operand
// was required. Totality is what lets prediction commit to an operator whose
// operand is missing, so the first diagnostic stays at the operand position
// instead of moving back to the operator.
absentOperand
  : {p.abortMissingOperand()} (EQ2 | EQ | NE | LT | LE | GT | GE | COMMA | RPAREN | RBRACKET | PIPE | EOF)
  ;

// --- aggregates ----------------------------------------------------------

aggList
  : agg (COMMA agg)*
  ;

// SPL2 requires the parentheses: `count()` aggregates the whole result set and
// `count(host)` aggregates a field. `stats count` is a syntax error.
agg
  : identLike LPAREN expr? RPAREN aliasClause?
  ;

aliasClause
  : AS identLike
  ;

byClause
  : BY fieldList
  ;

fieldList
  : field (COMMA field)*
  ;

field
  : identLike
  ;

compOp : EQ2 | EQ | NE | LT | LE | GT | GE ;
addOp : PLUS | MINUS ;
mulOp : STAR | SLASH | PERCENT ;
sign : MINUS | PLUS ;

identLike
  : IDENT | FROM | WHERE | STATS | BY | AS | AND | OR | NOT
  ;

// --- lexer ---------------------------------------------------------------
//
// The lexical rules below are identical to SPL1.g4's; see the header. Keywords
// are ASCII fragments rather than the caseInsensitive option so the generated
// automaton is stable across ANTLR versions and targets.

FROM  : [fF][rR][oO][mM] ;
WHERE : [wW][hH][eE][rR][eE] ;
STATS : [sS][tT][aA][tT][sS] ;
BY    : [bB][yY] ;
AS    : [aA][sS] ;
AND   : [aA][nN][dD] ;
OR    : [oO][rR] ;
NOT   : [nN][oO][tT] ;

// Durations are declared before numbers; maximal munch keeps `1mon` whole.
DURATION : DIGITS ('.' DIGITS)? EXPONENT? UNIT ;
NUMBER   : DIGITS ('.' DIGITS)? EXPONENT? ;

STRING : '"' DQ_BODY '"' | '\'' SQ_BODY '\'' ;

// A backslash escapes only the matching quote. Any other backslash stays a
// literal, including a run of them: `"a\\"` is unterminated because the last
// backslash escapes the closing quote.
fragment DQ_BODY : ('\\'+ '"' | '\\'+ ~["\\] | ~["\\])* ;
fragment SQ_BODY : ('\\'+ '\'' | '\\'+ ~['\\] | ~['\\])* ;

IDENT : [_@A-Za-z] [_@A-Za-z0-9.:]* ;

EQ2      : '==' ;
NE       : '!=' ;
LE       : '<=' ;
GE       : '>=' ;
EQ       : '=' ;
LT       : '<' ;
GT       : '>' ;
BANG     : '!' ;
PIPE     : '|' ;
COMMA    : ',' ;
LPAREN   : '(' ;
RPAREN   : ')' ;
LBRACKET : '[' ;
RBRACKET : ']' ;
PLUS     : '+' ;
MINUS    : '-' ;
STAR     : '*' ;
SLASH    : '/' ;
PERCENT  : '%' ;

WS : [ \t\r\n]+ -> skip ;

// The lexer never reports errors itself; the invalid-token pass in Go reports
// the first UNKNOWN token in source order with an exact position.
UNKNOWN : . ;

fragment DIGITS : [0-9]+ ;
fragment EXPONENT : [eE] [+-]? [0-9]+ ;
fragment UNIT : 'mon' | 'ms' | [smhdwy] ;
