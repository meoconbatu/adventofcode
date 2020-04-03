package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), ",")

	originalIns := make([]int64, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.ParseInt(ins[i], 10, 64)
	}
	copyIns := make([]int64, len(originalIns))
	for i := 0; i < len(originalIns); i++ {
		for j := 0; j < len(originalIns); j++ {
			copy(copyIns, originalIns)
			copyIns[1], copyIns[2] = int64(i), int64(j)
			process(copyIns, os.Stdin, os.Stdout)
			if copyIns[0] == 19690720 {
				fmt.Println(copyIns)
				return
			}
		}
	}
}
func process(ins []int64, r io.Reader, w io.Writer) {
	x, y, step := 0, 0, 0
	done := false
	maps := make(map[int64]int64)
	maps[0] = 0
	for {
		if done {
			return
		}
		x = y
		opcode := getOpcode(int(ins[x]))

		if opcode == 1 || opcode == 2 || opcode == 7 || opcode == 8 {
			step = 4
		} else if opcode == 5 || opcode == 6 {
			step = 3
		} else {
			step = 2
		}
		y = y + step
		if y > len(ins) {
			y = len(ins)
		}
		done = processIntcode(ins, maps, ins[x:y], &y, r, w)
	}
}
func getOpcode(in int) int {
	return int(math.Mod(float64(in), float64(100)))
}
func getPositionMode(in int) (out [4]int) {
	postModeStr := strconv.Itoa(in / 100)
	j := 1
	for i := len(postModeStr) - 1; i >= 0; i-- {
		tmp, _ := strconv.Atoi(string(postModeStr[i]))
		out[j] = tmp
		j++
	}
	return
}
func processIntcode(ins []int64, maps map[int64]int64, ins4 []int64, nextPoint *int, r io.Reader, w io.Writer) bool {
	// fmt.Println(ins)
	// fmt.Println(ins4)
	opcode := getOpcode(int(ins4[0]))
	positionModes := getPositionMode(int(ins4[0]))
	positions := make([]int64, 0)
	vals := make([]int64, 1)
	vals[0] = 0
	val := int64(0)
	switch opcode {
	case 1:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		pos := positions[3]
		if positionModes[3] == 2 {
			pos += maps[0]
		}
		if pos >= int64(len(ins)) {
			maps[pos] = vals[1] + vals[2]
		} else {
			ins[pos] = vals[1] + vals[2]
		}
	case 2:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		pos := positions[3]
		if positionModes[3] == 2 {
			pos += maps[0]
		}
		if pos >= int64(len(ins)) {
			maps[pos] = vals[1] * vals[2]
		} else {
			ins[pos] = vals[1] * vals[2]
		}
	case 3:
		pos1 := ins4[1]
		var i int64
		for {
			n, _ := fmt.Fscan(r, &i)
			if n == 0 {
				continue
			}
			break
		}
		pos := pos1
		if positionModes[1] == 2 {
			pos += maps[0]
		}
		if pos >= int64(len(ins)) {
			maps[pos] = i
		} else {
			ins[pos] = i
		}
	case 4:
		pos1 := ins4[1]
		switch positionModes[1] {
		case 0:
			if pos1 >= int64(len(ins)) {
				val = maps[pos1]
			} else {
				val = ins[pos1]
			}
		case 1:
			val = pos1
		case 2:
			newPos := pos1 + maps[0]
			if newPos >= int64(len(ins)) {
				val = maps[newPos]
			} else {
				val = ins[newPos]
			}
		}
		for {
			n, err := fmt.Fprintf(w, "%d", val)
			if n == 0 || err != nil {
				continue
			}
			break
		}
	case 5:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		if vals[1] != 0 {
			*nextPoint = int(vals[2])
		}
	case 6:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		if vals[1] == 0 {
			*nextPoint = int(vals[2])
		}
	case 7:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		xx := int64(0)
		if vals[1] < vals[2] {
			xx = 1
		} else {
			xx = 0
		}
		pos := positions[3]
		if positionModes[3] == 2 {
			pos += maps[0]
		}
		if pos >= int64(len(ins)) {
			maps[pos] = xx
		} else {
			ins[pos] = xx
		}
	case 8:
		positions = append(positions, ins4...)
		for i := 1; i < 3; i++ {
			switch positionModes[i] {
			case 0:
				if positions[i] >= int64(len(ins)) {
					val = maps[positions[i]]
				} else {
					val = ins[positions[i]]
				}
			case 1:
				val = positions[i]
			case 2:
				newPos := positions[i] + maps[0]
				if newPos >= int64(len(ins)) {
					val = maps[newPos]
				} else {
					val = ins[newPos]
				}
			}
			vals = append(vals, val)
		}
		xx := int64(0)
		if vals[1] == vals[2] {
			xx = 1
		} else {
			xx = 0
		}
		pos := positions[3]
		if positionModes[3] == 2 {
			pos += maps[0]
		}
		if pos >= int64(len(ins)) {
			maps[pos] = xx
		} else {
			ins[pos] = xx
		}
	case 9:
		pos1 := ins4[1]
		var val int64
		switch positionModes[1] {
		case 0:
			if pos1 >= int64(len(ins)) {
				val = maps[pos1]
			} else {
				val = ins[pos1]
			}
		case 1:
			val = pos1
		case 2:
			newPos := pos1 + maps[0]
			if newPos >= int64(len(ins)) {
				val = maps[newPos]
			} else {
				val = ins[newPos]
			}
		}
		maps[0] += val
	case 99:
		return true
	default:
		return true
	}
	return false
}
