package adventofcode2019

import (
	"io/ioutil"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	want := 315
	prog, err := ioutil.ReadFile("testdata/day13.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day13Part1(Day5, MustSplit(string(prog)))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
