package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay01Part1Examples(t *testing.T) {
	tests := []struct {
		mass uint
		fuel uint
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("mass=%d", tt.mass), func(t *testing.T) {
			input := fmt.Appendf(nil, "%d\n", tt.mass)
			got, err := Day01(input, true)
			if err != nil {
				t.Fatal(err)
			}
			if tt.fuel != got {
				t.Fatalf("want %d but got %d", tt.fuel, got)
			}
		})
	}
}

func TestDay01Part1(t *testing.T) {
	testSolver(t, 1, filename, true, Day01, uint(3231195))
}

func TestDay01Part2Examples(t *testing.T) {
	tests := []struct {
		mass uint
		fuel uint
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("mass=%d", tt.mass), func(t *testing.T) {
			input := fmt.Appendf(nil, "%d\n", tt.mass)
			got, err := Day01(input, false)
			if err != nil {
				t.Fatal(err)
			}
			if tt.fuel != got {
				t.Fatalf("want %d but got %d", tt.fuel, got)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	testSolver(t, 1, filename, false, Day01, uint(4843929))
}

func BenchmarkDay01Part1(b *testing.B) {
	benchSolver(b, 1, true, Day01)
}

func BenchmarkDay01Part2(b *testing.B) {
	benchSolver(b, 1, false, Day01)
}
