package adventofcode2019

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func exampleInput(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func input(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

func lineAsNumbers(line string) ([]int, error) {
	var (
		n   int
		ns  []int
		err error
	)
	for _, s := range strings.Fields(line) {
		n, err = strconv.Atoi(s)
		if err != nil {
			break
		}
		ns = append(ns, n)
	}
	return ns, err
}

// linesAsNumber converts strings into integer.
func linesAsNumbers(lines []string) ([]int, error) {
	var is []int
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			msg := "error in line %d: cannot convert %q to number"
			return is, fmt.Errorf(msg, i, lines[i])
		}
		is = append(is, n)
	}
	return is, nil
}

func numbersFromFilename(filename string) ([]int, error) {
	ls, err := linesFromFilename(filename)
	if err != nil {
		return nil, err
	}
	return linesAsNumbers(ls)
}
