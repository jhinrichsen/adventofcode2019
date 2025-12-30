package adventofcode2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Direction type
type Direction int

const (
	// Up represents 'U'
	Up Direction = iota
	// Right represents 'R'
	Right
	// Down represents 'D'
	Down
	// Left represents 'L'
	Left
)

// segment represents a horizontal or vertical line segment
type segment struct {
	x1, y1, x2, y2 int  // endpoints
	steps          int  // cumulative steps at start of segment
	horizontal     bool // true if horizontal (y1==y2)
}

// Day03 computes the minimal manhattan distance or minimal combined steps of two crossing wires
func Day03(wires []string, part1 bool) uint {
	// Parse both wires into segments
	segs1 := parseWireSegments(wires[0])
	segs2 := parseWireSegments(wires[1])

	minResult := math.MaxInt32

	// Check all segment pairs for intersections
	for _, s1 := range segs1 {
		for _, s2 := range segs2 {
			// Only horizontal-vertical pairs can intersect
			if s1.horizontal == s2.horizontal {
				continue
			}

			var hSeg, vSeg segment
			if s1.horizontal {
				hSeg, vSeg = s1, s2
			} else {
				hSeg, vSeg = s2, s1
			}

			// Check if they intersect
			// hSeg: horizontal line at y=hSeg.y1, x from min(x1,x2) to max(x1,x2)
			// vSeg: vertical line at x=vSeg.x1, y from min(y1,y2) to max(y1,y2)
			hMinX, hMaxX := hSeg.x1, hSeg.x2
			if hMinX > hMaxX {
				hMinX, hMaxX = hMaxX, hMinX
			}
			vMinY, vMaxY := vSeg.y1, vSeg.y2
			if vMinY > vMaxY {
				vMinY, vMaxY = vMaxY, vMinY
			}

			// Intersection point would be (vSeg.x1, hSeg.y1)
			ix, iy := vSeg.x1, hSeg.y1

			// Check if intersection point is within both segments
			if ix >= hMinX && ix <= hMaxX && iy >= vMinY && iy <= vMaxY {
				// Skip origin
				if ix == 0 && iy == 0 {
					continue
				}

				if part1 {
					dist := abs03(ix) + abs03(iy)
					if dist < minResult {
						minResult = dist
					}
				} else {
					// Calculate steps to intersection for each wire
					steps1 := s1.steps + abs03(ix-s1.x1) + abs03(iy-s1.y1)
					steps2 := s2.steps + abs03(ix-s2.x1) + abs03(iy-s2.y1)
					totalSteps := steps1 + steps2
					if totalSteps < minResult {
						minResult = totalSteps
					}
				}
			}
		}
	}

	return uint(minResult)
}

// parseWireSegments parses a wire string into segments using inline parsing
func parseWireSegments(wire string) []segment {
	// Estimate capacity: roughly one segment per 5 chars
	segs := make([]segment, 0, len(wire)/5+1)
	x, y, steps := 0, 0, 0

	i := 0
	for i < len(wire) {
		// Parse direction
		dir := wire[i]
		i++

		// Parse number
		num := 0
		for i < len(wire) && wire[i] >= '0' && wire[i] <= '9' {
			num = num*10 + int(wire[i]-'0')
			i++
		}

		// Skip comma
		if i < len(wire) && wire[i] == ',' {
			i++
		}

		// Create segment
		x1, y1 := x, y
		switch dir {
		case 'R':
			x += num
		case 'L':
			x -= num
		case 'U':
			y += num
		case 'D':
			y -= num
		}

		segs = append(segs, segment{
			x1:         x1,
			y1:         y1,
			x2:         x,
			y2:         y,
			steps:      steps,
			horizontal: dir == 'R' || dir == 'L',
		})
		steps += num
	}

	return segs
}

func abs03(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Day3Part1 computes the minimal manhattan distance of two crossing wires
func Day3Part1(wires []string) (int, error) {
	return int(Day03(wires, true)), nil
}

// Parse splits a path such as U32 into a direction North and a length 32
func Parse(path string) (Direction, int, error) {
	n, err := strconv.Atoi(path[1:])
	if err != nil {
		return Up, 0, err
	}
	switch path[0] {
	case 'U':
		return Up, n, nil
	case 'R':
		return Right, n, nil
	case 'D':
		return Down, n, nil
	case 'L':
		return Left, n, nil
	}
	return Up, 0, fmt.Errorf("illegal path: %q", path)
}

// Size calculates width and height of a wiring
func Size(wiring string) (int, int, error) {
	x, y, maxX, maxY := 0, 0, 0, 0
	for _, wire := range strings.Split(wiring, ",") {
		d, n, err := Parse(wire)
		if err != nil {
			return maxX, maxY, err
		}

		switch d {
		case Up:
			y += n
		case Down:
			y -= n
		case Right:
			x += n
		case Left:
			x -= n
		}

		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	return maxX, maxY, nil
}

// Day3Part2 computes the minimal combined steps for intersections
func Day3Part2(wires []string) (int, error) {
	return int(Day03(wires, false)), nil
}
