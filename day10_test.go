package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

const (
	day10Example1 = `
	.#..#
	.....
	#####
	....#
	...##
	`

	day10Example2 = `
	......#.#.
	#..#.#....
	..#######.
	.#.#.###..
	.#..#.....
	..#....#.#
	#..#....#.
	.##.#..###
	##...#..#.
	.#....####
	`

	day10Example3 = `
	#.#...#.#.
	.###....#.
	.#....#...
	##.#.#.#.#
	....#.#.#.
	.##..###.#
	..#...##..
	..##....##
	......#...
	.####.###.
	`

	day10Example4 = `
	.#..#..###
	####.###.#
	....###.#.
	..###.##.#
	##.##.#.#.
	....###..#
	..#.#..#.#
	#..#.#.###
	.##...##.#
	.....#.#..
	`

	day10Example5 = `
	.#..##.###...#######
	##.############..##.
	.#.######.########.#
	.###.#######.####.#.
	#####.##.#.##.###.##
	..#####..#.#########
	####################
	#.####....###.#.#.##
	##.#################
	#####.##.###..####..
	..######..##.#######
	####.##.####...##..#
	.#####..#.######.###
	##...#.##########...
	#.##########.#######
	.####.#.###.###.#.##
	....##.##.###..#####
	.#.#.###########.###
	#.#.#.#####.####.###
	###.##.####.##.#..##
	`

	day10Part2Example1 = `
	.#....#####...#..
	##...##.#####..##
	##...#...#.#####.
	..#.....X...###..
	..#.#.....#....##
	`
)

var day10Examples = []struct {
	asteroidMap string
	best        Asteroid
	bestCount   int
}{
	{day10Example1, 3 + 4i, 8},
	{day10Example2, 5 + 8i, 33},
	{day10Example3, 1 + 2i, 35},
	{day10Example4, 6 + 3i, 41},
	{day10Example5, 11 + 13i, 210},
}

func TestDay10Example1(t *testing.T) {
	as := ParseAsteroidMap([]byte(day10Example1))

	// Check number of asteroids
	if len(as) != 10 {
		t.Fatalf("want 10 but got %d", len(as))
	}
	second := 4 + 0i
	if as[1] != second {
		t.Fatalf("expected asteroid %+v at index 1, got %+v",
			second, as[1])

	}
}

func TestDay10Part1Examples(t *testing.T) {
	for i, tt := range day10Examples {
		id := fmt.Sprintf("Day10Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			as := ParseAsteroidMap([]byte(tt.asteroidMap))
			wantA, want := tt.best, tt.bestCount
			gotA, got := Day10Part1(as)
			if tt.best != gotA {
				t.Fatalf("%s: want %+v but got %+v",
					id, wantA, gotA)
			}
			if want != got {
				t.Fatalf("want %d but got %d", want, got)
			}
		})
	}
}

func TestDay10Part1(t *testing.T) {
	filename := "testdata/day10.txt"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	as := ParseAsteroidMap(buf)
	if err != nil {
		t.Fatal(err)
	}
	want := 267
	_, got := Day10Part1(as)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	filename := "testdata/day10.txt"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		b.Fatal(err)
	}
	as := ParseAsteroidMap(buf)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Day10Part1(as)
	}
}

func TestDay10Vaporize(t *testing.T) {
	ex := day10Examples[4]
	as := ParseAsteroidMap([]byte(ex.asteroidMap))
	got := center(vaporize(byPhase(center(as, ex.best))), -ex.best)

	wants := make(map[int]Asteroid)
	// The 1st asteroid to be vaporized is at 11,12
	wants[1-1] = 11 + 12i
	wants[2-1] = 12 + 1i
	wants[3-1] = 12 + 2i
	wants[10-1] = 12 + 8i
	wants[20-1] = 16 + 0i
	wants[50-1] = 16 + 9i
	wants[100-1] = 10 + 16i
	wants[199-1] = 9 + 6i
	wants[200-1] = 8 + 2i
	wants[201-1] = 10 + 9i
	wants[299-1] = 11 + 1i
	for k := range wants {
		if wants[k] != got[k] {
			t.Fatalf("want asteroid[%d] == %v but got %v", k,
				wants[k], got[k])
		}
	}
	if got[10-1] != 12+8i {
		t.Fatalf("want %v but got %v", got[10-1], 12+8i)
	}
}

func TestDay10Part2Example1(t *testing.T) {
	want := []Asteroid{
		8 + 1i,
		9 + 0i,
		9 + 1i,
		10 + 0i,
		9 + 2i,
		11 + 1i,
		12 + 1i,
		11 + 2i,
		15 + 1i,
	}
	as := ParseAsteroidMap([]byte(day10Part2Example1))
	base := 8 + 3i
	as = center(as, base)
	pgs := byPhase(as)

	Δ := len(as) - countAsteroids(pgs)
	if Δ != 0 {
		t.Fatalf("byPhase() lost %d asteroids", Δ)
	}

	got := vaporize(pgs)
	got = center(got, -base)

	// make sure no asteroids got lost
	Δ = len(as) - len(got)
	if Δ != 0 {
		t.Fatalf("vaporize() lost %d asteroids", Δ)
	}

	// check the first N known vaporized planets
	if !reflect.DeepEqual(want, got[:len(want)]) {
		t.Fatalf("want %+v but got %+v", want, got[:len(want)])
	}
}

func TestDay10Part2Example2(t *testing.T) {
	want := 802
	ex := day10Examples[4]
	as := ParseAsteroidMap([]byte(ex.asteroidMap))
	got := Day10Part2(as, ex.best)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay10Part2(t *testing.T) {
	want := 1309
	filename := "testdata/day10.txt"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	as := ParseAsteroidMap(buf)
	best, _ := Day10Part1(as)
	got := Day10Part2(as, best)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
