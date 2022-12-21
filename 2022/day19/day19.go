package day19

import (
	"fmt"
	"time"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day19 type
type Day19 struct{}

var dp map[int]int

// Part1 func
func (d Day19) Part1() {
	blueprints := readInput()
	rs := 0

	s := time.Now()
	for _, blueprint := range blueprints {
		rs += blueprint[0] * maxGeodes(blueprint, 24)
	}
	fmt.Println(rs, time.Since(s))
}

// Part2 func
func (d Day19) Part2() {
	blueprints := readInput()
	rs := 1

	s := time.Now()
	for i, blueprint := range blueprints {
		if i >= 3 {
			break
		}
		rs *= maxGeodes(blueprint, 32)
	}
	fmt.Println(rs, time.Since(s))
}

func readInput() [][7]int {
	scanner := utils.NewScanner(19)
	rs := make([][7]int, 0)
	for scanner.Scan() {
		var i, x1, x2, x3, x4, x5, x6 int
		fmt.Sscanf(scanner.Text(), "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.\n",
			&i, &x1, &x2, &x3, &x4, &x5, &x6)
		rs = append(rs, [7]int{i, x1, x2, x3, x4, x5, x6})
	}
	return rs
}

func maxGeodes(blueprint [7]int, mins int) int {
	s := time.Now()
	// ore, clay, obsidian, geode
	fmt.Println(blueprint)

	robots, gems := [4]int{1}, [4]int{}
	dp = make(map[int]int)
	rs := dfs(mins, gems, robots, blueprint)

	fmt.Println(blueprint[0], rs)
	fmt.Println(time.Since(s))
	return rs
}

func dfs(mins int, gems, robots [4]int, blueprint [7]int) int {
	if mins == 0 {
		return gems[3]
	}
	// TODO:
	// don't know why use string as a key (the line below), some blueprints are wrong
	// key := fmt.Sprintf("%v%v", gems, robots)
	key := (mins << 4)
	for i, robot := range robots {
		key = (key << 4) | robot
		key = (key << 4) | gems[i]
	}
	if v, ok := dp[key]; ok {
		return v
	}
	rs := 0
	if gems[0] >= blueprint[5] && gems[2] >= blueprint[6] {
		newgems := [4]int{gems[0] - blueprint[5] + robots[0], gems[1] + robots[1], gems[2] - blueprint[6] + robots[2], gems[3] + robots[3]}
		newrobots := [4]int{robots[0], robots[1], robots[2], robots[3] + 1}
		rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint))
	}
	if gems[0] >= blueprint[3] && gems[1] >= blueprint[4] {
		newgems := [4]int{gems[0] - blueprint[3] + robots[0], gems[1] - blueprint[4] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
		newrobots := [4]int{robots[0], robots[1], robots[2] + 1, robots[3]}
		rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint))
	}
	if gems[0] >= blueprint[2] {
		newgems := [4]int{gems[0] - blueprint[2] + robots[0], gems[1] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
		newrobots := [4]int{robots[0], robots[1] + 1, robots[2], robots[3]}
		rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint))
	}
	if gems[0] >= blueprint[1] {
		newgems := [4]int{gems[0] - blueprint[1] + robots[0], gems[1] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
		newrobots := [4]int{robots[0] + 1, robots[1], robots[2], robots[3]}
		rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint))
	}
	for i, robot := range robots {
		gems[i] += robot
	}
	rs = utils.Max(rs, dfs(mins-1, gems, robots, blueprint))
	dp[key] = rs
	return rs
}
