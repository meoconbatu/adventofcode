package day12

import (
	"fmt"
	"strings"
)

// Part2 func
func Part2() {
	inputFileName := "day12/input.txt"
	edges := readInput(inputFileName)
	output := part2Core(edges)
	fmt.Println(output)
}
func part2Core(edges map[string][]string) int {
	return findPath2(edges)
}
func findPath2(edges map[string][]string) int {
	numPath := 0
	dfs(edges, start, end, []string{}, &numPath, checkValidPath2)
	return numPath
}

func checkValidPath2(path []string) bool {
	m := make(map[string]int)
	twice := false
	for _, u := range path {
		m[u]++
		if strings.ToLower(u) != u {
			continue
		}
		if m[u] > 2 {
			return false
		}
		if m[u] == 2 {
			if twice || u == start || u == end {
				return false
			}
			twice = true
		}
	}
	return true
}
