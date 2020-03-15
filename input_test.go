package adventofcode2019

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// InputFilename returns the puzzle input filename for a given day.
func InputFilename(day int) string {
	return filepath.Join("testdata", fmt.Sprintf("day%d.txt", day))
}

// InputLines returns all lines from given filename.
func InputLines(day int) ([]string, error) {
	var lines []string
	f, err := os.Open(InputFilename(day))
	if err != nil {
		return lines, err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

// InputBuffer returns the puzzle input for a given day as bytes.
func InputBuffer(day int) ([]byte, error) {
	return ioutil.ReadFile(InputFilename(day))
}
