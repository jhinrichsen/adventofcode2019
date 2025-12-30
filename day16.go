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
	n := len(input)
	current := make([]int, n)
	next := make([]int, n)
	prefix := make([]int, n+1)

	for i, b := range input {
		current[i] = int(b - '0')
	}

	// Run 100 phases of FFT using prefix sums
	for range 100 {
		// Build prefix sums: prefix[i] = sum of current[0:i]
		prefix[0] = 0
		for i := range n {
			prefix[i+1] = prefix[i] + current[i]
		}

		// For each output position, process blocks of the pattern
		for i := range n {
			blockSize := i + 1
			sum := 0
			pos := i // Start at position i (skip first element of pattern)

			// Pattern cycles: +1 block, skip, -1 block, skip, +1 block, ...
			sign := 1
			for pos < n {
				// Add/subtract this block using prefix sums
				end := pos + blockSize
				if end > n {
					end = n
				}
				sum += sign * (prefix[end] - prefix[pos])
				pos += blockSize

				// Skip zero block
				pos += blockSize

				// Flip sign for next non-zero block
				sign = -sign
			}

			// Keep only ones digit (absolute value)
			if sum < 0 {
				sum = -sum
			}
			next[i] = sum % 10
		}

		// Swap current and next
		current, next = next, current
	}

	// Extract first 8 digits and convert to uint
	result := uint(0)
	for i := 0; i < 8 && i < n; i++ {
		result = result*10 + uint(current[i])
	}

	return result
}

func fftPart2(input []byte) uint {
	// Parse digits inline
	n := len(input)

	// Extract offset from first 7 digits
	offset := 0
	for i := 0; i < 7 && i < n; i++ {
		offset = offset*10 + int(input[i]-'0')
	}

	// Repeat input 10000 times
	repeatedLength := n * 10000

	// Key insight: if offset is in second half, FFT pattern simplifies
	// Each output digit is sum of all digits from that position to end (mod 10)
	// We only need to keep track of digits from offset to end

	if offset < repeatedLength/2 {
		// Offset is in first half - would need full FFT (very slow)
		// This shouldn't happen with valid puzzle inputs
		return 0
	}

	// Build the relevant portion (from offset to end) using bytes
	relevantLength := repeatedLength - offset
	signal := make([]byte, relevantLength)

	for i := range relevantLength {
		signal[i] = input[(offset+i)%n] - '0'
	}

	// Apply 100 phases using the simplified algorithm
	// For second half: output[i] = (sum of input[i:]) % 10
	for range 100 {
		// Calculate from right to left
		sum := 0
		for i := len(signal) - 1; i >= 0; i-- {
			sum += int(signal[i])
			signal[i] = byte(sum % 10)
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
