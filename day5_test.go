package adventofcode2019

import (
	"fmt"
	"testing"
)

var day5Part1Examples = []struct {
	in, out string
}{
	{"1,0,0,0,99", "2,0,0,0,99"},
	{"2,3,0,3,99", "2,3,0,6,99"},
	{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
	{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	{"1002,4,3,4,33", "1002,4,3,4,99"},
	{"1101,100,-1,4,0", "1101,100,-1,4,99"},
}

const (

	// The program will output 999 if the input value is below 8, output 1000
	// if the input value is equal to 8, or output 1001 if the input value is
	// greater than 8.
	day5Part2LargeExample = "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31," +
		"1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104," +
		"999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
)

var day5Part2Examples = []struct {
	prog        string
	in, out     int
	description string
}{
	{"3,9,8,9,10,9,4,9,99,-1,8", 0, 0, "0 == 8 ? 1 : 0 (position mode)"},
	{"3,9,8,9,10,9,4,9,99,-1,8", 7, 0, "7 == 8 ? 1 : 0 (position mode)"},
	{"3,9,8,9,10,9,4,9,99,-1,8", 8, 1, "8 == 8 ? 1 : 0 (position mode)"},
	{"3,9,8,9,10,9,4,9,99,-1,8", 9, 0, "9 == 8 ? 1 : 0 (position mode)"},

	{"3,9,7,9,10,9,4,9,99,-1,8", 0, 0, "0 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 7, 0, "7 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 8, 1, "8 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 9, 0, "9 < 8 ? 1 : 0 (position mode)"},

	{"3,3,1108,-1,8,3,4,3,99", 0, 0, "0 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 7, 0, "7 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 8, 1, "8 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 9, 0, "9 == 8 ? 1 : 0 (immediate mode)"},

	{"3,3,1107,-1,8,3,4,3,99", 0, 0, "0 < 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1107,-1,8,3,4,3,99", 7, 0, "7 < 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1107,-1,8,3,4,3,99", 8, 1, "8 < 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1107,-1,8,3,4,3,99", 9, 0, "9 < 8 ? 1 : 0 (immediate mode)"},

	{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 0, 0, "0 == 0 ? 0 : 1 (position mode)"},
	{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 1, 1, "1 == 0 ? 0 : 1 (position mode)"},
	{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 2, 1, "2 == 0 ? 0 : 1 (position mode)"},
	{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 3, 1, "3 == 0 ? 0 : 1 (position mode)"},

	{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, 0, "0 == 0 ? 0 : 1 (immediate mode)"},
	{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 1, 1, "1 == 0 ? 0 : 1 (immediate mode)"},
	{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 2, 1, "2 == 0 ? 0 : 1 (immediate mode)"},
	{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 3, 1, "3 == 0 ? 0 : 1 (immediate mode)"},

	{day5Part2LargeExample, 0, 999, "0 < 8: 999, == 8: 1000, > 8: 1001"},
	{day5Part2LargeExample, 7, 999, "7 < 8: 999, == 8: 1000, > 8: 1001"},
	{day5Part2LargeExample, 8, 1000, "8 < 8: 999, == 8: 1000, > 8: 1001"},
	{day5Part2LargeExample, 9, 1001, "9 < 8: 999, == 8: 1000, > 8: 1001"},
	{day5Part2LargeExample, 1000, 1001, "1000 < 8: 999, == 8: 1000, > 8: 1001"},
}

func TestDay5Part1Examples(t *testing.T) {
	for _, tt := range day5Part1Examples {
		id := fmt.Sprintf("Day5(%s)", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			proc := MustSplit(tt.in)
			Day5(proc, make(chan int), make(chan int))
			got := ToString(proc)
			if got != tt.out {
				t.Fatalf("%s: want %s but got %s", id,
					want, got)
			}
		})
	}
}

func channels() (chan int, chan int) {
	return make(chan int), make(chan int)
}

func TestDay5Part1(t *testing.T) {
	in, out := channels()
	lines, err := Lines("testdata/day5.txt")
	if err != nil {
		t.Fatal(err)
	}
	go Day5(MustSplit(lines[0]), in, out)
	in <- 1    // ID for air conditioner unit
	var dc int // diagnostic code

	// we do not know when output finishes, so we cannot abort on the first
	// non-0 value. Instead, keep a history accumulation
	var rcs []int
	for rc := range out {
		dc = rc
		rcs = append(rcs, rc)
	}
	// delete last entry, the diagnostic code
	rcs = rcs[:len(rcs)-1]

	// Check all rcs
	for i, rc := range rcs {
		if rc != 0 {
			t.Fatalf("rc #%d: want 0 but got %d", i, rc)
		}
	}

	// Check diagnostic code
	want := 16225258
	got := dc
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestIdentity(t *testing.T) {
	want := 42
	i, o := channels()
	// make sure Day5 implements interface
	var cpu IntCodeProcessor = Day5
	go cpu(MustSplit("3,0,4,0,99"), i, o)
	i <- want
	got := <-o
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay5Part2Examples(t *testing.T) {
	for _, tt := range day5Part2Examples[1:2] {
		id := fmt.Sprintf("Day5(%s)", tt.description)
		prog := MustSplit(tt.prog)
		t.Run(id, func(t *testing.T) {
			in, out := channels()
			go Day5(prog, in, out)
			in <- tt.in
			want := tt.out
			got := <-out
			if want != got {
				t.Fatalf("%s: want %d but got %d", id,
					want, got)
			}
		})
	}
}

func TestDay5Part2(t *testing.T) {
	want := 2808771
	got := -1
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
