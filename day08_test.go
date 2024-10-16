package adventofcode2019

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestDay8Part1(t *testing.T) {
	digits, err := ioutil.ReadFile(input(8))
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
	digits, err := ioutil.ReadFile(input(8))
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

func day8Part2Result() ([]byte, error) {
	filename := "testdata/day08-part2-result.txt"
	return ioutil.ReadFile(filename)
}

func TestDay8Part2(t *testing.T) {
	digits, err := ioutil.ReadFile(input(8))
	if err != nil {
		t.Fatal(err)
	}
	want, err := day8Part2Result()
	if err != nil {
		t.Fatal(err)
	}

	got, err := Day8Part2(digits)
	if err != nil {
		t.Fatal(err)
	}
	// well, this is ASCII art, so in absence of a package that can parse it
	// we need human interaction
	var humanReadable bytes.Buffer
	for i := 0; i < len(got); i++ {
		if got[i] == '0' {
			fmt.Fprintf(&humanReadable, "%s", string("  "))
		} else if got[i] == '1' {
			fmt.Fprintf(&humanReadable, "%s", string("X "))
		} else {
			fmt.Fprintf(&humanReadable, "%s", string("? "))
		}
		if i%25 == 24 {
			fmt.Fprintln(&humanReadable)
		}
	}
	if !reflect.DeepEqual(want, humanReadable.Bytes()) {
		fmt.Fprintf(os.Stderr, "want: %v\n", want)
		fmt.Fprintf(os.Stderr, " got: %v\n", humanReadable.Bytes())
		t.Fatal("want does not equal got")
	}
}

func BenchmarkDay8Part2(b *testing.B) {
	digits, err := ioutil.ReadFile(input(8))
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
