package adventofcode2019

// Day19 solves the "Tractor Beam" puzzle.
// It tests how many points are affected by a tractor beam.
func Day19(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	if part1 {
		return countBeamPoints(ic, 50), nil
	}
	return findSquare(ic, 100), nil
}

// testPoint checks if a point (x, y) is affected by the tractor beam
func testPoint(ic *intcode, x, y int) bool {
	ic.Reset()
	inputIdx := 0
	inputs := [2]int{x, y}

	for {
		state := ic.Step()
		switch state {
		case needsInput:
			ic.Input(inputs[inputIdx])
			inputIdx++
		case hasOutput:
			return ic.Output() == 1
		case halted:
			return false
		}
	}
}

// countBeamPoints counts how many points in a size×size grid are affected
func countBeamPoints(ic *intcode, size int) uint {
	count := uint(0)
	for y := range size {
		for x := range size {
			if testPoint(ic, x, y) {
				count++
			}
		}
	}
	return count
}

func findSquare(ic *intcode, square int) uint {
	// y represents the BOTTOM row of the square
	y := square - 1

	// Track the leftmost x position to avoid searching from 0 each time
	leftX := 0

	for {
		y++

		// Find the leftmost beam point in this row
		x := leftX
		found := false

		for x <= y*2 {
			if testPoint(ic, x, y) {
				found = true
				leftX = x
				break
			}
			x++
		}

		if !found {
			continue
		}

		// Check if top-right corner of square is in beam
		topRightX := x + square - 1
		topRightY := y - square + 1

		if topRightY >= 0 && testPoint(ic, topRightX, topRightY) {
			// Square fits! Return top-left corner value
			return uint(x*10000 + topRightY)
		}
	}
}
