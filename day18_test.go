package adventofcode2019

import (
	"fmt"
	"testing"
)

func TestDay18Part1Examples(t *testing.T) {
	wants := []uint{8, 86, 132, 136, 81}

	for i, want := range wants {
		example := uint8(i + 1)
		t.Run(fmt.Sprintf("example%d", example), func(t *testing.T) {
			testBytes(t, 18, func(d uint8) string { return exampleNFilename(d, example) }, true, Day18, want)
		})
	}
}

func TestDay18Part1(t *testing.T) {
	testBytes(t, 18, filename, true, Day18, 3962)
}

func BenchmarkDay18Part1(b *testing.B) {
	benchBytes(b, 18, true, Day18)
}
