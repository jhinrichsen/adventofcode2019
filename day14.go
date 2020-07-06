package adventofcode2019

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Bom map[string]int

// Day14 holds a list of all reactions the nanofactory offers.
type Day14 struct {
	reactions []reaction
	// store for rest, e.g. 3 if we need 7 but can only produce 10
	store Bom
}

// fuel returns the first reaction that has FUEL as output.
func (a Day14) fuel() (reaction, error) {
	for i := range a.reactions {
		r := a.reactions[i]
		if r.output.isFuel() {
			return r, nil
		}
	}
	return reaction{}, fmt.Errorf("no FUEL reaction")
}

func (a Day14) reduce(v value) ([]value, error) {
	for i := range a.reactions {
		r := a.reactions[i]
		if r.output.unit == v.unit {
			// factor * 5 >= 7
			factor := v.quantity / r.output.quantity
			if v.quantity%r.output.quantity > 0 {
				factor++
			}
			return mul(r.input, factor), nil
		}
	}
	return nil, fmt.Errorf("unit %q not found in reactions %+v",
		v.unit, a.reactions)
}

func mul(rs []value, n int) []value {
	for i := range rs {
		rs[i].mul(n)
	}
	return rs
}

type reaction struct {
	input  []value
	output value
}

type value struct {
	quantity int
	unit     string
}

func (a value) isFuel() bool {
	return a.unit == "FUEL"
}

func (a value) isOre() bool {
	return a.unit == "ORE"
}

func (a *value) mul(n int) {
	a.quantity *= n
}

// Day14Part1 returns number of Ore required to create 1 FUEL.
func Day14Part1(d Day14) (int, error) {
	fuel, err := d.fuel()
	if err != nil {
		return -1, fmt.Errorf("missing FUEL in %+v", d)
	}
	bom := make(Bom)
	for _, v := range fuel.input {
		bom[v.unit] = v.quantity
	}

	// reduce until only one unit left
	for len(bom) > 1 {
		for unit, quantity := range bom {
			v := value{quantity, unit}
			if v.isOre() {
				continue
			}
			vs, err := d.reduce(v)
			if err != nil {
				return -1,
					fmt.Errorf("cannot reduce %+v in %+v",
						v, d)
			}

			// add to map, optionally creating a new entry
			for _, v := range vs {
				if _, ok := bom[v.unit]; ok {
					bom[v.unit] += v.quantity
				} else {
					bom[v.unit] = v.quantity
				}
			}

			// reduced, remove original from bom
			delete(bom, unit)
		}
	}
	return bom["ORE"], nil
}

// NewDay14 returns a Day14 from a list of reactions.
func NewDay14(reactions io.Reader) (Day14, error) {
	var d Day14
	lines, err := linesFromReader(reactions)
	if err != nil {
		return d, err
	}
	for i := range lines {
		reaction, err := parseReaction(lines[i])
		if err != nil {
			return d, fmt.Errorf("error parsing line %d: %v",
				i, err)
		}
		d.reactions = append(d.reactions, reaction)
	}
	return d, nil
}

// parseReaction returns a struct from the textual representation such as
// 2 AB, 3 BC, 4 CA => 1 FUEL.
func parseReaction(s string) (reaction, error) {
	var r reaction
	ps := strings.Split(s, "=>")
	if len(ps) != 2 {
		return r, fmt.Errorf("reactions contains no => separator: %q",
			s)
	}

	// input
	is := strings.Split(ps[0], ",")
	if len(is) == 0 {
		return r, fmt.Errorf("missing input values: %q", is)
	}
	for i := range is {
		v, err := parseValue(is[i])
		if err != nil {
			return r, err
		}
		r.input = append(r.input, v)
	}

	// output
	v, err := parseValue(ps[1])
	if err != nil {
		return r, err
	}
	r.output = v
	return r, nil
}

// "2 AB" -> value{quantity: 2, unit: AB}
func parseValue(s string) (value, error) {
	var v value
	// parse 2 AB
	ss := strings.Fields(s)
	if len(ss) != 2 {
		return v, fmt.Errorf("missing quantity or unit: %q", ss)
	}
	// "2" -> 2
	qs := strings.TrimSpace(ss[0])
	q, err := strconv.Atoi(qs)
	if err != nil {
		return v, fmt.Errorf("error parsing quantity: %q", qs)
	}
	v.quantity = q
	v.unit = strings.TrimSpace(ss[1])
	return v, nil
}
