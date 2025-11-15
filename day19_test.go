package adventofcode2019

import (
	"testing"
)

func TestDay19Part1(t *testing.T) {
	testBytes(t, 19, filename, true, Day19, 160)
}

func TestDay19Part2(t *testing.T) {
	testBytes(t, 19, filename, false, Day19, 21062710)
}

func BenchmarkDay19Part1(b *testing.B) {
	benchBytes(b, 19, true, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchBytes(b, 19, false, Day19)
}
