package day13

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day13/input.txt"
	points, folds := readInput(inputFileName)
	output := part2Core(points, folds)
	fmt.Println(output)
}
func part2Core(points map[Point]struct{}, folds []Point) int {
	for _, f := range folds {
		fold(points, f)
	}
	// print to see the code (eight capital letters)
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if _, ok := points[Point{i, j}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	return 0
}
