package spl1

import "testing"

// benchmarkQuery exercises a realistic pipeline: implicit search, a time
// bound, an aggregation, arithmetic, and row shaping.
const benchmarkQuery = `index=web earliest=-24h (status=500 OR status=503) host=api* ` +
	`| stats count by host ` +
	`| eval kb = bytes / 1024 ` +
	`| sort limit=10 -kb | head 5`

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := Parse(benchmarkQuery); err != nil {
			b.Fatal(err)
		}
	}
}
