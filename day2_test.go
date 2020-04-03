package main

import (
	"os"
	"reflect"
	"testing"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		input          []int64
		expectedOutput []int64
	}{
		{[]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, []int64{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
		{[]int64{1, 0, 0, 0, 99}, []int64{2, 0, 0, 0, 99}},
		{[]int64{2, 4, 4, 5, 99, 0}, []int64{2, 4, 4, 5, 99, 9801}},
		{[]int64{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int64{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, tt := range tests {
		process(tt.input, os.Stdin, os.Stdout)
		if !reflect.DeepEqual(tt.input, tt.expectedOutput) {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", tt.input, tt.expectedOutput)
		}
	}
}
