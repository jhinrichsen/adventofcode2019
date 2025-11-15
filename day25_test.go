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

	got := Day25(prog, true)
	want := uint(0) // Update after running
	if want != 0 && got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
	if got == 0 {
		t.Fatal("No password found")
	}
	t.Logf("Password: %d", got)
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(25))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	for b.Loop() {
		Day25(prog, true)
	}
}
