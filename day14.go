package adventofcode2019

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

const (
	fuel = "FUEL"
	ore  = "ORE"
)

// Bom holds chemicals during resolution.
type Bom map[string]uint

// Day14 holds a list of all reactions the nanofactory offers.
type Day14 struct {
	reactions []reaction
	store     Bom
}

// sort reactions by resolution order.
func (a *Day14) sort() {
	ls := level(a.reactions)
	sort.Slice(a.reactions, func(i, j int) bool {
		left := ls[a.reactions[i].output.unit]
		right := ls[a.reactions[j].output.unit]
		return left > right
	})
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

type reaction struct {
	input  map[value]bool
	output value
}

type value struct {
	quantity uint
	unit     string
}

func (a value) isFuel() bool {
	return a.unit == "FUEL"
}

func (a value) isOre() bool {
	return a.unit == ore
}

// Day14Part1 returns number of Ore required to create 1 FUEL.
func Day14Part1(d Day14) (uint, error) {
	resolved := d.reactions[0]
	for _, r := range d.reactions {
		resolved = resolve(resolved, r)
	}
	return resolved.input[0].quantity, nil
}

func resolve(base, dissolve reaction) reaction {
	unit := dissolve.output.unit
	for _, r := range base.input {
		if r.unit == unit {
			factor := base.output.quantity / dissolve.output.quantity
			// add all ingredients one by one
			for _, input := range dissolve.input {
				// add to existing or new value?
				_, ok := base.input[unit]
				if !ok {
					base.input[unit] = 1
				}
				base.input[unit] += factor * input.quantity
			}
		}
	}
}

// NewDay14 returns a Day14 from a list of reactions. Reactions are sorted by
// evaluation level, i.e. highest level of indirection first.
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
	d.sort()
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
	v.quantity = uint(q)
	v.unit = strings.TrimSpace(ss[1])
	return v, nil
}

func level(rs []reaction) Bom {
	levels := make(Bom)
	levels[ore] = 0

	complete := false
	for !complete {
		complete = true
		for _, r := range rs {
			var max uint
			for _, v := range r.input {
				l, ok := levels[v.unit]
				if !ok {
					complete = false
				}
				if l > max {
					max = l
				}
			}
			levels[r.output.unit] = max + 1
		}
	}
	return levels
}
