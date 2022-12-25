package day25

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day25 type
type Day25 struct{}

// Part1 func
func (d Day25) Part1() {
	rs := readInput()
	fmt.Println(rs)
}

// Part2 func
func (d Day25) Part2() {

}
func readInput() string {
	scanner := utils.NewScanner(25)
	rs := 0
	for scanner.Scan() {
		s := scanner.Text()
		rs += toDecimal(s)
	}
	return toSnafu(rs)
}

// 122=020221220-0
func toDecimal(s string) int {
	k := 1
	rs := 0
	for i := len(s) - 1; i >= 0; i-- {
		var v int
		switch s[i] {
		case '=':
			v = -2
		case '-':
			v = -1
		default:
			v = int(s[i] - '0')
		}
		rs += k * v
		k *= 5
	}
	return rs
}
func toSnafu(decimal int) string {
	snafu := ""
	for decimal > 0 {
		snafu = fmt.Sprintf("%d%s", (decimal % 5), snafu)
		decimal /= 5
	}

	rs := ""
	remainder := '0'
	for i := len(snafu) - 1; i >= 0; i-- {
		v := int(snafu[i]-'0') + int(remainder-'0')
		remainder = '0'
		if v == 3 {
			rs = "=" + rs
			remainder = '1'
		} else if v == 4 {
			rs = "-" + rs
			remainder = '1'
		} else if v == 5 {
			rs = "0" + rs
			remainder = '1'
		} else {
			rs = fmt.Sprintf("%d%s", (v), rs)
		}
	}
	if remainder == '1' {
		rs = "1" + rs
	}
	return rs
}
