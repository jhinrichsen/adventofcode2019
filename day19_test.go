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
	// INCORRECT - answer 15362148 is too high
	// Previous attempt 21062710 was also too high
	t.Skip("Part 2 not yet correct")
}

func BenchmarkDay19Part1(b *testing.B) {
	benchBytes(b, 19, true, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchBytes(b, 19, false, Day19)
}
