package adventofcode2019

import (
	"bytes"
	"image"
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

type state struct {
	pos  image.Point
	keys uint32 // bitmask of collected keys
}

type queueItem struct {
	state
	steps uint
}

func solvePart1(maze Maze) uint {
	// Build key-to-bit mapping (deterministic order)
	keyList := make([]byte, 0, len(maze.keys))
	for key := range maze.keys {
		keyList = append(keyList, key)
	}
	// Sort to ensure deterministic bit assignment
	for i := range len(keyList) - 1 {
		for j := i + 1; j < len(keyList); j++ {
			if keyList[i] > keyList[j] {
				keyList[i], keyList[j] = keyList[j], keyList[i]
			}
		}
	}

	keyBits := make(map[byte]uint32)
	for i, key := range keyList {
		keyBits[key] = 1 << i
	}
	allKeys := (uint32(1) << len(keyList)) - 1

	// BFS with state = (position, collected keys)
	start := image.Point{X: maze.startX, Y: maze.startY}
	queue := []queueItem{{state{start, 0}, 0}}
	visited := make(map[state]bool)
	visited[state{start, 0}] = true

	dirs := []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		// Try all 4 directions
		for _, dir := range dirs {
			newPos := image.Point{X: item.pos.X + dir.X, Y: item.pos.Y + dir.Y}

			// Check bounds
			if newPos.X < 0 || newPos.X >= maze.dimX || newPos.Y < 0 || newPos.Y >= maze.dimY {
				continue
			}

			cell := maze.grid[newPos.Y][newPos.X]

			// Wall
			if cell == '#' {
				continue
			}

			// Door - check if we have the key
			if cell >= 'A' && cell <= 'Z' {
				keyNeeded := keyBits[cell-'A'+'a']
				if item.keys&keyNeeded == 0 {
					continue // don't have the key
				}
			}

			// Calculate new key state
			newKeys := item.keys
			if cell >= 'a' && cell <= 'z' {
				newKeys |= keyBits[cell]
			}

			// Check if we've collected all keys
			if newKeys == allKeys {
				return item.steps + 1
			}

			newState := state{newPos, newKeys}
			if visited[newState] {
				continue
			}

			visited[newState] = true
			queue = append(queue, queueItem{newState, item.steps + 1})
		}
	}

	return 0 // no solution found
}

func solvePart2(maze Maze) uint {
	// TODO: Implement part 2
	return 0
}
