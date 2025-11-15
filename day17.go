package adventofcode2019

import (
	"bytes"
)

// Day17 analyzes scaffolding map from ASCII camera
// Part 1: Sum of alignment parameters at intersections
// Part 2: Collect dust by visiting all scaffold
func Day17(program []byte, part1 bool) uint {
	code := MustSplit(string(bytes.TrimSpace(program)))

	if part1 {
		return calculateAlignmentSum(code)
	}
	return collectDust(code)
}

func calculateAlignmentSum(code IntCode) uint {
	// Run the Intcode program to get ASCII output
	input := make(chan int)
	output := make(chan int, 10000)

	prog := code.Copy()
	go Day5(prog, input, output)
	close(input)

	// Read ASCII output and build map
	var grid [][]byte
	var row []byte

	for val := range output {
		ch := byte(val)
		if ch == '\n' {
			if len(row) > 0 {
				grid = append(grid, row)
				row = nil
			}
		} else {
			row = append(row, ch)
		}
	}

	// Find intersections and calculate alignment parameters
	sum := uint(0)

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if isIntersection(grid, x, y) {
				// Alignment parameter is x * y
				sum += uint(x * y)
			}
		}
	}

	return sum
}

// isIntersection checks if position (x, y) is a scaffold intersection
// An intersection is a scaffold (#) with scaffolds on all 4 sides
func isIntersection(grid [][]byte, x, y int) bool {
	// Check if current position is a scaffold
	if grid[y][x] != '#' {
		return false
	}

	// Check all 4 directions
	if grid[y-1][x] != '#' { // North
		return false
	}
	if grid[y+1][x] != '#' { // South
		return false
	}
	if x > 0 && grid[y][x-1] != '#' { // West
		return false
	}
	if x < len(grid[y])-1 && grid[y][x+1] != '#' { // East
		return false
	}

	return true
}

func collectDust(code IntCode) uint {
	// Wake up the robot by changing address 0 from 1 to 2
	prog := code.Copy()
	prog[0] = 2

	input := make(chan int, 1000)
	output := make(chan int, 10000)

	go Day5(prog, input, output)

	// Movement routine - compressed from path analysis
	// Full path: R,8,L,12,R,8,R,8,L,12,R,8,L,10,L,10,R,8,L,12,L,12,L,10,R,10,L,10,L,10,R,8,L,12,L,12,L,10,R,10,L,10,L,10,R,8,R,8,L,12,R,8,L,12,L,12,L,10,R,10,R,8,L,12,R,8
	// Main routine: A,A,B,C,B,C,B,A,C,A
	// Function A: R,8,L,12,R,8
	// Function B: L,10,L,10,R,8
	// Function C: L,12,L,12,L,10,R,10
	// Video feed: n

	sendASCII := func(s string) {
		for _, ch := range s {
			input <- int(ch)
		}
		input <- 10 // newline
	}

	sendASCII("A,A,B,C,B,C,B,A,C,A")
	sendASCII("R,8,L,12,R,8")
	sendASCII("L,10,L,10,R,8")
	sendASCII("L,12,L,12,L,10,R,10")
	sendASCII("n")

	close(input)

	var lastOutput uint
	for val := range output {
		if val > 255 {
			// This is the dust collection amount (non-ASCII)
			lastOutput = uint(val)
		}
	}

	return lastOutput
}
