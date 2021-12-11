package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day10/input.txt"
	lines := readInput(inputFileName)
	output := part1Core(lines)
	fmt.Println(output)
}
func part1Core(lines []string) int {
	rs := 0
	for _, line := range lines {
		r := findFirstIllegalCharacter(line)
		switch r {
		case ')':
			rs += 3
		case ']':
			rs += 57
		case '}':
			rs += 1197
		case '>':
			rs += 25137
		}
	}
	return rs
}

func findFirstIllegalCharacter(line string) rune {
	runes := make([]rune, 0)
	for _, character := range line {
		if character == '(' || character == '[' || character == '<' || character == '{' {
			runes = append(runes, character)
			continue
		}
		if len(runes) == 0 {
			return character
		}
		lastCharacter := runes[len(runes)-1]
		if (character == ')' && lastCharacter != '(') || (character == ']' && lastCharacter != '[') ||
			(character == '}' && lastCharacter != '{') || (character == '>' && lastCharacter != '<') {
			return character
		}
		runes = runes[:len(runes)-1]
	}
	return 0
}

func readInput(inputFileName string) []string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
