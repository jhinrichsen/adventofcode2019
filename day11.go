package adventofcode2019

import (
	"bytes"
	"fmt"
	"math"
)

const (
	// no types because we want to pass it to the IntCode Computer
	colorBlack = 0
	colorWhite = 1
)

// RegistrationID is a 2D map of black (false) and white (true)
type RegistrationID map[complex128]bool

// box returns the integer away from 0 that can hold f.
func box(f float64) int {
	return int(math.Round(math.Ceil(f)))
}

// returns min and max as 2D coordinates (x/y)
func (a RegistrationID) dim() (min, max Point) {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := -math.MaxInt64, -math.MaxInt64
	for k := range a {
		x := box(real(k))
		if x < minX {
			minX = x
		} else if x > maxX {
			maxX = x
		}
		y := box(imag(k))
		if y < minY {
			minY = y
		} else if y > maxY {
			maxY = y
		}
	}
	return Point{minX - 2, minY - 2}, Point{maxX + 2, maxY + 2}
}

// pbm creats an image in portable bitmap format from an registrationIdentifier.
// https://en.wikipedia.org/wiki/Netpbm#File_formats
func (a RegistrationID) pbm() []byte {
	var buf bytes.Buffer
	// magic number
	fmt.Fprintln(&buf, "P1")
	min, max := a.dim()
	// width height
	fmt.Fprintf(&buf, "%d %d\n", max.x-min.x, max.y-min.y)
	for y := min.y; y < max.y; y++ {
		for x := min.x; x < max.x; x++ {
			// PBM uses 0 for white and 1 for black (inverse)
			var pbmCol int
			if b, ok := a[complex(float64(x), float64(y))]; ok {
				if b {

					pbmCol = colorWhite
				} else {
					pbmCol = colorBlack
				}
			} else {
				// no color, default
				pbmCol = colorBlack
			}
			fmt.Fprintf(&buf, "%d", pbmCol)
		}
		buf.WriteString("\n")
	}
	return buf.Bytes()
}

// Day11Part1 returns the number of painted tiles.
func Day11Part1(prog IntCode) int {
	return len(newRegistrationID(prog, colorBlack))
}

// Day11Part2 returns a PBM encoded registration ID.
// The image is horizontally flipped so em well you can read it from the inside
// of the transparent hull?
func Day11Part2(prog IntCode) []byte {
	return newRegistrationID(prog, colorWhite).pbm()
}

func newRegistrationID(prog IntCode, initialColor int) RegistrationID {
	in, out := channels()

	// boot emergency hull painting robot
	go Day5(prog, in, out)

	panels := make(RegistrationID)
	position := 0 + 0i
	left := 0 + 1i
	right := 0 - 1i
	// The robot starts facing up
	direction := left

	color := func() int {
		// Default is black
		if white := panels[position]; white {
			return colorWhite
		}
		return colorBlack
	}
	setColor := func(c int) {
		b := c == colorWhite
		panels[position] = b
	}
	translateTurn := func(i int) complex128 {
		const (
			// no types because we want to pass it to the IntCode Computer
			turnLeft  = 0
			turnRight = 1
		)
		if i == turnLeft {
			return left
		} else if i == turnRight {
			return right
		} else {
			panic(fmt.Errorf("unsupported turn value %d", i))
		}
	}

	in <- initialColor
	for col := range out {
		setColor(col)
		// make robot turn
		turn := translateTurn(<-out)
		direction *= turn
		// one step
		position += direction

		in <- color()
	}
	return panels
}
