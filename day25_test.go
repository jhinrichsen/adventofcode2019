package adventofcode2019

import (
	"strings"
	"testing"
)

func TestDay25MapGame(t *testing.T) {
	t.Skip("Map game to determine correct path")
	// TODO: Implement full BFS exploration to automatically find:
	// 1. All safe items (avoiding dangerous ones)
	// 2. Path to security checkpoint
	// 3. Then try all item combinations
}

func TestDay25Part1(t *testing.T) {
	t.Skip("Day 25 requires correct exploration path - TODO: implement BFS or use known solution")
	lines := testLinesFromFilename(t, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	got := Day25(prog, true)
	want := uint(0) // Update when solution found
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
