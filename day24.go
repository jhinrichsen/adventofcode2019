package adventofcode2019

import "math/bits"

// Day24 simulates bug evolution on a 5x5 grid.
// For part 1, it returns the biodiversity rating of the first repeated layout.
// For part 2, it returns the number of bugs after 200 minutes in recursive grids.
func Day24(lines []string, part1 bool) uint {
	if part1 {
		return findFirstRepeatingBiodiversity(lines)
	}
	return simulateRecursiveOptimized(lines, 200)
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

// countBugs counts the number of bugs in a grid.
func countBugs(grid uint) uint {
	return uint(bits.OnesCount(grid))
}

// Precomputed neighbor information for recursive grids.
// For each of 25 cells, stores masks for neighbors at same/outer/inner levels.
type day24NeighborInfo struct {
	sameLevelMask uint32 // Bitmask of same-level neighbors
	outerMask     uint32 // Bitmask of outer-level cells to check
	innerMask     uint32 // Bitmask of inner-level cells to check
}

var day24Neighbors [25]day24NeighborInfo

func init() {
	// Outer level tile indices for edge cells
	// Going left off grid → tile 11 (2*5+1), right → tile 13 (2*5+3)
	// Going up off grid → tile 7 (1*5+2), down → tile 17 (3*5+2)
	const outerLeft = 1 << (2*5 + 1)
	const outerRight = 1 << (2*5 + 3)
	const outerUp = 1 << (1*5 + 2)
	const outerDown = 1 << (3*5 + 2)

	// Inner level masks for cells adjacent to center
	// Right column (x=4): bits 4,9,14,19,24
	const innerRightCol = (1 << 4) | (1 << 9) | (1 << 14) | (1 << 19) | (1 << 24)
	// Left column (x=0): bits 0,5,10,15,20
	const innerLeftCol = (1 << 0) | (1 << 5) | (1 << 10) | (1 << 15) | (1 << 20)
	// Bottom row (y=4): bits 20,21,22,23,24
	const innerBottomRow = (1 << 20) | (1 << 21) | (1 << 22) | (1 << 23) | (1 << 24)
	// Top row (y=0): bits 0,1,2,3,4
	const innerTopRow = (1 << 0) | (1 << 1) | (1 << 2) | (1 << 3) | (1 << 4)

	for y := range 5 {
		for x := range 5 {
			bit := y*5 + x
			if x == 2 && y == 2 {
				continue // Skip center
			}

			var info day24NeighborInfo

			// Check 4 directions: up, down, left, right
			dirs := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
			for _, dir := range dirs {
				nx, ny := x+dir[0], y+dir[1]

				if nx == 2 && ny == 2 {
					// Going to center → need inner level
					if dir[0] == -1 { // moving left to center
						info.innerMask |= innerRightCol
					} else if dir[0] == 1 { // moving right to center
						info.innerMask |= innerLeftCol
					} else if dir[1] == -1 { // moving up to center
						info.innerMask |= innerBottomRow
					} else { // moving down to center
						info.innerMask |= innerTopRow
					}
				} else if nx < 0 {
					info.outerMask |= outerLeft
				} else if nx >= 5 {
					info.outerMask |= outerRight
				} else if ny < 0 {
					info.outerMask |= outerUp
				} else if ny >= 5 {
					info.outerMask |= outerDown
				} else {
					// Same level neighbor
					nbit := ny*5 + nx
					info.sameLevelMask |= (1 << nbit)
				}
			}

			day24Neighbors[bit] = info
		}
	}
}

// simulateRecursiveOptimized uses slices and precomputed neighbors.
func simulateRecursiveOptimized(lines []string, minutes int) uint {
	// Use slices with offset indexing instead of map
	// After N minutes, levels range from -N to +N
	const maxLevels = 201 // supports up to 200 minutes
	const offset = maxLevels

	// Double-buffer for grids: current and next
	grids := make([]uint32, 2*maxLevels+1)
	nextGrids := make([]uint32, 2*maxLevels+1)

	// Parse initial grid
	grids[offset] = uint32(parseGridPart2(lines))

	// Track active level range
	minLevel, maxLevel := 0, 0

	for range minutes {
		// Expand range to check adjacent levels
		checkMin := minLevel - 1
		checkMax := maxLevel + 1

		// Reset next grids in active range
		for level := checkMin; level <= checkMax; level++ {
			nextGrids[level+offset] = 0
		}

		newMinLevel, newMaxLevel := checkMax, checkMin // Will be updated

		for level := checkMin; level <= checkMax; level++ {
			idx := level + offset
			grid := grids[idx]
			outerGrid := grids[idx-1]
			innerGrid := grids[idx+1]

			var newGrid uint32

			for bit := range 25 {
				if bit == 12 { // Center (2,2)
					continue
				}

				info := &day24Neighbors[bit]
				adjacentBugs := bits.OnesCount32(grid&info.sameLevelMask) +
					bits.OnesCount32(outerGrid&info.outerMask) +
					bits.OnesCount32(innerGrid&info.innerMask)

				hasBug := (grid & (1 << bit)) != 0

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

			if newGrid != 0 {
				nextGrids[idx] = newGrid
				if level < newMinLevel {
					newMinLevel = level
				}
				if level > newMaxLevel {
					newMaxLevel = level
				}
			}
		}

		// Swap buffers
		grids, nextGrids = nextGrids, grids
		minLevel, maxLevel = newMinLevel, newMaxLevel
	}

	// Count total bugs across all levels
	total := uint(0)
	for level := minLevel; level <= maxLevel; level++ {
		total += uint(bits.OnesCount32(grids[level+offset]))
	}
	return total
}
