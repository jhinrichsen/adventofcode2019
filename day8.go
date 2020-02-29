package adventofcode2019

import "fmt"

func Day8Part1(digits []byte) (int, error) {
	width := 25
	height := 6
	size := width * height

	if len(digits)%size != 0 {
		return -1, fmt.Errorf("cannot pack %d digits into %dx%d layer",
			len(digits), width, height)
	}
	layers := len(digits) / size
	// count zeroes in layer N
	zeroes := make([]int, layers)
	for i := 0; i < len(digits); i++ {
		if digits[i] == '0' {
			zeroes[i/size]++
		}
	}
	// minimal layer
	min, minLayer := size, layers
	for i := range zeroes {
		if zeroes[i] < min {
			min = zeroes[i]
			minLayer = i
		}
	}
	ones := 0
	twos := 0
	for i := minLayer * size; i < (minLayer+1)*size; i++ {
		if digits[i] == '1' {
			ones++
		} else if digits[i] == '2' {
			twos++
		}
	}
	return ones * twos, nil
}
