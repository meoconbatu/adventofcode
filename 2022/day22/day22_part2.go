package day22

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

type Face struct {
	grid           []string
	l, r, u, d     int
	minRow, minCol int
}

func (f Face) String() string {
	rs := ""
	for _, s := range f.grid {
		rs += fmt.Sprintln(s)
	}
	rs += fmt.Sprintln(f.l, f.r, f.u, f.d)
	rs += fmt.Sprintln(f.minRow, f.minCol)
	rs += fmt.Sprintln()
	return rs
}

// Part2 func
func (d Day22) Part2() {
	faces, path := readInput2()

	pos := Position{0, 0, 0}
	igrid := 5
	num := 0
	for i, step := range path {
		if step >= '0' && step <= '9' {
			num = num*10 + int(step-'0')
		}
		if !(step >= '0' && step <= '9') || i == len(path)-1 {
			move2(faces, &pos, num, &igrid)
			turn2(faces, &pos, step, &igrid)
			num = 0
		}
	}
	// print(faces, positions)
	curx, cury := pos.x+faces[igrid].minRow, pos.y+faces[igrid].minCol
	fmt.Println(curx*1000 + cury*4 + int(pos.face))
}

func turn2(face []Face, position *Position, step rune, igrid *int) {
	if step != 'L' && step != 'R' {
		return
	}
	if step == 'R' {
		position.face++
	} else if step == 'L' {
		position.face--
	}
	position.face = (position.face%4 + 4) % 4
}

func move2(faces []Face, position *Position, num int, igrid *int) {
	for i := 0; i < num; i++ {
		var dx, dy int
		switch position.face {
		case 0:
			dy = 1
		case 2:
			dy = -1
		case 1:
			dx = 1
		case 3:
			dx = -1
		}
		newy := position.y + dy
		newx := position.x + dx
		igridnew := *igrid
		newturn := ""
		if newy >= 50 {
			igridnew = faces[*igrid].r
			if (*igrid == 1 && igridnew == 3) || (*igrid == 4 && igridnew == 6) {
				newx, newy = position.y, position.x
				newturn = "L"
			} else if (*igrid == 3 && igridnew == 6) || (*igrid == 6 && igridnew == 3) {
				newx, newy = 49-position.x, position.y
				newturn = "LL"
			} else {
				newx, newy = position.x, 0
			}
		} else if newy < 0 {
			igridnew = faces[*igrid].l
			if (*igrid == 1 && igridnew == 5) || (*igrid == 4 && igridnew == 2) {
				newx, newy = position.y, position.x
				newturn = "L"
			} else if (*igrid == 2 && igridnew == 5) || (*igrid == 5 && igridnew == 2) {
				newx, newy = 49-position.x, position.y
				newturn = "LL"
			} else {
				newx, newy = position.x, 49
			}
		} else if newx >= 50 {
			igridnew = faces[*igrid].d
			if (*igrid == 3 && igridnew == 1) || (*igrid == 6 && igridnew == 4) {
				newx, newy = position.y, position.x
				newturn = "R"
			} else {
				newx, newy = 0, position.y
			}
		} else if newx < 0 {
			igridnew = faces[*igrid].u
			if (*igrid == 2 && igridnew == 4) || (*igrid == 5 && igridnew == 1) {
				newx, newy = position.y, position.x
				newturn = "R"
			} else {
				newx, newy = 49, position.y
			}
		}
		if faces[igridnew].grid[newx][newy] == '#' {
			break
		} else {
			position.x, position.y = newx, newy
			*igrid = igridnew
		}
		for _, d := range newturn {
			turn2(faces, position, d, igrid)
		}
	}
}
func readInput2() ([]Face, string) {
	scanner := utils.NewScanner(22)
	faces := make([]Face, 7)

	grid5 := make([]string, 50)
	grid6 := make([]string, 50)
	minRow, minCol := 1, 0
	for i := 0; i < 50; i++ {
		scanner.Scan()
		s := scanner.Text()

		n := len(s)
		s = strings.TrimLeft(s, " ")
		nLeft := n - len(s) + 1
		s = strings.TrimRight(s, " ")
		minCol = nLeft
		grid5[i] = s[:50]
		grid6[i] = s[50:]
	}
	faces[5] = Face{grid5, 2, 6, 1, 4, minRow, minCol}
	faces[6] = Face{grid6, 5, 3, 1, 4, minRow, minCol + 50}

	grid4 := make([]string, 50)
	minRow, minCol = minRow+50, 0
	for i := 0; i < 50; i++ {
		scanner.Scan()
		s := scanner.Text()

		n := len(s)
		s = strings.TrimLeft(s, " ")
		nLeft := n - len(s) + 1
		s = strings.TrimRight(s, " ")
		minCol = nLeft
		grid4[i] = s
	}
	faces[4] = Face{grid4, 2, 6, 5, 3, minRow, minCol}

	grid2 := make([]string, 50)
	grid3 := make([]string, 50)
	minRow, minCol = minRow+50, 0
	for i := 0; i < 50; i++ {
		scanner.Scan()
		s := scanner.Text()

		n := len(s)
		s = strings.TrimLeft(s, " ")
		nLeft := n - len(s) + 1
		s = strings.TrimRight(s, " ")
		minCol = nLeft
		grid2[i] = s[:50]
		grid3[i] = s[50:]
	}
	faces[2] = Face{grid2, 5, 3, 4, 1, minRow, minCol}
	faces[3] = Face{grid3, 2, 6, 4, 1, minRow, minCol + 50}

	grid1 := make([]string, 50)
	minRow, minCol = minRow+50, 0
	for i := 0; i < 50; i++ {
		scanner.Scan()
		s := scanner.Text()

		n := len(s)
		s = strings.TrimLeft(s, " ")
		nLeft := n - len(s) + 1
		s = strings.TrimRight(s, " ")
		minCol = nLeft
		grid1[i] = s
	}
	faces[1] = Face{grid1, 5, 3, 2, 6, minRow, minCol}

	scanner.Scan()
	scanner.Scan()

	return faces, scanner.Text()
}
