package adventofcode2019

import (
	"fmt"
	"regexp"
	"strings"
)

// Day25 solves the Cryostasis text adventure.
func Day25(lines []string, part1 bool) uint {
	if !part1 {
		return 0
	}

	prog := MustSplit(strings.TrimSpace(lines[0]))

	// Build room graph using DFS
	graph := buildRoomGraph(prog)

	// Find all safe items and security checkpoint
	var items []string
	var securityRoom string
	var securityDir string

	dangerous := map[string]bool{
		"giant electromagnet": true,
		"escape pod":          true,
		"molten lava":         true,
		"photons":             true,
		"infinite loop":       true,
	}

	for roomName, room := range graph {
		if strings.Contains(roomName, "Security") {
			securityRoom = roomName
			for dir := range room.exits {
				securityDir = dir
				break
			}
		}
		for _, item := range room.items {
			if !dangerous[item] {
				items = append(items, item)
			}
		}
	}

	// Build path to collect all items and reach security
	path := buildCollectionPath(graph, items, securityRoom)

	// Try all item combinations
	return tryItemCombos(prog, items, path, securityDir)
}

type roomInfo struct {
	name  string
	exits map[string]string // direction -> room name
	items []string
}

// buildRoomGraph explores the ship and builds a complete room graph.
func buildRoomGraph(prog IntCode) map[string]*roomInfo {
	graph := make(map[string]*roomInfo)
	visited := make(map[string]bool)

	var dfs func(path []string)
	dfs = func(path []string) {
		output := executeCommands(prog, path)
		parsedRoom := parseRoom(output)

		// Add room to graph if not already there
		if _, exists := graph[parsedRoom.name]; !exists {
			graph[parsedRoom.name] = parsedRoom
		}

		// Get the room from the graph (in case it already existed)
		room := graph[parsedRoom.name]

		// Skip further exploration if we've already explored this room's exits
		if visited[room.name] {
			return
		}

		// Don't explore past security checkpoint
		if strings.Contains(room.name, "Security") {
			visited[room.name] = true
			return
		}

		visited[room.name] = true

		// Explore all exits
		for dir := range room.exits {
			nextPath := append(append([]string{}, path...), dir)
			nextOutput := executeCommands(prog, nextPath)
			nextRoom := parseRoom(nextOutput)

			// Record where this direction leads
			room.exits[dir] = nextRoom.name

			// Only continue DFS if we haven't visited this room yet
			if !visited[nextRoom.name] {
				dfs(nextPath)
			}
		}
	}

	dfs([]string{})
	return graph
}

// parseRoom extracts room data from game output.
func parseRoom(output string) *roomInfo {
	room := &roomInfo{exits: make(map[string]string)}

	// Extract room name (find the LAST occurrence since output may contain multiple rooms)
	var roomName string
	nameRe := regexp.MustCompile(`== (.+?) ==`)
	allMatches := nameRe.FindAllStringSubmatch(output, -1)
	if len(allMatches) > 0 {
		roomName = allMatches[len(allMatches)-1][1]
	}
	room.exits["_name"] = roomName // Hack to store name

	// Parse lines
	lines := strings.Split(output, "\n")
	inDoors, inItems := false, false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.Contains(trimmed, "Doors here lead:") {
			inDoors, inItems = true, false
			continue
		}
		if strings.Contains(trimmed, "Items here:") {
			inItems, inDoors = true, false
			continue
		}
		if trimmed == "" || strings.Contains(trimmed, "Command?") {
			inDoors, inItems = false, false
		}

		if inDoors && strings.HasPrefix(trimmed, "- ") {
			dir := strings.TrimPrefix(trimmed, "- ")
			room.exits[dir] = "" // Will be filled by DFS
		}
		if inItems && strings.HasPrefix(trimmed, "- ") {
			item := strings.TrimPrefix(trimmed, "- ")
			room.items = append(room.items, item)
		}
	}

	room.name = roomName
	delete(room.exits, "_name")
	return room
}

// buildCollectionPath creates path to collect all items and reach security.
func buildCollectionPath(graph map[string]*roomInfo, items []string, security string) []string {
	var path []string

	// Find each item's room
	itemRooms := make(map[string]string)
	for roomName, room := range graph {
		for _, item := range room.items {
			itemRooms[item] = roomName
		}
	}

	// Track current location
	currentRoom := "Hull Breach"

	// Collect all items
	for _, item := range items {
		roomName := itemRooms[item]
		roomPath := findRoomPath(graph, currentRoom, roomName)
		path = append(path, roomPath...)
		path = append(path, "take "+item)
		currentRoom = roomName
	}

	// Go to security
	secPath := findRoomPath(graph, currentRoom, security)
	path = append(path, secPath...)

	return path
}

// findRoomPath finds shortest path between two rooms using BFS.
func findRoomPath(graph map[string]*roomInfo, start, target string) []string {
	if start == target {
		return []string{}
	}

	type state struct {
		room string
		path []string
	}

	queue := []state{{room: start, path: []string{}}}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.room] {
			continue
		}
		visited[current.room] = true

		if current.room == target {
			return current.path
		}

		room := graph[current.room]
		if room == nil {
			continue
		}

		for dir, nextRoom := range room.exits {
			if nextRoom == "" {
				continue
			}
			newPath := append(append([]string{}, current.path...), dir)
			queue = append(queue, state{room: nextRoom, path: newPath})
		}
	}

	return []string{}
}

// tryItemCombos tries all combinations of items.
func tryItemCombos(prog IntCode, items []string, path []string, dir string) uint {
	for mask := range 1 << len(items) {
		cmds := make([]string, len(path))
		copy(cmds, path)

		// Drop all
		for _, item := range items {
			cmds = append(cmds, "drop "+item)
		}

		// Take selected
		for i, item := range items {
			if mask&(1<<i) != 0 {
				cmds = append(cmds, "take "+item)
			}
		}

		// Try security
		cmds = append(cmds, dir)

		output := executeCommands(prog, cmds)
		if pw := getPassword(output); pw != 0 {
			return pw
		}
	}
	return 0
}

func executeCommands(prog IntCode, commands []string) string {
	input := make(chan int, 10000)
	output := make(chan int, 10000)

	go Day5(prog.Copy(), input, output)

	var result strings.Builder
	cmdIdx := 0

	for {
		select {
		case val, ok := <-output:
			if !ok {
				return result.String()
			}
			result.WriteByte(byte(val))

			if strings.HasSuffix(result.String(), "Command?\n") {
				if cmdIdx < len(commands) {
					cmd := commands[cmdIdx]
					cmdIdx++
					for _, ch := range cmd {
						input <- int(ch)
					}
					input <- 10
				} else {
					close(input)
					return result.String()
				}
			}
		}
	}
}

func getPassword(output string) uint {
	re := regexp.MustCompile(`typing (\d+) on the keypad`)
	if matches := re.FindStringSubmatch(output); len(matches) > 1 {
		var pw uint
		fmt.Sscanf(matches[1], "%d", &pw)
		return pw
	}
	return 0
}
