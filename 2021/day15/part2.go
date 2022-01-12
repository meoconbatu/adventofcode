package day15

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day15/input.txt"
	grid := readInput(inputFileName)
	output := part2Core(grid)
	fmt.Println(output)
}
func part2Core(grid [][]int) int {
	fullGrid := getEntireCave(grid)
	return findLowestRiskPath(fullGrid)
}
func getEntireCave(grid [][]int) [][]int {
	rs := make([][]int, 5*len(grid))
	for i := 0; i < 5*len(grid); i++ {
		rs[i] = make([]int, 5*len(grid[0]))
	}
	for i := 0; i < 5*len(grid); i++ {
		for j := 0; j < 5*len(grid[0]); j++ {
			rs[i][j] = grid[i%len(grid)][j%len(grid[0])] + (i / len(grid)) + (j / len(grid))
			if rs[i][j] >= 10 {
				rs[i][j] = (rs[i][j] % 10) + 1
			}
		}
	}
	return rs
}
