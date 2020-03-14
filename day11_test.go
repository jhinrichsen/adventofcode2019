package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDay11Pbm(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := newRegistrationID(MustSplit(string(buf)), colorBlack)
	fmt.Println("Day 11, part 1")
	fmt.Println(string(got.pbm()))
}

func TestDay11Part1(t *testing.T) {
	// First try want := 9870
	// Second try want := 907
	want := 2343
	buf, err := ioutil.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
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
	buf, err := ioutil.ReadFile("testdata/day11.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day11Part2(MustSplit(string(buf)))
	fmt.Println("Day 11, part 2:")
	fmt.Println(string(got))
}
