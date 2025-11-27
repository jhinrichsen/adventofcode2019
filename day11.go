package adventofcode2019

import (
	"bytes"
	"fmt"
	"image"
)

const (
	// no types because we want to pass it to the IntCode Computer
	colorBlack = 0
	colorWhite = 1
)

// RegistrationID is a 2D map of black (false) and white (true)
type RegistrationID map[image.Point]bool

// returns min and max as 2D coordinates (x/y)
func (a RegistrationID) dim() (min, max image.Point) {
	minX, minY := int(^uint(0)>>1), int(^uint(0)>>1)       // max int
	maxX, maxY := -int(^uint(0)>>1)-1, -int(^uint(0)>>1)-1 // min int
	for k := range a {
		if k.X < minX {
			minX = k.X
		}
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y < minY {
			minY = k.Y
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}
	return image.Point{X: minX - 2, Y: minY - 2}, image.Point{X: maxX + 2, Y: maxY + 2}
}

// pbm creats an image in portable bitmap format from an registrationIdentifier.
// https://en.wikipedia.org/wiki/Netpbm#File_formats
func (a RegistrationID) pbm() []byte {
	var buf bytes.Buffer
	// magic number
	fmt.Fprintln(&buf, "P1")
	min, max := a.dim()
	// width height
	fmt.Fprintf(&buf, "%d %d\n", max.X-min.X, max.Y-min.Y)
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			// PBM uses 0 for white and 1 for black (inverse)
			var pbmCol int
			if b, ok := a[image.Point{X: x, Y: y}]; ok {
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
	position := image.Point{X: 0, Y: 0}
	// The robot starts facing up (negative Y direction)
	direction := image.Point{X: 0, Y: -1}

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
	// turnLeft rotates 90° counterclockwise: (x, y) → (-y, x)
	turnLeft := func(d image.Point) image.Point {
		return image.Point{X: -d.Y, Y: d.X}
	}
	// turnRight rotates 90° clockwise: (x, y) → (y, -x)
	turnRight := func(d image.Point) image.Point {
		return image.Point{X: d.Y, Y: -d.X}
	}
	translateTurn := func(i int) image.Point {
		const (
			// no types because we want to pass it to the IntCode Computer
			turn_Left  = 0
			turn_Right = 1
		)
		if i == turn_Left {
			return turnLeft(direction)
		} else if i == turn_Right {
			return turnRight(direction)
		}
		// Invalid turn value, keep current direction
		return direction
	}

	in <- initialColor
	for col := range out {
		setColor(col)
		// make robot turn
		direction = translateTurn(<-out)
		// one step
		position = position.Add(direction)

		in <- color()
	}
	return panels
}
