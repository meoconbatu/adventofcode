package day16

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day16 type
type Day16 struct{}

var (
	// tranforms describes the beam's behavior depends on what it encounters as it moves.
	// For example:
	// ">.": ">" 	means a rightward-moving beam that encounters empty space (.), it continues in the same direction (>).
	// ">|": "^v" 	means a rightward-moving beam that encounters a | splitter would split into two beams: one upward and one downward.
	tranforms = map[string]string{
		">.": ">", "<.": "<", "v.": "v", "^.": "^",
		">|": "^v", "<|": "^v", "^|": "^", "v|": "v",
		">-": ">", "<-": "<", "^-": "<>", "v-": "<>",
		">/": "^", "</": "v", "^/": ">", "v/": "<",
		">\\": "v", "<\\": "^", "^\\": "<", "v\\": ">",
	}
	directionToIndex = map[rune]int{'>': 1, '<': 2, '^': 3, 'v': 4}
	directionToPoint = map[rune][]int{
		'^': {-1, 0}, 'v': {1, 0}, '>': {0, 1}, '<': {0, -1},
	}
)

// Part1 func
func (d Day16) Part1() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	// the beam starting in the top-left heading right
	rs := totalEnergizedTile(grid, m, n, 0, 0, '>')
	fmt.Println(rs)
}

func totalEnergizedTile(grid []string, m, n, i, j int, d rune) int {
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}
	dfs(grid, visited, m, n, i, j, d)
	return countVisitedTiles(visited, m, n)
}
func dfs(grid []string, visited [][]int, m, n, x, y int, direction rune) {
	visited[x][y] |= (1 << directionToIndex[direction])
	for _, d := range tranforms[string(direction)+string(grid[x][y])] {
		newx, newy := directionToPoint[d][0]+x, directionToPoint[d][1]+y
		if newx < 0 || newy < 0 || newx >= m || newy >= n || visited[newx][newy]&(1<<directionToIndex[d]) > 0 {
			continue
		}
		dfs(grid, visited, m, n, newx, newy, d)
	}
}

func countVisitedTiles(visited [][]int, m, n int) int {
	rs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] > 0 {
				rs++
			}
		}
	}
	return rs
}

// Part2 func
func (d Day16) Part2() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	rs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				rs = utils.Max(rs, totalEnergizedTile(grid, m, n, i, j, 'v'))
			}
			if i == m-1 {
				rs = utils.Max(rs, totalEnergizedTile(grid, m, n, i, j, '^'))
			}
			if j == 0 {
				rs = utils.Max(rs, totalEnergizedTile(grid, m, n, i, j, '>'))
			}
			if j == n-1 {
				rs = utils.Max(rs, totalEnergizedTile(grid, m, n, i, j, '<'))
			}
		}
	}
	fmt.Println(rs)
}
func readInput() []string {
	scanner := utils.NewScanner(16)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func print(grid []string, visited [][]int, m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' || visited[i][j] == 0 {
				fmt.Print(string(grid[i][j]) + " ")
				continue
			}
			if visited[i][j]&(1<<1) > 0 {
				fmt.Print("> ")
			} else if visited[i][j]&(1<<2) > 0 {
				fmt.Print("< ")
			} else if visited[i][j]&(1<<3) > 0 {
				fmt.Print("^ ")
			} else if visited[i][j]&(1<<4) > 0 {
				fmt.Print("v ")
			}
		}
		fmt.Println()
	}
}
