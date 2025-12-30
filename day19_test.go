package adventofcode2019

import "testing"

func TestDay19Part1(t *testing.T) {
	testSolver(t, 19, filename, true, Day19, uint(160))
}

func TestDay19Part2(t *testing.T) {
	testSolver(t, 19, filename, false, Day19, uint(9441282))
}

func BenchmarkDay19Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 19)
	for b.Loop() {
		_, _ = Day19(buf, true)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 19)
	for b.Loop() {
		_, _ = Day19(buf, false)
	}
}
