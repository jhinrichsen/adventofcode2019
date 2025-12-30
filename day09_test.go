package adventofcode2019

import "testing"

func TestDay09Part1(t *testing.T) {
	testSolver(t, 9, filename, true, Day09, uint(2436480432))
}

func TestDay09Part2(t *testing.T) {
	testSolver(t, 9, filename, false, Day09, uint(45710))
}

func BenchmarkDay09Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 9)
	for b.Loop() {
		_, _ = Day09(buf, true)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 9)
	for b.Loop() {
		_, _ = Day09(buf, false)
	}
}
