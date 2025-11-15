package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay18Part1Examples(t *testing.T) {
	tests := []struct {
		filename string
		want     uint
	}{
		{example1Filename(18), 8},
		{example2Filename(18), 86},
		{example3Filename(18), 132},
		{example4Filename(18), 136},
		{example5Filename(18), 81},
	}

	for i, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			maze, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			got := Day18(maze, true)
			if got != tt.want {
				t.Fatalf("example %d: want %v but got %v", i+1, tt.want, got)
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
