package main

import "testing"

func BenchmarkDetectCPU(b *testing.B) {
	for range b.N {
		_ = 42
	}
}
