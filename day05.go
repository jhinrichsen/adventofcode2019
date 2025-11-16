package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	OpcodeAdd = 1
	OpcodeMul = 2

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

	// AdjustRelBase changes the relative base
	AdjustRelBase = 9

	OpcodeRet = 99
)

// IntCode is both a low-level interpreted language used in bootstrapping the
// BCPL compiler, and a virtual language in advent of code.
type IntCode []int

// Copy provides a master copy because most programs are self modifying
func (master IntCode) Copy() IntCode {
	prog := make(IntCode, len(master))
	copy(prog, master)
	return prog
}

// IntCodeProcessor can run IntCode
type IntCodeProcessor func(IntCode, <-chan int, chan<- int)

// ParameterMode is part of the opcode and controls parameter handling
type ParameterMode int

const (
	// PositionMode reads memory via index
	PositionMode ParameterMode = iota

	// ImmediateMode uses values directly
	ImmediateMode

	// RelativeMode addresses via relative base
	RelativeMode
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
	return boolean == 0
}

// True returns false for 0, true otherwise.
func True(boolean int) bool {
	return boolean != 0
}

// Day5 supports running IntCode for day 2 (ADD, MUL, RET) and adds input,
// output, parameter mode and immediate mode.
func Day5(program IntCode, input <-chan int, output chan<- int) {
	// instruction pointer, aka program counter
	ip := 0

	// relative position mode
	relBase := 0

	realloc := func(idx int) {
		// day 9 extension:
		// The computer's available memory should be much larger than
		// the initial program. Memory beyond the initial program starts
		// with the value 0 and can be read or written like any other
		// memory. (It is invalid to try to access memory at a negative
		// address, though.)
		if idx < 0 {
			s := fmt.Sprintf("trying to access address %d which "+
				"is not allowed (ip=%d, relBase=%d)",
				idx, ip, relBase)
			panic(s)
		}
		if idx >= len(program) {
			bigger := make(IntCode, idx+1)
			copy(bigger, program)
			program = bigger
		}
	}

	load := func(idx int, mode ParameterMode) int {
		// good old 6502
		lda := func(idx int) int {
			realloc(idx)
			return program[idx]
		}
		switch mode {
		case ImmediateMode:
			return lda(idx)
		case PositionMode:
			return lda(lda(idx))
		case RelativeMode:
			return lda(relBase + lda(idx))
		}
		return -1
	}

	store := func(idx int, val int, mode ParameterMode) {
		switch mode {
		case ImmediateMode:
			realloc(idx)
			program[idx] = val
		case PositionMode:
			adr := program[idx]
			realloc(adr)
			program[adr] = val
		case RelativeMode:
			adr := relBase + program[idx]
			realloc(adr)
			program[adr] = val
		}
	}

	halt := false
	for !halt {
		opcode, mode1, mode2, mode3 := instruction(program[ip])
		switch opcode {
		case OpcodeAdd:
			val := load(ip+1, mode1) + load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case OpcodeMul:
			val := load(ip+1, mode1) * load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case Input:
			val := <-input
			store(ip+1, val, mode1)
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
			val := Boolean(p1 < p2)
			store(ip+3, val, mode3)
			ip += 4
		case Equals:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			val := Boolean(p1 == p2)
			store(ip+3, val, mode3)
			ip += 4
		case AdjustRelBase:
			p1 := load(ip+1, mode1)
			relBase += p1
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

func instruction(instr int) (byte, ParameterMode, ParameterMode, ParameterMode) {
	var buf [5]byte
	DigitsInto(instr, buf[:])
	opcode := 10*buf[3] + buf[4]
	mode3 := ParameterMode(buf[0])
	mode2 := ParameterMode(buf[1])
	mode1 := ParameterMode(buf[2])
	return opcode, mode1, mode2, mode3
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

func channels() (chan int, chan int) {
	return make(chan int, 2), make(chan int, 2)
}
