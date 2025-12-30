package adventofcode2019

// ndigits returns number of digits that n consists of. 0 has 1 digit, 42 has
// 2 digits, -4711 has 4 digits.
func ndigits(n int) int {
	if n < 0 {
		n = -n
	}
	divcount := 1
	for n >= 10 {
		n /= 10
		divcount++
	}
	return divcount
}

// digits splits n into its digits
func digits(n int) []byte {
	if n < 0 {
		n = -n
	}
	buf := make([]byte, ndigits(n))
	digitsInto(n, buf)
	return buf
}

// digitsInto writes digits of n into buf
func digitsInto(n int, buf []byte) {
	if n < 0 {
		n = -n
	}
	for i := len(buf) - 1; i >= 0 && n > 0; i-- {
		buf[i] = byte(n % 10)
		n /= 10
	}
}
