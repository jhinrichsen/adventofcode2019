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

	// Build room graph using checkpoint-based exploration
	graph := buildRoomGraphFast(prog)

	// Find all safe items and security checkpoint
	var items []string
	var securityRoom string

	dangerous := map[string]bool{
		"giant electromagnet": true,
		"escape pod":          true,
		"molten lava":         true,
		"photons":             true,
		"infinite loop":       true,
	}

	var securityDirs []string
	for roomName, room := range graph {
		if strings.Contains(roomName, "Security") {
			securityRoom = roomName
			// Collect ALL security exits to try (fixes non-deterministic map iteration)
			for dir := range room.exits {
				securityDirs = append(securityDirs, dir)
			}
		}
		for _, item := range room.items {
			if !dangerous[item] {
				items = append(items, item)
			}
		}
	}

	if len(items) == 0 || securityRoom == "" {
		return 0
	}

	// Build path to collect all items and reach security
	path := buildCollectionPath(graph, items, securityRoom)

	// Try all item combinations with all possible security directions
	return tryItemCombos(prog, items, path, securityDirs)
}

// vmSnapshot represents a saved VM state
type vmSnapshot struct {
	mem     []int
	ip      int
	relBase int
}

func (s *vmSnapshot) copy() *vmSnapshot {
	newMem := make([]int, len(s.mem))
	copy(newMem, s.mem)
	return &vmSnapshot{
		mem:     newMem,
		ip:      s.ip,
		relBase: s.relBase,
	}
}

// buildRoomGraphFast explores using checkpoint-based VM copies
func buildRoomGraphFast(prog IntCode) map[string]*roomInfo {
	graph := make(map[string]*roomInfo)
	visited := make(map[string]bool)

	// Warmup: Run VM to first "Command?" prompt and get initial output
	warmSnapshot, initialOutput := warmupVM(prog)

	// oppositeDirection returns the reverse direction
	oppositeDirection := func(dir string) string {
		switch dir {
		case "north":
			return "south"
		case "south":
			return "north"
		case "east":
			return "west"
		case "west":
			return "east"
		}
		return ""
	}

	// DFS with checkpoints
	var explore func(checkpoint *vmSnapshot, roomName string)
	explore = func(checkpoint *vmSnapshot, roomName string) {
		if visited[roomName] {
			return
		}

		// Don't explore past security checkpoint
		if strings.Contains(roomName, "Security") {
			visited[roomName] = true
			return
		}

		visited[roomName] = true

		// Room should already be in graph from parent call
		room := graph[roomName]

		// Explore each exit
		for dir := range room.exits {
			// Copy checkpoint and send command
			vmCopy := checkpoint.copy()
			output := vmCopy.sendCommand(dir)
			nextRoom := parseRoom(output)

			// Check if VM died (no "Command?" means death)
			if !strings.Contains(output, "Command?") {
				continue // Skip dangerous path
			}

			// Add to graph if new
			if _, exists := graph[nextRoom.name]; !exists {
				graph[nextRoom.name] = nextRoom
			}

			// Record forward connection
			room.exits[dir] = nextRoom.name

			// Record reverse connection for pathfinding
			reverseDir := oppositeDirection(dir)
			nextGraphRoom := graph[nextRoom.name]
			if nextGraphRoom.exits[reverseDir] == "" {
				nextGraphRoom.exits[reverseDir] = room.name
			}

			// Recursively explore
			explore(vmCopy, nextRoom.name)
		}
	}

	// Parse initial room from warmup output
	startRoom := parseRoom(initialOutput)
	if startRoom.name == "" {
		return graph
	}
	graph[startRoom.name] = startRoom

	// Start exploration
	initCopy := warmSnapshot.copy()
	explore(initCopy, startRoom.name)

	return graph
}

// warmupVM runs program to first "Command?" prompt and returns snapshot + initial output
func warmupVM(prog IntCode) (*vmSnapshot, string) {
	// Allocate memory: program size + working space
	// TODO: Profile actual max address usage
	memSize := len(prog) * 2
	if memSize < 10000 {
		memSize = 10000
	}
	mem := make([]int, memSize)
	copy(mem, prog)

	vm := &vmSnapshot{
		mem:     mem,
		ip:      0,
		relBase: 0,
	}

	// Run until we hit first input, capturing output
	output := vm.runUntilInput()

	return vm, intsToString(output)
}

// runUntilInput runs VM until it needs input, returns output
func (vm *vmSnapshot) runUntilInput() []int {
	var output []int

	for {
		opcode, mode1, mode2, mode3 := instruction(vm.mem[vm.ip])

		switch opcode {
		case OpcodeAdd:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			vm.mem[addr] = p1 + p2
			vm.ip += 4

		case OpcodeMul:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			vm.mem[addr] = p1 * p2
			vm.ip += 4

		case Input:
			return output // Wait for input

		case Output:
			val := vm.loadParam(vm.ip+1, mode1)
			output = append(output, val)
			vm.ip += 2

		case JumpIfTrue:
			p1 := vm.loadParam(vm.ip+1, mode1)
			if p1 != 0 {
				vm.ip = vm.loadParam(vm.ip+2, mode2)
			} else {
				vm.ip += 3
			}

		case JumpIfFalse:
			p1 := vm.loadParam(vm.ip+1, mode1)
			if p1 == 0 {
				vm.ip = vm.loadParam(vm.ip+2, mode2)
			} else {
				vm.ip += 3
			}

		case LessThan:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			if p1 < p2 {
				vm.mem[addr] = 1
			} else {
				vm.mem[addr] = 0
			}
			vm.ip += 4

		case Equals:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			if p1 == p2 {
				vm.mem[addr] = 1
			} else {
				vm.mem[addr] = 0
			}
			vm.ip += 4

		case AdjustRelBase:
			vm.relBase += vm.loadParam(vm.ip+1, mode1)
			vm.ip += 2

		case OpcodeRet:
			return output

		default:
			panic(fmt.Sprintf("unknown opcode %d", vm.mem[vm.ip]))
		}
	}
}

// sendCommand sends a command and returns output
func (vm *vmSnapshot) sendCommand(cmd string) string {
	// Convert command to input
	input := []int{}
	for _, ch := range cmd {
		input = append(input, int(ch))
	}
	input = append(input, 10) // newline

	inputIdx := 0
	output := []int{}

	for {
		opcode, mode1, mode2, mode3 := instruction(vm.mem[vm.ip])

		switch opcode {
		case OpcodeAdd:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			vm.mem[addr] = p1 + p2
			vm.ip += 4

		case OpcodeMul:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			vm.mem[addr] = p1 * p2
			vm.ip += 4

		case Input:
			if inputIdx >= len(input) {
				// Input exhausted - return output so far
				return intsToString(output)
			}
			addr := vm.storeAddr(vm.ip+1, mode1)
			vm.mem[addr] = input[inputIdx]
			inputIdx++
			vm.ip += 2

		case Output:
			val := vm.loadParam(vm.ip+1, mode1)
			output = append(output, val)
			vm.ip += 2

		case JumpIfTrue:
			p1 := vm.loadParam(vm.ip+1, mode1)
			if p1 != 0 {
				vm.ip = vm.loadParam(vm.ip+2, mode2)
			} else {
				vm.ip += 3
			}

		case JumpIfFalse:
			p1 := vm.loadParam(vm.ip+1, mode1)
			if p1 == 0 {
				vm.ip = vm.loadParam(vm.ip+2, mode2)
			} else {
				vm.ip += 3
			}

		case LessThan:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			if p1 < p2 {
				vm.mem[addr] = 1
			} else {
				vm.mem[addr] = 0
			}
			vm.ip += 4

		case Equals:
			p1 := vm.loadParam(vm.ip+1, mode1)
			p2 := vm.loadParam(vm.ip+2, mode2)
			addr := vm.storeAddr(vm.ip+3, mode3)
			if p1 == p2 {
				vm.mem[addr] = 1
			} else {
				vm.mem[addr] = 0
			}
			vm.ip += 4

		case AdjustRelBase:
			vm.relBase += vm.loadParam(vm.ip+1, mode1)
			vm.ip += 2

		case OpcodeRet:
			return intsToString(output)

		default:
			panic(fmt.Sprintf("unknown opcode %d", vm.mem[vm.ip]))
		}
	}
}

func (vm *vmSnapshot) loadParam(addr int, mode ParameterMode) int {
	val := vm.mem[addr]
	switch mode {
	case ImmediateMode:
		return val
	case PositionMode:
		return vm.mem[val]
	case RelativeMode:
		return vm.mem[vm.relBase+val]
	}
	return 0
}

func (vm *vmSnapshot) storeAddr(addr int, mode ParameterMode) int {
	val := vm.mem[addr]
	switch mode {
	case PositionMode:
		return val
	case RelativeMode:
		return vm.relBase + val
	}
	return addr
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

	var dfs func(path []string, roomName string)
	dfs = func(path []string, roomName string) {
		// Skip if already visited
		if visited[roomName] {
			return
		}

		// Don't explore past security checkpoint
		if strings.Contains(roomName, "Security") {
			visited[roomName] = true
			return
		}

		visited[roomName] = true

		// Run Intcode once for this path to get room details
		output := executeCommands(prog, path)
		parsedRoom := parseRoom(output)

		// Add room to graph if not already there
		if _, exists := graph[parsedRoom.name]; !exists {
			graph[parsedRoom.name] = parsedRoom
		}

		// Explore all exits
		room := graph[parsedRoom.name]
		for dir := range room.exits {
			nextPath := append(append([]string{}, path...), dir)
			// Peek at next room to get its name for the graph
			nextOutput := executeCommands(prog, nextPath)
			nextRoom := parseRoom(nextOutput)

			// Record where this direction leads
			room.exits[dir] = nextRoom.name

			// Recurse with the room name to avoid re-execution
			dfs(nextPath, nextRoom.name)
		}
	}

	// Start exploration from Hull Breach
	output := executeCommands(prog, []string{})
	startRoom := parseRoom(output)
	dfs([]string{}, startRoom.name)
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

// tryItemCombos tries all combinations of items with all security directions.
func tryItemCombos(prog IntCode, items []string, path []string, dirs []string) uint {
	// Build input to reach security checkpoint with all items collected
	baseInput := buildInput(path)

	// Run VM once to security checkpoint
	mem := make([]int, 10000)
	copy(mem, prog)
	ip, relBase := runToCheckpoint(mem, baseInput)

	// Save checkpoint state (snapshot)
	checkpointMem := make([]int, len(mem))
	copy(checkpointMem, mem)

	// Try each security direction
	for _, dir := range dirs {
		// Try each item combination using snapshots
		for mask := range 1 << len(items) {
			// Restore checkpoint
			copy(mem, checkpointMem)

			// Build input for this combination
			var combInput []int

			// Drop all items
			for _, item := range items {
				combInput = appendCommand(combInput, "drop "+item)
			}

			// Take selected items
			for i, item := range items {
				if mask&(1<<i) != 0 {
					combInput = appendCommand(combInput, "take "+item)
				}
			}

			// Try security direction
			combInput = appendCommand(combInput, dir)

			// Run from checkpoint with this combination
			output := runFromCheckpoint(mem, ip, relBase, combInput)
			outputStr := intsToString(output)

			if pw := getPassword(outputStr); pw != 0 {
				return pw
			}
		}
	}
	return 0
}

// runToCheckpoint runs VM until input is exhausted, returns final ip and relBase
func runToCheckpoint(mem []int, input []int) (int, int) {
	ip := 0
	relBase := 0
	inputIdx := 0
	output := []int{}

	for {
		opcode, mode1, mode2, mode3 := instruction(mem[ip])

		switch opcode {
		case OpcodeAdd:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			mem[addr] = p1 + p2
			ip += 4

		case OpcodeMul:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			mem[addr] = p1 * p2
			ip += 4

		case Input:
			if inputIdx >= len(input) {
				// Input exhausted - we're at checkpoint
				return ip, relBase
			}
			addr := storeAddr(mem, ip+1, mode1, relBase)
			mem[addr] = input[inputIdx]
			inputIdx++
			ip += 2

		case Output:
			val := loadParam(mem, ip+1, mode1, relBase)
			output = append(output, val)
			ip += 2

		case JumpIfTrue:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			if p1 != 0 {
				ip = loadParam(mem, ip+2, mode2, relBase)
			} else {
				ip += 3
			}

		case JumpIfFalse:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			if p1 == 0 {
				ip = loadParam(mem, ip+2, mode2, relBase)
			} else {
				ip += 3
			}

		case LessThan:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			if p1 < p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case Equals:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			if p1 == p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case AdjustRelBase:
			relBase += loadParam(mem, ip+1, mode1, relBase)
			ip += 2

		case OpcodeRet:
			return ip, relBase

		default:
			panic(fmt.Sprintf("unknown opcode %d", mem[ip]))
		}
	}
}

// runFromCheckpoint continues VM from checkpoint with new input
func runFromCheckpoint(mem []int, startIP, startRelBase int, input []int) []int {
	ip := startIP
	relBase := startRelBase
	inputIdx := 0
	output := []int{}

	for {
		opcode, mode1, mode2, mode3 := instruction(mem[ip])

		switch opcode {
		case OpcodeAdd:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			mem[addr] = p1 + p2
			ip += 4

		case OpcodeMul:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			mem[addr] = p1 * p2
			ip += 4

		case Input:
			if inputIdx >= len(input) {
				return output
			}
			addr := storeAddr(mem, ip+1, mode1, relBase)
			mem[addr] = input[inputIdx]
			inputIdx++
			ip += 2

		case Output:
			val := loadParam(mem, ip+1, mode1, relBase)
			output = append(output, val)
			ip += 2

		case JumpIfTrue:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			if p1 != 0 {
				ip = loadParam(mem, ip+2, mode2, relBase)
			} else {
				ip += 3
			}

		case JumpIfFalse:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			if p1 == 0 {
				ip = loadParam(mem, ip+2, mode2, relBase)
			} else {
				ip += 3
			}

		case LessThan:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			if p1 < p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case Equals:
			p1 := loadParam(mem, ip+1, mode1, relBase)
			p2 := loadParam(mem, ip+2, mode2, relBase)
			addr := storeAddr(mem, ip+3, mode3, relBase)
			if p1 == p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case AdjustRelBase:
			relBase += loadParam(mem, ip+1, mode1, relBase)
			ip += 2

		case OpcodeRet:
			return output

		default:
			panic(fmt.Sprintf("unknown opcode %d", mem[ip]))
		}
	}
}

func loadParam(mem []int, addr int, mode ParameterMode, relBase int) int {
	val := mem[addr]
	switch mode {
	case ImmediateMode:
		return val
	case PositionMode:
		return mem[val]
	case RelativeMode:
		return mem[relBase+val]
	}
	return 0
}

func storeAddr(mem []int, addr int, mode ParameterMode, relBase int) int {
	val := mem[addr]
	switch mode {
	case PositionMode:
		return val
	case RelativeMode:
		return relBase + val
	}
	return addr
}

func buildInput(commands []string) []int {
	var input []int
	for _, cmd := range commands {
		input = appendCommand(input, cmd)
	}
	return input
}

func appendCommand(input []int, cmd string) []int {
	for _, ch := range cmd {
		input = append(input, int(ch))
	}
	return append(input, 10) // newline
}

func intsToString(output []int) string {
	result := make([]byte, len(output))
	for i, val := range output {
		result[i] = byte(val)
	}
	return string(result)
}

func executeCommands(prog IntCode, commands []string) string {
	// Build all input upfront
	var inputBytes []int
	for _, cmd := range commands {
		for _, ch := range cmd {
			inputBytes = append(inputBytes, int(ch))
		}
		inputBytes = append(inputBytes, 10) // newline
	}

	// Run synchronously
	output := runIntcodeSync(prog.Copy(), inputBytes)

	// Convert output to string
	result := make([]byte, len(output))
	for i, val := range output {
		result[i] = byte(val)
	}
	return string(result)
}

// runIntcodeSync runs IntCode synchronously without channels
func runIntcodeSync(program IntCode, input []int) []int {
	// Pre-allocate memory
	mem := make([]int, 10000)
	copy(mem, program)

	ip := 0
	relBase := 0
	inputIdx := 0
	output := make([]int, 0, 100000)

	for {
		opcode, mode1, mode2, mode3 := instruction(mem[ip])

		switch opcode {
		case OpcodeAdd:
			var p1, p2, addr int
			// Load p1
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			// Load p2
			switch mode2 {
			case ImmediateMode:
				p2 = mem[ip+2]
			case PositionMode:
				p2 = mem[mem[ip+2]]
			case RelativeMode:
				p2 = mem[relBase+mem[ip+2]]
			}
			// Store addr
			addr = mem[ip+3]
			if mode3 == RelativeMode {
				addr = relBase + addr
			}
			mem[addr] = p1 + p2
			ip += 4

		case OpcodeMul:
			var p1, p2, addr int
			// Load p1
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			// Load p2
			switch mode2 {
			case ImmediateMode:
				p2 = mem[ip+2]
			case PositionMode:
				p2 = mem[mem[ip+2]]
			case RelativeMode:
				p2 = mem[relBase+mem[ip+2]]
			}
			// Store addr
			addr = mem[ip+3]
			if mode3 == RelativeMode {
				addr = relBase + addr
			}
			mem[addr] = p1 * p2
			ip += 4

		case Input:
			if inputIdx >= len(input) {
				return output
			}
			addr := mem[ip+1]
			if mode1 == RelativeMode {
				addr = relBase + addr
			}
			mem[addr] = input[inputIdx]
			inputIdx++
			ip += 2

		case Output:
			var val int
			switch mode1 {
			case ImmediateMode:
				val = mem[ip+1]
			case PositionMode:
				val = mem[mem[ip+1]]
			case RelativeMode:
				val = mem[relBase+mem[ip+1]]
			}
			output = append(output, val)
			ip += 2

		case JumpIfTrue:
			var p1 int
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			if p1 != 0 {
				switch mode2 {
				case ImmediateMode:
					ip = mem[ip+2]
				case PositionMode:
					ip = mem[mem[ip+2]]
				case RelativeMode:
					ip = mem[relBase+mem[ip+2]]
				}
			} else {
				ip += 3
			}

		case JumpIfFalse:
			var p1 int
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			if p1 == 0 {
				switch mode2 {
				case ImmediateMode:
					ip = mem[ip+2]
				case PositionMode:
					ip = mem[mem[ip+2]]
				case RelativeMode:
					ip = mem[relBase+mem[ip+2]]
				}
			} else {
				ip += 3
			}

		case LessThan:
			var p1, p2, addr int
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			switch mode2 {
			case ImmediateMode:
				p2 = mem[ip+2]
			case PositionMode:
				p2 = mem[mem[ip+2]]
			case RelativeMode:
				p2 = mem[relBase+mem[ip+2]]
			}
			addr = mem[ip+3]
			if mode3 == RelativeMode {
				addr = relBase + addr
			}
			if p1 < p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case Equals:
			var p1, p2, addr int
			switch mode1 {
			case ImmediateMode:
				p1 = mem[ip+1]
			case PositionMode:
				p1 = mem[mem[ip+1]]
			case RelativeMode:
				p1 = mem[relBase+mem[ip+1]]
			}
			switch mode2 {
			case ImmediateMode:
				p2 = mem[ip+2]
			case PositionMode:
				p2 = mem[mem[ip+2]]
			case RelativeMode:
				p2 = mem[relBase+mem[ip+2]]
			}
			addr = mem[ip+3]
			if mode3 == RelativeMode {
				addr = relBase + addr
			}
			if p1 == p2 {
				mem[addr] = 1
			} else {
				mem[addr] = 0
			}
			ip += 4

		case AdjustRelBase:
			var val int
			switch mode1 {
			case ImmediateMode:
				val = mem[ip+1]
			case PositionMode:
				val = mem[mem[ip+1]]
			case RelativeMode:
				val = mem[relBase+mem[ip+1]]
			}
			relBase += val
			ip += 2

		case OpcodeRet:
			return output

		default:
			panic(fmt.Sprintf("unknown opcode %d at position %d", mem[ip], ip))
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
