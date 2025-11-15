package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay18Part1Examples(t *testing.T) {
	wants := []uint{8, 86, 132, 136, 81}

	for i, want := range wants {
		example := uint8(i + 1)
		t.Run(fmt.Sprintf("example%d", example), func(t *testing.T) {
			testBytes(t, 18, func(d uint8) string { return exampleNFilename(d, example) }, true, Day18, want)
		})
	}
}

func TestDay18Part1(t *testing.T) {
	buf := fileFromFilename(t, filename, 18)
	got := Day18(buf, true)
	t.Logf("Day 18 Part 1 result: %d", got)
	// TODO: Correct answer is less than 3962 (answer was too high)
	// Need to debug why BFS is giving a longer path than optimal
}

func BenchmarkDay18Part1(b *testing.B) {
	benchBytes(b, 18, true, Day18)
}
