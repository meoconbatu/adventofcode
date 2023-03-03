package day24

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

var final, x, y, z, w int
var dp map[string]bool
var w1s = [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
var w2s = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// Day24 struct
type Day24 struct {
}

// Part1 func
func (d Day24) Part1() {
	instructions := readInput()
	final = 0
	dp = make(map[string]bool)
	findWs(instructions, 0, 0, w1s)
	fmt.Println(final)
}

// Part2 func
func (d Day24) Part2() {
	instructions := readInput()
	x, y, z, w, final = 0, 0, 0, 0, 0
	dp = make(map[string]bool)
	findWs(instructions, 0, 0, w2s)
	fmt.Println(final)
}
func findWs(fs []*instruction, ifs, rs int, ws [9]int) bool {
	prev := ifs
	for ; ifs < len(fs) && fs[ifs] != nil; ifs++ {
		if err := (*fs[ifs]).calc(); err != nil {
			return false
		}
	}
	key := getKey(prev)
	if _, ok := dp[key]; ok {
		return false
	}
	tempx, tempy, tempz := x, y, z
	if ifs < len(fs) && fs[ifs] == nil {
		for i := 0; i < len(ws); i++ {
			w = ws[i]
			if findWs(fs, ifs+1, rs*10+w, ws) {
				return true
			}
			x, y, z = tempx, tempy, tempz
		}
	}
	if ifs == len(fs) {
		if z == 0 {
			final = rs
		}
		return z == 0
	}
	dp[key] = false
	return false
}

type instruction struct {
	x  *int
	a  *int
	op string
	s  string
}

func newInstruction(s string) *instruction {
	if s == "inp w" {
		return nil
	}
	ins := new(instruction)
	parts := strings.Split(s, " ")
	ins.op = parts[0]
	ins.s = s
	switch parts[1] {
	case "x":
		ins.x = &x
	case "y":
		ins.x = &y
	case "z":
		ins.x = &z
	case "w":
		ins.x = &w
	}

	switch parts[2] {
	case "x":
		ins.a = &x
	case "y":
		ins.a = &y
	case "z":
		ins.a = &z
	case "w":
		ins.a = &w
	default:
		num, _ := strconv.Atoi(parts[2])
		ins.a = &num
	}
	return ins
}
func (ins instruction) String() string {
	return "[" + ins.s + "]"
}
func (ins instruction) calc() error {
	switch ins.op {
	case "add":
		(*ins.x) = (*ins.x) + (*ins.a)
	case "mul":
		(*ins.x) = (*ins.x) * (*ins.a)
	case "div":
		if (*ins.a) == 0 {
			return errors.New("div zero")
		}
		(*ins.x) = (*ins.x) / (*ins.a)
	case "mod":
		if (*ins.a) < 0 {
			return errors.New("mod negative ")
		}
		(*ins.x) = (*ins.x) % (*ins.a)
	case "eql":
		if (*ins.x) == (*ins.a) {
			(*ins.x) = 1
		} else {
			(*ins.x) = 0
		}
	}
	return nil
}
func getKey(ifs int) string {
	return fmt.Sprintf("%d,%d", ifs, z)
}

func readInput() []*instruction {
	scanner := utils.NewScanner(24)
	instructions := make([]*instruction, 0)
	for scanner.Scan() {
		s := scanner.Text()
		instructions = append(instructions, newInstruction(s))
	}
	return instructions
}
