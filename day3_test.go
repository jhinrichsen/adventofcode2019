package adventofcode2019

import (
	"fmt"
	"testing"
)

var day3Examples = []struct {
	in  []string
	out int
}{
	{[]string{
		"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83",
	}, 159},
	{[]string{
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
	}, 135},
}

// Example stopped working after Part 1 - no clue
func Day3Part1Example(t *testing.T) {
	for i, tt := range day3Examples {
		id := fmt.Sprintf("Example #%d", i)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Day3(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("%s: want %d but got %d", id,
					want, got)
			}
		})
	}
}

func TestParse(t *testing.T) {
	wantD, wantN := Right, 10
	gotD, gotN, err := Parse("R10")
	if err != nil {
		t.Fatal(err)
	}
	if wantD != gotD {
		t.Fatalf("want direction %v but got %v", wantD, gotD)
	}
	if wantN != gotN {
		t.Fatalf("want length %v but got %v", wantN, gotN)
	}
}

func TestSize(t *testing.T) {
	wantX, wantY := 7, 3
	gotX, gotY, err := Size("R7,U1,L1,U1,L1,U1,L1")
	if err != nil {
		t.Fatal(err)
	}
	if wantX != gotX {
		t.Fatalf("want width %d but got %d", wantX, gotX)
	}
	if wantY != gotY {
		t.Fatalf("want height %d but got %d", wantY, gotY)
	}
}

func TestDay3Part1(t *testing.T) {
	lines, err := Lines("testdata/day3.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 248
	got, err := Day3(lines)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay3Part1(b *testing.B) {
	lines, err := Lines("testdata/day3.txt")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Day3(lines)
	}
}
