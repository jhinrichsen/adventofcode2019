package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

// Chemical represents a chemical with its quantity
type Chemical struct {
	Name     string
	Quantity uint
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

	quantity, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return Chemical{}, fmt.Errorf("invalid quantity: %s", parts[0])
	}

	return Chemical{
		Name:     parts[1],
		Quantity: uint(quantity),
	}, nil
}

// calculateOre calculates the minimum ORE needed to produce the target chemical
// using an iterative algorithm with explicit dependency tracking
func calculateOre(reactions ReactionMap, target string, targetQty uint) uint {
	// Track what we need to produce
	needs := make(map[string]uint)
	needs[target] = targetQty

	// Track surplus materials from over-production
	surplus := make(map[string]uint)

	// Iteratively resolve dependencies until only ORE remains
	for {
		// Find a chemical to process (not ORE)
		var chemical string
		for chem := range needs {
			if chem != "ORE" {
				chemical = chem
				break
			}
		}

		// If no non-ORE chemical found, we're done
		if chemical == "" {
			break
		}

		needed := needs[chemical]
		delete(needs, chemical)

		// Use surplus if available
		if surplus[chemical] > 0 {
			if surplus[chemical] >= needed {
				surplus[chemical] -= needed
				continue
			}
			needed -= surplus[chemical]
			surplus[chemical] = 0
		}

		// Get the reaction for this chemical
		reaction, ok := reactions[chemical]
		if !ok {
			// No reaction found - skip this chemical
			continue
		}

		// Calculate how many times we need to run this reaction
		outputQty := reaction.Output.Quantity
		times := (needed + outputQty - 1) / outputQty // Ceiling division

		// Track surplus from over-production
		produced := times * outputQty
		if produced > needed {
			surplus[chemical] += produced - needed
		}

		// Add input requirements to needs
		for _, input := range reaction.Inputs {
			requiredQty := input.Quantity * times
			needs[input.Name] += requiredQty
		}
	}

	return needs["ORE"]
}

// Day14 calculates either minimum ORE for 1 FUEL (part1) or
// maximum FUEL from 1 trillion ORE (part2)
func Day14(lines []string, part1 bool) uint {
	reactions, err := ParseReactions(lines)
	if err != nil {
		return 0
	}

	if part1 {
		// Part 1: minimum ORE needed to produce 1 FUEL
		return calculateOre(reactions, "FUEL", 1)
	}

	// Part 2: maximum FUEL that can be produced from 1 trillion ORE
	totalOre := uint(1000000000000) // 1 trillion

	// Binary search for the maximum fuel
	// First get the ORE per fuel as a baseline
	orePerFuel := calculateOre(reactions, "FUEL", 1)

	// Lower bound estimate
	low := totalOre / orePerFuel
	// Upper bound estimate (generous)
	high := low * 2

	var result uint
	for low <= high {
		mid := (low + high) / 2
		ore := calculateOre(reactions, "FUEL", mid)

		if ore <= totalOre {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}
