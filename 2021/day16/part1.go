package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Day16 struct {
}

// Part1 func
func (d Day16) Part1() {
	inputFileName := "day16/input.txt"
	message := readInput(inputFileName)
	output := part1Core(message)
	fmt.Println(output)
}
func part1Core(message string) int {
	binarySequence := toBinary(message)
	_, versionNumber := Parse(binarySequence)
	return versionNumber
}

func Parse(binarySequence string) (int, int) {
	versionNumber := toNumber(binarySequence[0:3])
	typeID := binarySequence[3:6]

	remain := binarySequence[6:]

	if typeID == "100" {
		rsn := 6
		for {
			group := remain[:5]
			remain = remain[5:]
			rsn += 5
			if group[0] == '0' {
				break
			}
		}
		return rsn, versionNumber
	}
	lengthTypeID := remain[0]
	if lengthTypeID == '0' {
		totalLengthInBit := toNumber(remain[1:16])
		remain = remain[16:]
		rsn, rsvn := 0, versionNumber
		for {
			n, vn := Parse(remain)
			totalLengthInBit -= n
			rsvn += vn
			rsn += n
			remain = remain[n:]
			if totalLengthInBit < 6 {
				break
			}
		}
		return rsn + 6 + 1 + 15, rsvn
	}
	numSubs := toNumber(remain[1:12])
	remain = remain[12:]
	rsn, rsvn := 0, versionNumber
	for i := 0; i < numSubs; i++ {
		n, vn := Parse(remain)
		rsvn += vn
		rsn += n
		remain = remain[n:]
	}
	return rsn + 6 + 1 + 11, rsvn
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
