package adventofcode2019

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

// exampleNFilename returns the filename for example N of a given day.
// Use this for generic access: exampleNFilename(18, 3) -> "testdata/day18_example3.txt"
func exampleNFilename(day uint8, example uint8) string {
	return fmt.Sprintf("testdata/day%02d_example%d.txt", int(day), int(example))
}

func example1Filename(day uint8) string {
	return exampleNFilename(day, 1)
}

func example2Filename(day uint8) string {
	return exampleNFilename(day, 2)
}

func example3Filename(day uint8) string {
	return exampleNFilename(day, 3)
}

func example4Filename(day uint8) string {
	return exampleNFilename(day, 4)
}

func example5Filename(day uint8) string {
	return exampleNFilename(day, 5)
}

func example6Filename(day uint8) string {
	return exampleNFilename(day, 6)
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", int(day))
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

func input(day int) string {
	return fmt.Sprintf("testdata/day%02d.txt", day)
}

// linesAsNumber converts strings into integer.
