package adventofcode2019

type Day13 struct {
	x, y int
}

type Day13Tile int
type Day13Board map[Day13]Day13Tile

const (
	Empty Day13Tile = iota
	Wall
	Block
	HorizontalPaddle
	Ball
)

// Blocks returns number of blocks on board.
func (a Day13Board) Blocks() int {
	n := 0
	for _, v := range a {
		if v == Block {
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
	var tup Day13
	m := make(Day13Board)
	for val := range out {
		// parse triples
		triplet := i % 3
		if triplet == 0 {
			tup.x = val
		} else if triplet == 1 {
			tup.y = val
		} else {
			m[tup] = Day13Tile(val)
		}
		i++
	}
	return m.Blocks()
}
