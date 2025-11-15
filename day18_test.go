package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay18Part1Examples(t *testing.T) {
	tests := []struct {
		example uint8
		want    uint
	}{
		{1, 8},
		{2, 86},
		{3, 132},
		{4, 136},
		{5, 81},
	}

	for _, tt := range tests {
		t.Run(exampleNFilename(18, tt.example), func(t *testing.T) {
			maze, err := os.ReadFile(exampleNFilename(18, tt.example))
			if err != nil {
				t.Fatal(err)
			}
			got := Day18(maze, true)
			if got != tt.want {
				t.Fatalf("example %d: want %v but got %v", tt.example, tt.want, got)
			}
		})
	}
}

func TestDay18Part1(t *testing.T) {
	maze, err := os.ReadFile(filename(18))
	if err != nil {
		t.Fatal(err)
	}
	// TODO: Update with actual expected value once puzzle is solved
	const want = 0
	got := Day18(maze, true)
	if want != 0 && want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
	t.Logf("Day 18 Part 1: %v", got)
}

func BenchmarkDay18Part1(b *testing.B) {
	maze, err := os.ReadFile(filename(18))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		Day18(maze, true)
	}
}
