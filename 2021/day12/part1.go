package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	start = "start"
	end   = "end"
)

// Part1 func
func Part1() {
	inputFileName := "day12/input.txt"
	edges := readInput(inputFileName)
	output := part1Core(edges)
	fmt.Println(output)
}
func part1Core(edges map[string][]string) int {
	return findPath(edges)
}
func findPath(edges map[string][]string) int {
	numPath := 0
	dfs(edges, start, end, []string{}, &numPath, checkValidPath)
	return numPath
}
func dfs(edges map[string][]string, v, target string, path []string, numPath *int, f func([]string) bool) {
	if v == target {
		*numPath++
		return
	}
	if !f(append(path, v)) {
		return
	}
	for _, u := range edges[v] {
		dfs(edges, u, target, append(path, v), numPath, f)
	}
}

func checkValidPath(path []string) bool {
	v := path[len(path)-1]
	if strings.ToLower(v) == v && visitedTimes(path) > 1 {
		return false
	}
	return true
}
func visitedTimes(path []string) int {
	v := path[len(path)-1]
	count := 0
	for _, u := range path {
		if v == u {
			count++
		}
	}
	return count
}

func readInput(inputFileName string) map[string][]string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	edges := make(map[string][]string)
	for scanner.Scan() {
		v := strings.Split(scanner.Text(), "-")
		v1, v2 := v[0], v[1]
		edges[v1] = append(edges[v1], v2)
		edges[v2] = append(edges[v2], v1)
	}
	return edges
}
