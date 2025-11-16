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

func TestDay1Part1(t *testing.T) {
	lines, err := linesFromFilename(input(1))
	if err != nil {
		t.Fatal(err)
	}
	want := 3231195
	got, err := Day1Part1(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

var part2Tests = []struct {
	mass int
	fuel int
}{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestDay1Part2Examples(t *testing.T) {
	for _, tt := range part2Tests {
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

func TestDay1Part2(t *testing.T) {
	lines, err := linesFromFilename(input(1))
	if err != nil {
		t.Fatal(err)
	}
	want := 4843929
	got, err := Day1Part2(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(1))
	for b.Loop() {
		_, _ = Day1Part1(lines)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
	lines := testLinesFromFilename(b, filename(1))
	for b.Loop() {
		_, _ = Day1Part2(lines)
	}
}
