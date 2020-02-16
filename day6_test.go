package adventofcode2019

import "testing"

func example(t *testing.T) Day6 {
	filename := "testdata/day6_example.txt"
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

func TestDay6OrbitCountCOM(t *testing.T) {
	d := example(t)
	want := 0
	got := d.OrbitCount(COM)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountD(t *testing.T) {
	d := example(t)
	want := 3
	got := d.OrbitCount("D")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountL(t *testing.T) {
	d := example(t)
	want := 7
	got := d.OrbitCount("L")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Example(t *testing.T) {
	d := example(t)
	want := 42
	got := d.OrbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part1(t *testing.T) {
	filename := "testdata/day6.txt"
	ss, err := Lines(filename)
	if err != nil {
		t.Fatal(err)
	}
	d, err := NewDay6(ss)
	if err != nil {
		t.Fatal(err)
	}
	want := 142497
	got := d.OrbitCountChecksum()

	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
