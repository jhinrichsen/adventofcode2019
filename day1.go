package adventofcode2019

// Fuel computes required fuel for given mass. Both mass and fuel have no units.
func Fuel(mass int) int {
	return mass/3 - 2
}
