package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day14 type
type Day14 struct{}

// Part1 func
func (d Day14) Part1() {
	grid, minp, maxp := readInput()
	minp.Y = 0

	s := utils.Point{X: 500, Y: 0}

	var rs int
	condfn := func(s, minp, maxp utils.Point) bool {
		return s.X >= minp.X && s.X <= maxp.X && s.Y >= minp.Y && s.Y <= maxp.Y
	}
	for rs = 0; fall(grid, s, minp, maxp, condfn); rs++ {
	}
	fmt.Println(rs)
	// print(grid, minp.X, maxp.X, 0, maxp.Y)
}

// Part2 func
func (d Day14) Part2() {
	grid, minp, maxp := readInput()
	s := utils.Point{X: 500, Y: 0}
	minp.X, minp.Y = 500-(maxp.Y+2), maxp.Y+2
	maxp.X, maxp.Y = 500+(maxp.Y+2), maxp.Y+2
	paint(grid, minp, maxp)

	var rs int
	condfn := func(s, minp, maxp utils.Point) bool {
		return s.Y <= maxp.Y
	}
	for rs = 0; fall(grid, s, minp, maxp, condfn); rs++ {

	}
	fmt.Println(rs)
	// print(grid, 500-(maxp.Y+2), 500+(maxp.Y+2), 0, maxp.Y+2)
}

func fall(grid map[utils.Point]byte, s utils.Point, minp, maxp utils.Point, f func(s, minp, maxp utils.Point) bool) bool {
	for f(s, minp, maxp) {
		ld, rd, d := utils.Point{X: s.X - 1, Y: s.Y + 1}, utils.Point{X: s.X + 1, Y: s.Y + 1}, utils.Point{X: s.X, Y: s.Y + 1}
		if _, ok := grid[d]; !ok {
			s.X, s.Y = d.X, d.Y
		} else if _, ok := grid[ld]; !ok {
			s.X, s.Y = ld.X, ld.Y
		} else if _, ok := grid[rd]; !ok {
			s.X, s.Y = rd.X, rd.Y
		} else {
			if grid[s] == 'o' {
				return false
			}
			grid[s] = 'o'
			return true
		}
	}
	return false
}
func readInput() (map[utils.Point]byte, utils.Point, utils.Point) {
	scanner := utils.NewScanner(14)

	grid := make(map[utils.Point]byte, 0)
	var a, b *utils.Point
	minp, maxp := utils.Point{X: math.MaxInt64, Y: math.MaxInt64}, utils.Point{X: 0, Y: 0}
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		for i, point := range points {
			var x, y int
			fmt.Sscanf(point, "%d,%d", &x, &y)
			minp.X, minp.Y = utils.Min(minp.X, x), utils.Min(minp.Y, y)
			maxp.X, maxp.Y = utils.Max(maxp.X, x), utils.Max(maxp.Y, y)
			if i == 0 {
				a = &utils.Point{X: x, Y: y}
			} else {
				b = &utils.Point{X: x, Y: y}
				paint(grid, *a, *b)
				a = b
			}
		}
	}
	return grid, minp, maxp
}
func print(grid map[utils.Point]byte, minp, maxp utils.Point) {
	for j := minp.Y; j <= maxp.Y; j++ {
		for i := minp.X; i <= maxp.X; i++ {
			if v, ok := grid[utils.Point{X: i, Y: j}]; ok {
				fmt.Printf("%s", string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func paint(grid map[utils.Point]byte, a, b utils.Point) {
	var fx, tx, fy, ty int
	if a.X < b.X {
		fx, tx = a.X, b.X
	} else if a.X > b.X {
		fx, tx = b.X, a.X
	}
	for i := fx; i <= tx && fx < tx; i++ {
		grid[utils.Point{X: i, Y: a.Y}] = '#'
	}
	if a.Y < b.Y {
		fy, ty = a.Y, b.Y
	} else if a.Y > b.Y {
		fy, ty = b.Y, a.Y
	}
	for i := fy; i <= ty && fy < ty; i++ {
		grid[utils.Point{X: a.X, Y: i}] = '#'
	}
}
