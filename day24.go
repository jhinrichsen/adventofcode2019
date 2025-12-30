package adventofcode2019

import "math/bits"

// Day24 simulates bug evolution on a 5x5 grid.
// Part 1: biodiversity rating of first repeated layout.
// Part 2: bug count after 200 minutes in recursive grids.
func Day24(lines []string, part1 bool) uint {
	grid := parseGrid24(lines)
	if part1 {
		return day24Part1(grid)
	}
	return day24Part2(grid, 200)
}

func parseGrid24(lines []string) uint32 {
	var grid uint32
	for y := range 5 {
		if y >= len(lines) {
			break
		}
		for x := range 5 {
			if x < len(lines[y]) && lines[y][x] == '#' {
				grid |= 1 << (y*5 + x)
			}
		}
	}
	return grid
}

func day24Part1(grid uint32) uint {
	seen := make(map[uint32]bool)
	seen[grid] = true

	for {
		grid = evolve(grid)
		if seen[grid] {
			return uint(grid)
		}
		seen[grid] = true
	}
}

func evolve(grid uint32) uint32 {
	var next uint32
	for bit := range 25 {
		x, y := bit%5, bit/5
		adj := 0
		if x > 0 && grid&(1<<(bit-1)) != 0 {
			adj++
		}
		if x < 4 && grid&(1<<(bit+1)) != 0 {
			adj++
		}
		if y > 0 && grid&(1<<(bit-5)) != 0 {
			adj++
		}
		if y < 4 && grid&(1<<(bit+5)) != 0 {
			adj++
		}

		hasBug := grid&(1<<bit) != 0
		if hasBug && adj == 1 {
			next |= 1 << bit
		} else if !hasBug && (adj == 1 || adj == 2) {
			next |= 1 << bit
		}
	}
	return next
}

func day24Part2(initial uint32, minutes int) uint {
	// Neighbor masks for recursive grids (precomputed)
	type neighbor struct{ same, outer, inner uint32 }
	var neighbors [25]neighbor

	// Outer level: which tile do we check when going off-grid?
	const outerL, outerR = 1 << 11, 1 << 13 // tiles (2,1) and (2,3)
	const outerU, outerD = 1 << 7, 1 << 17  // tiles (1,2) and (3,2)

	// Inner level: which edge/row do we check when entering center?
	const innerL = 1<<0 | 1<<5 | 1<<10 | 1<<15 | 1<<20   // left column
	const innerR = 1<<4 | 1<<9 | 1<<14 | 1<<19 | 1<<24   // right column
	const innerU = 1<<0 | 1<<1 | 1<<2 | 1<<3 | 1<<4      // top row
	const innerD = 1<<20 | 1<<21 | 1<<22 | 1<<23 | 1<<24 // bottom row

	for bit := range 25 {
		if bit == 12 {
			continue // center is recursive portal
		}
		x, y := bit%5, bit/5
		var n neighbor

		// Left neighbor
		if x == 0 {
			n.outer |= outerL
		} else if x == 3 && y == 2 {
			n.inner |= innerR
		} else {
			n.same |= 1 << (bit - 1)
		}

		// Right neighbor
		if x == 4 {
			n.outer |= outerR
		} else if x == 1 && y == 2 {
			n.inner |= innerL
		} else {
			n.same |= 1 << (bit + 1)
		}

		// Up neighbor
		if y == 0 {
			n.outer |= outerU
		} else if x == 2 && y == 3 {
			n.inner |= innerD
		} else {
			n.same |= 1 << (bit - 5)
		}

		// Down neighbor
		if y == 4 {
			n.outer |= outerD
		} else if x == 2 && y == 1 {
			n.inner |= innerU
		} else {
			n.same |= 1 << (bit + 5)
		}

		neighbors[bit] = n
	}

	// Levels array with offset (level 0 at index 201)
	const size = 403
	const offset = 201
	grids := make([]uint32, size)
	next := make([]uint32, size)
	grids[offset] = initial &^ (1 << 12) // clear center bit

	minLvl, maxLvl := 0, 0

	for range minutes {
		lo, hi := minLvl-1, maxLvl+1
		newMin, newMax := hi, lo

		for lvl := lo; lvl <= hi; lvl++ {
			idx := lvl + offset
			grid := grids[idx]
			outer := grids[idx-1]
			inner := grids[idx+1]

			var g uint32
			for bit := range 25 {
				if bit == 12 {
					continue
				}
				n := &neighbors[bit]
				adj := bits.OnesCount32(grid&n.same) +
					bits.OnesCount32(outer&n.outer) +
					bits.OnesCount32(inner&n.inner)

				hasBug := grid&(1<<bit) != 0
				if hasBug && adj == 1 {
					g |= 1 << bit
				} else if !hasBug && (adj == 1 || adj == 2) {
					g |= 1 << bit
				}
			}

			next[idx] = g
			if g != 0 {
				if lvl < newMin {
					newMin = lvl
				}
				if lvl > newMax {
					newMax = lvl
				}
			}
		}

		// Clear old range, swap buffers
		for lvl := lo; lvl <= hi; lvl++ {
			grids[lvl+offset] = 0
		}
		grids, next = next, grids
		minLvl, maxLvl = newMin, newMax
	}

	var total uint
	for lvl := minLvl; lvl <= maxLvl; lvl++ {
		total += uint(bits.OnesCount32(grids[lvl+offset]))
	}
	return total
}
