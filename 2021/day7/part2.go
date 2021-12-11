package day7

import (
	"fmt"
	"math"
)

// Part2 func
func Part2() {
	inputFileName := "day7/input.txt"
	nums := readInput(inputFileName)
	output := part2Core(nums)
	fmt.Println(output)
}
func part2Core(ins []int) int {
	return sumFuel2(ins)
}

func sumFuel2(nums []int) int {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += float64(nums[i])
	}
	median1 := int(math.Floor(sum / float64(len(nums))))
	median2 := median1 + 1
	return min(f(nums, median1), f(nums, median2))
}
func f(nums []int, median int) int {
	fuel := 0
	for i := 0; i < len(nums); i++ {
		temp := int(math.Abs(float64(nums[i] - median)))
		fuel += ((temp + 1) * temp) / 2
	}
	return fuel
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
