package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay16Part1(t *testing.T) {
	tests := []struct {
		input          string
		inputNPhase    int
		expectedOutput string
	}{
		{`12345678`, 1, "48226158"},
		{`12345678`, 2, "34040438"},
		{`12345678`, 3, "03415518"},
		{`12345678`, 4, "01029498"},
		{`80871224585914546619083218645595`, 100, "24176176"},
		{`19617804207202209144916044189917`, 100, "73745418"},
		{`69317163492948606335995924319873`, 100, "52432133"},
	}
	for _, tt := range tests {
		buf := &bytes.Buffer{}
		buf.WriteString(tt.input)
		result := day16Part1(buf, tt.inputNPhase)
		assert.Equal(t, tt.expectedOutput, result)
	}
}

func TestDay16Part2(t *testing.T) {
	tests := []struct {
		input          string
		inputNPhase    int
		expectedOutput string
	}{
		{`03036732577212944063491565474664`, 100, "84462026"},
		{`02935109699940807407585447034323`, 100, "78725270"},
		{`03081770884921959731165446850517`, 100, "53553731"},
	}
	for _, tt := range tests {
		buf := &bytes.Buffer{}
		buf.WriteString(tt.input)
		result := day16Part2(buf, tt.inputNPhase)
		assert.Equal(t, tt.expectedOutput, result)
	}
}
