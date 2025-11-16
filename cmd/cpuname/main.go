package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	// Run a minimal benchmark to get CPU info from Go's own detection
	// Use dedicated BenchmarkDetectCPU in this package
	cmd := exec.Command("go", "test", "-run=^$", "-bench=BenchmarkDetectCPU", "-benchtime=1ns", "./cmd/cpuname")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to detect CPU: %v\n", err)
		os.Exit(1)
	}

	// Extract CPU line: "cpu: Intel(R) Xeon(R) CPU @ 2.60GHz"
	re := regexp.MustCompile(`(?m)^cpu:\s+(.+)$`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		fmt.Fprintf(os.Stderr, "could not find CPU info in benchmark output\n")
		os.Exit(1)
	}

	// Clean up CPU name for filename use
	cpuName := matches[1]
	// Remove "CPU @ speed" suffix
	cpuName = regexp.MustCompile(`\s+CPU.*$`).ReplaceAllString(cpuName, "")
	// Remove special characters
	cpuName = regexp.MustCompile(`[()@/]`).ReplaceAllString(cpuName, "")
	// Replace spaces with underscores
	cpuName = strings.ReplaceAll(cpuName, " ", "_")
	// Collapse multiple underscores
	cpuName = regexp.MustCompile(`_+`).ReplaceAllString(cpuName, "_")
	// Trim trailing underscore
	cpuName = strings.TrimSuffix(cpuName, "_")

	fmt.Print(cpuName)
}
