package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBinary(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"D2FE28", "110100101111111000101000"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			assert.Equal(t, tt.expected, toBinary(tt.in))
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		in       string
		expected int
	}{
		{"D2FE28", 2021},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			_, out := Parse(tt.in)
			assert.Equal(t, tt.expected, out)
		})
	}
}
