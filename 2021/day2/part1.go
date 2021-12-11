package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type command struct {
	// forward, down, up
	direction string
	unit      int
}

// Part1 func
func Part1() {
	inputFileName := "day2/input.txt"
	commands := readInput(inputFileName)
	output := part1Core(commands)
	fmt.Println(output)
}
func part1Core(commands []command) int {
	horizontal, depth := getFinalPosition(commands)
	return horizontal * depth
}

// calculate the horizontal position and depth you would have
// after following the planned course (a series of commands)
// 	- forward X increases the horizontal position by X units.
//  - down X increases the depth by X units.
//  - up X decreases the depth by X units
func getFinalPosition(commands []command) (int, int) {
	horizontal, depth := 0, 0
	for _, cmd := range commands {
		switch cmd.direction {
		case "forward":
			horizontal += cmd.unit
		case "down":
			depth += cmd.unit
		case "up":
			depth -= cmd.unit
		}
	}
	return horizontal, depth
}

func readInput(inputFileName string) []command {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	commands := make([]command, 0)
	for scanner.Scan() {
		var direction string
		var unit int
		fmt.Sscanf(scanner.Text(), "%s %d", &direction, &unit)
		cmd := command{direction, unit}
		commands = append(commands, cmd)
	}
	return commands
}
