package day23

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day23 struct {
}

// #############
// #...........#
// ###D#A#D#C###
//   #C#A#B#B#
//   #########

// Part1 func
var rs int
var roomToEnergy = map[int]int{2: 1, 4: 10, 6: 100, 8: 1000}
var roomToChar = map[int]string{2: "A", 4: "B", 6: "C", 8: "D"}

func (d Day23) Part1() {
	rooms := map[int]int{2: 86, 4: 22, 6: 84, 8: 64}
	// rooms := map[int]int{2: 42, 4: 68, 6: 46, 8: 82}
	hallways := "00000000000"
	rs = math.MaxInt64
	fn(rooms, hallways, 0)
	fmt.Println(rs)
}
func fn(rooms map[int]int, hallways string, totalSteps int) {
	if totalSteps > rs {
		return
	}
	// print(rooms, hallways)
	// fmt.Println(totalSteps)
	end := 0
	for roomth, room := range rooms {
		one, two := room/10, room%10
		if two == roomth && one == roomth {
			end++
			continue
		}
		if one != 0 && (one != roomth || two != roomth) {
			if rooms[one] == 0 || (rooms[one]/10 == 0 && rooms[one]%10 == one) {
				if (one < roomth && hallways[one:roomth+1] == strings.Repeat("0", roomth-one+1)) ||
					(one > roomth && hallways[roomth:one+1] == strings.Repeat("0", one-roomth+1)) {
					copyRooms := copyMap(rooms)
					steps := (utils.Abs(roomth-one) + 2) * roomToEnergy[one]
					if copyRooms[one] == 0 {
						copyRooms[one] = one
						steps += roomToEnergy[one]
					} else if copyRooms[one]/10 == 0 {
						copyRooms[one] += one * 10
					}
					copyRooms[roomth] %= 10
					// fmt.Println("one: ", one, roomth)
					fn(copyRooms, hallways, totalSteps+steps)
				}
			} else {
				for i := 0; i < len(hallways); i++ {
					if hallways[i] != '0' || roomToEnergy[i] > 0 {
						continue
					}
					if (roomth < i && hallways[roomth:i+1] == strings.Repeat("0", i-roomth+1)) ||
						(roomth > i && hallways[i:roomth+1] == strings.Repeat("0", roomth-i+1)) {
						steps := (utils.Abs(i-roomth) + 1) * roomToEnergy[one]
						copyRooms := copyMap(rooms)
						copyRooms[roomth] %= 10
						// fmt.Println("i1: ", i, roomth, string(hallways[i]))
						fn(copyRooms, fmt.Sprintf("%s%d%s", hallways[:i], one, hallways[i+1:]), totalSteps+steps)
					}
				}
			}
		}
		if two != 0 && two != roomth && one == 0 {
			if rooms[two] == 0 || (rooms[two]/10 == 0 && rooms[two]%10 == two) {
				if (two < roomth && hallways[two:roomth+1] == strings.Repeat("0", roomth-two+1)) ||
					(two > roomth && hallways[roomth:two+1] == strings.Repeat("0", two-roomth+1)) {
					copyRooms := copyMap(rooms)
					steps := (utils.Abs(roomth-two) + 2 + 1) * roomToEnergy[two]
					if copyRooms[two] == 0 {
						copyRooms[two] = two
						steps += roomToEnergy[two]
					} else if copyRooms[two]/10 == 0 {
						copyRooms[two] += two * 10
					}
					copyRooms[roomth] /= 10
					// fmt.Println("two: ", two, roomth)
					fn(copyRooms, hallways, totalSteps+steps)
				}
			} else {
				for i := 0; i < len(hallways); i++ {
					if hallways[i] != '0' || roomToEnergy[i] > 0 {
						continue
					}
					if (roomth < i && hallways[roomth:i+1] == strings.Repeat("0", i-roomth+1)) ||
						(roomth > i && hallways[i:roomth+1] == strings.Repeat("0", roomth-i+1)) {
						steps := (utils.Abs(i-roomth) + 1 + 1) * roomToEnergy[two]
						copyRooms := copyMap(rooms)
						copyRooms[roomth] = 0
						// fmt.Println("i2: ", i, roomth, string(hallways[i]))
						fn(copyRooms, fmt.Sprintf("%s%d%s", hallways[:i], two, hallways[i+1:]), totalSteps+steps)
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
		if rooms[current] == 0 || (rooms[current]/10 == 0 && rooms[current]%10 == current) {
			if (current < i && hallways[current:i] == strings.Repeat("0", i-current)) ||
				(current > i && hallways[i+1:current+1] == strings.Repeat("0", current-i)) {
				copyRooms := copyMap(rooms)
				steps := (utils.Abs(i-current) + 1) * roomToEnergy[current]
				if copyRooms[current] == 0 {
					copyRooms[current] = current
					steps += roomToEnergy[current]
				} else if copyRooms[current]/10 == 0 {
					copyRooms[current] += current * 10
				}
				// fmt.Println("from i: ", i, hallways[i])
				fn(copyRooms, fmt.Sprintf("%s0%s", hallways[:i], hallways[i+1:]), totalSteps+steps)
			}
		}
	}
	if end == 4 {
		rs = utils.Min(rs, totalSteps)
		// fmt.Println("rs = ", rs)
	}
}
func print(rooms map[int]int, hallways string) {
	for i := 0; i < len(hallways); i++ {
		if hallways[i] != '0' {
			fmt.Print(roomToChar[int(hallways[i]-'0')])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
	for i := 0; i < len(hallways); i++ {
		if _, ok := rooms[i]; ok {
			if rooms[i]/10 == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(roomToChar[rooms[i]/10])
			}
		} else {
			fmt.Print("#")
		}
	}
	fmt.Println()
	for i := 0; i < len(hallways); i++ {
		if _, ok := rooms[i]; ok {
			if rooms[i]%10 == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(roomToChar[rooms[i]%10])
			}
		} else {
			fmt.Print("#")
		}
	}
	fmt.Println()
}
func copyMap[K, V comparable](s map[K]V) map[K]V {
	copyS := make(map[K]V)
	for k, v := range s {
		copyS[k] = v
	}
	return copyS
}

// Part2 func
func (d Day23) Part2() {
}
