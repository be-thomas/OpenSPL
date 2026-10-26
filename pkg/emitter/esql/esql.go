// Package esql emits Elastic ES|QL from the unified IR.
//
// The emitter never parses SPL. It only walks IR nodes, which keeps target
// quirks (schema-on-write fields, BUCKET, GROK) out of the parser and makes
// adding SPL2 support a front-end-only change.
package esql

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/mapping"
)

// Options configures a translation.
type Options struct {
	// Index overrides the index selection in the query.
	Index string
	// Span is the fallback timechart span when the query has no span=.
	Span string
	// Mapping supplies field/index dictionaries. Nil means mapping.Default().
	Mapping *mapping.Config
}

// Result is a translated query plus any warnings or errors.
type Result struct {
	Query       string
	Diagnostics []diag.Diagnostic
}

// Emit translates a pipeline to ES|QL.
func Emit(p *ast.Pipeline, opts Options) (*Result, error) {
	cfg := opts.Mapping
	if cfg == nil {
		cfg = mapping.Default()
	}
	e := &emitter{
		cfg:           cfg,
		indexOverride: opts.Index,
		span:          opts.Span,
		derived:       map[string]bool{},
	}
	return e.emit(p)
}

type emitter struct {
	cfg           *mapping.Config
	indexOverride string
	span          string

	lines      []string
	diags      []diag.Diagnostic
	derived    map[string]bool // columns created by bin/eval, keyed by SPL name
	fromSource string          // the source actually emitted in FROM
}

func (e *emitter) emit(p *ast.Pipeline) (*Result, error) {
	if p == nil || len(p.Stages) == 0 {
		return nil, e.errorf(ast.Pos{Line: 1, Col: 1}, "empty-query", "pipeline has no stages")
	}

	var (
		firstCleaned ast.Expr
		indexValues  []string
	)
	if s, ok := p.Stages[0].(*ast.SearchStage); ok {
		var err error
		firstCleaned, indexValues, err = extractIndex(s.Expr)
		if err != nil {
			e.record(err)
			return e.result(), err
		}
	}

	e.fromSource = e.resolveIndex(indexValues)
	e.lines = append(e.lines, "FROM "+e.fromSource)

	for i, st := range p.Stages {
		switch s := st.(type) {
		case *ast.SearchStage:
			expr := s.Expr
			if i == 0 {
				expr = firstCleaned
			} else {
				var late []string
				var err error
				expr, late, err = extractIndex(expr)
				if err != nil {
					e.record(err)
					return e.result(), err
				}
				if len(late) > 0 {
					e.warnf(s.Pos, "late-index-filter",
						"index filter after the first stage is ignored; FROM was already fixed to %s",
						e.fromSource)
				}
			}
			if err := e.emitSearch(s, expr); err != nil {
				return e.result(), err
			}
		case *ast.StatsStage:
			if err := e.emitStats(s); err != nil {
				return e.result(), err
			}
		case *ast.TimechartStage:
			if err := e.emitTimechart(s); err != nil {
				return e.result(), err
			}
		case *ast.EvalStage:
			if err := e.emitEval(s); err != nil {
				return e.result(), err
			}
		case *ast.RexStage:
			if err := e.emitRex(s); err != nil {
				return e.result(), err
			}
		case *ast.TableStage:
			e.emitTable(s)
		case *ast.FieldsStage:
			e.emitFields(s)
		case *ast.SortStage:
			e.emitSort(s)
		case *ast.HeadStage:
			e.lines = append(e.lines, fmt.Sprintf("| LIMIT %d", s.Count))
		case *ast.BinStage:
			if err := e.emitBin(s); err != nil {
				return e.result(), err
			}
		case *ast.RenameStage:
			if err := e.emitRename(s); err != nil {
				return e.result(), err
			}
		default:
			return e.result(), e.errorf(st.Position(), "unsupported-stage",
				"stage %q has no ES|QL emitter", st.StageName())
		}
	}

	diag.Sort(e.diags)
	if diag.HasErrors(e.diags) {
		return e.result(), &diag.Error{Diagnostics: e.diags}
	}
	return e.result(), nil
}

func (e *emitter) result() *Result {
	diag.Sort(e.diags)
	q := ""
	if !diag.HasErrors(e.diags) {
		q = strings.Join(e.lines, "\n")
	}
	return &Result{Query: q, Diagnostics: e.diags}
}

func (e *emitter) resolveIndex(values []string) string {
	if e.indexOverride != "" {
		return e.indexOverride
	}
	switch len(values) {
	case 0:
		if e.cfg.DefaultIndex == "" {
			return "*"
		}
		return e.cfg.DefaultIndex
	case 1:
		return e.cfg.Index(values[0])
	default:
		mapped := make([]string, 0, len(values))
		for _, v := range values {
			mapped = append(mapped, e.cfg.Index(v))
		}
		return strings.Join(mapped, ", ")
	}
}

func (e *emitter) emitSearch(s *ast.SearchStage, expr ast.Expr) error {
	var conds []string

	if s.Earliest != "" {
		c, err := e.timeCond(s.Earliest, true, s.Pos)
		if err != nil {
			return err
		}
		if c != "" {
			conds = append(conds, c)
		}
	}
	if s.Latest != "" {
		c, err := e.timeCond(s.Latest, false, s.Pos)
		if err != nil {
			return err
		}
		if c != "" {
			conds = append(conds, c)
		}
	}
	if expr != nil {
		if _, ok := expr.(*ast.WildcardLit); !ok {
			// Time bounds are joined with AND. Render a top-level OR at AND
			// precedence so the join cannot widen the time window.
			parent := 0
			if len(conds) > 0 {
				parent = opPrec("and")
			}
			if c := e.expr(expr, parent); c != "" && c != "TRUE" {
				conds = append(conds, c)
			}
		}
	}
	if len(conds) == 0 {
		return nil
	}
	e.lines = append(e.lines, "| WHERE "+strings.Join(conds, " AND "))
	return nil
}

var iso8601Re = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}([T ][0-9]{2}:[0-9]{2}(:[0-9]{2}(\.[0-9]+)?)?(Z|[+-][0-9]{2}:?[0-9]{2})?)?$`)

func (e *emitter) timeCond(raw string, lower bool, pos ast.Pos) (string, error) {
	v := strings.TrimSpace(raw)
	if v == "" || strings.EqualFold(v, "now") {
		return "", nil
	}

	op := "<="
	if lower {
		op = ">="
	}
	field := quoteIdent(e.cfg.Field("_time"))

	if strings.HasPrefix(v, "-") {
		span, err := esqlSpan(v[1:])
		if err != nil {
			return "", e.errorf(pos, "invalid-time-bound", "%v", err)
		}
		return fmt.Sprintf("%s %s NOW() - %s", field, op, span), nil
	}
	if iso8601Re.MatchString(v) {
		return fmt.Sprintf("%s %s TO_DATETIME(%s)", field, op, quoteString(v)), nil
	}
	return "", e.errorf(pos, "invalid-time-bound",
		"unsupported time bound %q; use now, a relative duration such as -24h, or an ISO 8601 timestamp", raw)
}

func (e *emitter) emitStats(s *ast.StatsStage) error {
	aggs, err := e.emitAggs(s.Aggs)
	if err != nil {
		return err
	}
	line := "| STATS " + strings.Join(aggs, ", ")
	if len(s.By) > 0 {
		fields, err := e.groupFields(s.By)
		if err != nil {
			return err
		}
		line += " BY " + fields
	}
	e.lines = append(e.lines, line)
	return nil
}

func (e *emitter) emitTimechart(s *ast.TimechartStage) error {
	span := s.Span
	if span == "" {
		span = e.span
	}
	if span == "" {
		span = e.cfg.Timechart.DefaultSpan
		e.warnf(s.Pos, "default-span",
			"timechart has no span; using the configured default %q", span)
	}
	espan, err := esqlSpan(span)
	if err != nil {
		return e.errorf(s.Pos, "invalid-span", "%v", err)
	}

	aggs, err := e.emitAggs(s.Aggs)
	if err != nil {
		return err
	}
	line := "| STATS " + strings.Join(aggs, ", ") +
		" BY BUCKET(" + quoteIdent(e.cfg.Field("_time")) + ", " + espan + ")"
	if len(s.By) > 0 {
		fields, err := e.groupFields(s.By)
		if err != nil {
			return err
		}
		line += ", " + fields
	}
	e.lines = append(e.lines, line)
	return nil
}

func (e *emitter) emitAggs(aggs []ast.Aggregate) ([]string, error) {
	out := make([]string, 0, len(aggs))
	for _, a := range aggs {
		fn := strings.ToLower(a.Func)
		var call string
		switch fn {
		case "count":
			if a.Arg == nil {
				call = "COUNT(*)"
			} else {
				call = "COUNT(" + e.expr(a.Arg, 0) + ")"
			}
		case "dc", "distinct_count", "count_distinct":
			if a.Arg == nil {
				return nil, e.errorf(a.Pos, "invalid-aggregation", "%s requires a field argument", a.Func)
			}
			call = "COUNT_DISTINCT(" + e.expr(a.Arg, 0) + ")"
		case "sum", "avg", "min", "max":
			if a.Arg == nil {
				return nil, e.errorf(a.Pos, "invalid-aggregation", "%s requires a field argument", a.Func)
			}
			call = strings.ToUpper(fn) + "(" + e.expr(a.Arg, 0) + ")"
		default:
			return nil, e.errorf(a.Pos, "unsupported-agg",
				"aggregation %q has no verified ES|QL mapping; use count, dc, sum, avg, min, or max", a.Func)
		}
		out = append(out, quoteIdent(aggregateAlias(a))+" = "+call)
	}
	return out, nil
}

func aggregateAlias(a ast.Aggregate) string {
	if a.Alias != "" {
		return a.Alias
	}
	fn := strings.ToLower(a.Func)
	if fn == "count" && a.Arg == nil {
		return "count"
	}
	if f, ok := a.Arg.(*ast.FieldRef); ok {
		return sanitizeAlias(fn + "_" + f.Name)
	}
	return fn
}

func sanitizeAlias(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteByte('_')
		}
	}
	out := strings.Trim(b.String(), "_")
	if out == "" {
		return "value"
	}
	if out[0] >= '0' && out[0] <= '9' {
		out = "_" + out
	}
	return out
}

func (e *emitter) groupFields(by []ast.Expr) (string, error) {
	parts := make([]string, 0, len(by))
	for _, ex := range by {
		f, ok := ex.(*ast.FieldRef)
		if !ok {
			return "", e.errorf(ex.Position(), "invalid-group-by",
				"group-by accepts field names only")
		}
		parts = append(parts, e.fieldName(f.Name))
	}
	return strings.Join(parts, ", "), nil
}

func (e *emitter) emitEval(s *ast.EvalStage) error {
	parts := make([]string, 0, len(s.Assignments))
	for _, a := range s.Assignments {
		parts = append(parts, quoteIdent(a.Name)+" = "+e.expr(a.Expr, 0))
		e.derived[a.Name] = true
	}
	e.lines = append(e.lines, "| EVAL "+strings.Join(parts, ", "))
	return nil
}

var namedCaptureRe = regexp.MustCompile(`\(\?<([A-Za-z_][A-Za-z0-9_]*)>([^()]*)\)`)

func (e *emitter) emitRex(s *ast.RexStage) error {
	grok, captures := regexToGrok(s.Pattern)
	if captures == 0 {
		return e.errorf(s.Pos, "rex-no-captures",
			"rex pattern must contain at least one named capture such as (?<user>\\S+)")
	}
	field := e.fieldName(s.Field)
	e.lines = append(e.lines, "| GROK "+field+" "+quoteString(grok))
	e.warnf(s.Pos, "rex-lossy",
		"rex was translated to a Grok pattern; verify the extracted values against the original search")
	return nil
}

// regexToGrok converts Splunk's `(?<name>pattern)` captures into Grok
// `%{TYPE:name}` captures. Anything that does not match a known primitive is
// mapped to DATA, which is lossy but keeps the pipeline valid; the emitter
// always warns.
func regexToGrok(pattern string) (string, int) {
	matches := namedCaptureRe.FindAllStringSubmatchIndex(pattern, -1)
	if len(matches) == 0 {
		return pattern, 0
	}
	var b strings.Builder
	last := 0
	for _, m := range matches {
		b.WriteString(pattern[last:m[0]])
		name := pattern[m[2]:m[3]]
		inner := pattern[m[4]:m[5]]
		b.WriteString("%{" + grokType(inner) + ":" + name + "}")
		last = m[1]
	}
	b.WriteString(pattern[last:])
	return b.String(), len(matches)
}

func grokType(pattern string) string {
	switch pattern {
	case `\S+`, `[^\s]+`, `[^ ]+`:
		return "NOTSPACE"
	case `\w+`:
		return "WORD"
	case `\d+`, `[0-9]+`:
		return "NUMBER"
	case `[A-Za-z]+`:
		return "WORD"
	case `.+`, `.*`, `.+?`, `.*?`:
		return "DATA"
	}
	return "DATA"
}

func (e *emitter) emitTable(s *ast.TableStage) {
	if s.Wildcard && len(s.Fields) == 0 {
		return // KEEP * is a no-op; omit it to keep queries readable.
	}
	fields := make([]string, 0, len(s.Fields))
	for _, f := range s.Fields {
		fields = append(fields, e.fieldName(f))
	}
	if len(fields) == 0 {
		return
	}
	e.lines = append(e.lines, "| KEEP "+strings.Join(fields, ", "))
}

func (e *emitter) emitFields(s *ast.FieldsStage) {
	if len(s.Include) > 0 {
		fields := make([]string, 0, len(s.Include))
		for _, f := range s.Include {
			fields = append(fields, e.fieldName(f))
		}
		e.lines = append(e.lines, "| KEEP "+strings.Join(fields, ", "))
	}
	if len(s.Exclude) > 0 {
		fields := make([]string, 0, len(s.Exclude))
		for _, f := range s.Exclude {
			fields = append(fields, e.fieldName(f))
		}
		e.lines = append(e.lines, "| DROP "+strings.Join(fields, ", "))
	}
}

func (e *emitter) emitSort(s *ast.SortStage) {
	if len(s.Fields) > 0 {
		parts := make([]string, 0, len(s.Fields))
		for _, f := range s.Fields {
			dir := "ASC"
			if f.Desc {
				dir = "DESC"
			}
			parts = append(parts, e.fieldName(f.Name)+" "+dir)
		}
		e.lines = append(e.lines, "| SORT "+strings.Join(parts, ", "))
	}
	if s.Limit > 0 {
		e.lines = append(e.lines, fmt.Sprintf("| LIMIT %d", s.Limit))
	}
}

func (e *emitter) emitBin(s *ast.BinStage) error {
	espan, err := esqlSpan(s.Span)
	if err != nil {
		return e.errorf(s.Pos, "invalid-span", "%v", err)
	}
	source := e.fieldName(s.Field)
	e.lines = append(e.lines, "| EVAL "+quoteIdent(s.Field)+" = BUCKET("+source+", "+espan+")")
	e.derived[s.Field] = true
	e.info(s.Pos, "bin-bucket",
		"bin is emitted as a derived column %q; downstream stages reference the bucketed value", s.Field)
	return nil
}

// emitRename translates `rename from AS to, ...` into one ES|QL RENAME
// command. ES|QL has no wildcard form and applies all pairs to the same input
// columns, so a wildcard, a field renamed twice, or a pair whose target is
// another pair's source is refused instead of silently changing meaning.
func (e *emitter) emitRename(s *ast.RenameStage) error {
	froms := make(map[string]bool, len(s.Renames))
	for _, r := range s.Renames {
		if strings.Contains(r.From, "*") {
			return e.errorf(r.FromPos, "rename-wildcard",
				"ES|QL RENAME has no wildcard form; %q names several columns", r.From)
		}
		if strings.Contains(r.To, "*") {
			return e.errorf(r.ToPos, "rename-wildcard",
				"ES|QL RENAME has no wildcard form; %q names several columns", r.To)
		}
	}
	for _, r := range s.Renames {
		if froms[r.From] {
			return e.errorf(r.FromPos, "rename-duplicate",
				"field %q is renamed more than once in one command", r.From)
		}
		froms[r.From] = true
	}
	for _, r := range s.Renames {
		if r.From == r.To {
			continue // Splunk treats this as a no-op.
		}
		if froms[r.To] {
			return e.errorf(r.ToPos, "rename-conflict",
				"field %q is also renamed in the same command; overlapping pairs are refused rather than applied in an unspecified order", r.To)
		}
	}

	parts := make([]string, 0, len(s.Renames))
	for _, r := range s.Renames {
		if r.From == r.To {
			continue
		}
		// The source is an existing column, so it goes through the mapping. The
		// target is the name the query chooses, so it is emitted as written;
		// mapping it would rename to a column the query never asked for.
		parts = append(parts, e.fieldName(r.From)+" AS "+quoteIdent(r.To))
	}
	if len(parts) == 0 {
		return nil
	}
	e.lines = append(e.lines, "| RENAME "+strings.Join(parts, ", "))
	return nil
}

func (e *emitter) fieldName(name string) string {
	if e.derived[name] {
		return quoteIdent(name)
	}
	return quoteIdent(e.cfg.Field(name))
}

// --- expressions ---------------------------------------------------------

func (e *emitter) expr(x ast.Expr, parent int) string {
	switch v := x.(type) {
	case *ast.FieldRef:
		return e.fieldName(v.Name)
	case *ast.StringLit:
		return quoteString(v.Value)
	case *ast.NumberLit:
		return v.Raw
	case *ast.DurationLit:
		if span, err := esqlSpan(v.Raw); err == nil {
			return span
		}
		return quoteString(v.Raw)
	case *ast.SearchTerm:
		term := strings.TrimSuffix(v.Value, "*")
		e.warnf(v.Pos, "fulltext-approximation",
			"free-text term %q is translated to MATCH on %q; verify analyzer and tokenization", v.Value, e.cfg.TextField)
		return "MATCH(" + quoteIdent(e.cfg.TextField) + ", " + quoteString(term) + ")"
	case *ast.WildcardLit:
		return "TRUE"
	case *ast.UnaryExpr:
		p := 6
		inner := e.expr(v.Operand, p)
		s := v.Op + inner
		if v.Op == "not" {
			s = "NOT " + inner
		}
		if p < parent {
			return "(" + s + ")"
		}
		return s
	case *ast.BinaryExpr:
		return e.emitBinary(v, parent)
	case *ast.CallExpr:
		return e.emitCall(v)
	default:
		return "NULL"
	}
}

func (e *emitter) emitBinary(b *ast.BinaryExpr, parent int) string {
	p := opPrec(b.Op)
	var s string
	if isComparison(b.Op) {
		s = e.emitComparison(b)
	} else {
		op := b.Op
		switch op {
		case "and":
			op = "AND"
		case "or":
			op = "OR"
		}
		rightParent := p
		if !isAssociative(b.Op) {
			rightParent = p + 1
		}
		s = e.expr(b.Left, p) + " " + op + " " + e.expr(b.Right, rightParent)
	}
	if p < parent {
		return "(" + s + ")"
	}
	return s
}

func (e *emitter) emitComparison(b *ast.BinaryExpr) string {
	left := e.expr(b.Left, 4)

	if _, ok := b.Right.(*ast.WildcardLit); ok {
		switch b.Op {
		case "=":
			return left + " IS NOT NULL"
		case "!=":
			return left + " IS NULL"
		}
	}
	if s, ok := b.Right.(*ast.StringLit); ok && strings.ContainsAny(s.Value, "*?") {
		switch b.Op {
		case "=":
			return left + " LIKE " + quoteString(s.Value)
		case "!=":
			return "NOT (" + left + " LIKE " + quoteString(s.Value) + ")"
		}
	}

	op := b.Op
	if op == "=" {
		op = "=="
	}
	return left + " " + op + " " + e.expr(b.Right, 4)
}

func (e *emitter) emitCall(c *ast.CallExpr) string {
	name := strings.ToLower(c.Name)
	args := make([]string, 0, len(c.Args))
	for _, a := range c.Args {
		args = append(args, e.expr(a, 0))
	}
	joined := strings.Join(args, ", ")

	switch name {
	case "if":
		if len(args) == 3 {
			return "CASE(" + joined + ")"
		}
	case "isnull":
		if len(args) == 1 {
			return args[0] + " IS NULL"
		}
	case "isnotnull", "notnull":
		if len(args) == 1 {
			return args[0] + " IS NOT NULL"
		}
	case "now":
		if len(args) == 0 {
			return "NOW()"
		}
	case "tonumber":
		return "TO_DOUBLE(" + joined + ")"
	case "tostring":
		return "TO_STRING(" + joined + ")"
	case "lower", "upper", "length", "abs", "round", "coalesce", "concat",
		"trim", "replace", "substring", "ceil", "floor", "sqrt", "pow":
		return strings.ToUpper(name) + "(" + joined + ")"
	}

	e.warnf(c.Pos, "unmapped-function",
		"function %q has no verified ES|QL mapping; passed through unchanged", c.Name)
	return strings.ToUpper(c.Name) + "(" + joined + ")"
}

func opPrec(op string) int {
	switch op {
	case "or":
		return 1
	case "and":
		return 2
	case "=", "!=", "<", "<=", ">", ">=":
		return 3
	case "+", "-":
		return 4
	case "*", "/", "%":
		return 5
	}
	return 7
}

func isComparison(op string) bool {
	switch op {
	case "=", "!=", "<", "<=", ">", ">=":
		return true
	}
	return false
}

func isAssociative(op string) bool {
	switch op {
	case "and", "or", "+", "*":
		return true
	}
	return false
}

// --- index extraction ----------------------------------------------------

// extractIndex lifts index= filters out of a search expression. A pure
// disjunction of any arity becomes a multi-index FROM; an OR that mixes index
// filters with row filters is refused, because ES|QL cannot express "scan
// these indexes, or keep rows matching another condition" without changing
// which rows are scanned.
func extractIndex(e ast.Expr) (ast.Expr, []string, error) {
	if e == nil {
		return nil, nil, nil
	}
	b, ok := e.(*ast.BinaryExpr)
	if !ok {
		return e, nil, nil
	}

	if b.Op == "and" {
		left, leftIdx, err := extractIndex(b.Left)
		if err != nil {
			return nil, nil, err
		}
		right, rightIdx, err := extractIndex(b.Right)
		if err != nil {
			return nil, nil, err
		}
		idx := append(leftIdx, rightIdx...)
		switch {
		case left == nil:
			return right, idx, nil
		case right == nil:
			return left, idx, nil
		default:
			return &ast.BinaryExpr{Op: "and", Left: left, Right: right, Pos: b.Pos}, idx, nil
		}
	}
	if b.Op == "or" {
		values, pure, err := indexOrValues(b)
		if err != nil {
			return nil, nil, err
		}
		if pure {
			return nil, values, nil
		}
		return b, nil, nil
	}
	if v, ok := indexEqValue(b); ok {
		return nil, []string{v}, nil
	}
	return b, nil, nil
}

// indexOrValues flattens an OR tree into index values. It reports pure=false
// when the OR contains no index filter at all; when it mixes one with other
// filters, it returns a positioned refusal.
func indexOrValues(e ast.Expr) ([]string, bool, error) {
	if v, ok := indexEqValue(e); ok {
		return []string{v}, true, nil
	}
	b, ok := e.(*ast.BinaryExpr)
	if !ok || b.Op != "or" {
		return nil, false, nil
	}

	left, leftPure, err := indexOrValues(b.Left)
	if err != nil {
		return nil, false, err
	}
	right, rightPure, err := indexOrValues(b.Right)
	if err != nil {
		return nil, false, err
	}
	if leftPure && rightPure {
		return append(left, right...), true, nil
	}
	if containsIndexFilter(b) {
		return nil, false, diag.NewError(b.Pos, "mixed-index-filter",
			"an OR cannot combine index= filters with other filters; ES|QL selects indexes in FROM, so use AND or put index= in a separate search")
	}
	return nil, false, nil
}

func containsIndexFilter(e ast.Expr) bool {
	switch v := e.(type) {
	case *ast.BinaryExpr:
		if _, ok := indexEqValue(v); ok {
			return true
		}
		return containsIndexFilter(v.Left) || containsIndexFilter(v.Right)
	case *ast.UnaryExpr:
		return containsIndexFilter(v.Operand)
	}
	return false
}

func indexEqValue(e ast.Expr) (string, bool) {
	b, ok := e.(*ast.BinaryExpr)
	if !ok || b.Op != "=" {
		return "", false
	}
	f, ok := b.Left.(*ast.FieldRef)
	if !ok || !strings.EqualFold(f.Name, "index") {
		return "", false
	}
	s, ok := b.Right.(*ast.StringLit)
	if !ok {
		return "", false
	}
	return s.Value, true
}

// --- spans, quoting, diagnostics ----------------------------------------

var spanRe = regexp.MustCompile(`^(\d+)(ms|s|m|h|d|w|mon|y)$`)

var spanUnits = map[string]string{
	"ms":  "milliseconds",
	"s":   "seconds",
	"m":   "minutes",
	"h":   "hours",
	"d":   "days",
	"w":   "weeks",
	"mon": "months",
	"y":   "years",
}

func esqlSpan(raw string) (string, error) {
	m := spanRe.FindStringSubmatch(strings.TrimSpace(raw))
	if m == nil {
		return "", fmt.Errorf("invalid span %q; use a duration such as 1m, 5m, 1h, or 1d", raw)
	}
	n, err := strconv.Atoi(m[1])
	if err != nil || n <= 0 {
		return "", fmt.Errorf("invalid span %q; the duration must be positive", raw)
	}
	name := spanUnits[m[2]]
	if n == 1 {
		name = strings.TrimSuffix(name, "s")
	}
	return fmt.Sprintf("%d %s", n, name), nil
}

func quoteString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for _, r := range s {
		switch r {
		case '\\':
			b.WriteString(`\\`)
		case '"':
			b.WriteString(`\"`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}

func quoteIdent(s string) string {
	if isSimpleIdent(s) {
		return s
	}
	return "`" + strings.ReplaceAll(s, "`", "``") + "`"
}

func isSimpleIdent(s string) bool {
	if s == "" {
		return false
	}
	for i, r := range s {
		ok := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_' ||
			(i > 0 && (r >= '0' && r <= '9' || r == '@'))
		if !ok {
			return false
		}
	}
	return true
}

func (e *emitter) errorf(pos ast.Pos, code, format string, args ...any) error {
	d := diag.Diagnostic{
		Severity: diag.SeverityError,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
		Pos:      pos,
	}
	e.diags = append(e.diags, d)
	return &diag.Error{Diagnostics: []diag.Diagnostic{d}}
}

// record adds the diagnostics carried by err to the emitter's list, so
// result() sees the failure and withholds the partial query. Errors that reach
// the emitter without diagnostics still get a positioned internal entry.
func (e *emitter) record(err error) {
	if ds, ok := diag.FromError(err); ok {
		e.diags = append(e.diags, ds...)
		return
	}
	e.diags = append(e.diags, diag.Diagnostic{
		Severity: diag.SeverityError,
		Code:     "internal",
		Message:  err.Error(),
		Pos:      ast.Pos{Line: 1, Col: 1},
	})
}

func (e *emitter) warnf(pos ast.Pos, code, format string, args ...any) {
	e.diags = append(e.diags, diag.Diagnostic{
		Severity: diag.SeverityWarning,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
		Pos:      pos,
	})
}

func (e *emitter) info(pos ast.Pos, code, format string, args ...any) {
	e.diags = append(e.diags, diag.Diagnostic{
		Severity: diag.SeverityInfo,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
		Pos:      pos,
	})
}
