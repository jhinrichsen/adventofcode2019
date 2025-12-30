package adventofcode2019

import "image"

// Day15 finds the minimum steps to the oxygen system (part1)
// or time to fill with oxygen (part2)
func Day15(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	// Direction: cmd, dx, dy, reverse cmd
	directions := []struct {
		cmd, dx, dy, reverse int
	}{
		{1, 0, -1, 2}, // north
		{2, 0, 1, 1},  // south
		{3, -1, 0, 4}, // west
		{4, 1, 0, 3},  // east
	}

	sendCommand := func(cmd int) int {
		for {
			state := ic.Step()
			switch state {
			case needsInput:
				ic.Input(cmd)
			case hasOutput:
				return ic.Output()
			case halted:
				return -1
			}
		}
	}

	// Explore entire maze, track distance to each cell
	type cell struct {
		pos   image.Point
		steps uint
	}

	dist := make(map[image.Point]uint)
	grid := make(map[image.Point]int) // 0=wall, 1=open, 2=oxygen

	start := image.Point{}
	dist[start] = 0
	grid[start] = 1

	queue := []cell{{pos: start, steps: 0}}
	currentPos := start
	var oxygenPos image.Point
	var oxygenSteps uint

	// BFS pathfinding to move droid
	moveTo := func(from, to image.Point) {
		if from == to {
			return
		}
		type node struct {
			pos image.Point
			cmd int
		}
		parent := make(map[image.Point]node)
		parent[from] = node{pos: from}
		q := []image.Point{from}

		for head := 0; head < len(q); head++ {
			cur := q[head]
			if cur == to {
				break
			}
			for _, dir := range directions {
				next := image.Point{X: cur.X + dir.dx, Y: cur.Y + dir.dy}
				if _, seen := parent[next]; seen {
					continue
				}
				if tile, ok := grid[next]; !ok || tile == 0 {
					continue
				}
				parent[next] = node{pos: cur, cmd: dir.cmd}
				q = append(q, next)
			}
		}

		// Reconstruct path
		var path []int
		for p := to; p != from; {
			n := parent[p]
			path = append(path, n.cmd)
			p = n.pos
		}
		// Execute in reverse
		for i := len(path) - 1; i >= 0; i-- {
			sendCommand(path[i])
		}
	}

	// Explore
	for head := 0; head < len(queue); head++ {
		cur := queue[head]

		if currentPos != cur.pos {
			moveTo(currentPos, cur.pos)
			currentPos = cur.pos
		}

		for _, dir := range directions {
			next := image.Point{X: cur.pos.X + dir.dx, Y: cur.pos.Y + dir.dy}
			if _, seen := dist[next]; seen {
				continue
			}

			status := sendCommand(dir.cmd)
			dist[next] = cur.steps + 1
			grid[next] = status

			if status == 0 { // wall
				continue
			}

			if status == 2 { // oxygen
				oxygenPos = next
				oxygenSteps = cur.steps + 1
				if part1 {
					return oxygenSteps, nil
				}
			}

			queue = append(queue, cell{pos: next, steps: cur.steps + 1})
			sendCommand(dir.reverse) // move back
		}
	}

	// Part 2: BFS from oxygen position to find max fill time
	filled := make(map[image.Point]bool)
	filled[oxygenPos] = true
	front := []image.Point{oxygenPos}
	var minutes uint

	for len(front) > 0 {
		var nextFront []image.Point
		for _, pos := range front {
			for _, dir := range directions {
				next := image.Point{X: pos.X + dir.dx, Y: pos.Y + dir.dy}
				if filled[next] {
					continue
				}
				if tile := grid[next]; tile == 0 {
					continue
				}
				filled[next] = true
				nextFront = append(nextFront, next)
			}
		}
		if len(nextFront) > 0 {
			minutes++
		}
		front = nextFront
	}

	return minutes, nil
}
