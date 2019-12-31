package adventofcode2019

import (
	"bytes"
	"fmt"
	"testing"
)

var digitExamples = []struct {
	in  int
	out []byte
}{
	{Lower, []byte{1, 3, 6, 8, 1, 8}},
	{Upper, []byte{6, 8, 5, 9, 7, 9}},
}

var day4ExamplesPart1 = []struct {
	in  int
	out bool
}{
	{111111, true},
	{223450, false},
	{123789, false},
}

var day4ExamplesPart2 = []struct {
	in  int
	out bool
}{
	{112233, true},
	{123444, false},
	{111122, true},
	{111223, true},
	{144456, false},
}

func TestDigits(t *testing.T) {
	for _, tt := range digitExamples {
		id := fmt.Sprintf("%d", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := Digits(tt.in)
			if !bytes.Equal(want, got) {
				t.Fatalf("%s: want %v but got %v", id,
					want, got)
			}
		})
	}
}

func TestDay4Part1Examples(t *testing.T) {
	// range criteria must not be used for examples
	crits := []Criteria{
		CritSixDigits,
		CritIncreasing,
		CritTwoOrMoreAdjacent,
	}
	for _, tt := range day4ExamplesPart1 {
		id := fmt.Sprintf("%d", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := MeetsCriteria(tt.in, Digits(tt.in), crits)
			if want != got {
				t.Fatalf("%s: want %v but got %v", id,
					want, got)
			}
		})
	}
}

func TestDay4Part1(t *testing.T) {
	want := 1919
	got := Day4Part1()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay4Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day4Part1()
	}
}

func TestDay4Part2Examples(t *testing.T) {
	// range criteria must not be used for examples
	crits := []Criteria{
		CritSixDigits,
		CritIncreasing,
		CritExactlyTwoAdjacent,
	}
	for _, tt := range day4ExamplesPart2 {
		id := fmt.Sprintf("%d", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got := MeetsCriteria(tt.in, Digits(tt.in), crits)
			if want != got {
				t.Fatalf("%s: want %v but got %v", id,
					want, got)
			}
		})
	}
}

func TestDay4Part2(t *testing.T) {
	want := 1291
	got := Day4Part2()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay4Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day4Part2()
	}
}
