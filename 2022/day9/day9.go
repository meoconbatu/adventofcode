package day9

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day9 type
type Day9 struct{}

type Point struct {
	x, y int
}

// Part1 func
func (d Day9) Part1() {
	scanner := utils.NewScanner(9)
	positions := make(map[Point]struct{})
	knots := make([]Point, 2)
	for scanner.Scan() {
		var direction string
		var numStep int
		fmt.Sscanf(scanner.Text(), "%s %d\n", &direction, &numStep)
		move(positions, knots, direction, numStep)
	}
	fmt.Printf("%d\n", len(positions))
}

// Part2 func
func (d Day9) Part2() {
	scanner := utils.NewScanner(9)
	positions := make(map[Point]struct{})
	knots := make([]Point, 10)
	for scanner.Scan() {
		var direction string
		var numStep int
		fmt.Sscanf(scanner.Text(), "%s %d\n", &direction, &numStep)
		move(positions, knots, direction, numStep)
	}
	fmt.Printf("%d\n", len(positions))
}

func move(positions map[Point]struct{}, knots []Point, direction string, numStep int) {
	for i := 0; i < numStep; i++ {
		switch direction {
		case "U":
			knots[0].y++
		case "D":
			knots[0].y--
		case "L":
			knots[0].x--
		case "R":
			knots[0].x++
		}
		for j := 0; j < len(knots)-1; j++ {
			head, tail := &knots[j], &knots[j+1]
			fn(head, tail)
		}
		positions[knots[len(knots)-1]] = struct{}{}
	}
}

func fn(head, tail *Point) {
	if head.x == tail.x {
		if head.y-tail.y == 2 {
			tail.y++
		} else if tail.y-head.y == 2 {
			tail.y--
		}
	}
	if head.y == tail.y {
		if head.x-tail.x == 2 {
			tail.x++
		} else if tail.x-head.x == 2 {
			tail.x--
		}
	}
	if utils.Abs(head.x-tail.x)+utils.Abs(head.y-tail.y) >= 3 {
		if tail.x < head.x {
			tail.x++
		} else if tail.x > head.x {
			tail.x--
		}
		if tail.y < head.y {
			tail.y++
		} else if tail.y > head.y {
			tail.y--
		}
	}
}
