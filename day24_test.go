package adventofcode2019

import (
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	lines := testLinesFromFilename(t, exampleFilename(24))
	got := Day24(lines, true)
	want := uint(2129920)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	testLines(t, 24, filename, true, Day24, 20751345)
}

func TestDay24Part2Example(t *testing.T) {
	lines := testLinesFromFilename(t, exampleFilename(24))
	grid := parseGrid24(lines)
	got := day24Part2(grid, 10)
	want := uint(99)
	if got != want {
		t.Fatalf("want %d bugs after 10 minutes but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	testLines(t, 24, filename, false, Day24, 1983)
}

func BenchmarkDay24Part1(b *testing.B) {
	benchLines(b, 24, true, Day24)
}

func BenchmarkDay24Part2(b *testing.B) {
	benchLines(b, 24, false, Day24)
}
