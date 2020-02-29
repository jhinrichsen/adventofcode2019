package adventofcode2019

import "fmt"

// Day8Part1 returns number of 1s multiplied by number of 2s in the layer having
// minimal number of 0s.
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

// Day8Part2 returns rendered layer.
func Day8Part2(digits []byte) ([]byte, error) {
	width := 25
	height := 6
	size := width * height
	rendered := make([]byte, size)

	if len(digits)%size != 0 {
		return rendered, fmt.Errorf("cannot pack %d digits into %dx%d layer",
			len(digits), width, height)
	}
	renderedPixel := func(i int) byte {
		// drill through layers as long as pixel is transparent
		// assume one pixel MUST be colored, otherwise we will run out
		// of layers and panic badly
		for digits[i] == '2' {
			i += size
		}
		return digits[i]
	}
	for pixel := range rendered {
		rendered[pixel] = renderedPixel(pixel)
	}
	return rendered, nil
}
