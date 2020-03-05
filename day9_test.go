package adventofcode2019

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestDay9Clone(t *testing.T) {
	want := MustSplit("109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99")
	in, out := channels()
	var proc IntCodeProcessor = Day5
	go proc(want, in, out)
	var got IntCode
	for ic := range out {
		got = append(got, ic)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestDay9Digits16(t *testing.T) {
	prog := MustSplit("1102,34915192,34915192,7,4,7,99,0")
	in, out := channels()
	var proc IntCodeProcessor = Day5
	// look mom - sync call
	proc(prog, in, out)
	want := 16
	got := Ndigits(<-out)
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestDay9LargeNumber(t *testing.T) {
	prog := MustSplit("104,1125899906842624,99")
	in, out := channels()
	var proc IntCodeProcessor = Day5
	// look mom - async call
	go proc(prog, in, out)
	want := 1125899906842624
	got := <-out
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func TestDay9Part1(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day9.txt")
	if err != nil {
		t.Fatal(err)
	}
	prog := MustSplit(string(buf))
	in, out := channels()
	in <- 1
	var proc IntCodeProcessor = Day5
	go proc(prog, in, out)
	wantLen := 1
	want := 2436480432
	var codes []int
	for got := range out {
		codes = append(codes, got)
	}
	gotLen := len(codes)
	if wantLen != gotLen {
		for _, opcode := range codes {
			fmt.Printf("broken opcode: %d\n", opcode)
		}
		t.Fatalf("want len %d, got len %d", wantLen, gotLen)
	}
	got := codes[len(codes)-1]
	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
