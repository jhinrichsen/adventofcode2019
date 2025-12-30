package adventofcode2019

import "image"

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
func Day15(program []byte, part1 bool) (uint, error) {
	ic, err := NewIntcode(program)
	if err != nil {
		return 0, err
	}

	if part1 {
		return exploreAndFindOxygen(ic), nil
	}
	return fillWithOxygen(ic), nil
}

// directions holds command and delta for each direction
var directions = []struct {
	cmd     int
	dx, dy  int
	reverse int
}{
	{North, 0, -1, South},
	{South, 0, 1, North},
	{West, -1, 0, East},
	{East, 1, 0, West},
}

// sendCommand sends a command and returns the status
func sendCommand(ic *Intcode, cmd int) int {
	for {
		state := ic.Step()
		switch state {
		case NeedsInput:
			ic.Input(cmd)
		case HasOutput:
			return ic.Output()
		case Halted:
			return -1
		}
	}
}

func exploreAndFindOxygen(ic *Intcode) uint {
	type state struct {
		pos   image.Point
		steps uint
	}

	visited := make(map[image.Point]bool)
	grid := make(map[image.Point]int)

	start := image.Point{X: 0, Y: 0}
	visited[start] = true
	grid[start] = Moved

	queue := make([]state, 0, 1000)
	queue = append(queue, state{pos: start, steps: 0})
	currentPos := start

	// moveTo moves the droid from current position to target
	moveTo := func(from, to image.Point) bool {
		if from == to {
			return true
		}
		path := findPath(from, to, grid)
		for _, cmd := range path {
			status := sendCommand(ic, cmd)
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

		if currentPos != current.pos {
			if !moveTo(currentPos, current.pos) {
				continue
			}
			currentPos = current.pos
		}

		for _, dir := range directions {
			nextPos := image.Point{
				X: current.pos.X + dir.dx,
				Y: current.pos.Y + dir.dy,
			}

			if visited[nextPos] {
				continue
			}

			status := sendCommand(ic, dir.cmd)
			visited[nextPos] = true
			grid[nextPos] = status

			if status == HitWall {
				continue
			}

			if status == FoundOxygen {
				return current.steps + 1
			}

			queue = append(queue, state{
				pos:   nextPos,
				steps: current.steps + 1,
			})

			// Move back
			sendCommand(ic, dir.reverse)
		}
	}

	return 0
}

// findPath returns sequence of commands to move from 'from' to 'to'
func findPath(from, to image.Point, grid map[image.Point]int) []int {
	if from == to {
		return nil
	}

	type parentInfo struct {
		pos image.Point
		cmd int
	}

	parent := make(map[image.Point]parentInfo)
	queue := make([]image.Point, 0, 100)
	queue = append(queue, from)
	parent[from] = parentInfo{pos: from, cmd: 0}

	pathDirs := []struct {
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

		for _, dir := range pathDirs {
			nextPos := image.Point{
				X: current.X + dir.dx,
				Y: current.Y + dir.dy,
			}

			if _, vis := parent[nextPos]; vis {
				continue
			}

			if tile, ok := grid[nextPos]; !ok || tile == HitWall {
				continue
			}

			parent[nextPos] = parentInfo{pos: current, cmd: dir.cmd}

			if nextPos == to {
				var path []int
				for p := nextPos; p != from; {
					info := parent[p]
					path = append(path, info.cmd)
					p = info.pos
				}
				for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
					path[i], path[j] = path[j], path[i]
				}
				return path
			}

			queue = append(queue, nextPos)
		}
	}

	return nil
}

func fillWithOxygen(ic *Intcode) uint {
	// First, explore the entire maze
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

	// Explore entire maze
	head := 0
	for head < len(queue) {
		current := queue[head]
		head++

		if currentPos != current.pos {
			path := findPath(currentPos, current.pos, grid)
			for _, cmd := range path {
				sendCommand(ic, cmd)
			}
			currentPos = current.pos
		}

		for _, dir := range directions {
			nextPos := image.Point{
				X: current.pos.X + dir.dx,
				Y: current.pos.Y + dir.dy,
			}

			if visited[nextPos] {
				continue
			}

			status := sendCommand(ic, dir.cmd)
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
			sendCommand(ic, dir.reverse)
		}
	}

	// Now simulate oxygen spreading using BFS
	oxygenVisited := make(map[image.Point]bool)
	oxygenQueue := []state{{pos: oxygenPos}}
	oxygenVisited[oxygenPos] = true

	var minutes uint

	for len(oxygenQueue) > 0 {
		levelSize := len(oxygenQueue)

		for i := 0; i < levelSize; i++ {
			current := oxygenQueue[0]
			oxygenQueue = oxygenQueue[1:]

			for _, dir := range directions {
				nextPos := image.Point{
					X: current.pos.X + dir.dx,
					Y: current.pos.Y + dir.dy,
				}

				if oxygenVisited[nextPos] {
					continue
				}

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
