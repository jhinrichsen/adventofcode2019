package adventofcode2019

import (
	"strconv"
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
func Day01(lines []string, part1 bool) uint {
	sum := uint(0)
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			continue // Skip invalid lines (AoC input is always valid)
		}
		mass := uint(n)
		if part1 {
			f := fuel(mass)
			if f > 0 {
				sum += uint(f)
			}
		} else {
			sum += completeFuel(mass)
		}
	}
	return sum
}

// Fuel computes required fuel for given mass (wrapper for tests)
func Fuel(mass int) int {
	return fuel(uint(mass))
}

// CompleteFuel computes required fuel for given mass, including fuel for fuel (wrapper for tests)
func CompleteFuel(mass int) int {
	return int(completeFuel(uint(mass)))
}
