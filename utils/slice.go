package utils

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// FindMinMax func return min and max number of ins array
func FindMinMax(ins []int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for i := 0; i < len(ins); i++ {
		if ins[i] > max {
			max = ins[i]
		}
		if ins[i] < min {
			min = ins[i]
		}
	}
	return min, max
}

// FindIntersection func find the intersection of two sorted slices.
func FindIntersection(s1, s2 []int) []int {
	i, j := 0, 0
	s := make([]int, 0)
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			s = append(s, s1[i])
			i++
			j++
		} else if s1[i] < s2[j] {
			i++
		} else {
			j++
		}
	}
	return s
}

// CopySlice return a copy of slice ins
func CopySlice(ins []int) []int {
	copyins := make([]int, len(ins))
	copy(copyins, ins)
	return copyins
}

func ParseIntSlice(numsStr string, sep string) []int {
	nums := make([]int, 0)
	for _, numStr := range strings.Split(numsStr, sep) {
		if numStr == "" {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	return nums
}
func SortReverse(nums []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
}
func Reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
