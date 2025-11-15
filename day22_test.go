package adventofcode2019

import (
	"slices"
	"testing"
)

const day22 = 22

func TestDay22Example1(t *testing.T) {
	lines := testLinesFromFilename(t, example1Filename(day22))
	deck := shuffleDeck(lines, 10)
	want := []uint{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}
	if !slices.Equal(deck, want) {
		t.Fatalf("want %v but got %v", want, deck)
	}
}

func TestDay22Example2(t *testing.T) {
	lines := testLinesFromFilename(t, example2Filename(day22))
	deck := shuffleDeck(lines, 10)
	want := []uint{3, 0, 7, 4, 1, 8, 5, 2, 9, 6}
	if !slices.Equal(deck, want) {
		t.Fatalf("want %v but got %v", want, deck)
	}
}

func TestDay22Example3(t *testing.T) {
	lines := testLinesFromFilename(t, example3Filename(day22))
	deck := shuffleDeck(lines, 10)
	want := []uint{6, 3, 0, 7, 4, 1, 8, 5, 2, 9}
	if !slices.Equal(deck, want) {
		t.Fatalf("want %v but got %v", want, deck)
	}
}

func TestDay22Example4(t *testing.T) {
	lines := testLinesFromFilename(t, example4Filename(day22))
	deck := shuffleDeck(lines, 10)
	want := []uint{9, 2, 5, 8, 1, 4, 7, 0, 3, 6}
	if !slices.Equal(deck, want) {
		t.Fatalf("want %v but got %v", want, deck)
	}
}

func TestDay22Part1(t *testing.T) {
	testLines(t, day22, filename, true, Day22, 6289)
}

func BenchmarkDay22Part1(b *testing.B) {
	benchLines(b, day22, true, Day22)
}
