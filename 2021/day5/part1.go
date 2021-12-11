package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Point type
type Point struct {
	x, y int
}

// Part1 func
func Part1() {
	inputFileName := "day5/input.txt"
	froms, tos := readInput(inputFileName)
	output := part1Core(froms, tos)
	fmt.Println(output)
}
func part1Core(froms, tos []Point) int {
	return countIntersect(froms, tos)
}

// the number of points where at least two lines overlap
func countIntersect(froms, tos []Point) int {
	m := make(map[Point]int)
	rs := 0
	for i := 0; i < len(froms); i++ {
		if froms[i].x != tos[i].x && froms[i].y != tos[i].y {
			continue
		}
		if froms[i].x == tos[i].x {
			for j := min(froms[i].y, tos[i].y); j <= max(froms[i].y, tos[i].y); j++ {
				m[Point{froms[i].x, j}]++
				if m[Point{froms[i].x, j}] == 2 {
					rs++
				}
			}
		} else if froms[i].y == tos[i].y {
			for j := min(froms[i].x, tos[i].x); j <= max(froms[i].x, tos[i].x); j++ {
				m[Point{j, froms[i].y}]++
				if m[Point{j, froms[i].y}] == 2 {
					rs++
				}
			}
		}
	}
	return rs
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 694,732 -> 290,328
func readInput(inputFileName string) ([]Point, []Point) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	froms, tos := make([]Point, 0), make([]Point, 0)

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " -> ")
		var x, y int
		fmt.Sscanf(points[0], "%d,%d", &x, &y)
		froms = append(froms, Point{x, y})
		fmt.Sscanf(points[1], "%d,%d", &x, &y)
		tos = append(tos, Point{x, y})
	}
	return froms, tos
}
