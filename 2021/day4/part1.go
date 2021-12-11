package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var winTemplates = []string{"0000000000000000000011111", "0000000000000001111100000", "0000000000111110000000000", "0000011111000000000000000", "1111100000000000000000000",
	"1000010000100001000010000", "0100001000010000100001000", "0010000100001000010000100", "0001000010000100001000010",
	"0000100001000010000100001"}

// Part1 func
func Part1() {
	inputFileName := "day4/input.txt"
	numbers, boards := readInput(inputFileName)
	output := part1Core(numbers, boards)
	fmt.Println(output)
}
func part1Core(numbers []int, boards [][]int) int {
	return draw(numbers, boards)
}

// find which board will win first, then calculate and return the final score
func draw(numbers []int, boards [][]int) int {
	marked := make([]int, len(boards))
	for _, num := range numbers {
		for i, board := range boards {
			for j := 0; j < len(board); j++ {
				if board[j] == num {
					marked[i] |= 1 << j
					if isWinner(marked[i]) {
						// calculate final score
						sumOfUnmaskNumber := 0
						for k := 0; k < len(board); k++ {
							if marked[i]&(1<<k) == 0 {
								sumOfUnmaskNumber += board[k]
							}
						}
						return num * sumOfUnmaskNumber
					}
				}
			}
		}
	}
	return 0
}

// the winner has at least one complete row or column of marked numbers
func isWinner(in int) bool {
	for _, template := range winTemplates {
		mask, err := strconv.ParseInt(template, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		if int64(in)&mask == mask {
			return true
		}
	}
	return false
}
func readInput(inputFileName string) ([]int, [][]int) {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)

	boards := make([][]int, 0)
	numbers := make([]int, 0)
	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), ",") {
		var num int
		fmt.Sscanf(c, "%d", &num)
		numbers = append(numbers, num)
	}
	var board []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			if len(board) > 0 {
				boards = append(boards, board)
			}
			board = make([]int, 0)
			continue
		}
		for _, c := range strings.Split(scanner.Text(), " ") {
			var num int
			if _, err := fmt.Sscanf(c, "%d", &num); err == nil {
				board = append(board, num)
			}
		}
	}
	if len(board) > 0 {
		boards = append(boards, board)
	}
	return numbers, boards
}
