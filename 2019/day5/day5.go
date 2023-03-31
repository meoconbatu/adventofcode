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
	input := 1
	var output int
	ic := day2.Init(nums, []int{input}, &output)
	ic.Run()
	fmt.Println(output)
}

// Part2 func
func (d Day5) Part2() {
	nums := readInput()
	input := 5
	var output int
	ic := day2.Init(nums, []int{input}, &output)
	ic.Run()
	fmt.Println(output)
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
