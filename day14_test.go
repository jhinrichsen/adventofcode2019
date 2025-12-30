package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay14Part1Examples(t *testing.T) {
	tests := []struct {
		filenameFunc func(uint8) string
		want         uint
	}{
		{example1Filename, 31},
		{example2Filename, 165},
		{example3Filename, 13312},
		{example4Filename, 13312},
		{example5Filename, 180697},
		{example6Filename, 2210736},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			testLines(t, 14, tt.filenameFunc, true, Day14, tt.want)
		})
	}
}

func TestDay14Part1(t *testing.T) {
	testLines(t, 14, filename, true, Day14, 337862)
}

func BenchmarkDay14Part1(b *testing.B) {
	benchLines(b, 14, true, Day14)
}

func TestDay14Part2Examples(t *testing.T) {
	tests := []struct {
		filenameFunc func(uint8) string
		want         uint
	}{
		{example3Filename, 82892753},
		{example4Filename, 82892753},
		{example5Filename, 5586022},
		{example6Filename, 460664},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("example%d", i+1), func(t *testing.T) {
			testLines(t, 14, tt.filenameFunc, false, Day14, tt.want)
		})
	}
}

func TestDay14Part2(t *testing.T) {
	testLines(t, 14, filename, false, Day14, 3687786)
}

func BenchmarkDay14Part2(b *testing.B) {
	benchLines(b, 14, false, Day14)
}

func TestParseReactions(t *testing.T) {
	input := []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
	}

	reactions, err := parseReactions(input)
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
