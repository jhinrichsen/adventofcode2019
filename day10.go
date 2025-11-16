package adventofcode2019

import (
	"image"
	"math"
	"sort"
)

// Asteroid is a type alias for image.Point, representing a 2D position.
type Asteroid = image.Point

// direction represents a normalized direction vector for line-of-sight.
type direction struct {
	dx, dy int
}

// normalize returns the normalized direction from one asteroid to another.
func normalize(dx, dy int) direction {
	if dx == 0 && dy == 0 {
		return direction{0, 0}
	}
	// Calculate GCD inline to avoid naming conflict with day12
	a, b := dx, dy
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	g := a
	if g == 0 {
		return direction{dx, dy}
	}
	return direction{dx / g, dy / g}
}

// ParseAsteroidMap parses newline separated strings into asteroids.
func ParseAsteroidMap(asteroids []byte) []Asteroid {
	isAsteroid := func(b byte) bool {
		return b == '#'
	}
	isEmpty := func(b byte) bool {
		return b == '.'
	}
	isNewline := func(b byte) bool {
		return b == '\n'
	}
	isWhitespace := func(b byte) bool {
		return !(isAsteroid(b) || isEmpty(b))
	}

	// overread any leading whitespace
	start := 0
	for start < len(asteroids) && isWhitespace(asteroids[start]) {
		start++
	}

	// Count asteroids first to preallocate
	count := 0
	for _, b := range asteroids[start:] {
		if isAsteroid(b) {
			count++
		}
	}

	as := make([]Asteroid, 0, count)
	y := 0
	x := 0
	for _, b := range asteroids[start:] {
		if isEmpty(b) {
			x++
		} else if isNewline(b) {
			x = 0
			y++
		} else if isAsteroid(b) {
			as = append(as, image.Point{X: x, Y: y})
			x++
		} else {
			// whitespace, ignore
		}
	}
	return as
}

// Day10Part1 returns the asteroid that can see most asteroids, and the number of
// visible asteroids.
func Day10Part1(as []Asteroid) (Asteroid, int) {
	var best Asteroid
	maxVisible := 0
	for i := range as {
		// map of normalized directions - each unique direction = one visible asteroid
		visible := make(map[direction]bool, len(as))
		for j := range as {
			// skip ourself
			if i == j {
				continue
			}
			dx := as[j].X - as[i].X
			dy := as[j].Y - as[i].Y
			dir := normalize(dx, dy)
			visible[dir] = true
		}
		// found a better location?
		if len(visible) > maxVisible {
			best = as[i]
			maxVisible = len(visible)
		}
	}
	return best, maxVisible
}

// asteroidWithAngle stores an asteroid with its angle and distance from base.
type asteroidWithAngle struct {
	asteroid Asteroid
	angle    float64
	dist     int
}

// Day10Part2 determines the 200th asteroid that gets vaporized.
func Day10Part2(as []Asteroid, base Asteroid) int {
	// Build list of asteroids with their angles and distances
	var targets []asteroidWithAngle
	for _, a := range as {
		if a == base {
			continue
		}
		dx := a.X - base.X
		dy := a.Y - base.Y
		// Calculate angle: atan2 with Y axis pointing down
		// Rotate by 90° so "up" (negative Y) is 0
		angle := math.Atan2(float64(dx), float64(-dy))
		if angle < 0 {
			angle += 2 * math.Pi
		}
		dist := dx*dx + dy*dy // squared distance is fine for comparison
		targets = append(targets, asteroidWithAngle{a, angle, dist})
	}

	// Sort by angle, then by distance
	sort.Slice(targets, func(i, j int) bool {
		if targets[i].angle != targets[j].angle {
			return targets[i].angle < targets[j].angle
		}
		return targets[i].dist < targets[j].dist
	})

	// Vaporize in order, cycling through angles
	var vaporized []Asteroid
	for len(targets) > 0 {
		lastAngle := -1.0
		i := 0
		for i < len(targets) {
			if targets[i].angle != lastAngle {
				// Vaporize this one
				vaporized = append(vaporized, targets[i].asteroid)
				lastAngle = targets[i].angle
				// Remove from targets
				targets = append(targets[:i], targets[i+1:]...)
			} else {
				// Same angle as last vaporized, skip for this rotation
				i++
			}
		}
	}

	// Return the 200th vaporized asteroid
	a := vaporized[199]
	return a.X*100 + a.Y
}
