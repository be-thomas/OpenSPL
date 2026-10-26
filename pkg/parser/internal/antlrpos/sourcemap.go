// Package antlrpos converts ANTLR token positions into the IR's byte-based
// positions.
//
// ANTLR's Go runtime indexes the input as a []rune, while pkg/ast reports
// 1-based lines, 1-based byte columns, and 0-based byte offsets. This package
// owns that conversion so the parser front-ends never do position arithmetic.
package antlrpos

import (
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"

	"github.com/be-thomas/openspl/pkg/ast"
)

// SourceMap maps rune indexes in a source string to byte offsets and lines.
//
// It is built once per parse. Invalid UTF-8 bytes decode as one U+FFFD rune
// each, exactly as ANTLR's input stream does, so every token start and stop
// index has a byte offset.
type SourceMap struct {
	src string

	// runeByte[r] is the byte offset of rune index r; it has one extra entry
	// holding len(src).
	runeByte []int

	// lineRune[i] is the rune index at which line i+1 starts, and lineByte
	// the matching byte offset. Line 0 starts at 0.
	lineRune []int
	lineByte []int
}

// New builds a source map for src.
func New(src string) *SourceMap {
	m := &SourceMap{src: src}
	m.runeByte = make([]int, 0, len(src)+1)
	m.lineRune = []int{0}
	m.lineByte = []int{0}

	for i := 0; i < len(src); {
		m.runeByte = append(m.runeByte, i)
		r, size := utf8.DecodeRuneInString(src[i:])
		if r == '\n' {
			m.lineRune = append(m.lineRune, len(m.runeByte))
			m.lineByte = append(m.lineByte, i+size)
		}
		i += size
	}
	m.runeByte = append(m.runeByte, len(src))
	return m
}

// RuneCount is the number of runes in the source.
func (m *SourceMap) RuneCount() int { return len(m.runeByte) - 1 }

// ByteOffset converts a rune index to a byte offset, clamping to the source.
func (m *SourceMap) ByteOffset(runeIndex int) int {
	if runeIndex <= 0 {
		return 0
	}
	if runeIndex >= len(m.runeByte) {
		return len(m.src)
	}
	return m.runeByte[runeIndex]
}

// Pos returns the IR position of a token's first character. The zero token and
// a nil token map to 1:1.
func (m *SourceMap) Pos(t antlr.Token) ast.Pos {
	if t == nil {
		return ast.Pos{Line: 1, Col: 1}
	}
	return m.PosAtOffset(m.ByteOffset(t.GetStart()))
}

// PosAtOffset returns the position of a byte offset.
func (m *SourceMap) PosAtOffset(offset int) ast.Pos {
	if offset < 0 {
		offset = 0
	}
	if offset > len(m.src) {
		offset = len(m.src)
	}
	line := 1
	for i := 1; i < len(m.lineByte); i++ {
		if m.lineByte[i] > offset {
			break
		}
		line = i + 1
	}
	col := offset - m.lineByte[line-1] + 1
	return ast.Pos{Line: line, Col: col, Offset: offset}
}

// PosAt converts a 1-based line and 0-based rune column, the coordinate pair an
// ANTLR error listener receives, into an IR position. Out-of-range coordinates
// clamp to the nearest valid position so a diagnostic never reports 0:0.
func (m *SourceMap) PosAt(line, col int) ast.Pos {
	if line < 1 {
		line = 1
	}
	if line > len(m.lineByte) {
		line = len(m.lineByte)
	}
	runeIndex := m.lineRune[line-1] + col
	if col < 0 {
		runeIndex = m.lineRune[line-1]
	}
	if runeIndex > m.RuneCount() {
		runeIndex = m.RuneCount()
	}
	return m.PosAtOffset(m.ByteOffset(runeIndex))
}

// Slice returns the source text covered by the run of tokens from first to
// last, inclusive. A nil or empty range yields "".
func (m *SourceMap) Slice(first, last antlr.Token) string {
	if first == nil || last == nil {
		return ""
	}
	start := m.ByteOffset(first.GetStart())
	stop := m.ByteOffset(last.GetStop() + 1)
	if stop < start || stop > len(m.src) {
		stop = len(m.src)
	}
	if start > stop {
		return ""
	}
	return m.src[start:stop]
}

// Adjacent reports whether b starts exactly where a ends.
func (m *SourceMap) Adjacent(a, b antlr.Token) bool {
	if a == nil || b == nil {
		return false
	}
	return a.GetStop()+1 == b.GetStart()
}

// Source returns the mapped source text.
func (m *SourceMap) Source() string { return m.src }
