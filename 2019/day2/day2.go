package day2

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day2 type
type Day2 struct{}

// Part1 func
func (d Day2) Part1() {
	nums := readInput()
	nums[1], nums[2] = 12, 2
	out := RunProgram(nums, nil)

	fmt.Println(out)
}

// RunProgram func
func RunProgram(nums []int, systemID *int) int {
	s := ""
	if systemID != nil {
		s = fmt.Sprintf("%d\n", *systemID)
	}
	bf := bytes.NewBufferString(s)
	i := 0
	for {
		op, paramsMode := getMode(nums[i])
		switch op {
		case 99:
			return nums[0]
		case 1:
			ia, ib, ic := nums[i+1], nums[i+2], nums[i+3]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			nums[ic] = a + b
			i += 4
		case 2:
			ia, ib, ic := nums[i+1], nums[i+2], nums[i+3]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			nums[ic] = a * b
			i += 4
		case 3:
			ia := nums[i+1]
			var input int
			inStr, _ := bf.ReadString('\n')
			inStr = strings.TrimRight(inStr, "\n")
			fmt.Sscanf(inStr, "%d", &input)
			nums[ia] = input
			i += 2
		case 4:
			ia := nums[i+1]
			a := ia
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			bf.WriteString(strconv.Itoa(a) + "\n")
			fmt.Println(a)
			i += 2
		case 5:
			ia, ib := nums[i+1], nums[i+2]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			if a != 0 {
				i = b
			} else {
				i += 3
			}
		case 6:
			ia, ib := nums[i+1], nums[i+2]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			if a == 0 {
				i = b
			} else {
				i += 3
			}
		case 7:
			ia, ib, ic := nums[i+1], nums[i+2], nums[i+3]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			if a < b {
				nums[ic] = 1
			} else {
				nums[ic] = 0
			}
			i += 4
		case 8:
			ia, ib, ic := nums[i+1], nums[i+2], nums[i+3]
			a, b := ia, ib
			if paramsMode&1 == 0 {
				a = nums[ia]
			}
			if paramsMode&(1<<1) == 0 {
				b = nums[ib]
			}
			if a == b {
				nums[ic] = 1
			} else {
				nums[ic] = 0
			}
			i += 4
		}
	}
}
func getMode(in int) (int, int) {
	op, temp := in%100, in/100
	paramsMode := 0
	for i := 0; temp > 0; i++ {
		if temp%10 == 1 {
			paramsMode |= (1 << i)
		}
		temp /= 10
	}
	return op, paramsMode
}

// Part2 func
func (d Day2) Part2() {
	nums := readInput()
	copyNums := make([]int, len(nums))
	for pos1 := 0; pos1 <= 99; pos1++ {
		for pos2 := 0; pos2 <= 99; pos2++ {
			copy(copyNums, nums)
			copyNums[1], copyNums[2] = pos1, pos2
			if RunProgram(copyNums, nil) == 19690720 {
				fmt.Println(100*pos1 + pos2)
				return
			}
		}
	}
}

func readInput() []int {
	scanner := utils.NewScanner(2)
	rs := make([]int, 0)
	for scanner.Scan() {
		numStrs := strings.Split(scanner.Text(), ",")
		var num int
		for _, numStr := range numStrs {
			fmt.Sscanf(numStr, "%d", &num)
			rs = append(rs, num)
		}
	}
	return rs
}
