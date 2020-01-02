package adventofcode2019

import (
	"reflect"
	"strconv"
	"testing"
)

var digitsExamples = []struct {
	in      int
	ndigits int
	digits  []byte
}{
	{-101, 3, []byte{1, 0, 1}},
	{-100, 3, []byte{1, 0, 0}},
	{-99, 2, []byte{9, 9}},
	{-1, 1, []byte{1}},
	{-0, 1, []byte{0}},
	{0, 1, []byte{0}},
	{1, 1, []byte{1}},
	{9, 1, []byte{9}},
	{10, 2, []byte{1, 0}},
	{99, 2, []byte{9, 9}},
	{100, 3, []byte{1, 0, 0}},
	{101, 3, []byte{1, 0, 1}},
	{1_000_000, 7, []byte{1, 0, 0, 0, 0, 0, 0}},
}

func TestDigits(t *testing.T) {
	for _, tt := range digitsExamples {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			n := Ndigits(tt.in)
			if tt.ndigits != n {
				t.Fatalf("want %d but got %d", tt.ndigits, n)
			}
			ds := Digits(tt.in)
			if !reflect.DeepEqual(tt.digits, ds) {
				t.Fatalf("want %v but got %v", tt.digits, ds)
			}
		})
	}
}
