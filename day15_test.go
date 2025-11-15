package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay15Part1(t *testing.T) {
	prog, err := os.ReadFile(filename(15))
	if err != nil {
		t.Fatal(err)
	}
	const want = 272
	got := Day15(prog, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay15Part2(t *testing.T) {
	prog, err := os.ReadFile(filename(15))
	if err != nil {
		t.Fatal(err)
	}
	const want = 398
	got := Day15(prog, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	prog, err := os.ReadFile(filename(15))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		Day15(prog, true)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	prog, err := os.ReadFile(filename(15))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		Day15(prog, false)
	}
}
