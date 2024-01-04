package day4

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day4 type
type Day4 struct{}

// Part1 func
func (d Day4) Part1() {
	scanner := utils.NewScanner(4)
	rs := 0
	for scanner.Scan() {
		_, nums1, nums2 := parseLine(scanner.Text())
		cnt := countWinningNumbers(nums1, nums2)
		rs += int(math.Pow(2, float64(cnt)-1))
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day4) Part2() {
	scanner := utils.NewScanner(4)
	m := make(map[int]int)
	for scanner.Scan() {
		id, nums1, nums2 := parseLine(scanner.Text())
		m[id] = countWinningNumbers(nums1, nums2)
	}
	fmt.Println(fn1(m))
	// fmt.Println(fn2(m))
}

func fn1(m map[int]int) int {
	dpp := make([]int, len(m)+1)
	for id := len(m); id > 0; id-- {
		for i := m[id]; i >= 1; i-- {
			dpp[id] += dpp[id+i] + 1
		}
	}
	rs := 0
	for i := 0; i < len(dpp); i++ {
		rs += dpp[i] + 1
	}
	return rs - 1
}

func fn2(m map[int]int) int {
	dp = make(map[int]int)
	rs := 0
	for id := range m {
		rs += dfs(m, id) + 1
	}
	return rs
}

var dp map[int]int

func dfs(m map[int]int, id int) int {
	if v, ok := dp[id]; ok {
		return v
	}
	rs := 0
	for i := 1; i <= m[id]; i++ {
		rs += dfs(m, id+i) + 1
	}
	dp[id] = rs
	return rs
}

func countWinningNumbers(nums1, nums2 []int) int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = true
	}
	rs := 0
	for _, num := range nums2 {
		if m[num] {
			rs++
		}
	}
	return rs
}

func parseLine(line string) (int, []int, []int) {
	var id int
	p := strings.Split(line, ": ")
	fmt.Sscanf(p[0], "Card %d", &id)
	nums := strings.Split(p[1], " | ")
	return id, utils.ParseIntSlice(nums[0], " "), utils.ParseIntSlice(nums[1], " ")
}
