package adventofcode2019

import (
	"fmt"
	"testing"
)

var part1Tests = []struct {
	mass int
	fuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestDay1Part1Examples(t *testing.T) {
	for _, tt := range part1Tests {
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
