package adventofcode2019

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
	return InputLines(InputFilename(day))
}

func InputLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	return Lines(f)
}

func Lines(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
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
