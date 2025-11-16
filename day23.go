package adventofcode2019

import (
	"sync"
	"time"
)

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

	// Packet queues for each computer
	queues := make([][]Packet, networkSize)
	var queuesMutex sync.Mutex

	// Create channels for each computer
	inputs := make([]chan int, networkSize)
	outputs := make([]chan int, networkSize)
	for i := range networkSize {
		inputs[i] = make(chan int, 10)
		outputs[i] = make(chan int, 10)
	}

	// Start all computers
	for i := range networkSize {
		go Day5(prog.Copy(), inputs[i], outputs[i])
		// Send network address as initial input
		inputs[i] <- i
	}

	// Input feeder goroutines
	for i := range networkSize {
		addr := i
		go func() {
			for {
				queuesMutex.Lock()
				if len(queues[addr]) > 0 {
					// Send packet from queue
					p := queues[addr][0]
					queues[addr] = queues[addr][1:]
					queuesMutex.Unlock()
					inputs[addr] <- p.X
					inputs[addr] <- p.Y
				} else {
					queuesMutex.Unlock()
					// Queue empty, send -1
					inputs[addr] <- -1
				}
				// Small delay to avoid busy loop
				time.Sleep(time.Microsecond)
			}
		}()
	}

	// Track NAT state for part 2
	var natPacket *Packet
	var lastNATY *int
	idleCycles := 0

	// Packet routing loop
	for {
		packetSent := false

		// Process outputs from all computers
		for addr := range networkSize {
			select {
			case dest := <-outputs[addr]:
				packetSent = true
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
					// Queue packet for destination
					queuesMutex.Lock()
					queues[dest] = append(queues[dest], Packet{X: x, Y: y})
					queuesMutex.Unlock()
				}
			default:
				// No output from this computer
			}
		}

		// Part 2: Check if network is idle
		if !part1 && natPacket != nil {
			// Check if all queues are empty
			queuesMutex.Lock()
			allQueuesEmpty := true
			for i := range networkSize {
				if len(queues[i]) > 0 {
					allQueuesEmpty = false
					break
				}
			}

			if !packetSent && allQueuesEmpty {
				idleCycles++
				// Network appears idle - wait a bit to confirm
				if idleCycles > 100 {
					// Send NAT packet to address 0
					y := natPacket.Y
					if lastNATY != nil && *lastNATY == y {
						queuesMutex.Unlock()
						return uint(y)
					}
					lastNATY = &y
					queues[0] = append(queues[0], *natPacket)
					idleCycles = 0
				}
			} else {
				idleCycles = 0
			}
			queuesMutex.Unlock()
		}

		// Small delay to avoid busy loop
		time.Sleep(time.Microsecond * 10)
	}
}
