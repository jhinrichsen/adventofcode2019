package adventofcode2019

const (
	blockTile  = 2
	paddleTile = 3
	ballTile   = 4
)

// Day13 runs the arcade game
func Day13(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	if part1 {
		return uint(day13Part1(ic)), nil
	}
	return uint(day13Part2(ic)), nil
}

func day13Part1(ic *intcode) int {
	blocks := 0
	outputIdx := 0
	var x, y int

	for {
		state := ic.Step()
		switch state {
		case hasOutput:
			val := ic.Output()
			switch outputIdx % 3 {
			case 0:
				x = val
			case 1:
				y = val
			case 2:
				_ = x // suppress unused
				_ = y
				if val == blockTile {
					blocks++
				}
			}
			outputIdx++
		case halted:
			return blocks
		}
	}
}

func day13Part2(ic *intcode) int {
	// Play for free
	ic.SetMem(0, 2)

	outputIdx := 0
	var x, y int
	var ballX, paddleX int
	score := 0

	for {
		state := ic.Step()
		switch state {
		case needsInput:
			// Move paddle towards ball
			joystick := 0
			if paddleX < ballX {
				joystick = 1
			} else if paddleX > ballX {
				joystick = -1
			}
			ic.Input(joystick)
		case hasOutput:
			val := ic.Output()
			switch outputIdx % 3 {
			case 0:
				x = val
			case 1:
				y = val
			case 2:
				if x == -1 && y == 0 {
					score = val
				} else if val == ballTile {
					ballX = x
				} else if val == paddleTile {
					paddleX = x
				}
			}
			outputIdx++
		case halted:
			return score
		}
	}
}
