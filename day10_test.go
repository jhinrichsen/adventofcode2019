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
	d:= NewDay10(example1)

	// Check number of asteroids
	if len(d.asteroids) != 10 {
		t.Fatalf("want 10 but got %d", len(d.asteroids))
	}
}
