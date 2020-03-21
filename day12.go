package adventofcode2019

import "fmt"

// DIM has three dimensions (x, y, z)
const (
	X   = 0
	Y   = 1
	Z   = 2
	DIM = 3
)

// D3 in outer space has 3 axis. 3D is not a valid Go identifier.
type D3 = [DIM]int

type moon struct {
	pos D3
	vel D3
}

func (a moon) energy() int {
	pot := energy(a.pos)
	kin := energy(a.vel)
	return pot * kin
}

func energy(d D3) int {
	sum := 0
	for i := 0; i < DIM; i++ {
		sum += abs(d[i])
	}
	return sum
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// pot A moon's potential energy is the sum of the absolute values of its x, y,
// and z position coordinates.
func (a moon) pot() int {
	sum := 0
	for i := 0; i < DIM; i++ {
		sum += abs(a.pos[i])
	}
	return sum
}

// gravity changes both moons, receiver and parameter.
func (a *moon) gravity(m2 *moon) {
	for dim := 0; dim < DIM; dim++ {
		if a.pos[dim] < m2.pos[dim] {
			a.vel[dim]++
			m2.vel[dim]--
		} else if a.pos[dim] > m2.pos[dim] {
			a.vel[dim]--
			m2.vel[dim]++
		} else {
			// no change
		}
	}
}

func (a *moon) velocity() {
	for dim := 0; dim < DIM; dim++ {
		a.pos[dim] += a.vel[dim]
	}
}

// String() returns the textual representation of a moon in format
// pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>
func (a moon) String() string {
	rep := func(d D3) string {
		return fmt.Sprintf("<x=%2d, y=%1d, z=%2d>", d[X], d[Y], d[Z])
	}
	return fmt.Sprintf("pos=%s, vel=%s", rep(a.pos), rep(a.vel))
}

type universe struct {
	moons []moon
}

// newUniverse creates a new universe with given positions and velocity 0 for
// all moons.
func newUniverse(positions []D3) universe {
	var u universe
	u.moons = make([]moon, len(positions))
	for i := 0; i < len(positions); i++ {
		u.moons[i].pos = positions[i]
		// velocity starts at 0
	}
	return u
}

func (a *universe) applyGravity() {
	//    A   B   C
	// A      *   *
	// B          *
	// C
	for i := 0; i < len(a.moons); i++ {
		for j := i + 1; j < len(a.moons); j++ {
			a.moons[i].gravity(&a.moons[j])
		}
	}
}

func (a *universe) applyVelocity() {
	for i := 0; i < len(a.moons); i++ {
		a.moons[i].velocity()
	}
}

func (a universe) step() {
	a.applyGravity()
	a.applyVelocity()
}

// energy of a universe is the sum of the energy of all moons.
func (a universe) energy() int {
	sum := 0
	for i := range a.moons {
		sum += a.moons[i].energy()
	}
	return sum
}
