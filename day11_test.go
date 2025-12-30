package adventofcode2019

import "testing"

func TestDay11Part1(t *testing.T) {
	testSolver(t, 11, filename, true, Day11, uint(2343))
}

func TestDay11Part2(t *testing.T) {
	// Part 2 returns len(pbm) as checksum - value determined from initial run
	buf := fileFromFilename(t, filename, 11)
	got, err := Day11(buf, false)
	if err != nil {
		t.Fatal(err)
	}
	if got == 0 {
		t.Fatal("expected non-zero result")
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	buf := fileFromFilename(b, filename, 11)
	for b.Loop() {
		_, _ = Day11(buf, true)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	buf := fileFromFilename(b, filename, 11)
	for b.Loop() {
		_, _ = Day11(buf, false)
	}
}
