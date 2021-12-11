package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Part1 func
func Part1() {
	inputFileName := "day3/input.txt"
	ins, bitLength := readInput(inputFileName)
	output := part1Core(ins, bitLength)
	fmt.Println(output)
}
func part1Core(ins []int, bitLength int) int {
	epsilon, gamma := generate(ins, bitLength)
	return epsilon * gamma
}

// generate two new binary numbers (called the gamma rate and the epsilon rate)
//
// gamma rate: the most common bit in the corresponding position of all numbers in the diagnostic report
// epsilon rate: the least common bit from each position.
func generate(ins []int, bitLength int) (int, int) {
	bit1s := make([]int, bitLength)
	for _, in := range ins {
		for i := bitLength - 1; i >= 0; i-- {
			if in&1 > 0 {
				bit1s[i]++
			}
			in >>= 1
		}
	}
	gammaStr, epsilonStr := "", ""
	for _, bit1 := range bit1s {
		if bit1 > len(ins)/2 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}
	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	return int(gamma), int(epsilon)
}
func readInput(inputFileName string) ([]int, int) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	ins := make([]int, 0)
	bitLength := 0
	for scanner.Scan() {
		s := scanner.Text()
		bitLength = len(s)
		in, _ := strconv.ParseUint(s, 2, 64)
		ins = append(ins, int(in))
	}
	return ins, bitLength
}
