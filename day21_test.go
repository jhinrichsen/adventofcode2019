package adventofcode2019

import (
	"testing"
)

// Note: Day 21 has no examples in the puzzle description.
// The puzzle requires figuring out the springscript logic.

func TestDay21Part1(t *testing.T) {
	testBytes(t, 21, filename, true, Day21, 19352493)
}

func TestDay21Part2(t *testing.T) {
	testBytes(t, 21, filename, false, Day21, 1141896219)
}

func BenchmarkDay21Part1(b *testing.B) {
	benchBytes(b, 21, true, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchBytes(b, 21, false, Day21)
}
