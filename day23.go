package adventofcode2019

// Day23 simulates a network of 50 Intcode computers.
// For part 1, it returns the Y value of the first packet sent to address 255.
// For part 2, it returns the first Y value delivered by the NAT twice in a row.
func Day23(program []byte, part1 bool) (uint, error) {
	ic, err := newIntcode(program)
	if err != nil {
		return 0, err
	}

	const networkSize = 50
	const natAddress = 255

	// Create 50 computers
	computers := make([]*intcode, networkSize)
	for i := range networkSize {
		computers[i] = ic.Clone()
	}

	// Packet queues for each computer (X,Y pairs)
	queues := make([][]int, networkSize)

	// Output buffers for each computer (collecting dest, X, Y)
	outBuffers := make([][]int, networkSize)

	// Track if each computer has received its address
	needsAddress := make([]bool, networkSize)
	for i := range networkSize {
		needsAddress[i] = true
	}

	// NAT state for part 2
	var natX, natY int
	hasNAT := false
	var lastNATY int
	hasLastNATY := false
	idleCycles := 0

	// Run each computer to next I/O point
	runToIO := func(addr int) state {
		for {
			state := computers[addr].Step()
			if state != running {
				return state
			}
		}
	}

	for {
		activity := false

		for addr := range networkSize {
			state := runToIO(addr)

			switch state {
			case needsInput:
				if needsAddress[addr] {
					computers[addr].Input(addr)
					needsAddress[addr] = false
					activity = true
				} else if len(queues[addr]) > 0 {
					computers[addr].Input(queues[addr][0])
					queues[addr] = queues[addr][1:]
					activity = true
				} else {
					computers[addr].Input(-1)
				}
			case hasOutput:
				activity = true
				outBuffers[addr] = append(outBuffers[addr], computers[addr].Output())
				if len(outBuffers[addr]) == 3 {
					dest := outBuffers[addr][0]
					x := outBuffers[addr][1]
					y := outBuffers[addr][2]
					outBuffers[addr] = outBuffers[addr][:0]

					if dest == natAddress {
						if part1 {
							return uint(y), nil
						}
						natX, natY = x, y
						hasNAT = true
					} else if dest >= 0 && dest < networkSize {
						queues[dest] = append(queues[dest], x, y)
					}
				}
			case halted:
				// Computer halted
			}
		}

		// Part 2: Check for network idle
		if !part1 && hasNAT && !activity {
			queuesEmpty := true
			for i := range networkSize {
				if len(queues[i]) > 0 {
					queuesEmpty = false
					break
				}
			}

			if queuesEmpty {
				idleCycles++
				if idleCycles > 2 {
					if hasLastNATY && lastNATY == natY {
						return uint(natY), nil
					}
					lastNATY = natY
					hasLastNATY = true
					queues[0] = append(queues[0], natX, natY)
					idleCycles = 0
				}
			} else {
				idleCycles = 0
			}
		} else {
			idleCycles = 0
		}
	}
}
