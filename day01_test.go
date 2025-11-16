package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay01Part1Examples(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, tt := range tests {
		id := fmt.Sprintf("Fuel(%d)", tt.mass)
		t.Run(id, func(t *testing.T) {
			want := tt.fuel
			got := Fuel(tt.mass)
			if want != got {
				t.Fatalf("%q: want %d but got %d", id,
					want, got)
			}
		})
	}
}

func TestDay01Part1(t *testing.T) {
	testLines(t, 1, filename, true, Day01, uint(3231195))
}

func TestDay01Part2Examples(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}
	for _, tt := range tests {
		id := fmt.Sprintf("Fuel(%d)", tt.mass)
		t.Run(id, func(t *testing.T) {
			want := tt.fuel
			got := CompleteFuel(tt.mass)
			if want != got {
				t.Fatalf("%q: want %d but got %d", id,
					want, got)
			}
		})
	}
}

func TestDay01Part2(t *testing.T) {
	testLines(t, 1, filename, false, Day01, uint(4843929))
}

func BenchmarkDay01Part1(b *testing.B) {
	benchLines(b, 1, true, Day01)
}

func BenchmarkDay01Part2(b *testing.B) {
	benchLines(b, 1, false, Day01)
}
