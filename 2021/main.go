package main

import (
	"log"

	"github.com/meoconbatu/adventofcode/config"
	"github.com/meoconbatu/adventofcode/utils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(conf.Dayth, conf.Session, false)

	day := utils.NewDay(conf.Dayth)
	day.Execute(conf.Part)
}
