package day12

import (
	"fmt"
	"math"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day12 type
type Day12 struct{}

type Point struct {
	x, y int
}

// Part1 func
func (d Day12) Part1() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	var x, y int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				x, y = i, j
				break
			}
		}
	}
	fmt.Println(bfs(grid, m, n, x, y, 'E'))
}
func bfs(grid [][]byte, m, n, x, y int, target byte) int {
	direction := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	steps := 0
	q := []int{x, y}
	visited[x][y] = true
	lenq := len(q)

	for len(q) > 0 {
		curx, cury := q[0], q[1]
		q = q[2:]
		lenq -= 2
		for _, d := range direction {
			newx, newy := curx+d[0], cury+d[1]
			if newx < 0 || newy < 0 || newx >= m || newy >= n || visited[newx][newy] {
				continue
			}
			cur, new := grid[curx][cury], grid[newx][newy]
			if curx == x && cury == y {
				cur = 'a'
			}
			if new == target {
				new = 'z'
			}
			if cur+1 < new {
				continue
			}
			if grid[newx][newy] == target {
				return steps + 1
			}
			q = append(q, newx, newy)
			visited[newx][newy] = true
		}
		if lenq == 0 {
			steps++
			lenq = len(q)
		}
	}
	return -1
}

// Part2 func
func (d Day12) Part2() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	rs := math.MaxInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' || grid[i][j] == 'a' {
				temp := bfs(grid, m, n, i, j, 'E')
				if temp == -1 {
					continue
				}
				rs = utils.Min(rs, temp)
			}
		}
	}
	fmt.Println(rs)
}

func readInput() [][]byte {
	scanner := utils.NewScanner(12)

	rs := make([][]byte, 0)
	for scanner.Scan() {
		rs = append(rs, []byte(scanner.Text()))
	}
	return rs
}
