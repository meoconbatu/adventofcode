package day23

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day23 struct
type Day23 struct {
}

// Part1 func
var roomToEnergy = map[int]int{2: 1, 4: 10, 6: 100, 8: 1000}
var roomToChar = map[int]string{2: "A", 4: "B", 6: "C", 8: "D"}
var initHallways = "00000000000"

// Part1 func
func (d Day23) Part1() {
	rooms := map[int][]int{2: {8, 6}, 4: {2, 2}, 6: {8, 4}, 8: {6, 4}}
	// rooms := map[int][]int{2: {4, 2}, 4: {6, 8}, 6: {4, 6}, 8: {8, 2}}
	fmt.Println(leastEnergy(rooms))
}
func leastEnergy(rooms map[int][]int) int {
	hallways := initHallways
	dp = make(map[string]int)
	return dfs(rooms, hallways, 0)
}

var dp map[string]int

func dfs(rooms map[int][]int, hallways string, totalSteps int) int {
	if hallways == initHallways && totalSteps > 0 {
		return 0
	}
	key := genKey(rooms, hallways)
	if val, ok := dp[key]; ok {
		return val
	}
	minTotalSteps := math.MaxInt64
	for roomth, room := range rooms {
		for iroomth, amphipod := range room {
			if !canMove(rooms, roomth, iroomth) {
				continue
			}
			if pos := findNewPosInRoom(rooms, amphipod); pos != -1 {
				if (amphipod < roomth && hallways[amphipod:roomth+1] == strings.Repeat("0", roomth-amphipod+1)) ||
					(amphipod > roomth && hallways[roomth:amphipod+1] == strings.Repeat("0", amphipod-roomth+1)) {
					steps := (utils.Abs(roomth-amphipod) + 2 + iroomth + pos) * roomToEnergy[amphipod]
					copyRooms := copyMap(rooms)
					copyRooms[amphipod][pos] = amphipod
					copyRooms[roomth][iroomth] = 0
					temp := dfs(copyRooms, hallways, totalSteps+steps)
					if temp != math.MaxInt64 {
						minTotalSteps = utils.Min(minTotalSteps, temp+steps)
					}
				}
			}
			for i := 0; i < len(hallways); i++ {
				if hallways[i] != '0' || roomToEnergy[i] > 0 {
					continue
				}
				if (roomth < i && hallways[roomth:i+1] == strings.Repeat("0", i-roomth+1)) ||
					(roomth > i && hallways[i:roomth+1] == strings.Repeat("0", roomth-i+1)) {
					steps := (utils.Abs(i-roomth) + 1 + iroomth) * roomToEnergy[amphipod]
					copyRooms := copyMap(rooms)
					copyRooms[roomth][iroomth] = 0
					temp := dfs(copyRooms, fmt.Sprintf("%s%d%s", hallways[:i], amphipod, hallways[i+1:]), totalSteps+steps)
					if temp != math.MaxInt64 {
						minTotalSteps = utils.Min(minTotalSteps, temp+steps)
					}
				}
			}
		}
	}
	for i := 0; i < len(hallways); i++ {
		if hallways[i] == '0' {
			continue
		}
		current := int(hallways[i] - '0')
		if pos := findNewPosInRoom(rooms, current); pos != -1 {
			if (current < i && hallways[current:i] == strings.Repeat("0", i-current)) ||
				(current > i && hallways[i+1:current+1] == strings.Repeat("0", current-i)) {
				steps := (utils.Abs(i-current) + 1 + pos) * roomToEnergy[current]
				copyRooms := copyMap(rooms)
				copyRooms[current][pos] = current
				temp := dfs(copyRooms, fmt.Sprintf("%s0%s", hallways[:i], hallways[i+1:]), totalSteps+steps)
				if temp != math.MaxInt64 {
					minTotalSteps = utils.Min(minTotalSteps, temp+steps)
				}
			}
		}
	}
	dp[key] = minTotalSteps
	return minTotalSteps
}

func genKey(rooms map[int][]int, hallways string) string {
	var sb strings.Builder
	sb.Grow(len(hallways) + 20)
	for k, vs := range rooms {
		sb.WriteByte(byte(k))
		for _, v := range vs {
			sb.WriteByte(byte(v))
		}
	}
	sb.WriteString(hallways)
	return sb.String()
}

func canMove(rooms map[int][]int, roomth, index int) bool {
	if rooms[roomth][index] == 0 {
		return false
	}
	for i := 0; i < index; i++ {
		if rooms[roomth][i] != 0 {
			return false
		}
	}
	for i := index; i < len(rooms[roomth]); i++ {
		if rooms[roomth][i] != 0 && rooms[roomth][i] != roomth {
			return true
		}
	}
	return false
}
func findNewPosInRoom(rooms map[int][]int, roomth int) int {
	pos := len(rooms[roomth]) - 1
	for i := len(rooms[roomth]) - 1; i >= 0; i-- {
		if rooms[roomth][i] != roomth && rooms[roomth][i] != 0 {
			return -1
		}
		if rooms[roomth][i] == roomth {
			pos--
		}
	}
	return pos
}
func print(rooms map[int][]int, hallways string) {
	for i := 0; i < len(hallways); i++ {
		if hallways[i] != '0' {
			fmt.Print(roomToChar[int(hallways[i]-'0')])
		} else {
			fmt.Print(".")
		}
	}
	for j := 0; j < len(rooms[2]); j++ {
		fmt.Println()
		for i := 0; i < len(hallways); i++ {
			if _, ok := rooms[i]; ok {
				if rooms[i][j] == 0 {
					fmt.Print(".")
				} else {
					fmt.Print(roomToChar[rooms[i][j]])
				}
			} else {
				fmt.Print("#")
			}
		}
	}
	fmt.Println()
}
func copyMap[K comparable, V any](s map[K][]V) map[K][]V {
	copyS := make(map[K][]V)
	for k, v := range s {
		newv := make([]V, len(v))
		copy(newv, v)
		copyS[k] = newv
	}
	return copyS
}

// Part2 func
func (d Day23) Part2() {
	rooms := map[int][]int{2: {8, 8, 8, 6}, 4: {2, 6, 4, 2}, 6: {8, 4, 2, 4}, 8: {6, 2, 6, 4}}
	// rooms := map[int][]int{2: {4, 8, 8, 2}, 4: {6, 6, 4, 8}, 6: {4, 4, 2, 6}, 8: {8, 2, 6, 2}}
	fmt.Println(leastEnergy(rooms))
}
