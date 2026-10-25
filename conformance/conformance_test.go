// Package conformance runs the checked-in conformance suite.
//
// Every case under cases/ is one SPL query and its expected result per target:
//
//	<name>.spl     the input query
//	<name>.esql    exact ES|QL output   (its presence makes the target applicable)
//	<name>.promql  exact PromQL output  (its presence makes the target applicable)
//	<name>.json    optional metadata: description, mapping, options, and the
//	               expected diagnostics or refusal per target
//
// Run with `go test ./conformance`. Regenerate goldens and expectation
// metadata with `go test ./conformance -update`.
package conformance

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/be-thomas/openspl/pkg/compiler"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/mapping"
)

var update = flag.Bool("update", false, "rewrite goldens and expectation metadata")

type expectation struct {
	Description string                  `json:"description,omitempty"`
	Mapping     string                  `json:"mapping,omitempty"`
	Options     *options                `json:"options,omitempty"`
	Targets     map[string]targetExpect `json:"targets,omitempty"`
}

type options struct {
	Metric string `json:"metric,omitempty"`
	Span   string `json:"span,omitempty"`
	Index  string `json:"index,omitempty"`
}

type targetExpect struct {
	Diagnostics []expectedDiag `json:"diagnostics,omitempty"`
	Error       *expectedDiag  `json:"error,omitempty"`
}

type expectedDiag struct {
	Code     string `json:"code"`
	Severity string `json:"severity,omitempty"`
	Line     int    `json:"line"`
	Col      int    `json:"col"`
}

func TestConformance(t *testing.T) {
	cases, orphans := discover(t, "cases")
	for _, orphan := range orphans {
		t.Errorf("orphan file has no matching .spl case: %s", orphan)
	}
	if len(cases) == 0 {
		t.Fatal("conformance suite is empty")
	}

	for _, src := range cases {
		src := src
		name := strings.TrimSuffix(strings.TrimPrefix(filepath.ToSlash(src), "cases/"), ".spl")
		t.Run(name, func(t *testing.T) {
			// Update mode mutates shared expectation state, so it stays serial.
			if !*update {
				t.Parallel()
			}
			runCase(t, src)
		})
	}
}

func discover(t *testing.T, root string) (cases, orphans []string) {
	t.Helper()

	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", root, err)
	}

	bases := map[string]bool{}
	for _, path := range files {
		if strings.HasSuffix(path, ".spl") {
			bases[strings.TrimSuffix(path, ".spl")] = true
			cases = append(cases, path)
		}
	}
	for _, path := range files {
		switch filepath.Ext(path) {
		case ".esql", ".promql", ".json":
			if base := strings.TrimSuffix(path, filepath.Ext(path)); !bases[base] {
				orphans = append(orphans, path)
			}
		}
	}
	sort.Strings(cases)
	sort.Strings(orphans)
	return cases, orphans
}

func runCase(t *testing.T, srcPath string) {
	t.Helper()

	src, err := os.ReadFile(srcPath)
	if err != nil {
		t.Fatalf("read case: %v", err)
	}
	base := strings.TrimSuffix(srcPath, ".spl")
	sidecarPath := base + ".json"
	exp := loadExpectation(t, sidecarPath)
	_, statErr := os.Stat(sidecarPath)
	hadSidecar := statErr == nil

	applicable := 0
	dirty := false

	for _, target := range compiler.Targets() {
		target := target
		goldenPath := base + "." + string(target)
		te, declared := exp.Targets[string(target)]

		_, goldenErr := os.Stat(goldenPath)
		hasGolden := goldenErr == nil
		if goldenErr != nil && !os.IsNotExist(goldenErr) {
			t.Fatalf("%s: %v", goldenPath, goldenErr)
		}
		if !declared && !hasGolden {
			continue
		}
		applicable++

		t.Run(string(target), func(t *testing.T) {
			if !*update {
				t.Parallel()
			}

			opts, err := caseOptions(base, exp)
			if err != nil {
				t.Fatalf("%v", err)
			}

			res, _, compileErr := compiler.Compile(string(src), target, opts)
			ds := collectDiagnostics(compileErr, res)

			if *update {
				exp.Targets = ensureTargets(exp.Targets)
				te = updateTarget(t, goldenPath, res, ds, te)
				if te.Error == nil && len(te.Diagnostics) == 0 {
					delete(exp.Targets, string(target))
				} else {
					exp.Targets[string(target)] = te
				}
				dirty = true
				return
			}

			if te.Error != nil {
				if hasGolden {
					t.Errorf("case expects an error but %s exists", goldenPath)
				}
				if _, ok := findDiag(ds, *te.Error); !ok {
					t.Fatalf("expected error %s at %d:%d, got:\n%s",
						te.Error.Code, te.Error.Line, te.Error.Col, render(ds))
				}
				return
			}

			if compileErr != nil || hasErrors(ds) {
				t.Fatalf("unexpected error:\n%s", render(ds))
			}
			if !hasGolden {
				t.Fatalf("missing golden %s; run go test ./conformance -update", goldenPath)
			}

			want, err := os.ReadFile(goldenPath)
			if err != nil {
				t.Fatalf("read golden: %v", err)
			}
			got := trimNewlines(res.Output)
			if got != trimNewlines(string(want)) {
				t.Errorf("output mismatch for %s\n--- got ---\n%s\n--- want ---\n%s",
					goldenPath, got, trimNewlines(string(want)))
			}

			gotDiags := normalizeDiagnostics(ds)
			wantDiags := te.Diagnostics
			if !sameDiags(gotDiags, wantDiags) {
				t.Errorf("diagnostics mismatch for %s\n got:\n%s\nwant:\n%s",
					goldenPath, renderExpected(gotDiags), renderExpected(wantDiags))
			}
		})
	}

	if applicable == 0 {
		t.Fatalf("%s is not exercised by any target", srcPath)
	}
	if *update && dirty && (hadSidecar || hasExpectations(exp)) {
		writeExpectation(t, sidecarPath, exp)
	}
}

// updateTarget refreshes one target. Targets become applicable only through an
// existing golden or an existing sidecar entry, so -update never widens a case
// to a target it did not already declare.
func updateTarget(t *testing.T, goldenPath string, res *compiler.Result, ds []diag.Diagnostic, te targetExpect) targetExpect {
	t.Helper()

	if hasErrors(ds) || res == nil {
		want := firstError(ds)
		if want == nil {
			t.Fatalf("compile failed without an error diagnostic")
		}
		if err := os.Remove(goldenPath); err != nil && !os.IsNotExist(err) {
			t.Fatalf("remove stale golden: %v", err)
		}
		te.Error = want
		te.Diagnostics = nil
		return te
	}

	te.Error = nil
	te.Diagnostics = normalizeDiagnostics(ds)
	if err := os.WriteFile(goldenPath, []byte(trimNewlines(res.Output)+"\n"), 0o644); err != nil {
		t.Fatalf("write golden: %v", err)
	}
	return te
}

func caseOptions(base string, exp expectation) (compiler.Options, error) {
	var opts compiler.Options
	if exp.Options != nil {
		opts.Metric = exp.Options.Metric
		opts.Span = exp.Options.Span
		opts.Index = exp.Options.Index
	}
	if exp.Mapping != "" {
		cfg, err := mapping.Load(filepath.Join(filepath.Dir(base), exp.Mapping))
		if err != nil {
			return opts, err
		}
		opts.Mapping = cfg
	}
	return opts, nil
}

func loadExpectation(t *testing.T, path string) expectation {
	t.Helper()
	var exp expectation

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return exp
	}
	if err != nil {
		t.Fatalf("read expectation: %v", err)
	}
	dec := json.NewDecoder(strings.NewReader(string(data)))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&exp); err != nil {
		t.Fatalf("%s: %v", path, err)
	}
	return exp
}

func writeExpectation(t *testing.T, path string, exp expectation) {
	t.Helper()
	data, err := json.MarshalIndent(exp, "", "  ")
	if err != nil {
		t.Fatalf("encode expectation: %v", err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		t.Fatalf("write expectation: %v", err)
	}
}

func ensureTargets(m map[string]targetExpect) map[string]targetExpect {
	if m == nil {
		return map[string]targetExpect{}
	}
	return m
}

func hasExpectations(exp expectation) bool {
	if exp.Description != "" || exp.Mapping != "" || exp.Options != nil {
		return true
	}
	for _, te := range exp.Targets {
		if te.Error != nil || len(te.Diagnostics) > 0 {
			return true
		}
	}
	return false
}

func collectDiagnostics(err error, res *compiler.Result) []diag.Diagnostic {
	var ds []diag.Diagnostic
	if res != nil {
		ds = append(ds, res.Diagnostics...)
	}
	if err != nil {
		if extra, ok := diag.FromError(err); ok {
			ds = append(ds, extra...)
		}
	}
	ds = dedupeDiagnostics(ds)
	diag.Sort(ds)
	return ds
}

func dedupeDiagnostics(ds []diag.Diagnostic) []diag.Diagnostic {
	type key struct {
		code     string
		severity diag.Severity
		line     int
		col      int
	}
	seen := map[key]bool{}
	out := ds[:0]
	for _, d := range ds {
		k := key{d.Code, d.Severity, d.Pos.Line, d.Pos.Col}
		if seen[k] {
			continue
		}
		seen[k] = true
		out = append(out, d)
	}
	return out
}

func hasErrors(ds []diag.Diagnostic) bool {
	for _, d := range ds {
		if d.Severity == diag.SeverityError {
			return true
		}
	}
	return false
}

func firstError(ds []diag.Diagnostic) *expectedDiag {
	for _, d := range ds {
		if d.Severity == diag.SeverityError {
			return &expectedDiag{Code: d.Code, Line: d.Pos.Line, Col: d.Pos.Col}
		}
	}
	return nil
}

func findDiag(ds []diag.Diagnostic, want expectedDiag) (diag.Diagnostic, bool) {
	for _, d := range ds {
		if d.Code == want.Code && d.Pos.Line == want.Line && d.Pos.Col == want.Col &&
			d.Severity == diag.SeverityError {
			return d, true
		}
	}
	return diag.Diagnostic{}, false
}

func normalizeDiagnostics(ds []diag.Diagnostic) []expectedDiag {
	out := make([]expectedDiag, 0, len(ds))
	for _, d := range ds {
		out = append(out, expectedDiag{
			Code:     d.Code,
			Severity: string(d.Severity),
			Line:     d.Pos.Line,
			Col:      d.Pos.Col,
		})
	}
	return out
}

func sameDiags(a, b []expectedDiag) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func render(ds []diag.Diagnostic) string {
	if len(ds) == 0 {
		return "(no diagnostics)"
	}
	return diag.Render(ds)
}

func renderExpected(ds []expectedDiag) string {
	if len(ds) == 0 {
		return "(none)"
	}
	var b strings.Builder
	for _, d := range ds {
		fmt.Fprintf(&b, "%d:%d %s %s\n", d.Line, d.Col, d.Severity, d.Code)
	}
	return strings.TrimRight(b.String(), "\n")
}

func trimNewlines(s string) string { return strings.TrimRight(s, "\n") }
