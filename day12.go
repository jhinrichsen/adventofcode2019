package adventofcode2019

import (
	"fmt"
	"strings"
)

// DIMS has three dimensions (x, y, z)
const (
	X     = 0
	Y     = 1
	Z     = 2
	DIMS  = 3
	MOONS = 4
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
	} else {
		// no change
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
	for i := range MOONS {
		sb.WriteString(a.moon(i))
		sb.WriteString("\n")
	}
	return sb.String()
}

func (a universe) moon(i int) string {
	return fmt.Sprintf("pos=<x=%2d, y=%1d, z=%2d>, "+
		"vel=<x=%2d, y=%1d, z=%2d>",
		a.moons[X][i].pos, a.moons[Y][i].pos, a.moons[Z][i].pos,
		a.moons[X][i].vel, a.moons[Y][i].vel, a.moons[Z][i].vel)
}

// gravity changes velocity of two moons for one axis
func (a *universe) gravity(d int, m1, m2 int) {
	if a.moons[d][m1].pos < a.moons[d][m2].pos {
		a.moons[d][m1].vel++
		a.moons[d][m2].vel--
	} else if a.moons[d][m1].pos > a.moons[d][m2].pos {
		a.moons[d][m1].vel--
		a.moons[d][m2].vel++
	} else {
		// no change
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
// a state it has been before, and returns 0 in case of overflow.
// some hints suggest that x, y, and z dimension all have their own cycle, and
// the total cycle can be deducted by multiplying (?) single cycles.
func (a universe) cycle() int {
	c := make(chan int, 3)
	fn := func(dim int, c chan int) {
		n := 0
		history := make(map[[4]point]bool)
		for {
			a.step(dim)
			n++
			d := a.dimension(dim)
			if _, ok := history[d]; ok {
				n--
				c <- n
				break
			}
			history[d] = true
		}
	}
	for dim := range DIMS {
		go fn(dim, c)
	}
	c1 := <-c
	c2 := <-c
	c3 := <-c
	// Multiplying creates a number ten times too high
	// d1 := gcd(c1, c2)
	d2 := gcd(c2, c3)
	d3 := gcd(c3, c1)
	return c1 * c2 / d2 * c3 / d3
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
	for i := range MOONS {
		for j := i + 1; j < MOONS; j++ {
			a.moons[dim][i].gravity(&a.moons[dim][j])
		}
	}

	// apply velocity
	for i := range MOONS {
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
	for j := range MOONS {
		var pot, kin int
		for i := range DIMS {
			moon := a.moons[i][j]
			pot += abs(moon.pos)
			kin += abs(moon.vel)
		}
		sum += pot * kin
	}
	return sum
}
