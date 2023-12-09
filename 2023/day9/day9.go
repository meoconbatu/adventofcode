package day9

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day9 type
type Day9 struct{}

// Part1 func
func (d Day9) Part1() {
	scanner := utils.NewScanner(9)
	rs := 0
	for scanner.Scan() {
		rs += getNextValue(utils.ParseIntSlice(scanner.Text()))
	}
	fmt.Println(rs)
}
func getNextValue(nums []int) int {
	for n := len(nums) - 1; ; n-- {
		isAllZero := true
		for i := 0; i < n; i++ {
			nums[i] = (nums[i+1] - nums[i])
			if nums[i] != 0 {
				isAllZero = false
			}
		}
		if isAllZero {
			break
		}
	}
	rs := 0
	for i := 0; i < len(nums); i++ {
		rs += nums[i]
	}
	return rs
}

// Part2 func
func (d Day9) Part2() {
	scanner := utils.NewScanner(9)
	rs := 0
	for scanner.Scan() {
		nums := utils.ParseIntSlice(scanner.Text())
		utils.Reverse(nums)
		rs += getNextValue(nums)
	}
	fmt.Println(rs)
}
