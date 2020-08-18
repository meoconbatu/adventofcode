package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	north = iota + 1
	south
	west
	east
)
const (
	barrier = iota
	free
	oxygenSystem
)

type position struct {
	pos    Point
	status int
	enable bool
}

func day15() {
	originalIns := readInputDay15()

	m, p := day15Part1(originalIns)

	day15Part2(m, p)
	// day15Part1(originalIns)
}
func readInputDay15() []int64 {
	file, err := os.Open("input15.txt")
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
	return originalIns
}

var temps []struct {
	newPosition   Point
	direction     int
	directionBack int
} = []struct {
	newPosition   Point
	direction     int
	directionBack int
}{
	{Point{0, 1}, north, south},
	{Point{0, -1}, south, north},
	{Point{1, 0}, east, west},
	{Point{-1, 0}, west, east},
}

func readOutputFromBuffer(buf *Buffer) int {
	var v int
	for {
		if n, _ := fmt.Fscan(buf, &v); n == 0 {
			continue
		}
		return v
	}
}
func day15Part1(ins []int64) (map[Point]int, Point) {
	in, out := &Buffer{}, &Buffer{}

	initialPos := position{Point{0, 0}, free, true}

	positions := make([]position, 0)
	positions = append(positions, initialPos)

	points := make(map[Point][]int)
	points[Point{0, 0}] = nil

	done := make(chan Point)

	var oxyPos Point

	noBarrierPoints := make(map[Point]int)
	noBarrierPoints[Point{0, 0}] = free

	go func() {
		i := 0
		for {
			i++
			canmove := false
			for i, p := range positions {
				if !p.enable {
					continue
				}
				// go to p
				steps := points[p.pos]
				for i := 0; i < len(steps); i++ {
					in.WriteString(strconv.Itoa(steps[i]))
					readOutputFromBuffer(out)
				}
				// go around
				for _, t := range temps {
					newPos := Point{t.newPosition.x + p.pos.x, t.newPosition.y + p.pos.y}
					if s, ok := points[newPos]; !ok {
						in.WriteString(strconv.Itoa(t.direction))
						v := readOutputFromBuffer(out)

						enable := false
						if v != barrier {
							enable = true
						}

						newPosition := position{newPos, v, enable}
						positions = append(positions, newPosition)

						newsteps := make([]int, len(steps)+1)
						copy(newsteps, steps)
						newsteps[len(newsteps)-1] = t.direction

						points[newPos] = newsteps

						if v == oxygenSystem {
							fmt.Println(len(newsteps))
							oxyPos = newPos
						}
						if v != barrier {
							noBarrierPoints[newPos] = v
							canmove = true
							in.WriteString(strconv.Itoa(t.directionBack))
							readOutputFromBuffer(out)
						}
					} else {
						newsteps := make([]int, len(steps)+1)
						copy(newsteps, steps)
						newsteps[len(newsteps)-1] = t.direction
						if len(s) > len(newsteps) {
							points[newPos] = newsteps
						}
					}
				}

				// go back to root
				for i := len(steps) - 1; i >= 0; i-- {
					var k int
					switch steps[i] {
					case north:
						k = south
					case south:
						k = north
					case east:
						k = west
					case west:
						k = east
					}
					in.WriteString(strconv.Itoa(k))
					readOutputFromBuffer(out)
				}
				positions[i].enable = false
			}
			if !canmove {
				done <- oxyPos
				return
			}
		}
	}()
	go func() {
		process(ins, in, out)
	}()

	return noBarrierPoints, <-done
}
func day15Part2(m map[Point]int, p Point) {
	pointDoneLst := make([]Point, 0)
	pointDoneLst = append(pointDoneLst, p)
	i := 0
	for {
		i++
		for _, p := range pointDoneLst {
			for _, t := range temps {
				newPos := Point{t.newPosition.x + p.x, t.newPosition.y + p.y}
				if v, ok := m[newPos]; ok {
					if v == free {
						m[newPos] = oxygenSystem
						pointDoneLst = append(pointDoneLst, newPos)
					}
				}
			}
		}
		if len(pointDoneLst) == len(m) {
			fmt.Println(i)
			return
		}
	}
}
