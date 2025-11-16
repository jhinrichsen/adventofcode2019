package adventofcode2019

import (
	"testing"
)

func TestDay11Pbm(t *testing.T) {
	buf := file(t, 11)
	_ = newRegistrationID(MustSplit(string(buf)), colorBlack)
	// TODO apply OCR
	// fmt.Println("Day 11, part 1")
	// fmt.Println(string(got.pbm()))
}

func TestDay11Part1(t *testing.T) {
	// First try want := 9870
	// Second try want := 907
	want := 2343
	buf := file(t, 11)
	got := Day11Part1(MustSplit(string(buf)))
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// TestMapOfBools makes sure both true and false values are counted
func TestMapOfBools(t *testing.T) {
	m := make(map[int]bool)

	if len(m) != 0 {
		t.Fatalf("want empty map but got %d", len(m))
	}

	m[0] = false
	if len(m) != 1 {
		t.Fatalf("want 1 map entry but got %d", len(m))
	}
}

func TestDay11Part2(t *testing.T) {
	buf := file(t, 11)
	_ = Day11Part2(MustSplit(string(buf)))
	// TODO apply OCR
	// fmt.Println("Day 11, part 2:")
	// fmt.Println(string(got))
}

func BenchmarkDay11Part1(b *testing.B) {
	buf := file(b, 11)
	master := MustSplit(string(buf))
	for range b.N {
		_ = Day11Part1(master.Copy())
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	buf := file(b, 11)
	master := MustSplit(string(buf))
	for range b.N {
		_ = Day11Part2(master.Copy())
	}
}
