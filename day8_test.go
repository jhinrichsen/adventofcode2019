package adventofcode2019

import (
	"fmt"
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

func BenchmarkDay8Part1(b *testing.B) {
	filename := "testdata/day8.txt"
	digits, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := Day8Part1(digits)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestDay8Part2(t *testing.T) {
	filename := "testdata/day8.txt"
	digits, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	_ = "GKCKH"
	// well, this is ASCII art, so in absence of a package that can parse it
	// we need human interaction
	got, err := Day8Part2(digits)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(got); i++ {
		if got[i] == '0' {
			fmt.Printf("%s", string("  "))
		} else if got[i] == '1' {
			fmt.Printf("%s", string("X "))
		} else {
			fmt.Printf("%s", string("? "))
			// t.Fatalf("bad color at index %d: %d", i, got[i])
		}
		if i%25 == 24 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func BenchmarkDay8Part2(b *testing.B) {
	filename := "testdata/day8.txt"
	digits, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := Day8Part2(digits)
		if err != nil {
			b.Fatal(err)
		}
	}
}
