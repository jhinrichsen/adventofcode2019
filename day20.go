package adventofcode2019

import (
	"bytes"
	"image"
)

// Day20 solves the "Donut Maze" puzzle.
func Day20(input []byte, part1 bool) uint {
	maze := parseMaze20(input)
	if part1 {
		return solveMaze20Part1(maze)
	}
	return 0 // TODO: Part 2
}

type Maze20 struct {
	grid    [][]byte
	dimX    int
	dimY    int
	portals map[string][]image.Point // portal name -> list of positions
	start   image.Point
	end     image.Point
}

func parseMaze20(input []byte) Maze20 {
	lines := bytes.Split(input, []byte{'\n'})
	dimY := len(lines)
	dimX := 0
	for _, line := range lines {
		if len(line) > dimX {
			dimX = len(line)
		}
	}

	maze := Maze20{
		grid:    make([][]byte, dimY),
		dimX:    dimX,
		dimY:    dimY,
		portals: make(map[string][]image.Point),
	}

	// Parse grid
	for y := range dimY {
		maze.grid[y] = make([]byte, dimX)
		for x := range dimX {
			if x < len(lines[y]) {
				maze.grid[y][x] = lines[y][x]
			} else {
				maze.grid[y][x] = ' '
			}
		}
	}

	// Find portals
	for y := range dimY {
		for x := range dimX {
			if maze.grid[y][x] >= 'A' && maze.grid[y][x] <= 'Z' {
				// Check if this is the start of a portal label
				maze.checkPortal(x, y)
			}
		}
	}

	return maze
}

func (m *Maze20) checkPortal(x, y int) {
	cell := m.grid[y][x]

	// Check horizontal portal (this letter + next letter)
	if x+1 < m.dimX && m.grid[y][x+1] >= 'A' && m.grid[y][x+1] <= 'Z' {
		label := string([]byte{cell, m.grid[y][x+1]})

		// Find the . adjacent to this portal
		var portalPos image.Point
		if x > 0 && m.grid[y][x-1] == '.' {
			portalPos = image.Point{X: x - 1, Y: y}
		} else if x+2 < m.dimX && m.grid[y][x+2] == '.' {
			portalPos = image.Point{X: x + 2, Y: y}
		} else {
			return
		}

		if label == "AA" {
			m.start = portalPos
		} else if label == "ZZ" {
			m.end = portalPos
		} else {
			m.portals[label] = append(m.portals[label], portalPos)
		}
	}

	// Check vertical portal (this letter + next letter)
	if y+1 < m.dimY && m.grid[y+1][x] >= 'A' && m.grid[y+1][x] <= 'Z' {
		label := string([]byte{cell, m.grid[y+1][x]})

		// Find the . adjacent to this portal
		var portalPos image.Point
		if y > 0 && m.grid[y-1][x] == '.' {
			portalPos = image.Point{X: x, Y: y - 1}
		} else if y+2 < m.dimY && m.grid[y+2][x] == '.' {
			portalPos = image.Point{X: x, Y: y + 2}
		} else {
			return
		}

		if label == "AA" {
			m.start = portalPos
		} else if label == "ZZ" {
			m.end = portalPos
		} else {
			m.portals[label] = append(m.portals[label], portalPos)
		}
	}
}

func solveMaze20Part1(maze Maze20) uint {
	// BFS from start to end
	type state struct {
		pos   image.Point
		steps uint
	}

	queue := []state{{maze.start, 0}}
	visited := make(map[image.Point]bool)
	visited[maze.start] = true

	dirs := []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.pos == maze.end {
			return curr.steps
		}

		// Try normal moves
		for _, dir := range dirs {
			newPos := image.Point{X: curr.pos.X + dir.X, Y: curr.pos.Y + dir.Y}

			if newPos.X < 0 || newPos.X >= maze.dimX || newPos.Y < 0 || newPos.Y >= maze.dimY {
				continue
			}

			if maze.grid[newPos.Y][newPos.X] != '.' {
				continue
			}

			if visited[newPos] {
				continue
			}

			visited[newPos] = true
			queue = append(queue, state{newPos, curr.steps + 1})
		}

		// Try portal teleport
		for _, positions := range maze.portals {
			if len(positions) == 2 {
				if positions[0] == curr.pos && !visited[positions[1]] {
					visited[positions[1]] = true
					queue = append(queue, state{positions[1], curr.steps + 1})
				} else if positions[1] == curr.pos && !visited[positions[0]] {
					visited[positions[0]] = true
					queue = append(queue, state{positions[0], curr.steps + 1})
				}
			}
		}
	}

	return 0
}
