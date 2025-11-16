package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

func main() {
	// Run a benchmark to get CPU info
	cmd := exec.Command("go", "test", "-run=^$", "-bench=BenchmarkDetectCPU", "-benchtime=1ns")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error running benchmark:", err)
		os.Exit(1)
	}

	// Extract CPU name from output
	re := regexp.MustCompile(`cpu:\s+(.+)`)
	matches := re.FindSubmatch(output)
	if len(matches) < 2 {
		fmt.Fprintln(os.Stderr, "Could not detect CPU")
		os.Exit(1)
	}

	cpuName := string(matches[1])

	// Clean up CPU name for use in filename
	// Remove speed info (e.g., "@ 2.60GHz", "CPU @ 2.40GHz")
	cpuName = regexp.MustCompile(`\s+@.*`).ReplaceAllString(cpuName, "")
	cpuName = regexp.MustCompile(`\s+CPU.*`).ReplaceAllString(cpuName, "")

	// Remove special characters
	cpuName = strings.ReplaceAll(cpuName, "(R)", "")
	cpuName = strings.ReplaceAll(cpuName, "(TM)", "")
	cpuName = strings.ReplaceAll(cpuName, "(", "")
	cpuName = strings.ReplaceAll(cpuName, ")", "")
	cpuName = strings.ReplaceAll(cpuName, "/", "")

	// Replace spaces with underscores
	cpuName = strings.ReplaceAll(cpuName, " ", "_")

	// Collapse multiple underscores
	cpuName = regexp.MustCompile(`_+`).ReplaceAllString(cpuName, "_")

	// Trim trailing underscores
	cpuName = strings.TrimSuffix(cpuName, "_")

	fmt.Println(cpuName)
}

// BenchmarkDetectCPU is a dummy benchmark used only for CPU detection
func BenchmarkDetectCPU(b *testing.B) {}
