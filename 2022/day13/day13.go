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
	l, ls := readInput()
	rs := 0
	for i, pair := range l {
		fmt.Println(i + 1)
		fmt.Println(ls[i])
		fmt.Println("--------------------------")
		if isInRightOrder(pair[0], pair[1]) == 1 {
			rs += i + 1
			fmt.Println("------>good")
		} else {
			fmt.Println("------>bad")
		}
	}
	fmt.Println(rs)
}

// Part2 func
func (d Day13) Part2() {
	l := readInput2()
	sort.Slice(l, func(i, j int) bool {
		return isInRightOrder(l[i], l[j]) == 1
	})
	fmt.Println(l)
}

func readInput2() []*List {
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
	var index int
	rs = append(rs, parse("[[2]]", &index))
	index = 0
	rs = append(rs, parse("[[6]]", &index))
	return rs
}
func readInput() ([][]*List, [][]string) {
	scanner := utils.NewScanner(13)

	rs := make([][]*List, 0)

	lstr := make([][]string, 0)
	for scanner.Scan() {
		l := make([]*List, 2)
		ll := make([]string, 2)
		var index, index2 int
		s := scanner.Text()
		l[0] = parse(s, &index)
		ll[0] = s
		scanner.Scan()
		s = scanner.Text()
		l[1] = parse(s, &index2)
		ll[1] = s
		rs = append(rs, l)
		lstr = append(lstr, ll)
		scanner.Scan()
	}
	return rs, lstr
}
func isInRightOrder(a, b *List) int {
	// fmt.Println(a, b)
	if a.IsInteger() && b.IsInteger() {
		if a.Integer < b.Integer {
			return 1
		}
		if a.Integer > b.Integer {
			return -1
		}
		return 0
	}
	var newa, newb *List
	if a.IsInteger() {
		newa = &List{-1, []*List{a}}
	} else {
		newa = a
	}
	if b.IsInteger() {
		newb = &List{-1, []*List{b}}
	} else {
		newb = b
	}
	for i := 0; i < len(newa.Lists) && i < len(newb.Lists); i++ {
		temp := isInRightOrder(newa.Lists[i], newb.Lists[i])
		if temp != 0 {
			return temp
		}
	}
	if len(newa.Lists) == len(newb.Lists) {
		return 0
	}
	if len(newa.Lists) < len(newb.Lists) {
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
func IsNumber(a byte) bool {
	return a >= '0' && a <= '9'
}
