package adventofcode2019

import (
	"testing"
)

func TestDay25Part1(t *testing.T) {
	testWithParser(t, 25, filename, true, NewDay25, Day25, 229384)
}

func BenchmarkDay25Part1(b *testing.B) {
	benchWithParser(b, 25, true, NewDay25, Day25)
}
