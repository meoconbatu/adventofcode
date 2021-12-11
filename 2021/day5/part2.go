package day5

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day5/input.txt"
	froms, tos := readInput(inputFileName)
	output := part2Core(froms, tos)
	fmt.Println(output)
}
func part2Core(froms, tos []Point) int {
	return countIntersect2(froms, tos)
}

// the number of points where at least two lines overlap
func countIntersect2(froms, tos []Point) int {
	m := make(map[Point]int)
	rs := 0
	for i := 0; i < len(froms); i++ {
		if froms[i].x == tos[i].x {
			for j := min(froms[i].y, tos[i].y); j <= max(froms[i].y, tos[i].y); j++ {
				m[Point{froms[i].x, j}]++
				if m[Point{froms[i].x, j}] == 2 {
					rs++
				}
			}
		} else if froms[i].y == tos[i].y {
			for j := min(froms[i].x, tos[i].x); j <= max(froms[i].x, tos[i].x); j++ {
				m[Point{j, froms[i].y}]++
				if m[Point{j, froms[i].y}] == 2 {
					rs++
				}
			}
		} else {
			j, k := froms[i].x, froms[i].y
			for (j >= tos[i].x && froms[i].x > tos[i].x) ||
				(j <= tos[i].x && froms[i].x < tos[i].x) {
				m[Point{j, k}]++
				if m[Point{j, k}] == 2 {
					rs++
				}
				if froms[i].x > tos[i].x {
					j--
				} else {
					j++
				}
				if froms[i].y > tos[i].y {
					k--
				} else {
					k++
				}
			}
		}

	}
	return rs
}
