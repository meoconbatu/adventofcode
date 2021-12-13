package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

// Part1 func
func Part1() {
	inputFileName := "day13/input.txt"
	points, folds := readInput(inputFileName)
	output := part1Core(points, folds)
	fmt.Println(output)
}
func part1Core(points map[Point]struct{}, folds []Point) int {
	return fold(points, folds[0])
}
func fold(points map[Point]struct{}, fold Point) int {
	for p := range points {
		if fold.x == 0 && p.y > fold.y {
			points[Point{p.x, fold.y - (p.y - fold.y)}] = struct{}{}
			delete(points, p)
		}
		if fold.y == 0 && p.x > fold.x {
			points[Point{fold.x - (p.x - fold.x), p.y}] = struct{}{}
			delete(points, p)
		}
	}
	return len(points)
}

func readInput(inputFileName string) (map[Point]struct{}, []Point) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	points := make(map[Point]struct{}, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		var x, y int
		fmt.Sscanf(s, "%d,%d", &x, &y)
		points[Point{x, y}] = struct{}{}
	}
	folds := make([]Point, 0)
	for scanner.Scan() {
		var val int
		n, _ := fmt.Sscanf(scanner.Text(), "fold along x=%d", &val)
		if n > 0 {
			folds = append(folds, Point{val, 0})
		} else {
			fmt.Sscanf(scanner.Text(), "fold along y=%d", &val)
			folds = append(folds, Point{0, val})
		}
	}
	return points, folds
}
