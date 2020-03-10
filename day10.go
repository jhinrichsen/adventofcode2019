package adventofcode2019

import (
	"log"
	"math"
	"math/cmplx"
	"sort"
)

// An Asteroid is identified by a dimensionless (x/y) position in a 2D space.
// Using a type _alias_ here to avoid casting
type Asteroid = complex128

// Day10 holds asteroids.
// Not concurrent access safe.
type Day10 struct {
	asteroids []Asteroid
}

// vaporized into black outer space
var vaporized = cmplx.Inf()

func IsVaporized(a Asteroid) bool {
	return a == vaporized
}

func Vaporize(a *Asteroid) {
	*a = vaporized
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
func (d Day10) Part1() (Asteroid, int) {
	var best Asteroid
	var maxVisible int
	for i := range d.asteroids {
		// map of angles
		visible := make(map[float64]Asteroid, len(d.asteroids))
		for j := range d.asteroids {
			// skip ourself
			if i == j {
				continue
			}
			rel := d.asteroids[j] - d.asteroids[i]
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
			best = d.asteroids[i]
			maxVisible = len(visible)
		}
	}
	return best, maxVisible
}

func hasRemaining(as []Asteroid) bool {
	for _, a := range as {
		if !IsVaporized(a) {
			return true
		}
	}
	return false
}

func (d Day10) vaporize(base Asteroid) []Asteroid {
	var order []Asteroid

	idx := d.findFirst(base)
	inc := func() {
		idx = (idx + 1) % len(d.asteroids)
	}
	remaining := len(d.asteroids)
	for remaining > 0 {
		a := d.asteroids[idx]
		phase := cmplx.Phase(a - base)
		order = append(order, a)
		Vaporize(&d.asteroids[idx])
		remaining--
		log.Printf("vaporized idx=%d, %+v, remaining=%d\n", idx, a, remaining)
		for j, a := range d.asteroids {
			if !IsVaporized(a) {
				log.Printf("remaining #%d: %v\n", j, a)
			}
		}
		inc()
		for {
			if IsVaporized(d.asteroids[idx]) {
				inc()
				continue
			}
			// skip same phase, those cannot be vaporized
			p := cmplx.Phase(d.asteroids[idx] - base)
			if p == phase {
				inc()
				continue
			}
			break
		}
	}
	return order
}

// Part2 determines the 200th asteroid that gets vaporized.
func (d Day10) Part2(base Asteroid) int {
	as := d.vaporize(base)
	a := as[199]
	// multiply its X coordinate by 100 and then add its Y coordinate
	return int(100.0*real(a) + imag(a))
}

// Sort arranges all asteroids by phase (level 1) and distance (level 2) with
// regard to base asteroid.
func (d Day10) sort(base Asteroid) {
	abs := func(idx int) float64 {
		return cmplx.Abs(d.asteroids[idx] - base)
	}
	phase := func(idx int) float64 {
		return cmplx.Phase(d.asteroids[idx] - base)
	}
	sort.Slice(d.asteroids, func(i, j int) bool {
		pi, pj := phase(i), phase(j)
		if pi < pj {
			return true
		}
		if pi > pj {
			return false
		}
		return abs(i) < abs(j)
	})
}

// find first asteroid in direction 'up' clockwise and return its
// index.
func (d Day10) findFirst(base Asteroid) int {
	d.sort(base)

	// 0 is (1,0)/ right/ 3 o'clock => up is -90°
	upPhase := -math.Pi / 2
	idx := 0
	for i, a := range d.asteroids {
		p := cmplx.Phase(a - base)
		if p >= upPhase {
			// ignore me myself and i
			if base == a {
				continue
			}
			idx = i
			break
		}
	}
	return idx
}
