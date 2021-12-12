package day11

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day11/input.txt"
	energyLevels := readInput(inputFileName)
	output := part2Core(energyLevels)
	fmt.Println(output)
}
func part2Core(energyLevels [][]int) int {
	rs := 0
	for countFlash(energyLevels) != 100 {
		step(energyLevels)
		rs++
	}
	return rs
}
