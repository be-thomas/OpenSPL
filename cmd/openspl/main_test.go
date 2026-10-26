package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

// runCaptured executes the CLI with stdout and stderr redirected. CLI output
// is small, so draining both pipes after the run cannot block.
func runCaptured(t *testing.T, args ...string) (code int, stdout, stderr string) {
	t.Helper()

	oldOut, oldErr := os.Stdout, os.Stderr
	rOut, wOut, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	rErr, wErr, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout, os.Stderr = wOut, wErr
	defer func() {
		os.Stdout, os.Stderr = oldOut, oldErr
	}()

	code = run(args)

	if err := wOut.Close(); err != nil {
		t.Fatal(err)
	}
	if err := wErr.Close(); err != nil {
		t.Fatal(err)
	}
	out, _ := io.ReadAll(rOut)
	errOut, _ := io.ReadAll(rErr)
	return code, string(out), string(errOut)
}

func TestCLITranspile(t *testing.T) {
	t.Run("success writes only the query", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", "index=web status=500")
		if code != 0 {
			t.Fatalf("exit = %d, stderr = %q", code, stderr)
		}
		if stdout != "FROM web\n| WHERE status == 500\n" {
			t.Fatalf("stdout = %q", stdout)
		}
		if stderr != "" {
			t.Fatalf("stderr = %q, want empty", stderr)
		}
	})

	t.Run("warnings survive an error and no query is printed", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile",
			"index=web | timechart count by host | stats median(bytes)")
		if code != 1 {
			t.Fatalf("exit = %d, want 1", code)
		}
		if stdout != "" {
			t.Fatalf("stdout = %q, want no partial query", stdout)
		}
		for _, want := range []string{"default-span", "unsupported-agg"} {
			if !strings.Contains(stderr, want) {
				t.Fatalf("stderr = %q, want %s", stderr, want)
			}
		}
	})

	t.Run("json reports success with a query", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", "--json",
			"index=web status=500")
		if code != 0 || stderr != "" {
			t.Fatalf("exit = %d, stderr = %q", code, stderr)
		}
		var payload struct {
			Query       string `json:"query"`
			OK          bool   `json:"ok"`
			Diagnostics []struct {
				Code string `json:"code"`
			} `json:"diagnostics"`
		}
		if err := json.Unmarshal([]byte(stdout), &payload); err != nil {
			t.Fatalf("stdout = %q: %v", stdout, err)
		}
		if !payload.OK || payload.Query == "" {
			t.Fatalf("payload = %+v, want ok with a query", payload)
		}
	})

	t.Run("json reports failure without a query", func(t *testing.T) {
		code, stdout, _ := runCaptured(t, "transpile", "--json", "-t", "promql",
			"index=web host=a host=b | stats count by host")
		if code != 1 {
			t.Fatalf("exit = %d, want 1", code)
		}
		var payload struct {
			Query       string `json:"query"`
			OK          bool   `json:"ok"`
			Diagnostics []struct {
				Code string `json:"code"`
			} `json:"diagnostics"`
		}
		if err := json.Unmarshal([]byte(stdout), &payload); err != nil {
			t.Fatalf("stdout = %q: %v", stdout, err)
		}
		if payload.OK || payload.Query != "" {
			t.Fatalf("payload = %+v, want failure without a query", payload)
		}
		found := false
		for _, d := range payload.Diagnostics {
			found = found || d.Code == "promql-conflicting-filters"
		}
		if !found {
			t.Fatalf("diagnostics = %+v, want promql-conflicting-filters", payload.Diagnostics)
		}
	})

	t.Run("invalid target is a usage error", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", "-t", "bogus", "index=web")
		if code != 2 || stdout != "" || !strings.Contains(stderr, "unknown target") {
			t.Fatalf("exit = %d, stdout = %q, stderr = %q", code, stdout, stderr)
		}
	})

	t.Run("empty query is a positioned error", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", " ")
		if code != 1 || stdout != "" || !strings.Contains(stderr, "1:1: error[empty-query]") {
			t.Fatalf("exit = %d, stdout = %q, stderr = %q", code, stdout, stderr)
		}
	})

	t.Run("strict treats warnings as errors", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", "--strict", "search error")
		if code != 1 || stdout != "" || !strings.Contains(stderr, "warnings treated as errors") {
			t.Fatalf("exit = %d, stdout = %q, stderr = %q", code, stdout, stderr)
		}
	})

	t.Run("spl-version 2 selects SPL2", func(t *testing.T) {
		code, stdout, stderr := runCaptured(t, "transpile", "--spl-version", "2",
			"from main | where status=500")
		if code != 0 || stderr != "" {
			t.Fatalf("exit = %d, stderr = %q", code, stderr)
		}
		if stdout != "FROM main\n| WHERE status == 500\n" {
			t.Fatalf("stdout = %q", stdout)
		}
	})

	t.Run("spl-version defaults to classic SPL", func(t *testing.T) {
		// `from` is a recognized classic SPL command, so without the flag this
		// is an unsupported-command refusal rather than an SPL2 parse.
		code, _, stderr := runCaptured(t, "transpile", "from main")
		if code != 1 || !strings.Contains(stderr, "unsupported-command") {
			t.Fatalf("exit = %d, stderr = %q", code, stderr)
		}
	})

	t.Run("unsupported spl-version is a usage error before reading input", func(t *testing.T) {
		for _, version := range []string{"0", "3", "two"} {
			code, stdout, stderr := runCaptured(t, "transpile", "--spl-version", version)
			if code != 2 || stdout != "" || !strings.Contains(stderr, "version") {
				t.Fatalf("version %q: exit = %d, stdout = %q, stderr = %q",
					version, code, stdout, stderr)
			}
		}
	})
}

func TestCLIAST(t *testing.T) {
	code, stdout, stderr := runCaptured(t, "ast", "index=web")
	if code != 0 || stderr != "" {
		t.Fatalf("exit = %d, stderr = %q", code, stderr)
	}
	var payload struct {
		Stages []struct {
			Expr *struct {
				Op string `json:"op"`
			} `json:"expr"`
		} `json:"stages"`
	}
	if err := json.Unmarshal([]byte(stdout), &payload); err != nil {
		t.Fatalf("stdout = %q: %v", stdout, err)
	}
	if len(payload.Stages) != 1 || payload.Stages[0].Expr == nil || payload.Stages[0].Expr.Op != "=" {
		t.Fatalf("payload = %+v", payload)
	}
}

func TestCLIASTSPL2(t *testing.T) {
	code, stdout, stderr := runCaptured(t, "ast", "--spl-version", "2", "from main")
	if code != 0 || stderr != "" {
		t.Fatalf("exit = %d, stderr = %q", code, stderr)
	}
	for _, want := range []string{`"op": "="`, `"value": "main"`} {
		if !strings.Contains(stdout, want) {
			t.Fatalf("stdout = %q, want %s", stdout, want)
		}
	}
}

func TestCLIUsage(t *testing.T) {
	if code, _, stderr := runCaptured(t); code != 2 || !strings.Contains(stderr, "Usage:") {
		t.Fatalf("no args: exit = %d, stderr = %q", code, stderr)
	}
	if code, _, stderr := runCaptured(t, "nope"); code != 2 || !strings.Contains(stderr, "unknown command") {
		t.Fatalf("bad command: exit = %d, stderr = %q", code, stderr)
	}
}
