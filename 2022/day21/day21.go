package day21

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day21 type
type Day21 struct{}

type Node struct {
	Name        string
	Value       *int
	Op          *string
	Left, Right *Node
}

func (root *Node) IsValid() bool {
	if root == nil {
		return true
	}
	if root.Value == nil && root.Op == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Value == nil {
		return false
	}
	return root.Left.IsValid() && root.Right.IsValid()
}

func (root *Node) String() string {
	if root.Left == nil && root.Right == nil && root.Value == nil {
		return "x"
	}
	if root.Value != nil {
		return fmt.Sprintf("%d", *root.Value)
	}
	return fmt.Sprintf("(%s%s%s)", root.Left.String(), *root.Op, root.Right.String())
}

// Part1 func
func (d Day21) Part1() {
	root, _ := readInput()
	calc(root)

	fmt.Println(*root.Value)
}

// Part2 func
func (d Day21) Part2() {
	root, nodes := readInput()
	me := nodes["humn"]
	me.Value = nil
	calc(root)

	if root.Left.Value == nil {
		root.Left.Value = root.Right.Value
		findX(root.Left)
	} else if root.Right.Value == nil {
		root.Right.Value = root.Left.Value
		findX(root.Right)
	}
	fmt.Println(*me.Value)
}

func calc(root *Node) *int {
	if root.Left == nil && root.Right == nil {
		return root.Value
	}
	l := calc(root.Left)
	r := calc(root.Right)

	if l == nil || r == nil {
		return nil
	}
	rs := 0
	switch *root.Op {
	case "+":
		rs = *l + *r
	case "-":
		rs = *l - *r
	case "/":
		rs = *l / *r
	case "*":
		rs = *l * *r
	}
	root.Value = &rs
	return &rs
}
func findX(root *Node) {
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left.Value == nil {
		leftVal := *root.Value
		rightVal := *root.Right.Value
		switch *root.Op {
		case "+":
			leftVal -= rightVal
		case "-":
			leftVal += rightVal
		case "/":
			leftVal *= rightVal
		case "*":
			leftVal /= rightVal
		}
		root.Left.Value = &leftVal
		findX(root.Left)
	}
	if root.Right.Value == nil {
		rightVal := *root.Value
		leftVal := *root.Left.Value
		switch *root.Op {
		case "+":
			rightVal -= leftVal
		case "-":
			rightVal = leftVal - rightVal
		case "/":
			rightVal = leftVal / rightVal
		case "*":
			rightVal = rightVal / leftVal
		}
		root.Right.Value = &rightVal
		findX(root.Right)
	}
}

func readInput() (*Node, map[string]*Node) {
	scanner := utils.NewScanner(21)
	nodes := make(map[string]*Node)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		name := parts[0][:len(parts[0])-1]
		if len(parts) == 2 {
			num, _ := strconv.Atoi(parts[1])
			node, ok := nodes[name]
			if !ok {
				nodes[name] = &Node{name, &num, nil, nil, nil}
			} else {
				node.Value = &num
			}
		} else if len(parts) == 4 {
			l, ok := nodes[parts[1]]
			if !ok {
				l = &Node{parts[1], nil, nil, nil, nil}
				nodes[parts[1]] = l
			}
			r, ok := nodes[parts[3]]
			if !ok {
				r = &Node{parts[3], nil, nil, nil, nil}
				nodes[parts[3]] = r
			}
			node, ok := nodes[name]
			if !ok {
				node = &Node{name, nil, &parts[2], l, r}
				nodes[name] = node
			} else {
				node.Left = l
				node.Right = r
				node.Op = &parts[2]
			}
		}
	}
	return nodes["root"], nodes
}
