// Package promql emits Prometheus PromQL from the unified IR.
//
// SPL and PromQL model different worlds: SPL returns rows (including raw log
// lines), PromQL returns time series. This emitter therefore performs a
// semantic check and refuses pipelines it cannot preserve, instead of
// producing a query that silently means something else.
package promql

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/mapping"
)

// Options configures a translation.
type Options struct {
	Metric  string
	Span    string
	Mapping *mapping.Config
}

// Result is a translated query plus diagnostics.
type Result struct {
	Query       string
	Diagnostics []diag.Diagnostic
}

// Emit translates a pipeline to a PromQL expression.
func Emit(p *ast.Pipeline, opts Options) (*Result, error) {
	cfg := opts.Mapping
	if cfg == nil {
		cfg = mapping.Default()
	}
	e := &emitter{cfg: cfg, metric: opts.Metric, span: opts.Span}
	return e.emit(p)
}

type emitter struct {
	cfg    *mapping.Config
	metric string
	span   string
	diags  []diag.Diagnostic
}

func (e *emitter) emit(p *ast.Pipeline) (*Result, error) {
	if p == nil || len(p.Stages) == 0 {
		return e.abort(ast.Pos{Line: 1, Col: 1}, "empty-query", "pipeline has no stages")
	}

	var (
		filters   []ast.Expr
		aggs      []ast.Aggregate
		by        []ast.Expr
		span      string
		sortStage *ast.SortStage
		headStage *ast.HeadStage
		seenAgg   bool
		aggPos    ast.Pos
	)

	for _, st := range p.Stages {
		switch s := st.(type) {
		case *ast.SearchStage:
			if seenAgg {
				return e.abort(s.Pos, "promql-search-after-agg",
					"filters after an aggregation operate on result rows and cannot be expressed in PromQL")
			}
			if s.Expr != nil {
				filters = append(filters, s.Expr)
			}
			if s.Earliest != "" || s.Latest != "" {
				e.info(s.Pos, "time-bounds-external",
					"earliest/latest are evaluation-window bounds; set them on the Prometheus query range instead of the expression")
			}
		case *ast.StatsStage:
			if seenAgg {
				return e.abort(s.Pos, "promql-multiple-aggs",
					"PromQL expresses a single aggregation; pipeline has more than one")
			}
			aggs, by, seenAgg, aggPos = s.Aggs, s.By, true, s.Pos
			e.warnf(s.Pos, "promql-window",
				"PromQL evaluates over a range window, not the whole search range; using the configured span")
		case *ast.TimechartStage:
			if seenAgg {
				return e.abort(s.Pos, "promql-multiple-aggs",
					"PromQL expresses a single aggregation; pipeline has more than one")
			}
			span = s.Span
			aggs, by, seenAgg, aggPos = s.Aggs, s.By, true, s.Pos
		case *ast.SortStage:
			if !seenAgg {
				e.warnf(s.Pos, "promql-sort-ignored",
					"sort before an aggregation is dropped; it cannot change the aggregated result")
				continue
			}
			sortStage = s
		case *ast.HeadStage:
			if !seenAgg {
				return e.abort(s.Pos, "promql-head-before-agg",
					"head before an aggregation changes which rows are aggregated and cannot be expressed in PromQL")
			}
			headStage = s
		default:
			return e.abort(st.Position(), "promql-unsupported-stage",
				"stage %q cannot be expressed in PromQL; PromQL targets time-series aggregations only", st.StageName())
		}
	}

	if !seenAgg {
		return e.abort(p.Stages[len(p.Stages)-1].Position(), "promql-no-aggregation",
			"PromQL can only express aggregations over time; add timechart or stats, or use --target=esql")
	}
	if len(aggs) == 0 {
		return e.abort(aggPos, "promql-no-aggregation", "the aggregation stage has no aggregations")
	}
	if len(aggs) > 1 {
		return e.abort(aggPos, "promql-multiple-aggs",
			"PromQL cannot combine %d aggregations in one expression; use one count aggregation", len(aggs))
	}
	agg := aggs[0]
	if fn := strings.ToLower(agg.Func); fn != "count" {
		return e.abort(agg.Pos, "promql-unsupported-agg",
			"aggregation %q has no faithful PromQL mapping; only count is supported because it maps to a counter increase", agg.Func)
	}
	if agg.Arg != nil {
		e.warnf(agg.Pos, "promql-count-field",
			"count(field) counts non-null values in SPL; the PromQL translation counts all series increments")
	}

	matchers, err := e.matchers(filters)
	if err != nil {
		return e.result(""), err
	}

	labels, err := e.groupLabels(by)
	if err != nil {
		return e.result(""), err
	}

	if span == "" {
		span = e.span
	}
	if span == "" {
		span = e.cfg.Timechart.DefaultSpan
		e.warnf(aggPos, "default-span",
			"no span was specified; using the configured default %q", span)
	}
	if err := validateRange(span); err != nil {
		return e.abort(aggPos, "invalid-span", "%v", err)
	}

	metric := e.metric
	if metric == "" {
		metric = e.cfg.Prometheus.Metric
	}
	fn := e.cfg.Prometheus.Function
	if fn != "increase" && fn != "rate" {
		return e.abort(aggPos, "invalid-mapping",
			"prometheus.function must be \"increase\" or \"rate\", got %q", fn)
	}

	selector := metric
	if len(matchers) > 0 {
		parts := make([]string, 0, len(matchers))
		for _, m := range matchers {
			parts = append(parts, fmt.Sprintf("%s%s%s", m.label, m.op, quote(m.value)))
		}
		selector += "{" + strings.Join(parts, ", ") + "}"
	}

	inner := fmt.Sprintf("%s(%s[%s])", fn, selector, span)
	query := "sum(" + inner + ")"
	if len(labels) > 0 {
		query = "sum by (" + strings.Join(labels, ", ") + ") (" + inner + ")"
	}

	query = e.applyOrdering(query, sortStage, headStage)

	diag.Sort(e.diags)
	if diag.HasErrors(e.diags) {
		return e.result(""), &diag.Error{Diagnostics: e.diags}
	}
	return e.result(query), nil
}

func (e *emitter) applyOrdering(query string, sortStage *ast.SortStage, headStage *ast.HeadStage) string {
	n := 0
	headGiven := false
	if headStage != nil {
		n = headStage.Count
		headGiven = true
		if headStage.CountSet && headStage.Count == 0 {
			e.warnf(headStage.Pos, "promql-head-ignored",
				"head 0 would drop every row and cannot be expressed in PromQL; it was ignored")
		}
	}

	if sortStage != nil {
		if len(sortStage.Fields) > 1 {
			e.warnf(sortStage.Pos, "promql-sort-fields",
				"PromQL can order by a single series value; additional sort fields were ignored")
		}
		if !headGiven {
			n = sortStage.Limit
		}
		if n > 0 && len(sortStage.Fields) > 0 {
			if sortStage.Fields[0].Desc {
				return fmt.Sprintf("topk(%d, %s)", n, query)
			}
			return fmt.Sprintf("bottomk(%d, %s)", n, query)
		}
		if n == 0 && len(sortStage.Fields) > 0 {
			if sortStage.Fields[0].Desc {
				return "sort_desc(" + query + ")"
			}
			return "sort(" + query + ")"
		}
		if len(sortStage.Fields) == 0 {
			e.warnf(sortStage.Pos, "promql-sort-limit-ignored",
				"a sort limit without a sort field cannot be expressed in PromQL; it was ignored")
		}
	}
	if n > 0 && sortStage == nil {
		e.warnf(headStage.Pos, "promql-head-ignored",
			"head without sort is an arbitrary truncation of rows and cannot be expressed in PromQL; it was ignored")
	}
	return query
}

// --- filters -------------------------------------------------------------

type matcher struct {
	label string
	op    string
	value string
	pos   ast.Pos
}

func (e *emitter) matchers(filters []ast.Expr) ([]matcher, error) {
	var collected []matcher
	for _, f := range filters {
		if err := e.collectMatchers(f, &collected); err != nil {
			return nil, err
		}
	}
	return mergeMatchers(collected)
}

func (e *emitter) collectMatchers(expr ast.Expr, out *[]matcher) error {
	switch v := expr.(type) {
	case nil:
		return nil
	case *ast.WildcardLit:
		return nil
	case *ast.BinaryExpr:
		switch v.Op {
		case "and":
			if err := e.collectMatchers(v.Left, out); err != nil {
				return err
			}
			return e.collectMatchers(v.Right, out)
		case "=", "!=":
			m, err := e.matcher(v)
			if err != nil {
				return err
			}
			*out = append(*out, m)
			return nil
		default:
			return e.fail(v.Pos, "promql-unsupported-filter",
				"operator %q cannot be expressed as a PromQL label matcher; only '=' and '!=' are supported", v.Op)
		}
	case *ast.UnaryExpr:
		if v.Op == "not" {
			b, ok := v.Operand.(*ast.BinaryExpr)
			if ok && (b.Op == "=" || b.Op == "!=") {
				flipped := *b
				if b.Op == "=" {
					flipped.Op = "!="
				} else {
					flipped.Op = "="
				}
				m, err := e.matcher(&flipped)
				if err != nil {
					return err
				}
				*out = append(*out, m)
				return nil
			}
		}
		return e.fail(v.Pos, "promql-unsupported-filter",
			"this negation cannot be expressed as a PromQL label matcher")
	case *ast.SearchTerm:
		return e.fail(v.Pos, "promql-fulltext",
			"free-text term %q cannot be expressed as a PromQL label matcher; add an explicit field filter", v.Value)
	default:
		return e.fail(expr.Position(), "promql-unsupported-filter",
			"this filter cannot be expressed as a PromQL label matcher")
	}
}

func (e *emitter) matcher(b *ast.BinaryExpr) (matcher, error) {
	field, ok := b.Left.(*ast.FieldRef)
	if !ok {
		return matcher{}, e.fail(b.Pos, "promql-unsupported-filter",
			"the left side of a PromQL label matcher must be a field")
	}

	label := e.cfg.Label(field.Name)
	if strings.EqualFold(field.Name, "index") {
		label = e.cfg.Prometheus.IndexLabel
		if label == "" {
			label = "job"
		}
	}

	var value string
	wildcard := false
	switch v := b.Right.(type) {
	case *ast.StringLit:
		value = v.Value
		wildcard = strings.ContainsAny(value, "*?")
	case *ast.NumberLit:
		value = v.Raw
	default:
		return matcher{}, e.fail(b.Pos, "promql-unsupported-filter",
			"PromQL label values must be string or numeric literals")
	}

	op := b.Op
	if wildcard {
		value = wildcardToRegex(value)
		if op == "=" {
			op = "=~"
		} else {
			op = "!~"
		}
	}
	return matcher{label: label, op: op, value: value, pos: b.Pos}, nil
}

// mergeMatchers collapses duplicate label matchers. PromQL forbids the same
// label twice in one selector, while SPL happily repeats a field.
func mergeMatchers(in []matcher) ([]matcher, error) {
	byLabel := map[string][]matcher{}
	var order []string
	for _, m := range in {
		if _, ok := byLabel[m.label]; !ok {
			order = append(order, m.label)
		}
		byLabel[m.label] = append(byLabel[m.label], m)
	}
	sort.Strings(order)

	out := make([]matcher, 0, len(order))
	for _, label := range order {
		ms := byLabel[label]
		if len(ms) == 1 {
			out = append(out, ms[0])
			continue
		}
		positive := true
		values := make([]string, 0, len(ms))
		for _, m := range ms {
			if m.op != "=" && m.op != "=~" {
				positive = false
			}
			if m.op == "=" {
				values = append(values, regexp.QuoteMeta(m.value))
			} else {
				values = append(values, m.value)
			}
		}
		if positive {
			values = dedupe(values)
			if len(values) == 1 {
				op := "=~"
				if allOp(ms, "=") {
					op = "="
				}
				out = append(out, matcher{label: label, op: op, value: ms[0].value})
				continue
			}
			// SPL filters are a conjunction. Two different positive values can
			// never match, so joining them with "|" would turn AND into OR and
			// silently widen the result set.
			return nil, diag.NewError(ms[0].pos, "promql-conflicting-filters",
				"label %q is filtered to different values; SPL treats repeated filters as a conjunction, which PromQL cannot express", label)
		}
		negative := true
		values = values[:0]
		for _, m := range ms {
			if m.op != "!=" && m.op != "!~" {
				negative = false
			}
			if m.op == "!=" {
				values = append(values, regexp.QuoteMeta(m.value))
			} else {
				values = append(values, m.value)
			}
		}
		if negative {
			values = dedupe(values)
			if len(values) == 1 && allOp(ms, "!=") {
				out = append(out, matcher{label: label, op: "!=", value: ms[0].value})
				continue
			}
			out = append(out, matcher{label: label, op: "!~", value: strings.Join(values, "|")})
			continue
		}
		return nil, diag.NewError(ms[0].pos, "promql-conflicting-filters",
			"label %q has conflicting matchers; combine them into a single filter", label)
	}
	return out, nil
}

// dedupe removes repeated matcher values while preserving first-seen order.
// SPL repeats filters freely; the label value set must not repeat.
func dedupe(values []string) []string {
	seen := make(map[string]bool, len(values))
	out := values[:0]
	for _, v := range values {
		if seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out
}

func allOp(ms []matcher, op string) bool {
	for _, m := range ms {
		if m.op != op {
			return false
		}
	}
	return true
}

func (e *emitter) groupLabels(by []ast.Expr) ([]string, error) {
	var labels []string
	seen := map[string]bool{}
	for _, ex := range by {
		f, ok := ex.(*ast.FieldRef)
		if !ok {
			return nil, e.fail(ex.Position(), "promql-unsupported-group",
				"PromQL grouping accepts field names only")
		}
		label := e.cfg.Label(f.Name)
		if !seen[label] {
			seen[label] = true
			labels = append(labels, label)
		}
	}
	return labels, nil
}

// --- helpers -------------------------------------------------------------

var promRangeRe = regexp.MustCompile(`^(\d+)(ms|s|m|h|d|w|y)$`)

func validateRange(span string) error {
	m := promRangeRe.FindStringSubmatch(strings.TrimSpace(span))
	if m == nil || m[1] == "0" {
		return fmt.Errorf("invalid PromQL range %q; use a duration such as 1m, 5m, 1h, or 1d", span)
	}
	return nil
}

func wildcardToRegex(s string) string {
	var b strings.Builder
	b.WriteString("^")
	for _, r := range s {
		switch r {
		case '*':
			b.WriteString(".*")
		case '?':
			b.WriteString(".")
		default:
			if strings.ContainsRune(`.+()|[]{}^$\`, r) {
				b.WriteByte('\\')
			}
			b.WriteRune(r)
		}
	}
	b.WriteString("$")
	return b.String()
}

func quote(s string) string {
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
		default:
			b.WriteRune(r)
		}
	}
	b.WriteByte('"')
	return b.String()
}

func (e *emitter) result(q string) *Result {
	diag.Sort(e.diags)
	return &Result{Query: q, Diagnostics: e.diags}
}

func (e *emitter) abort(pos ast.Pos, code, format string, args ...any) (*Result, error) {
	return e.result(""), e.fail(pos, code, format, args...)
}

func (e *emitter) fail(pos ast.Pos, code, format string, args ...any) error {
	d := diag.Diagnostic{
		Severity: diag.SeverityError,
		Code:     code,
		Message:  fmt.Sprintf(format, args...),
		Pos:      pos,
	}
	e.diags = append(e.diags, d)
	return &diag.Error{Diagnostics: []diag.Diagnostic{d}}
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
