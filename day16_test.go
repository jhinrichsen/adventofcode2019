package adventofcode2019

import (
	"os"
	"testing"
)

func TestDay16Part1Examples(t *testing.T) {
	tests := []struct {
		input  string
		phases int
		want   string
	}{
		{"12345678", 4, "01029498"},
	}

	for i, tt := range tests {
		t.Run(string(rune('a'+i)), func(t *testing.T) {
			digits := parseDigits([]byte(tt.input))
			for range tt.phases {
				digits = applyFFTPhase(digits)
			}
			// Convert to string for comparison
			result := ""
			for _, d := range digits {
				result += string(rune('0' + d))
			}
			if result != tt.want {
				t.Fatalf("want %v but got %v", tt.want, result)
			}
		})
	}
}

func TestDay16Part1ExamplesFirst8(t *testing.T) {
	tests := []struct {
		input string
		want  uint
	}{
		{"80871224585914546619083218645595", 24176176},
		{"19617804207202209144916044189917", 73745418},
		{"69317163492948606335995924319873", 52432133},
	}

	for i, tt := range tests {
		t.Run(string(rune('a'+i)), func(t *testing.T) {
			got := fftPart1([]byte(tt.input))
			if got != tt.want {
				t.Fatalf("want %v but got %v", tt.want, got)
			}
		})
	}
}

func TestDay16Part1(t *testing.T) {
	prog, err := os.ReadFile(filename(16))
	if err != nil {
		t.Fatal(err)
	}
	const want = 59281788
	got := Day16(prog, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	prog, err := os.ReadFile(filename(16))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		Day16(prog, true)
	}
}
