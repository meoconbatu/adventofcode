package day23

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day23 type
type Day23 struct{}

// Part1 func
func (d Day23) Part1() {
	grid, maxrow, maxcol := readInput()
	minp, maxp := utils.Point{0, 0}, utils.Point{maxrow, maxcol}

	idirection := 0
	for round := 0; round < 10; round++ {
		grid, _ = move(grid, &minp, &maxp, idirection)
		idirection = (idirection + 1) % 4
	}
	rs := 0
	for _, v := range grid {
		if v == 1 {
			rs++
		}
	}
	fmt.Println((maxp.X-minp.X)*(maxp.Y-minp.Y) - rs)
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
	grid, maxrow, maxcol := readInput()
	minp, maxp := utils.Point{X: 0, Y: 0}, utils.Point{X: maxrow, Y: maxcol}

	round := 0
	for idirection, numMove := 0, 1; numMove != 0; idirection = (idirection + 1) % 4 {
		grid, numMove = move(grid, &minp, &maxp, idirection)
		round++
	}
	fmt.Println(round)
}
func move(grid map[utils.Point]int, minp, maxp *utils.Point, idirection int) (map[utils.Point]int, int) {
	grid2 := make(map[utils.Point]int, maxp.X)
	from := make(map[utils.Point]utils.Point)
	// direction := "NSWE"
	xToNum, yToNum := make(map[int]int), make(map[int]int)
	for i := minp.X; i < maxp.X; i++ {
		for j := minp.Y; j < maxp.Y; j++ {
			p := utils.Point{i, j}
			if v, ok := grid[p]; !ok || (ok && v != 1) {
				continue
			}
			n := []utils.Point{{i - 1, j - 1}, {i - 1, j}, {i - 1, j + 1}}
			s := []utils.Point{{i + 1, j - 1}, {i + 1, j}, {i + 1, j + 1}}
			w := []utils.Point{{i - 1, j - 1}, {i, j - 1}, {i + 1, j - 1}}
			e := []utils.Point{{i - 1, j + 1}, {i, j + 1}, {i + 1, j + 1}}
			nswe := [][]utils.Point{n, s, w, e}
			masks := [4]int{}
			for ii, ps := range nswe {
				mask := 0
				for jj, p := range ps {
					if grid[p] == 1 {
						mask |= 1 << jj
					}
				}
				masks[ii] = mask
			}
			if masks[0]|masks[1]|masks[2]|masks[3] == 0 {
				grid2[p] = 1
				xToNum[p.X]++
				yToNum[p.Y]++
				continue
			}
			canMove := false
			for ii, icur := 0, idirection; ii < 4; ii, icur = ii+1, (icur+1)%4 {
				if masks[icur] == 0 {
					pp := p
					if icur == 0 {
						pp.X--
					} else if icur == 1 {
						pp.X++
					} else if icur == 2 {
						pp.Y--
					} else {
						pp.Y++
					}
					if v, ok := grid2[pp]; !ok {
						canMove = true
						grid2[pp] = 1
						from[pp] = p
						xToNum[pp.X]++
						yToNum[pp.Y]++
					} else if v == 1 {
						grid2[pp] = -1
						xToNum[pp.X]--
						yToNum[pp.Y]--

						grid2[from[pp]] = 1
						xToNum[from[pp].X]++
						yToNum[from[pp].Y]++

						delete(from, pp)
					}
					break
				}
			}
			if !canMove {
				grid2[p] = 1
				xToNum[p.X]++
				yToNum[p.Y]++
			}
		}
	}
	for i := -1; i <= 1; i++ {
		if xToNum[minp.X+i] > 0 {
			minp.X += i
			break
		}
	}
	for i := 1; i >= -1; i-- {
		if xToNum[maxp.X-1+i] > 0 {
			maxp.X += i
			break
		}
	}
	for i := -1; i <= 1; i++ {
		if yToNum[minp.Y+i] > 0 {
			minp.Y += i
			break
		}
	}
	for i := 1; i >= -1; i-- {
		if yToNum[maxp.Y-1+i] > 0 {
			maxp.Y += i
			break
		}
	}
	return grid2, len(from)
}
func readInput() (map[utils.Point]int, int, int) {
	scanner := utils.NewScanner(23)
	grid := make(map[utils.Point]int, 0)
	i := 0
	maxrow := 0
	maxcol := 0
	for scanner.Scan() {
		s := scanner.Text()
		for j, r := range s {
			if r == '#' {
				grid[utils.Point{i, j}] = 1
				maxrow = i
				maxcol = utils.Max(maxcol, j)
			}
		}
		i++
	}
	return grid, maxrow + 1, maxcol + 1
}
