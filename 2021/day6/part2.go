package day6

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day6/input.txt"
	nums := readInput(inputFileName)
	output := part2Core(nums)
	fmt.Println(output)
}
func part2Core(ins []int) int {
	return countLanternfish(ins, 256)
}
