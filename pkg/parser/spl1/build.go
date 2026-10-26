package spl1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/spl1/gen"
)

// builder lowers parsed rule contexts into the IR. The grammar decides what
// the query means; the builder reproduces the hand parser's defaults, value
// shapes, positions, and deferred checks, and reports the operand errors the
// grammar deliberately leaves optional.
type builder struct {
	m      *antlrpos.SourceMap
	toks   []antlr.Token
	parser *spl1gen.SPL1Parser
}

func (b *builder) pos(t antlr.Token) ast.Pos { return b.m.Pos(t) }

// text returns the source text of a rule context, sliced from the first to the
// last token so raw spellings, escapes, and wildcards survive untouched.
func (b *builder) text(ctx antlr.ParserRuleContext) string {
	return b.m.Slice(ctx.GetStart(), ctx.GetStop())
}

// after returns the token that follows t in the stream. The parser has already
// fetched every token, so the index is stable even past the current position.
func (b *builder) after(t antlr.Token) antlr.Token {
	if t == nil {
		return nil
	}
	i := t.GetTokenIndex() + 1
	if i >= 0 && i < len(b.toks) {
		return b.toks[i]
	}
	return t
}

// missingOperand reports the first token after an operator whose operand is
// absent, which is where the hand parser fails.
func (b *builder) missingOperand(operator antlr.Token) *diag.Error {
	t := b.after(operator)
	return diag.NewError(b.pos(t), "syntax",
		"expected an expression, found %s", b.describe(t))
}

func (b *builder) describe(t antlr.Token) string {
	if t == nil || t.GetTokenType() == antlr.TokenEOF {
		return "end of query"
	}
	if tok, err := convertToken(b.m, t); err == nil {
		return tok.describe()
	}
	return fmt.Sprintf("%q", t.GetText())
}

// --- stages ---------------------------------------------------------------

func (b *builder) searchStage(ctx spl1gen.IStageSearchContext) (ast.Stage, error) {
	expr, err := b.searchExpr(ctx.SearchExpr())
	if err != nil {
		return nil, err
	}
	st := &ast.SearchStage{Expr: expr, Pos: b.pos(ctx.GetStart())}
	if err := extractTimeBounds(st); err != nil {
		return nil, err
	}
	return st, nil
}

func (b *builder) command(ctx spl1gen.IStageCommandContext) (ast.Stage, error) {
	return b.commandBody(ctx.Command())
}

func (b *builder) commandBody(ctx spl1gen.ICommandContext) (ast.Stage, error) {
	switch c := ctx.(type) {
	case *spl1gen.CmdSearchContext:
		var expr ast.Expr
		if c.SearchExpr() != nil {
			var err error
			expr, err = b.searchExpr(c.SearchExpr())
			if err != nil {
				return nil, err
			}
		}
		st := &ast.SearchStage{Expr: expr, Pos: b.pos(c.GetStart())}
		if err := extractTimeBounds(st); err != nil {
			return nil, err
		}
		return st, nil

	case *spl1gen.CmdWhereContext:
		var expr ast.Expr
		if c.ValueExpr() != nil {
			var err error
			expr, err = b.valueExpr(c.ValueExpr())
			if err != nil {
				return nil, err
			}
		}
		st := &ast.SearchStage{Expr: expr, Pos: b.pos(c.GetStart())}
		if err := extractTimeBounds(st); err != nil {
			return nil, err
		}
		return st, nil

	case *spl1gen.CmdStatsContext:
		aggs, err := b.aggList(c.AggList())
		if err != nil {
			return nil, err
		}
		var by []ast.Expr
		if c.ByClause() != nil {
			by, err = b.byClause(c.ByClause())
			if err != nil {
				return nil, err
			}
		}
		return &ast.StatsStage{Aggs: aggs, By: by, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdTimechartContext:
		return b.timechart(c.GetStart(), c.TimechartBody())

	case *spl1gen.CmdEvalContext:
		var assignments []ast.Assignment
		for _, a := range c.AllAssignment() {
			expr, err := b.valueExpr(a.ValueExpr())
			if err != nil {
				return nil, err
			}
			assignments = append(assignments, ast.Assignment{
				Name: a.IdentLike().GetStart().GetText(),
				Expr: expr,
				Pos:  b.pos(a.IdentLike().GetStart()),
			})
		}
		return &ast.EvalStage{Assignments: assignments, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdRexContext:
		field := "_raw"
		if c.RexField() != nil {
			field = c.RexField().IdentLike().GetStart().GetText()
		}
		pattern := decodeString(c.STRING().GetSymbol().GetText())
		return &ast.RexStage{Field: field, Pattern: pattern, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdTableContext:
		var fields []string
		wildcard := false
		for _, item := range c.AllTableItem() {
			switch {
			case item.COMMA() != nil:
			case item.Star() != nil:
				wildcard = true
			case item.IdentLike() != nil:
				fields = append(fields, item.IdentLike().GetStart().GetText())
			}
		}
		if len(fields) == 0 && !wildcard {
			return nil, diag.NewError(b.pos(c.GetStart()), "syntax",
				"table requires at least one field")
		}
		return &ast.TableStage{Fields: fields, Wildcard: wildcard, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdFieldsContext:
		exclude := c.FieldsSign() != nil && c.FieldsSign().Minus() != nil
		var include, excludeFields []string
		for _, item := range c.AllFieldsItem() {
			if item.IdentLike() == nil {
				continue // a separating comma
			}
			name := item.IdentLike().GetStart().GetText()
			if exclude {
				excludeFields = append(excludeFields, name)
			} else {
				include = append(include, name)
			}
		}
		if len(include) == 0 && len(excludeFields) == 0 {
			return nil, diag.NewError(b.pos(c.GetStart()), "syntax",
				"fields requires at least one field")
		}
		return &ast.FieldsStage{Include: include, Exclude: excludeFields, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdSortContext:
		return b.sort(c.GetStart(), c.SortBody())

	case *spl1gen.CmdHeadContext:
		count, countSet := 10, false
		if c.LimitArg() != nil {
			n, err := b.limitArg(c.LimitArg(), "head limit", "head count")
			if err != nil {
				return nil, err
			}
			count, countSet = n, true
		}
		return &ast.HeadStage{Count: count, CountSet: countSet, Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.CmdBinContext:
		return b.bin(c.GetStart(), c.AllBinItem())

	case *spl1gen.CmdRenameContext:
		renames := make([]ast.Rename, 0, len(c.AllRenamePair()))
		for _, pair := range c.AllRenamePair() {
			from, fromPos := b.renameField(pair.RenameField(0))
			to, toPos := b.renameField(pair.RenameField(1))
			renames = append(renames, ast.Rename{
				From: from, FromPos: fromPos,
				To: to, ToPos: toPos,
			})
		}
		return &ast.RenameStage{Renames: renames, Pos: b.pos(c.GetStart())}, nil
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized command")
}

func (b *builder) timechart(cmd antlr.Token, body spl1gen.ITimechartBodyContext) (ast.Stage, error) {
	span := ""
	var aggs []ast.Aggregate
	var by []ast.Expr
	var err error

	switch c := body.(type) {
	case *spl1gen.TcSpanFirstContext:
		span = b.spanValue(c.TimechartSpan().SpanAssign().SpanValue())
		aggs, err = b.aggList(c.AggList())
		if err != nil {
			return nil, err
		}
		if c.ByClause() != nil {
			by, err = b.byClause(c.ByClause())
			if err != nil {
				return nil, err
			}
		}
	case *spl1gen.TcAggFirstContext:
		aggs, err = b.aggListNotSpan(c.AggListNotSpan())
		if err != nil {
			return nil, err
		}
		if c.ByClause() != nil {
			by, err = b.byClause(c.ByClause())
			if err != nil {
				return nil, err
			}
		}
	}
	return &ast.TimechartStage{Aggs: aggs, By: by, Span: span, Pos: b.pos(cmd)}, nil
}

func (b *builder) aggList(ctx spl1gen.IAggListContext) ([]ast.Aggregate, error) {
	var out []ast.Aggregate
	for _, a := range ctx.AllAgg() {
		agg, err := b.aggregate(a)
		if err != nil {
			return nil, err
		}
		out = append(out, agg)
	}
	return out, nil
}

func (b *builder) aggListNotSpan(ctx spl1gen.IAggListNotSpanContext) ([]ast.Aggregate, error) {
	var out []ast.Aggregate
	first := ctx.AggNotSpan()
	head, err := b.aggregateNotSpan(first)
	if err != nil {
		return nil, err
	}
	out = append(out, head)
	for _, a := range ctx.AllAgg() {
		agg, err := b.aggregate(a)
		if err != nil {
			return nil, err
		}
		out = append(out, agg)
	}
	return out, nil
}

func (b *builder) aggregate(ctx spl1gen.IAggContext) (ast.Aggregate, error) {
	return b.aggregateParts(ctx.IdentLike().GetStart(), ctx.LPAREN() != nil,
		ctx.ValueExpr(), ctx.AliasClause())
}

func (b *builder) aggregateNotSpan(ctx spl1gen.IAggNotSpanContext) (ast.Aggregate, error) {
	return b.aggregateParts(ctx.IdentLikeNotSpan().GetStart(), ctx.LPAREN() != nil,
		ctx.ValueExpr(), ctx.AliasClause())
}

func (b *builder) aggregateParts(fnTok antlr.Token, parenthesized bool,
	arg spl1gen.IValueExprContext, alias spl1gen.IAliasClauseContext) (ast.Aggregate, error) {

	agg := ast.Aggregate{
		Func: strings.ToLower(fnTok.GetText()),
		Pos:  b.pos(fnTok),
	}
	if parenthesized && arg != nil {
		expr, err := b.valueExpr(arg)
		if err != nil {
			return ast.Aggregate{}, err
		}
		agg.Arg = expr
	}
	if alias != nil {
		agg.Alias = alias.IdentLike().GetStart().GetText()
	}
	return agg, nil
}

func (b *builder) byClause(ctx spl1gen.IByClauseContext) ([]ast.Expr, error) {
	var out []ast.Expr
	for _, fe := range ctx.FieldExprList().AllFieldExpr() {
		if star := fe.Star(); star != nil {
			out = append(out, &ast.WildcardLit{Pos: b.pos(star.GetStart())})
			continue
		}
		tok := fe.IdentLike().GetStart()
		out = append(out, &ast.FieldRef{Name: tok.GetText(), Pos: b.pos(tok)})
	}
	return out, nil
}

func (b *builder) spanValue(ctx spl1gen.ISpanValueContext) string {
	if d := ctx.DURATION(); d != nil {
		return d.GetSymbol().GetText()
	}
	return ctx.IdentLike().GetStart().GetText()
}

func (b *builder) sort(cmd antlr.Token, body spl1gen.ISortBodyContext) (ast.Stage, error) {
	limit := 0
	var fields []ast.SortField

	switch c := body.(type) {
	case *spl1gen.SortLimitContext:
		n, err := b.limitArg(c.LimitArg(), "sort limit", "sort limit")
		if err != nil {
			return nil, err
		}
		limit = n
		items, err := b.sortItems(c.AllSortItem())
		if err != nil {
			return nil, err
		}
		fields = items
	case *spl1gen.SortFirstFieldContext:
		first, err := b.sortFirstItem(c.SortFirstItem())
		if err != nil {
			return nil, err
		}
		if first != nil {
			fields = append(fields, *first)
		}
		items, err := b.sortItems(c.AllSortItem())
		if err != nil {
			return nil, err
		}
		fields = append(fields, items...)
	}

	if len(fields) == 0 && limit == 0 {
		return nil, diag.NewError(b.pos(cmd), "syntax", "sort requires at least one field")
	}
	return &ast.SortStage{Fields: fields, Limit: limit, Pos: b.pos(cmd)}, nil
}

func (b *builder) sortFirstItem(ctx spl1gen.ISortFirstItemContext) (*ast.SortField, error) {
	if ctx.COMMA() != nil {
		return nil, nil
	}
	// A sign rules out the `limit` keyword reading, so `sort -limit` is a
	// descending sort on a field named `limit` and lands in the sign
	// alternative rather than in sortField.
	if word := ctx.IdentLike(); word != nil {
		tok := word.GetStart()
		return &ast.SortField{
			Name: tok.GetText(),
			Desc: ctx.Sign() != nil && ctx.Sign().Minus() != nil,
		}, nil
	}
	return b.sortField(ctx.SortField(), ctx.Sign()), nil
}

func (b *builder) sortField(ctx spl1gen.ISortFieldContext, sign spl1gen.ISignContext) *ast.SortField {
	tok := ctx.IdentLikeNotLimit().GetStart()
	return &ast.SortField{Name: tok.GetText(), Desc: sign != nil && sign.Minus() != nil}
}

func (b *builder) sortItems(items []spl1gen.ISortItemContext) ([]ast.SortField, error) {
	var out []ast.SortField
	for _, item := range items {
		if item.IdentLike() == nil {
			continue // a separating comma
		}
		tok := item.IdentLike().GetStart()
		out = append(out, ast.SortField{
			Name: tok.GetText(),
			Desc: item.Sign() != nil && item.Sign().Minus() != nil,
		})
	}
	return out, nil
}

func (b *builder) limitArg(ctx spl1gen.ILimitArgContext, whatKeyword, whatNumber string) (int, error) {
	switch c := ctx.(type) {
	case *spl1gen.LimitKeywordContext:
		return b.nonNegativeInt(c.LimitTail().Number().GetStart(), whatKeyword)
	case *spl1gen.LimitNumberContext:
		return b.nonNegativeInt(c.Number().GetStart(), whatNumber)
	}
	return 0, nil
}

func (b *builder) nonNegativeInt(tok antlr.Token, what string) (int, error) {
	n, err := strconv.Atoi(tok.GetText())
	if err != nil || n < 0 {
		return 0, diag.NewError(b.pos(tok), "syntax",
			"expected a non-negative integer for %s, found %q", what, tok.GetText())
	}
	return n, nil
}

// bin lowers the items the grammar accepted. The single-field rule is enforced
// by the grammar itself (binFieldSeen), so a second field has already been
// refused at its own token by the time this runs.
func (b *builder) bin(cmd antlr.Token, items []spl1gen.IBinItemContext) (ast.Stage, error) {
	field, span := "", ""
	for _, item := range items {
		switch c := item.(type) {
		case *spl1gen.BinSpanItemContext:
			span = b.spanValue(c.SpanAssign().SpanValue())
		case *spl1gen.BinFieldItemContext:
			field = c.BinField().IdentLikeNotSpan().GetStart().GetText()
		}
	}
	if field == "" {
		return nil, diag.NewError(b.pos(cmd), "syntax", "bin requires a field")
	}
	if span == "" {
		return nil, diag.NewError(b.pos(cmd), "syntax", "bin requires span=<duration>")
	}
	return &ast.BinStage{Field: field, Span: span, Pos: b.pos(cmd)}, nil
}

// renameField returns a rename field's name and the position of its first
// token. A quoted name keeps its decoded text and may contain spaces; an
// identifier keeps its raw spelling, including a tight trailing wildcard.
func (b *builder) renameField(ctx spl1gen.IRenameFieldContext) (string, ast.Pos) {
	if s := ctx.STRING(); s != nil {
		tok := s.GetSymbol()
		return decodeString(tok.GetText()), b.pos(tok)
	}
	return b.text(ctx), b.pos(ctx.GetStart())
}

// --- search expressions: comparison RHS words are literal values ----------

func (b *builder) searchExpr(ctx spl1gen.ISearchExprContext) (ast.Expr, error) {
	left, err := b.searchAnd(ctx.SearchAnd())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllSearchOrTail() {
		orTok := tail.OR().GetSymbol()
		if tail.SearchAnd() == nil {
			return nil, b.missingOperand(orTok)
		}
		right, err := b.searchAnd(tail.SearchAnd())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: "or", Left: left, Right: right, Pos: b.pos(orTok)}
	}
	return left, nil
}

func (b *builder) searchAnd(ctx spl1gen.ISearchAndContext) (ast.Expr, error) {
	left, err := b.searchNot(ctx.SearchNot())
	if err != nil {
		return nil, err
	}
	for _, part := range ctx.AllSearchAndPart() {
		switch p := part.(type) {
		case *spl1gen.SAndExplicitContext:
			andTok := p.AND().GetSymbol()
			if p.SearchNot() == nil {
				return nil, b.missingOperand(andTok)
			}
			right, err := b.searchNot(p.SearchNot())
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: b.pos(andTok)}
		case *spl1gen.SAndImplicitContext:
			right, err := b.searchNot(p.SearchNot())
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: left.Position()}
		}
	}
	return left, nil
}

func (b *builder) searchNot(ctx spl1gen.ISearchNotContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.SNotUnaryContext:
		opTok := notOperator(c.NOT(), c.BANG())
		if c.SearchNot() == nil {
			return nil, b.missingOperand(opTok)
		}
		operand, err := b.searchNot(c.SearchNot())
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: "not", Operand: operand, Pos: b.pos(opTok)}, nil
	case *spl1gen.SNotPrimaryContext:
		return b.searchComparison(c.SearchComparison())
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) searchComparison(ctx spl1gen.ISearchComparisonContext) (ast.Expr, error) {
	left, err := b.searchAdditive(ctx.SearchAdditive())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllSearchComparisonTail() {
		opTok := tail.ComparisonOp().GetStart()
		if tail.SearchOperand() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.searchOperand(tail.SearchOperand())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    comparisonOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) searchOperand(ctx spl1gen.ISearchOperandContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.SOpStringContext:
		tok := c.STRING().GetSymbol()
		return &ast.StringLit{Value: decodeString(tok.GetText()), Pos: b.pos(tok)}, nil
	case *spl1gen.SOpValueContext:
		return &ast.StringLit{Value: b.text(c), Pos: b.pos(c.GetStart())}, nil
	case *spl1gen.SOpExprContext:
		return b.searchOperandExpr(c.SearchOperandExpr())
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) searchAdditive(ctx spl1gen.ISearchAdditiveContext) (ast.Expr, error) {
	left, err := b.searchMultiplicative(ctx.SearchMultiplicative())
	if err != nil {
		return nil, err
	}
	return b.applySearchAddTails(left, ctx.AllSearchAddTail())
}

func (b *builder) searchMultiplicative(ctx spl1gen.ISearchMultiplicativeContext) (ast.Expr, error) {
	left, err := b.searchUnary(ctx.SearchUnary())
	if err != nil {
		return nil, err
	}
	return b.applySearchMulTails(left, ctx.AllSearchMulTail())
}

// searchOperandExpr lowers the expression reading of a comparison RHS. Its
// first operand is an operand that cannot begin a value run; the operator
// tails are the ordinary search tails.
func (b *builder) searchOperandExpr(ctx spl1gen.ISearchOperandExprContext) (ast.Expr, error) {
	left, err := b.searchOperandUnary(ctx.SearchOperandUnary())
	if err != nil {
		return nil, err
	}
	left, err = b.applySearchMulTails(left, ctx.AllSearchMulTail())
	if err != nil {
		return nil, err
	}
	return b.applySearchAddTails(left, ctx.AllSearchAddTail())
}

func (b *builder) searchOperandUnary(ctx spl1gen.ISearchOperandUnaryContext) (ast.Expr, error) {
	var signTok antlr.Token
	switch {
	case ctx.MINUS() != nil:
		signTok = ctx.MINUS().GetSymbol()
	case ctx.PLUS() != nil:
		signTok = ctx.PLUS().GetSymbol()
	default:
		return b.searchOperandPrimary(ctx.SearchOperandPrimary())
	}
	if ctx.SearchUnary() == nil {
		return nil, b.missingOperand(signTok)
	}
	operand, err := b.searchUnary(ctx.SearchUnary())
	if err != nil {
		return nil, err
	}
	if isMinusToken(signTok) {
		return &ast.UnaryExpr{Op: "-", Operand: operand, Pos: b.pos(signTok)}, nil
	}
	return operand, nil
}

func (b *builder) searchOperandPrimary(ctx spl1gen.ISearchOperandPrimaryContext) (ast.Expr, error) {
	tok := ctx.GetStart()
	switch tok.GetTokenType() {
	case spl1gen.SPL1ParserLPAREN:
		return b.searchExpr(ctx.SearchExpr())
	case spl1gen.SPL1ParserDURATION:
		return &ast.DurationLit{Raw: tok.GetText(), Pos: b.pos(tok)}, nil
	case spl1gen.SPL1ParserNUMBER:
		value, err := strconv.ParseFloat(tok.GetText(), 64)
		if err != nil {
			return nil, diag.NewError(b.pos(tok), "syntax", "invalid number %q", tok.GetText())
		}
		return &ast.NumberLit{Raw: tok.GetText(), Value: value, Pos: b.pos(tok)}, nil
	case spl1gen.SPL1ParserSTAR:
		return &ast.WildcardLit{Pos: b.pos(tok)}, nil
	case spl1gen.SPL1ParserLBRACKET:
		return nil, unsupportedSubsearch(b.pos(tok))
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) applySearchAddTails(left ast.Expr, tails []spl1gen.ISearchAddTailContext) (ast.Expr, error) {
	for _, tail := range tails {
		opTok := tail.AddOp().GetStart()
		if tail.SearchMultiplicative() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.searchMultiplicative(tail.SearchMultiplicative())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    additiveOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) applySearchMulTails(left ast.Expr, tails []spl1gen.ISearchMulTailContext) (ast.Expr, error) {
	for _, tail := range tails {
		opTok := tail.SearchMulOp().GetStart()
		if tail.SearchUnary() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.searchUnary(tail.SearchUnary())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    multiplicativeOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) searchUnary(ctx spl1gen.ISearchUnaryContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.SUnarySignContext:
		signTok := c.Sign().GetStart()
		operand, err := b.searchUnary(c.SearchUnary())
		if err != nil {
			return nil, err
		}
		if isMinusToken(signTok) {
			return &ast.UnaryExpr{Op: "-", Operand: operand, Pos: b.pos(signTok)}, nil
		}
		return operand, nil
	case *spl1gen.SUnaryPrimaryContext:
		return b.searchPrimary(c.SearchPrimary())
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) searchPrimary(ctx spl1gen.ISearchPrimaryContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.SPrimParenContext:
		return b.searchExpr(c.SearchExpr())

	case *spl1gen.SPrimStringContext:
		tok := c.STRING().GetSymbol()
		pos := b.pos(tok)
		value := decodeString(tok.GetText())
		if b.comparisonAhead(tok) {
			return &ast.FieldRef{Name: value, Pos: pos}, nil
		}
		return &ast.SearchTerm{Value: value, Pos: pos}, nil

	case *spl1gen.SPrimWordContext:
		word := c.ExprWord().GetStart()
		if b.comparisonAhead(word) {
			return &ast.FieldRef{Name: word.GetText(), Pos: b.pos(word)}, nil
		}
		return &ast.SearchTerm{Value: b.text(c), Pos: b.pos(word)}, nil

	case *spl1gen.SPrimNumberContext:
		tok := c.Number().GetStart()
		value, err := strconv.ParseFloat(tok.GetText(), 64)
		if err != nil {
			return nil, diag.NewError(b.pos(tok), "syntax", "invalid number %q", tok.GetText())
		}
		return &ast.NumberLit{Raw: tok.GetText(), Value: value, Pos: b.pos(tok)}, nil

	case *spl1gen.SPrimDurationContext:
		tok := c.DURATION().GetSymbol()
		return &ast.DurationLit{Raw: tok.GetText(), Pos: b.pos(tok)}, nil

	case *spl1gen.SPrimStarRunContext:
		return &ast.SearchTerm{Value: b.text(c), Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.SPrimStarContext:
		return &ast.WildcardLit{Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.SPrimSubsearchContext:
		return nil, unsupportedSubsearch(b.pos(c.GetStart()))
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

// comparisonAhead mirrors the hand parser's comparisonAhead(1): the token
// immediately after a word or string decides whether it names a field.
func (b *builder) comparisonAhead(t antlr.Token) bool {
	switch b.after(t).GetTokenType() {
	case spl1gen.SPL1ParserEQ2, spl1gen.SPL1ParserEQ, spl1gen.SPL1ParserNE,
		spl1gen.SPL1ParserLT, spl1gen.SPL1ParserLE, spl1gen.SPL1ParserGT,
		spl1gen.SPL1ParserGE:
		return true
	}
	return false
}

// --- value expressions: bare words are field references -------------------

func (b *builder) valueExpr(ctx spl1gen.IValueExprContext) (ast.Expr, error) {
	left, err := b.valueAnd(ctx.ValueAnd())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllValueOrTail() {
		orTok := tail.OR().GetSymbol()
		if tail.ValueAnd() == nil {
			return nil, b.missingOperand(orTok)
		}
		right, err := b.valueAnd(tail.ValueAnd())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: "or", Left: left, Right: right, Pos: b.pos(orTok)}
	}
	return left, nil
}

func (b *builder) valueAnd(ctx spl1gen.IValueAndContext) (ast.Expr, error) {
	left, err := b.valueNot(ctx.ValueNot())
	if err != nil {
		return nil, err
	}
	for _, part := range ctx.AllValueAndPart() {
		switch p := part.(type) {
		case *spl1gen.VAndExplicitContext:
			andTok := p.AND().GetSymbol()
			if p.ValueNot() == nil {
				return nil, b.missingOperand(andTok)
			}
			right, err := b.valueNot(p.ValueNot())
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: b.pos(andTok)}
		case *spl1gen.VAndImplicitContext:
			right, err := b.valueNot(p.ValueNot())
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: left.Position()}
		}
	}
	return left, nil
}

func (b *builder) valueNot(ctx spl1gen.IValueNotContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.VNotUnaryContext:
		opTok := notOperator(c.NOT(), c.BANG())
		if c.ValueNot() == nil {
			return nil, b.missingOperand(opTok)
		}
		operand, err := b.valueNot(c.ValueNot())
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: "not", Operand: operand, Pos: b.pos(opTok)}, nil
	case *spl1gen.VNotPrimaryContext:
		return b.valueComparison(c.ValueComparison())
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) valueComparison(ctx spl1gen.IValueComparisonContext) (ast.Expr, error) {
	left, err := b.valueAdd(ctx.ValueAdd())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllValueComparisonTail() {
		opTok := tail.ComparisonOp().GetStart()
		if tail.ValueAdd() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.valueAdd(tail.ValueAdd())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    comparisonOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) valueAdd(ctx spl1gen.IValueAddContext) (ast.Expr, error) {
	left, err := b.valueMul(ctx.ValueMul())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllValueAddTail() {
		opTok := tail.AddOp().GetStart()
		if tail.ValueMul() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.valueMul(tail.ValueMul())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    additiveOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) valueMul(ctx spl1gen.IValueMulContext) (ast.Expr, error) {
	left, err := b.valueUnary(ctx.ValueUnary())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllValueMulTail() {
		opTok := tail.MulOp().GetStart()
		if tail.ValueUnary() == nil {
			return nil, b.missingOperand(opTok)
		}
		right, err := b.valueUnary(tail.ValueUnary())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{
			Op:    multiplicativeOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) valueUnary(ctx spl1gen.IValueUnaryContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.VUnarySignContext:
		signTok := c.Sign().GetStart()
		operand, err := b.valueUnary(c.ValueUnary())
		if err != nil {
			return nil, err
		}
		if isMinusToken(signTok) {
			return &ast.UnaryExpr{Op: "-", Operand: operand, Pos: b.pos(signTok)}, nil
		}
		return operand, nil
	case *spl1gen.VUnaryPrimaryContext:
		return b.valuePrimary(c.ValuePrimary())
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

func (b *builder) valuePrimary(ctx spl1gen.IValuePrimaryContext) (ast.Expr, error) {
	switch c := ctx.(type) {
	case *spl1gen.VPrimParenContext:
		return b.valueExpr(c.ValueExpr())

	case *spl1gen.VPrimCallContext:
		name := c.ExprWord().GetStart()
		args := make([]ast.Expr, 0, len(c.AllValueExpr()))
		for _, arg := range c.AllValueExpr() {
			expr, err := b.valueExpr(arg)
			if err != nil {
				return nil, err
			}
			args = append(args, expr)
		}
		return &ast.CallExpr{Name: name.GetText(), Args: args, Pos: b.pos(name)}, nil

	case *spl1gen.VPrimWordContext:
		tok := c.ExprWord().GetStart()
		return &ast.FieldRef{Name: tok.GetText(), Pos: b.pos(tok)}, nil

	case *spl1gen.VPrimStringContext:
		tok := c.STRING().GetSymbol()
		return &ast.StringLit{Value: decodeString(tok.GetText()), Pos: b.pos(tok)}, nil

	case *spl1gen.VPrimNumberContext:
		tok := c.Number().GetStart()
		value, err := strconv.ParseFloat(tok.GetText(), 64)
		if err != nil {
			return nil, diag.NewError(b.pos(tok), "syntax", "invalid number %q", tok.GetText())
		}
		return &ast.NumberLit{Raw: tok.GetText(), Value: value, Pos: b.pos(tok)}, nil

	case *spl1gen.VPrimDurationContext:
		tok := c.DURATION().GetSymbol()
		return &ast.DurationLit{Raw: tok.GetText(), Pos: b.pos(tok)}, nil

	case *spl1gen.VPrimStarContext:
		return &ast.WildcardLit{Pos: b.pos(c.GetStart())}, nil

	case *spl1gen.VPrimSubsearchContext:
		return nil, unsupportedSubsearch(b.pos(c.GetStart()))
	}
	return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "syntax", "unrecognized expression")
}

// --- helpers --------------------------------------------------------------

func unsupportedSubsearch(pos ast.Pos) *diag.Error {
	return diag.NewError(pos, "unsupported-subsearch", "subsearches are not implemented yet")
}

func notOperator(not antlr.TerminalNode, bang antlr.TerminalNode) antlr.Token {
	if not != nil {
		return not.GetSymbol()
	}
	return bang.GetSymbol()
}

func comparisonOperator(t int) string {
	switch t {
	case spl1gen.SPL1ParserEQ2, spl1gen.SPL1ParserEQ:
		return "="
	case spl1gen.SPL1ParserNE:
		return "!="
	case spl1gen.SPL1ParserLT:
		return "<"
	case spl1gen.SPL1ParserLE:
		return "<="
	case spl1gen.SPL1ParserGT:
		return ">"
	case spl1gen.SPL1ParserGE:
		return ">="
	}
	return ""
}

func additiveOperator(t int) string {
	if isMinusTokenType(t) {
		return "-"
	}
	return "+"
}

func multiplicativeOperator(t int) string {
	switch t {
	// STAR_TERM is a multiplication operator in value position: the driver
	// retypes every star that touches the preceding word, and `eval x=a*b`
	// multiplies even though the star is glued to `a`. Only the term rules
	// treat that spelling as part of a search term.
	case spl1gen.SPL1ParserSTAR, spl1gen.SPL1ParserSTAR_TIGHT,
		spl1gen.SPL1ParserSTAR_TERM, spl1gen.SPL1ParserSTAR_OP:
		return "*"
	case spl1gen.SPL1ParserSLASH:
		return "/"
	case spl1gen.SPL1ParserPERCENT:
		return "%"
	}
	return ""
}

// isMinusToken reports whether tok is a minus, in either spelling. The driver
// retypes a minus that begins a value run as MINUS_TIGHT, so operator code has
// to accept both.
func isMinusToken(tok antlr.Token) bool { return isMinusTokenType(tok.GetTokenType()) }

func isMinusTokenType(t int) bool {
	return t == spl1gen.SPL1ParserMINUS || t == spl1gen.SPL1ParserMINUS_TIGHT
}

// extractTimeBounds lifts `earliest=` and `latest=` out of the search
// expression and into dedicated IR fields. Splunk treats them as global
// search-time modifiers, so they are lifted out of AND and OR branches alike.
func extractTimeBounds(st *ast.SearchStage) error {
	if st.Expr == nil {
		return nil
	}
	expr, err := stripTimeBounds(st.Expr, st)
	if err != nil {
		return err
	}
	st.Expr = expr
	return nil
}

func stripTimeBounds(e ast.Expr, st *ast.SearchStage) (ast.Expr, error) {
	if u, ok := e.(*ast.UnaryExpr); ok && u.Op == "not" {
		inner, err := stripTimeBounds(u.Operand, st)
		if err != nil {
			return nil, err
		}
		if inner == nil {
			return nil, nil
		}
		u.Operand = inner
		return u, nil
	}
	b, ok := e.(*ast.BinaryExpr)
	if !ok {
		return e, nil
	}
	if b.Op == "and" || b.Op == "or" {
		left, err := stripTimeBounds(b.Left, st)
		if err != nil {
			return nil, err
		}
		right, err := stripTimeBounds(b.Right, st)
		if err != nil {
			return nil, err
		}
		// A lifted bound is TRUE in the row expression. TRUE OR X is TRUE,
		// so an OR branch cannot fall back to its sibling; TRUE AND X is X.
		if b.Op == "or" {
			if left == nil || right == nil {
				return nil, nil
			}
		} else {
			if left == nil {
				return right, nil
			}
			if right == nil {
				return left, nil
			}
		}
		b.Left, b.Right = left, right
		return b, nil
	}
	if b.Op != "=" {
		return e, nil
	}
	field, ok := b.Left.(*ast.FieldRef)
	if !ok {
		return e, nil
	}
	name := strings.ToLower(field.Name)
	if name != "earliest" && name != "latest" {
		return e, nil
	}
	raw, ok := timeValue(b.Right)
	if !ok {
		return e, nil
	}
	current := &st.Earliest
	if name == "latest" {
		current = &st.Latest
	}
	if *current != "" && *current != raw {
		return nil, diag.NewError(b.Pos, "conflicting-time-bounds",
			"conflicting %s bounds %q and %q", name, *current, raw)
	}
	*current = raw
	return nil, nil
}

func timeValue(e ast.Expr) (string, bool) {
	switch v := e.(type) {
	case *ast.StringLit:
		return v.Value, true
	case *ast.NumberLit:
		return v.Raw, true
	case *ast.DurationLit:
		return v.Raw, true
	case *ast.UnaryExpr:
		if v.Op == "-" {
			if d, ok := v.Operand.(*ast.DurationLit); ok {
				return "-" + d.Raw, true
			}
		}
	}
	return "", false
}
