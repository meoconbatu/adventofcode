package day3

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day3/input.txt"
	ins, bitLength := readInput(inputFileName)
	output := part2Core(ins, bitLength)
	fmt.Println(output)
}
func part2Core(ins []int, bitLength int) int {
	epsilon, gamma := genenrate2(ins, bitLength)
	return epsilon * gamma
}

// calculate the oxygen generator rating and CO2 scrubber rating
//
// oxygen generator rating: determine the most common value (0 or 1) in the current bit position,
// and keep only numbers with that bit in that position.
// If 0 and 1 are equally common, keep values with a 1 in the position being considered.
//
// CO2 scrubber rating: determine the least common value (0 or 1) in the current bit position,
// and keep only numbers with that bit in that position.
// If 0 and 1 are equally common, keep values with a 0 in the position being considered.
func genenrate2(ins []int, bitLength int) (int, int) {
	oxygen := findMostCommonValue(ins, bitLength)
	co2 := findLeastCommonValue(ins, bitLength)
	return oxygen, co2
}
func findMostCommonValue(ins []int, bitLength int) int {
	return findCommonValue(ins, bitLength, func(a, b int) bool {
		if a >= b {
			return true
		}
		return false
	})
}
func findLeastCommonValue(ins []int, bitLength int) int {
	return findCommonValue(ins, bitLength, func(a, b int) bool {
		if a < b {
			return true
		}
		return false
	})
}
func findCommonValue(ins []int, bitLength int, f func(a, b int) bool) int {
	inscopy := make([]int, len(ins))
	copy(inscopy, ins)
	rs := 0
	for i := bitLength - 1; i >= 0; i-- {
		bit1s := make([]int, 0)
		bit0s := make([]int, 0)
		for _, in := range inscopy {
			if in&(1<<i) > 0 {
				bit1s = append(bit1s, in)
			} else {
				bit0s = append(bit0s, in)
			}
		}
		if f(len(bit1s), len(bit0s)) {
			inscopy = bit1s
		} else {
			inscopy = bit0s
		}
		if len(inscopy) == 1 {
			rs = inscopy[0]
		}
	}
	return rs
}
