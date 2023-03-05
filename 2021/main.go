package main

import (
	"fmt"
	"log"
	"time"

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
	"github.com/meoconbatu/adventofcode/2021/day25"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2021, conf.Dayth, conf.Session, false)

	s := time.Now()
	if conf.Dayth == 0 {
		for i := 1; i <= 25; i++ {
			day := NewDay(i)
			if day == nil {
				continue
			}
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
	case 25:
		return &adventofcode.Day{dayth, day25.Day25{}}
	}
	return nil
}
