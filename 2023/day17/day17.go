package day17

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/meoconbatu/adventofcode/utils"
)

var (
	directionToPoint = map[rune][]int{'^': {-1, 0}, 'v': {1, 0}, '>': {0, 1}, '<': {0, -1}}
	directionToIndex = map[rune]int{'>': 1, '<': 2, '^': 3, 'v': 0}
)

// Day17 type
type Day17 struct{}

// Part1 func
func (d Day17) Part1() {
	grid := readInput()
	fmt.Println(utils.Min(dijkstra(grid, 0, 0, '>', 1, 3), dijkstra(grid, 0, 0, 'v', 1, 3)))
}

// Part2 func
func (d Day17) Part2() {
	grid := readInput()
	fmt.Println(utils.Min(dijkstra(grid, 0, 0, 'v', 4, 10), dijkstra(grid, 0, 0, '>', 4, 10)))
}

func dijkstra(grid [][]int, x, y int, d rune, mi, ma int) int {
	m, n := len(grid), len(grid[0])
	cnts := make([][][4]int, m)
	for i := 0; i < m; i++ {
		cnts[i] = make([][4]int, n)
	}
	dists := make([][][4]int, m)
	for i := 0; i < m; i++ {
		dists[i] = make([][4]int, n)
		for j := 0; j < n; j++ {
			dists[i][j] = [4]int{math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64}
		}
	}
	dists[x][y] = [4]int{0, 0, 0, 0}
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{x: x, y: y, priority: 0, direction: d}
	heap.Init(&pq)
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		ux, uy, ud, udist := item.x, item.y, item.direction, item.priority
		for vd, p := range directionToPoint {
			if (ud == '>' && vd == '<') || (ud == '<' && vd == '>') || (ud == 'v' && vd == '^') || (ud == '^' && vd == 'v') ||
				ud == vd {
				continue
			}
			vx, vy, vdist, vcnt := ux, uy, udist, 0
			for i := 1; i <= ma; i++ {
				vx, vy = vx+p[0], vy+p[1]
				if vx < 0 || vy < 0 || vx >= m || vy >= n {
					continue
				}
				vdist += grid[vx][vy]
				vcnt++
				if i < mi {
					continue
				}
				if ivd := directionToIndex[vd]; vdist < dists[vx][vy][ivd] || cnts[vx][vy][ivd] > vcnt {
					dists[vx][vy][ivd] = utils.Min(dists[vx][vy][ivd], vdist)
					cnts[vx][vy][ivd] = vcnt
					heap.Push(&pq, &Item{x: vx, y: vy, priority: vdist, direction: vd})
				}
			}
		}
	}
	return utils.Min(utils.Min(dists[m-1][n-1][0], dists[m-1][n-1][1]), utils.Min(dists[m-1][n-1][3], dists[m-1][n-1][2]))
}

func readInput() [][]int {
	scanner := utils.NewScanner(17)
	grid := make([][]int, 0)
	for scanner.Scan() {
		grid = append(grid, utils.ParseIntSlice(scanner.Text(), ""))
	}
	return grid
}

func printInt(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 0 {
				fmt.Print("    .")
			} else {
				fmt.Printf("% 5d", (arr[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
func print(arr [][]rune) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Printf("%s", string(arr[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
