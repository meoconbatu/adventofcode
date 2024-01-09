package day14

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day14 type
type Day14 struct{}

// Part1 func
func (d Day14) Part1() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	tiltNorth(grid, m, n)
	fmt.Println(calcTotalLoad(grid, m, n))
}

// print the first 1000 total load to see the pattern
// after c cycle, the total load will repeat
// base on the input: i := (1000000000-124)%26
// the final result is at index i of the repeated array.
func (d Day14) Part2() {
	// grid := readInput()
	// m, n := len(grid), len(grid[0])
	// for i := 0; i < 1000; i++ {
	// tiltCycle(grid, m, n)
	// fmt.Println(calcTotalLoad(grid, m, n))
	// }
}
func calcTotalLoad(grid [][]byte, m, n int) int {
	rs := 0
	for i := 0; i < m; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if grid[i][j] == 'O' {
				cnt++
			}
		}
		rs += cnt * (m - i)
	}
	return rs
}

// rotate the matrix by 90 degree (clockwise)
func rotate(matrix [][]byte) {
	n := len(matrix) - 1
	for i := 0; i <= n/2; i++ {
		for j := i; j <= (n - i - 1); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = temp
		}
	}
}
func tiltCycle(grid [][]byte, m, n int) {
	for i := 0; i < 4; i++ {
		tiltNorth(grid, m, n)
		rotate(grid)
	}
}

// NEWS
func tiltNorth(grid [][]byte, m, n int) {
	for j := 0; j < n; j++ {
		isharp := -1
		cnt0 := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == 'O' {
				cnt0++
				grid[i][j] = '.'
			} else if grid[i][j] == '#' {
				rewriteRock(grid, &cnt0, isharp, i, j)
				isharp = i
			}
			if i == m-1 {
				rewriteRock(grid, &cnt0, isharp, i, j)
			}
		}
	}
}
func rewriteRock(grid [][]byte, cnt0 *int, isharp, i, j int) {
	if *cnt0 > 0 {
		k := isharp + 1
		for ; k <= isharp+*cnt0; k++ {
			grid[k][j] = 'O'
		}
		for ; k < i; k++ {
			grid[k][j] = '.'
		}
		*cnt0 = 0
	}
}
func readInput() [][]byte {
	scanner := utils.NewScanner(14)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		lines := strings.Split(scanner.Text(), "\n")
		for _, l := range lines {
			grid = append(grid, []byte(l))
		}
	}
	return grid
}
