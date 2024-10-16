package adventofcode2019

import (
	"testing"
)

func TestLinesFromFilename(t *testing.T) {
	lines, err := linesFromFilename("testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 1 {
		t.Fatalf("want 1 line but got %d", len(lines))
	}
}

func TestLinesAsNumbers(t *testing.T) {
	sample := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	ints, err := linesAsNumbers(sample)
	if err != nil {
		t.Fatal(err)
	}
	if len(ints) != len(sample) {
		t.Fatalf("want %d numbers but got %d", len(sample), len(ints))
	}
	for i := range sample {
		want := i + 1 // entries are 1-based
		got := ints[i]
		if want != got {
			t.Fatalf("line %d: want %d but got %d", i, want, got)
		}
	}
}
