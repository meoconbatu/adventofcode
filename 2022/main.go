package main

import (
	"log"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"

	"github.com/meoconbatu/adventofcode/2022/day1"
	"github.com/meoconbatu/adventofcode/2022/day2"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2022, conf.Dayth, conf.Session, false)

	day := NewDay(conf.Dayth)
	if day == nil {
		log.Fatal("not implement yet")
	}
	day.Execute(conf.Part)
}

// NewDay func
func NewDay(dayth int) *adventofcode.Day {
	switch dayth {
	case 1:
		return &adventofcode.Day{dayth, day1.Day1{}}
	case 2:
		return &adventofcode.Day{dayth, day2.Day2{}}
	}
	return nil
}
