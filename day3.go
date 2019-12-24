package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// Day3 computes the minimal manhattan distance of two crossing wires
func Day3(wires []string) (int, error) {
	return 0, nil
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
