package day19

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/2019/day2"
	"github.com/meoconbatu/adventofcode/utils"
)

// Day19 type
type Day19 struct{}

// Part1 func
func (d Day19) Part1() {
	nums := readInput()
	rs := 0
	copyNums := make([]int, len(nums))
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			copy(copyNums, nums)
			var output int
			ic := day2.Init(copyNums, []int{x, y}, &output)
			ic.Run()
			if output == 1 {
				// fmt.Print("#")
				rs++
			} else {
				// fmt.Print(".")
			}
		}
		// fmt.Println()
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day19) Part2() {
	nums := readInput()
	n := 100
	startX := 0
	for y := 0; ; y++ {
		for x := startX; x < startX+3; x++ {
			if do(nums, x, y) == 1 {
				startX = x
				break
			}
		}
		if do(nums, startX+n-1, y-n+1) == 1 {
			fmt.Println(startX*10000 + (y - n + 1))
			return
		}
	}
}

func do(nums []int, i, j int) int {
	copyNums := utils.CopySlice(nums)
	var output int
	ic := day2.Init(copyNums, []int{i, j}, &output)
	ic.Run()
	return output
}
func readInput() []int {
	scanner := utils.NewScanner(19)
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
