package adventofcode2019

type day13 struct {
	x, y int
}

func (a day13) isScore() bool {
	return a.x == -1 && a.y == 0
}

type day13Tile int
type day13Board map[day13]day13Tile

const (
	emptyTile day13Tile = iota
	wallTile
	blockTile
	paddleTile
	ballTile
)

// Blocks returns number of blocks on board.
func (a day13Board) Blocks() int {
	n := 0
	for _, v := range a {
		if v == blockTile {
			n++
		}
	}
	return n
}

// Day13Part1 returns number of blocks for arcade game.
func Day13Part1(cpu IntCodeProcessor, code IntCode) int {
	in, out := channels()
	go cpu(code, in, out)
	i := 0
	var tup day13
	m := make(day13Board)
	for val := range out {
		// parse triples
		triplet := i % 3
		if triplet == 0 {
			tup.x = val
		} else if triplet == 1 {
			tup.y = val
		} else {
			m[tup] = day13Tile(val)
		}
		i++
	}
	return m.Blocks()
}

// Memory address 0 represents the number of quarters that have been
// inserted; set it to 2 to play for free.
func playForFree(code *IntCode) {
	(*code)[0] = 2
}

type joystickDirection int

const (
	left joystickDirection = iota - 1
	neutral
	right
)

func movePaddle(paddle, ball day13) joystickDirection {
	if paddle.x == ball.x {
		return neutral
	}
	if paddle.x < ball.x {
		return right
	}
	return left
}

// Day13Part2 returns number of blocks for arcade game.
func Day13Part2(cpu IntCodeProcessor, code IntCode) int {
	playForFree(&code)

	in, out := channels()
	go cpu(code, in, out)

	// index continous output into triples
	tripleIndex := 0

	// save position of ball and paddle so we can simulate joystick
	var ball, paddle day13
	var tup day13
	score := 0

	for val := range out {
		if tripleIndex == 0 {
			tup.x = val
		} else if tripleIndex == 1 {
			tup.y = val
		} else if tup.isScore() {
			score = val
		} else {
			tile := day13Tile(val)
			if tile == ballTile {
				// save ball position
				ball = tup
				// empirically found that input channel won't
				// block if input is provided for each new ball
				// position
				in <- int(movePaddle(paddle, ball))
			} else if tile == paddleTile {
				// save paddle position
				paddle = tup
			}
		}
		tripleIndex++
		if tripleIndex == 3 {
			tripleIndex = 0
		}
	}
	return score
}
