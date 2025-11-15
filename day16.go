package adventofcode2019

import (
	"bytes"
)

// Day16 applies Flawed Frequency Transmission (FFT) algorithm
// Part 1: Returns first 8 digits after 100 phases
// Part 2: TBD
func Day16(input []byte, part1 bool) uint {
	input = bytes.TrimSpace(input)

	if part1 {
		return fftPart1(input)
	}
	return fftPart2(input)
}

func fftPart1(input []byte) uint {
	// Parse input to get digits
	digits := parseDigits(input)

	// Run 100 phases of FFT
	for range 100 {
		digits = applyFFTPhase(digits)
	}

	// Extract first 8 digits and convert to uint
	result := uint(0)
	for i := 0; i < 8 && i < len(digits); i++ {
		result = result*10 + uint(digits[i])
	}

	return result
}

func fftPart2(input []byte) uint {
	// TODO: Part 2 implementation
	return 0
}

// parseDigits converts byte string to slice of int digits
func parseDigits(input []byte) []int {
	digits := make([]int, len(input))
	for i, b := range input {
		digits[i] = int(b - '0')
	}
	return digits
}

// applyFFTPhase applies one phase of FFT to the input
func applyFFTPhase(input []int) []int {
	output := make([]int, len(input))

	for i := range output {
		output[i] = calculateOutputDigit(input, i)
	}

	return output
}

// calculateOutputDigit calculates a single output digit at position pos (0-indexed)
func calculateOutputDigit(input []int, pos int) int {
	basePattern := []int{0, 1, 0, -1}

	sum := 0
	patternIdx := 0
	repeatCount := 0
	skipFirst := true

	for inputIdx := range input {
		// Generate pattern value for this input position
		var patternValue int

		if skipFirst {
			skipFirst = false
			// Advance pattern as if we used one value
			repeatCount++
			if repeatCount > pos {
				repeatCount = 0
				patternIdx = (patternIdx + 1) % len(basePattern)
			}
		}

		patternValue = basePattern[patternIdx]
		sum += input[inputIdx] * patternValue

		// Advance pattern
		repeatCount++
		if repeatCount > pos {
			repeatCount = 0
			patternIdx = (patternIdx + 1) % len(basePattern)
		}
	}

	// Keep only the ones digit (absolute value)
	result := sum
	if result < 0 {
		result = -result
	}
	return result % 10
}
