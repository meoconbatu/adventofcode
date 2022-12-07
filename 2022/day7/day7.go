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
type LineType int

const (
	CD LineType = iota
	LS
	FILE
	FOLDER
)

var pathToNode = make(map[string]*Node)

// Part1 func
func (d Day7) Part1() {
	scanner := utils.NewScanner(7)
	var currNode *Node
	for scanner.Scan() {
		s := scanner.Text()
		tpe, name, size := getType(s)
		switch tpe {
		case CD:
			var currPath string
			switch name {
			case "/":
				currPath = "/"
			case "..":
				index := strings.LastIndex(currNode.Path[:len(currNode.Path)-1], "/")
				currPath = currNode.Path[:index+1]
			default:
				currPath = currNode.Path + name + "/"
			}
			currNode = getNodeByPath(currPath, 0, FOLDER)
		case FILE, FOLDER:
			node := getNodeByPath(currNode.Path+name+"/", float64(size), tpe)
			currNode.SubNodes = append(currNode.SubNodes, node)
		}
	}
	rs := 0.0
	calcSize(pathToNode["/"], &rs)
	fmt.Printf("%f\n", rs)
}
func getType(line string) (LineType, string, int) {
	if line[0] == '$' {
		if line[2:4] == "cd" {
			return CD, line[5:], 0
		}
		return LS, "", 0
	}
	parts := strings.Split(line, " ")
	if parts[0] == "dir" {
		return FOLDER, parts[1], 0
	}
	size, _ := strconv.Atoi(parts[0])
	return FILE, parts[1], size
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
func getNodeByPath(path string, size float64, tpe LineType) *Node {
	if pathToNode[path] == nil {
		if tpe == FOLDER {
			pathToNode[path] = &Node{path, nil, true, size}
		} else {
			pathToNode[path] = &Node{path, nil, false, size}
		}
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
