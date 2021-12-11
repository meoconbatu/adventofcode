package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day7/input.txt"
	nums := readInput(inputFileName)
	output := part1Core(nums)
	fmt.Println(output)
}
func part1Core(ins []int) int {
	return sumFuel(ins)
}

func sumFuel(nums []int) int {
	sort.Ints(nums)
	position := len(nums) / 2
	fuel := 0
	for i := 0; i < len(nums); i++ {
		fuel += int(math.Abs(float64(nums[i] - nums[position])))
	}
	return fuel
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
