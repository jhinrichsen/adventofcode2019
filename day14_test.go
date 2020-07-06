package adventofcode2019

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

var testDay14ExamplesWants = []int{
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

func TestExample1(t *testing.T) {
	want := testDay14ExamplesWants[0]
	f, err := os.Open(day14ExampleFilename(0))
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
		input: []value{
			{
				quantity: 2,
				unit:     "AB",
			},
			{
				quantity: 3,
				unit:     "BC",
			},
			{
				quantity: 4,
				unit:     "CA",
			},
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
	s := strings.Join([]string{
		"3 ORE => 5 a",
		"7 a => 1 FUEL",
	}, "\n")
	d, err := NewDay14(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}

	fuel, err := d.fuel()
	if err != nil {
		t.Fatal(err)
	}
	wantFuelInput := value{7, "a"}
	gotFuelInput := fuel.input[0]
	if wantFuelInput != gotFuelInput {
		t.Fatalf("want fuel value %v but got %v", wantFuelInput,
			gotFuelInput)
	}

	wantOreValue := value{6, "ORE"}
	gotOreValues, err := d.reduce(gotFuelInput)
	if err != nil {
		t.Fatal(err)
	}
	if wantOreValue != gotOreValues[0] {
		t.Fatalf("want %v but got %v", wantOreValue, gotOreValues[0])
	}
}
