package main

import (
	"log"

	"github.com/meoconbatu/adventofcode"
	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(2022, conf.Dayth, conf.Session, false)

	day := adventofcode.NewDay(conf.Dayth)
	if day == nil {
		log.Fatal("not implement yet")
	}
	day.Execute(conf.Part)
}
