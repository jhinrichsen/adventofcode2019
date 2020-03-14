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

// RegistrationIdentifier is a 2D map of black (false) and white (true)
type RegistrationIdentifier map[complex128]bool

func (a RegistrationIdentifier) dim() (complex128, complex128) {
	return -1 + -1i, 1 + 1i
}

// pbm creats an image in portable bitmap format from an registrationIdentifier.
// https://en.wikipedia.org/wiki/Netpbm#File_formats
func (a RegistrationIdentifier) pbm() []byte {
	var buf bytes.Buffer
	// magic number
	fmt.Fprintln(&buf, "P1")
	min, max := a.dim()
	box := func(f float64) int {
		return int(math.Round(math.Ceil(f)))
	}
	x0 := box(real(min))
	xn := box(real(max))
	y0 := box(imag(min))
	yn := box(imag(max))
	// width height
	fmt.Fprintf(&buf, "%d %d\n", xn-x0, yn-y0)
	for y := y0; y < yn; y++ {
		for x := x0; x < xn; x++ {
			// PBM uses 0 for white and 1 for black (inverse)
			var pbmCol int
			if _, white := a[complex(float64(x), float64(y))]; white {
				pbmCol = colorBlack
			} else {
				pbmCol = colorWhite
			}
			fmt.Fprintf(&buf, "%d", pbmCol)
		}
		buf.WriteString("\n")
	}
	return buf.Bytes()
}

// Day11Part1 returns the number of painted tiles.
func Day11Part1(prog IntCode) int {
	return len(registrationIdentifier(prog))
}

func registrationIdentifier(prog IntCode) RegistrationIdentifier {
	const (
		// no types because we want to pass it to the IntCode Computer
		turnLeft  = 0
		turnRight = 1
	)
	in, out := channels()

	// boot emergency hull painting robot
	go Day5(prog, in, out)

	panels := make(RegistrationIdentifier)
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
		if i == turnLeft {
			return left
		} else if i == turnRight {
			return right
		} else {
			panic(fmt.Errorf("unsupported turn value %d", i))
		}
	}

	in <- color()
	for col := range out {
		setColor(col)
		turnValue := <-out
		// make robot turn
		turn := translateTurn(turnValue)
		direction *= turn
		// one step
		position += direction

		in <- color()
	}
	return panels
}
