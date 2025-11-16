package adventofcode2019

import (
	"fmt"
)

const (
	opcodeAdd = 1
	opcodeMul = 2
	opcodeRet = 99
)

// parseIntcode parses comma-separated integers from []byte
func parseIntcode(input []byte) ([]int, error) {
	// Pre-allocate with estimated capacity based on input size
	opcodes := make([]int, 0, len(input)/4)
	num := 0
	hasDigits := false
	negative := false

	for i, b := range input {
		if b >= '0' && b <= '9' {
			num = num*10 + int(b-'0')
			hasDigits = true
		} else if b == '-' {
			negative = true
		} else if b == ',' || b == '\n' {
			if hasDigits {
				if negative {
					num = -num
				}
				opcodes = append(opcodes, num)
				num = 0
				hasDigits = false
				negative = false
			}
		} else {
			return nil, fmt.Errorf("unexpected byte at position %d: %q", i, b)
		}
	}

	// Handle last number
	if hasDigits {
		if negative {
			num = -num
		}
		opcodes = append(opcodes, num)
	}

	return opcodes, nil
}

// runIntcode executes the intcode program
func runIntcode(opcodes []int) {
	pc := 0
	for opcodes[pc] != opcodeRet {
		switch opcodes[pc] {
		case opcodeAdd:
			opcodes[opcodes[pc+3]] = opcodes[opcodes[pc+1]] + opcodes[opcodes[pc+2]]
			pc += 4
		case opcodeMul:
			opcodes[opcodes[pc+3]] = opcodes[opcodes[pc+1]] * opcodes[opcodes[pc+2]]
			pc += 4
		}
	}
}

// Day02 solves the 1202 Program Alarm puzzle
func Day02(input []byte, part1 bool) (uint, error) {
	master, err := parseIntcode(input)
	if err != nil {
		return 0, err
	}

	if part1 {
		opcodes := make([]int, len(master))
		copy(opcodes, master)
		opcodes[1] = 12
		opcodes[2] = 2
		runIntcode(opcodes)
		return uint(opcodes[0]), nil
	}

	// Part 2: Find noun and verb that produce output 19690720
	opcodes := make([]int, len(master))
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(opcodes, master)
			opcodes[1] = noun
			opcodes[2] = verb
			runIntcode(opcodes)
			if opcodes[0] == 19690720 {
				return uint(100*noun + verb), nil
			}
		}
	}

	return 0, fmt.Errorf("no solution found")
}
