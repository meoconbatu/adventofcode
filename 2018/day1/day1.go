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
		rs += toNumber(scanner.Text())
	}
	fmt.Println(rs)
}
func toNumber(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}

// Part2 func
func (d Day1) Part2() {
	ins := readInput()
	freqs := make(map[int]struct{})
	curFreq := 0
	for {
		for i := 0; i < len(ins); i++ {
			if _, ok := freqs[curFreq]; ok {
				fmt.Println(curFreq)
				return
			}
			freqs[curFreq] = struct{}{}
			curFreq += ins[i]
		}
	}
}

func readInput() []int {
	scanner := utils.NewScanner(1)
	ins := make([]int, 0)
	for scanner.Scan() {
		ins = append(ins, toNumber(scanner.Text()))
	}
	return ins
}
