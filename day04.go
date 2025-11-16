package adventofcode2019

const (
	// Lower range from puzzle input
	Lower = 136818
	// Upper range from puzzle input
	Upper = 685979
)

// Criteria is a function that tests an integer. For performance reasons,
// both numerical and digits are supplied.
type Criteria func(int, []byte) bool

// CritSixDigits It is a six-digit number
func CritSixDigits(n int, digits []byte) bool {
	return len(digits) == 6
}

// CritWithinRange The value is within the range given in your puzzle input
func CritWithinRange(n int, digits []byte) bool {
	return Lower <= n && n <= Upper
}

// CritTwoOrMoreAdjacent Two adjacent digits are the same (like 22 in 122345)
func CritTwoOrMoreAdjacent(n int, digits []byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] == digits[i] {
			return true
		}
	}
	return false
}

// CritIncreasing Going from left to right, the digits never decrease; they only ever
// increase or stay the same (like 111123 or 135679)
func CritIncreasing(n int, digits []byte) bool {
	for i := 1; i < len(digits); i++ {
		if digits[i-1] > digits[i] {
			return false
		}
	}
	return true
}

// MeetsCriteria returns true if all criteria are fulfilled for n, using short
// circuit evaluation
func MeetsCriteria(n int, digits []byte, crits []Criteria) bool {
	for _, f := range crits {
		if !f(n, digits) {
			return false
		}
	}
	return true
}

// CriteriaPart1 returns all required criteria for part 1
func CriteriaPart1() []Criteria {
	return []Criteria{CritSixDigits, CritWithinRange, CritTwoOrMoreAdjacent, CritIncreasing}
}

// Day4Part1 returns number of passwords between Lower and Upper that meet all
// criteria
func Day4Part1() int {
	count := 0
	crits := CriteriaPart1()
	var digits [6]byte // All numbers in range are 6 digits
	// this range selection makes Crit1 superfluous
	for n := Lower; n < Upper; n++ {
		DigitsInto(n, digits[:])
		if MeetsCriteria(n, digits[:], crits) {
			count++
		}
	}
	return count
}

// CritExactlyTwoAdjacent two adjacent matching digits are not part of a larger
// group of matching digits.
func CritExactlyTwoAdjacent(n int, digits []byte) bool {
	// it's a bit fiddly to check for different surrounding digits,
	// _and_ to keep an eye on invalid indices (too far to the left, too far
	// to the right), so we'll handle leading and trailing groups separately

	// start of digits
	// +---+---+---+...
	// | 0 | 1 | 2 |...
	// +---+---+---+...
	// | a | a | b |...
	// +---+---+---+...
	if (digits[0] == digits[1]) && (digits[1] != digits[2]) {
		return true
	}

	// everything in between
	// ...+---+---+---+---+...
	// ...|i-2|i-1| i |i+1|...
	// ...+---+---+---+---+...
	// ...| a | b | b | c |...
	// ...+---+---+---+---+...
	l := len(digits) - 1
	for i := 2; i < l; i++ {
		// a != b
		if (digits[i-2] != digits[i-1]) &&
			// b == b
			(digits[i-1] == digits[i]) &&
			// b != c
			(digits[i] != digits[i+1]) {
			return true
		}
	}

	// end of digts
	// ...+---+---+---+
	// ...|l-2|l-1| l |
	// ...+---+---+---+
	// ...| b | a | a |
	// ...+---+---+---+
	if (digits[l-2] != digits[l-1]) && (digits[l-1] == digits[l]) {
		return true
	}
	return false
}

// Day4Part2 returns number of passwords between Lower and Upper that meet all
// criteria
func Day4Part2() int {
	count := 0
	crits := CriteriaPart2()
	var digits [6]byte // All numbers in range are 6 digits
	// this range selection makes Crit1 superfluous
	for n := Lower; n < Upper; n++ {
		DigitsInto(n, digits[:])
		if MeetsCriteria(n, digits[:], crits) {
			count++
		}
	}
	return count
}

// CriteriaPart2 returns all required criteria for part 1
func CriteriaPart2() []Criteria {
	return []Criteria{CritSixDigits, CritWithinRange, CritIncreasing, CritExactlyTwoAdjacent}
}
