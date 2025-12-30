package adventofcode2019

import "errors"

// Day02 solves the 1202 Program Alarm puzzle
func Day02(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	if part1 {
		ic.SetMem(1, 12)
		ic.SetMem(2, 2)
		if _, err := ic.Run(); err != nil {
			return 0, err
		}
		return uint(ic.Mem(0)), nil
	}

	// Part 2: Find noun and verb that produce output 19690720
	for noun := range 100 {
		for verb := range 100 {
			ic.Reset()
			ic.SetMem(1, noun)
			ic.SetMem(2, verb)
			if _, err := ic.Run(); err != nil {
				continue
			}
			if ic.Mem(0) == 19690720 {
				return uint(100*noun + verb), nil
			}
		}
	}

	return 0, errors.New("no solution found")
}
