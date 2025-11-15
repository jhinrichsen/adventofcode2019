package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

// Chemical represents a chemical with its quantity
type Chemical struct {
	Name     string
	Quantity int
}

// Reaction represents a chemical reaction
type Reaction struct {
	Inputs []Chemical
	Output Chemical
}

// ReactionMap maps chemical names to their reactions
type ReactionMap map[string]Reaction

// ParseReactions parses the input lines into a reaction map
func ParseReactions(lines []string) (ReactionMap, error) {
	reactions := make(ReactionMap)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, "=>")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid reaction format: %s", line)
		}

		// Parse output
		output, err := parseChemical(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, err
		}

		// Parse inputs
		inputStrs := strings.Split(parts[0], ",")
		inputs := make([]Chemical, 0, len(inputStrs))
		for _, inputStr := range inputStrs {
			input, err := parseChemical(strings.TrimSpace(inputStr))
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, input)
		}

		reactions[output.Name] = Reaction{
			Inputs: inputs,
			Output: output,
		}
	}

	return reactions, nil
}

// parseChemical parses a string like "10 ORE" into a Chemical
func parseChemical(s string) (Chemical, error) {
	parts := strings.Fields(s)
	if len(parts) != 2 {
		return Chemical{}, fmt.Errorf("invalid chemical format: %s", s)
	}

	quantity, err := strconv.Atoi(parts[0])
	if err != nil {
		return Chemical{}, fmt.Errorf("invalid quantity: %s", parts[0])
	}

	return Chemical{
		Name:     parts[1],
		Quantity: quantity,
	}, nil
}

// CalculateOre calculates the minimum ORE needed to produce the target chemical
func CalculateOre(reactions ReactionMap, target string, targetQty int) int {
	// Track surplus materials
	surplus := make(map[string]int)

	return calculateOreRecursive(reactions, target, targetQty, surplus)
}

// calculateOreRecursive recursively calculates ORE needed
func calculateOreRecursive(reactions ReactionMap, chemical string, needed int, surplus map[string]int) int {
	// Base case: ORE is the raw material
	if chemical == "ORE" {
		return needed
	}

	// Use surplus if available
	if surplus[chemical] > 0 {
		if surplus[chemical] >= needed {
			surplus[chemical] -= needed
			return 0
		}
		needed -= surplus[chemical]
		surplus[chemical] = 0
	}

	// Get the reaction for this chemical
	reaction, ok := reactions[chemical]
	if !ok {
		panic(fmt.Sprintf("no reaction found for %s", chemical))
	}

	// Calculate how many times we need to run this reaction
	outputQty := reaction.Output.Quantity
	times := (needed + outputQty - 1) / outputQty // Ceiling division

	// Track surplus from over-production
	produced := times * outputQty
	if produced > needed {
		surplus[chemical] += produced - needed
	}

	// Calculate ORE needed for all inputs
	totalOre := 0
	for _, input := range reaction.Inputs {
		requiredQty := input.Quantity * times
		totalOre += calculateOreRecursive(reactions, input.Name, requiredQty, surplus)
	}

	return totalOre
}

// Day14Part1 calculates minimum ORE needed to produce 1 FUEL
func Day14Part1(lines []string) (int, error) {
	reactions, err := ParseReactions(lines)
	if err != nil {
		return 0, err
	}

	ore := CalculateOre(reactions, "FUEL", 1)
	return ore, nil
}

// Day14Part2 calculates maximum FUEL that can be produced from 1 trillion ORE
func Day14Part2(lines []string) (int, error) {
	reactions, err := ParseReactions(lines)
	if err != nil {
		return 0, err
	}

	totalOre := 1000000000000 // 1 trillion

	// Binary search for the maximum fuel
	// First get the ORE per fuel as a baseline
	orePerFuel := CalculateOre(reactions, "FUEL", 1)

	// Lower bound estimate
	low := totalOre / orePerFuel
	// Upper bound estimate (generous)
	high := low * 2

	result := 0
	for low <= high {
		mid := (low + high) / 2
		ore := CalculateOre(reactions, "FUEL", mid)

		if ore <= totalOre {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result, nil
}
