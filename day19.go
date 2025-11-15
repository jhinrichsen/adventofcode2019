package adventofcode2019

import (
	"strconv"
	"strings"
)

// Day19 solves the "Tractor Beam" puzzle.
// It tests how many points are affected by a tractor beam.
func Day19(input []byte, part1 bool) uint {
	program := parseIntCode(string(input))
	if part1 {
		return countBeamPoints(program, 50)
	}
	return findSquare(program, 100)
}

func parseIntCode(s string) IntCode {
	parts := strings.Split(strings.TrimSpace(s), ",")
	code := make(IntCode, len(parts))
	for i, part := range parts {
		code[i], _ = strconv.Atoi(part)
	}
	return code
}

// testPoint checks if a point (x, y) is affected by the tractor beam
func testPoint(program IntCode, x, y int) bool {
	in := make(chan int, 2)
	out := make(chan int, 1)

	in <- x
	in <- y
	close(in)

	go Day5(program.Copy(), in, out)

	result := <-out
	return result == 1
}

// countBeamPoints counts how many points in a size×size grid are affected
func countBeamPoints(program IntCode, size int) uint {
	count := uint(0)
	for y := range size {
		for x := range size {
			if testPoint(program, x, y) {
				count++
			}
		}
	}
	return count
}

// testPointFromMap checks a point against a precomputed beam map (for testing)
func testPointFromMap(beamMap [][]bool, x, y int) bool {
	if y < 0 || y >= len(beamMap) || x < 0 || x >= len(beamMap[y]) {
		return false
	}
	return beamMap[y][x]
}

// findSquareFromMap finds square in a precomputed map (for testing)
// Uses the same algorithm as findSquare to ensure consistency
func findSquareFromMap(beamMap [][]bool, square int) uint {
	// Start searching from a reasonable y position
	// y represents the BOTTOM row of the square
	y := square

	for y < len(beamMap) {
		// Find the leftmost beam point in this row (bottom-left of square)
		x := 0
		found := false
		for x < len(beamMap[y]) {
			if testPointFromMap(beamMap, x, y) {
				found = true
				break
			}
			x++
		}

		if !found {
			y++
			continue
		}

		// Check if top-right corner of square is in beam
		// Bottom-left is at (x, y)
		// Top-right should be at (x + square - 1, y - square + 1)
		topRightX := x + square - 1
		topRightY := y - square + 1

		if topRightY >= 0 && testPointFromMap(beamMap, topRightX, topRightY) {
			// Square fits! Return top-left corner value
			// Top-left is at (x, topRightY)
			return uint(x*10000 + topRightY)
		}

		y++
	}
	return 0
}

func findSquare(program IntCode, square int) uint {
	// Start searching from a reasonable y position
	// y represents the BOTTOM row of the square
	y := square - 1

	for {
		y++
		// Find the leftmost beam point in this row (bottom-left of square)
		x := 0

		// Find first beam point in this row
		found := false
		for {
			if testPoint(program, x, y) {
				found = true
				break
			}
			x++
			if x > y*3 {
				// Safety limit
				break
			}
		}

		if !found {
			continue
		}

		// Check if top-right corner of square is in beam
		// Bottom-left is at (x, y)
		// Top-right should be at (x + square - 1, y - square + 1)
		topRightX := x + square - 1
		topRightY := y - square + 1

		if topRightY >= 0 && testPoint(program, topRightX, topRightY) {
			// Square fits! Return top-left corner value
			// Top-left is at (x, topRightY)
			return uint(x*10000 + topRightY)
		}
	}
}
