// Corpus tests: the gathered SPL and SPL2 query collections.
//
// The corpus under corpus/ holds query collections gathered by
// scripts/gather-corpus.py. A collection is one of two kinds: documented
// collections are the example queries of the Splunk documentation, and derived
// collections are transformed from another source, which is currently the
// security_content detection searches with their macros expanded. Each
// collection records its dialect, kind, source versions, and - for a derived
// collection - the transformation and the licence in metadata.json, because a
// query file says nothing about the version or the provenance it belongs to.
//
// Three properties are tested:
//
//   - robustness: every query either parses or is refused with a positioned
//     diagnostic. A panic, an unpositioned error, or a position outside the
//     input is a failure, whatever the query says.
//   - stability: the outcome (accepted, or the diagnostic code of a refusal)
//     matches the recorded baseline, so any change in what the parser does
//     with a gathered query shows up as a reviewable diff.
//   - coverage: every implemented command starts at least one stage somewhere
//     in the collections of its dialect, so no command goes untested.
//
// The baseline is not a claim that the current result is right. A gathered
// query the parser refuses stays refused until the parser is fixed, and fixing
// it changes exactly one line of the baseline. Rewrite the baseline with
//
//	go test ./conformance -run TestCorpus -update-corpus
//
// and review the diff: every changed outcome is a behaviour change.
package conformance

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/be-thomas/openspl/pkg/compiler"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/parser/commands"
)

var updateCorpus = flag.Bool("update-corpus", false,
	"rewrite the corpus outcome baselines")

// corpusMeta is one collection's metadata.json.
type corpusMeta struct {
	Format     int            `json:"corpus_format"`
	Dialect    string         `json:"dialect"`
	Kind       string         `json:"kind"`
	Source     string         `json:"source"`
	BaseURL    string         `json:"base_url"`
	Versions   []string       `json:"versions"`
	Retrieved  string         `json:"retrieved"`
	Queries    int            `json:"queries"`
	PerVersion map[string]int `json:"queries_per_version"`
	Gatherer   string         `json:"gatherer"`
	Transform  string         `json:"transform"`
	License    string         `json:"license"`
	Note       string         `json:"note"`
}

// corpusQuery is one line of queries.jsonl.
type corpusQuery struct {
	Query    string   `json:"query"`
	Pages    []string `json:"pages"`
	Versions []string `json:"versions"`
	Expect   string   `json:"expect,omitempty"`
}

func TestCorpus(t *testing.T) {
	dirs := corpusDirs(t)
	if len(dirs) == 0 {
		t.Fatal("corpus is empty; run scripts/gather-corpus.py")
	}

	dialects := map[string]bool{}
	byDialect := map[string][]corpusQuery{}
	total := 0
	for _, dir := range dirs {
		meta, queries := loadCorpus(t, dir)
		dialects[meta.Dialect] = true
		byDialect[meta.Dialect] = append(byDialect[meta.Dialect], queries...)
		total += len(queries)
		t.Run(filepath.Base(dir), func(t *testing.T) {
			runCorpus(t, dir, meta, queries)
		})
	}
	for _, want := range []string{"spl1", "spl2"} {
		if !dialects[want] {
			t.Errorf("corpus has no %s collection", want)
		}
	}
	if total < 2000 {
		t.Errorf("corpus has %d queries; at least 2000 are expected", total)
	}
	// Coverage is per dialect rather than per collection: a collection is one
	// source, and a source that documents a single command should not have to
	// repeat every other command's example.
	for _, dialect := range []string{"spl1", "spl2"} {
		checkCoverage(t, dialect, byDialect[dialect])
	}
}

func corpusDirs(t *testing.T) []string {
	t.Helper()
	metas, err := filepath.Glob(filepath.Join("corpus", "*", "metadata.json"))
	if err != nil {
		t.Fatalf("glob corpus: %v", err)
	}
	dirs := make([]string, 0, len(metas))
	for _, path := range metas {
		dirs = append(dirs, filepath.Dir(path))
	}
	sort.Strings(dirs)
	return dirs
}

func loadCorpus(t *testing.T, dir string) (corpusMeta, []corpusQuery) {
	t.Helper()

	var meta corpusMeta
	raw, err := os.ReadFile(filepath.Join(dir, "metadata.json"))
	if err != nil {
		t.Fatalf("read metadata: %v", err)
	}
	if err := json.Unmarshal(raw, &meta); err != nil {
		t.Fatalf("%s: parse metadata: %v", dir, err)
	}
	checkMeta(t, dir, meta)

	f, err := os.Open(filepath.Join(dir, "queries.jsonl"))
	if err != nil {
		t.Fatalf("%s: open queries: %v", dir, err)
	}
	defer f.Close()

	var queries []corpusQuery
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1<<20)
	for line := 1; scanner.Scan(); line++ {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		var q corpusQuery
		if err := json.Unmarshal([]byte(text), &q); err != nil {
			t.Fatalf("%s: line %d: %v", dir, line, err)
		}
		if strings.TrimSpace(q.Query) == "" {
			t.Fatalf("%s: line %d: empty query", dir, line)
		}
		if len(q.Pages) == 0 {
			t.Fatalf("%s: line %d: query has no source page", dir, line)
		}
		if len(q.Versions) == 0 {
			t.Fatalf("%s: line %d: query has no version", dir, line)
		}
		for _, v := range q.Versions {
			if !contains(meta.Versions, v) {
				t.Fatalf("%s: line %d: version %q is not in metadata", dir, line, v)
			}
		}
		queries = append(queries, q)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("%s: read queries: %v", dir, err)
	}
	if meta.Queries != len(queries) {
		t.Fatalf("%s: metadata says %d queries, file has %d",
			dir, meta.Queries, len(queries))
	}
	return meta, queries
}

func checkMeta(t *testing.T, dir string, meta corpusMeta) {
	t.Helper()
	switch {
	case meta.Format != 1:
		t.Errorf("%s: corpus_format is %d, want 1", dir, meta.Format)
	case meta.Dialect != "spl1" && meta.Dialect != "spl2":
		t.Errorf("%s: dialect %q is not spl1 or spl2", dir, meta.Dialect)
	case meta.Source == "":
		t.Errorf("%s: source is empty", dir)
	case meta.BaseURL == "":
		t.Errorf("%s: base_url is empty", dir)
	case len(meta.Versions) == 0:
		t.Errorf("%s: versions is empty", dir)
	case meta.Retrieved == "":
		t.Errorf("%s: retrieved is empty", dir)
	case meta.Gatherer == "":
		t.Errorf("%s: gatherer is empty", dir)
	}
	checkMetaKind(t, dir, meta)
	if meta.Queries <= 0 {
		t.Errorf("%s: queries is %d", dir, meta.Queries)
	}
	for version, count := range meta.PerVersion {
		if !contains(meta.Versions, version) {
			t.Errorf("%s: queries_per_version has %q, not a listed version",
				dir, version)
		}
		if count < 0 || count > meta.Queries {
			t.Errorf("%s: queries_per_version[%q] = %d, want 0..%d",
				dir, version, count, meta.Queries)
		}
	}
}

// checkMetaKind enforces what a kind promises. A derived collection states how
// the source was transformed and under what licence, and carries that licence
// text beside the queries; a documented collection has no transformation to
// record, so a transform or licence field there is a copy-paste error rather
// than extra information.
func checkMetaKind(t *testing.T, dir string, meta corpusMeta) {
	t.Helper()
	switch meta.Kind {
	case "documented":
		if meta.Transform != "" || meta.License != "" {
			t.Errorf("%s: documented collection declares transform or license",
				dir)
		}
	case "derived":
		if meta.Transform == "" {
			t.Errorf("%s: derived collection does not say how it was transformed",
				dir)
		}
		if meta.License == "" {
			t.Errorf("%s: derived collection does not name a licence", dir)
		}
		if _, err := os.Stat(filepath.Join(dir, "LICENSE")); err != nil {
			t.Errorf("%s: derived collection has no LICENSE: %v", dir, err)
		}
	default:
		t.Errorf("%s: kind %q is not documented or derived", dir, meta.Kind)
	}
}

func runCorpus(t *testing.T, dir string, meta corpusMeta, queries []corpusQuery) {
	version := compiler.SPLVersion1
	if meta.Dialect == "spl2" {
		version = compiler.SPLVersion2
	}

	mismatches := 0
	reported := 0
	report := func(format string, args ...any) {
		reported++
		if reported <= 20 {
			t.Errorf(format, args...)
		}
	}
	for i, q := range queries {
		got, err := corpusOutcome(q.Query, version)
		if err != nil {
			// Rewriting a baseline cannot fix a panic or an unpositioned
			// diagnostic, so this stays an error in update mode too.
			report("%s: query %d %q: %v", dir, i+1, q.Query, err)
			continue
		}
		if q.Expect == got {
			continue
		}
		if *updateCorpus {
			// -update-corpus rewrites and reports nothing: the diff is the
			// output, and failing the run would make `make corpus-update`
			// exit non-zero exactly when a re-record is what is wanted.
			queries[i].Expect = got
			continue
		}
		mismatches++
		if q.Expect == "" {
			report("%s: query %d %q: no baseline outcome (got %q)",
				dir, i+1, q.Query, got)
			continue
		}
		report("%s: query %d %q: outcome %q, baseline says %q",
			dir, i+1, q.Query, got, q.Expect)
	}
	if mismatches > 20 {
		t.Errorf("%s: %d queries differ in total", dir, mismatches)
	}
	if reported > 20 {
		t.Errorf("%s: %d queries failed in total", dir, reported)
	}

	if *updateCorpus {
		writeCorpus(t, dir, queries)
	}
}

// corpusOutcome parses one query and describes the result as "accepted" or
// "error:<code>". A panic, an unpositioned error, or a position outside the
// input is an error: the corpus is a robustness test, not only a snapshot.
func corpusOutcome(src string, version int) (outcome string, err error) {
	defer func() {
		if r := recover(); r != nil {
			outcome, err = "", fmt.Errorf("panic: %v", r)
		}
	}()

	_, parseErr := compiler.ParseVersion(src, version)
	if parseErr == nil {
		return "accepted", nil
	}
	var de *diag.Error
	if !errors.As(parseErr, &de) {
		return "", fmt.Errorf("error is not a *diag.Error: %v", parseErr)
	}
	for _, d := range de.Diagnostics {
		if d.Severity != diag.SeverityError {
			continue
		}
		if d.Code == "" {
			return "", fmt.Errorf("diagnostic has no code: %v", d)
		}
		if !positionWithin(src, d.Pos.Line, d.Pos.Col) {
			return "", fmt.Errorf("diagnostic %s at %d:%d is outside the input",
				d.Code, d.Pos.Line, d.Pos.Col)
		}
		return "error:" + d.Code, nil
	}
	return "", fmt.Errorf("error has no error-severity diagnostic: %v", parseErr)
}

// positionWithin reports whether a 1-based byte position can name a place in
// src: a line that exists, and a column at most one past that line's end.
func positionWithin(src string, line, col int) bool {
	if line < 1 || col < 1 {
		return false
	}
	lines := strings.Split(src, "\n")
	if line > len(lines) {
		return false
	}
	return col <= len(lines[line-1])+1
}

func writeCorpus(t *testing.T, dir string, queries []corpusQuery) {
	t.Helper()
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	for _, q := range queries {
		if err := enc.Encode(q); err != nil {
			t.Fatalf("%s: encode query: %v", dir, err)
		}
	}
	path := filepath.Join(dir, "queries.jsonl")
	old, err := os.ReadFile(path)
	if err == nil && bytes.Equal(old, buf.Bytes()) {
		return
	}
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatalf("%s: write queries: %v", dir, err)
	}
}

// corpusCoverageExceptions names implemented commands the corpus cannot cover,
// with the reason. A command belongs here only when the documentation does not
// show it in the form the parser implements, so no documented query can
// exercise it.
var corpusCoverageExceptions = map[string]string{
	"spl2/where": "the SPL2 reference documents WHERE as a clause of FROM, " +
		"not as a piped command, so no documented query starts a stage with it",
}

// checkCoverage fails when an implemented command never appears as the start
// of a stage anywhere in the collections of its dialect. The corpus is the
// parser's broad input, so a command with no corpus query is a command nothing
// tests.
func checkCoverage(t *testing.T, dialect string, queries []corpusQuery) {
	t.Helper()

	implemented := commands.SPL1ImplementedNames()
	known := commands.SPL1Known
	if dialect == "spl2" {
		implemented = commands.SPL2ImplementedNames()
		known = commands.SPL2Known
	}

	counts := map[string]int{}
	for _, q := range queries {
		for _, stage := range stages(q.Query) {
			word := leadingWord(stage)
			if word == "" {
				continue
			}
			switch {
			case contains(implemented, word):
				counts[word]++
			case known(word):
				// A real command outside the supported subset: the
				// refusal is the expected outcome, not coverage.
			default:
				// Not a command name, so this stage is an implicit
				// search.
				counts["search"]++
			}
		}
	}

	var missing []string
	for _, name := range implemented {
		if counts[name] > 0 {
			continue
		}
		if reason, ok := corpusCoverageExceptions[dialect+"/"+name]; ok {
			t.Logf("%s: %-10s not covered: %s", dialect, name, reason)
			continue
		}
		missing = append(missing, name)
	}
	if len(missing) > 0 {
		t.Errorf("%s: implemented commands with no corpus query: %s",
			dialect, strings.Join(missing, ", "))
	}
	for _, name := range implemented {
		if _, ok := corpusCoverageExceptions[dialect+"/"+name]; ok {
			continue
		}
		t.Logf("%s: %-10s %d stages", dialect, name, counts[name])
	}
}

// stages splits a query on the pipes that separate stages, ignoring pipes
// inside strings and inside brackets. A subsearch stage is not a top-level
// stage, so it does not count as coverage.
func stages(src string) []string {
	var out []string
	var quote byte
	depth, start := 0, 0
	for i := 0; i < len(src); i++ {
		c := src[i]
		if quote != 0 {
			switch {
			case c == '\\':
				i++
			case c == quote:
				quote = 0
			}
			continue
		}
		switch c {
		case '"', '\'':
			quote = c
		case '[', '(', '{':
			depth++
		case ']', ')', '}':
			if depth > 0 {
				depth--
			}
		case '|':
			if depth == 0 {
				out = append(out, src[start:i])
				start = i + 1
			}
		}
	}
	return append(out, src[start:])
}

// leadingWord returns the first identifier of a stage, lowercased, or "" when
// the stage does not start with one.
func leadingWord(stage string) string {
	stage = strings.TrimSpace(stage)
	end := 0
	for end < len(stage) {
		c := stage[end]
		if c == '_' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' ||
			(end > 0 && c >= '0' && c <= '9') {
			end++
			continue
		}
		break
	}
	return strings.ToLower(stage[:end])
}

func contains(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}
