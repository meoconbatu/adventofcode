package day1

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day1/input.txt"
	ins := readInput(inputFileName)
	output := part2Core(ins)
	fmt.Print(output)
}

func part2Core(ins []int) int {
	return countIncreaseMeasurementByWindow(ins, 3)
}

// count the number of times the sum of measurements in three-measurement sliding window
// increases from the previous sum
func countIncreaseMeasurementByWindow(depths []int, window int) int {
	rs := 0
	for i := window; i < len(depths); i++ {
		if depths[i] > depths[i-window] {
			rs++
		}
	}
	return rs
}
