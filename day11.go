package adventofcode2019

import (
	"bytes"
	"fmt"
	"image"
)

const (
	colorBlack = 0
	colorWhite = 1
)

// registrationID is a 2D map of black (false) and white (true)
type registrationID map[image.Point]bool

// dim returns min and max as 2D coordinates (x/y)
func (a registrationID) dim() (min, max image.Point) {
	minX, minY := int(^uint(0)>>1), int(^uint(0)>>1)
	maxX, maxY := -int(^uint(0)>>1)-1, -int(^uint(0)>>1)-1
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

// pbm creates an image in portable bitmap format
func (a registrationID) pbm() []byte {
	var buf bytes.Buffer
	fmt.Fprintln(&buf, "P1")
	min, max := a.dim()
	fmt.Fprintf(&buf, "%d %d\n", max.X-min.X, max.Y-min.Y)
	for y := min.Y; y < max.Y; y++ {
		for x := min.X; x < max.X; x++ {
			var pbmCol int
			if b, ok := a[image.Point{X: x, Y: y}]; ok && b {
				pbmCol = colorWhite
			}
			fmt.Fprintf(&buf, "%d", pbmCol)
		}
		buf.WriteString("\n")
	}
	return buf.Bytes()
}

// Day11 runs the hull painting robot
func Day11(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	var initialColor int
	if !part1 {
		initialColor = colorWhite
	}

	panels := runRobot(ic, initialColor)

	if part1 {
		return uint(len(panels)), nil
	}
	// Part 2 returns a checksum of the image for testing
	return uint(len(panels.pbm())), nil
}

func runRobot(ic *intcode, initialColor int) registrationID {
	panels := make(registrationID)
	position := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 0, Y: -1} // facing up

	currentColor := initialColor
	outputCount := 0
	var paintColor int

	for {
		state := ic.Step()
		switch state {
		case needsInput:
			ic.Input(currentColor)
		case hasOutput:
			if outputCount%2 == 0 {
				// First output: color to paint
				paintColor = ic.Output()
			} else {
				// Second output: turn direction
				turn := ic.Output()
				panels[position] = paintColor == colorWhite

				// Turn and move
				if turn == 0 { // left
					direction = image.Point{X: -direction.Y, Y: direction.X}
				} else { // right
					direction = image.Point{X: direction.Y, Y: -direction.X}
				}
				position = position.Add(direction)

				// Get color of new position
				if panels[position] {
					currentColor = colorWhite
				} else {
					currentColor = colorBlack
				}
			}
			outputCount++
		case halted:
			return panels
		}
	}
}
