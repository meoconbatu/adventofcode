package day18

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day18 type
type Day18 struct{}

// Part1 func
func (d Day18) Part1() {
	cubes := readInput()
	fmt.Println(totalSurface(cubes))
}
func totalSurface(cubes [][3]int) int {
	ds := [][3]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {1, 1, 0}, {1, 0, 1}, {0, 1, 1}, {1, 1, 1}, {0, 0, 0}}
	rs := 0
	ps := make(map[[3]int][]int)
	n := 0
	for i, cube := range cubes {
		cubeToCount := make(map[int]int)
		for _, d := range ds {
			p := [3]int{cube[0] + d[0], cube[1] + d[1], cube[2] + d[2]}
			if _, ok := ps[p]; ok {
				for _, c := range ps[p] {
					cubeToCount[c]++
				}
			}
			ps[p] = append(ps[p], i)
		}
		if n == 0 {
			n = len(ps)
			continue
		}
		countToCube := make(map[int][]int)
		for c, count := range cubeToCount {
			countToCube[count] = append(countToCube[count], c)
		}
		rs += len(countToCube[4])
		n = len(ps)
	}
	return len(cubes)*6 - rs*2
}

// Part2 func
func (d Day18) Part2() {
	cubes := readInput()
	ps := getInsideCubes(cubes)
	fmt.Println(totalSurface(cubes) - totalSurface(ps))
}
func getInsideCubes(cubes [][3]int) [][3]int {
	cubeMap := make(map[[3]int]struct{})
	for _, c := range cubes {
		cubeMap[c] = struct{}{}
	}
	ps := make([][3]int, 0)
	for i := 0; i < 22; i++ {
		for j := 0; j < 22; j++ {
			for k := 0; k < 22; k++ {
				p := [3]int{i, j, k}
				if _, ok := cubeMap[p]; ok {
					continue
				}
				if isInside(cubes, p) && adjacentInside(cubes, cubeMap, p) {
					ps = append(ps, p)
				}
			}
		}
	}
	return ps
}

func adjacentInside(cubes [][3]int, cubeMap map[[3]int]struct{}, p [3]int) bool {
	dd := [][3]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {-1, 0, 0}, {0, 0, -1}, {0, -1, 0}}
	for _, d := range dd {
		a := [3]int{p[0] + d[0], p[1] + d[1], p[2] + d[2]}
		if _, ok := cubeMap[a]; ok {
			continue
		}
		if !isInside(cubes, [3]int{p[0] - 1, p[1], p[2]}) {
			return false
		}
	}
	return true
}
func isInside(cubes [][3]int, p [3]int) bool {
	rs := 0
	for _, c := range cubes {
		if c[0] < p[0] && c[1] == p[1] && c[2] == p[2] {
			rs |= 1 << 0
		} else if c[0] > p[0] && c[1] == p[1] && c[2] == p[2] {
			rs |= 1 << 1
		} else if c[0] == p[0] && c[1] < p[1] && c[2] == p[2] {
			rs |= 1 << 2
		} else if c[0] == p[0] && c[1] > p[1] && c[2] == p[2] {
			rs |= 1 << 3
		} else if c[0] == p[0] && c[1] == p[1] && c[2] < p[2] {
			rs |= 1 << 4
		} else if c[0] == p[0] && c[1] == p[1] && c[2] > p[2] {
			rs |= 1 << 5
		}
	}
	return rs == 0b111111
}

// 7,16,15
func readInput() [][3]int {
	scanner := utils.NewScanner(18)
	rs := make([][3]int, 0)
	for scanner.Scan() {
		var x, y, z int
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &x, &y, &z)
		rs = append(rs, [3]int{x, y, z})
	}
	return rs
}
