package day9

import (
	"fmt"
	"sort"
)

// Part2 func
func Part2() {
	inputFileName := "day9/input.txt"
	heightmap := readInput(inputFileName)
	output := part2Core(heightmap)
	fmt.Println(output)
}
func part2Core(heightmap [][]int) int {
	sum(heightmap)
	rs := make([]int, 0)
	for i := 0; i < len(lowPoints); i += 2 {
		rs = append(rs, loang(heightmap, lowPoints[i], lowPoints[i+1]))
	}
	sort.Ints(rs)
	return rs[len(rs)-1] * rs[len(rs)-2] * rs[len(rs)-3]
}

var lowPoints = make([]int, 0)

func loang(heightmap [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(heightmap) || y >= len(heightmap[x]) || heightmap[x][y] == -1 ||
		heightmap[x][y] == 9 {
		return 0
	}
	heightmap[x][y] = -1
	return 1 + loang(heightmap, x-1, y) + loang(heightmap, x+1, y) + loang(heightmap, x, y+1) +
		loang(heightmap, x, y-1)
}
