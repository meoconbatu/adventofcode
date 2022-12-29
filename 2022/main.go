package main

import (
	"fmt"
	"log"
	"time"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"

	"github.com/meoconbatu/adventofcode/2022/day1"
	"github.com/meoconbatu/adventofcode/2022/day10"
	"github.com/meoconbatu/adventofcode/2022/day11"
	"github.com/meoconbatu/adventofcode/2022/day12"
	"github.com/meoconbatu/adventofcode/2022/day13"
	"github.com/meoconbatu/adventofcode/2022/day14"
	"github.com/meoconbatu/adventofcode/2022/day15"
	"github.com/meoconbatu/adventofcode/2022/day16"
	"github.com/meoconbatu/adventofcode/2022/day17"
	"github.com/meoconbatu/adventofcode/2022/day18"
	"github.com/meoconbatu/adventofcode/2022/day19"
	"github.com/meoconbatu/adventofcode/2022/day2"
	"github.com/meoconbatu/adventofcode/2022/day20"
	"github.com/meoconbatu/adventofcode/2022/day21"
	"github.com/meoconbatu/adventofcode/2022/day22"
	"github.com/meoconbatu/adventofcode/2022/day23"
	"github.com/meoconbatu/adventofcode/2022/day24"
	"github.com/meoconbatu/adventofcode/2022/day25"
	"github.com/meoconbatu/adventofcode/2022/day3"
	"github.com/meoconbatu/adventofcode/2022/day4"
	"github.com/meoconbatu/adventofcode/2022/day5"
	"github.com/meoconbatu/adventofcode/2022/day6"
	"github.com/meoconbatu/adventofcode/2022/day7"
	"github.com/meoconbatu/adventofcode/2022/day8"
	"github.com/meoconbatu/adventofcode/2022/day9"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2022, conf.Dayth, conf.Session, false)
	s := time.Now()
	if conf.Dayth == 0 {
		for i := 1; i <= 25; i++ {
			day := NewDay(i)
			day.Execute(1)
			day.Execute(2)
		}
	} else {
		day := NewDay(conf.Dayth)
		day.Execute(conf.Part)
	}
	fmt.Println(time.Since(s))
}

// NewDay func
func NewDay(dayth int) *adventofcode.Day {
	switch dayth {
	case 1:
		return &adventofcode.Day{Dayth: dayth, Exec: day1.Day1{}}
	case 2:
		return &adventofcode.Day{Dayth: dayth, Exec: day2.Day2{}}
	case 3:
		return &adventofcode.Day{Dayth: dayth, Exec: day3.Day3{}}
	case 4:
		return &adventofcode.Day{Dayth: dayth, Exec: day4.Day4{}}
	case 5:
		return &adventofcode.Day{Dayth: dayth, Exec: day5.Day5{}}
	case 6:
		return &adventofcode.Day{Dayth: dayth, Exec: day6.Day6{}}
	case 7:
		return &adventofcode.Day{Dayth: dayth, Exec: day7.Day7{}}
	case 8:
		return &adventofcode.Day{Dayth: dayth, Exec: day8.Day8{}}
	case 9:
		return &adventofcode.Day{Dayth: dayth, Exec: day9.Day9{}}
	case 10:
		return &adventofcode.Day{Dayth: dayth, Exec: day10.Day10{}}
	case 11:
		return &adventofcode.Day{Dayth: dayth, Exec: day11.Day11{}}
	case 12:
		return &adventofcode.Day{Dayth: dayth, Exec: day12.Day12{}}
	case 13:
		return &adventofcode.Day{Dayth: dayth, Exec: day13.Day13{}}
	case 14:
		return &adventofcode.Day{Dayth: dayth, Exec: day14.Day14{}}
	case 15:
		return &adventofcode.Day{Dayth: dayth, Exec: day15.Day15{}}
	case 16:
		return &adventofcode.Day{Dayth: dayth, Exec: day16.Day16{}}
	case 17:
		return &adventofcode.Day{Dayth: dayth, Exec: day17.Day17{}}
	case 18:
		return &adventofcode.Day{Dayth: dayth, Exec: day18.Day18{}}
	case 19:
		return &adventofcode.Day{Dayth: dayth, Exec: day19.Day19{}}
	case 20:
		return &adventofcode.Day{Dayth: dayth, Exec: day20.Day20{}}
	case 21:
		return &adventofcode.Day{Dayth: dayth, Exec: day21.Day21{}}
	case 22:
		return &adventofcode.Day{Dayth: dayth, Exec: day22.Day22{}}
	case 23:
		return &adventofcode.Day{Dayth: dayth, Exec: day23.Day23{}}
	case 24:
		return &adventofcode.Day{Dayth: dayth, Exec: day24.Day24{}}
	case 25:
		return &adventofcode.Day{Dayth: dayth, Exec: day25.Day25{}}
	}
	return nil
}
