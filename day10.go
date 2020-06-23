package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func day10() {
	content, _ := ioutil.ReadFile("input10.txt")
	fmt.Println(day10Part1(string(content)))
	fmt.Println(day10Part2(string(content), 200))
}

func readMapToPoint(in string) []Point {
	points := make([]Point, 0)
	rows := strings.Split(in, "\n")
	for y, r := range rows {
		rowString := strings.Split(r, "")
		for x, r := range rowString {
			if r == "#" {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func day10Part1(content string) (Point, int) {
	points := readMapToPoint(content)
	seenPoint := make(map[Point][]Point, 0)
	for _, p1 := range points {
		p2s := seenPoint[p1]
		for _, p3 := range points {
			if p1 == p3 {
				continue
			}
			if p2s == nil {
				p2s = append(p2s, p3)
				continue
			}
			var temp *Point
			for i := 0; i < len(p2s); i++ {
				temp = isStraight(&p1, &p2s[i], &p3)
				if temp == &p2s[i] {
					break
				}
				if temp == &p3 {
					p2s[i] = p3
					break
				}
				temp = nil
			}
			if temp == nil {
				p2s = append(p2s, p3)
			}
		}
		seenPoint[p1] = p2s
	}
	maxSeenPoint := 0
	var p Point
	for key, value := range seenPoint {
		if maxSeenPoint < len(value) {
			maxSeenPoint = len(value)
			p = key
		}
	}
	return p, maxSeenPoint
}

var p1 Point

func day10Part2(content string, pointth int) *Point {
	p1, _ = day10Part1(content)
	points := readMapToPoint(content)
	angles := make(map[float64][]Point, 0)

	p2 := Point{p1.x, 0}
	v1 := Point{p2.x - p1.x, p2.y - p1.y}
	for _, p3 := range points {
		if p1 == p3 {
			continue
		}
		v2 := Point{p3.x - p1.x, p3.y - p1.y}
		angle := angle(v1, v2)

		angles[angle] = insertSorted(angles[angle], p3)
	}

	keys := make([]float64, 0, len(angles))
	for k := range angles {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	i := 1
	var pp Point
	for {
		for _, k := range keys {
			if len(angles[k]) > 0 {
				pp = angles[k][0]
				if i == pointth {
					return &pp
				}
				if len(angles[k]) == 1 {
					angles[k] = []Point{}
				} else {
					angles[k] = angles[k][1:]
				}
				i++
			}
		}
		if i == len(points) {
			return nil
		}
	}
}
func insertSorted(s []Point, e Point) []Point {
	i := sort.Search(len(s), func(i int) bool {
		return (s[i].x-p1.x)*(s[i].x-p1.x)+(s[i].y-p1.y)*(s[i].y-p1.y) >
			(e.x-p1.x)*(e.x-p1.x)+(e.y-p1.y)*(e.y-p1.y)
	})
	s = append(s, Point{0, 0})
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func isStraight(p1, p2, p3 *Point) *Point {
	p1p2 := Point{p2.x - p1.x, p2.y - p1.y}
	p1p3 := Point{p3.x - p1.x, p3.y - p1.y}
	if p1p2.x == 0 && p1p3.x == 0 {
		if p1p2.y/p1p3.y >= 0 {
			if math.Abs(float64(p1p2.y)) >= math.Abs(float64(p1p3.y)) {
				return p3
			}
			return p2
		}
		return p1
	}
	if p1p2.y == 0 && p1p3.y == 0 {
		if p1p2.x/p1p3.x >= 0 {
			if math.Abs(float64(p1p2.x)) >= math.Abs(float64(p1p3.x)) {
				return p3
			}
			return p2
		}
		return p1
	}
	if p1p2.x*p1p3.y == p1p2.y*p1p3.x {
		if float64(p1p2.x)/float64(p1p3.x) >= 0 {
			if math.Abs(float64(p1p2.x)) >= math.Abs(float64(p1p3.x)) ||
				math.Abs(float64(p1p2.y)) >= math.Abs(float64(p1p3.y)) {
				return p3
			}
			return p2
		}
		return p1
	}
	return nil
}
func angle(v1, v2 Point) float64 {
	angle := math.Atan2(float64(v1.x*v2.y-v1.y*v2.x), float64(v1.x*v2.x+v1.y*v2.y))
	if angle < 0 {
		angle = (math.Pi*2 + angle)
	}
	angle = angle * 360 / (2 * math.Pi)
	return angle
}
