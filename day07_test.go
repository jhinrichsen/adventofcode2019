package adventofcode2019

import (
	"fmt"
	"os"
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
	for i, tt := range day7Part1Examples {
		s := fmt.Sprintf("prog%d", i)
		id := fmt.Sprintf("Day7Part1(%s, %s)", s, tt.phases)
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
	buf, err := os.ReadFile(input(7))
	if err != nil {
		t.Fatal(err)
	}
	want := 24405
	got := Day7Part1(MustSplit(string(buf)), "01234")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	buf := fileFromFilename(t, filename, 7)
	want := uint(24405)
	got := Day07(buf, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 7)
	for b.Loop() {
		_ = Day07(buf, true)
	}
}

const (
	prog4 = "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001," +
		"28,-1,28,1005,28,6,99,0,0,5"
	prog5 = "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26," +
		"1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1," +
		"55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
)

var day7Part2Examples = []struct {
	phases string
	prog   string
	want   int
}{
	{"98765", prog4, 139629729},
	{"97856", prog5, 18216},
}

func TestDay7Part2Examples(t *testing.T) {
	for i, tt := range day7Part2Examples {
		s := fmt.Sprintf("prog%d", i)
		id := fmt.Sprintf("Day7Part2(%s, %s)", s, tt.phases)
		t.Run(id, func(t *testing.T) {
			got := Day7Part2(MustSplit(tt.prog), tt.phases)
			if tt.want != got {
				t.Fatalf("%s: want %d but got %d", id,
					tt.want, got)
			}
		})
	}
}

func TestDay7Part2(t *testing.T) {
	buf, err := os.ReadFile(input(7))
	if err != nil {
		t.Fatal(err)
	}
	want := 8271623
	got := Day7Part2(MustSplit(string(buf)), "56789")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	buf := fileFromFilename(t, filename, 7)
	want := uint(8271623)
	got := Day07(buf, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 7)
	for b.Loop() {
		_ = Day07(buf, false)
	}
}
