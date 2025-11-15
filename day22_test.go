package adventofcode2019

import (
	"fmt"
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
	tests := []struct {
		name         string
		filenameFunc func(uint8) string
		want         []uint
	}{
		{
			name:         "example1",
			filenameFunc: example1Filename,
			want:         []uint{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			name:         "example2",
			filenameFunc: example2Filename,
			want:         []uint{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			name:         "example3",
			filenameFunc: example3Filename,
			want:         []uint{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			name:         "example4",
			filenameFunc: example4Filename,
			want:         []uint{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := testLinesFromFilename(t, tt.filenameFunc(day22))
			deckSize := int64(10)

			// Use findCardAtPosition to verify we get the same result
			for pos := int64(0); pos < deckSize; pos++ {
				card := findCardAtPosition(lines, deckSize, 1, pos)
				if uint(card) != tt.want[pos] {
					t.Errorf("at position %d: want card %d but got %d", pos, tt.want[pos], card)
				}
			}
		})
	}
}

// TestDay22MultipleShuffles tests that the power transform works correctly
// by comparing against simulated multiple shuffles for a small deck.
func TestDay22MultipleShuffles(t *testing.T) {
	lines := testLinesFromFilename(t, example1Filename(day22))
	deckSize := int64(10)

	// Test with different numbers of shuffles
	for _, times := range []int64{1, 2, 3, 5, 10} {
		t.Run(fmt.Sprintf("shuffles=%d", times), func(t *testing.T) {
			// Simulate by tracking each card through multiple shuffles
			deck := make([]uint, deckSize)
			for cardNum := int64(0); cardNum < deckSize; cardNum++ {
				pos := uint(cardNum)
				// Apply shuffle 'times' times
				for range times {
					pos = trackCard(lines, uint(deckSize), pos)
				}
				// deck[pos] tells us which card is at position pos
				deck[pos] = uint(cardNum)
			}

			// Now verify our findCardAtPosition gives the same result
			for pos := int64(0); pos < deckSize; pos++ {
				card := findCardAtPosition(lines, deckSize, times, pos)
				if uint(card) != deck[pos] {
					t.Errorf("shuffle %d, pos %d: want card %d but got %d", times, pos, deck[pos], card)
				}
			}
		})
	}
}

func TestDay22Part2(t *testing.T) {
	testLines(t, day22, filename, false, Day22, 58348342289943)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchLines(b, day22, false, Day22)
}
