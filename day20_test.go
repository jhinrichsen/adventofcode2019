package adventofcode2019

import (
	"testing"
)

func TestDay20Part1Example1(t *testing.T) {
	testBytes(t, 20, example1Filename, true, Day20, 23)
}

func TestDay20Part1Example2(t *testing.T) {
	testBytes(t, 20, example2Filename, true, Day20, 58)
}

func TestDay20Part1(t *testing.T) {
	buf := fileFromFilename(t, filename, 20)
	got := Day20(buf, true)
	t.Logf("Day 20 Part 1 result: %d", got)
	// TODO: Update with expected value once verified
}

func BenchmarkDay20Part1(b *testing.B) {
	benchBytes(b, 20, true, Day20)
}
