// Package spl1 implements the classic Splunk SPL front-end.
//
// The parser lowers source syntax into the target-agnostic IR in pkg/ast. It
// deliberately knows nothing about ES|QL, PromQL, or any other back-end.
package spl1

import (
	"strconv"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
)

// implementedCommands are the stages this front-end can parse today.
var implementedCommands = map[string]bool{
	"search":    true,
	"where":     true,
	"stats":     true,
	"timechart": true,
	"eval":      true,
	"rex":       true,
	"table":     true,
	"fields":    true,
	"sort":      true,
	"head":      true,
	"bin":       true,
}

// knownCommands are real SPL commands we recognize but do not implement yet.
// Recognizing them produces a precise "not implemented" error instead of a
// confusing parse failure that treats the command as a search term.
var knownCommands = map[string]bool{
	"append": true, "appendcols": true, "addtotals": true, "analyzefields": true,
	"anomalydetection": true, "apply": true, "associate": true, "bucket": true,
	"chart": true, "cluster": true, "collect": true, "contingency": true,
	"correlate": true, "datamodel": true, "dbinspect": true, "dedup": true,
	"delete": true, "diff": true, "entitlement": true, "eventcount": true,
	"eventstats": true, "extract": true, "fieldformat": true, "fieldsummary": true,
	"filldown": true, "fillnull": true, "findtypes": true, "fit": true,
	"folderize": true, "foreach": true, "format": true, "from": true,
	"geostats": true, "head": true, "highlight": true, "inputcsv": true,
	"inputlookup": true, "iplocation": true, "join": true, "kmeans": true,
	"kvform": true, "loadjob": true, "localize": true, "lookup": true,
	"makecontinuous": true, "makeresults": true, "map": true, "mcollect": true,
	"metadata": true, "metasearch": true, "ml": true, "mstats": true,
	"multikmeans": true, "multisearch": true, "mvcombine": true, "mvexpand": true,
	"nomv": true, "outlier": true, "outputcsv": true, "outputlookup": true,
	"outputtext": true, "pivot": true, "predict": true, "rangemap": true,
	"rare": true, "regex": true, "reltime": true, "rename": true, "replace": true,
	"rest": true, "return": true, "reverse": true, "rex": true, "script": true,
	"scrub": true, "search": true, "select": true, "selfjoin": true,
	"sendemail": true, "set": true, "setfields": true, "spath": true,
	"stats": true, "strcat": true, "streamstats": true, "table": true,
	"tags": true, "tail": true, "timechart": true, "top": true,
	"transaction": true, "transpose": true, "trendline": true, "tscollect": true,
	"tstats": true, "typeahead": true, "typer": true, "union": true,
	"uniq": true, "untable": true, "where": true, "x11": true, "xmlkv": true,
	"xmlunescape": true, "xpath": true,
}

func isImplementedCommand(s string) bool { return implementedCommands[s] }
func isKnownCommand(s string) bool       { return knownCommands[s] }

type parser struct {
	toks []Token
	i    int

	// searchLiterals is true while parsing a `search` expression. In that
	// context an unquoted identifier on the right of a comparison is a value
	// (host=web means host == "web"), whereas in eval/where it is a field.
	searchLiterals bool
}

// Parse parses a classic SPL query into the unified IR.
func Parse(src string) (*ast.Pipeline, error) {
	if strings.TrimSpace(src) == "" {
		return nil, diag.NewError(ast.Pos{Line: 1, Col: 1}, "empty-query", "query is empty")
	}
	toks, err := Lex(src)
	if err != nil {
		return nil, err
	}
	p := &parser{toks: toks}
	return p.parsePipeline()
}

func (p *parser) parsePipeline() (*ast.Pipeline, error) {
	pl := &ast.Pipeline{}

	// A leading pipe is tolerated: `| stats count` means the same as
	// `stats count` with an implicit match-all search.
	if p.cur().Kind == Pipe {
		p.advance()
	}
	if p.cur().Kind == EOF {
		return nil, diag.NewError(p.cur().Pos, "empty-query", "query is empty")
	}

	first, err := p.parseFirstStage()
	if err != nil {
		return nil, err
	}
	pl.Stages = append(pl.Stages, first)

	for p.cur().Kind == Pipe {
		pipe := p.cur().Pos
		p.advance()
		if p.cur().Kind == EOF {
			return nil, diag.NewError(pipe, "syntax", "expected a command after '|'")
		}
		st, err := p.parseCommand()
		if err != nil {
			return nil, err
		}
		pl.Stages = append(pl.Stages, st)
	}
	if p.cur().Kind != EOF {
		return nil, diag.NewError(p.cur().Pos, "syntax",
			"unexpected %s; expected '|' or end of query", p.cur().describe())
	}
	return pl, nil
}

func (p *parser) parseFirstStage() (ast.Stage, error) {
	if p.cur().Kind == Ident {
		word := strings.ToLower(p.cur().Text)
		if isImplementedCommand(word) || isKnownCommand(word) {
			return p.parseCommand()
		}
	}
	return p.parseImplicitSearch()
}

func (p *parser) parseImplicitSearch() (ast.Stage, error) {
	pos := p.cur().Pos
	expr, err := p.parseExprWith(true)
	if err != nil {
		return nil, err
	}
	st := &ast.SearchStage{Expr: expr, Pos: pos}
	if err := extractTimeBounds(st); err != nil {
		return nil, err
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return st, nil
}

func (p *parser) parseCommand() (ast.Stage, error) {
	pos := p.cur().Pos
	if p.cur().Kind != Ident {
		return nil, diag.NewError(pos, "syntax", "expected a command, found %s", p.cur().describe())
	}
	name := strings.ToLower(p.cur().Text)
	p.advance()

	switch name {
	case "search":
		return p.parseSearchCommand(pos, true)
	case "where":
		return p.parseSearchCommand(pos, false)
	case "stats":
		return p.parseStats(pos)
	case "timechart":
		return p.parseTimechart(pos)
	case "eval":
		return p.parseEval(pos)
	case "rex":
		return p.parseRex(pos)
	case "table":
		return p.parseTable(pos)
	case "fields":
		return p.parseFields(pos)
	case "sort":
		return p.parseSort(pos)
	case "head":
		return p.parseHead(pos)
	case "bin":
		return p.parseBin(pos)
	default:
		if isKnownCommand(name) {
			return nil, diag.NewError(pos, "unsupported-command",
				"command %q is recognized but not implemented yet; see docs/COMPATIBILITY.md", name)
		}
		return nil, diag.NewError(pos, "unknown-command", "unknown command %q", name)
	}
}

func (p *parser) parseSearchCommand(pos ast.Pos, literals bool) (ast.Stage, error) {
	var expr ast.Expr
	if !p.atStageEnd() {
		var err error
		expr, err = p.parseExprWith(literals)
		if err != nil {
			return nil, err
		}
	}
	st := &ast.SearchStage{Expr: expr, Pos: pos}
	if err := extractTimeBounds(st); err != nil {
		return nil, err
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return st, nil
}

func (p *parser) parseStats(pos ast.Pos) (ast.Stage, error) {
	aggs, err := p.parseAggregates()
	if err != nil {
		return nil, err
	}
	var by []ast.Expr
	if p.matchKeyword("by") {
		by, err = p.parseFieldExprList()
		if err != nil {
			return nil, err
		}
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.StatsStage{Aggs: aggs, By: by, Pos: pos}, nil
}

func (p *parser) parseTimechart(pos ast.Pos) (ast.Stage, error) {
	span := ""
	if p.matchKeyword("span") {
		if err := p.expectKind(Eq, "'='"); err != nil {
			return nil, err
		}
		var err error
		span, err = p.parseSpanValue()
		if err != nil {
			return nil, err
		}
	}
	aggs, err := p.parseAggregates()
	if err != nil {
		return nil, err
	}
	var by []ast.Expr
	if p.matchKeyword("by") {
		by, err = p.parseFieldExprList()
		if err != nil {
			return nil, err
		}
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.TimechartStage{Aggs: aggs, By: by, Span: span, Pos: pos}, nil
}

func (p *parser) parseAggregates() ([]ast.Aggregate, error) {
	var aggs []ast.Aggregate
	for {
		if p.cur().Kind != Ident {
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected an aggregation function, found %s", p.cur().describe())
		}
		pos := p.cur().Pos
		fn := strings.ToLower(p.cur().Text)
		p.advance()

		var arg ast.Expr
		if p.matchKind(LParen) {
			if p.cur().Kind != RParen {
				var err error
				arg, err = p.parseExprWith(false)
				if err != nil {
					return nil, err
				}
			}
			if err := p.expectKind(RParen, "')'"); err != nil {
				return nil, err
			}
		}

		alias := ""
		if p.matchKeyword("as") {
			if p.cur().Kind != Ident {
				return nil, diag.NewError(p.cur().Pos, "syntax",
					"expected an alias after 'as', found %s", p.cur().describe())
			}
			alias = p.cur().Text
			p.advance()
		}

		aggs = append(aggs, ast.Aggregate{Func: fn, Arg: arg, Alias: alias, Pos: pos})

		if p.matchKind(Comma) {
			continue
		}
		break
	}
	if len(aggs) == 0 {
		return nil, diag.NewError(p.cur().Pos, "syntax", "stats requires at least one aggregation")
	}
	return aggs, nil
}

func (p *parser) parseFieldExprList() ([]ast.Expr, error) {
	var out []ast.Expr
	for {
		switch p.cur().Kind {
		case Ident:
			out = append(out, &ast.FieldRef{Name: p.cur().Text, Pos: p.cur().Pos})
			p.advance()
		case Star:
			out = append(out, &ast.WildcardLit{Pos: p.cur().Pos})
			p.advance()
		default:
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name, found %s", p.cur().describe())
		}
		if p.matchKind(Comma) {
			continue
		}
		if p.cur().Kind == Ident || p.cur().Kind == Star {
			continue // Splunk also accepts space-separated field lists.
		}
		break
	}
	if len(out) == 0 {
		return nil, diag.NewError(p.cur().Pos, "syntax", "expected at least one field after 'by'")
	}
	return out, nil
}

func (p *parser) parseEval(pos ast.Pos) (ast.Stage, error) {
	var assignments []ast.Assignment
	for {
		if p.cur().Kind != Ident {
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name in eval, found %s", p.cur().describe())
		}
		name := p.cur().Text
		namePos := p.cur().Pos
		p.advance()
		if err := p.expectKind(Eq, "'='"); err != nil {
			return nil, err
		}
		expr, err := p.parseExprWith(false)
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, ast.Assignment{Name: name, Expr: expr, Pos: namePos})
		if p.matchKind(Comma) {
			continue
		}
		break
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.EvalStage{Assignments: assignments, Pos: pos}, nil
}

func (p *parser) parseRex(pos ast.Pos) (ast.Stage, error) {
	field := "_raw"
	if p.matchKeyword("field") {
		if err := p.expectKind(Eq, "'='"); err != nil {
			return nil, err
		}
		if p.cur().Kind != Ident {
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name after 'field=', found %s", p.cur().describe())
		}
		field = p.cur().Text
		p.advance()
	}
	if p.cur().Kind != String {
		return nil, diag.NewError(p.cur().Pos, "syntax",
			"rex requires a quoted regular expression, found %s", p.cur().describe())
	}
	pattern := p.cur().Text
	p.advance()
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.RexStage{Field: field, Pattern: pattern, Pos: pos}, nil
}

func (p *parser) parseTable(pos ast.Pos) (ast.Stage, error) {
	var fields []string
	wildcard := false
	for !p.atStageEnd() {
		switch p.cur().Kind {
		case Comma:
			p.advance()
		case Star:
			wildcard = true
			p.advance()
		case Ident:
			fields = append(fields, p.cur().Text)
			p.advance()
		default:
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name in table, found %s", p.cur().describe())
		}
	}
	if len(fields) == 0 && !wildcard {
		return nil, diag.NewError(pos, "syntax", "table requires at least one field")
	}
	return &ast.TableStage{Fields: fields, Wildcard: wildcard, Pos: pos}, nil
}

func (p *parser) parseFields(pos ast.Pos) (ast.Stage, error) {
	exclude := false
	if p.matchKind(Minus) {
		exclude = true
	} else {
		p.matchKind(Plus)
	}

	var include, excludeFields []string
	for !p.atStageEnd() {
		if p.matchKind(Comma) {
			continue
		}
		if p.cur().Kind != Ident {
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name in fields, found %s", p.cur().describe())
		}
		if exclude {
			excludeFields = append(excludeFields, p.cur().Text)
		} else {
			include = append(include, p.cur().Text)
		}
		p.advance()
	}
	if len(include) == 0 && len(excludeFields) == 0 {
		return nil, diag.NewError(pos, "syntax", "fields requires at least one field")
	}
	return &ast.FieldsStage{Include: include, Exclude: excludeFields, Pos: pos}, nil
}

func (p *parser) parseSort(pos ast.Pos) (ast.Stage, error) {
	limit := 0
	if p.matchKeyword("limit") {
		if err := p.expectKind(Eq, "'='"); err != nil {
			return nil, err
		}
		n, err := p.parseNonNegativeInt("sort limit")
		if err != nil {
			return nil, err
		}
		limit = n
	} else if p.cur().Kind == Number {
		n, err := p.parseNonNegativeInt("sort limit")
		if err != nil {
			return nil, err
		}
		limit = n
	}

	var fields []ast.SortField
	for !p.atStageEnd() {
		if p.matchKind(Comma) {
			continue
		}
		desc := false
		if p.matchKind(Minus) {
			desc = true
		} else {
			p.matchKind(Plus)
		}
		if p.cur().Kind != Ident {
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field name in sort, found %s", p.cur().describe())
		}
		fields = append(fields, ast.SortField{Name: p.cur().Text, Desc: desc})
		p.advance()
	}
	if len(fields) == 0 && limit == 0 {
		return nil, diag.NewError(pos, "syntax", "sort requires at least one field")
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.SortStage{Fields: fields, Limit: limit, Pos: pos}, nil
}

func (p *parser) parseHead(pos ast.Pos) (ast.Stage, error) {
	count := 10 // Splunk's default
	countSet := false
	if p.matchKeyword("limit") {
		if err := p.expectKind(Eq, "'='"); err != nil {
			return nil, err
		}
		n, err := p.parseNonNegativeInt("head limit")
		if err != nil {
			return nil, err
		}
		count, countSet = n, true
	} else if p.cur().Kind == Number {
		n, err := p.parseNonNegativeInt("head count")
		if err != nil {
			return nil, err
		}
		count, countSet = n, true
	}
	if err := p.expectStageEnd(); err != nil {
		return nil, err
	}
	return &ast.HeadStage{Count: count, CountSet: countSet, Pos: pos}, nil
}

func (p *parser) parseBin(pos ast.Pos) (ast.Stage, error) {
	field, span := "", ""
	for !p.atStageEnd() {
		switch {
		case p.matchKeyword("span"):
			if err := p.expectKind(Eq, "'='"); err != nil {
				return nil, err
			}
			var err error
			span, err = p.parseSpanValue()
			if err != nil {
				return nil, err
			}
		case p.cur().Kind == Ident:
			if field != "" {
				return nil, diag.NewError(p.cur().Pos, "syntax",
					"bin accepts exactly one field; found %q after %q", p.cur().Text, field)
			}
			field = p.cur().Text
			p.advance()
		default:
			return nil, diag.NewError(p.cur().Pos, "syntax",
				"expected a field or span=<duration> in bin, found %s", p.cur().describe())
		}
	}
	if field == "" {
		return nil, diag.NewError(pos, "syntax", "bin requires a field")
	}
	if span == "" {
		return nil, diag.NewError(pos, "syntax", "bin requires span=<duration>")
	}
	return &ast.BinStage{Field: field, Span: span, Pos: pos}, nil
}

func (p *parser) parseSpanValue() (string, error) {
	switch p.cur().Kind {
	case Duration, Ident:
		v := p.cur().Text
		p.advance()
		return v, nil
	default:
		return "", diag.NewError(p.cur().Pos, "syntax",
			"expected a duration such as 5m, found %s", p.cur().describe())
	}
}

func (p *parser) parseNonNegativeInt(what string) (int, error) {
	if p.cur().Kind != Number {
		return 0, diag.NewError(p.cur().Pos, "syntax",
			"expected a non-negative integer for %s, found %s", what, p.cur().describe())
	}
	n, err := strconv.Atoi(p.cur().Text)
	if err != nil || n < 0 {
		return 0, diag.NewError(p.cur().Pos, "syntax",
			"expected a non-negative integer for %s, found %q", what, p.cur().Text)
	}
	p.advance()
	return n, nil
}

// --- expressions ---------------------------------------------------------

func (p *parser) parseExprWith(literals bool) (ast.Expr, error) {
	old := p.searchLiterals
	p.searchLiterals = literals
	defer func() { p.searchLiterals = old }()
	return p.parseOr()
}

func (p *parser) parseOr() (ast.Expr, error) {
	left, err := p.parseAnd()
	if err != nil {
		return nil, err
	}
	for p.matchKeyword("or") {
		pos := p.prevPos()
		right, err := p.parseAnd()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: "or", Left: left, Right: right, Pos: pos}
	}
	return left, nil
}

func (p *parser) parseAnd() (ast.Expr, error) {
	left, err := p.parseNot()
	if err != nil {
		return nil, err
	}
	for {
		if p.matchKeyword("and") {
			pos := p.prevPos()
			right, err := p.parseNot()
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: pos}
			continue
		}
		// SPL treats adjacent terms as an implicit AND.
		if p.startsTerm() {
			right, err := p.parseNot()
			if err != nil {
				return nil, err
			}
			left = &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: left.Position()}
			continue
		}
		break
	}
	return left, nil
}

func (p *parser) parseNot() (ast.Expr, error) {
	if p.matchKeyword("not") || p.matchKind(Bang) {
		pos := p.prevPos()
		operand, err := p.parseNot()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: "not", Operand: operand, Pos: pos}, nil
	}
	return p.parseComparison()
}

func (p *parser) parseComparison() (ast.Expr, error) {
	left, err := p.parseAdditive()
	if err != nil {
		return nil, err
	}
	for {
		var op string
		switch p.cur().Kind {
		case Eq:
			op = "="
		case Ne:
			op = "!="
		case Lt:
			op = "<"
		case Le:
			op = "<="
		case Gt:
			op = ">"
		case Ge:
			op = ">="
		default:
			return left, nil
		}
		pos := p.cur().Pos
		p.advance()
		right, err := p.parseComparisonOperand()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right, Pos: pos}
	}
}

func (p *parser) parseComparisonOperand() (ast.Expr, error) {
	// In a search expression `host=web*` and `host=web` use bare values, not
	// field references. Wildcards and embedded hyphens bind tightly to the
	// value token, so web-*, *error*, web*prod, and logs-2024-01 are single
	// values rather than arithmetic.
	if p.searchLiterals {
		if p.cur().Kind == String {
			pos, value := p.cur().Pos, p.cur().Text
			p.advance()
			return &ast.StringLit{Value: value, Pos: pos}, nil
		}
		pos := p.cur().Pos
		if value, ok := p.scanSearchValue(true); ok {
			return &ast.StringLit{Value: value, Pos: pos}, nil
		}
	}
	return p.parseAdditive()
}

// scanSearchValue consumes an unquoted search value starting at the current
// token. Components must touch: whitespace always ends the run. allowHyphen
// enables comparison-RHS values such as logs-2024-01; free-text terms leave a
// `-` to the additive parser, where it stays arithmetic. The cursor does not
// move when the function reports false.
func (p *parser) scanSearchValue(allowHyphen bool) (string, bool) {
	if !p.startsSearchValue(allowHyphen) {
		return "", false
	}
	tok := p.cur()

	var b strings.Builder
	b.WriteString(tok.Text)
	p.advance()

scan:
	for {
		prev, cur := p.prevTok(), p.cur()
		if cur.Kind == EOF || prev.Pos.Offset+len(prev.Text) != cur.Pos.Offset {
			break
		}
		switch cur.Kind {
		case Star:
			b.WriteString("*")
			p.advance()
		case Ident, Number:
			b.WriteString(cur.Text)
			p.advance()
		case Minus:
			if !allowHyphen {
				break scan
			}
			next := p.peek(1)
			if (next.Kind != Ident && next.Kind != Number && next.Kind != Star) ||
				cur.Pos.Offset+len(cur.Text) != next.Pos.Offset {
				break scan
			}
			b.WriteString("-")
			b.WriteString(next.Text)
			p.advance()
			p.advance()
		default:
			break scan
		}
	}
	return b.String(), true
}

// startsSearchValue reports whether the current token begins a tight value
// run. Numeric and signed starts are accepted only in comparison-RHS context
// (allowHyphen), so a bare `500` stays a NumberLit while `200-299` and `500*`
// are single literal values. A lone `*` stays a WildcardLit.
func (p *parser) startsSearchValue(allowHyphen bool) bool {
	switch p.cur().Kind {
	case Ident:
		return true
	case Number:
		// Only start a run when a continuation makes it a hyphen or wildcard
		// value; plain numbers keep their numeric literal semantics.
		return allowHyphen && p.adjacent(0, 1) &&
			(p.peek(1).Kind == Minus || p.peek(1).Kind == Star)
	case Star:
		if !p.adjacent(0, 1) {
			return false
		}
		switch p.peek(1).Kind {
		case Ident, Number:
			return true
		case Minus:
			if !allowHyphen || !p.adjacent(1, 1) {
				return false
			}
			switch p.peek(2).Kind {
			case Ident, Number, Star:
				return true
			}
		}
		return false
	case Minus:
		if !allowHyphen || !p.adjacent(0, 1) {
			return false
		}
		return p.peek(1).Kind == Ident || p.peek(1).Kind == Star
	}
	return false
}

func (p *parser) parseAdditive() (ast.Expr, error) {
	left, err := p.parseMultiplicative()
	if err != nil {
		return nil, err
	}
	for p.cur().Kind == Plus || p.cur().Kind == Minus {
		op := "+"
		if p.cur().Kind == Minus {
			op = "-"
		}
		pos := p.cur().Pos
		p.advance()
		right, err := p.parseMultiplicative()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right, Pos: pos}
	}
	return left, nil
}

func (p *parser) parseMultiplicative() (ast.Expr, error) {
	left, err := p.parseUnary()
	if err != nil {
		return nil, err
	}
	for p.cur().Kind == Star || p.cur().Kind == Slash || p.cur().Kind == Percent {
		var op string
		switch p.cur().Kind {
		case Star:
			op = "*"
		case Slash:
			op = "/"
		default:
			op = "%"
		}
		pos := p.cur().Pos
		p.advance()
		right, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		left = &ast.BinaryExpr{Op: op, Left: left, Right: right, Pos: pos}
	}
	return left, nil
}

func (p *parser) parseUnary() (ast.Expr, error) {
	if p.matchKind(Minus) {
		pos := p.prevPos()
		operand, err := p.parseUnary()
		if err != nil {
			return nil, err
		}
		return &ast.UnaryExpr{Op: "-", Operand: operand, Pos: pos}, nil
	}
	if p.matchKind(Plus) {
		return p.parseUnary()
	}
	return p.parsePrimary()
}

func (p *parser) parsePrimary() (ast.Expr, error) {
	tok := p.cur()
	switch tok.Kind {
	case Ident:
		// Base search has no functions: `index=web(status=500)` is a value
		// followed by a grouped term, not a call that swallows the filter.
		if !p.searchLiterals && p.callAhead(1) {
			return p.parseCall()
		}
		if p.searchLiterals {
			// A bare identifier followed by a comparison operator is a field;
			// otherwise it is a free-text search term, wildcards included.
			if p.comparisonAhead(1) {
				p.advance()
				return &ast.FieldRef{Name: tok.Text, Pos: tok.Pos}, nil
			}
			if value, ok := p.scanSearchValue(false); ok {
				return &ast.SearchTerm{Value: value, Pos: tok.Pos}, nil
			}
		}
		p.advance()
		return &ast.FieldRef{Name: tok.Text, Pos: tok.Pos}, nil

	case String:
		if p.searchLiterals {
			if p.comparisonAhead(1) {
				p.advance()
				return &ast.FieldRef{Name: tok.Text, Pos: tok.Pos}, nil
			}
			p.advance()
			return &ast.SearchTerm{Value: tok.Text, Pos: tok.Pos}, nil
		}
		p.advance()
		return &ast.StringLit{Value: tok.Text, Pos: tok.Pos}, nil

	case Number:
		value, err := strconv.ParseFloat(tok.Text, 64)
		if err != nil {
			return nil, diag.NewError(tok.Pos, "syntax", "invalid number %q", tok.Text)
		}
		p.advance()
		return &ast.NumberLit{Raw: tok.Text, Value: value, Pos: tok.Pos}, nil

	case Duration:
		p.advance()
		return &ast.DurationLit{Raw: tok.Text, Pos: tok.Pos}, nil

	case Star:
		if p.searchLiterals {
			if value, ok := p.scanSearchValue(false); ok {
				return &ast.SearchTerm{Value: value, Pos: tok.Pos}, nil
			}
		}
		p.advance()
		return &ast.WildcardLit{Pos: tok.Pos}, nil

	case LParen:
		p.advance()
		expr, err := p.parseOr()
		if err != nil {
			return nil, err
		}
		if err := p.expectKind(RParen, "')'"); err != nil {
			return nil, err
		}
		return expr, nil

	case LBracket:
		return nil, diag.NewError(tok.Pos, "unsupported-subsearch",
			"subsearches are not implemented yet")

	default:
		return nil, diag.NewError(tok.Pos, "syntax",
			"expected an expression, found %s", tok.describe())
	}
}

func (p *parser) parseCall() (ast.Expr, error) {
	name := p.cur()
	p.advance() // identifier
	p.advance() // '('

	var args []ast.Expr
	if p.cur().Kind != RParen {
		for {
			arg, err := p.parseOr()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
			if p.matchKind(Comma) {
				continue
			}
			break
		}
	}
	if err := p.expectKind(RParen, "')'"); err != nil {
		return nil, err
	}
	return &ast.CallExpr{Name: name.Text, Args: args, Pos: name.Pos}, nil
}

// --- helpers -------------------------------------------------------------

func (p *parser) cur() Token { return p.toks[p.i] }

func (p *parser) peek(n int) Token {
	idx := p.i + n
	if idx < 0 {
		idx = 0
	}
	if idx >= len(p.toks) {
		return p.toks[len(p.toks)-1]
	}
	return p.toks[idx]
}

func (p *parser) advance() {
	if p.i < len(p.toks)-1 {
		p.i++
	}
}

func (p *parser) prevPos() ast.Pos {
	if p.i == 0 {
		return p.toks[0].Pos
	}
	return p.toks[p.i-1].Pos
}

// prevTok returns the last consumed token. At the start of input there is
// none, so it returns an EOF token that never touches the current one.
func (p *parser) prevTok() Token {
	if p.i == 0 {
		return Token{Kind: EOF}
	}
	return p.toks[p.i-1]
}

func (p *parser) matchKind(k Kind) bool {
	if p.cur().Kind == k {
		p.advance()
		return true
	}
	return false
}

func (p *parser) matchKeyword(word string) bool {
	if p.cur().Kind == Ident && strings.EqualFold(p.cur().Text, word) {
		p.advance()
		return true
	}
	return false
}

func (p *parser) expectKind(k Kind, what string) error {
	if p.cur().Kind != k {
		return diag.NewError(p.cur().Pos, "syntax", "expected %s, found %s", what, p.cur().describe())
	}
	p.advance()
	return nil
}

func (p *parser) atStageEnd() bool {
	return p.cur().Kind == Pipe || p.cur().Kind == EOF
}

func (p *parser) expectStageEnd() error {
	if !p.atStageEnd() {
		return diag.NewError(p.cur().Pos, "syntax",
			"unexpected %s; expected '|' or end of query", p.cur().describe())
	}
	return nil
}

func (p *parser) startsTerm() bool {
	switch p.cur().Kind {
	case Ident:
		word := strings.ToLower(p.cur().Text)
		return word != "and" && word != "or" && word != "by" && word != "as"
	case String, Number, Duration, LParen, LBracket, Star:
		return true
	}
	return false
}

func (p *parser) comparisonAhead(n int) bool {
	switch p.peek(n).Kind {
	case Eq, Ne, Lt, Le, Gt, Ge:
		return true
	}
	return false
}

func (p *parser) callAhead(n int) bool {
	return p.peek(n).Kind == LParen
}

// adjacent reports whether token i ends exactly where token i+n starts.
func (p *parser) adjacent(i, n int) bool {
	a, b := p.peek(i), p.peek(i+n)
	return a.Pos.Offset+len(a.Text) == b.Pos.Offset
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
