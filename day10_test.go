package adventofcode2019

import (
	"fmt"
	"image"
	"testing"
)

var day10Examples = []struct {
	filenameFunc func(uint8) string
	best         image.Point
	bestCount    int
}{
	{example1Filename, image.Point{X: 3, Y: 4}, 8},
	{example2Filename, image.Point{X: 5, Y: 8}, 33},
	{example3Filename, image.Point{X: 1, Y: 2}, 35},
	{example4Filename, image.Point{X: 6, Y: 3}, 41},
	{example5Filename, image.Point{X: 11, Y: 13}, 210},
}

func TestDay10Example1(t *testing.T) {
	buf := fileFromFilename(t, example1Filename, 10)
	as := ParseAsteroidMap(buf)

	// Check number of asteroids
	if len(as) != 10 {
		t.Fatalf("want 10 but got %d", len(as))
	}
	second := image.Point{X: 4, Y: 0}
	if as[1] != second {
		t.Fatalf("expected asteroid %+v at index 1, got %+v",
			second, as[1])

	}
}

func TestDay10Part1Examples(t *testing.T) {
	for i, tt := range day10Examples {
		id := fmt.Sprintf("Day10Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			buf := fileFromFilename(t, tt.filenameFunc, 10)
			as := ParseAsteroidMap(buf)
			wantA, want := tt.best, tt.bestCount
			gotA, got := Day10Part1(as)
			if tt.best != gotA {
				t.Fatalf("%s: want %+v but got %+v",
					id, wantA, gotA)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay10Part1(t *testing.T) {
	buf := fileFromFilename(t, filename, 10)
	want := uint(267)
	got := Day10(buf, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 10)
	for b.Loop() {
		_ = Day10(buf, true)
	}
}

func TestDay10Part2Example(t *testing.T) {
	want := 802
	ex := day10Examples[4]
	buf := fileFromFilename(t, ex.filenameFunc, 10)
	as := ParseAsteroidMap(buf)
	got := Day10Part2(as, ex.best)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2(t *testing.T) {
	buf := fileFromFilename(t, filename, 10)
	want := uint(1309)
	got := Day10(buf, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 10)
	for b.Loop() {
		_ = Day10(buf, false)
	}
}
