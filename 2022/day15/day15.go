package day15

import (
	"fmt"
	"sort"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day15 type
type Day15 struct{}

// Part1 func
func (d Day15) Part1() {
	sensorToBeacon := readInput()
	line := 2000000
	xs := part(sensorToBeacon, line)

	rs := 0
	for _, x := range xs {
		rs += x.Y - x.X + 1
	}
	visitedBeacon := make(map[utils.Point]bool)
	for _, b := range sensorToBeacon {
		for _, x := range xs {
			if !visitedBeacon[b] && b.Y == line && b.X >= x.X && b.X <= x.Y {
				visitedBeacon[b] = true
				rs--
				break
			}
		}
	}
	fmt.Println(rs)
}
func part(sensorToBeacon map[utils.Point]utils.Point, line int) []utils.Point {
	xs := make([]utils.Point, 0)
	for s, b := range sensorToBeacon {
		x := fn(s, b, line)
		if x == nil {
			continue
		}
		xs = append(xs, *x)
	}
	if len(xs) == 0 {
		return nil
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].X < xs[j].X || (xs[i].X == xs[j].X && xs[i].Y < xs[j].Y)
	})
	reunion(&xs)
	return xs
}
func reunion(xs *[]utils.Point) {
	for i := 0; i < len(*xs); i++ {
		for j := i + 1; j < len((*xs)); j++ {
			if (*xs)[i].Y >= (*xs)[j].X && (*xs)[i].X <= (*xs)[j].Y {
				(*xs)[i].X = utils.Min((*xs)[i].X, (*xs)[j].X)
				(*xs)[i].Y = utils.Max((*xs)[i].Y, (*xs)[j].Y)
				copy((*xs)[j:], (*xs)[j+1:])
				(*xs) = (*xs)[:len((*xs))-1]
				j--
			}
		}
	}
}

// Part2 func
func (d Day15) Part2() {
	maxline := 4000000
	sensorToBeacon := readInput()
	for line := 0; line <= maxline; line++ {
		xs := part(sensorToBeacon, line)
		if len(xs) == 2 {
			fmt.Println((xs[1].X-1)*maxline + line)
			break
		}
	}
}

func fn(s, b utils.Point, liney int) *utils.Point {
	distance := utils.Abs(s.X-b.X) + utils.Abs(s.Y-b.Y)
	l, r, u, d := utils.Point{X: s.X - distance, Y: s.Y}, utils.Point{X: s.X + distance, Y: s.Y},
		utils.Point{X: s.X, Y: s.Y - distance}, utils.Point{X: s.X, Y: s.Y + distance}
	if u.Y <= liney && s.Y >= liney {
		return findIntersection(u, l, r, liney)
	} else if s.Y <= liney && d.Y >= liney {
		return findIntersection(d, l, r, liney)
	}
	return nil
}

func findIntersection(d, l, r utils.Point, liney int) *utils.Point {
	x1 := ((l.Y-d.Y)*d.X - (l.X-d.X)*d.Y + (l.X-d.X)*liney) / (l.Y - d.Y)
	x2 := ((r.Y-d.Y)*d.X - (r.X-d.X)*d.Y + (r.X-d.X)*liney) / (r.Y - d.Y)
	return &utils.Point{X: x1, Y: x2}
}

// Sensor at x=3890859, y=2762958: closest beacon is at x=4037927, y=2985317
func readInput() map[utils.Point]utils.Point {
	scanner := utils.NewScanner(15)
	grid := make(map[utils.Point]utils.Point)
	for scanner.Scan() {
		var xs, ys, xb, yb int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		grid[utils.Point{X: xs, Y: ys}] = utils.Point{X: xb, Y: yb}
	}
	return grid
}
