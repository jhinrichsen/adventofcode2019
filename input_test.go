package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// InputFilename returns the puzzle input filename for a given day.
func InputFilename(day int) string {
	return TestdataFilename(fmt.Sprintf("day%d.txt", day))
}

func TestdataFilename(filename string) string {
	return filepath.Join("testdata", filename)
}

// InputLinesForDay returns all lines from given filename.
func InputLinesForDay(day int) ([]string, error) {
	return linesFromFilename(InputFilename(day))
}

// InputBuffer returns the puzzle input for a given day as bytes.
func InputBuffer(day int) ([]byte, error) {
	return ioutil.ReadFile(InputFilename(day))
}
