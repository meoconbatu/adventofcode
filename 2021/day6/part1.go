package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var mem = make(map[string]int)

// Part1 func
func Part1() {
	inputFileName := "day6/input.txt"
	nums := readInput(inputFileName)
	output := part1Core(nums)
	fmt.Println(output)
}
func part1Core(ins []int) int {
	return countLanternfish(ins, 80)
}

func countLanternfish(nums []int, dayth int) int {
	sum := 0
	for _, num := range nums {
		sum += f(num, num+1, dayth)
	}
	return sum
}
func f(num int, day, dayth int) int {
	if day > dayth {
		return 1
	}
	key1 := fmt.Sprintf("%d#%d", 6, day+7)
	key2 := fmt.Sprintf("%d#%d", 8, day+9)
	a, b := mem[key1], mem[key2]
	if a == 0 {
		a = f(6, day+7, dayth)
		mem[key1] = a
	}
	if b == 0 {
		b = f(8, day+9, dayth)
		mem[key2] = b
	}
	return a + b
}
func readInput(inputFileName string) []int {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	nums := make([]int, 0)
	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), ",") {
		var num int
		fmt.Sscanf(c, "%d", &num)
		nums = append(nums, num)
	}
	return nums
}
