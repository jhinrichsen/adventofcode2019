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

	// JumpIfTrue sets the instruction pointer to the value from the second
	// parameter if the first parameter is non-zero.
	JumpIfTrue = 5

	// JumpIfFalse sets the instruction pointer to the value from the second
	// parameter if the first parameter is zero.
	JumpIfFalse = 6

	// LessThan stores 1 in the position given by the third parameter if
	// the first parameter is less than the second parameter, otherwise
	// it stores 0.
	LessThan = 7

	// Equals stores 1 in the position given by the third parameter if
	// the first parameter is equal to the second parameter, otherwise
	// it stores 0.
	Equals = 8
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

// Boolean is a C style Boolean: false -> 0, true -> 1.
func Boolean(b bool) int {
	if b {
		return 1
	}
	return 0
}

// False returns false for 0, true otherwise.
func False(boolean int) bool {
	if boolean == 0 {
		return true
	}
	return false
}

// True returns false for 0, true otherwise.
func True(boolean int) bool {
	return !False(boolean)
}

// Day5 supports running IntCode for day 2 (ADD, MUL, RET) and adds input,
// output, parameter mode and immediate mode.
func Day5(program IntCode, input <-chan int, output chan<- int) {
	load := func(ip int, mode ParameterMode) int {
		// assume immediate mode
		val := program[ip]
		if mode == PositionMode {
			// nope, one more indirection
			val = program[val]
		}
		return val
	}
	// instruction pointer, aka program counter
	ip := 0
	halt := false
	for !halt {
		opcode, mode1, mode2 := instruction(program[ip])
		switch opcode {
		case OpcodeAdd:
			val := load(ip+1, mode1) + load(ip+2, mode2)
			program[program[ip+3]] = val
			ip += 4
		case OpcodeMul:
			val := load(ip+1, mode1) * load(ip+2, mode2)
			program[program[ip+3]] = val
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
		case JumpIfTrue:
			p := load(ip+1, mode1)
			if True(p) {
				ip = load(ip+2, mode2)
				// No IP alignment for jumps
				continue
			}
			ip += 3
		case JumpIfFalse:
			p := load(ip+1, mode1)
			if False(p) {
				ip = load(ip+2, mode2)
				// No IP alignment for jumps
				continue
			}
			ip += 3
		case LessThan:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			p3 := load(ip+3, ImmediateMode)
			val := Boolean(p1 < p2)
			program[p3] = val
			ip += 4
		case Equals:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			p3 := load(ip+3, ImmediateMode)
			val := Boolean(p1 == p2)
			program[p3] = val
			ip += 4

		case OpcodeRet:
			halt = true
		default:
			panic(fmt.Sprintf("unknown opcode %d at position %d",
				program[ip], ip))
		}
	}
	close(output)
}

func instruction(instr int) (byte, ParameterMode, ParameterMode) {
	var buf [5]byte
	DigitsInto(instr, buf[:])
	opcode := 10*buf[3] + buf[4]
	mode2 := ParameterMode(buf[1])
	mode1 := ParameterMode(buf[2])
	return opcode, mode1, mode2
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
