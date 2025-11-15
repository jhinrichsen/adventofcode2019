package adventofcode2019

import (
	"testing"
)

func TestDay19Part1(t *testing.T) {
	testBytes(t, 19, filename, true, Day19, 160)
}

func TestDay19Part2(t *testing.T) {
	buf := fileFromFilename(t, filename, 19)
	got := Day19(buf, false)
	t.Logf("Day 19 Part 2 result: %d", got)
	// Previous attempts (all too high):
	// - 21062710 (checked wrong corners)
	// - 15362148 (checked top-right and bottom-left from top-left)
}

func BenchmarkDay19Part1(b *testing.B) {
	benchBytes(b, 19, true, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchBytes(b, 19, false, Day19)
}
