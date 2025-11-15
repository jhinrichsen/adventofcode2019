package adventofcode2019

import (
	"testing"
)

const day25Part1Want = 229384

func TestDay25Part1(t *testing.T) {
	testWithParser(t, 25, filename, true, NewDay25, Day25, day25Part1Want)
}

func BenchmarkDay25Part1(b *testing.B) {
	benchWithParser(b, 25, true, NewDay25, Day25)
}
