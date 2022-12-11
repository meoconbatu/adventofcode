package day10

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day10 type
type Day10 struct{}

type Point struct {
	x, cycleth int
}

var points = make([]Point, 0)

// Part1 func
func (d Day10) Part1() {
	scanner := utils.NewScanner(10)
	points = append(points, Point{1, 0})
	target := []int{20, 60, 100, 140, 180, 220}
	rs := 0
	for scanner.Scan() {
		var instruction string
		var value int
		fmt.Sscanf(scanner.Text(), "%s %d\n", &instruction, &value)
		exec(&points, instruction, value)
	}
	itarget := 0
	for i := 0; i < len(points)-1 && itarget < len(target); i++ {
		if points[i].cycleth <= target[itarget] && points[i+1].cycleth > target[itarget] {
			rs += target[itarget] * points[i].x
			itarget++
		} else if points[i+1].cycleth == target[itarget] {
			rs += target[itarget] * points[i].x
			itarget++
		}
	}
	fmt.Printf("%d\n", rs)
}
func exec(points *[]Point, instruction string, value int) {
	n := len(*points)
	newx := (*points)[n-1].x + value
	cycleth := 2
	if instruction == "noop" {
		cycleth = 1
	}
	newcycleth := (*points)[n-1].cycleth + cycleth
	(*points) = append((*points), Point{newx, newcycleth})
}

// Part2 func
func (d Day10) Part2() {
	d.Part1()
	ip := 0
	for i := points[ip].cycleth; i < points[len(points)-1].cycleth; i++ {
		if i%40 >= points[ip].x-1 && i%40 <= points[ip].x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if ip+1 < len(points) && i+1 == points[ip+1].cycleth {
			ip++
		}
		if i%40 == 39 {
			fmt.Println()
		}
	}
}
