package adventofcode2019

import (
	"bytes"
	"container/heap"
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
	grid    [][]byte
	dimX    int
	dimY    int
	startX  int
	startY  int
	keys    map[byte]bool
	robots4 [4]image.Point // For part 2
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

	// Find key positions
	keyPos := make(map[byte]image.Point)
	for y := range maze.dimY {
		for x := range maze.dimX {
			cell := maze.grid[y][x]
			if cell >= 'a' && cell <= 'z' {
				keyPos[cell] = image.Point{X: x, Y: y}
			}
		}
	}

	// Precompute distances from start and each key to all other keys
	distances := make(map[byte]map[byte]pathInfo)
	startPos := image.Point{X: maze.startX, Y: maze.startY}
	distances['@'] = bfsFrom(maze, startPos, keyBits)
	for key, pos := range keyPos {
		distances[key] = bfsFrom(maze, pos, keyBits)
	}

	// Dijkstra on key graph: state = (current key, collected keys)
	type keyState struct {
		at        byte
		collected uint32
	}

	dist := make(map[keyState]uint)
	startState := keyState{'@', 0}
	dist[startState] = 0

	// Use proper heap
	pq := &day18Heap{{startState, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(day18HeapItem)

		// Skip if we've found a better path
		if d, ok := dist[curr.state]; ok && curr.dist > d {
			continue
		}

		// Check if done
		if curr.state.collected == allKeys {
			return curr.dist
		}

		// Try going to each uncollected key
		paths := distances[curr.state.at]
		for nextKey, path := range paths {
			keyBit := keyBits[nextKey]

			// Already collected
			if curr.state.collected&keyBit != 0 {
				continue
			}

			// Don't have required keys to pass doors
			if path.requiredKeys&curr.state.collected != path.requiredKeys {
				continue
			}

			newState := keyState{nextKey, curr.state.collected | keyBit}
			newDist := curr.dist + path.dist

			if d, ok := dist[newState]; !ok || newDist < d {
				dist[newState] = newDist
				heap.Push(pq, day18HeapItem{newState, newDist})
			}
		}
	}

	return 0 // no solution found
}

// day18HeapItem for priority queue
type day18HeapItem struct {
	state struct {
		at        byte
		collected uint32
	}
	dist uint
}

// day18Heap implements heap.Interface
type day18Heap []day18HeapItem

func (h day18Heap) Len() int           { return len(h) }
func (h day18Heap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h day18Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *day18Heap) Push(x any) {
	*h = append(*h, x.(day18HeapItem))
}

func (h *day18Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// modifyMazeForPart2 transforms the maze by replacing the 3x3 area around @ with 4 robots
func modifyMazeForPart2(maze *Maze) {
	// Find @ and replace 3x3 area:
	// ... becomes @#@
	// .@.         ###
	// ...         @#@
	cx, cy := maze.startX, maze.startY

	// Set the pattern
	maze.grid[cy-1][cx-1] = '@'
	maze.grid[cy-1][cx] = '#'
	maze.grid[cy-1][cx+1] = '@'

	maze.grid[cy][cx-1] = '#'
	maze.grid[cy][cx] = '#'
	maze.grid[cy][cx+1] = '#'

	maze.grid[cy+1][cx-1] = '@'
	maze.grid[cy+1][cx] = '#'
	maze.grid[cy+1][cx+1] = '@'

	// Set robot positions (top-left, top-right, bottom-left, bottom-right)
	maze.robots4[0] = image.Point{X: cx - 1, Y: cy - 1}
	maze.robots4[1] = image.Point{X: cx + 1, Y: cy - 1}
	maze.robots4[2] = image.Point{X: cx - 1, Y: cy + 1}
	maze.robots4[3] = image.Point{X: cx + 1, Y: cy + 1}
}

// pathInfo stores distance and required keys to reach a destination
type pathInfo struct {
	dist         uint
	requiredKeys uint32
}

// bfsFrom computes distances from a position to all reachable keys
func bfsFrom(maze Maze, start image.Point, keyBits map[byte]uint32) map[byte]pathInfo {
	result := make(map[byte]pathInfo)
	visited := make(map[image.Point]bool)

	type bfsState struct {
		pos          image.Point
		dist         uint
		requiredKeys uint32
	}

	queue := make([]bfsState, 0, 1000)
	queue = append(queue, bfsState{start, 0, 0})
	visited[start] = true
	dirs := []image.Point{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	head := 0
	for head < len(queue) {
		curr := queue[head]
		head++

		for _, dir := range dirs {
			newPos := image.Point{X: curr.pos.X + dir.X, Y: curr.pos.Y + dir.Y}

			if newPos.X < 0 || newPos.X >= maze.dimX || newPos.Y < 0 || newPos.Y >= maze.dimY {
				continue
			}
			if visited[newPos] {
				continue
			}

			cell := maze.grid[newPos.Y][newPos.X]
			if cell == '#' {
				continue
			}

			visited[newPos] = true
			newRequired := curr.requiredKeys

			// If it's a door, we need the corresponding key
			if cell >= 'A' && cell <= 'Z' {
				newRequired |= keyBits[cell-'A'+'a']
			}

			// If it's a key, record the path to it
			if cell >= 'a' && cell <= 'z' {
				result[cell] = pathInfo{curr.dist + 1, newRequired}
			}

			queue = append(queue, bfsState{newPos, curr.dist + 1, newRequired})
		}
	}

	return result
}

type state4 struct {
	robotKeys [4]byte // which key each robot is at ('@' for start)
	collected uint32
}

func solvePart2(maze Maze) uint {
	modifyMazeForPart2(&maze)

	// Build key list and bits
	keyList := make([]byte, 0, len(maze.keys))
	for key := range maze.keys {
		keyList = append(keyList, key)
	}
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

	// Build position map
	keyPos := make(map[byte]image.Point)
	for y := range maze.dimY {
		for x := range maze.dimX {
			cell := maze.grid[y][x]
			if cell >= 'a' && cell <= 'z' {
				keyPos[cell] = image.Point{X: x, Y: y}
			}
		}
	}

	// Precompute distances from each robot start and from each key
	distances := make(map[byte]map[byte]pathInfo)

	// From each robot starting position
	for i, startPos := range maze.robots4 {
		robotID := byte('@' + i) // '@', 'A', 'B', 'C' for 4 robots
		distances[robotID] = bfsFrom(maze, startPos, keyBits)
	}

	// From each key position
	for key, pos := range keyPos {
		distances[key] = bfsFrom(maze, pos, keyBits)
	}

	// Dijkstra instead of recursive DP
	dist := make(map[state4]uint)
	initialState := state4{
		robotKeys: [4]byte{'@', '@' + 1, '@' + 2, '@' + 3},
		collected: 0,
	}
	dist[initialState] = 0

	pq := &day18Heap4{day18HeapItem4{initialState, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(day18HeapItem4)

		// Skip if we've found a better path
		if d, ok := dist[curr.state]; ok && curr.dist > d {
			continue
		}

		// Check if done
		if curr.state.collected == allKeys {
			return curr.dist
		}

		// Try moving each robot to an uncollected key
		for i := range 4 {
			currentKey := curr.state.robotKeys[i]
			paths := distances[currentKey]

			for nextKey, path := range paths {
				keyBit := keyBits[nextKey]

				// Already collected
				if curr.state.collected&keyBit != 0 {
					continue
				}

				// Don't have required keys
				if path.requiredKeys&curr.state.collected != path.requiredKeys {
					continue
				}

				// Move robot i to nextKey
				newState := curr.state
				newState.robotKeys[i] = nextKey
				newState.collected |= keyBit
				newDist := curr.dist + path.dist

				if d, ok := dist[newState]; !ok || newDist < d {
					dist[newState] = newDist
					heap.Push(pq, day18HeapItem4{newState, newDist})
				}
			}
		}
	}

	return 0
}

// day18HeapItem4 for Part 2 priority queue
type day18HeapItem4 struct {
	state state4
	dist  uint
}

// day18Heap4 for Part 2
type day18Heap4 []day18HeapItem4

func (h day18Heap4) Len() int           { return len(h) }
func (h day18Heap4) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h day18Heap4) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *day18Heap4) Push(x any) {
	*h = append(*h, x.(day18HeapItem4))
}

func (h *day18Heap4) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
