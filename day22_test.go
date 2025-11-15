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

// TestDay22InverseLogic verifies the inverse transformation logic works correctly
// by checking that we can find which card is at each position for a small deck.
func TestDay22InverseLogic(t *testing.T) {
	lines := testLinesFromFilename(t, example1Filename(day22))
	deckSize := int64(10)

	// First, get the final deck configuration
	want := []uint{0, 3, 6, 9, 2, 5, 8, 1, 4, 7}

	// Now use findCardAtPosition to verify we get the same result
	for pos := int64(0); pos < deckSize; pos++ {
		card := findCardAtPosition(lines, deckSize, 1, pos)
		if uint(card) != want[pos] {
			t.Errorf("at position %d: want card %d but got %d", pos, want[pos], card)
		}
	}
}

func TestDay22Part2(t *testing.T) {
	testLines(t, day22, filename, false, Day22, 69676926565412)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchLines(b, day22, false, Day22)
}
