package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharToInt(t *testing.T) {
	expected := 1
	for i := byte('a'); i <= 'z'; i++ {
		assert.Equal(t, expected, charToInt(i))
		expected++
	}
	for i := byte('A'); i <= 'Z'; i++ {
		assert.Equal(t, expected, charToInt(i))
		expected++
	}
}
func TestGetPriority(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
		{"ZrLtpLnlfQJqnfJtpLnZlrqdNNGqcDNNFFTNDzzMMTMsMNMs", 17},
		{"lqgqsgvjVtbMDzzbtcDQ", 20},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.expected, GetPriority(tt.in))
		})
	}
}
