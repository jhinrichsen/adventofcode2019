package adventofcode2019

import (
	"math/rand/v2"
)

// Day7Part1 returns max thrust for given program, using a permutation of all
// phases.
func Day7Part1(prog IntCode, phases string) int {
	max := 0
	perms := permutations(phases)
	for perm := range perms {
		// fmt.Printf("calculating permutation %q\n", perm)
		// convert "0" -> 0, "1" -> 1, ...
		buf := []byte(perm)
		for i := 0; i < len(buf); i++ {
			buf[i] -= '0'
		}
		in, out := amps(prog, buf)
		// fmt.Printf("0 -> %+v\n", in)
		in <- 0
		// fmt.Printf("waiting for %+v\n", out)
		n := <-out
		// fmt.Printf("read %d from %+v\n", n, out)
		// fmt.Printf("permutation %q returns %d\n", perm, n)
		if n > max {
			max = n
		}
	}
	return max
}

func amps(prog IntCode, phases []byte) (chan<- int, <-chan int) {
	var proc IntCodeProcessor = Day5
	var connect, lastConnect chan int
	input, output := channels()
	// fmt.Printf("created input=%+v, output=%+v\n", input, output)
	for i, phase := range phases {
		if i == 0 {
			connect = make(chan int)
			lastConnect = connect
			// fmt.Printf("created connect=%+v\n", connect)
			// fmt.Printf("creating first amp %+v -> %+v\n",
			// input, connect)
			go proc(prog.Copy(), input, connect)
			input <- int(phase)
			// fmt.Printf("amp %d init: %d -> %+v\n",
			// i, int(phase), input)
		} else if i+1 == len(phases) {
			// fmt.Printf("creating last amp %+v -> %+v\n",
			// lastConnect, output)
			go proc(prog.Copy(), lastConnect, output)
			lastConnect <- int(phase)
			// fmt.Printf("amp %d init: %d -> %+v\n",
			// i, int(phase), lastConnect)
		} else {
			connect = make(chan int)
			// fmt.Printf("creating amp #%d: %+v -> %+v\n",
			// i, lastConnect, connect)
			go proc(prog.Copy(), lastConnect, connect)
			lastConnect <- int(phase)
			// fmt.Printf("amp %d init: %d -> %+v\n",
			// i, int(phase), lastConnect)
			lastConnect = connect
		}
	}
	return input, output
}

func permutations(s string) map[string]bool {
	m := make(map[string]bool, fac(len(s)))
	// Heap or Yates would be better but require more coding
	buf := []byte(s)
	for len(m) < fac(len(s)) {
		rand.Shuffle(len(s), func(i, j int) {
			buf[i], buf[j] = buf[j], buf[i]
		})
		m[string(buf)] = true
	}
	return m
}

func fac(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// Day7Part2 returns the maximum thrust for a feedback loop.
func Day7Part2(prog IntCode, phases string) int {
	max := 0
	perms := permutations(phases)
	for perm := range perms {
		// convert "0" -> 0, "1" -> 1, ...
		buf := []byte(perm)
		for i := 0; i < len(buf); i++ {
			buf[i] -= '0'
		}
		in, out := amps(prog, buf)
		in <- 0

		// feedback loop
		for n := range out {
			if n > max {
				max = n
			}
			in <- n
		}
	}
	return max
}

// NewDay07 parses IntCode program from input
var NewDay07 = func(input []byte) IntCode {
	return MustSplit(string(input))
}

// Day07 computes maximum thruster signal for amplifier circuits
func Day07(input []byte, part1 bool) uint {
	prog := NewDay07(input)
	if part1 {
		return uint(Day7Part1(prog, "01234"))
	}
	return uint(Day7Part2(prog, "56789"))
}
