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
	digits := parseDigits(input)

	// Extract offset from first 7 digits
	offset := 0
	for i := 0; i < 7 && i < len(digits); i++ {
		offset = offset*10 + digits[i]
	}

	// Repeat input 10000 times
	repeatedLength := len(digits) * 10000

	// Key insight: if offset is in second half, FFT pattern simplifies
	// Each output digit is sum of all digits from that position to end (mod 10)
	// We only need to keep track of digits from offset to end

	if offset < repeatedLength/2 {
		// Offset is in first half - would need full FFT (very slow)
		// This shouldn't happen with valid puzzle inputs
		return 0
	}

	// Build the relevant portion (from offset to end)
	relevantLength := repeatedLength - offset
	signal := make([]int, relevantLength)

	for i := 0; i < relevantLength; i++ {
		actualPos := offset + i
		signal[i] = digits[actualPos%len(digits)]
	}

	// Apply 100 phases using the simplified algorithm
	// For second half: output[i] = (sum of input[i:]) % 10
	for range 100 {
		// Calculate from right to left
		sum := 0
		for i := len(signal) - 1; i >= 0; i-- {
			sum += signal[i]
			signal[i] = sum % 10
		}
	}

	// Extract first 8 digits of the result
	result := uint(0)
	for i := 0; i < 8 && i < len(signal); i++ {
		result = result*10 + uint(signal[i])
	}

	return result
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
