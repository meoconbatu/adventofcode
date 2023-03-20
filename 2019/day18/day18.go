package day18

import (
	"fmt"
	"math"
	"strings"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day18 type
type Day18 struct{}

// Part1 func
func (d Day18) Part1() {
	tunnels := readInput()
	fmt.Println(getMinStep(tunnels, "@"))
}

// Part2 func
func (d Day18) Part2() {
	tunnels := readInput()
	for i, tunnel := range tunnels {
		if idx := strings.Index(tunnel, "@"); idx != -1 {
			tunnels[i-1] = tunnels[i-1][:idx-1] + "@#]" + tunnels[i-1][idx+2:]
			tunnels[i] = tunnel[:idx-1] + "###" + tunnel[idx+2:]
			tunnels[i+1] = tunnels[i+1][:idx-1] + "[#_" + tunnels[i+1][idx+2:]
			break
		}
	}
	fmt.Println(getMinStep(tunnels, "@[]_"))
}

var dp map[int]int
func getMinStep(tunnels []string, entrance string) int {
	dp = make(map[int]int)
	edges, keys := initMap(tunnels)
	return dfs(edges, entrance, entrance, 0, keys, 0)
}

func dfs(edges map[int][]int, entrance, prevEntrance string, holdingKeys, keys int, step int) int {
	if holdingKeys == keys {
		return 0
	}
	key := holdingKeys*100000000 + entranceToInt(prevEntrance)
	if val, ok := dp[key]; ok {
		return val
	}
	minStep := math.MaxInt64
	for j := 0; j < len(entrance); j++ {
		u := int(entrance[j] - '@')
		for i := 0; i < len(edges[u]); i++ {
			v, currentStep := i, edges[u][i]
			if edges[u][i] == 0 || (isDoor(v) && holdingKeys&(1<<(v+'@'+32-'a')) == 0) || isEntrance(v) {
				continue
			}
			newEdges := removeEdge(edges, u, v, currentStep)
			newEdges = deleteDoor(newEdges, v-32)
			temp := dfs(newEdges, entrance, prevEntrance[:j]+string(byte(v+'@'))+prevEntrance[j+1:], holdingKeys|(1<<(v+'@'-'a')), keys, step+currentStep)
			if temp != math.MaxInt64 {
				temp += currentStep
			}
			minStep = utils.Min(minStep, temp)
		}
	}
	dp[key] = minStep
	return minStep
}
func entranceToInt(ent string) int {
	rs := 0
	for _, e := range ent {
		rs = rs*58 + int(e-'@')
	}
	return rs
}
func deleteDoor(edges map[int][]int, u int) map[int][]int {
	newEdges := copyMap(edges)
	for i := 0; i < len(newEdges[u]); i++ {
		if newEdges[u][i] == 0 || i == u {
			continue
		}
		for j := i + 1; j < len(newEdges[u]); j++ {
			if newEdges[u][j] == 0 || j == u {
				continue
			}
			if newEdges[i][j] > 0 {
				newEdges[i][j] = utils.Min(newEdges[i][j], newEdges[i][u]+newEdges[u][j])
			} else {
				newEdges[i][j] = newEdges[i][u] + newEdges[u][j]
			}
			newEdges[j][i] = newEdges[i][j]
		}
	}
	for i := 0; i < len(newEdges[u]); i++ {
		if u < len(newEdges[i]) && newEdges[i][u] > 0 {
			newEdges[i][u] = 0
		}
	}
	delete(newEdges, u)
	return newEdges
}
func print(edges map[int][]int) {
	for door, path := range edges {
		fmt.Printf("[%s:", string(byte(door+'@')))
		for i := 0; i < len(path); i++ {
			if path[i] == 0 {
				continue
			}
			fmt.Printf("%s(%d),", string(byte(i+'@')), path[i])
		}
		fmt.Printf("] ")
	}
	fmt.Println()
}

func removeEdge(edges map[int][]int, u, v, weight int) map[int][]int {
	newEdges := deleteDoor(edges, u)
	newEdges[u] = newEdges[v]
	for i := 0; i < len(newEdges[v]); i++ {
		if newEdges[v][i] > 0 {
			newEdges[u][i], newEdges[i][u] = newEdges[v][i], newEdges[v][i]
			newEdges[i][v] = 0
		}
	}
	delete(newEdges, v)
	return newEdges
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

func isDoor(v int) bool {
	return v+'@' >= 'A' && v+'@' <= 'Z'
}

func isEntrance(v int) bool {
	return !((v+'@' >= 'A' && v+'@' <= 'Z') || (v+'@' >= 'a' && v+'@' <= 'z'))
}

func initMap(tunnels []string) (map[int][]int, int) {
	edges := make(map[int][]int)
	keys := 0
	for i := 0; i < len(tunnels); i++ {
		for j := 0; j < len(tunnels[i]); j++ {
			if tunnels[i][j] != '#' && tunnels[i][j] != '.' {
				edges[int(tunnels[i][j]-'@')] = findPaths(tunnels, int(tunnels[i][j]), i, j)
				if tunnels[i][j] >= 'a' && tunnels[i][j] <= 'z' {
					keys |= 1 << (tunnels[i][j] - 'a')
				}
			}
		}
	}
	return edges, keys
}

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func findPaths(tunnels []string, start, startx, starty int) []int {
	q := []int{startx, starty, 0}
	paths := make([]int, 0)
	visited := make(map[int]bool)
	for len(q) > 0 {
		x, y, step := q[0], q[1], q[2]
		q = q[3:]
		for _, d := range directions {
			newx, newy := x+d[0], y+d[1]
			if newx >= len(tunnels) || newy >= len(tunnels[0]) || newx < 0 || newy < 0 || visited[newx*100+newy] || tunnels[newx][newy] == '#' {
				continue
			}
			visited[newx*100+newy] = true
			if tunnels[newx][newy] != '.' && tunnels[newx][newy] != byte(start) {
				paths = append(paths, int(tunnels[newx][newy]), step+1)
			} else {
				q = append(q, newx, newy, step+1)
			}
		}
	}
	rs := make([]int, 'z'-'@'+1)
	for i := 0; i < len(paths); i += 2 {
		rs[paths[i]-'@'] = paths[i+1]
	}
	return rs
}

func readInput() []string {
	scanner := utils.NewScanner(18)
	rs := make([]string, 0)
	for scanner.Scan() {
		rs = append(rs, scanner.Text())
	}
	return rs
}
