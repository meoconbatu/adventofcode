package day4

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day4 type
type Day4 struct{}

// Part1 func
func (d Day4) Part1() {
	scanner := utils.NewScanner(4)
	rs := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if isFullyContain(a, b, c, d) {
			rs++
		}
	}
	fmt.Println(rs)
}

func isFullyContain(a, b, c, d int) bool {
	if (a <= c && b >= d) || (c <= a && d >= b) {
		return true
	}
	return false
}

// Part2 func
func (d Day4) Part2() {
	scanner := utils.NewScanner(4)
	rs := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if isRangeContain(a, b, c, d) {
			rs++
		}
	}
	fmt.Println(rs)
}
func isRangeContain(a, b, c, d int) bool {
	if c <= b && d >= a {
		return true
	}
	return false
}
