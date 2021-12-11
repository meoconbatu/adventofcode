package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day8/input.txt"
	patterns, outputs := readInput(inputFileName)
	output := part1Core(patterns, outputs)
	fmt.Println(output)
}
func part1Core(patterns, outputs [][]string) int {
	return count(outputs)
}

//digits 1, 4, 7, or 8
func count(outputs [][]string) int {
	rs := 0
	for _, output := range outputs {
		for _, digit := range output {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				rs++
			}
		}
	}
	return rs
}

// dcga cadgbfe gecba cbfde eda cdbea gbadfe fegcba bedgca da | bgefdac bdace ad agcd
func readInput(inputFileName string) ([][]string, [][]string) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	patterns := make([][]string, 0)
	outputs := make([][]string, 0)
	for scanner.Scan() {
		ins := strings.Split(scanner.Text(), " | ")
		signals := strings.Split(ins[0], " ")
		patterns = append(patterns, signals)

		output := strings.Split(ins[1], " ")
		outputs = append(outputs, output)
	}
	return patterns, outputs
}
