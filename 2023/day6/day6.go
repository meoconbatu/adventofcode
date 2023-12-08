package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day6 type
type Day6 struct{}

// Part1 func
func (d Day6) Part1() {
	scanner := utils.NewScanner(6)
	scanner.Scan()
	p := strings.Split(scanner.Text(), ": ")
	times := parseLstNumStr(p[1])

	scanner.Scan()
	p = strings.Split(scanner.Text(), ": ")
	distances := parseLstNumStr(p[1])

	rs := 1
	for i := 0; i < len(times); i++ {
		temp := waysToWin(times[i], distances[i])
		rs *= temp
	}
	fmt.Println(rs)
}
func waysToWin(time, distance int) int {
	rs := 0
	for i := 0; i <= time; i++ {
		d := i * (time - i)
		if d > distance {
			rs++
		}
	}
	return rs
}
func waysToWin2(time, distance int) int {
	l, r := 0, time/2
	for l < r {
		i := (l + r) / 2
		if d := i * (time - i); d > distance {
			r = i
		} else {
			l = i + 1
		}
	}
	return time + 1 - 2*l
}
func parseLstNumStr(numsStr string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Split(numsStr, " ") {
		if numStr == "" {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}

// Part2 func
func (d Day6) Part2() {
	scanner := utils.NewScanner(6)
	scanner.Scan()
	p := strings.Split(scanner.Text(), ": ")
	times := parseLstNumStr(strings.Replace(p[1], " ", "", -1))

	scanner.Scan()
	p = strings.Split(scanner.Text(), ": ")
	distances := parseLstNumStr(strings.Replace(p[1], " ", "", -1))

	rs := 1
	for i := 0; i < len(times); i++ {
		rs *= waysToWin2(times[i], distances[i])
	}
	fmt.Println(rs)
}
