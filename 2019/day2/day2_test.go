package day2

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunProgram(t *testing.T) {
	tests := []struct {
		in     string
		expect string
	}{
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50"},
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
		{"1002,4,3,4,33", "1002,4,3,4,99"},
		{"1101,100,-1,4,0", "1101,100,-1,4,99"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			nums := stringToArray(tt.in)
			// RunProgram(nums, nil, nil)
			ic := Init(nums, nil, nil)
			ic.Run()
			assert.Equal(t, tt.expect, toString(nums))
		})
	}
}
func TestRunProgram_Input(t *testing.T) {
	tests := []struct {
		in       string
		systemID int
		expect   int
	}{
		{"3,9,8,9,10,9,4,9,99,-1,8", 8, 1},
		{"3,9,8,9,10,9,4,9,99,-1,8", 7, 0},
		{"3,9,7,9,10,9,4,9,99,-1,8", 7, 1},
		{"3,3,1108,-1,8,3,4,3,99", 7, 0},
		{"3,3,1107,-1,8,3,4,3,99", 7, 1},
		{"3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", 7, 1},
		{"3,3,1105,-1,9,1101,0,0,12,4,12,99,1", 0, 0},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 7, 999},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 9, 1001},
		{"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", 8, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var actual int
			nums := stringToArray(tt.in)
			// RunProgram(nums, &tt.systemID, &actual)
			ic := Init(nums, []int{tt.systemID}, &actual)
			ic.Run()
			assert.Equal(t, tt.expect, actual)
		})
	}
}
func TestRunProgram_Offset(t *testing.T) {
	tests := []struct {
		in     string
		expect int
	}{
		{"1102,34915192,34915192,7,4,7,99,0", 34915192 * 34915192},
		{"104,1125899906842624,99", 1125899906842624},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			var actual int
			nums := stringToArray(tt.in)
			// RunProgram(nums, nil, &actual)
			ic := Init(nums, nil, &actual)
			ic.Run()
			assert.Equal(t, tt.expect, actual)
		})
	}
}
func toString(nums []int) string {
	rs := ""
	for i := 0; i < len(nums); i++ {
		rs += fmt.Sprintf("%d,", nums[i])
	}
	return strings.TrimRight(rs, ",")
}
