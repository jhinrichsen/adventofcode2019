package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay02Part1Examples(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		id := fmt.Sprintf("Example(%s)", tt.in)
		t.Run(id, func(t *testing.T) {
			ic, err := NewIntcode([]byte(tt.in))
			if err != nil {
				t.Fatal(err)
			}
			ic.Run()
			for i, want := range tt.out {
				if got := ic.Mem(i); got != want {
					t.Fatalf("%s: mem[%d] want %d but got %d", id, i, want, got)
				}
			}
		})
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
