package adventofcode

import (
	"fmt"
	"time"
)

// Exec interface
type Exec interface {
	Part1()
	Part2()
}

// Day type
type Day struct {
	Dayth int
	Exec
}

// Execute func
func (d Day) Execute(part int) {
	fmt.Printf("🎄day %d, part %d: \n", d.Dayth, part)
	s := time.Now()
	if part == 1 {
		d.Part1()
	} else {
		d.Part2()
	}
	fmt.Println(time.Since(s))
}
