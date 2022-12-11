package utils

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindRightMostSetBit(num int) int {
	i := -1
	for num > 0 {
		i++
		if num&1 == 1 {
			break
		}
		num >>= 1
	}
	return i
}
func Abs(a int) int {
	return int(math.Abs(float64(a)))
}
