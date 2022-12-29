package day5

import (
	"fmt"
	"log"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day5 type
type Day5 struct{}

func newCrates() []*utils.Stack[byte] {
	return []*utils.Stack[byte]{
		nil,
		utils.NewStack([]byte{'R', 'G', 'H', 'Q', 'S', 'B', 'T', 'N'}),
		utils.NewStack([]byte{'H', 'S', 'F', 'D', 'P', 'Z', 'J'}),
		utils.NewStack([]byte{'Z', 'H', 'V'}),
		utils.NewStack([]byte{'M', 'Z', 'J', 'F', 'G', 'H'}),
		utils.NewStack([]byte{'T', 'Z', 'C', 'D', 'L', 'M', 'S', 'R'}),
		utils.NewStack([]byte{'M', 'T', 'W', 'V', 'H', 'Z', 'J'}),
		utils.NewStack([]byte{'T', 'F', 'P', 'L', 'Z'}),
		utils.NewStack([]byte{'Q', 'V', 'W', 'S'}),
		utils.NewStack([]byte{'W', 'H', 'L', 'M', 'T', 'D', 'N', 'C'}),
	}
}

// Part1 func
func (d Day5) Part1() {
	scanner := utils.NewScanner(5)
	rs := ""
	crates := newCrates()
	for scanner.Scan() {
		var quantity, from, to int
		n, _ := fmt.Sscanf(scanner.Text(), "move %d from %d to %d\n", &quantity, &from, &to)
		if n != 3 {
			continue
		}
		move(crates, from, to, quantity)
	}
	for i := 1; i < len(crates); i++ {
		crate, err := crates[i].Top()
		if err != nil {
			log.Fatalln(err, crate)
		}
		rs += string(crate)
	}
	fmt.Println(rs)
}

func move(crates []*utils.Stack[byte], from, to, quantity int) {
	for i := 0; i < quantity; i++ {
		crate, err := crates[from].Pop()
		if err != nil {
			log.Fatalln(err, crate, from, to)
		}
		crates[to].Push(crate)
	}
}

// Part2 func
func (d Day5) Part2() {
	scanner := utils.NewScanner(5)
	rs := ""
	crates := newCrates()
	for scanner.Scan() {
		var quantity, from, to int
		n, _ := fmt.Sscanf(scanner.Text(), "move %d from %d to %d\n", &quantity, &from, &to)
		if n != 3 {
			continue
		}
		moveMulti(crates, from, to, quantity)
	}
	for i := 1; i < len(crates); i++ {
		crate, err := crates[i].Top()
		if err != nil {
			log.Fatalln(err, crate)
		}
		rs += string(crate)
	}
	fmt.Println(rs)
}

func moveMulti(crates []*utils.Stack[byte], from, to, quantity int) {
	subCrates, err := crates[from].PopMulti(quantity)
	if err != nil {
		log.Fatalln(err, subCrates, from, to)
	}
	crates[to].PushMulti(subCrates)
}
