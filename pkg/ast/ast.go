// Package ast defines OpenSPL's unified intermediate representation (IR).
//
// Every front-end (classic SPL today, SPL2 next) lowers into these nodes and
// every back-end (ES|QL, PromQL) reads only these nodes. Front-ends never know
// about a target; back-ends never know about source syntax. Adding a language
// version or a target is therefore an additive change, not a rewrite.
package ast

// Pos is a 1-based line and byte column plus a byte offset into the source
// query. Columns count bytes, not runes, so a multibyte character advances the
// column by its UTF-8 length. It is carried on every node so diagnostics can
// point at the offending token.
type Pos struct {
	Line   int `json:"line"`
	Col    int `json:"col"`
	Offset int `json:"offset"`
}

// Pipeline is an ordered list of stages, mirroring SPL's `|` pipeline.
type Pipeline struct {
	Stages []Stage `json:"stages"`
}

// Stage is one `|`-delimited command.
type Stage interface {
	StageName() string
	Position() Pos
}

// SearchStage filters rows. In classic SPL the leading search is implicit;
// `| search` and `| where` produce the same node (with different expression
// semantics resolved by the parser: search values are literals, where
// expressions may compare fields).
type SearchStage struct {
	Expr     Expr   `json:"expr,omitempty"`
	Earliest string `json:"earliest,omitempty"`
	Latest   string `json:"latest,omitempty"`
	Pos      Pos    `json:"pos"`
}

func (*SearchStage) StageName() string { return "search" }
func (s *SearchStage) Position() Pos   { return s.Pos }

// StatsStage groups and aggregates rows.
type StatsStage struct {
	Aggs []Aggregate `json:"aggs"`
	By   []Expr      `json:"by,omitempty"`
	Pos  Pos         `json:"pos"`
}

func (*StatsStage) StageName() string { return "stats" }
func (s *StatsStage) Position() Pos   { return s.Pos }

// TimechartStage is stats over time buckets. Span is the raw SPL span
// (for example "5m"); the emitter resolves defaults.
type TimechartStage struct {
	Aggs []Aggregate `json:"aggs"`
	By   []Expr      `json:"by,omitempty"`
	Span string      `json:"span,omitempty"`
	Pos  Pos         `json:"pos"`
}

func (*TimechartStage) StageName() string { return "timechart" }
func (s *TimechartStage) Position() Pos   { return s.Pos }

// Aggregate is one aggregation expression, optionally aliased.
type Aggregate struct {
	Func  string `json:"func"`
	Arg   Expr   `json:"arg,omitempty"`
	Alias string `json:"alias,omitempty"`
	Pos   Pos    `json:"pos"`
}

// EvalStage adds or replaces fields.
type EvalStage struct {
	Assignments []Assignment `json:"assignments"`
	Pos         Pos          `json:"pos"`
}

func (*EvalStage) StageName() string { return "eval" }
func (s *EvalStage) Position() Pos   { return s.Pos }

// Assignment is one `name = expression` pair inside eval.
type Assignment struct {
	Name string `json:"name"`
	Expr Expr   `json:"expr"`
	Pos  Pos    `json:"pos"`
}

// RexStage extracts fields with a regular expression.
type RexStage struct {
	Field   string `json:"field"`
	Pattern string `json:"pattern"`
	Pos     Pos    `json:"pos"`
}

func (*RexStage) StageName() string { return "rex" }
func (s *RexStage) Position() Pos   { return s.Pos }

// TableStage selects and orders columns.
type TableStage struct {
	Fields   []string `json:"fields,omitempty"`
	Wildcard bool     `json:"wildcard,omitempty"`
	Pos      Pos      `json:"pos"`
}

func (*TableStage) StageName() string { return "table" }
func (s *TableStage) Position() Pos   { return s.Pos }

// FieldsStage keeps (Include) or removes (Exclude) fields.
type FieldsStage struct {
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
	Pos     Pos      `json:"pos"`
}

func (*FieldsStage) StageName() string { return "fields" }
func (s *FieldsStage) Position() Pos   { return s.Pos }

// SortStage orders rows. Limit, when non-zero, is emitted as a trailing LIMIT.
type SortStage struct {
	Fields []SortField `json:"fields,omitempty"`
	Limit  int         `json:"limit,omitempty"`
	Pos    Pos         `json:"pos"`
}

func (*SortStage) StageName() string { return "sort" }
func (s *SortStage) Position() Pos   { return s.Pos }

// SortField is one sort key.
type SortField struct {
	Name string `json:"name"`
	Desc bool   `json:"desc,omitempty"`
}

// HeadStage caps the number of rows. CountSet distinguishes `head 0` from an
// omitted count, which defaults to Splunk's 10.
type HeadStage struct {
	Count    int  `json:"count"`
	CountSet bool `json:"count_set,omitempty"`
	Pos      Pos  `json:"pos"`
}

func (*HeadStage) StageName() string { return "head" }
func (s *HeadStage) Position() Pos   { return s.Pos }

// BinStage buckets a field into fixed-width spans.
type BinStage struct {
	Field string `json:"field"`
	Span  string `json:"span"`
	Pos   Pos    `json:"pos"`
}

func (*BinStage) StageName() string { return "bin" }
func (s *BinStage) Position() Pos   { return s.Pos }

// RenameStage renames one or more fields. Splunk applies every pair in a
// single command, so the pairs travel together rather than as separate stages.
type RenameStage struct {
	Renames []Rename `json:"renames"`
	Pos     Pos      `json:"pos"`
}

func (*RenameStage) StageName() string { return "rename" }
func (s *RenameStage) Position() Pos   { return s.Pos }

// Rename is one `from AS to` pair. The two positions point at the field
// tokens, so a refusal can name the field that caused it.
type Rename struct {
	From    string `json:"from"`
	FromPos Pos    `json:"from_pos"`
	To      string `json:"to"`
	ToPos   Pos    `json:"to_pos"`
}

// Expr is a scalar expression in the IR.
type Expr interface {
	exprNode()
	Position() Pos
}

// FieldRef references a column.
type FieldRef struct {
	Name string `json:"name"`
	Pos  Pos    `json:"pos"`
}

func (*FieldRef) exprNode()       {}
func (e *FieldRef) Position() Pos { return e.Pos }

// StringLit is a quoted (or, in search context, unquoted) string value.
type StringLit struct {
	Value string `json:"value"`
	Pos   Pos    `json:"pos"`
}

func (*StringLit) exprNode()       {}
func (e *StringLit) Position() Pos { return e.Pos }

// NumberLit is a numeric literal; Raw preserves the source spelling.
type NumberLit struct {
	Raw   string  `json:"raw"`
	Value float64 `json:"value"`
	Pos   Pos     `json:"pos"`
}

func (*NumberLit) exprNode()       {}
func (e *NumberLit) Position() Pos { return e.Pos }

// DurationLit is a time span such as 5m or 24h.
type DurationLit struct {
	Raw string `json:"raw"`
	Pos Pos    `json:"pos"`
}

func (*DurationLit) exprNode()       {}
func (e *DurationLit) Position() Pos { return e.Pos }

// SearchTerm is a bare full-text term such as `timeout`.
type SearchTerm struct {
	Value string `json:"value"`
	Pos   Pos    `json:"pos"`
}

func (*SearchTerm) exprNode()       {}
func (e *SearchTerm) Position() Pos { return e.Pos }

// WildcardLit is a bare `*`.
type WildcardLit struct {
	Pos Pos `json:"pos"`
}

func (*WildcardLit) exprNode()       {}
func (e *WildcardLit) Position() Pos { return e.Pos }

// BinaryExpr is a comparison, boolean, or arithmetic operation.
type BinaryExpr struct {
	Op    string `json:"op"`
	Left  Expr   `json:"left"`
	Right Expr   `json:"right"`
	Pos   Pos    `json:"pos"`
}

func (*BinaryExpr) exprNode()       {}
func (e *BinaryExpr) Position() Pos { return e.Pos }

// UnaryExpr is a prefix `-`, `!`, or `NOT`.
type UnaryExpr struct {
	Op      string `json:"op"`
	Operand Expr   `json:"operand"`
	Pos     Pos    `json:"pos"`
}

func (*UnaryExpr) exprNode()       {}
func (e *UnaryExpr) Position() Pos { return e.Pos }

// CallExpr is a function call such as if(...) or lower(...).
type CallExpr struct {
	Name string `json:"name"`
	Args []Expr `json:"args,omitempty"`
	Pos  Pos    `json:"pos"`
}

func (*CallExpr) exprNode()       {}
func (e *CallExpr) Position() Pos { return e.Pos }
