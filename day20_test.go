package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay20Part1Example1(t *testing.T) {
	testBytes(t, 20, example1Filename, true, Day20, 23)
}

func TestDay20Part1Example2(t *testing.T) {
	testBytes(t, 20, example2Filename, true, Day20, 58)
}

func TestDay20Part1(t *testing.T) {
	testBytes(t, 20, filename, true, Day20, 638)
}

func TestDay20Part2Example(t *testing.T) {
	buf, err := os.ReadFile("testdata/day20_part2_example.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(buf, false)
	const want = 396
	if got != want {
		t.Fatalf("Part 2 example: want %v but got %v", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	testBytes(t, 20, filename, false, Day20, 7844)
}

func BenchmarkDay20Part1(b *testing.B) {
	benchBytes(b, 20, true, Day20)
}

func BenchmarkDay20Part2(b *testing.B) {
	benchBytes(b, 20, false, Day20)
}
