package adventofcode2019

import (
	"strconv"
	"strings"
)

// Day22 solves the space card shuffle puzzle.
// For part 1, it returns the position of card 2019 after shuffling a deck of 10007 cards.
func Day22(lines []string, part1 bool) uint {
	if !part1 {
		panic("part 2 not implemented")
	}

	// For the main puzzle, track card 2019 in a deck of 10007
	deckSize := uint(10007)
	cardNumber := uint(2019)

	position := trackCard(lines, deckSize, cardNumber)
	return position
}

// trackCard follows a specific card through all shuffle operations
// and returns its final position in the deck.
func trackCard(lines []string, deckSize, cardNumber uint) uint {
	// Card starts at position equal to its number (factory order)
	pos := cardNumber

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "deal into new stack" {
			// Reverse: position becomes (deckSize - 1 - position)
			pos = deckSize - 1 - pos
		} else if strings.HasPrefix(line, "cut ") {
			// Cut N cards
			nStr := strings.TrimPrefix(line, "cut ")
			n, _ := strconv.Atoi(nStr)
			// Position shifts by -N (with wrapping)
			if n >= 0 {
				pos = (pos - uint(n) + deckSize) % deckSize
			} else {
				pos = (pos + uint(-n)) % deckSize
			}
		} else if strings.HasPrefix(line, "deal with increment ") {
			// Deal with increment N
			nStr := strings.TrimPrefix(line, "deal with increment ")
			n, _ := strconv.Atoi(nStr)
			// Position multiplies by N (mod deckSize)
			pos = (pos * uint(n)) % deckSize
		}
	}

	return pos
}

// shuffleDeck performs all shuffle operations and returns the final deck state.
// This is used for testing with small decks.
func shuffleDeck(lines []string, deckSize uint) []uint {
	deck := make([]uint, deckSize)
	for i := range deckSize {
		deck[i] = i
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line == "deal into new stack" {
			// Reverse the deck
			for i, j := uint(0), deckSize-1; i < j; i, j = i+1, j-1 {
				deck[i], deck[j] = deck[j], deck[i]
			}
		} else if strings.HasPrefix(line, "cut ") {
			// Cut N cards
			nStr := strings.TrimPrefix(line, "cut ")
			n, _ := strconv.Atoi(nStr)

			if n >= 0 {
				// Cut from top
				cut := uint(n) % deckSize
				newDeck := make([]uint, deckSize)
				copy(newDeck, deck[cut:])
				copy(newDeck[deckSize-cut:], deck[:cut])
				deck = newDeck
			} else {
				// Cut from bottom
				cut := uint(-n) % deckSize
				newDeck := make([]uint, deckSize)
				copy(newDeck, deck[deckSize-cut:])
				copy(newDeck[cut:], deck[:deckSize-cut])
				deck = newDeck
			}
		} else if strings.HasPrefix(line, "deal with increment ") {
			// Deal with increment N
			nStr := strings.TrimPrefix(line, "deal with increment ")
			n, _ := strconv.Atoi(nStr)
			increment := uint(n)

			newDeck := make([]uint, deckSize)
			pos := uint(0)
			for i := range deckSize {
				newDeck[pos] = deck[i]
				pos = (pos + increment) % deckSize
			}
			deck = newDeck
		}
	}

	return deck
}
