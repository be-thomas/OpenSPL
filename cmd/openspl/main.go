// Command openspl transpiles classic Splunk SPL into Elastic ES|QL or
// Prometheus PromQL.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/be-thomas/openspl/pkg/ast"
	"github.com/be-thomas/openspl/pkg/compiler"
	"github.com/be-thomas/openspl/pkg/diag"
	"github.com/be-thomas/openspl/pkg/mapping"
	"github.com/be-thomas/openspl/pkg/version"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) == 0 {
		usage(os.Stderr)
		return 2
	}
	switch args[0] {
	case "transpile", "translate":
		return cmdTranspile(args[1:])
	case "ast":
		return cmdAST(args[1:])
	case "targets":
		return cmdTargets()
	case "version", "--version", "-v":
		fmt.Println(version.String())
		return 0
	case "help", "-h", "--help":
		usage(os.Stdout)
		return 0
	default:
		fmt.Fprintf(os.Stderr, "openspl: unknown command %q\n\n", args[0])
		usage(os.Stderr)
		return 2
	}
}

func usage(w io.Writer) {
	fmt.Fprint(w, `openspl - deterministic Splunk SPL to ES|QL and PromQL transpiler

Usage:
  openspl transpile [flags] <query>
  openspl transpile [flags] -f query.spl
  echo '<query>' | openspl transpile [flags]
  openspl ast [flags] <query>
  openspl targets
  openspl version

Transpile flags:
  -t, --target esql|promql   output target (default "esql")
      --mapping file.json    field and index mapping (default: built-in ECS map)
      --index name           override the index selected by the query
      --metric name          PromQL metric name (default from mapping)
      --span duration        PromQL range window / timechart default (e.g. 5m)
      --spl-version 1|2      input dialect (default 1: classic SPL; 2: SPL2)
  -f, --file file            read the query from a file ("-" for stdin)
      --json                 emit machine-readable JSON
      --strict               treat warnings as errors

AST flags:
      --spl-version 1|2      input dialect (default 1: classic SPL; 2: SPL2)

Examples:
  openspl transpile 'index=web status=500 | stats count by host'
  openspl transpile -t promql 'index=metrics | timechart span=5m count by host'
  openspl transpile -f savedsearch.spl --index logs-web-*
`)
}

func cmdTranspile(args []string) int {
	flagArgs, positional := splitFlags(args)

	fs := flag.NewFlagSet("transpile", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	target := "esql"
	splVersion := compiler.SPLVersion1
	var mappingPath, index, metric, span, file string
	var jsonOut, strict bool

	fs.StringVar(&target, "target", "esql", "output target")
	fs.StringVar(&target, "t", "esql", "output target")
	fs.StringVar(&mappingPath, "mapping", "", "mapping file")
	fs.StringVar(&index, "index", "", "index override")
	fs.StringVar(&metric, "metric", "", "PromQL metric name")
	fs.StringVar(&span, "span", "", "span or range")
	fs.IntVar(&splVersion, "spl-version", compiler.SPLVersion1, "input SPL dialect version")
	fs.StringVar(&file, "file", "", "query file")
	fs.StringVar(&file, "f", "", "query file")
	fs.BoolVar(&jsonOut, "json", false, "JSON output")
	fs.BoolVar(&strict, "strict", false, "treat warnings as errors")

	if err := fs.Parse(flagArgs); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			usage(os.Stdout)
			return 0
		}
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 2
	}

	t, err := compiler.ParseTarget(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 2
	}

	if !supportedSPLVersion(splVersion) {
		fmt.Fprintf(os.Stderr, "openspl: unsupported SPL version %d; supported versions: 1, 2\n", splVersion)
		return 2
	}

	var cfg *mapping.Config
	if mappingPath != "" {
		cfg, err = mapping.Load(mappingPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "openspl: mapping: %v\n", err)
			return 2
		}
	} else {
		cfg = mapping.Default()
	}

	src, err := readQuery(file, positional)
	if err != nil {
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 2
	}

	res, _, err := compiler.Compile(src, t, compiler.Options{
		Index:      index,
		Metric:     metric,
		Span:       span,
		Mapping:    cfg,
		SPLVersion: splVersion,
	})

	diagnostics := []diag.Diagnostic{}
	if res != nil {
		diagnostics = res.Diagnostics
	}
	if err != nil {
		if ds, ok := diag.FromError(err); ok {
			diagnostics = diag.Merge(diagnostics, ds)
		} else if len(diagnostics) == 0 {
			diagnostics = []diag.Diagnostic{{
				Severity: diag.SeverityError,
				Code:     "internal",
				Message:  err.Error(),
				Pos:      ast.Pos{Line: 1, Col: 1},
			}}
		}
	}

	warningsAsErrors := strict && hasWarnings(diagnostics)

	if jsonOut {
		payload := struct {
			Target      string            `json:"target"`
			Query       string            `json:"query,omitempty"`
			Diagnostics []diag.Diagnostic `json:"diagnostics,omitempty"`
			OK          bool              `json:"ok"`
		}{
			Target:      string(t),
			Diagnostics: diagnostics,
			OK:          err == nil && !warningsAsErrors,
		}
		if err == nil && res != nil {
			payload.Query = res.Output
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if encErr := enc.Encode(payload); encErr != nil {
			fmt.Fprintf(os.Stderr, "openspl: %v\n", encErr)
			return 1
		}
		if err != nil || warningsAsErrors {
			return 1
		}
		return 0
	}

	if len(diagnostics) > 0 {
		fmt.Fprintln(os.Stderr, diag.Render(diagnostics))
	}
	if err != nil {
		return 1
	}
	if warningsAsErrors {
		fmt.Fprintln(os.Stderr, "openspl: warnings treated as errors (--strict)")
		return 1
	}
	fmt.Println(res.Output)
	return 0
}

func cmdAST(args []string) int {
	flagArgs, positional := splitFlags(args)

	fs := flag.NewFlagSet("ast", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	splVersion := compiler.SPLVersion1
	var file string
	fs.StringVar(&file, "file", "", "query file")
	fs.StringVar(&file, "f", "", "query file")
	fs.IntVar(&splVersion, "spl-version", compiler.SPLVersion1, "input SPL dialect version")

	if err := fs.Parse(flagArgs); err != nil {
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 2
	}

	if !supportedSPLVersion(splVersion) {
		fmt.Fprintf(os.Stderr, "openspl: unsupported SPL version %d; supported versions: 1, 2\n", splVersion)
		return 2
	}

	src, err := readQuery(file, positional)
	if err != nil {
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 2
	}
	p, err := compiler.ParseVersion(src, splVersion)
	if err != nil {
		ds, ok := diag.FromError(err)
		if !ok {
			ds = []diag.Diagnostic{{
				Severity: diag.SeverityError,
				Code:     "internal",
				Message:  err.Error(),
				Pos:      ast.Pos{Line: 1, Col: 1},
			}}
		}
		fmt.Fprintln(os.Stderr, diag.Render(ds))
		return 1
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(p); err != nil {
		fmt.Fprintf(os.Stderr, "openspl: %v\n", err)
		return 1
	}
	return 0
}

func cmdTargets() int {
	fmt.Println("TARGET   ENGINE                STATUS")
	fmt.Println("esql     Elastic ES|QL         supported")
	fmt.Println("promql   Prometheus PromQL     supported (time-series aggregations only)")
	return 0
}

func readQuery(file string, positional []string) (string, error) {
	if file != "" {
		if file == "-" {
			return readStdin()
		}
		b, err := os.ReadFile(file)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	if len(positional) > 0 {
		return strings.Join(positional, " "), nil
	}
	return readStdin()
}

func readStdin() (string, error) {
	info, err := os.Stdin.Stat()
	if err == nil && info.Mode()&os.ModeCharDevice != 0 {
		return "", errors.New("no query provided; pass a query, use -f, or pipe one on stdin")
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	src := string(b)
	if strings.TrimSpace(src) == "" {
		return "", errors.New("no query provided")
	}
	return src, nil
}

// supportedSPLVersion lists the dialects the CLI accepts. Zero is the
// compiler's "default" spelling for classic SPL, but the flag is explicit about
// its value, so only 1 and 2 are accepted here.
func supportedSPLVersion(v int) bool {
	return v == compiler.SPLVersion1 || v == compiler.SPLVersion2
}

func hasWarnings(ds []diag.Diagnostic) bool {
	for _, d := range ds {
		if d.Severity == diag.SeverityWarning {
			return true
		}
	}
	return false
}

// splitFlags lets flags appear before or after the query. Go's flag package
// stops at the first positional argument, which would otherwise force users to
// remember flag ordering.
func splitFlags(args []string) (flagArgs, positional []string) {
	takesValue := map[string]bool{
		"target": true, "t": true,
		"mapping":     true,
		"index":       true,
		"metric":      true,
		"span":        true,
		"spl-version": true,
		"file":        true, "f": true,
	}
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			positional = append(positional, args[i+1:]...)
			break
		}
		if strings.HasPrefix(a, "-") && a != "-" {
			flagArgs = append(flagArgs, a)
			name := strings.TrimLeft(a, "-")
			if strings.Contains(name, "=") {
				continue
			}
			if takesValue[name] && i+1 < len(args) && (!strings.HasPrefix(args[i+1], "-") || args[i+1] == "-") {
				flagArgs = append(flagArgs, args[i+1])
				i++
			}
			continue
		}
		positional = append(positional, a)
	}
	return flagArgs, positional
}
