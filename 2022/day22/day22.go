package day22

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day22 type
type Day22 struct{}

type Grid struct {
	walls          map[utils.Point]struct{}
	minRow, maxRow int
	minCol, maxCol int
	m, n           int
}
type Position struct {
	x, y int
	face byte
}

var positions = make(map[utils.Point]Position)

// Part1 func
func (d Day22) Part1() {
	grids, path := readInput()
	num := 0
	pos := Position{0, 0, 0}
	igrid := 0
	for i, step := range path {
		if step >= '0' && step <= '9' {
			num = num*10 + int(step-'0')
		}
		if !(step >= '0' && step <= '9') || i == len(path)-1 {
			move(grids, &pos, num, &igrid)
			turn(grids, &pos, step, &igrid)
			num = 0
		}
	}
	curx, cury := pos.x+grids[igrid].minRow, pos.y+grids[igrid].minCol
	fmt.Println(curx*1000 + cury*4 + int(pos.face))
}

func print(grids []Grid, positions map[utils.Point]Position) {
	for _, grid := range grids {
		for i := grid.minRow; i <= grid.maxRow; i++ {
			for j := 1; j <= grid.maxCol; j++ {
				if j >= grid.minCol && j <= grid.maxCol {
					if _, ok := grid.walls[utils.Point{X: i - grid.minRow, Y: j - grid.minCol}]; ok {
						fmt.Print("#")
					} else if pos, ok := positions[utils.Point{X: i, Y: j}]; ok {
						switch pos.face {
						case 1:
							fmt.Print("v")
						case 0:
							fmt.Print(">")
						case 2:
							fmt.Print("<")
						case 3:
							fmt.Print("^")
						}
					} else {
						fmt.Print(".")
					}
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}
func turn(grids []Grid, position *Position, step rune, igrid *int) {
	if step != 'L' && step != 'R' {
		return
	}
	if step == 'R' {
		position.face++
	} else if step == 'L' {
		position.face--
	}
	position.face = (position.face%4 + 4) % 4
	positions[utils.Point{X: position.x + grids[*igrid].minRow, Y: position.y + grids[*igrid].minCol}] = Position{position.x + grids[*igrid].minRow, position.y + grids[*igrid].minCol, position.face}

}

func moveHorizontally(grids []Grid, position *Position, num int, igrid *int, direction int) {
	for i := 0; i < num; i++ {
		newy := position.y + direction
		if newy >= grids[*igrid].n {
			newy = 0
		} else if newy < 0 {
			newy = grids[*igrid].n - 1
		}
		newp := utils.Point{X: position.x, Y: newy}
		if _, ok := grids[*igrid].walls[newp]; ok {
			break
		} else {
			position.y = newy
		}
		positions[utils.Point{X: position.x + grids[*igrid].minRow, Y: position.y + grids[*igrid].minCol}] = Position{position.x + grids[*igrid].minRow, position.y + grids[*igrid].minCol, position.face}
	}
}
func moveVertically(grids []Grid, position *Position, num int, igrid *int, direction int) {
	n := len(grids)
	for i := 0; i < num; i++ {
		newx, newy := position.x+direction, position.y

		delta := 0
		if newx >= grids[*igrid].m {
			delta = 1
		} else if newx < 0 {
			delta = -1
		}
		iprev := *igrid
		if delta != 0 {
			for *igrid = *igrid + delta; ; *igrid += delta {
				*igrid = (*igrid%n + n) % n
				newy = (position.y + grids[iprev].minCol - grids[*igrid].minCol)
				if newy >= 0 && newy < grids[*igrid].n {
					break
				}
			}
		}
		if delta == 1 {
			newx = 0
		} else if delta == -1 {
			newx = grids[*igrid].m - 1
		}
		newp := utils.Point{X: newx, Y: newy}
		if _, ok := grids[*igrid].walls[newp]; ok {
			*igrid = iprev
			break
		} else {
			position.x, position.y = newx, newy
		}
		positions[utils.Point{X: position.x + grids[*igrid].minRow, Y: position.y + grids[*igrid].minCol}] = Position{position.x + grids[*igrid].minRow, position.y + grids[*igrid].minCol, position.face}
	}
}
func move(grids []Grid, position *Position, num int, igrid *int) {
	switch position.face {
	case 0:
		moveHorizontally(grids, position, num, igrid, 1)
	case 2:
		moveHorizontally(grids, position, num, igrid, -1)
	case 1:
		moveVertically(grids, position, num, igrid, 1)
	case 3:
		moveVertically(grids, position, num, igrid, -1)
	}
}

func readInput() ([]Grid, string) {
	scanner := utils.NewScanner(22)
	grids := make([]Grid, 0)

	walls := make(map[utils.Point]struct{})
	var minRow, maxRow, minCol, maxCol int
	path := ""
	for scanner.Scan() {
		s := scanner.Text()

		n := len(s)
		s = strings.TrimLeft(s, " ")
		nLeft := n - len(s) + 1
		s = strings.TrimRight(s, " ")
		nRight := n - len(s) - nLeft + 1
		nMid := n - nLeft - nRight + 1

		if !(minCol == nLeft && maxCol == nLeft+nMid-1) {
			if minCol != 0 || nMid == 0 {
				grid := Grid{walls, minRow, maxRow, minCol, maxCol, maxRow - minRow + 1, maxCol - minCol + 1}
				grids = append(grids, grid)
				if nMid == 0 {
					scanner.Scan()
					path = scanner.Text()
					break
				}
			}
			minRow = maxRow + 1
			maxRow = minRow
			minCol, maxCol = nLeft, nLeft+nMid-1
			walls = make(map[utils.Point]struct{})
		} else if minCol == nLeft && maxCol == nLeft+nMid-1 {
			maxRow++
		}
		for i, r := range s {
			if r == '#' {
				walls[utils.Point{X: maxRow - minRow, Y: i}] = struct{}{}
			}
		}
	}
	return grids, path
}
