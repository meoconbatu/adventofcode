package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day11 type
type Day11 struct{}

type Monkey struct {
	items       []int
	op          string
	b           string
	divisibleBy int
	throwTos    []int
}

// Part1 func
func (d Day11) Part1() {
	monkeys, _ := readInput()
	n := len(monkeys)
	rs := make([]int, n)
	for i := 0; i < 20; i++ {
		inpect(monkeys, rs)
	}
	sort.Ints(rs)
	fmt.Printf("%v\n", rs[n-1]*rs[n-2])
}

func inpect(monkeys []*Monkey, rs []int) {
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			rs[i]++
			num, _ := strconv.Atoi(monkey.b)
			worryLevel := item
			if monkey.op == "+" {
				worryLevel += num
			} else {
				if monkey.b == "old" {
					worryLevel *= worryLevel
				} else {
					worryLevel *= num
				}
			}
			worryLevel /= 3
			if worryLevel%monkey.divisibleBy == 0 {
				monkeys[monkey.throwTos[0]].items = append(monkeys[monkey.throwTos[0]].items, worryLevel)
			} else {
				monkeys[monkey.throwTos[1]].items = append(monkeys[monkey.throwTos[1]].items, worryLevel)
			}
		}
		monkey.items = nil
	}
}

// Part2 func
func (d Day11) Part2() {
	monkeys, m := readInput()
	n := len(monkeys)
	rs := make([]int, n)
	for i, monkey := range monkeys {
		for _, item := range monkey.items {
			for divisibleBy := range m {
				m[divisibleBy] = item % divisibleBy
			}
			fn(monkeys, m, i, 10000, rs)
		}
	}
	sort.Ints(rs)
	fmt.Printf("%v\n", rs[n-1]*rs[n-2])
}

func fn(monkeys []*Monkey, m map[int]int, index, roundth int, rs []int) {
	if roundth == 0 {
		return
	}
	monkey := monkeys[index]
	rs[index]++
	num, _ := strconv.Atoi(monkey.b)
	for divisibleBy, remainder := range m {
		if monkey.op == "+" {
			m[divisibleBy] = (remainder + num%divisibleBy) % divisibleBy
		} else {
			if monkey.b == "old" {
				m[divisibleBy] = (remainder * (remainder % divisibleBy)) % divisibleBy
			} else {
				m[divisibleBy] = (remainder * (num % divisibleBy)) % divisibleBy
			}
		}
	}
	to := 0
	if m[monkey.divisibleBy] != 0 {
		to = 1
	}
	if monkey.throwTos[to] > index {
		fn(monkeys, m, monkey.throwTos[to], roundth, rs)
	} else {
		fn(monkeys, m, monkey.throwTos[to], roundth-1, rs)
	}
}

func readInput() ([]*Monkey, map[int]int) {
	scanner := utils.NewScanner(11)
	monkeys := make([]*Monkey, 0)
	m := make(map[int]int)
	for scanner.Scan() {
		monkey := new(Monkey)

		scanner.Scan()
		parts := strings.Split(scanner.Text(), ":")
		items := make([]int, 0)
		for _, item := range strings.Split(strings.Trim(parts[1], " "), ",") {
			num, _ := strconv.Atoi(strings.Trim(item, " "))
			items = append(items, num)
		}
		monkey.items = items

		scanner.Scan()
		parts = strings.Split(scanner.Text(), ":")
		op := strings.Trim(strings.Trim(parts[1], " "), " ")
		monkey.op, monkey.b = getOp(op)

		scanner.Scan()
		var num int
		fmt.Sscanf(scanner.Text(), "   Test: divisible by %d\n", &num)
		monkey.divisibleBy = num
		m[monkey.divisibleBy] = 0

		scanner.Scan()
		var numt int
		fmt.Sscanf(scanner.Text(), "	  If true: throw to monkey %d\n", &numt)
		monkey.throwTos = append(monkey.throwTos, numt)

		scanner.Scan()
		var numf int
		fmt.Sscanf(scanner.Text(), "	  If false: throw to monkey %d\n", &numf)
		monkey.throwTos = append(monkey.throwTos, numf)

		monkeys = append(monkeys, monkey)
		scanner.Scan()
	}
	return monkeys, m
}

func getOp(op string) (string, string) {
	parts := strings.Split(op, "=")
	parts[1] = strings.TrimSpace(parts[1])
	parts = strings.Split(parts[1], " ")
	return parts[1], parts[2]
}
