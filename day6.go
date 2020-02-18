package adventofcode2019

import (
	"fmt"
	"strings"
)

// Center of Mass
const COM = "COM"

type Day6 struct {
	orbits map[string]string
}

func NewDay6(ss []string) (Day6, error) {
	d := Day6{
		orbits: make(map[string]string, len(ss)),
	}
	for i, s := range ss {
		parts := strings.Split(s, ")")
		if len(parts) != 2 {
			return d, fmt.Errorf("bad input in line %d: %s", i, s)
		}

		// Save object
		d.orbits[parts[1]] = parts[0]
	}
	return d, nil
}

func (a Day6) Orbit(object string) string {
	return a.orbits[object]
}

func (a Day6) OrbitCount(object string) int {
	n := 0
	for object != COM {
		object = a.Orbit(object)
		n++
	}
	return n
}

func (a Day6) OrbitCountChecksum() int {
	sum := 0
	for object := range a.orbits {
		n := a.OrbitCount(object)
		sum += n
	}
	return sum
}

func (a Day6) CommonOrbit(object1, object2 string) string {
	// align both objects to same orbit distance
	for a.OrbitCount(object1) > a.OrbitCount(object2) {
		object1 = a.Orbit(object1)
	}
	for a.OrbitCount(object2) > a.OrbitCount(object1) {
		object2 = a.Orbit(object2)
	}
	for object1 != COM && object2 != COM {
		if object1 == object2 {
			return object1
		}
		object1 = a.Orbit(object1)
		object2 = a.Orbit(object2)
	}
	return COM
}

// Between the objects they are orbiting - not between YOU and SAN.
func (a Day6) Transfers(object1, object2 string) int {
	c := a.CommonOrbit(object1, object2)
	nc := a.OrbitCount(c)
	return (a.OrbitCount(object1) - 1 - nc) + (a.OrbitCount(object2) - 1 - nc)
}
