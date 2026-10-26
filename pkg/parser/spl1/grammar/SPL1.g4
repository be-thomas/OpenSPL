// SPL1.g4 — classic Splunk SPL, parsed by ANTLR.
//
// The generated Go parser replaces the hand-written lexer and recursive
// descent parser. Three things stay outside the grammar on purpose:
//
//   - Stage sequencing. `stageCommand`/`stageSearch` are entry rules the Go
//     driver calls one stage at a time so a semantic error in stage 1 is
//     reported before a grammar error in stage 3, and so an unimplemented
//     command name can be classified before it is parsed.
//   - String decoding and value slicing. The lexer keeps raw token text; the
//     lowering code slices values from the source and decodes escapes.
//   - Command bodies that are empty on purpose (`table`, `fields`, `sort`,
//     `bin`) are accepted by the grammar and refused by the lowering at the
//     command token, matching the hand parser's positions.
//
// Every operand and every keyword value is required in the grammar. ANTLR
// then rejects a missing operand at the very token where the hand parser
// rejects it, and the fail-fast error listener aborts at that first error, so
// no later grammar error can overtake it.
//
// Classic SPL is context sensitive: an unquoted word on the right of a
// comparison in a search expression is a literal value, while the same word
// in `where`, `eval`, or an aggregate argument is a field reference. The
// grammar keeps two parallel expression rule sets (`search*` and `value*`)
// instead of a mutable mode flag. Keyword tokens stay usable as identifiers
// through `identLike`, so `by=1`, `table where`, and `eval search=...` work.

grammar SPL1;

// Tight token types. The driver retypes a STAR that touches a preceding word
// as STAR_TERM, because the wildcard belongs to the search term (`error*`), and
// the STAR, MINUS, and NUMBER tokens that begin a literal value run on the
// right of a comparison as STAR_TIGHT, MINUS_TIGHT, and NUMBER_TIGHT (see
// tight.go). It also retypes NOT after an arithmetic or comparison operator as
// NOT_WORD, because there the word is a field name (`1+not`), while at the
// start of an expression it is the negation operator and commits. These are
// not lexer rules: the lexer emits the ordinary token, and the retyping is a
// token-stream pass, because the decision depends on the surrounding tokens,
// past the end of the token the lexer is matching.
tokens { NUMBER_TIGHT, STAR_TIGHT, MINUS_TIGHT, STAR_TERM, NOT_WORD }

@parser::members {
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
}

// --- entry rules ---------------------------------------------------------
//
// The driver calls one of these per stage and inspects stageEnd to decide
// whether another stage follows.

stageCommand : command stageEnd ;

// An implicit search always has an expression; empty input and a leading pipe
// are handled by the driver before this rule is called.
stageSearch  : searchExpr stageEnd ;

stageEnd : PIPE | EOF ;

command
  : SEARCH searchExpr?                      #CmdSearch
  | WHERE valueExpr?                        #CmdWhere
  | STATS aggList byClause?                 #CmdStats
  | TIMECHART timechartBody                 #CmdTimechart
  | EVAL assignment (COMMA assignment)*     #CmdEval
  | REX rexField? STRING                    #CmdRex
  | TABLE tableItem*                        #CmdTable
  | FIELDS fieldsSign? fieldsItem*          #CmdFields
  | SORT sortBody                           #CmdSort
  | HEAD limitArg?                          #CmdHead
  | BIN binItem*                            #CmdBin
  ;

// Every operator rule below is gated on its own token type. Without the gate,
// ANTLR predicts the whole loop with unlimited lookahead and, seeing that the
// operand is missing, refuses to enter the loop; the error then lands on the
// operator. The hand parser commits to the operator and reports the missing
// operand after it (`a AND` at end of query, `a = = b` at the second '=').
// The gates make the grammar do the same, and they keep `a * b` arithmetic
// instead of letting '*' start an implicit-AND term.

// --- search expression: comparison RHS values are literals ---------------

searchExpr
  : searchAnd searchOrTail*
  ;

searchOrTail
  : {p.at(SPL1ParserOR)}? OR (searchAnd | absentOperand)
  ;

searchAnd
  : searchNot searchAndPart*
  ;

searchAndPart
  : {p.at(SPL1ParserAND)}? AND (searchNot | absentOperand) #SAndExplicit
  | {p.startsTerm()}? searchNot                   #SAndImplicit
  ;

searchNot
  : (NOT | BANG) (searchNot | absentOperand) #SNotUnary
  | searchComparison                             #SNotPrimary
  ;

searchComparison
  : searchAdditive searchComparisonTail*
  ;

searchComparisonTail
  : {p.atCompare()}? comparisonOp (searchOperand | absentOperand)
  ;

// A comparison RHS is a quoted string, a tight literal value run, or a search
// expression. The three alternatives must be decided by token type alone.
// ANTLR's SLL prediction never uses a semantic predicate to prune a
// simulation, so alternatives that share a first token and differ only by
// predicates cannot be told apart: every surviving alternative dies together
// at the end of the stage and the parser reports no viable alternative. That
// is why the value-run starts are distinct token types rather than a predicate
// over STAR, MINUS, and NUMBER. searchOperandExpr excludes those types, so the
// first sets stay disjoint and no predicate is needed at this decision.
searchOperand
  : STRING                                  #SOpString
  | searchOperandValue                      #SOpValue
  | searchOperandExpr                       #SOpExpr
  ;

// A value run starts on a word, or on a wildcard, hyphen, or number the driver
// marked tight, and runs on while its components touch. The alternatives are
// disjoint by token type.
searchOperandValue
  : identLike rhsTail*
  | NUMBER_TIGHT rhsTail*
  | STAR_TIGHT (identLike | number | minus (identLike | number | star)) rhsTail*
  | MINUS_TIGHT (identLike | star) rhsTail*
  ;

// Components must touch; whitespace ends the run.
rhsTail
  : {p.touchingPrev()}? star
  | {p.touchingPrev()}? identLike
  | {p.touchingPrev()}? number
  | {p.hyphenTail()}? minus (identLike | number | star)
  ;

// The expression alternative of a comparison RHS: an additive chain whose
// first operand can never begin a value run, so the alternatives above keep
// their own first tokens. Later operands are ordinary search expressions.
searchOperandExpr
  : searchOperandUnary searchMulTail* searchAddTail*
  ;

searchOperandUnary
  : (MINUS | PLUS) (searchUnary | absentOperand)
  | searchOperandPrimary
  ;

searchOperandPrimary
  : LPAREN searchExpr RPAREN
  | DURATION
  | NUMBER
  | STAR
  | subsearch
  ;

searchAdditive
  : searchMultiplicative searchAddTail*
  ;

searchAddTail
  : {p.atAdd()}? addOp (searchMultiplicative | absentOperand)
  ;

searchMultiplicative
  : searchUnary searchMulTail*
  ;

searchMulTail
  : {p.atSearchMul()}? searchMulOp (searchUnary | absentOperand)
  ;

searchUnary
  : sign (searchUnary | absentOperand)     #SUnarySign
  | searchPrimary                           #SUnaryPrimary
  ;

searchPrimary
  : LPAREN searchExpr RPAREN                #SPrimParen
  | STRING                                  #SPrimString
  | exprWord searchTermTail*                #SPrimWord
  | number                                  #SPrimNumber
  | DURATION                                #SPrimDuration
  | {p.termStarRun()}? star searchTermTail+ #SPrimStarRun
  | star                                    #SPrimStar
  | subsearch                               #SPrimSubsearch
  ;

// Free-text runs merge only adjacent wildcards and words: `a-b` stays
// arithmetic in term position. The driver retypes a STAR that touches the
// preceding word as STAR_TERM, so `a* | stats` is the term `a*` followed by a
// stage; without that the star could also begin a multiplication whose operand
// is missing.
searchTermTail
  : {p.touchingPrev()}? STAR_TERM
  | {p.touchingPrev()}? identLike
  | {p.touchingPrev()}? number
  ;

// --- value expression: bare words are field references -------------------

valueExpr
  : valueAnd valueOrTail*
  ;

valueOrTail
  : {p.at(SPL1ParserOR)}? OR (valueAnd | absentOperand)
  ;

valueAnd
  : valueNot valueAndPart*
  ;

valueAndPart
  : {p.at(SPL1ParserAND)}? AND (valueNot | absentOperand) #VAndExplicit
  | {p.startsTerm()}? valueNot                #VAndImplicit
  ;

valueNot
  : (NOT | BANG) (valueNot | absentOperand) #VNotUnary
  | valueComparison                          #VNotPrimary
  ;

valueComparison
  : valueAdd valueComparisonTail*
  ;

valueComparisonTail
  : {p.atCompare()}? comparisonOp (valueAdd | absentOperand)
  ;

valueAdd
  : valueMul valueAddTail*
  ;

valueAddTail
  : {p.atAdd()}? addOp (valueMul | absentOperand)
  ;

valueMul
  : valueUnary valueMulTail*
  ;

valueMulTail
  : {p.atMul()}? mulOp (valueUnary | absentOperand)
  ;

valueUnary
  : sign (valueUnary | absentOperand)      #VUnarySign
  | valuePrimary                            #VUnaryPrimary
  ;

valuePrimary
  : LPAREN valueExpr RPAREN                 #VPrimParen
  | {p.callAhead()}? exprWord LPAREN (valueExpr (COMMA valueExpr)*)? RPAREN #VPrimCall
  | exprWord                                #VPrimWord
  | STRING                                  #VPrimString
  | number                                  #VPrimNumber
  | DURATION                                #VPrimDuration
  | star                                    #VPrimStar
  | subsearch                               #VPrimSubsearch
  ;

// A subsearch consumes the rest of its stage so the lowering layer can refuse
// it at the bracket instead of reporting a grammar error further in.
subsearch
  : LBRACKET {p.abortSubsearch()} (~(PIPE | EOF))*
  ;

// absentOperand keeps every expression rule total: at a token that cannot
// begin a primary it ends the parse immediately, at the token where the hand
// parser fails. Totality is what lets prediction commit to an operator whose
// operand is missing, so the first diagnostic stays at the operand position
// instead of moving back to the operator.
absentOperand
  : {p.abortMissingOperand()} (EQ2 | EQ | NE | LT | LE | GT | GE | COMMA | RPAREN | RBRACKET | PIPE | EOF)
  ;

comparisonOp : EQ2 | EQ | NE | LT | LE | GT | GE ;

// `star`, `minus`, and `number` accept both spellings of a token, because the
// driver retypes only the ones that begin a value run. Every rule below them
// is shared by the value-run and expression reads, and none of them changes
// meaning with the tightness.
star  : STAR | STAR_TIGHT | STAR_TERM ;
minus : MINUS | MINUS_TIGHT ;
number : NUMBER | NUMBER_TIGHT ;

addOp : PLUS | minus ;
mulOp : star | SLASH | PERCENT ;
// Term-position multiplication excludes STAR_TERM: there the star already
// belongs to the preceding term.
searchMulOp : STAR | STAR_TIGHT | SLASH | PERCENT ;
sign : minus | PLUS ;

// --- identifiers and command bodies --------------------------------------

identLike
  : IDENT | SEARCH | WHERE | STATS | TIMECHART | EVAL | REX | TABLE | FIELDS
  | SORT | HEAD | BIN | BY | AS | SPAN | LIMIT | FIELD | AND | OR | NOT
  | NOT_WORD
  ;

// exprWord is identLike where a bare `not` cannot start an expression: there
// the keyword is the negation operator, and the driver retypes a `not` that
// follows an operator as NOT_WORD so it stays available as a field name.
exprWord
  : IDENT | SEARCH | WHERE | STATS | TIMECHART | EVAL | REX | TABLE | FIELDS
  | SORT | HEAD | BIN | BY | AS | SPAN | LIMIT | FIELD | AND | OR | NOT_WORD
  ;

aggList
  : agg (COMMA agg)*
  ;

agg
  : identLike (LPAREN valueExpr? RPAREN)? aliasClause?
  ;

// `timechart` may not take an aggregation named `span`: the first word after
// `timechart` is either the span keyword or an aggregation.
aggListNotSpan
  : aggNotSpan (COMMA agg)*
  ;

aggNotSpan
  : identLikeNotSpan (LPAREN valueExpr? RPAREN)? aliasClause?
  ;

aliasClause
  : AS identLike
  ;

byClause
  : BY fieldExprList
  ;

timechartBody
  : {p.at(SPL1ParserSPAN)}? timechartSpan aggList byClause?   #TcSpanFirst
  | aggListNotSpan byClause?                                  #TcAggFirst
  ;

timechartSpan
  : SPAN spanAssign
  ;

spanAssign
  : EQ spanValue
  ;

rexField
  : {p.at(SPL1ParserFIELD)}? FIELD EQ identLike
  ;

fieldExprList
  : fieldExpr (COMMA fieldExpr | fieldExpr)*
  ;

fieldExpr
  : identLike
  | star
  ;

assignment
  : identLike EQ valueExpr
  ;

spanValue
  : DURATION
  | identLike
  ;

tableItem
  : COMMA
  | star
  | identLike
  ;

fieldsSign
  : minus
  | PLUS
  ;

fieldsItem
  : COMMA
  | identLike
  ;

sortBody
  : limitArg sortItem*                      #SortLimit
  | sortFirstItem sortItem*                 #SortFirstField
  |                                         #SortEmpty
  ;

sortItem
  : COMMA
  | sign? identLike
  ;

// `limit` is the keyword only in the first body position; after a limit or a
// field it is an ordinary field name, so only the first item excludes it.
sortFirstItem
  : COMMA
  | sign? sortField
  ;

sortField
  : identLikeNotLimit
  ;

limitArg
  : {p.at(SPL1ParserLIMIT)}? LIMIT limitTail #LimitKeyword
  | number                                  #LimitNumber
  ;

limitTail
  : EQ number
  ;

binItem
  : {p.at(SPL1ParserSPAN)}? SPAN spanAssign  #BinSpanItem
  | binField                                 #BinFieldItem
  ;

// `span` is always the keyword inside `bin`, so it is not a field name here.
binField
  : identLikeNotSpan
  ;

// `span` excluded: `bin span` starts a span assignment, never a field.
identLikeNotSpan
  : IDENT | SEARCH | WHERE | STATS | TIMECHART | EVAL | REX | TABLE | FIELDS
  | SORT | HEAD | BIN | BY | AS | LIMIT | FIELD | AND | OR | NOT
  ;

// `limit` excluded: `sort limit` starts the limit argument, never a field.
identLikeNotLimit
  : IDENT | SEARCH | WHERE | STATS | TIMECHART | EVAL | REX | TABLE | FIELDS
  | SORT | HEAD | BIN | BY | AS | SPAN | FIELD | AND | OR | NOT
  ;

// --- lexer ---------------------------------------------------------------

// Keywords are ASCII fragments rather than the caseInsensitive option so the
// generated automaton is stable across ANTLR versions and targets.
SEARCH    : [sS][eE][aA][rR][cC][hH] ;
WHERE     : [wW][hH][eE][rR][eE] ;
STATS     : [sS][tT][aA][tT][sS] ;
TIMECHART : [tT][iI][mM][eE][cC][hH][aA][rR][tT] ;
EVAL      : [eE][vV][aA][lL] ;
REX       : [rR][eE][xX] ;
TABLE     : [tT][aA][bB][lL][eE] ;
FIELDS    : [fF][iI][eE][lL][dD][sS] ;
SORT      : [sS][oO][rR][tT] ;
HEAD      : [hH][eE][aA][dD] ;
BIN       : [bB][iI][nN] ;
BY        : [bB][yY] ;
AS        : [aA][sS] ;
SPAN      : [sS][pP][aA][nN] ;
LIMIT     : [lL][iI][mM][iI][tT] ;
FIELD     : [fF][iI][eE][lL][dD] ;
AND       : [aA][nN][dD] ;
OR        : [oO][rR] ;
NOT       : [nN][oO][tT] ;

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
