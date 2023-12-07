package day1

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day1 type
type Day1 struct{}

// Part1 func
func (d Day1) Part1() {
	scanner := utils.NewScanner(1)
	rs := 0
	for scanner.Scan() {
		rs += getCalibrationValue(scanner.Text(), getDigit)
	}
	fmt.Println(rs)
}
func getCalibrationValue(s string, fn func(s string, i int) int) int {
	a, b := 0, 0
	for i := 0; i < len(s); i++ {
		if d := fn(s, i); d > 0 {
			a = d
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if d := fn(s, i); d > 0 {
			b = d
			break
		}
	}
	return a*10 + b
}
func getDigit(s string, i int) int {
	if s[i] >= '0' && s[i] <= '9' {
		return int(s[i] - '0')
	}
	return 0
}

// Part2 func
func (d Day1) Part2() {
	scanner := utils.NewScanner(1)
	rs := 0
	for scanner.Scan() {
		rs += getCalibrationValue(scanner.Text(), getDigit2)
	}
	fmt.Println(rs)
}

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func getDigit2(s string, i int) int {
	for j := 0; j < len(digits); j++ {
		if d := getDigit(s, i); d > 0 {
			return d
		}
		if i+len(digits[j]) <= len(s) && s[i:i+len(digits[j])] == digits[j] {
			return j + 1
		}
	}
	return 0
}
