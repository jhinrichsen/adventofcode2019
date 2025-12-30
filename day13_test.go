package adventofcode2019

import "testing"

func TestDay13Part1(t *testing.T) {
	testSolver(t, 13, filename, true, Day13, uint(315))
}

func TestDay13Part2(t *testing.T) {
	testSolver(t, 13, filename, false, Day13, uint(16171))
}

func BenchmarkDay13Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 13)
	for b.Loop() {
		_, _ = Day13(buf, true)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 13)
	for b.Loop() {
		_, _ = Day13(buf, false)
	}
}
