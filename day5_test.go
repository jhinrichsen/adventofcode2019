package adventofcode2019

import (
	"fmt"
	"testing"
)

var day5Examples = []struct {
	in, out string
}{
	{"1,0,0,0,99", "2,0,0,0,99"},
	{"2,3,0,3,99", "2,3,0,6,99"},
	{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
	{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	{"1002,4,3,4,33", "1002,4,3,4,99"},
	{"1101,100,-1,4,0", "1101,100,-1,4,99"},
}

func TestDay5Part1Examples(t *testing.T) {
	for _, tt := range day5Examples {
		id := fmt.Sprintf("Runs(%s)", tt.in)
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
	in, out := make(chan int), make(chan int)
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
	i, o := make(chan int), make(chan int)
	// make sure Day5 implements interface
	var cpu IntCodeProcessor = Day5
	go cpu(MustSplit("3,0,4,0,99"), i, o)
	i <- want
	got := <-o
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
