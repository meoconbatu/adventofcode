package main

import (
	"fmt"
	"log"
	"time"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/2020/day20"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2020, conf.Dayth, conf.Session, false)
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
	case 20:
		return &adventofcode.Day{Dayth: dayth, Exec: day20.Day20{}}
	}
	return nil
}
