package spl1

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/be-thomas/openspl/pkg/diag"
)

// FuzzParseSPL1 checks the front-end's contract for arbitrary input: it
// returns either a pipeline or a *diag.Error carrying at least one positioned
// diagnostic. A panic, an unpositioned error, or a nil pipeline with a nil
// error is a failure. The seed corpus is the interesting part of the language,
// because random bytes mostly exercise the invalid-token pass.
func FuzzParseSPL1(f *testing.F) {
	seeds := []string{
		`index=web status=500 | stats count by host`,
		`search error timeout`,
		`search "quoted phrase" host=web*`,
		`index=a OR index=b OR index=c | fields - _raw`,
		`earliest=-24h (status=500 OR status=503)`,
		`| stats count`,
		`|`,
		` `,
		`index=web host=*-* status=200-299`,
		`index=web | eval kb=bytes/1024 | sort limit=10 -kb | head 5`,
		`index=web earliest=-24h | timechart span=5m count by host`,
		`index=web | rex field=_raw "(?<user>\w+)"`,
		`index=web | bin span=1h _time | stats count by _time`,
		`index=web | table host, source`,
		`index=web | where a=1 AND b=2`,
		`index=web | stats dc(host) AS users BY status`,
		`msg="東京" host=web`,
		"index=web\r\n| stats count\r\n",
		`host=a+b`,
		`a [search b]`,
		`join [search x]`,
		`index=web | search index=c status=500`,
		`5month 1mo 1e+foo`,
		`"unterminated`,
		`stats count,`,
		`sort limit 1.5`,
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
