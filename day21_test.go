package adventofcode2019

import (
	"testing"
)

// Note: Day 21 has no examples in the puzzle description.
// The puzzle requires figuring out the springscript logic.

func TestDay21Part1(t *testing.T) {
	buf := fileFromFilename(t, filename, 21)
	got := Day21(buf, true)
	t.Logf("Day 21 Part 1 hull damage: %d", got)
	// TODO: Update with expected value once confirmed
}

func BenchmarkDay21Part1(b *testing.B) {
	benchBytes(b, 21, true, Day21)
}

func BenchmarkDay21Part2(b *testing.B) {
	benchBytes(b, 21, false, Day21)
}
