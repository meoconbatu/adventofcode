package day14

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day14/input.txt"
	template, rules := readInput(inputFileName)
	output := part2Core(template, rules)
	fmt.Println(output)
}
func part2Core(template string, rules map[string]string) int {
	newTemplate := applyRule(template, rules, 40)
	least, most := countCommonElement(template, newTemplate)

	return most - least
}
