package adventofcode2019

import "testing"

func TestDay07Part1(t *testing.T) {
	testSolver(t, 7, filename, true, Day07, uint(24405))
}

func TestDay07Part2(t *testing.T) {
	testSolver(t, 7, filename, false, Day07, uint(8271623))
}

func BenchmarkDay07Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 7)
	for b.Loop() {
		_, _ = Day07(buf, true)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 7)
	for b.Loop() {
		_, _ = Day07(buf, false)
	}
}
