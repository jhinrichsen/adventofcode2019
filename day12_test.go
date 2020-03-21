package adventofcode2019

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

const day12Example1Steps = 10

// TestDay12GanymedeCallisto tests if Ganymede has an x position of 3, and Callisto has
// a x position of 5, then Ganymede's x velocity changes by +1 (because 5 > 3)
// and Callisto's x velocity changes by -1 (because 3 < 5). However, if the
// positions on a given axis are the same, the velocity on that axis does not
// change for that pair of moons.
func TestDay12GanymedeCallisto(t *testing.T) {
	v := 42
	ganymede := moon{
		pos: D3{3, v, 0},
	}
	callisto := moon{
		pos: D3{5, ganymede.pos[Y], 0},
	}
	ganymede.gravity(&callisto)
	if ganymede.vel[X] != 1 {
		t.Fatalf("ganymede's X velocity: want %d but got %d",
			0, ganymede.vel[X])
	}
	if callisto.vel[X] != -1 {
		t.Fatalf("callisto's X position: want %d but got %d",
			-1, callisto.vel[X])
	}

	if ganymede.vel[Y] != 0 {
		t.Fatalf("ganymede's Y velocity: want %d but got %d",
			0, ganymede.vel[Y])
	}
	if callisto.vel[Y] != 0 {
		t.Fatalf("callisto's Y velocity: want %d but got %d",
			0, callisto.vel[X])
	}
}

// TestDay12Europa simply add the velocity of each moon to its own position. For
// example, if Europa has a position of x=1, y=2, z=3 and a velocity of x=-2,
// y=0,z=3, then its new position would be x=-1, y=2, z=6. This process does not
// modify the velocity of any moon.
func TestDay12Europa(t *testing.T) {
	want := D3{-1, 2, 6}
	europa := moon{
		pos: D3{1, 2, 3},
		vel: D3{-2, 0, 3},
	}
	europa.velocity()
	got := europa.pos
	if want != got {
		t.Fatalf("want %+v but got %+v", want, got)
	}

}

func day12Example1Universe() (universe, error) {
	return day12FromFile(TestdataFilename("day12_example1_input.txt"))
}

func day12FromFile(filename string) (universe, error) {
	input, err := InputLines(filename)
	if err != nil {
		return universe{}, err
	}
	positions, err := ParseUsingParser(input)
	if err != nil {
		return universe{}, err
	}
	return newUniverse(positions), nil
}

func TestDay12Example1Timeline(t *testing.T) {
	wantLines, err := InputLines(TestdataFilename("day12_example1_output.txt"))
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
		for j := 0; j < len(u.moons); j++ {
			m := u.moons[j]
			sb.WriteString(m.String())
			sb.WriteString("\n")
		}
		// separating line except for the last
		if i+1 < n {
			sb.WriteString("\n")
		}
		u.step()
	}
	// fmt.Fprintf(os.Stdout, "%s", sb.String())
	gotLines, err := Lines(strings.NewReader(sb.String()))
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
	input, err := InputLines(TestdataFilename(filename))
	if err != nil {
		return 0, err
	}
	positions, err := ParseUsingParser(input)
	if err != nil {
		return 0, err
	}
	u := newUniverse(positions)
	for i := 0; i < steps; i++ {
		u.step()
	}
	return u.energy(), nil
}

func TestDay12Example1Energy(t *testing.T) {
	want := 179
	got, err := EnergyFromFile("day12_example1_input.txt", day12Example1Steps)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Example2(t *testing.T) {
	want := 1940
	got, err := EnergyFromFile("day12_example2.txt", 100)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	want := 7471
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
	ps, err := ParseUsingParser(lines)
	if err != nil {
		t.Fatal(err)
	}
	newUniverse(ps)
}

func ParseUsingRegexp(lines []string) ([]D3, error) {
	var ps []D3
	r := regexp.MustCompile(`^<x=([\d-]+), y=([\d-]+), z=([\d-]+)>$`)
	for _, line := range lines {
		ss := r.FindAllStringSubmatch(line, 3)
		var nums [3]int
		var err error
		for i := 0; i < 3; i++ {
			nums[i], err = strconv.Atoi(ss[0][i+1])
			if err != nil {
				return ps, err
			}
		}
		ps = append(ps, D3{nums[0], nums[1], nums[2]})
	}
	return ps, nil
}

func ParseUsingParser(lines []string) ([]D3, error) {
	var ps []D3
	isNumeric := func(c byte) bool {
		return c == '-' || c >= '0' && c <= '9'
	}
	for _, line := range lines {
		buf := []byte(line)
		var nums D3
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
			return n, err
		}
		for num := 0; num < len(nums); num++ {
			var err error
			nums[num], err = nextNum()
			if err != nil {
				return ps, err
			}
		}
		ps = append(ps, nums)
	}
	return ps, nil
}

func BenchmarkParseUsingRegexp(b *testing.B) {
	lines, err := InputLinesForDay(12)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		ParseUsingRegexp(lines)
	}
}

func BenchmarkParseUsingParser(b *testing.B) {
	lines, err := InputLinesForDay(12)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		ParseUsingParser(lines)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	input, err := InputLines(TestdataFilename("day12.txt"))
	if err != nil {
		b.Fatal(err)
	}
	positions, err := ParseUsingParser(input)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		u := newUniverse(positions)
		for j := 0; j < 1000; j++ {
			u.step()
		}
		u.energy()
	}
}

/*
func TestDay12Part2Example(t *testing.T) {
	u, err := day12FromFile(TestdataFilename("day12_example1_input.txt"))
	if err != nil {
		t.Fatal(err)
	}
	states := make(map[universe]bool)
}
*/
