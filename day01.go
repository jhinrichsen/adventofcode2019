package adventofcode2019

import (
	"fmt"
)

// Day01 returns sum of fuel for all modules
func Day01(input []byte, part1 bool) (uint, error) {
	// fuel computes required fuel for given mass
	fuel := func(mass uint) int {
		return int(mass)/3 - 2
	}

	// completeFuel computes required fuel for given mass, including fuel for fuel.
	// Loop terminates when f <= 0 (fuel can be negative for small masses).
	completeFuel := func(mass uint) (sum uint) {
		for f := fuel(mass); f > 0; f = fuel(uint(f)) {
			sum += uint(f)
		}
		return
	}

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

	// .editorconfig guarantees insert_final_newline, so error if digits remain
	if hasDigits {
		return 0, fmt.Errorf("input does not end with newline")
	}

	return sum, nil
}
