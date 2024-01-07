package day12

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day12 type
type Day12 struct{}

var dp map[int]int

// Part1 func
func (d Day12) Part1() {
	scanner := utils.NewScanner(12)
	rs := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		nums := utils.ParseIntSlice(parts[1], ",")
		rs += countAllArrangement(parts[0], nums)
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day12) Part2() {
	scanner := utils.NewScanner(12)
	rs := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		left := strings.Repeat(parts[0]+"?", 5)
		left = left[:len(left)-1]

		right := strings.Repeat(parts[1]+",", 5)
		right = right[:len(right)-1]

		nums := utils.ParseIntSlice(right, ",")

		rs += countAllArrangement(left, nums)
	}
	fmt.Println(rs)
}

// countAllArrangement counts all of the different arrangements of operational and broken springs
// that meet the given criteria
func countAllArrangement(s string, nums []int) int {
	dp = make(map[int]int)
	return dfs(s, nums, 0, 0, -1)
}
func dfs(s string, nums []int, idx int, cnum, cidx int) int {
	if idx == len(s) {
		if cidx == len(nums) || (cidx == len(nums)-1 && cnum == nums[cidx]) {
			return 1
		}
		return 0
	}
	key := cnum*110*110 + idx*110 + cidx
	if v, ok := dp[key]; ok {
		return v
	}
	rs := 0
	if s[idx] == '?' {
		if cidx == -1 || (cidx < len(nums) && cnum == nums[cidx]) || cnum == 0 {
			if cidx == -1 || cnum == 0 {
				rs += dfs(s[:idx]+"."+s[idx+1:], nums, idx+1, 0, cidx)
			} else {
				rs += dfs(s[:idx]+"."+s[idx+1:], nums, idx+1, 0, cidx+1)
			}
		}
		if cidx == -1 || (cidx < len(nums) && cnum < nums[cidx]) {
			if cidx == -1 {
				cidx = 0
			}
			rs += dfs(s[:idx]+"#"+s[idx+1:], nums, idx+1, cnum+1, cidx)
		}
	} else if s[idx] == '.' {
		if cidx == -1 || (cidx < len(nums) && cnum == nums[cidx]) || cnum == 0 {
			if cidx == -1 || cnum == 0 {
				rs += dfs(s, nums, idx+1, 0, cidx)
			} else {
				rs += dfs(s, nums, idx+1, 0, cidx+1)
			}
		}
	} else {
		if cidx == -1 || (cidx < len(nums) && cnum < nums[cidx]) {
			if cidx == -1 {
				cidx = 0
			}
			rs += dfs(s, nums, idx+1, cnum+1, cidx)
		}
	}
	dp[key] = rs
	return rs
}
