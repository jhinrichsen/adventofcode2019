package adventofcode2019

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Day25 runs the Cryostasis text adventure and returns the password.
func Day25(prog IntCode, part1 bool) uint {
	if !part1 {
		return 0 // No part 2 for day 25
	}

	// Script to collect all safe items and reach checkpoint
	// This path was determined through exploration
	script := `north
take mutex
north
take semiconductor
south
west
take astrolabe
south
take sand
north
east
south
west
take antenna
north
take klein bottle
south
south
take spool of cat6
west
take tambourine
west`

	items := []string{
		"mutex",
		"semiconductor",
		"astrolabe",
		"sand",
		"antenna",
		"klein bottle",
		"spool of cat6",
		"tambourine",
	}

	// Try all 256 combinations
	for mask := range 256 {
		commands := strings.Split(script, "\n")

		// Drop all items
		for _, item := range items {
			commands = append(commands, "drop "+item)
		}

		// Take selected items
		for i, item := range items {
			if mask&(1<<i) != 0 {
				commands = append(commands, "take "+item)
			}
		}

		// Try to go through checkpoint
		commands = append(commands, "west")

		output := runWithCommands(prog, commands)
		if password := extractPassword(output); password != 0 {
			return password
		}
	}

	return 0
}

// runWithCommands runs the Intcode program with a list of commands.
func runWithCommands(prog IntCode, commands []string) string {
	input := make(chan int, 1000)
	output := make(chan int, 1000)

	go func() {
		Day5(prog.Copy(), input, output)
		close(output)
	}()

	var result strings.Builder
	cmdIndex := 0
	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-timeout:
			return result.String()

		case val, ok := <-output:
			if !ok {
				return result.String()
			}
			result.WriteByte(byte(val))

			// Check if waiting for command
			if strings.HasSuffix(result.String(), "Command?\n") {
				if cmdIndex < len(commands) {
					cmd := commands[cmdIndex]
					cmdIndex++

					for _, ch := range cmd {
						input <- int(ch)
					}
					input <- 10
				} else {
					// No more commands
					return result.String()
				}
			}
		}
	}
}

// extractPassword extracts the password from output.
func extractPassword(output string) uint {
	re := regexp.MustCompile(`typing (\d+) on the keypad`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		var password uint
		fmt.Sscanf(matches[1], "%d", &password)
		return password
	}
	return 0
}
