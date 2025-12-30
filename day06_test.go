package adventofcode2019

import "testing"

func day06Puzzle(tb testing.TB, filename string) Day06Puzzle {
	tb.Helper()
	lines := testLinesFromFilename(tb, filename)
	puzzle, err := NewDay06(lines)
	if err != nil {
		tb.Fatal(err)
	}
	return puzzle
}

func TestDay6OrbitC(t *testing.T) {
	d := day06Puzzle(t, exampleFilename(6))
	want := "B"
	got := d.orbit("C")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay6OrbitC(b *testing.B) {
	d := day06Puzzle(b, exampleFilename(6))
	for b.Loop() {
		d.orbit("C")
	}
}

func TestDay6OrbitCountCOM(t *testing.T) {
	d := day06Puzzle(t, exampleFilename(6))
	want := 0
	got := d.orbitCount(com)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountD(t *testing.T) {
	d := day06Puzzle(t, exampleFilename(6))
	want := 3
	got := d.orbitCount("D")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountL(t *testing.T) {
	d := day06Puzzle(t, exampleFilename(6))
	want := 7
	got := d.orbitCount("L")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Example(t *testing.T) {
	d := day06Puzzle(t, exampleFilename(6))
	want := 42
	got := d.orbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part1(t *testing.T) {
	d := day06Puzzle(t, filename(6))
	want := 142497
	got := d.orbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part1(t *testing.T) {
	testWithParser(t, 6, filename, true, NewDay06, Day06, uint(142497))
}

func BenchmarkDay06Part1(b *testing.B) {
	benchWithParser(b, 6, true, NewDay06, Day06)
}

func TestDay6CommonOrbit(t *testing.T) {
	d := day06Puzzle(t, example2Filename(6))
	want := "D"
	got := d.commonOrbit("YOU", "SAN")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay6CommonOrbit(b *testing.B) {
	d := day06Puzzle(b, example2Filename(6))
	for b.Loop() {
		d.commonOrbit("YOU", "SAN")
	}
}

func TestDay6Part2Example(t *testing.T) {
	d := day06Puzzle(t, example2Filename(6))
	want := 4
	got := d.transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part2(t *testing.T) {
	d := day06Puzzle(t, filename(6))
	want := 301
	got := d.transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay06Part2(t *testing.T) {
	testWithParser(t, 6, filename, false, NewDay06, Day06, uint(301))
}

func BenchmarkDay06Part2(b *testing.B) {
	benchWithParser(b, 6, false, NewDay06, Day06)
}
