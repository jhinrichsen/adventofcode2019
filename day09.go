package adventofcode2019

// NewDay09 parses IntCode program from input
var NewDay09 = func(input []byte) IntCode {
	return MustSplit(string(input))
}

// Day09 runs the BOOST program and returns the output
func Day09(input []byte, part1 bool) uint {
	program := NewDay09(input)
	in, out := channels()

	go Day5(program, in, out)

	if part1 {
		in <- 1
	} else {
		in <- 2
	}

	// Return the output value
	result := <-out
	return uint(result)
}
