package adventofcode2019

import "errors"

// state represents the current state of the Intcode machine after a Step.
type state int

const (
	running    state = iota // Still executing, call Step again
	needsInput              // Waiting for input, call Input then Step
	hasOutput               // Output available, call Output then Step
	halted                  // Program finished (opcode 99)
)

// intcode is a synchronous intcode virtual machine.
// Use Step for fine-grained control or Run for batch execution.
type intcode struct {
	original []int // pristine copy for Reset
	mem      []int // working memory
	ip       int   // instruction pointer
	relBase  int   // relative base for mode 2
	output   int   // last output value
	state    state // current state
	dirty    bool  // true if program memory was modified
}

// newIntcode parses the input and returns a new Intcode machine.
func newIntcode(input []byte) (*intcode, error) {
	// Count commas to pre-allocate
	count := 1
	for _, b := range input {
		if b == ',' {
			count++
		}
	}

	original := make([]int, 0, count)
	num := 0
	negative := false
	hasDigits := false

	for _, b := range input {
		switch {
		case b >= '0' && b <= '9':
			num = num*10 + int(b-'0')
			hasDigits = true
		case b == '-':
			negative = true
		case b == ',' || b == '\n':
			if hasDigits {
				if negative {
					num = -num
				}
				original = append(original, num)
				num = 0
				negative = false
				hasDigits = false
			}
		}
	}
	if hasDigits {
		if negative {
			num = -num
		}
		original = append(original, num)
	}

	ic := &intcode{
		original: original,
		mem:      make([]int, len(original)),
	}
	copy(ic.mem, original)
	return ic, nil
}

// Reset restores the machine to its initial state.
func (ic *intcode) Reset() {
	// Only copy memory if it was modified
	if ic.dirty {
		// Reuse existing memory if capacity is sufficient
		if cap(ic.mem) >= len(ic.original) {
			ic.mem = ic.mem[:len(ic.original)]
		} else {
			ic.mem = make([]int, len(ic.original))
		}
		copy(ic.mem, ic.original)
		ic.dirty = false
	}
	ic.ip = 0
	ic.relBase = 0
	ic.output = 0
	ic.state = running
}

// Clone returns a fresh Intcode machine sharing the same parsed program.
func (ic *intcode) Clone() *intcode {
	clone := &intcode{
		original: ic.original, // share original (never modified)
		mem:      make([]int, len(ic.original)),
	}
	copy(clone.mem, ic.original)
	return clone
}

// Mem returns the value at memory address addr.
func (ic *intcode) Mem(addr int) int {
	if addr >= len(ic.mem) {
		return 0
	}
	return ic.mem[addr]
}

// SetMem sets the value at memory address addr.
func (ic *intcode) SetMem(addr, val int) {
	ic.grow(addr)
	ic.markDirty(addr)
	ic.mem[addr] = val
}

// Output returns the last output value.
func (ic *intcode) Output() int {
	return ic.output
}

// Input provides a value for the next input instruction.
func (ic *intcode) Input(val int) {
	if ic.state != needsInput {
		return
	}
	opcode := ic.mem[ic.ip] % 100
	if opcode != 3 {
		return
	}
	addr := ic.writeAddr(1)
	ic.grow(addr)
	ic.markDirty(addr)
	ic.mem[addr] = val
	ic.ip += 2
	ic.state = running
}

// Step executes one instruction and returns the new state.
func (ic *intcode) Step() state {
	if ic.state == halted || ic.state == needsInput {
		return ic.state
	}

	opcode := ic.mem[ic.ip] % 100

	switch opcode {
	case 1: // add
		a := ic.read(1)
		b := ic.read(2)
		addr := ic.writeAddr(3)
		ic.grow(addr)
		ic.markDirty(addr)
		ic.mem[addr] = a + b
		ic.ip += 4

	case 2: // multiply
		a := ic.read(1)
		b := ic.read(2)
		addr := ic.writeAddr(3)
		ic.grow(addr)
		ic.markDirty(addr)
		ic.mem[addr] = a * b
		ic.ip += 4

	case 3: // input
		ic.state = needsInput
		return ic.state

	case 4: // output
		ic.output = ic.read(1)
		ic.ip += 2
		ic.state = hasOutput
		return ic.state

	case 5: // jump-if-true
		if ic.read(1) != 0 {
			ic.ip = ic.read(2)
		} else {
			ic.ip += 3
		}

	case 6: // jump-if-false
		if ic.read(1) == 0 {
			ic.ip = ic.read(2)
		} else {
			ic.ip += 3
		}

	case 7: // less than
		addr := ic.writeAddr(3)
		ic.grow(addr)
		ic.markDirty(addr)
		if ic.read(1) < ic.read(2) {
			ic.mem[addr] = 1
		} else {
			ic.mem[addr] = 0
		}
		ic.ip += 4

	case 8: // equals
		addr := ic.writeAddr(3)
		ic.grow(addr)
		ic.markDirty(addr)
		if ic.read(1) == ic.read(2) {
			ic.mem[addr] = 1
		} else {
			ic.mem[addr] = 0
		}
		ic.ip += 4

	case 9: // adjust relative base
		ic.relBase += ic.read(1)
		ic.ip += 2

	case 99: // halt
		ic.state = halted
		return ic.state
	}

	ic.state = running
	return ic.state
}

// errNeedsInput is returned when Run exhausts inputs before the program halts.
var errNeedsInput = errors.New("program needs input but none provided")

// Run executes the program with the given inputs and returns all outputs.
// Returns ErrNeedsInput if the program needs more inputs than provided.
func (ic *intcode) Run(inputs ...int) ([]int, error) {
	// Fast path for programs with no I/O (like Day 2)
	if len(inputs) == 0 {
		return ic.runNoIO()
	}

	inputIdx := 0
	var outputs []int

	for {
		state := ic.Step()
		switch state {
		case halted:
			return outputs, nil
		case needsInput:
			if inputIdx < len(inputs) {
				ic.Input(inputs[inputIdx])
				inputIdx++
			} else {
				return outputs, errNeedsInput
			}
		case hasOutput:
			outputs = append(outputs, ic.output)
			ic.state = running
		}
	}
}

// runNoIO is an optimized path for programs without input/output.
func (ic *intcode) runNoIO() ([]int, error) {
	mem := ic.mem
	ip := 0

	for {
		op := mem[ip]
		opcode := op % 100

		// Fast path for mode 0 (position mode) - most common case
		if op < 100 {
			switch opcode {
			case 1: // add
				addr := mem[ip+3]
				ic.markDirty(addr)
				mem[addr] = mem[mem[ip+1]] + mem[mem[ip+2]]
				ip += 4
			case 2: // multiply
				addr := mem[ip+3]
				ic.markDirty(addr)
				mem[addr] = mem[mem[ip+1]] * mem[mem[ip+2]]
				ip += 4
			case 99:
				ic.ip = ip
				ic.state = halted
				return nil, nil
			default:
				ic.ip = ip
				return ic.runWithStep(nil)
			}
			continue
		}

		// Slow path with mode handling
		switch opcode {
		case 1: // add
			a := ic.readAt(ip, 1)
			b := ic.readAt(ip, 2)
			addr := ic.writeAddrAt(ip, 3)
			ic.grow(addr)
			ic.markDirty(addr)
			mem = ic.mem
			mem[addr] = a + b
			ip += 4
		case 2: // multiply
			a := ic.readAt(ip, 1)
			b := ic.readAt(ip, 2)
			addr := ic.writeAddrAt(ip, 3)
			ic.grow(addr)
			ic.markDirty(addr)
			mem = ic.mem
			mem[addr] = a * b
			ip += 4
		case 99:
			ic.ip = ip
			ic.state = halted
			return nil, nil
		default:
			ic.ip = ip
			return ic.runWithStep(nil)
		}
	}
}

// runWithStep continues execution using Step() for complex programs.
func (ic *intcode) runWithStep(inputs []int) ([]int, error) {
	inputIdx := 0
	var outputs []int

	for {
		state := ic.Step()
		switch state {
		case halted:
			return outputs, nil
		case needsInput:
			if inputIdx < len(inputs) {
				ic.Input(inputs[inputIdx])
				inputIdx++
			} else {
				return outputs, errNeedsInput
			}
		case hasOutput:
			outputs = append(outputs, ic.output)
			ic.state = running
		}
	}
}

// readAt reads parameter n at instruction pointer ip.
func (ic *intcode) readAt(ip, n int) int {
	mode := (ic.mem[ip] / pow10(n+1)) % 10
	param := ic.mem[ip+n]
	switch mode {
	case 0:
		if param >= len(ic.mem) {
			return 0
		}
		return ic.mem[param]
	case 1:
		return param
	case 2:
		addr := ic.relBase + param
		if addr >= len(ic.mem) {
			return 0
		}
		return ic.mem[addr]
	}
	return 0
}

// writeAddrAt returns write address for parameter n at instruction pointer ip.
func (ic *intcode) writeAddrAt(ip, n int) int {
	mode := (ic.mem[ip] / pow10(n+1)) % 10
	param := ic.mem[ip+n]
	switch mode {
	case 2:
		return ic.relBase + param
	default:
		return param
	}
}

// read returns the value of parameter n based on its mode.
func (ic *intcode) read(n int) int {
	mode := (ic.mem[ic.ip] / pow10(n+1)) % 10
	param := ic.mem[ic.ip+n]

	switch mode {
	case 0: // position
		if param >= len(ic.mem) {
			return 0
		}
		return ic.mem[param]
	case 1: // immediate
		return param
	case 2: // relative
		addr := ic.relBase + param
		if addr >= len(ic.mem) {
			return 0
		}
		return ic.mem[addr]
	}
	return 0
}

// writeAddr returns the address where parameter n should write.
func (ic *intcode) writeAddr(n int) int {
	mode := (ic.mem[ic.ip] / pow10(n+1)) % 10
	param := ic.mem[ic.ip+n]

	switch mode {
	case 0: // position
		return param
	case 2: // relative
		return ic.relBase + param
	}
	return param
}

// grow expands memory if needed.
func (ic *intcode) grow(addr int) {
	if addr >= len(ic.mem) {
		newMem := make([]int, addr+1)
		copy(newMem, ic.mem)
		ic.mem = newMem
	}
}

// markDirty sets the dirty flag if writing to original program space
func (ic *intcode) markDirty(addr int) {
	if addr < len(ic.original) {
		ic.dirty = true
	}
}

// pow10table for fast mode extraction.
var pow10table = [5]int{1, 10, 100, 1000, 10000}

// pow10 returns 10^n.
func pow10(n int) int {
	return pow10table[n]
}
