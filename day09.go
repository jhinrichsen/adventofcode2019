package adventofcode2019

// Day09 runs the BOOST program and returns the output
func Day09(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	var inputVal int
	if part1 {
		inputVal = 1
	} else {
		inputVal = 2
	}

	outputs, err := ic.Run(inputVal)
	if err != nil {
		return 0, err
	}

	if len(outputs) > 0 {
		return uint(outputs[len(outputs)-1]), nil
	}
	return 0, nil
}
