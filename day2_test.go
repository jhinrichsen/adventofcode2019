package adventofcode2019

import (
	"fmt"
	"reflect"
	"testing"
)

var examples = []struct {
	in, out string
}{
	{"1,0,0,0,99", "2,0,0,0,99"},
	{"2,3,0,3,99", "2,3,0,6,99"},
	{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
	{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
}

func TestDay2Len(t *testing.T) {
	want := 99
	opcodes, err := Split(examples[0].in)
	if err != nil {
		t.Fatal(err)
	}
	got := Len(opcodes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay2Split(t *testing.T) {
	want := []int{1, 0, 0, 0, 99}
	got, err := Split(examples[0].in)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay2Part1Examples(t *testing.T) {
	for _, tt := range examples {
		id := fmt.Sprintf("Runs(%s)", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Runs(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("%s: want %s but got %s", id,
					want, got)
			}
		})
	}
}

func TestDay2Part1(t *testing.T) {
	lines, err := Lines("testdata/day2.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Restore 1202 program alarm
	opcodes, err := Split(lines[0])
	if err != nil {
		t.Fatal(err)
	}
	opcodes[1] = 12
	opcodes[2] = 2
	opcodes, err = Run(opcodes)
	if err != nil {
		t.Fatal(err)
	}
	want := 3562624
	got := opcodes[0]
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
