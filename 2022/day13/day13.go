package day13

import (
	"fmt"
	"sort"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day13 type
type Day13 struct{}

type List struct {
	Integer int
	Lists   []*List
}

// IsInteger func
func (l *List) IsInteger() bool {
	return l.Lists == nil
}
func (l *List) String() string {
	if l.IsInteger() {
		return fmt.Sprintf("%d", l.Integer)
	}
	s := ""
	for _, l := range l.Lists {
		s += l.String() + ","
	}
	return "[" + strings.Trim(s, ",") + "]"
}

// Part1 func
func (d Day13) Part1() {
	pairs := readInput()
	rs := 0
	for i := 0; i < len(pairs)-1; i += 2 {
		if isInRightOrder(pairs[i], pairs[i+1]) == 1 {
			rs += i/2 + 1
		}
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day13) Part2() {
	pairs := readInput()
	dividers := []string{"[[2]]", "[[6]]"}
	for _, d := range dividers {
		index := 1
		pairs = append(pairs, parse(d, &index))
	}

	sort.Slice(pairs, func(i, j int) bool {
		return isInRightOrder(pairs[i], pairs[j]) == 1
	})

	rs := 1
	for i, pair := range pairs {
		for _, d := range dividers {
			if pair.String() == d {
				rs *= (i + 1)
				break
			}
		}
	}
	fmt.Println(rs)
}

func readInput() []*List {
	scanner := utils.NewScanner(13)
	rs := make([]*List, 0)
	for scanner.Scan() {
		var index, index2 int
		s := scanner.Text()
		rs = append(rs, parse(s, &index))
		scanner.Scan()
		s = scanner.Text()
		rs = append(rs, parse(s, &index2))
		scanner.Scan()
	}
	return rs
}
func isInRightOrder(a, b *List) int {
	if a.IsInteger() && b.IsInteger() {
		if a.Integer < b.Integer {
			return 1
		}
		if a.Integer > b.Integer {
			return -1
		}
		return 0
	}
	newa, newb := a, b
	if a.IsInteger() {
		newa = &List{-1, []*List{a}}
	}
	if b.IsInteger() {
		newb = &List{-1, []*List{b}}
	}
	na, nb := len(newa.Lists), len(newb.Lists)
	for i := 0; i < na && i < nb; i++ {
		temp := isInRightOrder(newa.Lists[i], newb.Lists[i])
		if temp != 0 {
			return temp
		}
	}
	if na == nb {
		return 0
	}
	if na < nb {
		return 1
	}
	return -1
}
func parse(s string, index *int) *List {
	l := make([]*List, 0)
	for *index < len(s) {
		a := s[*index]
		if IsNumber(a) {
			num := int(s[*index] - '0')
			*index++
			if IsNumber(s[*index]) {
				num = num*10 + int(s[*index]-'0')
				*index++
			}
			l = append(l, &List{num, nil})
		} else if a == ',' {
			*index++
			continue
		} else if a == '[' {
			*index++
			l = append(l, parse(s, index))
		} else if a == ']' {
			break
		}
	}
	*index++
	return &List{-1, l}
}

// IsNumber func
func IsNumber(a byte) bool {
	return a >= '0' && a <= '9'
}
