package main

import (
	"log"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/2021/day16"
	"github.com/meoconbatu/adventofcode/2021/day17"
	"github.com/meoconbatu/adventofcode/2021/day18"
	"github.com/meoconbatu/adventofcode/2021/day19"
	"github.com/meoconbatu/adventofcode/2021/day20"
	"github.com/meoconbatu/adventofcode/2021/day21"
	"github.com/meoconbatu/adventofcode/2021/day22"
	"github.com/meoconbatu/adventofcode/2021/day23"
	"github.com/meoconbatu/adventofcode/2021/day24"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2021, conf.Dayth, conf.Session, false)

	day := NewDay(conf.Dayth)
	if day == nil {
		log.Fatal("not implement yet")
	}
	if conf.Part == 0 {
		day.Execute(1)
		day.Execute(2)
	} else {
		day.Execute(conf.Part)
	}
}

// NewDay func
func NewDay(dayth int) *adventofcode.Day {
	switch dayth {
	case 16:
		return &adventofcode.Day{dayth, day16.Day16{}}
	case 17:
		return &adventofcode.Day{dayth, day17.Day17{}}
	case 18:
		return &adventofcode.Day{dayth, day18.Day18{}}
	case 19:
		return &adventofcode.Day{dayth, day19.Day19{}}
	case 20:
		return &adventofcode.Day{dayth, day20.Day20{}}
	case 21:
		return &adventofcode.Day{dayth, day21.Day21{}}
	case 22:
		return &adventofcode.Day{dayth, day22.Day22{}}
	case 23:
		return &adventofcode.Day{dayth, day23.Day23{}}
	case 24:
		return &adventofcode.Day{dayth, day24.Day24{}}
	}
	return nil
}
