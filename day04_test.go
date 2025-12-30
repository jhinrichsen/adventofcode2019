package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay04Part1Examples(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
	}
	crits := []criteria{critSixDigits, critIncreasing, critTwoOrMoreAdjacent}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := meetsCriteria(tt.in, digits(tt.in), crits)
			if tt.out != got {
				t.Fatalf("want %v but got %v", tt.out, got)
			}
		})
	}
}

func TestDay04Part1(t *testing.T) {
	want := uint(1919)
	got := Day04(true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	for b.Loop() {
		_ = Day04(true)
	}
}

func TestDay04Part2Examples(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
		{111223, true},
		{144456, false},
	}
	crits := []criteria{critSixDigits, critIncreasing, critExactlyTwoAdjacent}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := meetsCriteria(tt.in, digits(tt.in), crits)
			if tt.out != got {
				t.Fatalf("want %v but got %v", tt.out, got)
			}
		})
	}
}

func TestDay04Part2(t *testing.T) {
	want := uint(1291)
	got := Day04(false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	for b.Loop() {
		_ = Day04(false)
	}
}
