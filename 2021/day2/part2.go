package day2

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day2/input.txt"
	commands := readInput(inputFileName)
	output := part2Core(commands)
	fmt.Println(output)
}
func part2Core(commands []command) int {
	horizontal, depth := getFinalPosition2(commands)
	return horizontal * depth
}

// calculate the horizontal position and depth you would have
// after following the planned course (a series of commands)
//  - down X increases your aim by X units.
//  - up X decreases your aim by X units.
//  - forward X does two things:
//  	- It increases your horizontal position by X units.
//  	- It increases your depth by your aim multiplied by X.
func getFinalPosition2(commands []command) (int, int) {
	horizontal, depth, aim := 0, 0, 0
	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			horizontal += cmd.unit
			depth += aim * cmd.unit
		case "down":
			aim += cmd.unit
		case "up":
			aim -= cmd.unit
		}
	}
	return horizontal, depth
}
