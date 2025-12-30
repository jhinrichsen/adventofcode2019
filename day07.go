package adventofcode2019

// Day07 computes maximum thruster signal for amplifier circuits
func Day07(input []byte, part1 bool) (uint, error) {
	ic, err := NewIntcode(input)
	if err != nil {
		return 0, err
	}
	if part1 {
		return uint(day7Part1(ic)), nil
	}
	return uint(day7Part2(ic)), nil
}

func day7Part1(ic *Intcode) int {
	maxThrust := 0
	phases := []int{0, 1, 2, 3, 4}

	permute(phases, func(perm []int) {
		signal := 0
		for _, phase := range perm {
			ic.Reset()
			outputs, _ := ic.Run(phase, signal)
			if len(outputs) > 0 {
				signal = outputs[0]
			}
		}
		if signal > maxThrust {
			maxThrust = signal
		}
	})

	return maxThrust
}

func day7Part2(ic *Intcode) int {
	maxThrust := 0
	phases := []int{5, 6, 7, 8, 9}

	permute(phases, func(perm []int) {
		// Create 5 amplifiers
		amps := make([]*Intcode, 5)
		for i := range 5 {
			amps[i] = ic.Clone()
		}

		// Initialize each amp with its phase setting
		for i, phase := range perm {
			runUntilNeedsInput(amps[i])
			amps[i].Input(phase)
		}

		// Run feedback loop
		signal := 0
		lastOutput := 0
		halted := 0

		for halted < 5 {
			for i := range 5 {
				if amps[i] == nil {
					continue
				}

				// Run until needs input or has output or halted
				for {
					state := amps[i].Step()
					switch state {
					case NeedsInput:
						amps[i].Input(signal)
					case HasOutput:
						signal = amps[i].Output()
						lastOutput = signal
						goto nextAmp
					case Halted:
						halted++
						amps[i] = nil
						goto nextAmp
					}
				}
			nextAmp:
			}
		}

		if lastOutput > maxThrust {
			maxThrust = lastOutput
		}
	})

	return maxThrust
}

func runUntilNeedsInput(ic *Intcode) {
	for {
		state := ic.Step()
		if state == NeedsInput || state == Halted {
			return
		}
	}
}

// permute calls f with each permutation of a
func permute(a []int, f func([]int)) {
	permuteHelper(a, len(a), f)
}

func permuteHelper(a []int, k int, f func([]int)) {
	if k == 1 {
		f(a)
		return
	}
	permuteHelper(a, k-1, f)
	for i := range k - 1 {
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}
		permuteHelper(a, k-1, f)
	}
}
