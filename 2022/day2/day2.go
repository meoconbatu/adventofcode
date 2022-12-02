package day2

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day2 type
type Day2 struct{}

// Part1 func
func (d Day2) Part1() {
	scanner := utils.NewScanner(2)
	rs := 0
	scores := [3][3]int{{3, 6, 0}, {0, 3, 6}, {6, 0, 3}}
	for scanner.Scan() {
		var opponent, me byte
		fmt.Sscanf(scanner.Text(), "%c %c\n", &opponent, &me)
		i, j := opponent-'A', me-'X'
		rs += scores[i][j] + int(j) + 1
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day2) Part2() {
	scanner := utils.NewScanner(2)
	rs := 0
	scores := [3][3]int{{3, 1, 2}, {1, 2, 3}, {2, 3, 1}}
	for scanner.Scan() {
		var opponent, me byte
		fmt.Sscanf(scanner.Text(), "%c %c\n", &opponent, &me)
		i, j := opponent-'A', me-'X'
		rs += scores[i][j] + int(j)*3
	}
	fmt.Println(rs)
}
