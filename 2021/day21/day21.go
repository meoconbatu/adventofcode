package day21

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day21 struct {
}

// Part1 func
func (d Day21) Part1() {
	playerPositions := readInput()
	spaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	scores := make([]int, len(playerPositions))
	dice := 6
	curPlayer := 0
	times := 0
	for {
		times++
		playerPositions[curPlayer] = (playerPositions[curPlayer] + dice) % len(spaces)
		scores[curPlayer] += spaces[playerPositions[curPlayer]]
		if scores[curPlayer] >= 1000 {
			break
		}
		dice += 9
		curPlayer = (curPlayer + 1) % len(playerPositions)
	}
	rs := 0
	for _, p := range scores {
		if p < 1000 {
			rs += p * times * 3
		}
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day21) Part2() {
	playerPositions := readInput()
	spaces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a, b := dicefn(spaces, playerPositions[0], playerPositions[1], 0, 0, 1)
	fmt.Println(utils.Max(a, b))
}

var dp = make(map[int][2]int)

func dicefn(spaces []int, pos1, pos2, score1, score2, curPlayer int) (int, int) {
	if score1 >= 21 {
		return 1, 0
	}
	if score2 >= 21 {
		return 0, 1
	}
	key := (((((pos1<<4|pos2)<<5)|score1)<<5)|score2)<<2 | curPlayer
	if v, ok := dp[key]; ok {
		return v[0], v[1]
	}
	win1, win2 := 0, 0
	for dice1 := 1; dice1 <= 3; dice1++ {
		for dice2 := 1; dice2 <= 3; dice2++ {
			for dice3 := 1; dice3 <= 3; dice3++ {
				dice := dice1 + dice2 + dice3
				if curPlayer == 1 {
					newpos1 := (pos1 + dice) % len(spaces)
					newscore1 := spaces[newpos1] + score1
					a, b := dicefn(spaces, newpos1, pos2, newscore1, score2, 2)
					win1, win2 = win1+a, win2+b
				}
				if curPlayer == 2 {
					newpos2 := (pos2 + dice) % len(spaces)
					newscore2 := spaces[newpos2] + score2
					a, b := dicefn(spaces, pos1, newpos2, score1, newscore2, 1)
					win1, win2 = win1+a, win2+b
				}
			}
		}
	}
	dp[key] = [2]int{win1, win2}
	return win1, win2
}
func readInput() []int {
	scanner := utils.NewScanner(21)
	startingPositions := make([]int, 0)
	for scanner.Scan() {
		var i, num int
		fmt.Sscanf(scanner.Text(), "Player %d starting position: %d\n", &i, &num)
		startingPositions = append(startingPositions, num-1)
	}
	return startingPositions
}
