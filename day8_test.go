package main

import "testing"

func TestDay8Part2(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"0222112222120000", "0110"},
	}
	for _, tt := range tests {
		result := part82([]rune(tt.input), 2, 2)
		if result != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %s, want: %s.", result, tt.expectedOutput)
		}
	}
}
