package day5

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/2019/day2"
	"github.com/meoconbatu/adventofcode/utils"
)

// Day5 type
type Day5 struct{}

// Part1 func
func (d Day5) Part1() {
	nums := readInput()
	systempID := 1
	out := day2.RunProgram(nums, &systempID)
	fmt.Println(out)
}

// Part2 func
func (d Day5) Part2() {
	nums := readInput()
	systempID := 5
	out := day2.RunProgram(nums, &systempID)
	fmt.Println(out)
}

func readInput() []int {
	scanner := utils.NewScanner(5)
	rs := make([]int, 0)
	for scanner.Scan() {
		numStrs := strings.Split(scanner.Text(), ",")
		var num int
		for _, numStr := range numStrs {
			fmt.Sscanf(numStr, "%d", &num)
			rs = append(rs, num)
		}
	}
	return rs
}
