package adventofcode2019

import "testing"

func TestDay15Part1(t *testing.T) {
	testSolver(t, 15, filename, true, Day15, uint(272))
}

func TestDay15Part2(t *testing.T) {
	testSolver(t, 15, filename, false, Day15, uint(398))
}

func BenchmarkDay15Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 15)
	for b.Loop() {
		_, _ = Day15(buf, true)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 15)
	for b.Loop() {
		_, _ = Day15(buf, false)
	}
}
