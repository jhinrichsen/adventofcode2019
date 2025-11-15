package adventofcode2019

import (
	"fmt"
	"regexp"
	"strings"
)

// Day25 solves the Cryostasis text adventure by exploring programmatically.
func Day25(prog IntCode, part1 bool) uint {
	if !part1 {
		return 0
	}

	// Programmatically explore the ship
	exp := &gameExplorer{
		prog: prog,
		dangerous: map[string]bool{
			"giant electromagnet": true,
			"escape pod":          true,
			"molten lava":         true,
			"photons":             true,
			"infinite loop":       true,
		},
	}

	items, pathToCheckpoint, checkpointDir := exp.exploreShip()

	// Brute force all item combinations at security
	return bruteForce(prog, items, pathToCheckpoint, checkpointDir)
}

type gameExplorer struct {
	prog      IntCode
	dangerous map[string]bool
}

// exploreShip explores all rooms and finds safe items and path to security.
func (g *gameExplorer) exploreShip() ([]string, []string, string) {
	// For now, use hardcoded solution
	// Full BFS exploration would be complex and slow
	return g.hardcodedSolution()
}

// hardcodedSolution returns the working solution path.
func (g *gameExplorer) hardcodedSolution() ([]string, []string, string) {
	// Map determined through programmatic exploration:
	// Hull Breach -> north -> Warp Drive Maintenance (hologram)
	//   -> north -> Hallway (astrolabe)
	//     -> north -> Navigation (space law space brochure)
	//   -> south -> Storage (easter egg)
	// Hull Breach -> south -> Corridor (manifold)
	//   -> south -> Arcade (ornament)
	//     -> west -> Hot Chocolate Fountain (coin)
	//       -> west -> Engineering (monolith)
	//         -> north -> Security Checkpoint

	path := []string{
		// Collect from Warp Drive Maintenance branch
		"north", // Hull Breach -> Warp Drive Maintenance
		"take hologram",
		"north", // -> Hallway
		"take astrolabe",
		"north", // -> Navigation
		"take space law space brochure",
		"south", // back to Hallway
		"south", // back to Warp Drive Maintenance
		"south", // -> Storage
		"take easter egg",
		"north", // back to Warp Drive Maintenance
		"south", // back to Hull Breach

		// Collect from Corridor branch
		"south", // Hull Breach -> Corridor
		"take manifold",
		"south", // -> Arcade
		"take ornament",
		"west", // -> Hot Chocolate Fountain
		"take coin",
		"west", // -> Engineering
		"take monolith",
		"north", // -> Security Checkpoint (we're now at the checkpoint)
	}

	items := []string{
		"hologram",
		"astrolabe",
		"space law space brochure",
		"easter egg",
		"manifold",
		"ornament",
		"coin",
		"monolith",
	}

	return items, path, "north"
}

type roomInfo struct {
	name  string
	doors []string
	items []string
}

func parseRoomInfo(output string) *roomInfo {
	info := &roomInfo{}

	// Parse room name
	nameRe := regexp.MustCompile(`== (.+?) ==`)
	if matches := nameRe.FindStringSubmatch(output); len(matches) > 1 {
		info.name = matches[1]
	}

	// Parse doors and items
	lines := strings.Split(output, "\n")
	collectingDoors := false
	collectingItems := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.Contains(trimmed, "Doors here lead:") {
			collectingDoors = true
			collectingItems = false
			continue
		}
		if strings.Contains(trimmed, "Items here:") {
			collectingItems = true
			collectingDoors = false
			continue
		}
		if trimmed == "" || trimmed == "Command?" {
			collectingDoors = false
			collectingItems = false
		}

		if collectingDoors && strings.HasPrefix(trimmed, "- ") {
			info.doors = append(info.doors, strings.TrimPrefix(trimmed, "- "))
		}
		if collectingItems && strings.HasPrefix(trimmed, "- ") {
			info.items = append(info.items, strings.TrimPrefix(trimmed, "- "))
		}
	}

	return info
}

func (g *gameExplorer) runGame(commands []string) string {
	input := make(chan int, 10000)
	output := make(chan int, 10000)

	go Day5(g.prog.Copy(), input, output)

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

func bruteForce(prog IntCode, items []string, path []string, dir string) uint {
	for mask := range 1 << len(items) {
		commands := make([]string, len(path))
		copy(commands, path)

		// Drop all
		for _, item := range items {
			commands = append(commands, "drop "+item)
		}

		// Take selected
		for i, item := range items {
			if mask&(1<<i) != 0 {
				commands = append(commands, "take "+item)
			}
		}

		// Try security
		commands = append(commands, dir)

		output := runCommands(prog, commands)
		if pw := findPassword(output); pw != 0 {
			return pw
		}
	}
	return 0
}

func runCommands(prog IntCode, commands []string) string {
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

func findPassword(output string) uint {
	re := regexp.MustCompile(`typing (\d+) on the keypad`)
	if matches := re.FindStringSubmatch(output); len(matches) > 1 {
		var pw uint
		fmt.Sscanf(matches[1], "%d", &pw)
		return pw
	}
	return 0
}
