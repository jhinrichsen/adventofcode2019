package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	cmd := exec.Command("go", "test", "-run=^$", "-bench=BenchmarkDetectCPU", "-benchtime=1ns", "./cmd/cpuname")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to detect CPU: %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`(?m)^cpu:\s+(.+)$`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) < 2 {
		fmt.Fprintf(os.Stderr, "could not find CPU info in benchmark output\n")
		os.Exit(1)
	}

	cpuName := matches[1]
	// Remove only parentheses (interfere with Make)
	cpuName = strings.ReplaceAll(cpuName, "(", "")
	cpuName = strings.ReplaceAll(cpuName, ")", "")
	// Replace spaces with underscores (Make requirement)
	cpuName = strings.ReplaceAll(cpuName, " ", "_")

	fmt.Print(cpuName)
}
