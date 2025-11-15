package adventofcode2019

import (
	"strings"
	"testing"
)

func TestDay23Part1(t *testing.T) {
	lines := testLinesFromFilename(t, filename(23))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	got := Day23(prog, true)
	want := uint(19530)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay23Part2(t *testing.T) {
	lines := testLinesFromFilename(t, filename(23))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	got := Day23(prog, false)
	want := uint(12725)
	if got != want {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := testLinesFromFilename(b, filename(23))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	for b.Loop() {
		Day23(prog, true)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := testLinesFromFilename(b, filename(23))
	prog := MustSplit(strings.TrimSpace(lines[0]))
	for b.Loop() {
		Day23(prog, false)
	}
}
