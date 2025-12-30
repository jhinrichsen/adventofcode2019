package adventofcode2019

import "testing"

// Note: Day 21 has no examples in the puzzle description.
// The puzzle requires figuring out the springscript logic.

func TestDay21Part1(t *testing.T) {
	testSolver(t, 21, filename, true, Day21, uint(19352493))
}

func TestDay21Part2(t *testing.T) {
	testSolver(t, 21, filename, false, Day21, uint(1141896219))
}

func BenchmarkDay21Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 21)
	for b.Loop() {
		_, _ = Day21(buf, true)
	}
}

func BenchmarkDay21Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 21)
	for b.Loop() {
		_, _ = Day21(buf, false)
	}
}
