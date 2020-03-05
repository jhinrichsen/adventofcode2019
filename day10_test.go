package adventofcode2019

import (
	"testing"
)

func TestExample1(t *testing.T) {
	example1 := []byte(`
	.#..#
	.....
	#####
	....#
	...##
	`)
	d := NewDay10(example1)

	// Check number of asteroids
	if len(d.asteroids) != 10 {
		t.Fatalf("want 10 but got %d", len(d.asteroids))
	}
	second := Asteroid{4, 0}
	if d.asteroids[1] != second {
		t.Fatalf("expected asteroid %+v at index 1, got %+v",
			second, d.asteroids[1])

	}
	a8 := Asteroid{3, 4}
	if d.asteroids[8] != a8 {
		t.Fatalf("expected asteroid %+v at index 1, got %+v",
			a8, d.asteroids[8])

	}
}
