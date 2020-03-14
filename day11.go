package adventofcode2019

import "fmt"

// Day11Part1 returns the number of painted tiles.
func Day11Part1(prog IntCode) int {
	const (
		// no types because we want to pass it to the IntCode Computer
		colorBlack = 0
		colorWhite = 1

		// no types because we want to pass it to the IntCode Computer
		turnLeft  = 0
		turnRight = 1
	)
	in, out := channels()

	// boot emergency hull painting robot
	go Day5(prog, in, out)

	colors := make(map[complex128]bool)
	position := 0 + 0i
	left := 0 + 1i
	right := 0 - 1i
	// The robot starts facing up
	direction := left

	color := func() int {
		// Default is black
		if white := colors[position]; white {
			return colorWhite
		}
		return colorBlack
	}
	setColor := func(c int) {
		b := c == colorWhite
		colors[position] = b
	}
	translateTurn := func(i int) complex128 {
		if i == turnLeft {
			return left
		} else if i == turnRight {
			return right
		} else {
			panic(fmt.Errorf("unsupported turn value %d", i))
		}
	}

	in <- color()
	for col := range out {
		setColor(col)
		turnValue := <-out
		// make robot turn
		turn := translateTurn(turnValue)
		direction *= turn
		// one step
		position += direction

		in <- color()
	}
	return len(colors)
}
