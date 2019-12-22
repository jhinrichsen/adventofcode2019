package adventofcode2019

import (
	"strconv"
)

// CompleteFuel computes required fuel for given mass, including fuel for fuel.
// Both mass and fuel have no units.
func CompleteFuel(mass int) int {
	sum := 0
	f := mass
again:
	f = Fuel(f)
	if f > 0 {
		sum += f
		goto again
	}
	return sum
}

// Fuel computes required fuel for given mass. Both mass and fuel have no units.
func Fuel(mass int) int {
	return mass/3 - 2
}

// Day1Part1 returns sum of fuel for all modules
func Day1Part1(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return sum, err
		}
		sum += Fuel(n)
	}
	return sum, nil
}

// Day1Part2 returns sum of fuel and transitively the required fuel for the
// fuel for all modules
func Day1Part2(lines []string) (int, error) {
	sum := 0
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return sum, err
		}
		sum += CompleteFuel(n)
	}
	return sum, nil
}
