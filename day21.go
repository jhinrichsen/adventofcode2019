package adventofcode2019

import "strings"

// Day21 solves the "Springdroid Adventure" puzzle.
// Part 1 uses WALK mode, Part 2 uses RUN mode.
func Day21(input []byte, part1 bool) (uint, error) {
	ic, err := NewIntcode(input)
	if err != nil {
		return 0, err
	}

	var springscript string
	if part1 {
		// Part 1: WALK mode - can see 4 tiles ahead (A,B,C,D)
		// Jump if there's a hole in next 3 tiles AND ground at D to land on
		// Logic: (!A OR !B OR !C) AND D
		springscript = strings.Join([]string{
			"NOT A J", // J = !A
			"NOT B T", // T = !B
			"OR T J",  // J = !A OR !B
			"NOT C T", // T = !C
			"OR T J",  // J = !A OR !B OR !C
			"AND D J", // J = (!A OR !B OR !C) AND D
			"WALK",
		}, "\n") + "\n"
	} else {
		// Part 2: RUN mode - can see 9 tiles ahead (A,B,C,D,E,F,G,H,I)
		// Jump if there's a hole in next 3 tiles AND ground at D AND can continue from D
		// After landing at D, we need either E (to walk) or H (to jump again)
		// Logic: (!A OR !B OR !C) AND D AND (E OR H)
		springscript = strings.Join([]string{
			"NOT A J", // J = !A
			"NOT B T", // T = !B
			"OR T J",  // J = !A OR !B
			"NOT C T", // T = !C
			"OR T J",  // J = !A OR !B OR !C
			"AND D J", // J = (!A OR !B OR !C) AND D
			"NOT H T", // T = !H
			"NOT T T", // T = H
			"OR E T",  // T = H OR E
			"AND T J", // J = (!A OR !B OR !C) AND D AND (H OR E)
			"RUN",
		}, "\n") + "\n"
	}

	return executeSpringdroid(ic, springscript), nil
}

func executeSpringdroid(ic *Intcode, springscript string) uint {
	scriptIdx := 0
	var lastOutput int

	for {
		state := ic.Step()
		switch state {
		case NeedsInput:
			if scriptIdx < len(springscript) {
				ic.Input(int(springscript[scriptIdx]))
				scriptIdx++
			}
		case HasOutput:
			lastOutput = ic.Output()
		case Halted:
			return uint(lastOutput)
		}
	}
}
