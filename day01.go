package adventofcode2019

import (
	"fmt"
)

// fuel computes required fuel for given mass. Both mass and fuel have no units.
func fuel(mass uint) int {
	return int(mass)/3 - 2
}

// completeFuel computes required fuel for given mass, including fuel for fuel.
func completeFuel(mass uint) uint {
	sum := uint(0)
	for f := fuel(mass); f > 0; f = fuel(uint(f)) {
		sum += uint(f)
	}
	return sum
}

// Day01 returns sum of fuel for all modules
func Day01(input []byte, part1 bool) (uint, error) {
	sum := uint(0)
	num := 0
	hasDigits := false

	for _, b := range input {
		if b >= '0' && b <= '9' {
			num = num*10 + int(b-'0')
			hasDigits = true
		} else if b == '\n' {
			if hasDigits {
				mass := uint(num)
				if part1 {
					f := fuel(mass)
					if f > 0 {
						sum += uint(f)
					}
				} else {
					sum += completeFuel(mass)
				}
				num = 0
				hasDigits = false
			}
		} else {
			return 0, fmt.Errorf("unexpected byte: %q", b)
		}
	}

	// Handle last number if no trailing newline
	if hasDigits {
		mass := uint(num)
		if part1 {
			f := fuel(mass)
			if f > 0 {
				sum += uint(f)
			}
		} else {
			sum += completeFuel(mass)
		}
	}

	return sum, nil
}

// Fuel computes required fuel for given mass (wrapper for tests)
func Fuel(mass int) int {
	return fuel(uint(mass))
}

// CompleteFuel computes required fuel for given mass, including fuel for fuel (wrapper for tests)
func CompleteFuel(mass int) int {
	return int(completeFuel(uint(mass)))
}
