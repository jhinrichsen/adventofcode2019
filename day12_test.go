package adventofcode2019

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

const day12Example1Steps = 10

// TestDay12GanymedeCallisto tests if Ganymede has an x position of 3, and Callisto has
// a x position of 5, then Ganymede's x velocity changes by +1 (because 5 > 3)
// and Callisto's x velocity changes by -1 (because 3 < 5). However, if the
// positions on a given axis are the same, the velocity on that axis does not
// change for that pair of moons.
func TestDay12GanymedeCallisto(t *testing.T) {
	const (
		v        = 42
		ganymede = 0
		callisto = 1
	)
	var u universe

	u.moons[X][ganymede].pos = 3
	u.moons[Y][ganymede].pos = v
	u.moons[Z][ganymede].pos = 0

	u.moons[X][callisto].pos = 5
	u.moons[Y][callisto].pos = v
	u.moons[Z][callisto].pos = 0
	for dim := 0; dim < DIMS; dim++ {
		u.gravity(dim, ganymede, callisto)
	}
	w1 := u.moons[X][ganymede].vel
	const g1 = 1
	if w1 != g1 {
		t.Fatalf("ganymede's X velocity: want %d but got %d", w1, g1)
	}

	if u.moons[X][callisto].vel != -1 {
		t.Fatalf("callisto's X velocity: want %d but got %d",
			-1, u.moons[X][callisto].vel)
	}
	if u.moons[Y][ganymede].vel != 0 {
		t.Fatalf("ganymede's Y velocity: want %d but got %d",
			0, u.moons[Y][ganymede].vel)
	}
	if u.moons[Y][callisto].vel != 0 {
		t.Fatalf("callisto's Y velocity: want %d but got %d",
			0, u.moons[Y][callisto].vel)
	}
}

// TestDay12Europa simply add the velocity of each moon to its own position. For
// example, if Europa has a position of x=1, y=2, z=3 and a velocity of x=-2,
// y=0,z=3, then its new position would be x=-1, y=2, z=6. This process does not
// modify the velocity of any moon.
func TestDay12Europa(t *testing.T) {
	want := [DIMS]point{
		{-1, -2},
		{2, 0},
		{6, 3},
	}
	got := [DIMS]point{
		{1, -2},
		{2, 0},
		{3, 3},
	}
	for i := 0; i < DIMS; i++ {
		got[i].velocity()
	}
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}
}

func day12Example1Universe() (universe, error) {
	return day12FromFile(TestdataFilename("day12_example1_input.txt"))
}

func day12FromFile(filename string) (universe, error) {
	input, err := linesFromFilename(filename)
	if err != nil {
		return universe{}, err
	}
	return parseUsingParser(input)
}

func TestDay12Example1Timeline(t *testing.T) {
	wantLines, err := linesFromFilename(TestdataFilename("day12_example1_output.txt"))
	if err != nil {
		t.Fatal(err)
	}
	u, err := day12Example1Universe()
	if err != nil {
		t.Fatal(err)
	}
	var sb strings.Builder
	// one more because we print then step
	n := day12Example1Steps + 1
	for i := 0; i < n; i++ {
		fmt.Fprintf(&sb, "After %d steps:\n", i)
		sb.WriteString(u.String())
		// separating line except for the last
		if i+1 < n {
			sb.WriteString("\n")
		}
		for dim := 0; dim < DIMS; dim++ {
			u.step(dim)
		}
	}
	// fmt.Fprintf(os.Stdout, "%s", sb.String())
	gotLines, err := linesFromReader(strings.NewReader(sb.String()))
	if err != nil {
		t.Fatal(err)
	}

	// compare output lines
	if len(wantLines) != len(gotLines) {
		t.Fatalf("number of lines: want %d but got %d",
			len(wantLines), len(gotLines))
	}
	trim := func(s string) string {
		return strings.ReplaceAll(s, " ", "")
	}
	for i := 0; i < len(wantLines); i++ {
		// original output has variable length output, which we ignore
		w := trim(wantLines[i])
		g := trim(gotLines[i])
		if w != g {
			t.Fatalf("line %d: want %q but got %q",
				i, w, g)

		}
	}
}

// energy returns the total energy of a universe identified by its moon
// positions from file and the number of steps in a timeline.
func EnergyFromFile(filename string, steps int) (int, error) {
	input, err := linesFromFilename(TestdataFilename(filename))
	if err != nil {
		return 0, err
	}
	u, err := parseUsingParser(input)
	if err != nil {
		return 0, err
	}
	for ; steps > 0; steps-- {
		for dim := 0; dim < DIMS; dim++ {
			u.step(dim)
		}
	}
	return u.energy(), nil
}

func TestDay12Example1Energy(t *testing.T) {
	const want = 179
	got, err := EnergyFromFile("day12_example1_input.txt", day12Example1Steps)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Example2(t *testing.T) {
	const want = 1940
	got, err := EnergyFromFile("day12_example2_input.txt", 100)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 7471
	got, err := EnergyFromFile("day12.txt", 1000)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Parse(t *testing.T) {
	lines, err := InputLinesForDay(12)
	if err != nil {
		t.Fatal(err)
	}
	u, err := parseUsingParser(lines)
	if err != nil {
		t.Fatal(err)
	}
	want := 2
	got := u.moons[2][3].pos
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func parseUsingParser(lines []string) (universe, error) {
	var u universe
	isNumeric := func(c byte) bool {
		return c == '-' || c >= '0' && c <= '9'
	}
	for j, line := range lines {
		buf := []byte(line)
		idx := 0
		l := len(line)
		nextNum := func() (int, error) {
			// skip any non-numeric characters
			for !isNumeric(buf[idx]) && idx < l {
				idx++
			}
			from := idx
			for isNumeric(buf[idx]) && idx < l {
				idx++
			}
			n, err := strconv.Atoi(string(buf[from:idx]))
			return int(n), err
		}
		for i := 0; i < DIMS; i++ {
			var err error
			u.moons[i][j].pos, err = nextNum()
			if err != nil {
				return u, err
			}
		}
	}
	return u, nil
}

func BenchmarkParseUsingParser(b *testing.B) {
	lines, err := InputLinesForDay(12)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		parseUsingParser(lines)
	}
}

func BenchmarkDay12Example2(b *testing.B) {
	input, err := linesFromFilename(TestdataFilename("day12.txt"))
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		u, err := parseUsingParser(input)
		if err != nil {
			b.Fatal(err)
		}
		for j := 0; j < 1000; j++ {
			for dim := 0; dim < DIMS; dim++ {
				u.step(dim)
			}
		}
		u.energy()
	}
}

func TestDay12Part2Example1(t *testing.T) {
	want := 2772
	u, err := day12FromFile(TestdataFilename("day12_example1_input.txt"))
	if err != nil {
		t.Fatal(err)
	}
	got := u.cycle()
	if want != got {
		t.Fatalf("want %d but got %d, factor %f", want, got,
			float64(got)/float64(want))
	}
}

func TestDay12Part2Example2(t *testing.T) {
	want := 4686774924
	u, err := day12FromFile(TestdataFilename("day12_example2_input.txt"))
	if err != nil {
		t.Fatal(err)
	}
	got := u.cycle()
	if want != got {
		t.Fatalf("want %d but got %d, factor %f", want, got,
			float64(got)/float64(want))
	}
}

func TestDay12Part2(t *testing.T) {
	want := 376243355967784
	u, err := day12FromFile(TestdataFilename("day12.txt"))
	if err != nil {
		t.Fatal(err)
	}
	got := u.cycle()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Sizeof(t *testing.T) {
	var u universe
	n := unsafe.Sizeof(u)
	log.Printf("sizeof universe: %d\n", n)
}
