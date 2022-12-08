package day8

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day8 type
type Day8 struct{}

// Part1 func
func (d Day8) Part1() {
	scanner := utils.NewScanner(8)
	trees := make([][]byte, 0)
	for scanner.Scan() {
		s := scanner.Text()
		trees = append(trees, []byte(s))
	}

	fmt.Printf("%d\n", countVisibleTrees(trees))
}
func countVisibleTrees(trees [][]byte) int {
	rs := 0
	m, n := len(trees), len(trees[0])
	prevs := make([]byte, n)
	visibles := make([][]bool, m)

	for i := 0; i < m; i++ {
		visibles[i] = make([]bool, n)
	}
	copy(prevs, trees[0])
	for i := 1; i < m-1; i++ {
		maxRow := trees[i][0]
		for j := 1; j < n-1; j++ {
			rs += fn(trees, visibles, prevs, &maxRow, i, j)
		}
	}
	copy(prevs, trees[n-1])
	for i := m - 2; i > 0; i-- {
		maxRow := trees[i][n-1]
		for j := n - 2; j > 0; j-- {
			rs += fn(trees, visibles, prevs, &maxRow, i, j)
		}
	}
	return rs + (m+n-2)*2
}
func fn(trees [][]byte, visibles [][]bool, prevs []byte, maxRow *byte, i, j int) int {
	rs := 0
	if trees[i][j] > prevs[j] || trees[i][j] > *maxRow {
		if !visibles[i][j] {
			rs = 1
		}
		if trees[i][j] > *maxRow {
			*maxRow = trees[i][j]
		}
		if trees[i][j] > prevs[j] {
			prevs[j] = trees[i][j]
		}
		visibles[i][j] = true

	}
	return rs
}

// Part2 func
func (d Day8) Part2() {
	scanner := utils.NewScanner(8)
	trees := make([][]byte, 0)
	for scanner.Scan() {
		s := scanner.Text()
		trees = append(trees, []byte(s))
	}
	fmt.Printf("%d\n", countVisibleTrees2(trees))
}

func countVisibleTrees2(trees [][]byte) int {
	rs := 0
	m, n := len(trees), len(trees[0])
	scores := make([][]int, m)
	for i := 0; i < m; i++ {
		scores[i] = make([]int, n)
	}
	for i := 1; i < m-1; i++ {
		st := utils.NewStack([]Tree{{trees[i][0], 0}})
		for j := 1; j < n-1; j++ {
			scores[i][j] = fn2(st, trees[i][j])
		}
	}
	for i := 1; i < m-1; i++ {
		st := utils.NewStack([]Tree{{trees[i][n-1], 0}})
		for j := n - 2; j > 0; j-- {
			scores[i][j] *= fn2(st, trees[i][j])
		}
	}
	for j := 1; j < n-1; j++ {
		st := utils.NewStack([]Tree{{trees[0][j], 0}})
		for i := 1; i < m-1; i++ {
			scores[i][j] *= fn2(st, trees[i][j])
		}
	}
	for j := 1; j < n-1; j++ {
		st := utils.NewStack([]Tree{{trees[m-1][j], 0}})
		for i := m - 2; i > 0; i-- {
			scores[i][j] *= fn2(st, trees[i][j])
			rs = utils.Max(rs, scores[i][j])
		}
	}
	return rs
}
func fn2(st *utils.Stack[Tree], height byte) int {
	numSmaller := 1
	for curr, err := st.Top(); err == nil && curr.height < height; {
		if curr.height < height {
			numSmaller += curr.value
		}
		st.Pop()
		curr, err = st.Top()
	}
	st.Push(Tree{height, numSmaller})
	return numSmaller
}

type Tree struct {
	height byte
	value  int
}
