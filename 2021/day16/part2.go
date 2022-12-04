package day16

import (
	"fmt"
	"math"

	"github.com/meoconbatu/adventofcode/utils"
)

// Part2 func
func (d Day16) Part2() {
	inputFileName := "day16/input.txt"
	message := readInput(inputFileName)
	output := part2Core(message)
	fmt.Println(output)
}
func part2Core(message string) int {
	binarySequence := toBinary(message)
	_, _, value := Parse2(binarySequence)
	return value
}

// Parse2 func
func Parse2(binarySequence string) (int, int, int) {
	versionNumber := toNumber(binarySequence[0:3])
	typeID := toNumber(binarySequence[3:6])

	remain := binarySequence[6:]

	if typeID == 4 {
		rsn := 6
		valueStr := ""
		for {
			group := remain[:5]
			remain = remain[5:]
			valueStr += group[1:]
			rsn += 5
			if group[0] == '0' {
				break
			}
		}
		return rsn, versionNumber, toNumber(valueStr)
	}
	lengthTypeID := remain[0]
	nTotal, versionTotal, valueTotal := 0, versionNumber, initValueTotal(typeID)
	if lengthTypeID == '0' {
		totalLengthInBit := toNumber(remain[1:16])
		remain = remain[16:]
		for {
			n, version, value := Parse2(remain)
			totalLengthInBit -= n
			versionTotal += version
			nTotal += n
			remain = remain[n:]
			calcVal(typeID, &valueTotal, value)
			if totalLengthInBit < 6 {
				break
			}
		}

		return nTotal + 6 + 1 + 15, versionTotal, valueTotal
	}
	numSubs := toNumber(remain[1:12])
	remain = remain[12:]
	for i := 0; i < numSubs; i++ {
		n, version, value := Parse2(remain)
		versionTotal += version
		nTotal += n
		remain = remain[n:]
		calcVal(typeID, &valueTotal, value)
	}
	return nTotal + 6 + 1 + 11, versionTotal, valueTotal
}
func initValueTotal(typeID int) int {
	switch typeID {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return math.MaxInt64
	case 3:
		return math.MinInt64
	default:
		return -1
	}
}
func calcVal(typeID int, valueTotal *int, value int) {
	switch typeID {
	case 0:
		*valueTotal += value
	case 1:
		*valueTotal *= value
	case 2:
		*valueTotal = utils.Min(*valueTotal, value)
	case 3:
		*valueTotal = utils.Max(*valueTotal, value)
	case 5:
		if *valueTotal == -1 {
			*valueTotal = value
		} else if *valueTotal > value {
			*valueTotal = 1
		} else {
			*valueTotal = 0
		}
	case 6:
		if *valueTotal == -1 {
			*valueTotal = value
		} else if *valueTotal < value {
			*valueTotal = 1
		} else {
			*valueTotal = 0
		}
	case 7:

		if *valueTotal == -1 {
			*valueTotal = value
		} else if *valueTotal == value {
			*valueTotal = 1
		} else {
			*valueTotal = 0
		}
	}
}
