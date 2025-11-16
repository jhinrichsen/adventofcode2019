package adventofcode2019

import (
	"testing"
)

func TestDay25Part1(t *testing.T) {
	testLines(t, 25, filename, true, Day25, 229384)
}

func BenchmarkDay25Part1(b *testing.B) {
	benchLines(b, 25, true, Day25)
}
