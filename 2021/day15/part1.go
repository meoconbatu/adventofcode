package day15

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Item struct {
	I, J  int
	Value int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value < pq[j].Value
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Part1 func
func Part1() {
	inputFileName := "day15/input.txt"
	grid := readInput(inputFileName)
	output := part1Core(grid)
	fmt.Println(output)
}
func part1Core(grid [][]int) int {
	return findLowestRiskPath(grid)
}

var directions = [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func findLowestRiskPath(grid [][]int) int {
	n := len(grid)

	q := make(PriorityQueue, 0)
	heap.Init(&q)

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = math.MaxInt64

		}
	}

	dist[0][0] = 0
	heap.Push(&q, &Item{0, 0, dist[0][0]})
	x, y := 0, 0

	for q.Len() > 0 {
		item := heap.Pop(&q).(*Item)
		x, y = item.I, item.J

		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if nx < 0 || ny < 0 || nx >= n || ny >= n {
				continue
			}
			alt := dist[x][y] + grid[nx][ny]
			if alt < dist[nx][ny] {
				dist[nx][ny] = alt
				heap.Push(&q, &Item{nx, ny, alt})
			}
		}
	}
	return dist[n-1][n-1]
}

func readInput(inputFileName string) [][]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	grid := make([][]int, 0)
	for scanner.Scan() {
		rowStr := strings.Split(scanner.Text(), "")
		row := make([]int, 0)
		var num int
		for _, numStr := range rowStr {
			fmt.Sscanf(numStr, "%d", &num)
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	return grid
}
