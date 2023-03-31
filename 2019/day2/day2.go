package day2

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day2 type
type Day2 struct{}

// Part1 func
func (d Day2) Part1() {
	nums := readInput()
	nums[1], nums[2] = 12, 2
	ic := Init(nums, nil, nil)
	ic.Run()

	fmt.Println(nums[0])
}

// RunProgram func
func RunProgram(nums []int, input, output *int) {
	offset := 0
	i := 0
	for {
		op, paramsMode := getMode(nums[i])
		switch op {
		case 99:
			return
		case 1:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			nums[nums[i+3]] = vals[0] + vals[1]
			i += 4
		case 2:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			nums[nums[i+3]] = vals[0] * vals[1]
			i += 4
		case 3:
			ia := nums[i+1]
			nums[ia] = *input
			i += 2
		case 4:
			vals := getParamsValue(nums, nums[i+1:i+2], paramsMode, offset)
			*output = vals[0]
			i += 2
		case 5:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			if vals[0] != 0 {
				i = vals[1]
			} else {
				i += 3
			}
		case 6:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			if vals[0] == 0 {
				i = vals[1]
			} else {
				i += 3
			}
		case 7:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			if vals[0] < vals[1] {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			i += 4
		case 8:
			vals := getParamsValue(nums, nums[i+1:i+3], paramsMode, offset)
			if vals[0] == vals[1] {
				nums[nums[i+3]] = 1
			} else {
				nums[nums[i+3]] = 0
			}
			i += 4
		case 9:
			vals := getParamsValue(nums, nums[i+1:i+2], paramsMode, offset)
			offset += vals[0]
		}
	}
}
func getParamsValue(nums, params []int, paramsMode, offset int) []int {
	rs := make([]int, len(params))
	for i := 0; i < len(rs); i++ {
		var val int
		switch paramsMode % 10 {
		case 0:
			val = nums[params[i]]
		case 2:
			val = nums[params[i]+offset]
		case 1:
			val = params[i]
		}
		paramsMode /= 10
		rs[i] = val
	}
	return rs
}
func getMode(in int) (int, int) {
	return in % 100, in / 100
}

// Part2 func
func (d Day2) Part2() {
	nums := readInput()
	copyNums := make([]int, len(nums))
	for pos1 := 0; pos1 <= 99; pos1++ {
		for pos2 := 0; pos2 <= 99; pos2++ {
			copy(copyNums, nums)
			copyNums[1], copyNums[2] = pos1, pos2
			ic := Init(copyNums, nil, nil)
			ic.Run()
			if copyNums[0] == 19690720 {
				fmt.Println(100*pos1 + pos2)
				return
			}
		}
	}
}

func readInput() []int {
	scanner := utils.NewScanner(2)
	scanner.Scan()
	return stringToArray(scanner.Text())
}
func stringToArray(s string) []int {
	rs := make([]int, 0)
	numStrs := strings.Split(s, ",")
	var num int
	for _, numStr := range numStrs {
		fmt.Sscanf(numStr, "%d", &num)
		rs = append(rs, num)
	}
	return rs
}
