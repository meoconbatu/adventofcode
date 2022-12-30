package day6

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day6 type
type Day6 struct{}

// Part1 func
func (d Day6) Part1() {
	scanner := utils.NewScanner(6)
	rs := 0
	for scanner.Scan() {
		rs = distinctSubSet(scanner.Text(), 4)
	}
	fmt.Println(rs)
}

func distinctSubSet(s string, lenSubSet int) int {
	chars := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		chars[i] = -1
	}
	n := 0
	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		if chars[val] == -1 || chars[val] < i-n {
			chars[val] = i
			n++
		} else {
			n = utils.Min(i-chars[val], n)
			chars[val] = i
		}
		if n == lenSubSet {
			return i + 1
		}
	}
	return -1
}

// Part2 func
func (d Day6) Part2() {
	scanner := utils.NewScanner(6)
	rs := 0
	for scanner.Scan() {
		rs = distinctSubSet(scanner.Text(), 14)
	}
	fmt.Println(rs)
}
