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

func main() {
	file, err := os.Open("input7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), ",")

	originalIns := make([]int, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.Atoi(ins[i])
	}

	copyIns1 := make([]int, len(originalIns))
	copy(copyIns1, originalIns)
	part1(copyIns1)

	copyIns2 := make([]int, len(originalIns))
	copy(copyIns2, originalIns)
	part2(copyIns2)

}
func part1(originalIns []int) {
	max := 0
	for i := 0; i <= 4; i = i + 1 {
		for j := 0; j <= 4; j = j + 1 {
			for k := 0; k <= 4; k = k + 1 {
				for l := 0; l <= 4; l = l + 1 {
					for m := 0; m <= 4; m = m + 1 {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							settingSequence := []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), strconv.Itoa(l), strconv.Itoa(m)}
							result := processExtendMulti(originalIns, settingSequence)
							if max < result {
								fmt.Println(settingSequence)
								max = result
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

func processExtendMulti(signals []int, settingSequence []string) int {
	temp1, temp2 := &bytes.Buffer{}, &bytes.Buffer{}
	temp2.WriteString("0")
	// fmt.Println(settingSequence)
	for _, sq := range settingSequence {
		copyIns := make([]int, len(signals))
		copy(copyIns, signals)

		temp1.WriteString(sq)
		temp1.WriteString("\n")

		var ii []byte
		fmt.Fscanln(temp2, &ii)
		temp1.Write(ii)

		process(copyIns, temp1, temp2)
		// fmt.Println(temp2.String())
	}
	result, _ := strconv.Atoi(temp2.String())
	return result
}
func part2(originalIns []int) {
	max := 0
	for i := 5; i <= 9; i = i + 1 {
		for j := 5; j <= 9; j = j + 1 {
			for k := 5; k <= 9; k = k + 1 {
				for l := 5; l <= 9; l = l + 1 {
					for m := 5; m <= 9; m = m + 1 {
						if i != j && i != k && i != l && i != m && j != k && j != l && j != m && k != l && k != m && l != m {
							settingSequence := []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), strconv.Itoa(l), strconv.Itoa(m)}
							result := processExtendMultiLoop(originalIns, settingSequence)
							if max < result {
								fmt.Println(settingSequence)
								max = result
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}
func processExtendMultiLoop(signals []int, settingSequence []string) int {
	done := make(chan int)
	donez := make(chan int)
	ins := []*Buffer{&Buffer{}, &Buffer{}, &Buffer{}, &Buffer{}, &Buffer{}}
	outs := []*Buffer{&Buffer{}, &Buffer{}, &Buffer{}, &Buffer{}, &Buffer{}}
	var wg sync.WaitGroup
	// settingSequence := []string{strconv.Itoa(i), strconv.Itoa(j), strconv.Itoa(k), strconv.Itoa(l), strconv.Itoa(m)}
	// fmt.Println(settingSequence)
	for i, sq := range settingSequence {
		wg.Add(1)
		copyIns := make([]int, len(signals))
		copy(copyIns, signals)
		if i == 0 {
			ins[i].WriteString(sq + "\n0")
		} else {
			ins[i].WriteString(sq + "\n")
		}
		go func(ii int, sq string) {
			defer wg.Done()
			process(copyIns, ins[ii], outs[ii])
		}(i, sq)
	}
	go func() {
		i := 0
		for {
			select {
			case <-done:
				donez <- i
				return
			default:
				var v int
				if n, _ := fmt.Fscan(outs[i], &v); n == 0 {
					continue
				}
				if i = i + 1; i == len(settingSequence) {
					i = 0
				}
				ins[i].WriteString(strconv.Itoa(v))
			}
		}
	}()
	wg.Wait()
	done <- 1
	xx := <-donez
	result := 0
	if xx == 4 {
		result, _ = strconv.Atoi(strings.Trim(outs[len(settingSequence)-1].String(), "\n"))
	} else {
		result, _ = strconv.Atoi(strings.Trim(ins[0].String(), "\n"))
	}
	return result
}
