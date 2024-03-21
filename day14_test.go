package adventofcode2019

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testDay14ExamplesWants = []uint{
	31,
	165,
	13_312,
	180_697,
	2_210_736,
}

func day14ExampleFilename(i int) string {
	// filename index is 1-based
	return fmt.Sprintf("testdata/day14_example%d.txt", i+1)
}

func day14Example1(t *testing.T) Day14 {
	f, err := os.Open(day14ExampleFilename(0))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	d, err := NewDay14(f)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func TestDay14Example1(t *testing.T) {
	want := testDay14ExamplesWants[0]
	d := day14Example1(t)
	got, err := Day14Part1(d)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func NoTestDay14Examples(t *testing.T) {
	for i, want := range testDay14ExamplesWants {
		id := fmt.Sprintf("Day14Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			f, err := os.Open(day14ExampleFilename(i))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			d, err := NewDay14(f)
			if err != nil {
				t.Fatal(err)
			}
			got, err := Day14Part1(d)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestParseReaction(t *testing.T) {
	want := reaction{
		input: map[value]bool{
			{
				quantity: 2,
				unit:     "AB",
			}: true,
			{
				quantity: 3,
				unit:     "BC",
			}: true,
			{
				quantity: 4,
				unit:     "CA",
			}: true,
		},
		output: value{
			quantity: 1,
			unit:     "FUEL",
		},
	}
	got, err := parseReaction("2 AB, 3 BC, 4 CA => 1 FUEL")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestParseValue(t *testing.T) {
	want := value{
		quantity: 2,
		unit:     "AB",
	}
	got, err := parseValue("2 AB")
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestFuel(t *testing.T) {
	want := true
	got := value{
		unit: "FUEL",
	}.isFuel()
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestResolve(t *testing.T) {
	r1 := reaction{
		map[value]bool{{3, ore}: true},
		value{5, "A"},
	}
	r2 := reaction{
		map[value]bool{{7, "A"}: true},
		value{1, fuel},
	}
	r := resolve(r1, r2)
	const want = 6
	got := r.input[ore].quantity
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Level(t *testing.T) {
	const want = 5
	d := day14Example1(t)
	ls := level(d.reactions)
	got := ls["FUEL"]
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay14Sort(t *testing.T) {
	const want = "FUEL"
	got := day14Example1(t).reactions[0].output.unit
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}
