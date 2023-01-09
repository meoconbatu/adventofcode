package day22

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day22 struct {
}

// Part1 func
func (d Day22) Part1() {
	readInput()
}

// Part2 func
func (d Day22) Part2() {
	readInput2()
}

type Point struct {
	x, y, z int
}

func readInput() {
	scanner := utils.NewScanner(22)
	points := make(map[Point]struct{})
	for scanner.Scan() {
		var set string
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(scanner.Text(), "%s x=%d..%d,y=%d..%d,z=%d..%d\n", &set, &x1, &x2, &y1, &y2, &z1, &z2)
		for x := x1; x <= x2; x++ {
			if x < -50 || x > 50 {
				continue
			}
			for y := y1; y <= y2; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for z := z1; z <= z2; z++ {
					if z < -50 || z > 50 {
						continue
					}
					p := Point{x, y, z}
					if set == "on" {
						points[p] = struct{}{}
					} else {
						delete(points, p)
					}
				}
			}
		}
	}
	fmt.Println(len(points))
}

func readInput2() {
	scanner := utils.NewScanner(22)
	rs := 0
	intersections := make([][]int, 0)
	for scanner.Scan() {
		var set string
		var x1, x2, y1, y2, z1, z2 int
		s := scanner.Text()
		fmt.Sscanf(s, "%s x=%d..%d,y=%d..%d,z=%d..%d\n", &set, &x1, &x2, &y1, &y2, &z1, &z2)
		isOn := false
		if set == "on" {
			isOn = true
		}
		rs += fn(&intersections, isOn, []int{x1, x2, y1, y2, z1, z2})
	}
	fmt.Println(rs)
}

func fn(intersections *[][]int, isOn bool, rectangle []int) int {
	rs := area(rectangle)
	if !isOn {
		rs = 0
	}
	newIntersections := make([][]int, 0)
	for _, intersection := range *intersections {
		newIntersection := getIntersection(intersection[1:], rectangle)
		if newIntersection == nil {
			continue
		}
		rs += intersection[0] * -1 * area(*newIntersection)
		newIntersections = append(newIntersections, append([]int{intersection[0] * -1}, *newIntersection...))
	}
	*intersections = append(*intersections, newIntersections...)
	if isOn {
		*intersections = append(*intersections, append([]int{1}, rectangle...))
	}
	return rs
}

func getIntersection(a, b []int) *[]int {
	p := []int{utils.Max(a[0], b[0]), utils.Min(a[1], b[1]), utils.Max(a[2], b[2]), utils.Min(a[3], b[3]), utils.Max(a[4], b[4]), utils.Min(a[5], b[5])}
	if p[0] > p[1] || p[2] > p[3] || p[4] > p[5] {
		return nil
	}
	return &p
}
func area(x []int) int {
	return (x[1] - x[0] + 1) * (x[3] - x[2] + 1) * (x[5] - x[4] + 1)
}
