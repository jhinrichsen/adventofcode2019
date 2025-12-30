package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

// chemical represents a chemical with its quantity
type chemical struct {
	Name     string
	Quantity uint
}

// reaction represents a chemical reaction
type reaction struct {
	Inputs []chemical
	Output chemical
}

// reactionMap maps chemical names to their reactions
type reactionMap map[string]reaction

// parseReactions parses the input lines into a reaction map
func parseReactions(lines []string) (reactionMap, error) {
	reactions := make(reactionMap)

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
		inputs := make([]chemical, 0, len(inputStrs))
		for _, inputStr := range inputStrs {
			input, err := parseChemical(strings.TrimSpace(inputStr))
			if err != nil {
				return nil, err
			}
			inputs = append(inputs, input)
		}

		reactions[output.Name] = reaction{
			Inputs: inputs,
			Output: output,
		}
	}

	return reactions, nil
}

// parseChemical parses a string like "10 ORE" into a Chemical
// Inline parsing to avoid allocations from strings.Fields
func parseChemical(s string) (chemical, error) {
	// Skip leading whitespace
	i := 0
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	// Parse quantity
	start := i
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	if i == start {
		return chemical{}, fmt.Errorf("invalid chemical format: %s", s)
	}
	quantity, err := strconv.ParseUint(s[start:i], 10, 64)
	if err != nil {
		return chemical{}, err
	}
	// Skip whitespace between quantity and name
	for i < len(s) && (s[i] == ' ' || s[i] == '\t') {
		i++
	}
	// Rest is name (trim trailing whitespace)
	end := len(s)
	for end > i && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return chemical{
		Name:     s[i:end],
		Quantity: uint(quantity),
	}, nil
}

// calculateOre calculates the minimum ORE needed to produce the target chemical
// using an iterative algorithm with explicit dependency tracking
func calculateOre(reactions reactionMap, target string, targetQty uint) uint {
	needs := make(map[string]uint)
	surplus := make(map[string]uint)
	return calculateOreWithMaps(reactions, target, targetQty, needs, surplus)
}

// calculateOreWithMaps allows reusing maps to avoid allocations
func calculateOreWithMaps(reactions reactionMap, target string, targetQty uint, needs, surplus map[string]uint) uint {
	// Clear maps for reuse
	clear(needs)
	clear(surplus)
	needs[target] = targetQty

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
	reactions, err := parseReactions(lines)
	if err != nil {
		return 0
	}

	if part1 {
		// Part 1: minimum ORE needed to produce 1 FUEL
		return calculateOre(reactions, "FUEL", 1)
	}

	// Part 2: maximum FUEL that can be produced from 1 trillion ORE
	totalOre := uint(1000000000000) // 1 trillion

	// Reusable maps for binary search
	needs := make(map[string]uint)
	surplus := make(map[string]uint)

	// First get the ORE per fuel as a baseline
	orePerFuel := calculateOreWithMaps(reactions, "FUEL", 1, needs, surplus)

	// Lower bound estimate
	low := totalOre / orePerFuel
	// Upper bound estimate (generous)
	high := low * 2

	var result uint
	for low <= high {
		mid := (low + high) / 2
		ore := calculateOreWithMaps(reactions, "FUEL", mid, needs, surplus)

		if ore <= totalOre {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return result
}
