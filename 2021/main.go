package main

import (
	"fmt"
	"log"

	"github.com/meoconbatu/adventofcode/2021/config"
	"github.com/meoconbatu/adventofcode/2021/day1"
	"github.com/meoconbatu/adventofcode/2021/day10"
	"github.com/meoconbatu/adventofcode/2021/day11"
	"github.com/meoconbatu/adventofcode/2021/day12"
	"github.com/meoconbatu/adventofcode/2021/day2"
	"github.com/meoconbatu/adventofcode/2021/day3"
	"github.com/meoconbatu/adventofcode/2021/day4"
	"github.com/meoconbatu/adventofcode/2021/day5"
	"github.com/meoconbatu/adventofcode/2021/day6"
	"github.com/meoconbatu/adventofcode/2021/day7"
	"github.com/meoconbatu/adventofcode/2021/day8"
	"github.com/meoconbatu/adventofcode/2021/day9"
	"github.com/meoconbatu/adventofcode/2021/utils"
)

var funcMap map[string]interface{}

func init() {
	funcMap = map[string]interface{}{
		"011": day1.Part1, "012": day1.Part2,
		"021": day2.Part1, "022": day2.Part2,
		"031": day3.Part1, "032": day3.Part2,
		"041": day4.Part1, "042": day4.Part2,
		"051": day5.Part1, "052": day5.Part2,
		"061": day6.Part1, "062": day6.Part2,
		"071": day7.Part1, "072": day7.Part2,
		"081": day8.Part1, "082": day8.Part2,
		"091": day9.Part1, "092": day9.Part2,
		"101": day10.Part1, "102": day10.Part2,
		"111": day11.Part1, "112": day11.Part2,
		"121": day12.Part1, "122": day12.Part2,
		// "131": day13.Part1, "132": day13.Part2,
		// "141": day14.Part1, "142": day14.Part2,
		// "151": day15.Part1, "152": day15.Part2,
		// "161": day16.Part1, "162": day16.Part2,
		// "171": day17.Part1, "172": day17.Part2,
		// "181": day18.Part1, "182": day18.Part2,
		// "191": day19.Part1, "192": day19.Part2,
		// "201": day20.Part1, // "202": day20.Part2,
		// "211": day21.Part1, // "202": day20.Part2,
		// "221": day22.Part1, "222": day22.Part2,
		// "231": day23.Part1, "232": day23.Part2,
		// "241": day24.Part1, "242": day24.Part2,
		// "251": day25.Part1, //"252": day25.Part2,
	}
}
func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	utils.GetInputFile(conf.Dayth, conf.Session, false)
	execute(conf.Dayth, conf.Part)
}
func execute(dayth, part int) {
	key := fmt.Sprintf("%02d%d", dayth, part)
	if f, ok := funcMap[key]; ok {
		f.(func())()
	} else {
		log.Fatalf("Not support day %d, part %d yet.", dayth, part)
	}
}
