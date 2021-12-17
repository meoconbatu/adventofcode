package day14

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day14/input.txt"
	template, rules := readInput(inputFileName)
	output := part1Core(template, rules)
	fmt.Println(output)
}
func part1Core(template string, rules map[string]string) int {
	newTemplate := applyRule(template, rules, 10)
	least, most := countCommonElement(template, newTemplate)
	return most - least
}
func countCommonElement(template string, m map[string]int) (int, int) {
	counts := make(map[byte]int)
	min, max := math.MaxInt64, 0
	for k, v := range m {
		counts[k[0]] += v
	}
	counts[template[len(template)-1]]++

	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
func applyRule(template string, rules map[string]string, numStep int) map[string]int {
	counts := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		counts[string(template[i])+string(template[i+1])]++
	}

	for step := 0; step < numStep; step++ {
		tempCounts := make(map[string]int)
		for k, v := range counts {
			if r, ok := rules[k]; ok {
				tempCounts[string(k[0])+r] += v
				tempCounts[r+string(k[1])] += v
			} else {
				tempCounts[k] = v
			}
		}
		counts = tempCounts
	}
	return counts
}

func readInput(inputFileName string) (string, map[string]string) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	template := scanner.Text()

	scanner.Scan()

	rules := make(map[string]string, 0)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " -> ")
		rules[elements[0]] = elements[1]
	}

	return template, rules
}
