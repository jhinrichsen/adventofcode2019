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
func Day5(program IntCode, input <-chan int, output chan<- int) {
	load := func(ip int, mode ParameterMode) int {
		if mode == ImmediateMode {
			return program[ip]
		}
		return program[program[ip]]
	}
	store := func(ip int, n int) {
		program[program[ip]] = n
	}
	// instruction pointer, aka program counter
	ip := 0
	halt := false
	for !halt {
		instruction := make([]byte, 5)
		DigitsInto(program[ip], instruction)
		opcode := 10*instruction[3] + instruction[4]
		mode2 := ParameterMode(instruction[1])
		mode1 := ParameterMode(instruction[2])
		switch opcode {
		case OpcodeAdd:
			v := load(ip+1, mode1) + load(ip+2, mode2)
			store(ip+3, v)
			ip += 4
		case OpcodeMul:
			v := load(ip+1, mode1) * load(ip+2, mode2)
			store(ip+3, v)
			ip += 4
		case Input:
			adr := program[ip+1]
			val := <-input
			program[adr] = val
			ip += 2
		case Output:
			val := load(ip+1, mode1)
			output <- val
			ip += 2
		case OpcodeRet:
			halt = true
		default:
			panic(fmt.Sprintf("unknown opcode %d at position %d",
				program[ip], ip))
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
