package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	empty int = iota
	wall
	block
	paddle
	ball
)

func day13Part1() (map[Point]int, Point, Point) {
	file, err := os.Open("input13_part1.txt")
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
	buf := &bytes.Buffer{}
	process(originalIns, os.Stdin, buf)

	outs := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

	originalOuts := make([]int64, len(outs))
	for i := range outs {
		originalOuts[i], _ = strconv.ParseInt(outs[i], 10, 64)

	}
	paddlePos, ballPos := Point{0, 0}, Point{0, 0}
	maps := make(map[Point]int)
	for i := 0; i+3 <= len(originalOuts); i = i + 3 {
		p := Point{int(originalOuts[i]), int(originalOuts[i+1])}
		maps[p] = int(originalOuts[i+2])
		if int(originalOuts[i+2]) == paddle {
			paddlePos = p
		} else if int(originalOuts[i+2]) == ball {
			ballPos = p
		}
	}
	n := 0
	for _, v := range maps {
		if v == block {
			n++
		}
	}
	// fmt.Println(n)
	// fmt.Println(originalOuts)

	return maps, paddlePos, ballPos
}
func drawGame(originalOuts []int64) {
	for i := 0; i+3 <= len(originalOuts); i = i + 3 {
		switch int(originalOuts[i+2]) {
		case empty:
			fmt.Printf(" ")
		case wall:
			fmt.Printf("|")
		case block:
			fmt.Printf("~")
		case paddle:
			fmt.Printf("=")
		case ball:
			fmt.Printf("o")
		}
		if int(originalOuts[i+1]) == 43 || int(originalOuts[i]) == 43 {
			fmt.Println()
		}
	}
}
func day13Part2() {
	file, err := os.Open("input13_part2.txt")
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
	tempIns := make([]int64, len(ins))
	copy(tempIns, originalIns)
	for {
		remainBlocks, score := day13Part2RunGame(tempIns)
		if remainBlocks == 0 {
			fmt.Println(score)
			break
		}
	}
}
func day13Part2RunGame(ins []int64) (int, int) {
	in, out := &Buffer{}, &Buffer{}
	done := make(chan struct{}, 2)
	done3 := make(chan struct{}, 1)

	maps, paddlePos, ballPos := day13Part1()
	data := make(chan int, 300)
	temps := make([]int, 3)
	score := 0
	go func() {
		i := 0
		for t := range data {
			temps[i] = t
			if i == 2 {
				pp := Point{temps[i-2], temps[i-1]}
				if temps[i-2] == -1 && temps[i-1] == 0 {
					score = temps[i]
				}
				if temps[i] == ball {
					ballPos = pp
				} else if temps[i] == paddle {
					paddlePos = pp
				} else {
					maps[pp] = temps[i]
				}
				i = 0
			} else {
				i++
			}
		}
		done3 <- struct{}{}
	}()
	var mutex = &sync.Mutex{}

	wait := 0
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				if len(in.String()) > 0 {
					continue
				}
				mutex.Lock()
				x := wait
				mutex.Unlock()
				if x < 5000 {
					continue
				}
				mutex.Lock()
				wait = 0
				mutex.Unlock()
				if paddlePos.x < ballPos.x {
					in.WriteString("1")
				} else if paddlePos.x == ballPos.x {
					in.WriteString("0")
				} else {
					in.WriteString("-1")
				}
			}
		}
	}()
	go func() {
		i := 0
		for {
			select {
			case <-done:
				close(data)
				return
			default:
				var v int
				if n, _ := fmt.Fscan(out, &v); n == 0 {
					mutex.Lock()
					wait++
					mutex.Unlock()
					continue
				}
				mutex.Lock()
				wait = 0
				mutex.Unlock()
				t := v
				temps[i] = t

				if i == 2 {
					pp := Point{temps[i-2], temps[i-1]}
					if temps[i-2] == -1 && temps[i-1] == 0 {
						score = temps[i]
					}
					if temps[i] == ball {
						ballPos = pp
					} else if temps[i] == paddle {
						paddlePos = pp
					} else {
						maps[pp] = temps[i]
					}
					i = 0
				} else {
					i++
				}
			}
		}
	}()

	process(ins, in, out)

	done <- struct{}{}
	done <- struct{}{}
	<-done3

	n := 0
	for _, v := range maps {
		if v == block {
			n++
		}
	}

	// fmt.Println("score=", score)
	// fmt.Println("n=", n)
	// fmt.Println("ball=", ballPos)
	// fmt.Println("paddle=", paddlePos)
	return n, score
}
