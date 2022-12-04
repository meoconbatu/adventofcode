package day3

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day3 type
type Day3 struct{}

// Part1 func
func (d Day3) Part1() {
	scanner := utils.NewScanner(3)
	rs := 0
	for scanner.Scan() {
		rs += GetPriority(scanner.Text())
	}
	fmt.Println(rs)
}

// GetPriority func
func GetPriority(rucksack string) int {
	n := len(rucksack)
	compartment1, compartment2 := rucksack[:n/2], rucksack[n/2:]
	num1 := 0
	for _, c := range compartment1 {
		num1 |= (1 << charToInt(byte(c)))
	}
	num2 := 0
	for _, c := range compartment2 {
		num2 |= (1 << charToInt(byte(c)))
	}
	return utils.FindRightMostSetBit(num1 & num2)
}
func charToInt(c byte) int {
	if c >= 'a' {
		return int(c) - 'a' + 1
	}
	return int(c) - 'A' + 27
}

// Part2 func
func (d Day3) Part2() {
	scanner := utils.NewScanner(3)
	rs := 0
	strs := make([]string, 3)
	i := 0
	for scanner.Scan() {
		strs[i%3] = scanner.Text()
		if i%3 == 2 {
			rs += GetPriority2(strs)
		}
		i++
	}
	fmt.Println(rs)
}

// GetPriority2 func
func GetPriority2(rucksacks []string) int {
	rs := 0xFFFFFFFFFFFFFFF
	for _, rucksack := range rucksacks {
		num := 0
		for _, c := range rucksack {
			num |= (1 << charToInt(byte(c)))
		}
		rs &= num
	}
	return utils.FindRightMostSetBit(rs)
}
