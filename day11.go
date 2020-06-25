package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	up = iota
	left
	down
	right
)

const (
	// BLACK var
	BLACK = iota
	// WHITE var
	WHITE
)

type direction int
type robot struct {
	currentPositition Point
	face              direction
	color             int
}

func day11() {
	file, err := os.Open("input11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), ",")

	originalIns := make([]int64, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.ParseInt(ins[i], 10, 64)
	}
	robot := robot{Point{0, 0}, up, BLACK}
	out, _ := day11Part1(originalIns, robot)
	fmt.Println(out)

	robot.color = WHITE
	_, points := day11Part1(originalIns, robot)
	day11Part2(points)
}

func day11Part1(signals []int64, robot robot) (int, map[Point]int) {
	done, donez := make(chan struct{}), make(chan struct{})
	in, out := &Buffer{}, &Buffer{}

	points := make(map[Point]int, 0)

	in.WriteString(strconv.Itoa(robot.color))
	go func() {
		process(signals, in, out)
		done <- struct{}{}
	}()

	go func() {
		i := 1
		for {
			select {
			case <-done:
				donez <- struct{}{}
				break
			default:
				var v int
				if n, _ := fmt.Fscan(out, &v); n == 0 {
					continue
				}
				if i == 1 {
					robot.color = v
					points[robot.currentPositition] = robot.color
					i++
				} else {
					i = 1
					if v == 0 {
						switch robot.face {
						case up:
							robot.currentPositition = Point{robot.currentPositition.x - 1, robot.currentPositition.y}
							robot.face = left
						case left:
							robot.currentPositition = Point{robot.currentPositition.x, robot.currentPositition.y - 1}
							robot.face = down
						case down:
							robot.currentPositition = Point{robot.currentPositition.x + 1, robot.currentPositition.y}
							robot.face = right
						case right:
							robot.currentPositition = Point{robot.currentPositition.x, robot.currentPositition.y + 1}
							robot.face = up
						}
					} else {
						switch robot.face {
						case up:
							robot.currentPositition = Point{robot.currentPositition.x + 1, robot.currentPositition.y}
							robot.face = right
						case left:
							robot.currentPositition = Point{robot.currentPositition.x, robot.currentPositition.y + 1}
							robot.face = up
						case down:
							robot.currentPositition = Point{robot.currentPositition.x - 1, robot.currentPositition.y}
							robot.face = left
						case right:
							robot.currentPositition = Point{robot.currentPositition.x, robot.currentPositition.y - 1}
							robot.face = down
						}
					}
					in.WriteString(strconv.Itoa(points[robot.currentPositition]))
				}

			}
		}
	}()
	<-donez
	result := len(points)

	return result, points
}
func day11Part2(points map[Point]int) {
	minX, maxX := math.MaxInt64, math.MinInt64
	minY, maxY := math.MaxInt64, math.MinInt64
	for key := range points {
		if key.x < minX {
			minX = key.x
		} else if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		} else if key.y > maxY {
			maxY = key.y
		}
	}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if value, ok := points[Point{x, y}]; ok && value == WHITE {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
}
