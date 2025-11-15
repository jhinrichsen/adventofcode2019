package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay17Part1(t *testing.T) {
	prog, err := os.ReadFile(filename(17))
	if err != nil {
		t.Fatal(err)
	}
	const want = 5972
	got := Day17(prog, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay17Part1(b *testing.B) {
	prog, err := os.ReadFile(filename(17))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		Day17(prog, true)
	}
}
