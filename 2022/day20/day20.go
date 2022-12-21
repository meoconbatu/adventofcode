package day20

import (
	"fmt"
	"strconv"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day20 type
type Day20 struct{}

type Node struct {
	Value      int
	Next, Prev *Node
}

func (root *Node) String() string {
	slow, fast := root, root
	rs := strconv.Itoa(slow.Value)
	for slow.Next != fast.Next.Next {
		slow = slow.Next
		fast = fast.Next.Next
		rs = fmt.Sprintf("%s,%d", rs, slow.Value)
	}
	return rs
}
func (root *Node) Len() int {
	slow, fast := root, root
	rs := 1
	for slow.Next != fast.Next.Next {
		slow = slow.Next
		fast = fast.Next.Next
		rs++
	}
	return rs
}

// Part1 func
func (d Day20) Part1() {
	root, nodes := readInput()
	n := len(nodes)
	rs := 0
	for _, node := range nodes {
		move(node, n)
	}
	p := root
	for i := 1; i <= 3000; i++ {
		p = p.Next
		if i%1000 == 0 {
			rs += p.Value
		}
	}
	fmt.Println(rs)
}

func move(node *Node, n int) {
	if node.Value == 0 {
		return
	}
	p := node
	cycle := node.Value % (n - 1)
	if cycle > 0 {
		for i := 0; i < cycle; i++ {
			p = p.Next
		}
	} else if cycle < 0 {
		for i := -cycle; i >= 0; i-- {
			p = p.Prev
		}
	} else {
		return
	}
	prev := node.Prev
	prev.Next = node.Next
	node.Next.Prev = prev

	next := p.Next
	p.Next = node
	node.Prev = p
	node.Next = next
	next.Prev = node
}

// Part2 func
func (d Day20) Part2() {
	root, nodes := readInput()
	n := len(nodes)
	rs := 0
	for i := 0; i < 10; i++ {
		for _, node := range nodes {
			if i == 0 {
				node.Value *= 811589153
			}
			move(node, n)
		}
	}
	p := root
	for i := 1; i <= 3000; i++ {
		p = p.Next
		if i%1000 == 0 {
			rs += p.Value
		}
	}
	fmt.Println(rs)
}

func readInput() (*Node, []*Node) {
	scanner := utils.NewScanner(20)
	nodes := make([]*Node, 0)
	var root *Node
	i := 0
	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d\n", &num)
		node := &Node{num, nil, nil}
		nodes = append(nodes, node)
		if i > 0 {
			nodes[i-1].Next = node
			node.Prev = nodes[i-1]
		}
		i++
		if num == 0 {
			root = node
		}
	}
	nodes[len(nodes)-1].Next = nodes[0]
	nodes[0].Prev = nodes[len(nodes)-1]
	return root, nodes
}
