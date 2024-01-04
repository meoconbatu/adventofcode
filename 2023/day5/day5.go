package day5

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day5 type
type Day5 struct{}

// Part1 func
func (d Day5) Part1() {
	b, _ := os.ReadFile(fmt.Sprintf("day%d/input.txt", 5))
	parts := strings.Split(string(b), "\n\n")

	p := strings.Split(parts[0], ": ")
	seeds := utils.ParseIntSlice(p[1], " ")

	m := make([][][]int, 0)
	for _, part := range parts[1:] {
		m = append(m, parseRanges(part))
	}
	fmt.Println(findLowestLocation(seeds, m))
}
func findLowestLocation(seeds []int, m [][][]int) int {
	rs := math.MaxInt64
	for _, seed := range seeds {
		rs = utils.Min(rs, dfs(seed, m, 0))
	}
	return rs
}
func dfs(start int, m [][][]int, th int) int {
	if th == len(m) {
		return start
	}
	phase := m[th]
	idx := sort.Search(len(phase), func(i int) bool {
		return phase[i][1] >= start
	})
	if idx < len(phase) && start == phase[idx][1] {
		start = phase[idx][0]
	} else if idx > 0 && phase[idx-1][1]+phase[idx-1][2]-1 >= start {
		start = phase[idx-1][0] + (start - phase[idx-1][1])
	}
	return dfs(start, m, th+1)
}
func parseRanges(s string) [][]int {
	lines := strings.Split(string(s), "\n")
	rs := make([][]int, 0)
	for _, line := range lines[1:] {
		rs = append(rs, utils.ParseIntSlice(line, " "))
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][1] < rs[j][1]
	})
	return rs
}

// Part2 func
func (d Day5) Part2() {
	b, _ := os.ReadFile(fmt.Sprintf("day%d/input.txt", 5))
	parts := strings.Split(string(b), "\n\n")

	p := strings.Split(parts[0], ": ")
	seeds := utils.ParseIntSlice(p[1], " ")

	m := make([][][]int, 0)
	for _, part := range parts[1:] {
		m = append(m, parseRanges2(part))
	}
	fmt.Println(findLowestLocation2(seeds, m))
}

func parseRanges2(s string) [][]int {
	lines := strings.Split(string(s), "\n")
	rs := make([][]int, 0)
	for _, line := range lines[1:] {
		rs = append(rs, utils.ParseIntSlice(line, " "))
	}
	sort.Slice(rs, func(i, j int) bool {
		return rs[i][0] < rs[j][0]
	})
	return rs
}

func findLowestLocation2(seeds []int, m [][][]int) int {
	nseeds := make([]int, 0)
	for i := 0; i < len(seeds); i += 2 {
		nseeds = append(nseeds, seeds[i], seeds[i]+seeds[i+1])
	}
	sort.Ints(nseeds)

	for i := 0; ; i++ {
		temp := reverse(i, m)
		if idx := sort.SearchInts(nseeds, temp); idx%2 == 1 {
			return i
		}
	}
}
func reverse(dest int, m [][][]int) int {
	for i := len(m) - 1; i >= 0; i-- {
		phase := m[i]
		idx := sort.Search(len(phase), func(i int) bool {
			return phase[i][0] >= dest
		})
		if idx < len(phase) && dest == phase[idx][0] {
			dest = phase[idx][1]
			continue
		}
		if idx > 0 && phase[idx-1][0]+phase[idx-1][2]-1 >= dest {
			dest = phase[idx-1][1] + (dest - phase[idx-1][0])
		}
	}
	return dest
}
