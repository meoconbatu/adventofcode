package day23

import (
	"fmt"
	"math"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day23 type
type Day23 struct{}

// Part1 func
func (d Day23) Part1() {
	grid := readInput()

	idirection := 0
	for round := 0; round < 10; round++ {
		grid, _ = move(grid, idirection)
		idirection = (idirection + 1) % 4
	}
	rs := 0
	minp, maxp := utils.Point{X: math.MaxInt64, Y: math.MaxInt64}, utils.Point{X: math.MinInt64, Y: math.MinInt64}
	for pnum, v := range grid {
		if v == 1 {
			px, py := mapIntToXY(pnum)
			minp.X = utils.Min(minp.X, px)
			minp.Y = utils.Min(minp.Y, py)
			maxp.X = utils.Max(maxp.X, px)
			maxp.Y = utils.Max(maxp.Y, py)
			rs++
		}
	}
	fmt.Println((1+maxp.X-minp.X)*(1+maxp.Y-minp.Y) - rs)
}

func print(grid map[utils.Point]int, minp, maxp utils.Point) {
	for i := minp.X; i < maxp.X; i++ {
		for j := minp.Y; j < maxp.Y; j++ {
			if v, ok := grid[utils.Point{X: i, Y: j}]; ok && v == 1 {
				fmt.Printf("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Part2 func
func (d Day23) Part2() {
	grid := readInput()

	round := 0
	for idirection, numMove := 0, 1; numMove != 0; idirection = (idirection + 1) % 4 {
		grid, numMove = move(grid, idirection)
		round++
	}
	fmt.Println(round)
}

func mapXYToInt(x, y int) int {
	min := -200
	m := 400
	return (x-min)*m + y - min
}
func mapIntToXY(num int) (int, int) {
	min := -200
	m := 400
	return num/m + min, num%m + min
}

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func move(grid map[int]byte, idirection int) (map[int]byte, int) {
	grid2 := make(map[int]byte)
	from := make(map[int]int)
	for pnum, v := range grid {
		if v != 1 {
			continue
		}
		px, py := mapIntToXY(pnum)
		totalMask, masks := getMask(grid, px, py)
		canMove := false
		for ii, icur := 0, idirection; totalMask != 0 && ii < 4; ii, icur = ii+1, (icur+1)%4 {
			if masks[icur] != 0 {
				continue
			}
			ppx, ppy := px+directions[icur][0], py+directions[icur][1]
			ppnum := mapXYToInt(ppx, ppy)
			if v, ok := grid2[ppnum]; !ok {
				canMove = true
				grid2[ppnum] = 1
				from[ppnum] = pnum
			} else if v == 1 {
				grid2[ppnum] = 2
				grid2[from[ppnum]] = 1
				delete(from, ppnum)
			}
			break
		}
		if !canMove {
			grid2[pnum] = 1
		}
	}
	return grid2, len(from)
}
func getMask(grid map[int]byte, x, y int) (int, [4]int) {
	n := []int{mapXYToInt(x-1, y-1), mapXYToInt(x-1, y), mapXYToInt(x-1, y+1)}
	s := []int{mapXYToInt(x+1, y-1), mapXYToInt(x+1, y), mapXYToInt(x+1, y+1)}
	w := []int{mapXYToInt(x-1, y-1), mapXYToInt(x, y-1), mapXYToInt(x+1, y-1)}
	e := []int{mapXYToInt(x-1, y+1), mapXYToInt(x, y+1), mapXYToInt(x+1, y+1)}
	nswe := [][]int{n, s, w, e}
	masks := [4]int{}
	totalMask := 0
	for ii, ps := range nswe {
		for jj, pp := range ps {
			if grid[pp] == 1 {
				masks[ii] |= 1 << jj
			}
		}
		totalMask |= masks[ii]
	}
	return totalMask, masks
}
func readInput() map[int]byte {
	scanner := utils.NewScanner(23)
	grid := make(map[int]byte, 0)
	for i := 0; scanner.Scan(); i++ {
		s := scanner.Text()
		for j, r := range s {
			if r == '#' {
				num := mapXYToInt(i, j)
				grid[num] = 1
			}
		}
	}
	return grid
}
