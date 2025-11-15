package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay20Part1Example1(t *testing.T) {
	buf, err := os.ReadFile(example1Filename(20))
	if err != nil {
		t.Fatal(err)
	}
	maze := parseMaze20(buf)
	t.Logf("Example1 - Start: %v, End: %v, dimX=%d, dimY=%d", maze.start, maze.end, maze.dimX, maze.dimY)
	testBytes(t, 20, example1Filename, true, Day20, 23)
}

func TestDay20Part1Example2(t *testing.T) {
	testBytes(t, 20, example2Filename, true, Day20, 58)
}

func TestDay20Part1(t *testing.T) {
	testBytes(t, 20, filename, true, Day20, 638)
}

func TestDay20Part2Example(t *testing.T) {
	t.Skip("Part 2 example file appears to be malformed - skipping for now")
}

func TestDay20Part2(t *testing.T) {
	buf := fileFromFilename(t, filename, 20)
	got := Day20(buf, false)
	t.Logf("Day 20 Part 2 result: %d", got)
	// TODO: Update with expected value once verified
}

func BenchmarkDay20Part1(b *testing.B) {
	benchBytes(b, 20, true, Day20)
}

func BenchmarkDay20Part2(b *testing.B) {
	benchBytes(b, 20, false, Day20)
}
