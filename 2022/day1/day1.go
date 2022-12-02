package day1

import (
	"fmt"
	"sort"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day1 type
type Day1 struct{}

// Part1 func
func (d Day1) Part1() {
	scanner := utils.NewScanner(1)
	rs := 0
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			rs = utils.Max(rs, sum)
			sum = 0
			continue
		}
		var num int
		fmt.Sscanf(scanner.Text(), "%d\n", &num)
		sum += num
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day1) Part2() {
	scanner := utils.NewScanner(1)
	nums := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			nums = append(nums, sum)
			sum = 0
			continue
		}
		var num int
		fmt.Sscanf(scanner.Text(), "%d\n", &num)
		sum += num
	}
	sort.Ints(nums)
	n := len(nums)
	fmt.Println(nums[n-1] + nums[n-2] + nums[n-3])
}
