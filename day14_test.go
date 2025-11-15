package adventofcode2019

import (
	"fmt"
	"testing"
)

var day14Part1Examples = []struct {
	filename string
	want     uint
}{
	{"testdata/day14_example1.txt", 31},
	{"testdata/day14_example2.txt", 165},
	{"testdata/day14_example3.txt", 13312},
	{"testdata/day14_example4.txt", 13312},
	{"testdata/day14_example5.txt", 180697},
	{"testdata/day14_example6.txt", 2210736},
}

func TestDay14Part1Examples(t *testing.T) {
	for i, tt := range day14Part1Examples {
		id := fmt.Sprintf("Day14Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			lines, err := linesFromFilename(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			want := tt.want
			got := Day14(lines, true)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay14Part1(t *testing.T) {
	lines, err := linesFromFilename(input(14))
	if err != nil {
		t.Fatal(err)
	}
	want := uint(337862)
	got := Day14(lines, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	lines, err := linesFromFilename(input(14))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Day14(lines, true)
	}
}

var day14Part2Examples = []struct {
	filename string
	want     uint
}{
	{"testdata/day14_example3.txt", 82892753},
	{"testdata/day14_example4.txt", 82892753},
	{"testdata/day14_example5.txt", 5586022},
	{"testdata/day14_example6.txt", 460664},
}

func TestDay14Part2Examples(t *testing.T) {
	for i, tt := range day14Part2Examples {
		id := fmt.Sprintf("Day14Part2 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			lines, err := linesFromFilename(tt.filename)
			if err != nil {
				t.Fatal(err)
			}
			want := tt.want
			got := Day14(lines, false)
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay14Part2(t *testing.T) {
	lines, err := linesFromFilename(input(14))
	if err != nil {
		t.Fatal(err)
	}
	want := uint(3687786)
	got := Day14(lines, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	lines, err := linesFromFilename(input(14))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Day14(lines, false)
	}
}

func TestParseReactions(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
	}

	reactions, err := ParseReactions(input)
	if err != nil {
		t.Fatal(err)
	}

	// Check ORE reaction
	aReaction, ok := reactions["A"]
	if !ok {
		t.Fatal("expected reaction for A")
	}
	if aReaction.Output.Quantity != 10 {
		t.Errorf("expected output quantity 10, got %d", aReaction.Output.Quantity)
	}
	if len(aReaction.Inputs) != 1 {
		t.Errorf("expected 1 input, got %d", len(aReaction.Inputs))
	}
	if aReaction.Inputs[0].Name != "ORE" {
		t.Errorf("expected input ORE, got %s", aReaction.Inputs[0].Name)
	}

	// Check C reaction
	cReaction, ok := reactions["C"]
	if !ok {
		t.Fatal("expected reaction for C")
	}
	if len(cReaction.Inputs) != 2 {
		t.Errorf("expected 2 inputs, got %d", len(cReaction.Inputs))
	}
}

// Example usage for documentation
func ExampleDay14() {
	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}
	ore := Day14(input, true)
	fmt.Println(ore)
	// Output: 31
}
