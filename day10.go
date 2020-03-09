package adventofcode2019

import "math/cmplx"

// An Asteroid is identified by a dimensionless (x/y) position in a 2D space.
// Using a type _alias_ here to avoid casting
type Asteroid = complex128

// Day10 holds asteroids.
type Day10 struct {
	asteroids []Asteroid
}

// NewDay10 parses newline separated strings into a Day 10 struct.
func NewDay10(asteroids []byte) Day10 {
	isAsteroid := func(b byte) bool {
		return b == '#'
	}
	isEmpty := func(b byte) bool {
		return b == '.'
	}
	isNewline := func(b byte) bool {
		return b == '\n'
	}
	isWhitespace := func(b byte) bool {
		return !(isAsteroid(b) || isEmpty(b))
	}

	// overread any leading whitespace
	start := 0
	for isWhitespace(asteroids[start]) {
		start++
	}
	var d Day10
	y := 0
	x := 0
	for _, b := range asteroids[start:] {
		if isEmpty(b) {
			x++
		} else if isNewline(b) {
			x = 0
			y++
		} else if isAsteroid(b) {
			a := complex(float64(x), float64(y))
			d.asteroids = append(d.asteroids, a)
			x++
		} else {
			// whitespace, ignore
		}
	}
	return d
}

// Part1 returns the asteroid that can see most asteroids, and the number of
// visible asteroids.
func (a Day10) Part1() (Asteroid, int) {
	var best Asteroid
	var maxVisible int
	for i := range a.asteroids {
		// map of angles
		visible := make(map[float64]Asteroid, len(a.asteroids))
		for j := range a.asteroids {
			// skip ourself
			if i == j {
				continue
			}
			rel := a.asteroids[j] - a.asteroids[i]
			r, φ := cmplx.Polar(rel)
			// already an asteroid at same angle?
			if a, ok := visible[φ]; ok {
				// is it closer?
				if r < cmplx.Abs(a) {
					// yes, make us visible, hide other
					visible[φ] = a
				}
			} else {
				// no, just save us
				visible[φ] = rel
			}
		}
		// found a better planet?
		if len(visible) > maxVisible {
			best = a.asteroids[i]
			maxVisible = len(visible)
		}
	}
	return best, maxVisible
}
