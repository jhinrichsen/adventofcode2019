package adventofcode2019

import (
	"bytes"
	"image"
)

// Direction commands for the repair droid
const (
	North = 1
	South = 2
	West  = 3
	East  = 4
)

// Status codes from the droid
const (
	HitWall     = 0
	Moved       = 1
	FoundOxygen = 2
)

// Day15 finds the minimum steps to the oxygen system (part1)
// or time to fill with oxygen (part2)
func Day15(program []byte, part1 bool) uint {
	code := MustSplit(string(bytes.TrimSpace(program)))

	if part1 {
		return exploreAndFindOxygen(code)
	}
	return fillWithOxygen(code)
}

func exploreAndFindOxygen(code IntCode) uint {
	// Create channels for Intcode computer
	input := make(chan int, 10)
	output := make(chan int, 10)

	// Run the Intcode computer in a goroutine
	prog := code.Copy()
	go Day5(prog, input, output)

	// BFS to explore the maze
	type state struct {
		pos   image.Point
		steps uint
	}

	// Track visited positions and the grid
	visited := make(map[image.Point]bool)
	grid := make(map[image.Point]int) // stores tile type

	// Start at origin
	start := image.Point{X: 0, Y: 0}
	visited[start] = true
	grid[start] = Moved

	// Queue for BFS
	queue := make([]state, 0, 1000)
	queue = append(queue, state{pos: start, steps: 0})
	currentPos := start

	// Direction mappings
	directions := []struct {
		cmd     int
		dx, dy  int
		reverse int
	}{
		{North, 0, -1, South},
		{South, 0, 1, North},
		{West, -1, 0, East},
		{East, 1, 0, West},
	}

	// Move droid from current position to target position
	moveTo := func(from, to image.Point) bool {
		// Find path using BFS on known grid
		if from == to {
			return true
		}

		// Simple path finding - for each step, move closer
		path := findPath(from, to, grid)
		for _, cmd := range path {
			input <- cmd
			status := <-output
			if status == HitWall {
				return false
			}
		}
		return true
	}

	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		// Move droid to current position if needed
		if currentPos != current.pos {
			if !moveTo(currentPos, current.pos) {
				continue
			}
			currentPos = current.pos
		}

		// Try all four directions
		for _, dir := range directions {
			nextPos := image.Point{
				X: current.pos.X + dir.dx,
				Y: current.pos.Y + dir.dy,
			}

			// Skip if already visited
			if visited[nextPos] {
				continue
			}

			// Send movement command
			input <- dir.cmd
			status := <-output

			// Mark as visited and update grid
			visited[nextPos] = true
			grid[nextPos] = status

			if status == HitWall {
				// Hit a wall, don't add to queue
				continue
			}

			// Successfully moved (status 1 or 2)
			if status == FoundOxygen {
				// Found the oxygen system!
				close(input)
				return current.steps + 1
			}

			// Add to queue for further exploration
			queue = append(queue, state{
				pos:   nextPos,
				steps: current.steps + 1,
			})

			// Move back to continue exploring from current position
			input <- dir.reverse
			<-output
		}
	}

	close(input)
	return 0
}

// findPath returns sequence of commands to move from 'from' to 'to'
func findPath(from, to image.Point, grid map[image.Point]int) []int {
	if from == to {
		return nil
	}

	type pathState struct {
		pos  image.Point
		path []int
	}

	visited := make(map[image.Point]bool)
	queue := make([]pathState, 0, 100)
	queue = append(queue, pathState{pos: from, path: nil})
	visited[from] = true

	directions := []struct {
		cmd    int
		dx, dy int
	}{
		{North, 0, -1},
		{South, 0, 1},
		{West, -1, 0},
		{East, 1, 0},
	}

	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		for _, dir := range directions {
			nextPos := image.Point{
				X: current.pos.X + dir.dx,
				Y: current.pos.Y + dir.dy,
			}

			if visited[nextPos] {
				continue
			}

			// Only move through known non-wall tiles
			if tile, ok := grid[nextPos]; !ok || tile == HitWall {
				continue
			}

			visited[nextPos] = true
			newPath := append([]int{}, current.path...)
			newPath = append(newPath, dir.cmd)

			if nextPos == to {
				return newPath
			}

			queue = append(queue, pathState{
				pos:  nextPos,
				path: newPath,
			})
		}
	}

	return nil
}

func fillWithOxygen(code IntCode) uint {
	// First, explore the entire maze
	input := make(chan int, 10)
	output := make(chan int, 10)

	prog := code.Copy()
	go Day5(prog, input, output)

	// Build complete map of the maze
	grid := make(map[image.Point]int)
	visited := make(map[image.Point]bool)

	start := image.Point{X: 0, Y: 0}
	visited[start] = true
	grid[start] = Moved

	type state struct {
		pos image.Point
	}

	queue := make([]state, 0, 1000)
	queue = append(queue, state{pos: start})
	currentPos := start
	var oxygenPos image.Point

	directions := []struct {
		cmd     int
		dx, dy  int
		reverse int
	}{
		{North, 0, -1, South},
		{South, 0, 1, North},
		{West, -1, 0, East},
		{East, 1, 0, West},
	}

	// Explore entire maze
	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		// Move droid to current position if needed
		if currentPos != current.pos {
			path := findPath(currentPos, current.pos, grid)
			for _, cmd := range path {
				input <- cmd
				<-output
			}
			currentPos = current.pos
		}

		// Try all four directions
		for _, dir := range directions {
			nextPos := image.Point{
				X: current.pos.X + dir.dx,
				Y: current.pos.Y + dir.dy,
			}

			if visited[nextPos] {
				continue
			}

			input <- dir.cmd
			status := <-output

			visited[nextPos] = true
			grid[nextPos] = status

			if status == HitWall {
				continue
			}

			if status == FoundOxygen {
				oxygenPos = nextPos
			}

			queue = append(queue, state{pos: nextPos})

			// Move back
			input <- dir.reverse
			<-output
		}
	}

	close(input)

	// Now simulate oxygen spreading using BFS
	// Oxygen spreads to all adjacent non-wall tiles
	oxygenVisited := make(map[image.Point]bool)
	oxygenQueue := []state{{pos: oxygenPos}}
	oxygenVisited[oxygenPos] = true

	var minutes uint

	for len(oxygenQueue) > 0 {
		// Process all positions at current minute
		levelSize := len(oxygenQueue)

		for i := 0; i < levelSize; i++ {
			current := oxygenQueue[0]
			oxygenQueue = oxygenQueue[1:]

			// Spread to all 4 directions
			for _, dir := range directions {
				nextPos := image.Point{
					X: current.pos.X + dir.dx,
					Y: current.pos.Y + dir.dy,
				}

				if oxygenVisited[nextPos] {
					continue
				}

				// Only spread to non-wall tiles
				if tile, ok := grid[nextPos]; !ok || tile == HitWall {
					continue
				}

				oxygenVisited[nextPos] = true
				oxygenQueue = append(oxygenQueue, state{pos: nextPos})
			}
		}

		if len(oxygenQueue) > 0 {
			minutes++
		}
	}

	return minutes
}
