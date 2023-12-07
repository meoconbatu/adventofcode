package day2

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day2 type
type Day2 struct{}

// Part1 func
func (d Day2) Part1() {
	scanner := utils.NewScanner(2)
	rs := 0
	for scanner.Scan() {
		var id int
		p := strings.Split(scanner.Text(), ": ")
		fmt.Sscanf(p[0], "Game %d", &id)
		if isPossibleGame(p[1]) {
			rs += id
		}
	}
	fmt.Println(rs)
}

var rules = map[string]int{"red": 12, "green": 13, "blue": 14}

func isPossibleGame(s string) bool {
	sets := strings.Split(s, ";")
	for _, set := range sets {
		numCubes := strings.Split(set, ", ")
		for _, cube := range numCubes {
			var num int
			var typ string
			fmt.Sscanf(cube, "%d %s\n", &num, &typ)
			if num > rules[typ] {
				return false
			}
		}
	}
	return true
}

// Part2 func
func (d Day2) Part2() {
	scanner := utils.NewScanner(2)
	rs := 0
	for scanner.Scan() {
		p := strings.Split(scanner.Text(), ": ")
		rs += powerOfMinimumSet(p[1])
	}
	fmt.Println(rs)
}

func powerOfMinimumSet(s string) int {
	sets := strings.Split(s, "; ")
	m := make(map[string]int)
	for _, set := range sets {
		numCubes := strings.Split(set, ", ")
		for _, cube := range numCubes {
			var num int
			var typ string
			fmt.Sscanf(cube, "%d %s\n", &num, &typ)
			m[typ] = utils.Max(m[typ], num)
		}
	}
	return utils.Max(m["red"], 1) * utils.Max(m["green"], 1) * utils.Max(m["blue"], 1)
}
