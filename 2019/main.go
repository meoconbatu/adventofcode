package main

import (
	"fmt"
	"log"
	"time"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/2019/day18"
	"github.com/meoconbatu/adventofcode/2019/day19"
	"github.com/meoconbatu/adventofcode/2019/day2"
	"github.com/meoconbatu/adventofcode/2019/day5"
	"github.com/meoconbatu/adventofcode/2019/day9"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2019, conf.Dayth, conf.Session, false)
	s := time.Now()
	if conf.Dayth == 0 {
		for i := 1; i <= 25; i++ {
			day := NewDay(i)
			day.Execute(1)
			day.Execute(2)
		}
	} else {
		day := NewDay(conf.Dayth)
		if day == nil {
			log.Fatalln("not implement yet")
		}
		day.Execute(conf.Part)
	}
	fmt.Println(time.Since(s))
}

// NewDay func
func NewDay(dayth int) *adventofcode.Day {
	switch dayth {
	case 2:
		return &adventofcode.Day{Dayth: dayth, Exec: day2.Day2{}}
	case 5:
		return &adventofcode.Day{Dayth: dayth, Exec: day5.Day5{}}
	case 9:
		return &adventofcode.Day{Dayth: dayth, Exec: day9.Day9{}}
	case 18:
		return &adventofcode.Day{Dayth: dayth, Exec: day18.Day18{}}
	case 19:
		return &adventofcode.Day{Dayth: dayth, Exec: day19.Day19{}}
	}
	return nil
}
