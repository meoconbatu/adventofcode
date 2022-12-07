package day7

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day7 type
type Day7 struct{}

// Node type
type Node struct {
	Path     string
	SubNodes []*Node
	IsFolder bool
	Size     float64
}

var pathToNode = make(map[string]*Node)

// Part1 func
func (d Day7) Part1() {
	scanner := utils.NewScanner(7)
	var currNode *Node
	for scanner.Scan() {
		s := scanner.Text()
		if isCommand(s) {
			if s[2:4] == "cd" {
				to := s[5:]
				var currPath string
				switch to {
				case "/":
					currPath = "/"
				case "..":
					index := strings.LastIndex(currNode.Path[:len(currNode.Path)-1], "/")
					currPath = currNode.Path[:index+1]
				default:
					currPath = currNode.Path + to + "/"
				}
				currNode = getFolderByPath(currPath)
			}
		} else {
			outs := strings.Split(s, " ")
			var node *Node
			if outs[0] == "dir" {
				node = getFolderByPath(currNode.Path + outs[1] + "/")
			} else {
				size, _ := strconv.Atoi(outs[0])
				node = getFileByPath(currNode.Path+outs[1]+"/", float64(size))
			}
			currNode.SubNodes = append(currNode.SubNodes, node)
		}
	}
	rs := 0.0
	calcSize(pathToNode["/"], &rs)
	fmt.Printf("%f\n", rs)
}
func isCommand(s string) bool {
	return s[0] == '$'
}
func calcSize(root *Node, rs *float64) float64 {
	if !root.IsFolder {
		return root.Size
	}
	total := 0.0
	for _, node := range root.SubNodes {
		total += calcSize(node, rs)
	}
	root.Size = total
	if total < 100000 {
		*rs += total
	}
	return total
}
func getFolderByPath(path string) *Node {
	if pathToNode[path] == nil {
		pathToNode[path] = &Node{path, nil, true, 0.0}
	}
	return pathToNode[path]
}

func getFileByPath(path string, size float64) *Node {
	if pathToNode[path] == nil {
		pathToNode[path] = &Node{path, nil, false, size}
	}
	return pathToNode[path]
}

// Part2 func
func (d Day7) Part2() {
	d.Part1()
	root := pathToNode["/"]
	free := 70000000 - root.Size
	need := 30000000 - free
	rs := 70000000.0
	findSmallest(root, need, &rs)
	fmt.Printf("%f\n", rs)
}

func findSmallest(root *Node, need float64, rs *float64) {
	for _, node := range root.SubNodes {
		if node.IsFolder && node.Size >= need {
			*rs = math.Min(*rs, node.Size)
		}
		findSmallest(node, need, rs)
	}
}
