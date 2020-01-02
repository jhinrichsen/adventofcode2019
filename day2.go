package adventofcode2019

import (
	"strconv"
	"strings"
)

const (
	// OpcodeAdd adds parameter
	OpcodeAdd = 1
	// OpcodeMul multiplies parameter
	OpcodeMul = 2
	// OpcodeRet halts the program
	OpcodeRet = 99
)

// Len calculates them maximum index into an array, i.e. the len of an array
// that can hold opcodes.
func Len(opcodes []int) int {
	// just return max no matter if index or opcode (99)
	var max int
	for _, n := range opcodes {
		if n > max {
			max = n
		}
	}
	return max
}

// Split converts a comma separated list into an array
func Split(s string) ([]int, error) {
	var opcodes []int
	for _, opcode := range strings.Split(s, ",") {
		n, err := strconv.Atoi(opcode)
		if err != nil {
			return opcodes, err
		}
		opcodes = append(opcodes, n)
	}
	return opcodes, nil
}

// ToString returns a comma separated list of opcodes
func ToString(opcodes []int) string {
	var ss []string
	for _, opcode := range opcodes {
		ss = append(ss, strconv.Itoa(opcode))
	}
	return strings.Join(ss, ",")
}

// Run executes opcodes
func Run(opcodes []int) ([]int, error) {
	pc := 0
	for opcodes[pc] != OpcodeRet {
		switch opcodes[pc] {
		case OpcodeAdd:
			opcodes[opcodes[pc+3]] = opcodes[opcodes[pc+1]] + opcodes[opcodes[pc+2]]
			pc += 4
		case OpcodeMul:
			opcodes[opcodes[pc+3]] = opcodes[opcodes[pc+1]] * opcodes[opcodes[pc+2]]
			pc += 4
		}
	}
	return opcodes, nil
}

// Runs executes a comma separated list of opcodes
func Runs(s string) (string, error) {
	opcodes, err := Split(s)
	if err != nil {
		return "", nil
	}
	opcodes, err = Run(opcodes)
	if err != nil {
		return "", err
	}
	return ToString(opcodes), nil
}
