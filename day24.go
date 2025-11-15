package adventofcode2019

// Day24 simulates bug evolution on a 5x5 grid.
// For part 1, it returns the biodiversity rating of the first repeated layout.
// For part 2, it returns the number of bugs after 200 minutes in recursive grids.
func Day24(lines []string, part1 bool) uint {
	if part1 {
		return findFirstRepeatingBiodiversity(lines)
	}
	return simulateRecursive(lines, 200)
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

// simulateRecursive simulates bug evolution across recursive grid levels.
func simulateRecursive(lines []string, minutes int) uint {
	// Use a map to track grids at each depth level
	grids := make(map[int]uint)
	grids[0] = parseGridPart2(lines)

	for range minutes {
		grids = evolveRecursive(grids)
	}

	// Count total bugs across all levels
	total := uint(0)
	for _, grid := range grids {
		total += countBugs(grid)
	}
	return total
}

// parseGridPart2 parses the initial grid, treating center tile as always empty.
func parseGridPart2(lines []string) uint {
	grid := uint(0)

	for y := range 5 {
		if y >= len(lines) {
			break
		}
		line := []byte(lines[y])
		for x := range 5 {
			// Skip center tile (it's recursive, always empty)
			if x == 2 && y == 2 {
				continue
			}

			if x < len(line) && (line[x] == '#' || line[x] == '?') {
				bit := y*5 + x
				grid |= (1 << bit)
			}
		}
	}

	return grid
}

// evolveRecursive evolves all grid levels simultaneously.
func evolveRecursive(grids map[int]uint) map[int]uint {
	// Find min and max levels, and expand range to check adjacent levels
	minLevel, maxLevel := 0, 0
	for level := range grids {
		if level < minLevel {
			minLevel = level
		}
		if level > maxLevel {
			maxLevel = level
		}
	}

	// Expand the range to check outer and inner levels
	minLevel--
	maxLevel++

	newGrids := make(map[int]uint)

	for level := minLevel; level <= maxLevel; level++ {
		grid := grids[level] // Will be 0 if not present
		newGrid := uint(0)

		for y := range 5 {
			for x := range 5 {
				// Skip center tile
				if x == 2 && y == 2 {
					continue
				}

				bit := y*5 + x
				hasBug := (grid & (1 << bit)) != 0
				adjacentBugs := countAdjacentBugsRecursive(grids, x, y, level)

				if hasBug {
					if adjacentBugs == 1 {
						newGrid |= (1 << bit)
					}
				} else {
					if adjacentBugs == 1 || adjacentBugs == 2 {
						newGrid |= (1 << bit)
					}
				}
			}
		}

		if newGrid != 0 {
			newGrids[level] = newGrid
		}
	}

	return newGrids
}

// countAdjacentBugsRecursive counts bugs in adjacent cells across recursive levels.
func countAdjacentBugsRecursive(grids map[int]uint, x, y, level int) int {
	count := 0

	// Check all 4 directions
	directions := [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]

		// Check if we're going to the center tile
		if nx == 2 && ny == 2 {
			// Add tiles from inner level (level + 1)
			innerGrid := grids[level+1]
			if dir[0] == -1 { // moving left to center, so right column of inner
				for innerY := range 5 {
					bit := innerY*5 + 4
					if (innerGrid & (1 << bit)) != 0 {
						count++
					}
				}
			} else if dir[0] == 1 { // moving right to center, so left column of inner
				for innerY := range 5 {
					bit := innerY*5 + 0
					if (innerGrid & (1 << bit)) != 0 {
						count++
					}
				}
			} else if dir[1] == -1 { // moving up to center, so bottom row of inner
				for innerX := range 5 {
					bit := 4*5 + innerX
					if (innerGrid & (1 << bit)) != 0 {
						count++
					}
				}
			} else if dir[1] == 1 { // moving down to center, so top row of inner
				for innerX := range 5 {
					bit := 0*5 + innerX
					if (innerGrid & (1 << bit)) != 0 {
						count++
					}
				}
			}
		} else if nx < 0 || nx >= 5 || ny < 0 || ny >= 5 {
			// Out of bounds, so we need to look at outer level (level - 1)
			outerGrid := grids[level-1]
			var bit int
			if nx < 0 { // left edge, so tile to the left of center in outer
				bit = 2*5 + 1
			} else if nx >= 5 { // right edge, so tile to the right of center in outer
				bit = 2*5 + 3
			} else if ny < 0 { // top edge, so tile above center in outer
				bit = 1*5 + 2
			} else if ny >= 5 { // bottom edge, so tile below center in outer
				bit = 3*5 + 2
			}
			if (outerGrid & (1 << bit)) != 0 {
				count++
			}
		} else {
			// Normal case, same level
			bit := ny*5 + nx
			if (grids[level] & (1 << bit)) != 0 {
				count++
			}
		}
	}

	return count
}

// countBugs counts the number of bugs in a grid.
func countBugs(grid uint) uint {
	count := uint(0)
	for i := range 25 {
		if (grid & (1 << i)) != 0 {
			count++
		}
	}
	return count
}
