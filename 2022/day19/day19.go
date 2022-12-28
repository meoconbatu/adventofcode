package day19

import (
	"fmt"
	"math"
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
		temp := maxGeodes(blueprint, 24)
		rs += blueprint[0] * temp
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
	// ore, clay, obsidian, geode
	robots, gems := [4]int{1}, [4]int{}
	dp = make(map[int]int)
	g := 0
	// dfs(mins, gems, robots, blueprint, &g)
	return dfs(mins, gems, robots, blueprint, &g)
}
func bfs(blueprint [7]int) int {
	q := []int{1, 0}
	mins := 24
	lenq := 2
	rs := 0
	visited := make(map[[3]int]bool)
	for mins >= 0 {
		robots, gems := q[0], q[1]
		q = q[2:]
		lenq -= 2
		if mins == 0 {
			rs = utils.Max(rs, get(gems, 3))
		}
		if get(gems, 0) >= blueprint[5] && get(gems, 2) >= blueprint[6] {
			newgems := gems + robots - int(math.Pow(1000, 0))*blueprint[5] - int(math.Pow(1000, 2))*blueprint[6]
			newrobots := robots + int(math.Pow(1000, 3)*1)
			if !visited[[3]int{newrobots, newgems, mins}] {
				q = append(q, newrobots, newgems)
				visited[[3]int{newrobots, newgems, mins}] = true
			}

		} else {
			if get(gems, 0) >= blueprint[3] && get(gems, 1) >= blueprint[4] && (get(robots, 2) < blueprint[6] || get(robots, 2)*mins+get(gems, 2) < mins*blueprint[6]) {
				newgems := gems + robots - int(math.Pow(1000, 0))*blueprint[3] - int(math.Pow(1000, 1))*blueprint[4]
				newrobots := robots + int(math.Pow(1000, 2)*1)
				if !visited[[3]int{newrobots, newgems, mins}] {
					q = append(q, newrobots, newgems)
					visited[[3]int{newrobots, newgems, mins}] = true
				}
			}
			if get(gems, 0) >= blueprint[2] && (get(robots, 1) < blueprint[4] || get(robots, 1)*mins+get(gems, 1) < mins*blueprint[4]) {
				newgems := gems + robots - int(math.Pow(1000, 0))*blueprint[2]
				newrobots := robots + int(math.Pow(1000, 1)*1)
				if !visited[[3]int{newrobots, newgems, mins}] {
					q = append(q, newrobots, newgems)
					visited[[3]int{newrobots, newgems, mins}] = true
				}
			}
			if get(gems, 0) >= blueprint[1] && (get(robots, 0) < blueprint[3]+blueprint[5] || get(robots, 0)*mins+get(gems, 0) < mins*blueprint[3]) {
				newgems := gems + robots - int(math.Pow(1000, 0))*blueprint[1]
				newrobots := robots + 1
				if !visited[[3]int{newrobots, newgems, mins}] {
					q = append(q, newrobots, newgems)
					visited[[3]int{newrobots, newgems, mins}] = true
				}
			}
			if get(robots, 0)*mins+get(gems, 0) < mins*(blueprint[2]+blueprint[3]+blueprint[5]) {
				newgems := gems + robots
				newrobots := robots
				if !visited[[3]int{newrobots, newgems, mins}] {
					q = append(q, newrobots, newgems)
					visited[[3]int{newrobots, newgems, mins}] = true
				}
			}
		}
		if lenq == 0 {
			mins--
			lenq = len(q)
		}
	}
	return rs
}
func get(num, index int) int {
	for i := 0; i < index; i++ {
		num /= 1000
	}
	return num % 1000
}
func dfs(mins int, gems, robots [4]int, blueprint [7]int, maxGeode *int) int {
	if mins == 0 {
		*maxGeode = utils.Max(*maxGeode, gems[3])
		return gems[3]
	}
	// TODO:
	// don't know why use string as a key (the line below), some blueprints are wrong
	// key := fmt.Sprintf("%v%v", gems, robots)

	key := mins
	for i := 0; i < 3; i++ {
		key = (key << 4) | robots[i]
		key = (key << 4) | gems[i]
	}
	if v, ok := dp[key]; ok {
		return v
	}
	if gems[3]+robots[3]+robots[3]*mins+(mins*(mins-1))/2 <= *maxGeode {
		return 0
	}
	rs := 0
	if gems[0] >= blueprint[5] && gems[2] >= blueprint[6] {
		newgems := [4]int{gems[0] - blueprint[5] + robots[0], gems[1] + robots[1], gems[2] - blueprint[6] + robots[2], gems[3] + robots[3]}
		newrobots := [4]int{robots[0], robots[1], robots[2], robots[3] + 1}
		rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint, maxGeode))
	} else {
		if gems[0] >= blueprint[3] && gems[1] >= blueprint[4] && (robots[2] < blueprint[6] || robots[2]*mins+gems[2] < mins*blueprint[6]) {
			newgems := [4]int{gems[0] - blueprint[3] + robots[0], gems[1] - blueprint[4] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
			newrobots := [4]int{robots[0], robots[1], robots[2] + 1, robots[3]}
			rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint, maxGeode))
		}
		if gems[0] >= blueprint[2] && (robots[1] < blueprint[4] || robots[1]*mins+gems[1] < mins*blueprint[4]) {
			newgems := [4]int{gems[0] - blueprint[2] + robots[0], gems[1] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
			newrobots := [4]int{robots[0], robots[1] + 1, robots[2], robots[3]}
			rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint, maxGeode))
		}
		if gems[0] >= blueprint[1] && (robots[0] < blueprint[3]+blueprint[5] || robots[0]*mins+gems[0] < mins*blueprint[3]) {
			newgems := [4]int{gems[0] - blueprint[1] + robots[0], gems[1] + robots[1], gems[2] + robots[2], gems[3] + robots[3]}
			newrobots := [4]int{robots[0] + 1, robots[1], robots[2], robots[3]}
			rs = utils.Max(rs, dfs(mins-1, newgems, newrobots, blueprint, maxGeode))
		}
		if robots[0]*mins+gems[0] < mins*(blueprint[2]+blueprint[3]+blueprint[5]) {
			for i, robot := range robots {
				gems[i] += robot
			}
			rs = utils.Max(rs, dfs(mins-1, gems, robots, blueprint, maxGeode))
		}
	}
	dp[key] = rs
	return rs
}
