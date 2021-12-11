package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		ins  string
		out  int
	}{
		{"base case", "input_test.txt", 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ins, l := readInput(tt.ins)
			got := part1Core(ins, l)
			assert.EqualValues(t, tt.out, got)
		})
	}
}
