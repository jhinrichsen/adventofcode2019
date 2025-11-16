package adventofcode2019

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDay02Part1Examples(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}
	for _, tt := range tests {
		id := fmt.Sprintf("Runs(%s)", tt.in)
		t.Run(id, func(t *testing.T) {
			want := tt.out
			got, err := Runs(tt.in)
			if err != nil {
				t.Fatal(err)
			}
			if want != got {
				t.Fatalf("%s: want %s but got %s", id,
					want, got)
			}
		})
	}
}

func TestDay02Len(t *testing.T) {
	want := 99
	opcodes, err := Split("1,0,0,0,99")
	if err != nil {
		t.Fatal(err)
	}
	got := Len(opcodes)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Split(t *testing.T) {
	want := []int{1, 0, 0, 0, 99}
	got, err := Split("1,0,0,0,99")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay02Part1(t *testing.T) {
	testSolver(t, 2, filename, true, Day02, uint(3562624))
}

func TestDay02Part2(t *testing.T) {
	testSolver(t, 2, filename, false, Day02, uint(8298))
}

func BenchmarkDay02Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 2)
	for b.Loop() {
		_, _ = Day02(buf, true)
	}
}

func BenchmarkDay02Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 2)
	for b.Loop() {
		_, _ = Day02(buf, false)
	}
}
