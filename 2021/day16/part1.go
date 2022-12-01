package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Part1 func
func Part1() {
	inputFileName := "day16/input.txt"
	message := readInput(inputFileName)
	output := part1Core(message)
	fmt.Println(output)
}
func part1Core(message string) int {
	binarySequence := toBinary(message)
	fmt.Println(binarySequence)
	parse(binarySequence)
	return versionNumber
}

var versionNumber int

func parse(binarySequence string) {
	version := binarySequence[0:3]
	versionNumber += toNumber(version)

	typeID := binarySequence[3:6]
	remain := binarySequence[6:]
	if typeID == "100" {
		if len(remain)%4 != 0 {
			remain = strings.Repeat("0", 4-len(remain)%4) + remain
		}
		for {
			group := remain[:5]
			remain = remain[5:]
			if group[0] == '0' {
				break
			}
		}
	} else {
		lengthTypeID := remain[0]
		if lengthTypeID == '0' {
			// totalLengthInBit := remain[1:16]
		} else {
			// numSubs := remain[1:12]
		}

	}
}
func toNumber(binaryStr string) int {
	num, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	return int(num)
}
func toBinary(message string) string {
	rs := ""
	for _, r := range message {
		num, err := strconv.ParseInt(string(r), 16, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		rs += fmt.Sprintf("%04b", num)
	}
	return rs
}
func readInput(inputFileName string) string {
	f, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("No input file: %s", inputFileName)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text()
}
