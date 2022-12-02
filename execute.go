package adventofcode

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/2022/day1"
	"github.com/meoconbatu/adventofcode/2022/day2"
)

// Exec interface
type Exec interface {
	Part1()
	Part2()
}

// Day type
type Day struct {
	dayth int
	Exec
}

// NewDay func
func NewDay(dayth int) *Day {
	switch dayth {
	case 1:
		return &Day{dayth, day1.Day1{}}
	case 2:
		return &Day{dayth, day2.Day2{}}
	}
	return nil
}

// Execute func
func (d Day) Execute(part int) {
	fmt.Printf("day %d, part %d: \n", d.dayth, part)
	if part == 1 {
		d.Part1()
	} else {
		d.Part2()
	}
}
