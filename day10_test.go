package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestPolar(t *testing.T) {
	r, φ := Polar(1, 1)
	log.Printf("(%d,%d) -> r=%v,=%v\n", 1, 1, r, φ)
	r, φ = Polar(1, -1)
	r, φ = Polar(-1, -1)
	r, φ = Polar(-1, 1)
}

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
	second := Asteroid{4, 0}
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
	{day10Example1, Asteroid{3, 4}, 8},
	{day10Example2, Asteroid{5, 8}, 33},
	{day10Example3, Asteroid{1, 2}, 35},
	{day10Example4, Asteroid{6, 3}, 41},
	{day10Example5, Asteroid{11, 13}, 210},
}

func TestDay10Examples(t *testing.T) {
	for i, tt := range day10Examples {
		id := fmt.Sprintf("Day10Part1 example #%d", i+1)
		t.Run(id, func(t *testing.T) {
			d := NewDay10([]byte(tt.asteroidMap))
			gotA, got := d.Best()
			if tt.best != gotA {
				t.Fatalf("%d: want %+v but got %+v",
					tt.best.x, tt.best.y, gotA)
			}
			if tt.bestCount != got {
				t.Fatalf("want %d but got %d", tt.bestCount, got)
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
	_, got := d.Best()
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
		d.Best()
	}
}
