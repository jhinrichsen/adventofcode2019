package adventofcode2019

import (
	"fmt"
	"io/ioutil"
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
)

func TestExample1(t *testing.T) {
	d := NewDay10([]byte(day10Example1))

	// Check number of asteroids
	if len(d.asteroids) != 10 {
		t.Fatalf("want 10 but got %d", len(d.asteroids))
	}
	second := 4 + 0i
	if d.asteroids[1] != second {
		t.Fatalf("expected asteroid %+v at index 1, got %+v",
			second, d.asteroids[1])

	}
}

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

func TestDay10Examples(t *testing.T) {
	for i, tt := range day10Examples {
		id := fmt.Sprintf("Day10Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			d := NewDay10([]byte(tt.asteroidMap))
			wantA, want := tt.best, tt.bestCount
			gotA, got := d.Part1()
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
	d := NewDay10(buf)
	if err != nil {
		t.Fatal(err)
	}
	want := 267
	_, got := d.Part1()
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
	d := NewDay10(buf)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		d.Part1()
	}
}
