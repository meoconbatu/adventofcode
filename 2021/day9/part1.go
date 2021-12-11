package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day9/input.txt"
	heightmap := readInput(inputFileName)
	output := part1Core(heightmap)
	fmt.Println(output)
}
func part1Core(heightmap [][]int) int {
	return sum(heightmap)
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func sum(heightmap [][]int) int {
	rs := 0
	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			isLowPoint := true
			for _, direciton := range directions {
				if direciton[0]+i < 0 || direciton[1]+j < 0 || direciton[0]+i >= len(heightmap) ||
					direciton[1]+j >= len(heightmap[i]) {
					continue
				}
				if heightmap[i][j] >= heightmap[i+direciton[0]][j+direciton[1]] {
					isLowPoint = false
				}
			}

			if isLowPoint {
				rs += heightmap[i][j] + 1
				lowPoints = append(lowPoints, i, j)
			}
		}
	}
	return rs
}

func readInput(inputFileName string) [][]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	heightmap := make([][]int, 0)
	for scanner.Scan() {
		nums := make([]int, 0)
		for _, c := range strings.Split(scanner.Text(), "") {
			var num int
			fmt.Sscanf(c, "%d", &num)
			nums = append(nums, num)
		}
		heightmap = append(heightmap, nums)
	}
	return heightmap
}
