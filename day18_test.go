package adventofcode2019

import (
	"fmt"
	"os"
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

func TestDay18Part2Examples(t *testing.T) {
	tests := []struct {
		filename string
		want     uint
	}{
		{"testdata/day18_part2_example1.txt", 8},
		{"testdata/day18_part2_example2.txt", 24},
		{"testdata/day18_part2_example3.txt", 32},
		{"testdata/day18_part2_example4.txt", 72},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			maze, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			got := Day18(maze, false)
			if got != tt.want {
				t.Fatalf("example %d: want %v but got %v", i+1, tt.want, got)
			}
		})
	}
}

func TestDay18Part2(t *testing.T) {
	buf := fileFromFilename(t, filename, 18)
	got := Day18(buf, false)
	t.Logf("Day 18 Part 2 result: %d", got)
	// TODO: Update with expected value once verified
}

func BenchmarkDay18Part2(b *testing.B) {
	benchBytes(b, 18, false, Day18)
}
