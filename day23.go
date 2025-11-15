package adventofcode2019

// Packet represents a network packet with X and Y values.
type Packet struct {
	X, Y int
}

// Day23 simulates a network of 50 Intcode computers.
// For part 1, it returns the Y value of the first packet sent to address 255.
// For part 2, it returns the first Y value delivered by the NAT twice in a row.
func Day23(prog IntCode, part1 bool) uint {
	const networkSize = 50
	const natAddress = 255

	// Create channels for each computer
	inputs := make([]chan int, networkSize)
	outputs := make([]chan int, networkSize)
	for i := range networkSize {
		inputs[i] = make(chan int, 1000)
		outputs[i] = make(chan int, 1000)
	}

	// Start all computers
	for i := range networkSize {
		go Day5(prog.Copy(), inputs[i], outputs[i])
		// Send network address as initial input
		inputs[i] <- i
	}

	// Track NAT state for part 2
	var natPacket *Packet
	var lastNATY *int
	idleCount := 0

	// Packet routing loop
	for {
		allIdle := true

		// Process outputs from all computers
		for addr := range networkSize {
			select {
			case dest := <-outputs[addr]:
				allIdle = false
				// Read X and Y values
				x := <-outputs[addr]
				y := <-outputs[addr]

				if dest == natAddress {
					if part1 {
						return uint(y)
					}
					// Part 2: NAT stores the packet
					natPacket = &Packet{X: x, Y: y}
				} else if dest >= 0 && dest < networkSize {
					// Send to destination computer
					inputs[dest] <- x
					inputs[dest] <- y
				}
			default:
				// No output from this computer, send -1 if idle
				select {
				case inputs[addr] <- -1:
				default:
				}
			}
		}

		// Part 2: Check if network is idle and NAT should send packet
		if !part1 {
			if allIdle {
				idleCount++
				if idleCount > 10 && natPacket != nil {
					// Network is idle, send NAT packet to address 0
					if lastNATY != nil && *lastNATY == natPacket.Y {
						return uint(natPacket.Y)
					}
					lastNATY = &natPacket.Y
					inputs[0] <- natPacket.X
					inputs[0] <- natPacket.Y
					idleCount = 0
				}
			} else {
				idleCount = 0
			}
		}
	}
}
