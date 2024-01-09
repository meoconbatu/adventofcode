package day15

import (
	"fmt"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

const MINLINE = 0
const MAXLINE = 4000000

// Day15 type
type Day15 struct{}

// Part1 func
func (d Day15) Part1() {
	steps := readInput()
	rs := 0
	for _, step := range steps {
		rs += hashMap(step)
	}
	fmt.Println(rs)
}

func hashMap(step string) int {
	rs := 0
	for _, c := range step {
		rs = ((rs + int(c)) * 17) % 256
	}
	return rs
}

type Lens struct {
	label       string
	focalLength int
}

// Part2 func
func (d Day15) Part2() {
	steps := readInput()
	rs := totalFocusingPower(steps)
	fmt.Println(rs)
}
func totalFocusingPower(steps []string) int {
	boxes := make([][]Lens, 256)
	for _, step := range steps {
		if strings.Contains(step, "-") {
			label := step[:len(step)-1]
			boxId := hashMap(label)
			idx := -1
			for i := 0; i < len(boxes[boxId]); i++ {
				if boxes[boxId][i].label == label {
					idx = i
					break
				}
			}
			if idx >= 0 {
				if idx+1 < len(boxes[boxId]) {
					copy(boxes[boxId][idx:], boxes[boxId][idx+1:])
				}
				boxes[boxId] = boxes[boxId][:len(boxes[boxId])-1]
			}
		} else {
			var focalLength int
			parts := strings.Split(step, "=")
			label := parts[0]
			fmt.Sscanf(parts[1], "%d", &focalLength)
			boxId := hashMap(label)
			idx := -1
			for i := 0; i < len(boxes[boxId]); i++ {
				if boxes[boxId][i].label == label {
					idx = i
					break
				}
			}
			if idx >= 0 {
				boxes[boxId][idx].focalLength = focalLength
			} else {
				boxes[boxId] = append(boxes[boxId], Lens{label, focalLength})
			}
		}
	}
	rs := 0
	for i, lenslst := range boxes {
		for j, lens := range lenslst {
			rs += (i + 1) * (j + 1) * lens.focalLength
		}
	}
	return rs
}
func readInput() []string {
	scanner := utils.NewScanner(15)
	scanner.Scan()
	rs := strings.Split(scanner.Text(), ",")
	return rs
}
