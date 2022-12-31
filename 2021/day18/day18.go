package day18

import (
	"fmt"
	"strconv"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day18 struct {
}

type Node struct {
	val         int
	left, right *Node
}

// Part1 func
func (d Day18) Part1() {
	lines := readInput()
	var left, right *Node
	for _, s := range lines {
		if left == nil {
			left, _ = buildTree(s, 0)
		} else {
			right, _ = buildTree(s, 0)
			left = reductionTree(&Node{0, left, right})
		}
	}
	fmt.Println(magnitude(left))
}

// Part2 func
func (d Day18) Part2() {
	lines := readInput()
	rs := 0
	for i, si := range lines {
		for j, sj := range lines {
			if i == j {
				continue
			}
			nodei, _ := buildTree(si, 0)
			nodej, _ := buildTree(sj, 0)
			tree := reductionTree(&Node{0, nodei, nodej})
			rs = utils.Max(rs, magnitude(tree))
		}
	}
	fmt.Println(rs)
}
func buildTree(s string, i int) (*Node, int) {
	if s[i] >= '0' && s[i] <= '9' {
		return &Node{int(s[i] - '0'), nil, nil}, 1
	}
	node := &Node{-1, nil, nil}
	l, ln := buildTree(s, i+1)
	node.left = l
	r, rn := buildTree(s, i+ln+2)
	node.right = r
	return node, ln + rn + 3
}
func magnitude(tree *Node) int {
	if tree.left == nil && tree.right == nil {
		return tree.val
	}
	return 3*magnitude(tree.left) + 2*magnitude(tree.right)
}
func reductionTree(tree *Node) *Node {
	for {
		for explode(tree, tree, nil, nil, 0) {
		}
		if !split(tree, tree) {
			break
		}
	}
	return tree
}
func split(root, curr *Node) bool {
	if curr.left == nil {
		if curr.val >= 10 {
			lval := curr.val / 2
			rval := lval + (curr.val - lval*2)
			curr.left = &Node{lval, nil, nil}
			curr.right = &Node{rval, nil, nil}
			curr.val = -1
			return true
		}
		return false
	}
	if split(root, curr.left) {
		return true
	}
	return split(root, curr.right)
}
func explode(root, curr, pl, pr *Node, height int) bool {
	if curr.left == nil {
		return false
	}
	if height >= 4 && curr.left.left == nil {
		if pr != nil {
			addLeftMost(root, pr.right, curr.right.val)
		}
		if pl != nil {
			addRightMost(root, pl.left, curr.left.val)
		}
		curr.val, curr.left, curr.right = 0, nil, nil
		return true
	}

	if explode(root, curr.left, pl, curr, height+1) {
		return true
	}
	return explode(root, curr.right, curr, pr, height+1)
}
func (root *Node) String() string {
	if root.left == nil {
		return strconv.Itoa(root.val)
	}
	return fmt.Sprintf("[%s,%s]", (root.left.String()), (root.right.String()))
}
func addRightMost(root, node *Node, value int) bool {
	for node.right != nil {
		node = node.right
	}
	node.val += value
	return node.val >= 10
}

func addLeftMost(root, node *Node, value int) bool {
	for node.left != nil {
		node = node.left
	}
	node.val += value
	return node.val >= 10
}

func readInput() []string {
	scanner := utils.NewScanner(18)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
