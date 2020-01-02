package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Input takes a single integer as input and saves it to the position
	// given by its only parameter.
	Input = 3

	// Output outputs the value of its only parameter.
	Output = 4
)

// IntCode is both a low-level interpreted language used in bootstrapping the
// BCPL compiler, and a virtual language in advent of code.
type IntCode []int

// IntCodeProcessor can run IntCode
type IntCodeProcessor func(IntCode, <-chan int, chan<- int)

// ParameterMode is part of the opcode and controls parameter handling
type ParameterMode int

const (
	// PositionMode reads memory via index
	PositionMode = 0
	// ImmediateMode uses values directly
	ImmediateMode = 1
)

// Day5 supports running IntCode for day 2 (ADD, MUL, RET) and adds input,
// output, parameter mode and immediate mode.
func Day5(ic IntCode, input <-chan int, output chan<- int) {
	load := func(pc int, mode ParameterMode) int {
		if mode == ImmediateMode {
			return ic[pc]
		}
		return ic[ic[pc]]
	}
	store := func(pc int, n int) {
		ic[ic[pc]] = n
	}
	pc := 0
	halt := false
	for !halt {
		instruction := make([]byte, 5)
		DigitsInto(ic[pc], instruction)
		opcode := 10*instruction[3] + instruction[4]
		mode2 := ParameterMode(instruction[1])
		mode1 := ParameterMode(instruction[2])
		switch opcode {
		case OpcodeAdd:
			v := load(pc+1, mode1) + load(pc+2, mode2)
			store(pc+3, v)
			pc += 4
		case OpcodeMul:
			v := load(pc+1, mode1) * load(pc+2, mode2)
			store(pc+3, v)
			pc += 4
		case Input:
			adr := ic[pc+1]
			val := <-input
			ic[adr] = val
			pc += 2
		case Output:
			val := load(pc+1, mode1)
			output <- val
			pc += 2
		case OpcodeRet:
			halt = true
		default:
			panic(fmt.Sprintf("unknown opcode %d at position %d",
				ic[pc], pc))
		}
	}
	close(output)
}

// MustSplit converts IntCode in string representation (a comma separated list
// of token) into IntCode, and panics if conversion fails.
func MustSplit(program string) (ic IntCode) {
	for _, s := range strings.Split(program, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ic = append(ic, n)
	}
	return
}
