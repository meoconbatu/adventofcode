package day4

import (
	"fmt"
)

// Part2 func
func Part2() {
	inputFileName := "day4/input.txt"
	numbers, boards := readInput(inputFileName)
	output := part2Core(numbers, boards)
	fmt.Println(output)
}
func part2Core(numbers []int, boards [][]int) int {
	return draw2(numbers, boards)
}

func draw2(numbers []int, boards [][]int) int {
	rs := make([]int, len(boards))
	winner := make([]int, len(boards))
	out := 0
	for _, num := range numbers {
		for i, board := range boards {
			if winner[i] > 0 {
				continue
			}
			for j := 0; j < len(board); j++ {
				if board[j] == num {
					rs[i] |= 1 << j
					if isWinner(rs[i]) {
						sum := 0
						for k := 0; k < len(board); k++ {
							if rs[i]&(1<<k) == 0 {
								sum += board[k]
							}
						}
						out = num * sum
						winner[i] = out
					}
				}
			}
		}
	}
	return out
}
