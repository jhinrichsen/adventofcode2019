package adventofcode2019

import (
	"bytes"
)

// Day18 solves the "Many-Worlds Interpretation" puzzle.
// It finds the minimum number of steps to collect all keys in a maze with doors.
func Day18(input []byte, part1 bool) uint {
	maze := parseMaze(input)
	if part1 {
		return solvePart1(maze)
	}
	return solvePart2(maze)
}

type Maze struct {
	grid   [][]byte
	dimX   int
	dimY   int
	startX int
	startY int
	keys   map[byte]bool
}

func parseMaze(input []byte) Maze {
	lines := bytes.Split(bytes.TrimSpace(input), []byte{'\n'})
	dimY := len(lines)
	if dimY == 0 {
		return Maze{}
	}
	dimX := len(lines[0])

	maze := Maze{
		grid: make([][]byte, dimY),
		dimX: dimX,
		dimY: dimY,
		keys: make(map[byte]bool),
	}

	for y := range dimY {
		maze.grid[y] = make([]byte, dimX)
		for x := range dimX {
			if x < len(lines[y]) {
				cell := lines[y][x]
				maze.grid[y][x] = cell

				if cell == '@' {
					maze.startX = x
					maze.startY = y
				} else if cell >= 'a' && cell <= 'z' {
					maze.keys[cell] = true
				}
			}
		}
	}

	return maze
}

func solvePart1(maze Maze) uint {
	// TODO: Implement BFS/Dijkstra to find shortest path collecting all keys
	return 0
}

func solvePart2(maze Maze) uint {
	// TODO: Implement part 2
	return 0
}
