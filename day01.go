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
		n, _ := strconv.Atoi(line)
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

// Legacy functions for backwards compatibility
func Fuel(mass int) int {
	return fuel(uint(mass))
}

func CompleteFuel(mass int) int {
	return int(completeFuel(uint(mass)))
}

func Day1Part1(lines []string) (int, error) {
	return int(Day01(lines, true)), nil
}

func Day1Part2(lines []string) (int, error) {
	return int(Day01(lines, false)), nil
}
