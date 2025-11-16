package adventofcode2019

import (
	"testing"
)

func TestDay02Part1(t *testing.T) {
	testSolver(t, 2, filename, true, Day02, uint(3562624))
}

func TestDay02Part2(t *testing.T) {
	testSolver(t, 2, filename, false, Day02, uint(8298))
}

func BenchmarkDay02Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 2)
	for b.Loop() {
		_, _ = Day02(buf, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 2)
	for b.Loop() {
		_, _ = Day02(buf, false)
	}
}
