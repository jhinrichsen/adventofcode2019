package adventofcode2019

import (
	"io/ioutil"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	filename := "testdata/day8.txt"
	digits, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := 1463
	got, err := Day8Part1(digits)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
