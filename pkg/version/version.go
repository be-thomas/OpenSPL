// Package version exposes build metadata set via -ldflags.
package version

import "fmt"

var (
	// Version is the semantic version, set at build time.
	Version = "0.1.0-dev"
	// Commit is the git SHA, set at build time.
	Commit = "none"
	// Date is the build date, set at build time.
	Date = "unknown"
)

// String renders a single-line version string.
func String() string {
	return fmt.Sprintf("openspl %s (commit %s, built %s)", Version, Commit, Date)
}
