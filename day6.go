package adventofcode2019

import (
	"fmt"
	"strings"
)

// Center of Mass
const COM = "COM"

type Objects map[string]bool

type Day6 struct {
	orbiters map[string]Objects
	objects  Objects
}

func NewDay6(ss []string) (Day6, error) {
	d := Day6{
		objects:  make(map[string]bool, 2*len(ss)),
		orbiters: make(map[string]Objects),
	}
	for i, s := range ss {
		parts := strings.Split(s, ")")
		if len(parts) != 2 {
			return d, fmt.Errorf("bad input in line %d: %s", i, s)
		}

		// Save object
		d.objects[parts[0]] = true
		d.objects[parts[1]] = true

		// Its orbiter
		if os, ok := d.orbiters[parts[0]]; ok {
			// add orbiter to existing entry
			os[parts[1]] = true
		} else {
			// add new orbit
			os := make(Objects)
			os[parts[1]] = true
			d.orbiters[parts[0]] = os
		}
	}
	return d, nil
}

func (a Day6) OrbitCount(object string) int {
	n := 0
	for object != COM {
		for k, v := range a.orbiters {
			if _, ok := v[object]; ok {
				object = k
				n++
			}
		}
	}
	return n
}

func (a Day6) OrbitCountChecksum() int {
	sum := 0
	for object := range a.objects {
		n := a.OrbitCount(object)
		sum += n
	}
	return sum
}
