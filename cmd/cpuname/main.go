package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sanitizeCPUName(input string) string {
	var result strings.Builder

	for _, r := range input {
		switch {
		case r == ' ':
			result.WriteRune('-')
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'):
			result.WriteRune(r)
		case r == '@' || r == '_' || r == '-':
			result.WriteRune(r)
		default:
			// Skip any other character
		}
	}

	return result.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu: ") {
			cpuName := strings.TrimPrefix(line, "cpu: ")
			cpuName = strings.TrimSpace(cpuName)
			fmt.Print(sanitizeCPUName(cpuName))
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "no 'cpu: ' line found in input\n")
	os.Exit(1)
}
