package adventofcode2019

import "testing"

func TestDay23Part1(t *testing.T) {
	testSolver(t, 23, filename, true, Day23, uint(19530))
}

func TestDay23Part2(t *testing.T) {
	testSolver(t, 23, filename, false, Day23, uint(12725))
}

func BenchmarkDay23Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 23)
	for b.Loop() {
		_, _ = Day23(buf, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 23)
	for b.Loop() {
		_, _ = Day23(buf, false)
	}
}
