package day11

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day11 type
type Day11 struct{}

// Part1 func
func (d Day11) Part1() {
	grid := readInput()
	fmt.Println(fn(grid, 1))
}

// Part2 func
func (d Day11) Part2() {
	grid := readInput()
	fmt.Println(fn(grid, 1000000-1))
}

func fn(grid []string, x int) int {
	m, n := len(grid), len(grid[0])
	points := make([][]int, 0)

	rows := make([]bool, m)
	cols := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				points = append(points, []int{i, j})
				rows[i], cols[j] = true, true
			}
		}
	}
	rs := 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			rs += utils.Abs(points[i][0]-points[j][0]) + utils.Abs(points[i][1]-points[j][1])
			for ii := utils.Min(points[j][0], points[i][0]); ii <= utils.Max(points[j][0], points[i][0]); ii++ {
				if !rows[ii] {
					rs += x
				}
			}
			for ii := utils.Min(points[j][1], points[i][1]); ii <= utils.Max(points[j][1], points[i][1]); ii++ {
				if !cols[ii] {
					rs += x
				}
			}
		}
	}
	return rs
}
func readInput() []string {
	scanner := utils.NewScanner(11)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}
