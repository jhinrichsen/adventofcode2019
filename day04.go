package adventofcode2019

// criteria is a function that tests an integer with its digits.
type criteria func(int, []byte) bool

func critSixDigits(n int, digits []byte) bool {
	return len(digits) == 6
}

func critTwoOrMoreAdjacent(n int, digits []byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] == digits[i] {
			return true
		}
	}
	return false
}

func critIncreasing(n int, digits []byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] > digits[i] {
			return false
		}
	}
	return true
}

func critExactlyTwoAdjacent(n int, digits []byte) bool {
	// Start of digits: aa != a
	if (digits[0] == digits[1]) && (digits[1] != digits[2]) {
		return true
	}

	// Middle: != aa !=
	for i := 2; i < 5; i++ {
		if (digits[i-2] != digits[i-1]) &&
			(digits[i-1] == digits[i]) &&
			(digits[i] != digits[i+1]) {
			return true
		}
	}

	// End of digits: != aa
	if (digits[3] != digits[4]) && (digits[4] == digits[5]) {
		return true
	}
	return false
}

func meetsCriteria(n int, digits []byte, crits []criteria) bool {
	for _, f := range crits {
		if !f(n, digits) {
			return false
		}
	}
	return true
}

// Day04 returns number of passwords that meet all criteria
func Day04(part1 bool) uint {
	const lower, upper = 136818, 685979
	count := uint(0)
	digits := make([]byte, 6)

	for n := lower; n < upper; n++ {
		digitsInto(n, digits)

		// Check increasing first (fastest rejection)
		increasing := true
		for i := 1; i < 6; i++ {
			if digits[i-1] > digits[i] {
				increasing = false
				break
			}
		}
		if !increasing {
			continue
		}

		if part1 {
			// Check for any two adjacent digits
			hasAdjacent := false
			for i := 1; i < 6; i++ {
				if digits[i-1] == digits[i] {
					hasAdjacent = true
					break
				}
			}
			if hasAdjacent {
				count++
			}
		} else {
			// Check for exactly two adjacent digits
			if critExactlyTwoAdjacent(0, digits) {
				count++
			}
		}
	}
	return count
}
