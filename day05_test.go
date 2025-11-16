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
	AirContitionerUnit        = 1
	ThermalRadiatorController = 5

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

	{"3,9,7,9,10,9,4,9,99,-1,8", 0, 1, "0 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 7, 1, "7 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 8, 0, "8 < 8 ? 1 : 0 (position mode)"},
	{"3,9,7,9,10,9,4,9,99,-1,8", 9, 0, "9 < 8 ? 1 : 0 (position mode)"},

	{"3,3,1108,-1,8,3,4,3,99", 0, 0, "0 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 7, 0, "7 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 8, 1, "8 == 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1108,-1,8,3,4,3,99", 9, 0, "9 == 8 ? 1 : 0 (immediate mode)"},

	{"3,3,1107,-1,8,3,4,3,99", 0, 1, "0 < 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1107,-1,8,3,4,3,99", 7, 1, "7 < 8 ? 1 : 0 (immediate mode)"},
	{"3,3,1107,-1,8,3,4,3,99", 8, 0, "8 < 8 ? 1 : 0 (immediate mode)"},
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

func TestDay5Part1(t *testing.T) {
	want := 16225258
	in, out := channels()
	lines, err := linesFromFilename(input(5))
	if err != nil {
		t.Fatal(err)
	}
	go Day5(MustSplit(lines[0]), in, out)
	in <- AirContitionerUnit
	got := last(t, out)
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

// TestDay5Example2 runs a little program to check multiply.
// For example, consider the program 1002,4,3,4,33.
//
// The first instruction, 1002,4,3,4, is a multiply instruction - the rightmost
// two digits of the first value, 02, indicate opcode 2, multiplication. Then,
// going right to left, the parameter modes are 0 (hundreds digit), 1 (thousands
// digit), and 0 (ten-thousands digit, not present and therefore zero):
// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,
//
//	omitted due to being a leading zero
//
// This instruction multiplies its first two parameters. The first parameter, 4
// in position mode, works like it did before - its value is the value stored at
// address 4 (33). The second parameter, 3 in immediate mode, simply has value
// 3. The result of this operation, 33 * 3 = 99, is written according to the
// third parameter, 4 in position mode, which also works like it did before - 99
// is written to address 4.
func TestDay5Part2Example(t *testing.T) {
	prog := MustSplit("1002,4,3,4,33")
	in, out := channels()
	Day5(prog, in, out)
	// we would not be here if last code wasn't 99
	want := 99
	got := prog[4]
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay5Part2Examples(t *testing.T) {
	for _, tt := range day5Part2Examples {
		id := fmt.Sprintf("Day5(%s)", tt.description)
		prog := MustSplit(tt.prog)
		t.Run(id, func(t *testing.T) {
			in, out := channels()
			go Day5(prog, in, out)
			in <- tt.in
			want := tt.out
			// one or multiple outputs that must be 0?
			var got int
			if tt.prog == day5Part2LargeExample {
				got = last(t, out)
			} else {
				got = <-out
			}
			if want != got {
				t.Fatalf("%s: want %d but got %d", id,
					want, got)
			}
		})
	}
}

func last(t *testing.T, c chan int) int {
	var dc int // diagnostic code

	// we do not know when output finishes, so we cannot abort on the first
	// non-0 value. Instead, keep a history accumulation
	var rcs []int
	for rc := range c {
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
	return dc
}

func TestDay5Part2(t *testing.T) {
	want := 2808771
	in, out := channels()
	lines, err := linesFromFilename(input(5))
	if err != nil {
		t.Fatal(err)
	}
	go Day5(MustSplit(lines[0]), in, out)
	in <- ThermalRadiatorController
	got := <-out
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay05Part2(b *testing.B) {
	lines, err := linesFromFilename(input(5))
	if err != nil {
		b.Fatal(err)
	}
	master := MustSplit(lines[0])
	for i := 0; i < b.N; i++ {
		in, out := channels()
		// Run each step in its own copy
		go Day5(master.Copy(), in, out)
		in <- ThermalRadiatorController
		<-out
	}
}

func BenchmarkDay05Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(5))
	master := MustSplit(lines[0])
	for b.Loop() {
		in, out := channels()
		go Day5(master.Copy(), in, out)
		in <- AirContitionerUnit
		// Drain all output
		for range out {
		}
	}
}
