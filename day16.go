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

var basePattern = []int64{0, 1, 0, -1}

func day16() {

	file, err := os.Open("input16.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// fmt.Println(day16Part1(file, 100))
	fmt.Println(day16Part2(file, 100))
}
func day16Part1(in io.Reader, phaseth int) string {
	scanner := bufio.NewScanner(in)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), "")

	originalIns := make([]int64, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.ParseInt(ins[i], 10, 64)
	}
	copyIns := make([]int64, len(ins))

	for i := 1; i <= phaseth; i++ {
		for j := 0; j < len(originalIns); j++ {
			temp := int64(0)
			for k := 0; k < len(originalIns); k++ {
				kk := (k + 1) / (j + 1)
				kk = kk % 4
				temp += originalIns[k] * basePattern[kk]
			}
			copyIns[j] = int64(math.Abs(float64(temp))) % 10
		}
		copy(originalIns, copyIns)
	}
	var s []string
	for i := 0; i < 8; i++ {
		s = append(s, strconv.Itoa(int(originalIns[i])))
	}
	return strings.Join(s, "")
}

func day16Part2(in io.Reader, phaseth int) string {
	scanner := bufio.NewScanner(in)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), "")

	originalIns := make([]int64, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.ParseInt(ins[i], 10, 64)
	}

	var ss []string
	for i := 0; i < 7; i++ {
		ss = append(ss, strconv.Itoa(int(originalIns[i])))
	}
	offset, _ := strconv.Atoi(strings.Join(ss, ""))

	tempIns := make([]int64, len(ins))
	copy(tempIns, originalIns)
	
	for i := 1; i < 10000; i++ {
		originalIns = append(originalIns, tempIns...)
	}
	
	copyIns := make([]int64, len(originalIns))

	
	for i := 1; i <= phaseth; i++ {
		tempBefore10 := int64(0)
		for j := len(originalIns)-1;j>= offset; j-- {
			temp := originalIns[j]*basePattern[1]+tempBefore10
			
			copyIns[j] = int64(math.Abs(float64(temp))) % 10
			
			tempBefore10 = temp
		}
		copy(originalIns, copyIns)
	}

	var s []string
	for i := offset; i < offset+8; i++ {
		s = append(s, strconv.Itoa(int(originalIns[i])))
	}
	return strings.Join(s, "")
}
