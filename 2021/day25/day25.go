package day25

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day25 struct
type Day25 struct {
}

// Part1 func
func (d Day25) Part1() {
	board := readInput()
	move(board)
}
func move(board [][]byte) {
	rs := 0
	for {
		// fmt.Println(rs)
		// print(board)
		rs++
		if moveEast(board)+moveSouth(board) == 0 {
			break
		}
	}
	fmt.Println(rs)
}

func print(arr [][]byte) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%s", string(arr[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

// >
func moveEast(board [][]byte) int {
	m, n := len(board), len(board[0])
	steps := 0
	for i := 0; i < m; i++ {
		head := board[i][0]
		for j := 0; j < n; j++ {
			if board[i][j] == '>' {
				if ((j+1)%n != 0 && board[i][(j+1)%n] == '.') || ((j+1)%n == 0 && head == '.') {
					board[i][j], board[i][(j+1)%n] = '.', '>'
					steps++
					j++
				}
			}
		}
	}
	return steps
}

// v
func moveSouth(board [][]byte) int {
	m, n := len(board), len(board[0])
	steps := 0
	for j := 0; j < n; j++ {
		head := board[0][j]
		for i := 0; i < m; i++ {
			if board[i][j] == 'v' {
				if ((i+1)%m != 0 && board[(i+1)%m][j] == '.') || ((i+1)%m == 0 && head == '.') {
					board[i][j], board[(i+1)%m][j] = '.', 'v'
					steps++
					i++
				}
			}
		}
	}
	return steps
}

// Part2 func
func (d Day25) Part2() {
}

func readInput() [][]byte {
	scanner := utils.NewScanner(25)
	rs := make([][]byte, 0)
	for scanner.Scan() {
		s := scanner.Text()
		rs = append(rs, []byte(s))
	}
	return rs
}
