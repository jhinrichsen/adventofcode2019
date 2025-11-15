package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay14Part1Example1(t *testing.T) {
	testLines(t, 14, example1Filename, true, Day14, 31)
}

func TestDay14Part1Example2(t *testing.T) {
	testLines(t, 14, example2Filename, true, Day14, 165)
}

func TestDay14Part1Example3(t *testing.T) {
	testLines(t, 14, example3Filename, true, Day14, 13312)
}

func TestDay14Part1Example4(t *testing.T) {
	testLines(t, 14, example4Filename, true, Day14, 13312)
}

func TestDay14Part1Example5(t *testing.T) {
	testLines(t, 14, example5Filename, true, Day14, 180697)
}

func TestDay14Part1Example6(t *testing.T) {
	testLines(t, 14, example6Filename, true, Day14, 2210736)
}

func TestDay14Part1(t *testing.T) {
	testLines(t, 14, filename, true, Day14, 337862)
}

func BenchmarkDay14Part1(b *testing.B) {
	benchLines(b, 14, true, Day14)
}

func TestDay14Part2Example3(t *testing.T) {
	testLines(t, 14, example3Filename, false, Day14, 82892753)
}

func TestDay14Part2Example4(t *testing.T) {
	testLines(t, 14, example4Filename, false, Day14, 82892753)
}

func TestDay14Part2Example5(t *testing.T) {
	testLines(t, 14, example5Filename, false, Day14, 5586022)
}

func TestDay14Part2Example6(t *testing.T) {
	testLines(t, 14, example6Filename, false, Day14, 460664)
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
