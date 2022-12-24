package day24

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day24 type
type Day24 struct{}

// Part1 func
func (d Day24) Part1() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	start, end := [2]int{0, 1}, [2]int{m - 1, n - 2}

	r1, _ := bfs(grid, start, end, m, n)
	fmt.Println(r1)
}

// Part2 func
func (d Day24) Part2() {
	grid := readInput()
	m, n := len(grid), len(grid[0])
	start, end := [2]int{0, 1}, [2]int{m - 1, n - 2}

	r1, grid1 := bfs(grid, start, end, m, n)
	r2, grid2 := bfs(grid1, end, start, m, n)
	r3, _ := bfs(grid2, start, end, m, n)
	fmt.Println(r1 + r2 + r3)
}
func bfs(grid [][]int, start, end [2]int, m, n int) (int, [][]int) {
	q := [][2]int{start}
	lenq := 1

	qMap := make(map[[2]int][]int)
	qMap[start] = []int{1}

	roundToGrid := make(map[int][][]int)
	roundToGrid[0] = grid

	round := 1
	roundToGrid[round] = move(grid, m, n)

	visited := make(map[int]bool)
	for {
		curX, curY := q[0][0], q[0][1]
		curRound := qMap[[2]int{curX, curY}][0]
		curgrid := roundToGrid[curRound]

		visited[key(curRound, curX, curY)] = true

		qMap[[2]int{curX, curY}] = qMap[[2]int{curX, curY}][1:]
		q = q[1:]
		lenq--
		directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {0, 0}}
		if curgrid[curX][curY] != 0 {
			directions = directions[:len(directions)-1]
		}
		for _, d := range directions {
			newX, newY := curX+d[0], curY+d[1]
			if newX < 0 || newY < 0 || newX >= m || newY >= n || curgrid[newX][newY] != 0 || visited[key(round+1, newX, newY)] {
				continue
			}
			if newX == end[0] && newY == end[1] {
				return round, roundToGrid[round]
			}
			q = append(q, [2]int{newX, newY})
			qMap[[2]int{newX, newY}] = append(qMap[[2]int{newX, newY}], round+1)
			visited[key(round+1, newX, newY)] = true

		}
		if lenq == 0 {
			lenq = len(q)
			round++
			curgrid = move(curgrid, m, n)
			roundToGrid[round] = curgrid
		}
	}
}
func key(round, x, y int) int {
	return (round+1)*100000000 + x*100000 + y
}
func print(grid [][]int, m, n, x, y int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == x && j == y {
				fmt.Printf("E")
			} else if grid[i][j] == -1 {
				fmt.Printf("#")
			} else if grid[i][j] == 0 {
				fmt.Print(".")
			} else {
				ones := 0
				num := grid[i][j]
				for num > 0 {
					if num&1 == 1 {
						ones++
					}
					num >>= 1
				}
				if ones > 1 {
					fmt.Printf("%d", ones)
				} else {
					if grid[i][j] == 0b0001 {
						fmt.Print("^")
					} else if grid[i][j] == 0b0010 {
						fmt.Print(">")
					} else if grid[i][j] == 0b0100 {
						fmt.Print("v")
					} else if grid[i][j] == 0b1000 {
						fmt.Print("<")
					}
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func move(grid [][]int, m, n int) [][]int {
	grid2 := make([][]int, m)
	directions := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for i := range grid2 {
		grid2[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 || grid[i][j] == -1 {
				if grid[i][j] == -1 {
					grid2[i][j] = grid[i][j]
				}
				continue
			}
			for k := 0; k < 4; k++ {
				if grid[i][j]&(1<<k) > 0 {
					newi, newj := i, j
					for {
						newi, newj = newi+directions[k][0], newj+directions[k][1]
						newi, newj = (newi%m+m)%m, (newj%n+n)%n
						if grid[newi][newj] != -1 {
							grid2[newi][newj] |= (1 << k)
							break
						}
					}
				}
			}
		}
	}
	return grid2
}

func readInput() [][]int {
	scanner := utils.NewScanner(24)
	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		s := scanner.Text()
		for _, c := range s {
			switch c {
			case '^':
				row = append(row, 1<<0)
			case '>':
				row = append(row, 1<<1)
			case 'v':
				row = append(row, 1<<2)
			case '<':
				row = append(row, 1<<3)
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, -1)
			}
		}
		grid = append(grid, row)
	}
	return grid
}
