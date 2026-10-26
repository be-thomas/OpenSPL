package spl2

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/internal/antlrpos"
	"github.com/be-thomas/openspl/pkg/parser/spl2/gen"
)

// builder lowers parsed rule contexts into the IR. The grammar decides what
// the query means; the builder carries the SPL2 defaults (aggregate name
// normalization, alias spelling), converts positions through the source map,
// and reports the semantic refusals the grammar deliberately leaves out.
type builder struct {
	m      *antlrpos.SourceMap
	toks   []antlr.Token
	parser *spl2gen.SPL2Parser
}

func (b *builder) pos(t antlr.Token) ast.Pos { return b.m.Pos(t) }

// text returns the source text of a rule context, sliced from the first to the
// last token so a dataset name keeps its exact spelling.
func (b *builder) text(ctx antlr.ParserRuleContext) string {
	return b.m.Slice(ctx.GetStart(), ctx.GetStop())
}

// --- stages ---------------------------------------------------------------

// from lowers `from <dataset>` to the same IR shape classic SPL produces for
// `index=<dataset>`, so the emitters resolve the index through one path.
func (b *builder) from(ctx spl2gen.IStageFromContext) (ast.Stage, error) {
	fromTok := ctx.FROM().GetSymbol()
	name := ctx.IndexName()
	first := name.GetStart()
	dataset := b.text(name)

	if strings.ContainsAny(dataset, ":") {
		return nil, diag.NewError(b.pos(first), "unsupported-dataset",
			"dataset %q is namespaced; only plain index names are supported", dataset)
	}

	expr := &ast.BinaryExpr{
		Op:    "=",
		Left:  &ast.FieldRef{Name: "index", Pos: b.pos(fromTok)},
		Right: &ast.StringLit{Value: dataset, Pos: b.pos(first)},
		Pos:   b.pos(fromTok),
	}
	return &ast.SearchStage{Expr: expr, Pos: b.pos(fromTok)}, nil
}

func (b *builder) where(ctx spl2gen.IStageWhereContext) (ast.Stage, error) {
	whereTok := ctx.WHERE().GetSymbol()
	expr, err := b.expr(ctx.Expr())
	if err != nil {
		return nil, err
	}
	return &ast.SearchStage{Expr: expr, Pos: b.pos(whereTok)}, nil
}

func (b *builder) stats(ctx spl2gen.IStageStatsContext) (ast.Stage, error) {
	statsTok := ctx.STATS().GetSymbol()
	aggs, err := b.aggList(ctx.AggList())
	if err != nil {
		return nil, err
	}
	var by []ast.Expr
	if clause := ctx.ByClause(); clause != nil {
		by, err = b.byClause(clause)
		if err != nil {
			return nil, err
		}
	}
	return &ast.StatsStage{Aggs: aggs, By: by, Pos: b.pos(statsTok)}, nil
}

// --- aggregates -----------------------------------------------------------

func (b *builder) aggList(ctx spl2gen.IAggListContext) ([]ast.Aggregate, error) {
	out := make([]ast.Aggregate, 0, len(ctx.AllAgg()))
	for _, a := range ctx.AllAgg() {
		agg, err := b.agg(a)
		if err != nil {
			return nil, err
		}
		out = append(out, agg)
	}
	return out, nil
}

func (b *builder) agg(ctx spl2gen.IAggContext) (ast.Aggregate, error) {
	fnTok := ctx.IdentLike().GetStart()
	agg := ast.Aggregate{
		Func: normalizeAggFunc(strings.ToLower(fnTok.GetText())),
		Pos:  b.pos(fnTok),
	}
	if arg := ctx.Expr(); arg != nil {
		expr, err := b.expr(arg)
		if err != nil {
			return ast.Aggregate{}, err
		}
		agg.Arg = expr
	}
	if alias := ctx.AliasClause(); alias != nil {
		agg.Alias = alias.IdentLike().GetStart().GetText()
	}
	return agg, nil
}

// normalizeAggFunc maps the SPL2 spellings of an aggregate onto the IR name the
// emitters know. An unknown name is left as written: the emitter refuses it
// with its own diagnostic.
func normalizeAggFunc(name string) string {
	switch name {
	case "distinct_count":
		return "dc"
	}
	return name
}

func (b *builder) byClause(ctx spl2gen.IByClauseContext) ([]ast.Expr, error) {
	list := ctx.FieldList()
	out := make([]ast.Expr, 0, len(list.AllField()))
	for _, fe := range list.AllField() {
		tok := fe.IdentLike().GetStart()
		out = append(out, &ast.FieldRef{Name: tok.GetText(), Pos: b.pos(tok)})
	}
	return out, nil
}

// --- expressions ----------------------------------------------------------

func (b *builder) expr(ctx spl2gen.IExprContext) (ast.Expr, error) {
	return b.orExpr(ctx.OrExpr())
}

func (b *builder) orExpr(ctx spl2gen.IOrExprContext) (ast.Expr, error) {
	left, err := b.andExpr(ctx.AndExpr())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllOrTail() {
		right, err := b.andExpr(tail.AndExpr())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: "or", Left: left, Right: right, Pos: b.pos(tail.OR().GetSymbol())}
	}
	return left, nil
}

func (b *builder) andExpr(ctx spl2gen.IAndExprContext) (ast.Expr, error) {
	left, err := b.notExpr(ctx.NotExpr())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllAndTail() {
		right, err := b.notExpr(tail.NotExpr())
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: b.pos(tail.AND().GetSymbol())}
	}
	return left, nil
}

func (b *builder) notExpr(ctx spl2gen.INotExprContext) (ast.Expr, error) {
	if ctx.Comparison() != nil {
		return b.comparison(ctx.Comparison())
	}
	opTok := notOperator(ctx.NOT(), ctx.BANG())
	operand, err := b.notExpr(ctx.NotExpr())
	if err != nil {
		return nil, err
	}
	return &ast.UnaryExpr{Op: "not", Operand: operand, Pos: b.pos(opTok)}, nil
}

func (b *builder) comparison(ctx spl2gen.IComparisonContext) (ast.Expr, error) {
	left, err := b.additive(ctx.Additive())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllComparisonTail() {
		right, err := b.additive(tail.Additive())
		if err != nil {
			return nil, err
		}
		opTok := tail.CompOp().GetStart()
		left = &ast.BinaryExpr{
			Op:    comparisonOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) additive(ctx spl2gen.IAdditiveContext) (ast.Expr, error) {
	left, err := b.multiplicative(ctx.Multiplicative())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllAdditiveTail() {
		right, err := b.multiplicative(tail.Multiplicative())
		if err != nil {
			return nil, err
		}
		opTok := tail.AddOp().GetStart()
		left = &ast.BinaryExpr{
			Op:    additiveOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) multiplicative(ctx spl2gen.IMultiplicativeContext) (ast.Expr, error) {
	left, err := b.unary(ctx.Unary())
	if err != nil {
		return nil, err
	}
	for _, tail := range ctx.AllMultiplicativeTail() {
		right, err := b.unary(tail.Unary())
		if err != nil {
			return nil, err
		}
		opTok := tail.MulOp().GetStart()
		left = &ast.BinaryExpr{
			Op:    multiplicativeOperator(opTok.GetTokenType()),
			Left:  left,
			Right: right,
			Pos:   b.pos(opTok),
		}
	}
	return left, nil
}

func (b *builder) unary(ctx spl2gen.IUnaryContext) (ast.Expr, error) {
	if ctx.Primary() != nil {
		return b.primary(ctx.Primary())
	}
	signTok := ctx.Sign().GetStart()
	operand, err := b.unary(ctx.Unary())
	if err != nil {
		return nil, err
	}
	op := "+"
	if signTok.GetTokenType() == spl2gen.SPL2ParserMINUS {
		op = "-"
	}
	return &ast.UnaryExpr{Op: op, Operand: operand, Pos: b.pos(signTok)}, nil
}

func (b *builder) primary(ctx spl2gen.IPrimaryContext) (ast.Expr, error) {
	switch ctx.GetStart().GetTokenType() {
	case spl2gen.SPL2ParserLPAREN:
		return b.expr(ctx.Expr(0))

	case spl2gen.SPL2ParserSTRING:
		tok := ctx.STRING().GetSymbol()
		return &ast.StringLit{Value: decodeString(tok.GetText()), Pos: b.pos(tok)}, nil

	case spl2gen.SPL2ParserNUMBER:
		tok := ctx.NUMBER().GetSymbol()
		value, err := strconv.ParseFloat(tok.GetText(), 64)
		if err != nil {
			return nil, diag.NewError(b.pos(tok), "syntax", "invalid number %q", tok.GetText())
		}
		return &ast.NumberLit{Raw: tok.GetText(), Value: value, Pos: b.pos(tok)}, nil

	case spl2gen.SPL2ParserDURATION:
		tok := ctx.DURATION().GetSymbol()
		return &ast.DurationLit{Raw: tok.GetText(), Pos: b.pos(tok)}, nil

	case spl2gen.SPL2ParserLBRACKET:
		// The grammar aborts at the bracket; this path is unreachable but
		// keeps the switch total without a nil dereference.
		return nil, diag.NewError(b.pos(ctx.GetStart()), "unsupported-subsearch",
			"subsearches are not supported")
	}

	name := ctx.IdentLike().GetStart()
	if ctx.LPAREN() == nil {
		return &ast.FieldRef{Name: name.GetText(), Pos: b.pos(name)}, nil
	}

	args := make([]ast.Expr, 0, len(ctx.AllExpr()))
	for _, arg := range ctx.AllExpr() {
		expr, err := b.expr(arg)
		if err != nil {
			return nil, err
		}
		args = append(args, expr)
	}
	return &ast.CallExpr{Name: name.GetText(), Args: args, Pos: b.pos(name)}, nil
}

// --- helpers --------------------------------------------------------------

func notOperator(not antlr.TerminalNode, bang antlr.TerminalNode) antlr.Token {
	if not != nil {
		return not.GetSymbol()
	}
	return bang.GetSymbol()
}

func comparisonOperator(t int) string {
	switch t {
	case spl2gen.SPL2ParserEQ2, spl2gen.SPL2ParserEQ:
		return "="
	case spl2gen.SPL2ParserNE:
		return "!="
	case spl2gen.SPL2ParserLT:
		return "<"
	case spl2gen.SPL2ParserLE:
		return "<="
	case spl2gen.SPL2ParserGT:
		return ">"
	case spl2gen.SPL2ParserGE:
		return ">="
	}
	return ""
}

func additiveOperator(t int) string {
	if t == spl2gen.SPL2ParserMINUS {
		return "-"
	}
	return "+"
}

func multiplicativeOperator(t int) string {
	switch t {
	case spl2gen.SPL2ParserSTAR:
		return "*"
	case spl2gen.SPL2ParserSLASH:
		return "/"
	case spl2gen.SPL2ParserPERCENT:
		return "%"
	}
	return ""
}
