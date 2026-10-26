// Package lexerparity holds the test that keeps the lexical rules duplicated in
// SPL1.g4 and SPL2.g4 from drifting apart.
//
// The Go target's grammar-import output crosses package boundaries awkwardly,
// so each combined grammar declares its own identifier, number, duration,
// string, operator, whitespace, and UNKNOWN rules. The two front-ends must
// tokenize a shared corpus identically, because a difference here would change
// expressions or dataset names between dialects for reasons that have nothing
// to do with either dialect's semantics.
//
// The package has no API and no non-test files beyond this one.
package lexerparity
