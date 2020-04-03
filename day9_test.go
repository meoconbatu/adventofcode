package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestProcessDay9(t *testing.T) {
	tests := []struct {
		input1         []int64
		input2         string
		expectedOutput int64
	}{
		{[]int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, "", 0},
		{[]int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, "", 34915192 * 34915192},
		{[]int64{104, 1125899906842624, 99}, "", 1125899906842624},
	}
	for _, tt := range tests {
		temp := &bytes.Buffer{}
		temp.WriteString(tt.input2)
		process(tt.input1, temp, temp)
		result, _ := strconv.Atoi(temp.String())
		if int64(result) != tt.expectedOutput {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", result, tt.expectedOutput)
		}
	}
}
