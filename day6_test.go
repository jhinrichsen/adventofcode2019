package adventofcode2019

import "testing"

func day6FromFile(t *testing.T, filename string) Day6 {
	ss, err := Lines(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay6(ss)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func example1(t *testing.T) Day6 {
	return day6FromFile(t, "testdata/day6_example.txt")
}

func example2(t *testing.T) Day6 {
	return day6FromFile(t, "testdata/day6_example2.txt")
}

func TestDay6OrbitC(t *testing.T) {
	want := "B"
	d := example1(t)
	got := d.Orbit("C")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay6OrbitCountCOM(t *testing.T) {
	d := example1(t)
	want := 0
	got := d.OrbitCount(COM)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountD(t *testing.T) {
	d := example1(t)
	want := 3
	got := d.OrbitCount("D")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountL(t *testing.T) {
	d := example1(t)
	want := 7
	got := d.OrbitCount("L")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Example(t *testing.T) {
	d := example1(t)
	want := 42
	got := d.OrbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part1(t *testing.T) {
	d := day6FromFile(t, "testdata/day6.txt")
	want := 142497
	got := d.OrbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6CommonOrbit(t *testing.T) {
	d := example2(t)
	want := "D"
	got := d.CommonOrbit("YOU", "SAN")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func TestDay6Part2Example(t *testing.T) {
	d := day6FromFile(t, "testdata/day6_example2.txt")
	want := 4
	got := d.Transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part2(t *testing.T) {
	d := day6FromFile(t, "testdata/day6.txt")
	want := 301
	got := d.Transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
