package adventofcode2019

import (
	"fmt"
	"strings"
)

// com is the Center of Mass
const com = "COM"

// Day06Puzzle holds an object and its orbit.
type Day06Puzzle struct {
	orbits map[string]string
}

// NewDay06 creates a new Orbit from a list of 'a)b' lines.
func NewDay06(ss []string) (Day06Puzzle, error) {
	d := Day06Puzzle{
		orbits: make(map[string]string, len(ss)),
	}
	for i, s := range ss {
		// Inline parsing instead of strings.Split
		idx := strings.IndexByte(s, ')')
		if idx < 0 {
			return d, fmt.Errorf("bad input in line %d: %s", i, s)
		}
		// Save object: orbits[child] = parent
		d.orbits[s[idx+1:]] = s[:idx]
	}
	return d, nil
}

// orbit returns the orbit of an object.
func (a Day06Puzzle) orbit(object string) string {
	return a.orbits[object]
}

// orbitCount returns the number of orbits of a given object.
func (a Day06Puzzle) orbitCount(object string) int {
	n := 0
	for object != com {
		object = a.orbit(object)
		n++
	}
	return n
}

// orbitCountChecksum returns the checksum for a complete orbit.
func (a Day06Puzzle) orbitCountChecksum() int {
	// Memoize orbit counts to avoid repeated tree traversals
	cache := make(map[string]int, len(a.orbits))
	var orbitCount func(object string) int
	orbitCount = func(object string) int {
		if object == com {
			return 0
		}
		if n, ok := cache[object]; ok {
			return n
		}
		n := 1 + orbitCount(a.orbits[object])
		cache[object] = n
		return n
	}

	sum := 0
	for object := range a.orbits {
		sum += orbitCount(object)
	}
	return sum
}

// commonOrbit returns the nearest orbit of two objects, at least COM.
func (a Day06Puzzle) commonOrbit(object1, object2 string) string {
	// align both objects to same orbit distance
	for a.orbitCount(object1) > a.orbitCount(object2) {
		object1 = a.orbit(object1)
	}
	for a.orbitCount(object2) > a.orbitCount(object1) {
		object2 = a.orbit(object2)
	}
	for object1 != com && object2 != com {
		if object1 == object2 {
			return object1
		}
		object1 = a.orbit(object1)
		object2 = a.orbit(object2)
	}
	return com
}

// transfers counts the number of hops between an object up to the nearest
// common orbit, and then down to the the second object.
func (a Day06Puzzle) transfers(object1, object2 string) int {
	c := a.commonOrbit(object1, object2)
	nc := a.orbitCount(c)
	return (a.orbitCount(object1) - 1 - nc) + (a.orbitCount(object2) - 1 - nc)
}

// Day06 solves Universal Orbit Map puzzle
func Day06(puzzle Day06Puzzle, part1 bool) uint {
	if part1 {
		return uint(puzzle.orbitCountChecksum())
	}
	return uint(puzzle.transfers("YOU", "SAN"))
}
