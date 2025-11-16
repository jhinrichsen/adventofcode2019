package adventofcode2019

import (
	"math"
	"math/cmplx"
	"sort"
)

// An Asteroid is identified by a dimensionless (x/y) position in a 2D space.
// Using a type _alias_ here to avoid casting
type Asteroid = complex128

// ParseAsteroidMap parses newline separated strings into asteroids.
func ParseAsteroidMap(asteroids []byte) []Asteroid {
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
	var as []Asteroid
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
			as = append(as, a)
			x++
		} else {
			// whitespace, ignore
		}
	}
	return as
}

// Day10Part1 returns the asteroid that can see most asteroids, and the number of
// visible asteroids.
func Day10Part1(as []Asteroid) (Asteroid, int) {
	var best Asteroid
	var maxVisible int
	for i := range as {
		// map of angles
		visible := make(map[float64]Asteroid, len(as))
		for j := range as {
			// skip ourself
			if i == j {
				continue
			}
			rel := as[j] - as[i]
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
			best = as[i]
			maxVisible = len(visible)
		}
	}
	return best, maxVisible
}

// Day10Part2 determines the 200th asteroid that gets vaporized.
func Day10Part2(as []Asteroid, base Asteroid) int {
	as = center(vaporize(byPhase(center(as, base))), -base)
	a := as[199]
	// multiply its X coordinate by 100 and then add its Y coordinate
	return int(100.0*real(a) + imag(a))
}

type phaseGroup struct {
	phase     float64
	asteroids []Asteroid
}

// byPhase arranges normalized asteroids by phase (level 1) and distance (level
// 2) with regard to base asteroid. The phase group will contain all asteroids
// except for (0/0), i.e. len(as) -1 == len(phaseGroup).
func byPhase(as []Asteroid) []phaseGroup {
	var pgs []phaseGroup

	findIndex := func(φ float64) int {
		for i := 0; i < len(pgs); i++ {
			if pgs[i].phase == φ {
				return i
			}
		}
		return -1
	}
	// sweep #1: group by phase
	for _, a := range as {
		// ignore (0/0)
		if a == 0+0i {
			continue
		}
		φ := cmplx.Phase(a)
		idx := findIndex(φ)
		if idx == -1 {
			pgs = append(pgs, phaseGroup{φ, []Asteroid{a}})
		} else {
			pgs[idx].asteroids = append(pgs[idx].asteroids, a)
		}
	}

	// sweep #2: sort by phase asc
	sort.Slice(pgs, func(i, j int) bool {
		return pgs[i].phase < pgs[j].phase
	})

	// sweep #3: sort each phase by distance asc
	for i := range pgs {
		sort.Slice(pgs[i].asteroids, func(j, k int) bool {
			as := pgs[i].asteroids
			return cmplx.Abs(as[j]) < cmplx.Abs(as[k])
		})
	}
	return pgs
}

// countAsteroids returns number of asteroids
func countAsteroids(pgs []phaseGroup) int {
	n := 0
	for _, pg := range pgs {
		n += len(pg.asteroids)
	}
	return n
}

func center(as []Asteroid, base Asteroid) []Asteroid {
	var cas []Asteroid
	for i := range as {
		cas = append(cas, as[i]-base)
	}
	return cas
}

// find first asteroid in direction 'up' clockwise and return its
// index.
func findFirst(pgs []phaseGroup) int {
	// 0 is (1,0)/ right/ 3 o'clock => up is -90°
	upPhase := -math.Pi / 2
	idx := 0
	for i, pg := range pgs {
		if pg.phase >= upPhase {
			idx = i
			break
		}
	}
	return idx
}

func vaporize(pgs []phaseGroup) []Asteroid {
	var order []Asteroid
	idx := findFirst(pgs)

	n := countAsteroids(pgs)
	for i := 0; i < n; i++ {
		a := pgs[idx].asteroids[0]
		order = append(order, a)

		// housekeeping: remove asteroid from phase group
		if len(pgs[idx].asteroids) == 1 {
			// last asteroid for this phase, remove empty phase
			pgs = append(pgs[:idx], pgs[idx+1:]...)
			// don't increment index, next idx just popped in place
		} else {
			// remove first asteroid, leave the rest
			pgs[idx].asteroids = pgs[idx].asteroids[1:]
			// next phase
			idx++
		}
		// do not use %, len(pgs) = 0 in last run
		if idx == len(pgs) {
			idx = 0
		}
	}
	return order
}

// NewDay10 parses asteroid map from input
var NewDay10 = ParseAsteroidMap

// Day10 solves Monitoring Station puzzle
func Day10(input []byte, part1 bool) uint {
	asteroids := NewDay10(input)
	base, maxVisible := Day10Part1(asteroids)

	if part1 {
		return uint(maxVisible)
	}

	return uint(Day10Part2(asteroids, base))
}
