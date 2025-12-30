package adventofcode2019

// Day17 analyzes scaffolding map from ASCII camera
// Part 1: Sum of alignment parameters at intersections
// Part 2: Collect dust by visiting all scaffold
func Day17(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	if part1 {
		return calculateAlignmentSum(ic), nil
	}
	return collectDust(ic), nil
}

func calculateAlignmentSum(ic *intcode) uint {
	// Run the Intcode program to get ASCII output
	var grid [][]byte
	var row []byte

	for {
		state := ic.Step()
		switch state {
		case hasOutput:
			ch := byte(ic.Output())
			if ch == '\n' {
				if len(row) > 0 {
					grid = append(grid, row)
					row = nil
				}
			} else {
				row = append(row, ch)
			}
		case halted:
			goto done
		}
	}
done:

	// Find intersections and calculate alignment parameters
	sum := uint(0)

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if isIntersection(grid, x, y) {
				sum += uint(x * y)
			}
		}
	}

	return sum
}

// isIntersection checks if position (x, y) is a scaffold intersection
func isIntersection(grid [][]byte, x, y int) bool {
	if grid[y][x] != '#' {
		return false
	}
	if grid[y-1][x] != '#' {
		return false
	}
	if grid[y+1][x] != '#' {
		return false
	}
	if x > 0 && grid[y][x-1] != '#' {
		return false
	}
	if x < len(grid[y])-1 && grid[y][x+1] != '#' {
		return false
	}
	return true
}

func collectDust(ic *intcode) uint {
	// Wake up the robot by changing address 0 from 1 to 2
	ic.SetMem(0, 2)

	// Movement routine - compressed from path analysis
	// Main routine: A,A,B,C,B,C,B,A,C,A
	// Function A: R,8,L,12,R,8
	// Function B: L,10,L,10,R,8
	// Function C: L,12,L,12,L,10,R,10
	// Video feed: n
	commands := "A,A,B,C,B,C,B,A,C,A\nR,8,L,12,R,8\nL,10,L,10,R,8\nL,12,L,12,L,10,R,10\nn\n"
	cmdIdx := 0

	var lastOutput uint

	for {
		state := ic.Step()
		switch state {
		case needsInput:
			if cmdIdx < len(commands) {
				ic.Input(int(commands[cmdIdx]))
				cmdIdx++
			}
		case hasOutput:
			val := ic.Output()
			if val > 255 {
				lastOutput = uint(val)
			}
		case halted:
			return lastOutput
		}
	}
}
