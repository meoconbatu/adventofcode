package day8

import (
	"fmt"
	"strings"
)

// Part2 func
func Part2() {
	inputFileName := "day8/input.txt"
	patterns, outputs := readInput(inputFileName)
	output := part2Core(patterns, outputs)
	fmt.Println(output)
}
func part2Core(patterns, outputs [][]string) int {
	return count2(patterns, outputs)
}

//digits 1, 4, 7, or 8
func count2(patterns, outputs [][]string) int {
	rs := 0
	for i, pattern := range patterns {
		nums := make([]string, 10)
		m := make(map[int][]string)
		for _, digit := range pattern {
			m[len(digit)] = append(m[len(digit)], digit)
		}
		nums[1] = m[2][0]
		nums[4] = m[4][0]
		nums[7] = m[3][0]
		nums[8] = m[7][0]
		a := diff(nums[1], nums[7])
		bd := diff(a, diff(nums[4], nums[7]))
		cf := nums[1]
		bcef := union(diff(m[5][0], m[5][1]), diff(m[5][0], m[5][2]))
		be := diff(cf, bcef)
		b := same(be, bd)
		d := diff(bd, b)
		e := diff(be, b)
		for _, num := range m[5] {
			if strings.Contains(num, e) {
				nums[2] = num
			} else if strings.Contains(num, b) {
				nums[5] = num
			} else {
				nums[3] = num
			}
		}

		for _, num := range m[6] {
			if !strings.Contains(num, e) {
				nums[9] = num
			} else if strings.Contains(num, d) {
				nums[6] = num
			} else {
				nums[0] = num
			}
		}
		temp := 0
		for j, num := range nums {
			if len(outputs[i][0]) == len(num) && same(num, outputs[i][0]) == num {
				temp += 1000 * j
			}
			if len(outputs[i][1]) == len(num) && same(num, outputs[i][1]) == num {
				temp += 100 * j
			}
			if len(outputs[i][2]) == len(num) && same(num, outputs[i][2]) == num {
				temp += 10 * j
			}
			if len(outputs[i][3]) == len(num) && same(num, outputs[i][3]) == num {
				temp += j
			}
		}
		rs += temp
	}
	return rs
}
func union(w1, w2 string) string {
	rs := ""
	for _, r := range w1 {
		if !strings.Contains(rs, string(r)) {
			rs += string(r)
		}
	}
	for _, r := range w2 {
		if !strings.Contains(rs, string(r)) {
			rs += string(r)
		}
	}
	return rs
}
func same(w1, w2 string) string {
	rs := ""
	for _, r := range w1 {
		if strings.Contains(w2, string(r)) {
			rs += string(r)
		}
	}
	return rs
}

func diff(w1, w2 string) string {
	rs := ""
	for _, r := range w1 {
		if !strings.Contains(w2, string(r)) {
			rs += string(r)
		}
	}
	for _, r := range w2 {
		if !strings.Contains(w1, string(r)) {
			rs += string(r)
		}
	}
	return rs
}
