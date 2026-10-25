package compiler

import (
	"strings"
	"testing"

	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/mapping"
)

func mustCompile(t *testing.T, src string, target Target, opts Options) *Result {
	t.Helper()
	res, _, err := Compile(src, target, opts)
	if err != nil {
		t.Fatalf("Compile(%q) failed: %v", src, err)
	}
	return res
}

func hasCode(ds []diag.Diagnostic, code string) bool {
	for _, d := range ds {
		if d.Code == code {
			return true
		}
	}
	return false
}

func messageFor(ds []diag.Diagnostic, code string) string {
	for _, d := range ds {
		if d.Code == code {
			return d.Message
		}
	}
	return ""
}

func TestESQLGolden(t *testing.T) {
	tests := []struct {
		name string
		src  string
		opts Options
		want string
	}{
		{
			name: "search and stats",
			src:  `index=web status=500 | stats count by host`,
			want: "FROM web\n| WHERE status == 500\n| STATS count = COUNT(*) BY `host.name`",
		},
		{
			name: "time bounds and timechart",
			src:  `index=web earliest=-24h | timechart span=5m count by host`,
			want: "FROM web\n" +
				"| WHERE `@timestamp` >= NOW() - 24 hours\n" +
				"| STATS count = COUNT(*) BY BUCKET(`@timestamp`, 5 minutes), `host.name`",
		},
		{
			name: "full text and field mapping",
			src:  `search error | table host, source`,
			want: "FROM logs-*\n| WHERE MATCH(message, \"error\")\n| KEEP `host.name`, `log.file.path`",
		},
		{
			name: "multi index",
			src:  `index=a OR index=b | fields - _raw`,
			want: "FROM a, b\n| DROP message",
		},
		{
			name: "wildcard comparison",
			src:  `index=web host=web*`,
			want: "FROM web\n| WHERE `host.name` LIKE \"web*\"",
		},
		{
			name: "eval sort and head",
			src:  `index=web | eval kb=bytes/1024 | sort limit=10 -kb | head 5`,
			want: "FROM web\n| EVAL kb = bytes / 1024\n| SORT kb DESC\n| LIMIT 10\n| LIMIT 5",
		},
		{
			name: "bin keeps a derived column",
			src:  `index=web | bin span=1h _time | stats count by _time`,
			want: "FROM web\n| EVAL _time = BUCKET(`@timestamp`, 1 hour)\n| STATS count = COUNT(*) BY _time",
		},
		{
			name: "index override",
			src:  `index=web | stats count`,
			opts: Options{Index: "logs-web-*"},
			want: "FROM logs-web-*\n| STATS count = COUNT(*)",
		},
		{
			name: "mapping indexes",
			src:  `index=web | stats count`,
			opts: Options{Mapping: &mapping.Config{
				TimeField: "@timestamp",
				TextField: "message",
				Indexes:   map[string]string{"web": "logs-web-*"},
				Prometheus: mapping.Prometheus{
					Metric: "events_total", Function: "increase", IndexLabel: "job",
				},
				Timechart: mapping.Timechart{DefaultSpan: "5m"},
			}},
			want: "FROM logs-web-*\n| STATS count = COUNT(*)",
		},
		{
			name: "three index disjunction",
			src:  `index=a OR index=b OR index=c | stats count`,
			want: "FROM a, b, c\n| STATS count = COUNT(*)",
		},
		{
			name: "span option supplies the timechart fallback",
			src:  `index=web | timechart count by host`,
			opts: Options{Span: "1h"},
			want: "FROM web\n| STATS count = COUNT(*) BY BUCKET(`@timestamp`, 1 hour), `host.name`",
		},
		{
			name: "time bounds parenthesize a top-level or",
			src:  `index=web earliest=-24h (status=500 OR status=503)`,
			want: "FROM web\n| WHERE `@timestamp` >= NOW() - 24 hours AND (status == 500 OR status == 503)",
		},
		{
			name: "lifted time bounds parenthesize a top-level or",
			src:  `earliest=-24h status=500 OR status=503`,
			want: "FROM logs-*\n| WHERE `@timestamp` >= NOW() - 24 hours AND (status == 500 OR status == 503)",
		},
		{
			name: "adjacent wildcard and hyphen runs stay literal",
			src:  `index=web host=*-* status=200-299`,
			want: "FROM web\n| WHERE `host.name` LIKE \"*-*\" AND status == \"200-299\"",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := mustCompile(t, tc.src, ESQL, tc.opts)
			if res.Output != tc.want {
				t.Fatalf("output:\n%s\nwant:\n%s", res.Output, tc.want)
			}
		})
	}
}

func TestESQLDiagnostics(t *testing.T) {
	t.Run("full text approximation warns", func(t *testing.T) {
		res := mustCompile(t, `search timeout`, ESQL, Options{})
		if !hasCode(res.Diagnostics, "fulltext-approximation") {
			t.Fatalf("diagnostics = %#v", res.Diagnostics)
		}
	})

	t.Run("rex warns about grok", func(t *testing.T) {
		res := mustCompile(t, `index=web | rex field=_raw "(?<user>\w+)"`, ESQL, Options{})
		if !hasCode(res.Diagnostics, "rex-lossy") {
			t.Fatalf("diagnostics = %#v", res.Diagnostics)
		}
		if !strings.Contains(res.Output, `GROK message "%{WORD:user}"`) {
			t.Fatalf("output = %q", res.Output)
		}
	})

	t.Run("bin reports a derived column", func(t *testing.T) {
		res := mustCompile(t, `index=web | bin span=1h _time`, ESQL, Options{})
		if !hasCode(res.Diagnostics, "bin-bucket") {
			t.Fatalf("diagnostics = %#v", res.Diagnostics)
		}
	})

	t.Run("invalid duration fails", func(t *testing.T) {
		_, _, err := Compile(`index=web | timechart span=abc count`, ESQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "invalid-span") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("unmapped aggregation fails", func(t *testing.T) {
		_, _, err := Compile(`index=web | stats median(bytes)`, ESQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "unsupported-agg") {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestESQLIndexFilters(t *testing.T) {
	t.Run("mixed index OR is refused", func(t *testing.T) {
		res, _, err := Compile(`index=web OR host=api`, ESQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "mixed-index-filter") {
			t.Fatalf("error = %v", err)
		}
		if res == nil || res.Output != "" {
			t.Fatalf("refused query produced output: %#v", res)
		}
	})

	t.Run("late index filter names the emitted FROM", func(t *testing.T) {
		res := mustCompile(t, `index=a OR index=b | search index=c status=500`, ESQL, Options{})
		msg := messageFor(res.Diagnostics, "late-index-filter")
		if !strings.Contains(msg, "a, b") {
			t.Fatalf("late-index-filter = %q, want the emitted FROM", msg)
		}
		if want := "FROM a, b\n| WHERE status == 500"; res.Output != want {
			t.Fatalf("output = %q, want %q", res.Output, want)
		}
	})

	t.Run("late mixed OR is refused without a partial query", func(t *testing.T) {
		res, _, err := Compile(`index=a | search index=b OR host=x`, ESQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "mixed-index-filter") {
			t.Fatalf("error = %v", err)
		}
		if res == nil || res.Output != "" {
			t.Fatalf("refused query produced output: %#v", res)
		}
	})
}

func TestDiagnosticsSurviveErrors(t *testing.T) {
	res, _, err := Compile(`index=web | timechart count by host | stats median(bytes)`, ESQL, Options{})
	if err == nil {
		t.Fatal("expected an error")
	}
	if !hasCode(res.Diagnostics, "default-span") {
		t.Fatalf("warning was dropped: %#v", res.Diagnostics)
	}
	if !hasCode(res.Diagnostics, "unsupported-agg") {
		t.Fatalf("error diagnostic is missing: %#v", res.Diagnostics)
	}
	if res.Output != "" {
		t.Fatalf("failed compile produced output: %q", res.Output)
	}
}

func TestPromQL(t *testing.T) {
	t.Run("timechart becomes a rate", func(t *testing.T) {
		res := mustCompile(t, `index=metrics | timechart span=5m count by host`, PromQL, Options{})
		want := `sum by (host) (increase(splunk_events_total{job="metrics"}[5m]))`
		if res.Output != want {
			t.Fatalf("output = %q, want %q", res.Output, want)
		}
	})

	t.Run("duplicate filters collapse", func(t *testing.T) {
		res := mustCompile(t, `index=metrics job=metrics | stats count by host`, PromQL, Options{})
		want := `sum by (host) (increase(splunk_events_total{job="metrics"}[5m]))`
		if res.Output != want {
			t.Fatalf("output = %q, want %q", res.Output, want)
		}
	})

	t.Run("distinct positive filters are refused", func(t *testing.T) {
		res, _, err := Compile(`index=web host=a host=b | stats count by host`, PromQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "promql-conflicting-filters") {
			t.Fatalf("error = %v", err)
		}
		if res.Output != "" {
			t.Fatalf("output = %q, want no partial query", res.Output)
		}
	})

	t.Run("negative filters merge", func(t *testing.T) {
		res := mustCompile(t, `index=web status!=500 status!=503 | stats count by host`, PromQL, Options{})
		if !strings.Contains(res.Output, `status!~"500|503"`) {
			t.Fatalf("output = %q", res.Output)
		}
	})

	t.Run("filters become label matchers", func(t *testing.T) {
		res := mustCompile(t, `index=web host=web* status=500 | stats count by host`, PromQL, Options{})
		if !strings.Contains(res.Output, `host=~"^web.*$"`) {
			t.Fatalf("output = %q", res.Output)
		}
		if !strings.Contains(res.Output, `status="500"`) {
			t.Fatalf("output = %q", res.Output)
		}
	})

	t.Run("metric and span overrides", func(t *testing.T) {
		res := mustCompile(t, `index=metrics | stats count by host`, PromQL, Options{
			Metric: "custom_total",
			Span:   "1h",
		})
		if !strings.Contains(res.Output, "custom_total") || !strings.Contains(res.Output, "[1h]") {
			t.Fatalf("output = %q", res.Output)
		}
	})

	t.Run("limit becomes topk", func(t *testing.T) {
		res := mustCompile(t, `index=web | stats count by host | sort limit=3 -count`, PromQL, Options{})
		if !strings.HasPrefix(res.Output, "topk(3, ") {
			t.Fatalf("output = %q", res.Output)
		}
	})

	t.Run("head zero is ignored with a warning", func(t *testing.T) {
		res := mustCompile(t, `index=web | stats count by host | head 0`, PromQL, Options{})
		if !hasCode(res.Diagnostics, "promql-head-ignored") {
			t.Fatalf("diagnostics = %#v", res.Diagnostics)
		}
		if strings.Contains(res.Output, "topk") {
			t.Fatalf("output = %q, want no topk", res.Output)
		}
	})

	t.Run("sort limit without fields is ignored with a warning", func(t *testing.T) {
		res := mustCompile(t, `index=web | stats count by host | sort limit=3`, PromQL, Options{})
		if !hasCode(res.Diagnostics, "promql-sort-limit-ignored") {
			t.Fatalf("diagnostics = %#v", res.Diagnostics)
		}
		if strings.Contains(res.Output, "topk") {
			t.Fatalf("output = %q, want no topk", res.Output)
		}
	})

	t.Run("non-count aggregation is refused", func(t *testing.T) {
		_, _, err := Compile(`index=web | stats avg(bytes)`, PromQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "promql-unsupported-agg") {
			t.Fatalf("error = %v", err)
		}
	})

	t.Run("raw pipeline is refused", func(t *testing.T) {
		_, _, err := Compile(`index=web status=500 | table host`, PromQL, Options{})
		if err == nil {
			t.Fatal("expected an error")
		}
		ds, _ := diag.FromError(err)
		if !hasCode(ds, "promql-unsupported-stage") {
			t.Fatalf("error = %v", err)
		}
	})
}

func TestParseTarget(t *testing.T) {
	for _, name := range []string{"esql", "ESQL", "elastic", "elasticsearch"} {
		if got, err := ParseTarget(name); err != nil || got != ESQL {
			t.Errorf("ParseTarget(%q) = %q, %v", name, got, err)
		}
	}
	for _, name := range []string{"promql", "PromQL", "prometheus", "prom"} {
		if got, err := ParseTarget(name); err != nil || got != PromQL {
			t.Errorf("ParseTarget(%q) = %q, %v", name, got, err)
		}
	}
	if _, err := ParseTarget("kusto"); err == nil {
		t.Fatal("expected an unknown target to fail")
	}
}
