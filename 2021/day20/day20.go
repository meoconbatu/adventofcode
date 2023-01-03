package day20

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day20 struct {
}

// Part1 func
func (d Day20) Part1() {
	algo, img := readInput()
	fmt.Println(applyAlgo(img, algo, 2))

}

// Part2 func
func (d Day20) Part2() {
	algo, img := readInput()
	fmt.Println(applyAlgo(img, algo, 50))
}

func applyAlgo(img [][]byte, algo string, times int) int {
	ch := byte('.')
	for k := 0; k < times; k++ {
		img = extendGrid(img, 2, ch)
		if ch == '.' {
			ch = '#'
		} else {
			ch = '.'
		}
		oimg := make([][]byte, len(img))
		for i := 0; i < len(img); i++ {
			oimg[i] = make([]byte, len(img[i]))
			for j := 0; j < len(img[i]); j++ {
				oimg[i][j] = transform(algo, img, i, j, ch)
			}
		}
		img = oimg
	}
	rs := 0
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			if img[i][j] == '#' {
				rs++
			}
		}
	}
	return rs
}
func transform(algo string, grid [][]byte, x, y int, ch byte) byte {
	if x-1 < 0 || y-1 < 0 || x+1 >= len(grid) || y+1 >= len(grid[x]) {
		return ch
	}
	num := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			num <<= 1
			if grid[i+x][j+y] == '#' {
				num |= 1
			}
		}
	}
	return algo[num]
}

func readInput() (string, [][]byte) {
	scanner := utils.NewScanner(20)

	scanner.Scan()
	algo := scanner.Text()
	grid := make([][]byte, 0)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		grid = append(grid, []byte(scanner.Text()))
	}
	return algo, grid
}
func newLine(n int, ch byte) []byte {
	line := make([]byte, n)
	for i := 0; i < n; i++ {
		line[i] = ch
	}
	return line
}
func extendGrid(grid [][]byte, k int, ch byte) [][]byte {
	m, n := len(grid), len(grid[0])
	rs := make([][]byte, m+4)
	mrs, nrs := m+4, n+4
	for i := 0; i < mrs; i++ {
		rs[i] = newLine(nrs, ch)
		if i >= 2 && i < mrs-2 {
			copy(rs[i][2:nrs-2], grid[i-2])
		}
	}
	return rs
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
