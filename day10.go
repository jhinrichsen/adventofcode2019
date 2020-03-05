package adventofcode2019

// Day10 holds a map of asteroids.
type Day10 struct {
	asteroids []Asteroid
}

// An Asteroid is identified by a dimensionless (x/y) position in a 2D space.
type Asteroid struct {
	x, y int
}

// NewDay10 parses newline separated strings into a Day 10 struct.
func NewDay10(asteroids []byte) Day10 {
	isAsteroid := func (b byte) bool {
		return b == '#'
	}
	isEmpty := func (b byte) bool {
		return b == '.'
	}
	isNewline := func (b byte) bool {
		return b == '\n'
	}

	var d Day10
	y := 0
	x := 0
	for i := range asteroids {
		b := asteroids[i]
		if isEmpty(b) {
			x++
		} else if isNewline(b) {
			y++
		} else if isAsteroid(b) {
			d.asteroids = append(d.asteroids, Asteroid{x, y})
		} else {
			// whitespace, ignore
		}
	}
	return d
}