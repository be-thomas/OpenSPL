// Corpus compile report.
//
// The corpus tests assert outcomes; this report measures them. It answers
// three questions for every collection: how many queries parse to the IR, how
// many of those the IR translates to each back-end, and how many translate
// end-to-end. Run it with
//
//	make corpus-report
//
// and paste the printed table into README.md's "Conformance report" section.
package conformance

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/compiler"
)

var corpusReport = flag.Bool("corpus-report", false,
	"print the corpus compile report")

// reportTally counts one row of the report.
type reportTally struct {
	queries int
	parsed  int
	esql    int
	promql  int
}

func (r *reportTally) add(other reportTally) {
	r.queries += other.queries
	r.parsed += other.parsed
	r.esql += other.esql
	r.promql += other.promql
}

type reportRow struct {
	name  string
	tally reportTally
}

func TestCorpusReport(t *testing.T) {
	if !*corpusReport {
		t.Skip("set -corpus-report to print the compile report")
	}
	var rows []reportRow
	var total reportTally
	for _, dir := range corpusDirs(t) {
		meta, queries := loadCorpus(t, dir)
		got := measure(t, meta, queries)
		rows = append(rows, reportRow{name: filepath.Base(dir), tally: got})
		total.add(got)
	}
	rows = append(rows, reportRow{name: "**Total**", tally: total})

	fmt.Println("| Collection | Queries | → IR | SPL→IR | → ES\\|QL | SPL→ES\\|QL | IR→ES\\|QL | → PromQL | SPL→PromQL | IR→PromQL |")
	fmt.Println("|---|---|---|---|---|---|---|---|---|---|")
	for _, row := range rows {
		t := row.tally
		fmt.Printf("| %s | %s | %s | %s | %s | %s | %s | %s | %s | %s |\n",
			row.name, count(t.queries), count(t.parsed),
			percent(t.parsed, t.queries),
			count(t.esql), percent(t.esql, t.queries), percent(t.esql, t.parsed),
			count(t.promql), percent(t.promql, t.queries),
			percent(t.promql, t.parsed))
	}
}

// measure counts, for one collection, the queries that parse and the queries
// that also emit a non-empty query for each back-end.
func measure(t *testing.T, meta corpusMeta, queries []corpusQuery) reportTally {
	t.Helper()
	version := compiler.SPLVersion1
	if meta.Dialect == "spl2" {
		version = compiler.SPLVersion2
	}
	var tally reportTally
	for _, q := range queries {
		tally.queries++
		p, err := compiler.ParseVersion(q.Query, version)
		if err != nil {
			continue
		}
		tally.parsed++
		if emits(p, compiler.ESQL) {
			tally.esql++
		}
		if emits(p, compiler.PromQL) {
			tally.promql++
		}
	}
	return tally
}

func emits(p *ast.Pipeline, target compiler.Target) bool {
	res, err := compiler.Transpile(p, target, compiler.Options{})
	if err != nil || res == nil {
		return false
	}
	return strings.TrimSpace(res.Output) != ""
}

// count formats a count with thousands separators, so the printed table can be
// pasted into the README unchanged.
func count(n int) string {
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var b strings.Builder
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(r)
	}
	return b.String()
}

func percent(n, of int) string {
	if of == 0 {
		return "—"
	}
	return fmt.Sprintf("%.1f%%", 100*float64(n)/float64(of))
}
