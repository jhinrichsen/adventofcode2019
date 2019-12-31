package adventofcode2019

import (
	"fmt"
	"math"
	"math/bits"
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
	maxX, maxY, err := MaxSize(wires)
	if err != nil {
		return 0, err
	}

	// Create a board that is large enough to always fit the wiring
	b := Board((1+maxX)*2, (1+maxY)*2)
	if err := Walk(b, wires); err != nil {
		return 0, err
	}

	min := MinimalDistance(b)
	return min, nil
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

// MinimalDistance returns minimal manhattan distance of all crossings
func MinimalDistance(b [][]uint) int {
	min := math.MaxInt64
	lx, ly := len(b[0]), len(b)
	centerX, centerY := lx/2, ly/2
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	for y := 0; y < ly; y++ {
		for x := 0; x < lx; x++ {
			// More than one bit set?
			if bits.OnesCount(b[y][x]) > 1 {
				manhattanDistance := abs(x-centerX) + abs(y-centerY)
				// ignore center spot itself
				if manhattanDistance > 0 && manhattanDistance < min {
					min = manhattanDistance
				}
			}
		}
	}
	return min
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
	return Up, 0, fmt.Errorf("Illegal path %q", path)
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
