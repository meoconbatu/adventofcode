package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Part1 func
func Part1() {
	inputFileName := "day1/input.txt"
	ins := readInput(inputFileName)
	output := part1Core(ins)
	fmt.Println(output)
}

func part1Core(ins []int) int {
	return countIncreaseMeasurement(ins)
}

// count the number of times a depth measurement increases from the previous measurement
// (There is no measurement before the first measurement)
func countIncreaseMeasurement(depths []int) int {
	rs := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			rs++
		}
	}
	return rs
}

func readInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	ins := make([]int, 0)
	for scanner.Scan() {
		in := 0
		_, err := fmt.Sscanf(scanner.Text(), "%d", &in)
		if err != nil {
			log.Fatalf("Invalid input: %s", scanner.Text())
		}
		ins = append(ins, in)
	}
	return ins
}
