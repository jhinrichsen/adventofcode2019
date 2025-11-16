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

// Day3Part1 computes the minimal manhattan distance of two crossing wires
func Day3Part1(wires []string) (int, error) {
	// Use sparse maps instead of dense 2D arrays
	wireMaps := make([]map[Point]bool, len(wires))
	for i := range wireMaps {
		wireMaps[i] = make(map[Point]bool)
	}

	// Mark all positions for each wire
	for i, wire := range wires {
		x, y := 0, 0
		ws := strings.Split(wire, ",")
		for _, w := range ws {
			d, n, err := Parse(w)
			if err != nil {
				return 0, err
			}
			for range n {
				switch d {
				case Up:
					y++
				case Down:
					y--
				case Right:
					x++
				case Left:
					x--
				}
				wireMaps[i][Point{x, y}] = true
			}
		}
	}

	// Find intersections and minimum Manhattan distance
	min := math.MaxInt
	for p := range wireMaps[0] {
		// Check if this point is also in other wires
		allWiresHit := true
		for i := 1; i < len(wireMaps); i++ {
			if !wireMaps[i][p] {
				allWiresHit = false
				break
			}
		}
		if allWiresHit {
			manhattan := absDay3(p.x) + absDay3(p.y)
			if manhattan > 0 && manhattan < min {
				min = manhattan
			}
		}
	}
	return min, nil
}

func absDay3(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Board creates a two dimensional arrray. The board created will have double
// width and double height, so that no negative indices are used when wiring.
func Board(x, y int) [][]uint {
	b := make([][]uint, y)
	for i := 0; i < y; i++ {
		b[i] = make([]uint, x)
	}
	return b
}

// MaxSize calculates the size of wirings
func MaxSize(wirings []string) (int, int, error) {
	maxX, maxY := 0, 0
	for _, wiring := range wirings {
		x, y, err := Size(wiring)
		if err != nil {
			return maxX, maxY, err
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

// Walk dots all wires onto a board
func Walk(b [][]uint, wires []string) error {
	for i, wire := range wires {
		id := uint(1 << i)
		// center in middle of board
		px, py := len(b[0])/2, len(b)/2
		ws := strings.Split(wire, ",")
		for _, w := range ws {
			d, n, err := Parse(w)
			if err != nil {
				return err
			}
			switch d {
			case Up:
				for y := py; y < py+n; y++ {
					b[y][px] |= id
				}
				py += n
			case Down:
				for y := py; y < py-n; y-- {
					b[y][px] |= id
				}
				py -= n
			case Right:
				for x := px; x < px+n; x++ {
					b[py][x] |= id
				}
				px += n
			case Left:
				for x := px; x < px-n; x-- {
					b[py][x] |= id
				}
				px -= n
			}
		}
	}
	return nil
}

// Point holds a (x/y) position
type Point struct {
	x, y int
}

// Marker holds steps needed to reach position (x/y)
type marker map[Point]int

// Store will save the first, and only the first, number of steps for (x/y)
func storeOnce(m marker, x, y, steps int) {
	p := Point{x, y}
	// do nothing if position already visited
	if _, ok := m[p]; ok {
		return
	}
	m[p] = steps
}

// Day3Part2 computes the minimal combined steps for intersections
func Day3Part2(wires []string) (int, error) {
	boards := make([]marker, len(wires))
	for i := 0; i < len(boards); i++ {
		boards[i] = make(marker)
	}

	// Transform wires into marker maps
	for i, wire := range wires {
		board := boards[i]
		x, y, steps := 0, 0, 0
		ws := strings.Split(wire, ",")
		for _, w := range ws {
			d, n, err := Parse(w)
			if err != nil {
				return -1, err
			}
			for j := 0; j < n; j++ {
				storeOnce(board, x, y, steps)
				steps++
				// Next time i write this part i will decode the
				// RLE (run length encoding) into distinct (x/y)
				// deltas - turtle graphics ;-)
				switch d {
				case Right:
					x++
				case Left:
					x--
				case Up:
					y++
				case Down:
					y--
				}
			}
		}
	}

	// Find all intersections, and the sum of steps
	intersections := make(marker)
	// use wiring of first board as reference
	for refpos, refsteps := range boards[0] {
		// ignore center point (0/0)
		if refpos.x == 0 && refpos.y == 0 {
			continue
		}
		// check for collision
		for _, board := range boards[1:] {
			if steps, ok := board[refpos]; ok {
				intersections[refpos] = refsteps + steps
			}
		}
	}

	// find lowest intersection
	min := math.MaxInt32
	for _, sum := range intersections {
		if sum < min {
			min = sum
		}
	}

	return min, nil
}
