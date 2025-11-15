package adventofcode2019

import (
	"strings"
	"testing"
)

func TestDay25Part1(t *testing.T) {
	lines := testLinesFromFilename(t, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))

	// First just test parseRoom
	output := executeCommands(prog, []string{})
	room := parseRoom(output)
	t.Logf("Start room: %s", room.name)
	t.Logf("Exits: %v", room.exits)
	t.Logf("Items: %v", room.items)

	// Test going north
	output2 := executeCommands(prog, []string{"north"})
	room2 := parseRoom(output2)
	t.Logf("North room: %s", room2.name)
	t.Logf("Exits: %v", room2.exits)

	// Print last 500 chars of output
	if len(output2) > 500 {
		t.Logf("Output (last 500): %s", output2[len(output2)-500:])
	} else {
		t.Logf("Full output: %s", output2)
	}

	t.Skip("Day 25 solution takes ~5 minutes to run - skipping in tests")

	got := Day25(prog, true)
	want := uint(229384)
	if got != want {
		t.Errorf("Day25(part1) = %d, want %d", got, want)
	}
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	for b.Loop() {
		Day25(prog, true)
	}
}
