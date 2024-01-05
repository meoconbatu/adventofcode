package day10

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day10 type
type Day10 struct{}

const (
	out  = math.MinInt64
	in   = math.MaxInt64
	null = -9999999
)

var (
	pointToNextTiles = map[utils.Point]string{
		{X: 0, Y: -1}: "-LFS", //L
		{X: 0, Y: 1}:  "-7JS", //R
		{X: -1, Y: 0}: "|7FS", //U
		{X: 1, Y: 0}:  "|JLS", //D
	}

	tileToPoints = map[byte][]utils.Point{
		'S': {{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}},
		'-': {{X: 0, Y: -1}, {X: 0, Y: 1}},
		'7': {{X: 0, Y: -1}, {X: 1, Y: 0}},
		'|': {{X: -1, Y: 0}, {X: 1, Y: 0}},
		'L': {{X: -1, Y: 0}, {X: 0, Y: 1}},
		'J': {{X: 0, Y: -1}, {X: -1, Y: 0}},
		'F': {{X: 0, Y: 1}, {X: 1, Y: 0}},
	}
)

// Part1 func
func (d Day10) Part1() {
	grid, startp := readInput()
	fmt.Println(cntStep(grid, startp))
}

// cntStep counts the number of steps along the loop
// required to reach the point farthest from the starting position
func cntStep(grid []string, startP utils.Point) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs1(grid, m, n, startP.X, startP.Y, visited) / 2
}

func dfs1(grid []string, m, n, x, y int, visited [][]bool) int {
	if visited[x][y] {
		if grid[x][y] == 'S' {
			return 0
		}
		return math.MinInt64
	}
	visited[x][y] = true
	rs := math.MinInt64
	for _, d := range tileToPoints[grid[x][y]] {
		cx, cy := x+d.X, y+d.Y
		if cx < 0 || cy < 0 || cx >= m || cy >= n || grid[cx][cy] == '.' || !strings.Contains(pointToNextTiles[d], string(grid[cx][cy])) {
			continue
		}
		rs = utils.Max(rs, dfs1(grid, m, n, cx, cy, visited)+1)

	}
	visited[x][y] = false
	return rs
}

// Part2 func
func (d Day10) Part2() {
	grid, startp := readInput()

	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	loopPoints := dfs2(grid, m, n, startp.X, startp.Y, visited)
	fmt.Println(cntEnclosedTiles(loopPoints, m, n))
}

// dfs2 finds the loop points
func dfs2(grid []string, m, n, x, y int, visited [][]bool) []utils.Point {
	if visited[x][y] {
		if grid[x][y] == 'S' {
			return []utils.Point{}
		}
		return nil
	}
	visited[x][y] = true
	rs := []utils.Point{}
	for _, d := range tileToPoints[grid[x][y]] {
		cx, cy := x+d.X, y+d.Y
		if cx < 0 || cy < 0 || cx >= m || cy >= n || grid[cx][cy] == '.' || !strings.Contains(pointToNextTiles[d], string(grid[cx][cy])) {
			continue
		}
		rss := dfs2(grid, m, n, cx, cy, visited)
		if rss != nil {
			if len(rs) < len(rss)+1 {
				rs = append(rss, utils.Point{X: cx, Y: cy})
			}
		}
	}
	visited[x][y] = false
	return rs
}

// cntEnclosedTiles counts the tiles enclosed by the loop
//  1. mark the loop using increasing numbers
//  2. from 4 edges, mark all the points outside the loop with out const
//  3. with the remain points, use isInside to check if it's really inside the loop or not
func cntEnclosedTiles(loopPoints []utils.Point, m, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	// step 1
	i := 1
	for _, p := range loopPoints {
		grid[p.X][p.Y] = i
		i++
	}
	// step 2
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 {
				if grid[i][j] != 0 {
					continue
				}
				bfs(grid, m, n, i, j, out)
			}
		}
	}
	// step 3
	rs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			if isInside(grid, i, j, m, n) {
				rs++
				grid[i][j] = in
			} else {
				grid[i][j] = out
			}
		}
	}
	return rs
}

// isInside return true if a point p(x,y) inside the loop, else return false
// count the number of points n at which the loop intersects the x-axis from the left to point p.
// p is inside if n is odd
func isInside(grid [][]int, x, y, m, n int) bool {
	cur := null
	cnt := 0
	for i := 0; i < y; {
		if grid[x][i] == out || grid[x][i] == in {
			i++
			continue
		}
		if cur == null {
			cur = grid[x][i]
			i++
			continue
		}
		if utils.Abs(grid[x][i]-cur) != 1 {
			cur = grid[x][i]
			cnt++
			i++
			continue
		}
		diff := grid[x][i] - cur
		l, r := i-1, i
		for j := i + 1; j < y; j++ {
			if (grid[x][j] - grid[x][j-1]) != diff {
				break
			}
			r = j
		}
		if !(grid[x][l]-grid[x-1][l] == diff && grid[x-1][r]-grid[x][r] == diff) &&
			!(grid[x][l]-grid[x+1][l] == diff && grid[x+1][r]-grid[x][r] == diff) {
			cnt++
		}
		i = r + 1
		cur = null
	}
	if cur != null {
		cnt++
	}
	return cnt%2 == 1
}
func bfs(grid [][]int, m, n, x, y, val int) {
	for _, d := range tileToPoints['S'] {
		cx, cy := x+d.X, y+d.Y
		if cx < 0 || cy < 0 || cx >= m || cy >= n || grid[cx][cy] != 0 {
			continue
		}
		grid[cx][cy] = val
		bfs(grid, m, n, cx, cy, val)
	}
}

func readInput() ([]string, utils.Point) {
	scanner := utils.NewScanner(10)
	grid := make([]string, 0)
	i := 0
	x, y := 0, 0
	for scanner.Scan() {
		s := scanner.Text()
		grid = append(grid, s)
		if idx := strings.Index(s, "S"); idx != -1 {
			x, y = i, idx
		}
		i++
	}
	return grid, utils.Point{X: x, Y: y}
}
func print(grid [][]int, m, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%-5d", (grid[i][j]))
		}
		fmt.Println()
	}
}
