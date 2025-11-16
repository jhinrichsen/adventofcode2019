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
	return runBeamTest(program.Copy(), x, y) == 1
}

// runBeamTest runs the IntCode program directly without channels for better performance
func runBeamTest(program IntCode, x, y int) int {
	ip := 0
	relBase := 0
	inputIdx := 0
	inputs := [2]int{x, y}

	realloc := func(idx int) {
		if idx < 0 {
			panic("negative address")
		}
		if idx >= len(program) {
			bigger := make(IntCode, idx+1)
			copy(bigger, program)
			program = bigger
		}
	}

	load := func(idx int, mode ParameterMode) int {
		lda := func(idx int) int {
			realloc(idx)
			return program[idx]
		}
		switch mode {
		case ImmediateMode:
			return lda(idx)
		case PositionMode:
			return lda(lda(idx))
		case RelativeMode:
			return lda(relBase + lda(idx))
		}
		return -1
	}

	store := func(idx int, val int, mode ParameterMode) {
		switch mode {
		case ImmediateMode:
			realloc(idx)
			program[idx] = val
		case PositionMode:
			adr := program[idx]
			realloc(adr)
			program[adr] = val
		case RelativeMode:
			adr := relBase + program[idx]
			realloc(adr)
			program[adr] = val
		}
	}

	for {
		opcode, mode1, mode2, mode3 := instruction(program[ip])
		switch opcode {
		case OpcodeAdd:
			val := load(ip+1, mode1) + load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case OpcodeMul:
			val := load(ip+1, mode1) * load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case Input:
			val := inputs[inputIdx]
			inputIdx++
			store(ip+1, val, mode1)
			ip += 2
		case Output:
			val := load(ip+1, mode1)
			return val
		case JumpIfTrue:
			p := load(ip+1, mode1)
			if True(p) {
				ip = load(ip+2, mode2)
				continue
			}
			ip += 3
		case JumpIfFalse:
			p := load(ip+1, mode1)
			if False(p) {
				ip = load(ip+2, mode2)
				continue
			}
			ip += 3
		case LessThan:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			val := Boolean(p1 < p2)
			store(ip+3, val, mode3)
			ip += 4
		case Equals:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			val := Boolean(p1 == p2)
			store(ip+3, val, mode3)
			ip += 4
		case AdjustRelBase:
			p1 := load(ip+1, mode1)
			relBase += p1
			ip += 2
		case OpcodeRet:
			return 0
		}
	}
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

	// Track the leftmost x position to avoid searching from 0 each time
	leftX := 0

	for {
		y++

		// Find the leftmost beam point in this row (bottom-left of square)
		// Start from the previous row's leftX since the beam only expands
		x := leftX
		found := false

		// Find first beam point in this row
		for x <= y*2 {
			if testPoint(program, x, y) {
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
