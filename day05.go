package adventofcode2019

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	opcodeAdd = 1
	opcodeMul = 2

	// opInput takes a single integer as input and saves it to the position
	// given by its only parameter.
	opInput = 3

	// opOutput outputs the value of its only parameter.
	opOutput = 4

	// jumpIfTrue sets the instruction pointer to the value from the second
	// parameter if the first parameter is non-zero.
	jumpIfTrue = 5

	// jumpIfFalse sets the instruction pointer to the value from the second
	// parameter if the first parameter is zero.
	jumpIfFalse = 6

	// lessThan stores 1 in the position given by the third parameter if
	// the first parameter is less than the second parameter, otherwise
	// it stores 0.
	lessThan = 7

	// equals stores 1 in the position given by the third parameter if
	// the first parameter is equal to the second parameter, otherwise
	// it stores 0.
	equals = 8

	// adjustRelBase changes the relative base
	adjustRelBase = 9

	opcodeRet = 99
)

// intCode is both a low-level interpreted language used in bootstrapping the
// BCPL compiler, and a virtual language in advent of code.
type intCode []int

// Copy provides a master copy because most programs are self modifying
func (master intCode) Copy() intCode {
	prog := make(intCode, len(master))
	copy(prog, master)
	return prog
}

// intCodeProcessor can run IntCode
type intCodeProcessor func(intCode, <-chan int, chan<- int)

// parameterMode is part of the opcode and controls parameter handling
type parameterMode int

const (
	// positionMode reads memory via index
	positionMode parameterMode = iota

	// immediateMode uses values directly
	immediateMode

	// relativeMode addresses via relative base
	relativeMode
)

// boolean is a C style boolean: false -> 0, true -> 1.
func boolean(b bool) int {
	if b {
		return 1
	}
	return 0
}

// isFalse returns false for 0, true otherwise.
func isFalse(boolean int) bool {
	return boolean == 0
}

// isTrue returns false for 0, true otherwise.
func isTrue(boolean int) bool {
	return boolean != 0
}

// day5 supports running IntCode for day 2 (ADD, MUL, RET) and adds input,
// output, parameter mode and immediate mode.
func day5(program intCode, input <-chan int, output chan<- int) {
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
			// Negative address - skip reallocation
			return
		}
		if idx >= len(program) {
			bigger := make(intCode, idx+1)
			copy(bigger, program)
			program = bigger
		}
	}

	load := func(idx int, mode parameterMode) int {
		// good old 6502
		lda := func(idx int) int {
			realloc(idx)
			return program[idx]
		}
		switch mode {
		case immediateMode:
			return lda(idx)
		case positionMode:
			return lda(lda(idx))
		case relativeMode:
			return lda(relBase + lda(idx))
		}
		return -1
	}

	store := func(idx int, val int, mode parameterMode) {
		switch mode {
		case immediateMode:
			realloc(idx)
			program[idx] = val
		case positionMode:
			adr := program[idx]
			realloc(adr)
			program[adr] = val
		case relativeMode:
			adr := relBase + program[idx]
			realloc(adr)
			program[adr] = val
		}
	}

	halt := false
	for !halt {
		opcode, mode1, mode2, mode3 := instruction(program[ip])
		switch opcode {
		case opcodeAdd:
			val := load(ip+1, mode1) + load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case opcodeMul:
			val := load(ip+1, mode1) * load(ip+2, mode2)
			store(ip+3, val, mode3)
			ip += 4
		case opInput:
			val := <-input
			store(ip+1, val, mode1)
			ip += 2
		case opOutput:
			val := load(ip+1, mode1)
			output <- val
			ip += 2
		case jumpIfTrue:
			p := load(ip+1, mode1)
			if isTrue(p) {
				ip = load(ip+2, mode2)
				// No IP alignment for jumps
				continue
			}
			ip += 3
		case jumpIfFalse:
			p := load(ip+1, mode1)
			if isFalse(p) {
				ip = load(ip+2, mode2)
				// No IP alignment for jumps
				continue
			}
			ip += 3
		case lessThan:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			val := boolean(p1 < p2)
			store(ip+3, val, mode3)
			ip += 4
		case equals:
			p1 := load(ip+1, mode1)
			p2 := load(ip+2, mode2)
			val := boolean(p1 == p2)
			store(ip+3, val, mode3)
			ip += 4
		case adjustRelBase:
			p1 := load(ip+1, mode1)
			relBase += p1
			ip += 2
		case opcodeRet:
			halt = true
		default:
			// Unknown opcode - halt execution gracefully
			halt = true
		}
	}
	close(output)
}

func instruction(instr int) (byte, parameterMode, parameterMode, parameterMode) {
	opcode := byte(instr % 100)
	instr /= 100
	mode1 := parameterMode(instr % 10)
	instr /= 10
	mode2 := parameterMode(instr % 10)
	mode3 := parameterMode(instr / 10)
	return opcode, mode1, mode2, mode3
}

// mustSplit converts IntCode in string representation (a comma separated list
// of token) into IntCode. Invalid values are converted to 0.
func mustSplit(program string) (ic intCode) {
	for _, s := range strings.Split(program, ",") {
		n, err := strconv.Atoi(s)
		if err != nil {
			n = 0
		}
		ic = append(ic, n)
	}
	return
}

// toString converts opcodes back to comma-separated string
func toString(opcodes []int) string {
	result := ""
	for i, opcode := range opcodes {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", opcode)
	}
	return result
}

func channels() (chan int, chan int) {
	return make(chan int, 2), make(chan int, 2)
}

// Day05 runs the diagnostic program and returns the diagnostic code
func Day05(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	var systemID int
	if part1 {
		systemID = 1 // Air conditioner unit
	} else {
		systemID = 5 // Thermal radiator controller
	}

	outputs, err := ic.Run(systemID)
	if err != nil {
		return 0, err
	}

	if len(outputs) == 0 {
		return 0, nil
	}
	return uint(outputs[len(outputs)-1]), nil
}
