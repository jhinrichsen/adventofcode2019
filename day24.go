package adventofcode2019

// Day24 simulates bug evolution on a 5x5 grid.
// For part 1, it returns the biodiversity rating of the first repeated layout.
func Day24(lines []string, part1 bool) uint {
	if part1 {
		return findFirstRepeatingBiodiversity(lines)
	}
	return 0 // Part 2 not yet implemented
}

// findFirstRepeatingBiodiversity simulates bug evolution until a layout repeats.
func findFirstRepeatingBiodiversity(lines []string) uint {
	const gridSize = 5

	// Parse initial grid into bitmask
	grid := parseGrid(lines)

	// Track seen states
	seen := make(map[uint]bool)
	seen[grid] = true

	// Simulate until we find a repeat
	for {
		grid = evolve(grid, gridSize)
		if seen[grid] {
			return grid
		}
		seen[grid] = true
	}
}

// parseGrid converts a slice of strings into a bitmask representation.
func parseGrid(lines []string) uint {
	grid := uint(0)
	bit := uint(0)

	for y := range 5 {
		if y >= len(lines) {
			break
		}
		line := []byte(lines[y])
		for x := range 5 {
			if x < len(line) && line[x] == '#' {
				grid |= (1 << bit)
			}
			bit++
		}
	}

	return grid
}

// evolve simulates one minute of bug evolution.
func evolve(grid uint, size int) uint {
	newGrid := uint(0)

	for y := range size {
		for x := range size {
			bit := y*size + x
			hasBug := (grid & (1 << bit)) != 0
			adjacentBugs := countAdjacentBugs(grid, x, y, size)

			if hasBug {
				// Bug survives only if exactly 1 adjacent bug
				if adjacentBugs == 1 {
					newGrid |= (1 << bit)
				}
			} else {
				// Empty space becomes infested if 1 or 2 adjacent bugs
				if adjacentBugs == 1 || adjacentBugs == 2 {
					newGrid |= (1 << bit)
				}
			}
		}
	}

	return newGrid
}

// countAdjacentBugs counts bugs in the 4 adjacent cells (not diagonals).
func countAdjacentBugs(grid uint, x, y, size int) int {
	count := 0

	// Check all 4 directions
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < size && ny >= 0 && ny < size {
			bit := ny*size + nx
			if (grid & (1 << bit)) != 0 {
				count++
			}
		}
	}

	return count
}
