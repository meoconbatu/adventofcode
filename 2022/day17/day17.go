package day17

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day17 type
type Day17 struct{}

// ####

// .#.
// ###
// .#.

// ..#
// ..#
// ###

// #
// #
// #
// #

// ##
// ##
// Part1 func

var rockTypes = [][]utils.Point{
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}},
	{{X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 2}},
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2}},
	{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}},
	{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
}

type Rock []utils.Point

func NewRock(rock []utils.Point, x, y int) Rock {
	newrock := make([]utils.Point, len(rock))
	for i := range rock {
		newrock[i].X = rock[i].X + x
		newrock[i].Y = rock[i].Y + y
	}
	return newrock
}
func (r *Rock) Move(tower []int, move int) {
	rock := []utils.Point(*r)
	for i := range rock {
		if (rock[i].Y < len(tower) && rock[i].X+move >= 0 && tower[rock[i].Y]&(1<<(rock[i].X+move)) > 0) || rock[i].X+move < 0 || rock[i].X+move >= 7 {
			return
		}
	}
	for i := range rock {
		rock[i].X += move
	}
}
func (r *Rock) Down(tower []int, move int) bool {
	rock := []utils.Point(*r)
	for i := range rock {
		if (rock[i].Y-move >= 0 && rock[i].Y-move < len(tower) && tower[rock[i].Y-move]&(1<<rock[i].X) > 0) || rock[i].Y-move < 0 {
			return false
		}
	}
	for i := range rock {
		rock[i].Y -= move
	}
	return true
}

// Part1 func
func (d Day17) Part1() {
	pattern := readInput()
	fmt.Println(fall(pattern, rockTypes, 2022))
}

func fall(patterns string, rocks [][]utils.Point, rockth int) int {
	tower := make([]int, 0)
	maxHeight := 0
	times := (rockth - ((rockth-1744)/1725)*1725 - 1744 - 1)
	for i, ipattern := 0, 0; i < times; i++ {
		rock := NewRock(rocks[i%5], 2, maxHeight+3)
		for cont := true; cont; cont = rock.Down(tower, 1) {
			step := -1
			if patterns[ipattern] == '>' {
				step = 1
			}
			ipattern = (ipattern + 1) % len(patterns)
			rock.Move(tower, step)
		}
		for _, p := range rock {
			if p.Y >= len(tower) {
				tower = append(tower, make([]int, p.Y+1)...)
			}
			tower[p.Y] |= (1 << p.X)
			maxHeight = utils.Max(maxHeight, p.Y+1)
		}
	}
	// print(maxHeight, tower,0)
	return maxHeight + ((rockth-1744)/1725)*2659 + 2690
}

// Part2 func
func (d Day17) Part2() {
	pattern := readInput()
	fmt.Println(fall(pattern, rockTypes, 1000000000000))
}
func readInput() string {
	scanner := utils.NewScanner(17)
	scanner.Scan()
	return scanner.Text()
}
func print(maxHeight int, tower []int, delta int) {
	for i := maxHeight; i >= 0; i-- {
		for j := 0; j < 7; j++ {
			if tower[i-delta]&(1<<j) > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
