package adventofcode2019

import (
	"bytes"
	"fmt"
	"image"
)

// Day20 solves the "Donut Maze" puzzle.
//
// Part 2 Note: Initial implementation had off-by-one error in isOuterPortal()
// which caused incorrect answer of 2520. Fixed by changing boundary checks
// from `>` to `>=`. Correct answer: 7844.
func Day20(input []byte, part1 bool) uint {
	maze := parseMaze20(input)
	if part1 {
		return solveMaze20Part1(maze)
	}
	return solveMaze20Part2(maze)
}

type Maze20 struct {
	grid      [][]byte
	dimX      int
	dimY      int
	portals   map[string][]image.Point // portal name -> list of positions
	start     image.Point
	end       image.Point
	innerEdge image.Point // Track inner boundary for determining inner/outer portals
	outerEdge image.Point
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

	// Find inner/outer boundaries
	// Outer boundary is near edges, inner boundary is around the donut hole
	maze.findBoundaries()

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

func (m *Maze20) findBoundaries() {
	// Find the donut hole (area with no passages in the middle)
	minHoleX, maxHoleX := m.dimX, 0
	minHoleY, maxHoleY := m.dimY, 0

	for y := range m.dimY {
		for x := range m.dimX {
			if m.grid[y][x] == ' ' && x > 2 && x < m.dimX-2 && y > 2 && y < m.dimY-2 {
				if x < minHoleX {
					minHoleX = x
				}
				if x > maxHoleX {
					maxHoleX = x
				}
				if y < minHoleY {
					minHoleY = y
				}
				if y > maxHoleY {
					maxHoleY = y
				}
			}
		}
	}

	m.innerEdge = image.Point{X: minHoleX, Y: minHoleY}
	m.outerEdge = image.Point{X: maxHoleX, Y: maxHoleY}
}

func (m *Maze20) isOuterPortal(pos image.Point) bool {
	// Check if position is on outer edge of maze (near the border)
	return pos.X <= 2 || pos.X >= m.dimX-3 || pos.Y <= 2 || pos.Y >= m.dimY-4
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

func solveMaze20Part2(maze Maze20) uint {
	// BFS from start to end with recursive levels
	type state struct {
		pos   image.Point
		level int
		steps uint
	}

	queue := []state{{maze.start, 0, 0}}
	visited := make(map[string]bool)
	visitKey := func(pos image.Point, level int) string {
		return fmt.Sprintf("%d,%d,%d", pos.X, pos.Y, level)
	}
	visited[visitKey(maze.start, 0)] = true

	dirs := []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	const maxLevel = 500 // Reasonable depth limit

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check if we reached the end at level 0
		if curr.pos == maze.end && curr.level == 0 {
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

			key := visitKey(newPos, curr.level)
			if visited[key] {
				continue
			}

			visited[key] = true
			queue = append(queue, state{newPos, curr.level, curr.steps + 1})
		}

		// Try portal teleport
		for _, positions := range maze.portals {
			if len(positions) != 2 {
				continue
			}

			var fromPos, toPos image.Point
			var isOuter bool

			if positions[0] == curr.pos {
				fromPos = positions[0]
				toPos = positions[1]
				isOuter = maze.isOuterPortal(fromPos)
			} else if positions[1] == curr.pos {
				fromPos = positions[1]
				toPos = positions[0]
				isOuter = maze.isOuterPortal(fromPos)
			} else {
				continue
			}

			// Determine new level
			newLevel := curr.level
			if isOuter {
				// Outer portal: go up a level (outward)
				newLevel--
				if newLevel < 0 {
					// Can't go beyond level 0
					continue
				}
			} else {
				// Inner portal: go down a level (inward)
				newLevel++
				if newLevel > maxLevel {
					// Don't go too deep
					continue
				}
			}

			key := visitKey(toPos, newLevel)
			if visited[key] {
				continue
			}

			visited[key] = true
			queue = append(queue, state{toPos, newLevel, curr.steps + 1})
		}
	}

	return 0
}
