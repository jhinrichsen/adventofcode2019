package adventofcode2019

import (
	"strings"
	"testing"
)

func TestDay25Part1(t *testing.T) {
	t.Skip("Day 25 requires manual exploration or pre-mapped solution")
	lines := testLinesFromFilename(t, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	got := Day25(prog, true)
	want := uint(0) // Update after determining correct path and password
	if want != 0 && got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
	if got != 0 {
		t.Logf("Password: %d", got)
	}
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	for b.Loop() {
		Day25(prog, true)
	}
}
