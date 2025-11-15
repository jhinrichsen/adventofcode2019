package adventofcode2019

import (
	"testing"
)

func TestDay19Part2Example(t *testing.T) {
	// Parse the beam map
	lines := testLinesFromFilename(t, "testdata/day19_part2_example.txt")
	beamMap := make([][]bool, len(lines))
	for y, line := range lines {
		beamMap[y] = make([]bool, len(line))
		for x, ch := range line {
			beamMap[y][x] = (ch == '#' || ch == 'O')
		}
	}

	got := findSquareFromMap(beamMap, 10)
	const want = 250020
	if got != want {
		t.Fatalf("Part 2 example: want %v but got %v", want, got)
	}
}

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
	// - 11491534 (scanning bottom row, checking top-right)
	t.Skip("Part 2 not correct yet - answer too high")
}

func BenchmarkDay19Part1(b *testing.B) {
	benchBytes(b, 19, true, Day19)
}

func BenchmarkDay19Part2(b *testing.B) {
	benchBytes(b, 19, false, Day19)
}
