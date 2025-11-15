package main

import "testing"

func BenchmarkDetectCPU(b *testing.B) {
	// This benchmark exists solely to provide CPU detection
	// The benchmark output includes: cpu: <CPU Name>
	for b.Loop() {
	}
}
