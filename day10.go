package adventofcode2019

import (
	"math"
)

// Polar converts cartesian into polar coordinates
// https://de.wikipedia.org/wiki/Polarkoordinaten#Umrechnung_von_kartesischen_Koordinaten_in_Polarkoordinaten
func Polar(x, y int) (float64, float64) {
	r := math.Sqrt(float64(x*x + y*y))
	fx := float64(x)
	fy := float64(y)
	var φ float64

	if x > 0 && y >= 0 {
		φ = math.Atan(fy / fx)
	} else if x > 0 && y < 0 {
		φ = math.Atan(fy/fx) + 2*math.Pi
	} else if x < 0 {
		φ = math.Atan(fy/fx) + math.Pi
	} else if x == 0 && y > 0 {
		φ = math.Pi / 2
	} else /* x == 0 && y < 0 */ {
		φ = 3 * math.Pi / 2
	}
	return r, φ
}

// Day10 holds a cartesian map of asteroids.
type Day10 struct {
	asteroids []Asteroid
}

// An Asteroid is identified by a dimensionless (x/y) position in a 2D space.
type Asteroid struct {
	x, y int
}

// Distance returns the cartesian distance (dx, dy) from a to a2.
func (a Asteroid) Distance(a2 Asteroid) (int, int) {
	return a2.x - a.x, a2.y - a.y
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
			d.asteroids = append(d.asteroids, Asteroid{x, y})
			x++
		} else {
			// whitespace, ignore
		}
	}
	return d
}

// Best returns the asteroid that can see most asteroids, and the number of
// visible asteroids.
func (a Day10) Best() (Asteroid, int) {
	var best Asteroid
	var maxVisible int
	// convert relative distance (dx/dy) into (phi, len)
	for i := range a.asteroids {
		// map of angles and one corresopnding min distance
		visible := make(map[float64]float64, len(a.asteroids))
		for j := range a.asteroids {
			// skip ourself
			if i == j {
				continue
			}
			r, φ := Polar(a.asteroids[i].Distance(a.asteroids[j]))
			// already an asteroid at same angle?
			if l, ok := visible[φ]; ok {
				// are we closer?
				if r < l {
					// yes, make us visible, hide other
					visible[φ] = r
				}
			} else {
				// no, just save us
				visible[φ] = r
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
