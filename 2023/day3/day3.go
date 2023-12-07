package day3

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day3 type
type Day3 struct{}

// Part1 func
func (d Day3) Part1() {
	scanner := utils.NewScanner(3)
	rs := 0
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	rs = sumAllNumbers(grid)
	fmt.Println(rs)
}
func sumAllNumbers(grid []string) int {
	rs := 0
	for i := 0; i < len(grid); i++ {
		num := 0
		l := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= '0' && grid[i][j] <= '9' {
				if num == 0 {
					l = j
				}
				num *= 10
				num += int(grid[i][j] - '0')
			}
			if (!(grid[i][j] >= '0' && grid[i][j] <= '9') || j == len(grid[i])-1) && num > 0 {
				if isAdjacentToSymbol(grid, i, j, l) {
					rs += num
				}
				num = 0
			}
		}
	}
	return rs
}

func isAdjacentToSymbol(grid []string, i, j, l int) bool {
	for k := -1; k <= 1; k++ {
		for kk := l - 1; kk <= j && kk < len(grid[i]); kk++ {
			if i+k >= 0 && i+k < len(grid) && kk >= 0 && grid[i+k][kk] != '.' && !(grid[i+k][kk] >= '0' && grid[i+k][kk] <= '9') {
				return true
			}
		}
	}
	return false
}

// Part2 func
func (d Day3) Part2() {
	scanner := utils.NewScanner(3)
	rs := 0
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	rs = sumAllGearRatios(grid)
	fmt.Println(rs)
}

func sumAllGearRatios(grid []string) int {
	rs := 0
	m := make(map[[2]int][]int)
	for i := 0; i < len(grid); i++ {
		num := 0
		l := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= '0' && grid[i][j] <= '9' {
				if num == 0 {
					l = j
				}
				num *= 10
				num += int(grid[i][j] - '0')
			}
			if (!(grid[i][j] >= '0' && grid[i][j] <= '9') || j == len(grid[i])-1) && num > 0 {
				if x, y := getAdjacentAsteriskPos(grid, i, j, l); x != -1 {
					m[[2]int{x, y}] = append(m[[2]int{x, y}], num)
				}
				num = 0
			}
		}
	}
	for _, gear := range m {
		if len(gear) == 2 {
			rs += gear[0] * gear[1]
		}
	}
	return rs
}
func getAdjacentAsteriskPos(grid []string, i, j, l int) (int, int) {
	for k := -1; k <= 1; k++ {
		for kk := l - 1; kk <= j && kk < len(grid[i]); kk++ {
			if i+k >= 0 && i+k < len(grid) && kk >= 0 && grid[i+k][kk] != '.' && !(grid[i+k][kk] >= '0' && grid[i+k][kk] <= '9') {
				if grid[i+k][kk] == '*' {
					return i + k, kk
				}
				return -1, -1
			}
		}
	}
	return -1, -1
}
