package adventofcode2019

import "testing"

func TestDay17Part1(t *testing.T) {
	testSolver(t, 17, filename, true, Day17, uint(5972))
}

func TestDay17Part2(t *testing.T) {
	testSolver(t, 17, filename, false, Day17, uint(933214))
}

func BenchmarkDay17Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 17)
	for b.Loop() {
		_, _ = Day17(buf, true)
	}
}

func BenchmarkDay17Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 17)
	for b.Loop() {
		_, _ = Day17(buf, false)
	}
}
