package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	want := 315
	prog, err := os.ReadFile(input(13))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13Part1(Day5, MustSplit(string(prog)))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	want := 16171
	prog, err := os.ReadFile(input(13))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13Part2(Day5, MustSplit(string(prog)))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
