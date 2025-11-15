package adventofcode2019

import (
	"fmt"
	"slices"
	"testing"
)

func TestDay22Part1Examples(t *testing.T) {
	tests := []struct {
		filenameFunc func(uint8) string
		want         []uint
	}{
		{
			filenameFunc: example1Filename,
			want:         []uint{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			filenameFunc: example2Filename,
			want:         []uint{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			filenameFunc: example3Filename,
			want:         []uint{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			filenameFunc: example4Filename,
			want:         []uint{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			lines := testLinesFromFilename(t, tt.filenameFunc(22))
			deck := shuffleDeck(lines, 10)
			if !slices.Equal(deck, tt.want) {
				t.Fatalf("want %v but got %v", tt.want, deck)
			}
		})
	}
}

func TestDay22Part1(t *testing.T) {
	testLines(t, 22, filename, true, Day22, 6289)
}

func BenchmarkDay22Part1(b *testing.B) {
	benchLines(b, 22, true, Day22)
}

// TestDay22Part2InverseLogic verifies the inverse transformation logic works correctly
// by checking that we can find which card is at each position for a small deck.
func TestDay22Part2InverseLogic(t *testing.T) {
	tests := []struct {
		filenameFunc func(uint8) string
		want         []uint
	}{
		{
			filenameFunc: example1Filename,
			want:         []uint{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			filenameFunc: example2Filename,
			want:         []uint{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			filenameFunc: example3Filename,
			want:         []uint{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			filenameFunc: example4Filename,
			want:         []uint{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			lines := testLinesFromFilename(t, tt.filenameFunc(22))
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

// TestDay22Part2MultipleShuffles tests that the power transform works correctly
// by comparing against simulated multiple shuffles for a small deck.
func TestDay22Part2MultipleShuffles(t *testing.T) {
	lines := testLinesFromFilename(t, example1Filename(22))
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
	testLines(t, 22, filename, false, Day22, 58348342289943)
}

func BenchmarkDay22Part2(b *testing.B) {
	benchLines(b, 22, false, Day22)
}
