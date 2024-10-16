package adventofcode2019

import (
	"fmt"
	"reflect"
	"testing"
)

var day2Examples = []struct {
	in, out string
}{
	{"1,0,0,0,99", "2,0,0,0,99"},
	{"2,3,0,3,99", "2,3,0,6,99"},
	{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
	{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
}

func TestDay2Len(t *testing.T) {
	want := 99
	opcodes, err := Split(day2Examples[0].in)
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
	got, err := Split(day2Examples[0].in)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay2Part1Examples(t *testing.T) {
	for _, tt := range day2Examples {
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
	lines, err := linesFromFilename(input(2))
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

func TestDay2Part2(t *testing.T) {
	lines, err := linesFromFilename(input(2))
	if err != nil {
		t.Fatal(err)
	}

	var solution int
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// Restore 1202 program alarm
			opcodes, err := Split(lines[0])
			if err != nil {
				t.Fatal(err)
			}

			opcodes[1] = noun
			opcodes[2] = verb
			opcodes, err = Run(opcodes)
			if err != nil {
				t.Fatal(err)
			}
			want := 19690720
			got := opcodes[0]
			if want == got {
				solution = 100*noun + verb
				goto solved
			}
		}
	}
	t.Fatalf("no solution found for noun 0..99, verb 0..99")
solved:
	want := 8298
	got := solution
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
