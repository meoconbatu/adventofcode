package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day1 type
type Day1 struct{}

// Part1 func
func (d Day1) Part1() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	scanner := bufio.NewScanner(f)

	rs := 0
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			rs = utils.Max(rs, sum)
			sum = 0
			continue
		}
		var num int
		fmt.Sscanf(scanner.Text(), "%d\n", &num)
		sum += num
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day1) Part2() {
	f, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	scanner := bufio.NewScanner(f)

	nums := make([]int, 0)
	sum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			nums = append(nums, sum)
			sum = 0
			continue
		}
		var num int
		fmt.Sscanf(scanner.Text(), "%d\n", &num)
		sum += num
	}
	sort.Ints(nums)
	n := len(nums)
	fmt.Println(nums[n-1] + nums[n-2] + nums[n-3])
}
