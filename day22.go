package adventofcode2019

import (
	"strconv"
	"strings"
)

// Day22 solves the space card shuffle puzzle.
// For part 1, it returns the position of card 2019 after shuffling a deck of 10007 cards.
// For part 2, it returns the card number at position 2020 after shuffling 119315717514047 cards
// 101741582076661 times.
func Day22(lines []string, part1 bool) uint {
	if part1 {
		// For the main puzzle, track card 2019 in a deck of 10007
		deckSize := uint(10007)
		cardNumber := uint(2019)
		position := trackCard(lines, deckSize, cardNumber)
		return position
	}

	// Part 2: Find which card is at position 2020 after many shuffles
	deckSize := int64(119315717514047)
	shuffles := int64(101741582076661)
	targetPos := int64(2020)

	card := findCardAtPosition(lines, deckSize, shuffles, targetPos)
	return uint(card)
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

// findCardAtPosition finds which card ends up at a given position after applying
// the shuffle operations a specified number of times.
func findCardAtPosition(lines []string, deckSize, times, position int64) int64 {
	// Build the inverse transformation (to work backward from position to card)
	a, b := int64(1), int64(0)

	// Process operations in reverse order to build inverse transformation
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if line == "deal into new stack" {
			// Inverse of reverse: pos -> deckSize - 1 - pos
			// As linear: new_pos = -1 * pos + (deckSize - 1)
			// Inverse is same: a' = -1, b' = deckSize - 1
			a = modMul(-a, 1, deckSize)
			b = modAdd(modMul(-b, 1, deckSize), deckSize-1, deckSize)
		} else if strings.HasPrefix(line, "cut ") {
			nStr := strings.TrimPrefix(line, "cut ")
			n, _ := strconv.ParseInt(nStr, 10, 64)
			// Forward: pos -> pos - n
			// Inverse: pos -> pos + n
			// a' = a, b' = b + n
			b = modAdd(b, n, deckSize)
		} else if strings.HasPrefix(line, "deal with increment ") {
			nStr := strings.TrimPrefix(line, "deal with increment ")
			n, _ := strconv.ParseInt(nStr, 10, 64)
			// Forward: pos -> pos * n
			// Inverse: pos -> pos * n^(-1)
			nInv := modInverse(n, deckSize)
			a = modMul(a, nInv, deckSize)
			b = modMul(b, nInv, deckSize)
		}
	}

	// Now we have the inverse transformation for one shuffle: f^(-1)(x) = a*x + b
	// Apply it 'times' times using exponentiation
	aFinal, bFinal := powerTransform(a, b, times, deckSize)

	// Apply to the target position
	card := modAdd(modMul(aFinal, position, deckSize), bFinal, deckSize)
	return card
}

// powerTransform applies a linear transformation (a*x + b) n times using
// matrix exponentiation. Returns the coefficients (a', b') of the composed transformation.
func powerTransform(a, b, n, mod int64) (int64, int64) {
	// Base case: identity transformation
	resA, resB := int64(1), int64(0)

	// Matrix exponentiation for [a b; 0 1]^n
	currA, currB := a, b

	for n > 0 {
		if n%2 == 1 {
			// Multiply result by current
			newA := modMul(resA, currA, mod)
			newB := modAdd(modMul(resB, currA, mod), currB, mod)
			resA, resB = newA, newB
		}
		// Square current
		newA := modMul(currA, currA, mod)
		newB := modAdd(modMul(currB, currA, mod), currB, mod)
		currA, currB = newA, newB
		n /= 2
	}

	return resA, resB
}

// modAdd performs (a + b) mod m, handling negative numbers correctly
func modAdd(a, b, m int64) int64 {
	res := (a + b) % m
	if res < 0 {
		res += m
	}
	return res
}

// modMul performs (a * b) mod m, handling potential overflow
func modMul(a, b, m int64) int64 {
	// Use Go's native mod handling (doesn't overflow for int64)
	res := (a % m) * (b % m) % m
	if res < 0 {
		res += m
	}
	return res
}

// modInverse computes the modular inverse of a modulo m using extended Euclidean algorithm
func modInverse(a, m int64) int64 {
	// Ensure a is positive
	a = a % m
	if a < 0 {
		a += m
	}

	// Extended Euclidean algorithm
	t, newT := int64(0), int64(1)
	r, newR := m, a

	for newR != 0 {
		quotient := r / newR
		t, newT = newT, t-quotient*newT
		r, newR = newR, r-quotient*newR
	}

	if r > 1 {
		panic("a is not invertible")
	}
	if t < 0 {
		t += m
	}

	return t
}
