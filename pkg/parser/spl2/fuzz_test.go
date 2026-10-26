package spl2

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/be-thomas/openspl/pkg/diag"
)

// FuzzParseSPL2 checks the SPL2 subset's contract for arbitrary input: it
// returns either a pipeline or a *diag.Error carrying at least one positioned
// diagnostic. A panic, an unpositioned error, or a nil pipeline with a nil
// error is a failure. Everything outside the subset must be refused rather
// than approximated, so a successful parse is also checked for its shape: the
// first stage is always `from`.
func FuzzParseSPL2(f *testing.F) {
	seeds := []string{
		`from main`,
		`from main | where status=500`,
		`from main | where status=500 AND host="web" | stats count() AS n BY host`,
		`from logs-2024-01 | stats dc(user) BY host`,
		`from logs*`,
		`from *`,
		`from datamodel:events`,
		`from main | eval x=1`,
		`from main | where status=*`,
		`from main | stats count(*)`,
		`from a | from b`,
		`from main |`,
		`| from main`,
		`index=web | stats count`,
		` `,
		`from main | where (a=1`,
		`from main | where a in (1,2)`,
		`from main | stats count()`,
		`from main | where a=1 or b=2`,
		`from main | where not a=1`,
		`from main | stats sum(bytes) AS total BY host, status`,
		`from "quoted"`,
		`from main | stats`,
		"from main\r\n| where status=500\r\n",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	f.Fuzz(func(t *testing.T, src string) {
		// Keep the corpus bounded: the point is parser shape, not a stress test.
		if len(src) > 1<<16 || !utf8.ValidString(src) {
			t.Skip()
		}

		p, err := Parse(src)
		if err != nil {
			assertDiagError(t, src, err)
			return
		}
		if p == nil {
			t.Fatalf("Parse(%q) returned no pipeline and no error", src)
		}
		if len(p.Stages) == 0 {
			t.Fatalf("Parse(%q) returned an empty pipeline", src)
		}
		if got := p.Stages[0].StageName(); got != "search" {
			t.Fatalf("Parse(%q) starts with stage %q, want the lowered from search", src, got)
		}
		for _, s := range p.Stages {
			if pos := s.Position(); pos.Line < 1 || pos.Col < 1 || pos.Offset < 0 {
				t.Fatalf("Parse(%q) produced stage position %+v", src, pos)
			}
		}
	})
}

// assertDiagError requires a positioned *diag.Error. The parser is expected to
// report invalid input rather than fail in a way the caller cannot print.
func assertDiagError(t *testing.T, src string, err error) {
	t.Helper()

	ds, ok := diag.FromError(err)
	if !ok {
		t.Fatalf("Parse(%q) returned %T %v, want *diag.Error", src, err, err)
	}
	if len(ds) == 0 {
		t.Fatalf("Parse(%q) returned an empty diagnostic list", src)
	}
	for _, d := range ds {
		if d.Pos.Line < 1 || d.Pos.Col < 1 || d.Pos.Offset < 0 {
			t.Fatalf("Parse(%q) produced diagnostic %+v at %+v", src, d, d.Pos)
		}
		if d.Code == "" || strings.TrimSpace(d.Code) != d.Code {
			t.Fatalf("Parse(%q) produced diagnostic with code %q", src, d.Code)
		}
	}
}
