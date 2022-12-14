package day14

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day14 type
type Day14 struct{}

type Point struct {
	x, y int
}

// Part1 func
func (d Day14) Part1() {
	grid, minp, maxp := readInput()
	minp.y = 0

	s := Point{500, 0}

	var rs int
	for rs = 0; fall(grid, s, minp, maxp, func(s, minp, maxp Point) bool {
		return s.x >= minp.x && s.x <= maxp.x && s.y >= minp.y && s.y <= maxp.y
	}); rs++ {
	}
	fmt.Println(rs)
	// print(grid, minp.x, maxp.x, 0, maxp.y)
}

// Part2 func
func (d Day14) Part2() {
	grid, minp, maxp := readInput()
	s := Point{500, 0}
	minp.x, minp.y = 500-(maxp.y+2), maxp.y+2
	maxp.x, maxp.y = 500+(maxp.y+2), maxp.y+2
	paint(grid, minp, maxp)

	var rs int
	for rs = 0; fall(grid, s, minp, maxp, func(s, minp, maxp Point) bool {
		return s.y <= maxp.y
	}); rs++ {

	}
	fmt.Println(rs)
	// print(grid, 500-(maxp.y+2), 500+(maxp.y+2), 0, maxp.y+2)
}

func fall(grid map[Point]byte, s Point, minp, maxp Point, f func(s, minp, maxp Point) bool) bool {
	for f(s, minp, maxp) {
		ld, rd, d := Point{s.x - 1, s.y + 1}, Point{s.x + 1, s.y + 1}, Point{s.x, s.y + 1}
		if _, ok := grid[d]; !ok {
			s.x, s.y = d.x, d.y
		} else if _, ok := grid[ld]; !ok {
			s.x, s.y = ld.x, ld.y
		} else if _, ok := grid[rd]; !ok {
			s.x, s.y = rd.x, rd.y
		} else {
			if grid[Point{s.x, s.y}] == 'o' {
				return false
			}
			grid[Point{s.x, s.y}] = 'o'
			return true
		}
	}
	return false
}
func readInput() (map[Point]byte, Point, Point) {
	scanner := utils.NewScanner(14)

	grid := make(map[Point]byte, 0)
	var a, b *Point
	minp, maxp := Point{math.MaxInt64, math.MaxInt64}, Point{0, 0}
	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		var x, y int
		for i, point := range points {
			fmt.Sscanf(point, "%d,%d", &x, &y)
			minp.x, minp.y = utils.Min(minp.x, x), utils.Min(minp.y, y)
			maxp.x, maxp.y = utils.Max(maxp.x, x), utils.Max(maxp.y, y)
			if i == 0 {
				a = &Point{x, y}
			} else {
				b = &Point{x, y}
				paint(grid, *a, *b)
				a = b
			}
		}
	}
	return grid, minp, maxp
}
func print(grid map[Point]byte, minp, maxp Point) {
	for j := minp.y; j <= maxp.y; j++ {
		for i := minp.x; i <= maxp.x; i++ {
			if v, ok := grid[Point{i, j}]; ok {
				fmt.Printf("%s", string(v))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func paint(grid map[Point]byte, a, b Point) {
	var fx, tx, fy, ty int
	if a.x < b.x {
		fx, tx = a.x, b.x
	} else if a.x > b.x {
		fx, tx = b.x, a.x
	}
	for i := fx; i <= tx && fx < tx; i++ {
		grid[Point{i, a.y}] = '#'
	}
	if a.y < b.y {
		fy, ty = a.y, b.y
	} else if a.y > b.y {
		fy, ty = b.y, a.y
	}
	for i := fy; i <= ty && fy < ty; i++ {
		grid[Point{a.x, i}] = '#'
	}
}
