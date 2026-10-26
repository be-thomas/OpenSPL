package spl2

import "testing"

// benchmarkQuery exercises every stage of the subset: a dataset, a compound
// condition, and a grouped aggregation with an alias.
const benchmarkQuery = `from logs-2024-01 | where status=500 AND host="web" ` +
	`| stats count() AS failures, dc(user) AS users BY host, status`

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := Parse(benchmarkQuery); err != nil {
			b.Fatal(err)
		}
	}
}
