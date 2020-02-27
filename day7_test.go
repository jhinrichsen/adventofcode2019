package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	prog1 = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	prog2 = "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23," +
		"23,4,23,99,0,0"
	prog3 = "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002," +
		"33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
)

var day7Part1Examples = []struct {
	phases string
	prog   string
	want   int
}{
	{"43210", prog1, 43210},
	{"01234", prog2, 54321},
	{"10432", prog3, 65210},
}

func TestDay7Part1Examples(t *testing.T) {
	for _, tt := range day7Part1Examples {
		id := fmt.Sprintf("Day7(%s, %s)", "prog1", tt.phases)
		t.Run(id, func(t *testing.T) {
			got := Day7Part1(MustSplit(tt.prog), tt.phases)
			if tt.want != got {
				t.Fatalf("%s: want %d but got %d", id,
					tt.want, got)
			}
		})
	}
}

func TestFac(t *testing.T) {
	want := 120
	got := fac(5)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay7Part1(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day7.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 24405
	got := Day7Part1(MustSplit(string(buf)), "01234")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay7Part1(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/day7.txt")
	if err != nil {
		b.Fatal(err)
	}
	prog := MustSplit(string(buf))
	for i := 0; i < b.N; i++ {
		Day7Part1(prog, "01234")
	}
}
