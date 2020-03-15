package adventofcode2019

import "testing"

func day6FromFile(filename string) (Day6, error) {
	ss, err := InputLines(6)
	if err != nil {
		return Day6{}, err
	}
	d, err := NewDay6(ss)
	if err != nil {
		return d, err
	}
	return d, nil
}

func example1() (Day6, error) {
	return day6FromFile("testdata/day6_example.txt")
}

func example1b(b *testing.B) Day6 {
	d, err := example1()
	if err != nil {
		b.Fatal(err)
	}
	return d
}

func example1t(t *testing.T) Day6 {
	d, err := example1()
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func example2() (Day6, error) {
	return day6FromFile("testdata/day6_example2.txt")
}

func example2t(t *testing.T) Day6 {
	d, err := example2()
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func example2b(b *testing.B) Day6 {
	d, err := example2()
	if err != nil {
		b.Fatal(err)
	}
	return d
}

func TestDay6OrbitC(t *testing.T) {
	d := example1t(t)
	want := "B"
	got := d.Orbit("C")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay6OrbitC(b *testing.B) {
	d := example1b(b)
	for i := 0; i < b.N; i++ {
		d.Orbit("C")
	}
}

func TestDay6OrbitCountCOM(t *testing.T) {
	d := example1t(t)
	want := 0
	got := d.OrbitCount(COM)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountD(t *testing.T) {
	d := example1t(t)
	want := 3
	got := d.OrbitCount("D")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6OrbitCountL(t *testing.T) {
	d := example1t(t)
	want := 7
	got := d.OrbitCount("L")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Example(t *testing.T) {
	d := example1t(t)
	want := 42
	got := d.OrbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part1(t *testing.T) {
	d, err := day6FromFile("testdata/day6.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 142497
	got := d.OrbitCountChecksum()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay6Part1(b *testing.B) {
	d, err := day6FromFile("testdata/day6.txt")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		d.OrbitCountChecksum()
	}
}

func TestDay6CommonOrbit(t *testing.T) {
	d := example2t(t)
	want := "D"
	got := d.CommonOrbit("YOU", "SAN")
	if want != got {
		t.Fatalf("want %q but got %q", want, got)
	}
}

func BenchmarkDay6CommonOrbit(b *testing.B) {
	d := example2b(b)
	for i := 0; i < b.N; i++ {
		d.CommonOrbit("YOU", "SAN")
	}
}

func TestDay6Part2Example(t *testing.T) {
	d := example2t(t)
	want := 4
	got := d.Transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part2(t *testing.T) {
	d, err := day6FromFile("testdata/day6.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 301
	got := d.Transfers("YOU", "SAN")
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay6Part2(b *testing.B) {
	d, err := day6FromFile("testdata/day6.txt")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		d.Transfers("YOU", "SAN")
	}
}
