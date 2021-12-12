package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day11/input.txt"
	energyLevels := readInput(inputFileName)
	output := part1Core(energyLevels)
	fmt.Println(output)
}
func part1Core(energyLevels [][]int) int {
	rs := 0
	for i := 0; i < 100; i++ {
		step(energyLevels)
		rs += countFlash(energyLevels)
	}
	return rs
}

func step(energyLevels [][]int) {
	q := make([]int, 0)
	for i := 0; i < len(energyLevels); i++ {
		for j := 0; j < len(energyLevels[0]); j++ {
			energyLevels[i][j]++
			if energyLevels[i][j] > 9 {
				energyLevels[i][j] = 0
				q = append(q, i, j)
			}
		}
	}
	for len(q) > 0 {
		i, j := q[0], q[1]
		q = q[2:]
		for m := -1; m <= 1; m++ {
			for n := -1; n <= 1; n++ {
				if m+i < 0 || n+j < 0 || m+i >= len(energyLevels) || n+j >= len(energyLevels[0]) || (m == 0 && n == 0) || energyLevels[i+m][j+n] == 0 {
					continue
				}
				energyLevels[i+m][j+n]++
				if energyLevels[i+m][j+n] > 9 {
					energyLevels[i+m][j+n] = 0
					q = append(q, i+m, j+n)
				}
			}
		}
	}
}
func countFlash(energyLevels [][]int) int {
	rs := 0
	for i := 0; i < len(energyLevels); i++ {
		for j := 0; j < len(energyLevels[0]); j++ {
			if energyLevels[i][j] == 0 {
				rs++
			}
		}
	}
	return rs
}

func print(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%d", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func readInput(inputFileName string) [][]int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	energyLevels := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, c := range strings.Split(scanner.Text(), "") {
			var num int
			if _, err := fmt.Sscanf(c, "%d", &num); err == nil {
				row = append(row, num)
			}
		}
		energyLevels = append(energyLevels, row)
	}
	return energyLevels
}
