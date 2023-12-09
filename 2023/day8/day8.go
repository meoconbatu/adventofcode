package day8

import (
	"fmt"
	"regexp"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day8 type
type Day8 struct{}

// Part1 func
func (d Day8) Part1() {
	instructions, edges := readInput()

	re := regexp.MustCompile(`ZZZ`)
	fmt.Println(stepsToReachTarget(instructions, edges, "AAA", re))
}

// Part2 func
func (d Day8) Part2() {
	instructions, edges := readInput()

	starts := make([]string, 0)
	for u := range edges {
		if u[2] == 'A' {
			starts = append(starts, u)
		}
	}

	re := regexp.MustCompile(`Z$`)

	steps := make([]int, 0)
	for j := 0; j < len(starts); j++ {
		steps = append(steps, stepsToReachTarget(instructions, edges, starts[j], re))
	}
	fmt.Println(utils.LCM(steps[0], steps[1], steps[2:]...))
}

func stepsToReachTarget(instructions string, edges map[string][]string, src string, re *regexp.Regexp) int {
	rs := 0
	u := src
	for i := 0; i < len(instructions); i = (i + 1) % len(instructions) {
		if instructions[i] == 'L' {
			u = edges[u][0]
		} else {
			u = edges[u][1]
		}
		rs++
		if re.MatchString(u) {
			return rs
		}
	}
	return 0
}

func readInput() (string, map[string][]string) {
	scanner := utils.NewScanner(8)

	scanner.Scan()
	instructions := scanner.Text()

	re := regexp.MustCompile(`(.{3}) = \((.{3}), (.{3})\)`)
	edges := make(map[string][]string)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		parts := re.FindStringSubmatch(scanner.Text())
		u, l, r := parts[1], parts[2], parts[3]
		edges[u] = append(edges[u], l, r)
	}
	return instructions, edges
}
