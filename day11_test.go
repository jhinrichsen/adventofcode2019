package adventofcode2019

import (
	"io/ioutil"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
	// First try want := 9870
	// Second try want := 907
	want := 2343
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
