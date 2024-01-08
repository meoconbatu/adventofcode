package day13

import (
	"fmt"
	"os"
	"strings"
)

// Day13 type
type Day13 struct{}

func (d Day13) Part1() {
	grids := readInput()
	rs := 0
	for _, grid := range grids {
		rs += findReflectionLine(grid)
	}
	fmt.Println(rs)
}

// findReflectionLine finds the line of reflection in the pattern (grid)
// and return the line index i (start from 1)
// if the line is vertical, i < 100
// else if the line is horizontal then i >= 100
func findReflectionLine(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	rows, cols := toRowCol(grid)
	for i := 0; i < n-1; i++ {
		if isVerticalMirror(rows, cols, i, i+1) {
			return i + 1
		}
	}
	for i := 0; i < m-1; i++ {
		if isHorizontalMirror(rows, cols, i, i+1) {
			return (i + 1) * 100
		}
	}
	return -1
}
func toRowCol(grid [][]byte) ([]string, []string) {
	m, n := len(grid), len(grid[0])
	rows := make([]string, m)
	for i := 0; i < m; i++ {
		rows[i] = string(grid[i])
	}
	cols := make([]string, n)
	for j := 0; j < n; j++ {
		for i := 0; i < m; i++ {
			cols[j] += string(rows[i][j])
		}
	}
	return rows, cols
}
func isVerticalMirror(rows, cols []string, l, r int) bool {
	i, j := l, r
	for ; i >= 0 && j < len(cols); i, j = i-1, j+1 {
		if cols[i] != cols[j] {
			break
		}
	}
	return i == -1 || j == len(cols)
}

func isHorizontalMirror(rows, cols []string, u, d int) bool {
	i, j := u, d
	for ; i >= 0 && j < len(rows); i, j = i-1, j+1 {
		if rows[i] != rows[j] {
			break
		}
	}
	return i == -1 || j == len(rows)
}

// Part2 func
func (d Day13) Part2() {
	grids := readInput()
	rs := 0
	for _, grid := range grids {
		iline := findReflectionLine(grid)
		iNewLine := findNewReflectionLine(grid, iline)
		rs += iNewLine
	}
	fmt.Println(rs)
}

// findNewReflectionLine finds the different line of reflection after changing a single smudge
func findNewReflectionLine(grid [][]byte, iline int) int {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			updateSmudge(grid, i, j)
			if iNewLines := findAllReflectionLines(grid); iNewLines != nil {
				for _, iNewLine := range iNewLines {
					if iline != iNewLine {
						return iNewLine
					}
				}
			}
			updateSmudge(grid, i, j)
		}
	}
	return -1
}
func updateSmudge(grid [][]byte, i, j int) {
	if grid[i][j] == '.' {
		grid[i][j] = '#'
	} else {
		grid[i][j] = '.'
	}
}
func findAllReflectionLines(grid [][]byte) []int {
	m, n := len(grid), len(grid[0])
	rows, cols := toRowCol(grid)

	ilines := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if isVerticalMirror(rows, cols, i, i+1) {
			ilines = append(ilines, i+1)
		}
	}
	for i := 0; i < m-1; i++ {
		if isHorizontalMirror(rows, cols, i, i+1) {
			ilines = append(ilines, (i+1)*100)
		}
	}
	return ilines
}
func readInput() [][][]byte {
	b, _ := os.ReadFile(fmt.Sprintf("day%d/input.txt", 13))
	parts := strings.Split(string(b), "\n\n")
	grids := make([][][]byte, 0)
	for _, p := range parts {
		lines := strings.Split(p, "\n")
		grid := make([][]byte, 0)
		for _, l := range lines {
			grid = append(grid, []byte(l))
		}
		grids = append(grids, grid)
	}
	return grids
}
