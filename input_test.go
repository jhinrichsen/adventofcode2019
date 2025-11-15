package adventofcode2019

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func testLinesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	lines := testLinesFromReader(tb, f)
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return lines
}

func testLinesFromReader(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", int(day))
}

func example1Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example1.txt", int(day))
}

func example2Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example2.txt", int(day))
}

func example3Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example3.txt", int(day))
}

func example4Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example4.txt", int(day))
}

func example5Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example5.txt", int(day))
}

func example6Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example6.txt", int(day))
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", int(day))
}

// file reads the main input file bytes for day N (zero-padded).
func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return buf
}

// fileFromFilename reads file bytes using a filename function (e.g., filename or exampleFilename).
func fileFromFilename(tb testing.TB, filenameFunc func(uint8) string, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filenameFunc(day))
	if err != nil {
		tb.Fatal(err)
	}
	if b, ok := tb.(*testing.B); ok {
		b.ResetTimer()
	}
	return buf
}

// Backward-compatible functions for old test files (int-based API)

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
	if err := sc.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func exampleInput(day int) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", day)
}

func input(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
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
