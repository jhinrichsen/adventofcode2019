package adventofcode2019

// Abs returns absolute value for integer.
func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// Ndigits returns number of digits that n consists of. 0 has 1 digit, 42 has
// 2 digits, -4711 has 4 digits.
// Naming convention taken from Julia
func Ndigits(n int) int {
	n = Abs(n)
	divcount := 1
	for n >= 10 {
		n /= 10
		divcount++
	}
	return divcount
}

// Digits splits n into its digits
func Digits(n int) []byte {
	n = Abs(n)
	buf := make([]byte, Ndigits(n))
	DigitsInto(n, buf)
	return buf
}

// DigitsInto will write digits of n into buf
func DigitsInto(n int, buf []byte) {
	n = Abs(n)
	idx := len(buf) - 1
	for n > 0 {
		d := n % 10
		buf[idx] = byte(d)
		idx--
		n /= 10
	}
}
