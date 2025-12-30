package adventofcode2019

import (
	"fmt"
	"strings"
)

// dims has three dimensions (x, y, z)
const (
	x     = 0
	y     = 1
	z     = 2
	dims  = 3
	moons = 4
)

type point struct {
	pos, vel int
}

func (a *point) gravity(b *point) {
	if a.pos < b.pos {
		a.vel++
		b.vel--
	} else if a.pos > b.pos {
		a.vel--
		b.vel++
	}
}

func (a *point) velocity() {
	a.pos += a.vel
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// +---------+---------+---------+-------+
// | moon 0  | moon 1  | moon 2  | moon 3| X
// +---------+---------+---------+-------+
// | moon 0  | moon 1  | moon 2  | moon 3| Y
// +---------+---------+---------+-------+
// | moon 0  | moon 1  | moon 2  | moon 3| Z
// +---------+---------+---------+-------+
type universe struct {
	moons [3][4]point
}

// String() returns the textual representation in format
// pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>
func (a universe) String() string {
	var sb strings.Builder
	for i := range moons {
		sb.WriteString(a.moon(i))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (a universe) moon(i int) string {
	return fmt.Sprintf("pos=<x=%2d, y=%1d, z=%2d>, "+
		"vel=<x=%2d, y=%1d, z=%2d>",
		a.moons[x][i].pos, a.moons[y][i].pos, a.moons[z][i].pos,
		a.moons[x][i].vel, a.moons[y][i].vel, a.moons[z][i].vel)
}

// gravity changes velocity of two moons for one axis
func (a *universe) gravity(d int, m1, m2 int) {
	if a.moons[d][m1].pos < a.moons[d][m2].pos {
		a.moons[d][m1].vel++
		a.moons[d][m2].vel--
	} else if a.moons[d][m1].pos > a.moons[d][m2].pos {
		a.moons[d][m1].vel--
		a.moons[d][m2].vel++
	}
}

func (a universe) dimension(dim int) [4]point {
	return [4]point{
		a.moons[dim][0],
		a.moons[dim][1],
		a.moons[dim][2],
		a.moons[dim][3],
	}
}

// cycle returns the number of discrete steps it takes a universe to return to
// its initial state. Each dimension cycles independently, so we find each
// dimension's cycle length and compute LCM.
func (a universe) cycle() int {
	// Find cycle for each dimension by counting steps to return to initial state
	cycleDim := func(dim int) int {
		initial := a.dimension(dim)
		u := a
		n := 0
		for {
			u.step(dim)
			n++
			if u.dimension(dim) == initial {
				return n
			}
		}
	}

	c1 := cycleDim(x)
	c2 := cycleDim(y)
	c3 := cycleDim(z)

	// Calculate LCM(c1, c2, c3) = LCM(LCM(c1, c2), c3)
	lcm12 := c1 * c2 / gcd(c1, c2)
	return lcm12 * c3 / gcd(lcm12, c3)
}

func gcd(a, b int) int {
	if a == 0 {
		return abs(b)
	}
	if b == 0 {
		return abs(a)
	}
	for {
		h := a % b
		a = b
		b = h
		if b == 0 {
			break
		}
	}
	return abs(a)
}

func (a *universe) step(dim int) {
	// apply gravity
	//    A   B   C
	// A      *   *
	// B          *
	// C
	for i := range moons {
		for j := i + 1; j < moons; j++ {
			a.moons[dim][i].gravity(&a.moons[dim][j])
		}
	}

	// apply velocity
	for i := range moons {
		a.moons[dim][i].velocity()
	}
}

// energy of a universe is the sum of the energy of all moons. The total energy
// for a single moon is its potential energy multiplied by its kinetic energy. A
// moon's potential energy is the sum of the absolute values of its x, y, and
// z position coordinates. A moon's kinetic energy is the sum of the absolute
// values of its velocity coordinates.
func (a universe) energy() int {
	var sum int
	for j := range moons {
		var pot, kin int
		for i := range dims {
			moon := a.moons[i][j]
			pot += abs(moon.pos)
			kin += abs(moon.vel)
		}
		sum += pot * kin
	}
	return sum
}
