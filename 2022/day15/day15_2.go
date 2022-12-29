package day15

import (
	"github.com/meoconbatu/adventofcode/utils"
)

func findDistressPointOptimize(sensorToBeacon map[utils.Point]utils.Point) *utils.Point {
	sensorToBorders := make(map[utils.Point][4]utils.Point)
	for sensor, beacon := range sensorToBeacon {
		distance := getDistance(sensor, beacon)
		l, r, u, d := utils.Point{X: sensor.X - distance, Y: sensor.Y}, utils.Point{X: sensor.X + distance, Y: sensor.Y},
			utils.Point{X: sensor.X, Y: sensor.Y - distance}, utils.Point{X: sensor.X, Y: sensor.Y + distance}
		sensorToBorders[sensor] = [4]utils.Point{l, r, u, d}
	}
	for sensor1, border1s := range sensorToBorders {
		for sensor2, border2s := range sensorToBorders {
			if sensor1 == sensor2 {
				continue
			}
			if p := find(sensorToBeacon, sensor1, border1s, border2s); p != nil {
				return p
			}
		}
	}
	return nil
}
func getDistance(a, b utils.Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}
func find(sensorToBeacon map[utils.Point]utils.Point, sensor utils.Point, border1s, border2s [4]utils.Point) *utils.Point {
	line1s := [][2]utils.Point{{border1s[0], border1s[2]}, {border1s[0], border1s[3]}, {border1s[1], border1s[2]}, {border1s[1], border1s[3]}}
	line2s := [][2]utils.Point{{border2s[0], border2s[2]}, {border2s[0], border2s[3]}, {border2s[1], border2s[2]}, {border2s[1], border2s[3]}}
	for _, line1 := range line1s {
		for _, line2 := range line2s {
			if p := getIntersectionPoint(sensor, line1[0], line1[1], line2[0], line2[1]); p != nil {
				if pp := findDistressFromIntersectionPoint(sensorToBeacon, *p); pp != nil {
					return pp
				}
			}
		}
	}

	return nil
}

func findDistressFromIntersectionPoint(sensorToBeacon map[utils.Point]utils.Point, p utils.Point) *utils.Point {
	directions := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for _, d := range directions {
		cur := utils.Point{X: p.X + d[0], Y: p.Y + d[1]}
		found := false
		for sensor, beacon := range sensorToBeacon {
			distance := getDistance(sensor, beacon)
			if getDistance(sensor, cur) <= distance {
				found = true
				break
			}
		}
		if !found && cur.Y >= MINLINE && cur.Y <= MAXLINE && cur.X >= MINLINE && cur.X <= MAXLINE {
			return &cur
		}
	}
	return nil
}
func getIntersectionPoint(s, x1, x2, x3, x4 utils.Point) *utils.Point {
	a1 := (x2.Y - x1.Y) / (x2.X - x1.X)
	b1 := x1.Y - a1*x1.X
	a2 := (x4.Y - x3.Y) / (x4.X - x3.X)
	b2 := x3.Y - a2*x3.X
	if a1 == a2 {
		return nil
	}
	x := (b2 - b1) / (a1 - a2)
	y := a2*x + b2
	rs := utils.Point{X: x, Y: y}
	if getDistance(s, x1) != getDistance(s, rs) {
		return nil
	}
	return &rs
}
