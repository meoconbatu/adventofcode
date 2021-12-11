package day10

import (
	"fmt"
	"sort"
)

var score = map[rune]int{'[': 2, '(': 1, '{': 3, '<': 4}

// Part2 func
func Part2() {
	inputFileName := "day10/input.txt"
	lines := readInput(inputFileName)
	output := part2Core(lines)
	fmt.Println(output)
}

func part2Core(lines []string) int {
	sums := make([]int, 0)
	for _, line := range lines {
		runes := findIncompleteLine(line)
		if runes == nil {
			continue
		}
		sum := 0
		for i := len(runes) - 1; i >= 0; i-- {
			sum = sum*5 + score[runes[i]]
		}
		sums = append(sums, sum)
	}
	sort.Ints(sums)
	return sums[len(sums)/2]
}

func findIncompleteLine(line string) []rune {
	runes := make([]rune, 0)
	for _, character := range line {
		if character == '(' || character == '[' || character == '<' || character == '{' {
			runes = append(runes, character)
			continue
		}
		if len(runes) == 0 {
			return nil
		}
		lastCharacter := runes[len(runes)-1]
		if (character == ')' && lastCharacter != '(') || (character == ']' && lastCharacter != '[') ||
			(character == '}' && lastCharacter != '{') || (character == '>' && lastCharacter != '<') {
			return nil
		}
		runes = runes[:len(runes)-1]
	}
	return runes
}
