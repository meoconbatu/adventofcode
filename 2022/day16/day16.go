package day16

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day16 type
type Day16 struct{}

// Part1 func
func (d Day16) Part1() {
	vertices, edgestemp, indexes := readInput()
	edges := make([][]int, len(vertices))
	for u, vs := range edgestemp {
		for _, v := range vs {
			edges[u] = append(edges[u], indexes[v])
		}
	}
	rs := 0
	dist := floyd(vertices, edges)
	dfs(vertices, dist, indexes["AA"], 0, 30, &rs, make([]bool, len(vertices)), "0")
	fmt.Println(rs)
}
func dfs(vertices []int, dist [][]int, u, total, minute int, rs *int, visited []bool, path string) {
	*rs = utils.Max(*rs, total)
	if minute <= 0 {
		return
	}
	visited[u] = true
	for v := 0; v < len(dist); v++ {
		if dist[u][v] > 0 && !visited[v] && vertices[v] > 0 {
			dfs(vertices, dist, v, total+(minute-dist[u][v]-1)*vertices[v], minute-dist[u][v]-1, rs, visited, fmt.Sprintf("%s->%d(%d)", path, v, vertices[v]))
		}
	}
	visited[u] = false
}
func floyd(vertices []int, edges [][]int) [][]int {
	dist := make([][]int, len(vertices))
	for i := 0; i < len(vertices); i++ {
		dist[i] = make([]int, len(vertices))
		for j := 0; j < len(vertices); j++ {
			dist[i][j] = math.MaxInt32
		}
	}
	for u, vs := range edges {
		for _, v := range vs {
			dist[u][v] = 1
		}
	}
	for v := range vertices {
		dist[v][v] = 0
	}
	for k := 0; k < len(vertices); k++ {
		for i := 0; i < len(vertices); i++ {
			for j := 0; j < len(vertices); j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

var dp = make(map[int]int)

// Part2 func
func (d Day16) Part2() {
	vertices, edgestemp, indexes := readInput()
	edges := make([][]int, len(vertices))
	for u, vs := range edgestemp {
		for _, v := range vs {
			edges[u] = append(edges[u], indexes[v])
		}
	}
	rs := 0
	dist := floyd(vertices, edges)
	for i := 1; i <= 15; i++ {
		dfs2(vertices, dist, indexes["AA"], 0, 26, make([]bool, len(vertices)), i, 0)
	}
	dp[0] = 0
	for v1, rs1 := range dp {
		for v2, rs2 := range dp {
			if v1&v2 == 0 {
				rs = utils.Max(rs, rs1+rs2)
			}
		}
	}
	fmt.Println(rs)
}
func dfs2(vertices []int, dist [][]int, u, total, minute int, visited []bool, target int, vs int) {
	if target == 0 && minute >= 0 {
		dp[vs] = utils.Max(dp[vs], total)
		return
	}
	if minute <= 0 {
		return
	}
	visited[u] = true
	for v := 0; v < len(dist); v++ {
		if dist[u][v] > 0 && !visited[v] && vertices[v] > 0 {
			dfs2(vertices, dist, v, total+(minute-dist[u][v]-1)*vertices[v], minute-dist[u][v]-1, visited, target-1, vs|(1<<v))
		}
	}
	visited[u] = false
}

func readInput() ([]int, map[int][]string, map[string]int) {
	scanner := utils.NewScanner(16)
	edges := make(map[int][]string)
	vertices := make([]int, 0)
	indexes := make(map[string]int)
	i := 0
	for scanner.Scan() {
		re := regexp.MustCompile(`Valve (.*) has flow rate=(.*); (tunnels|tunnel) (lead|leads) to (valves|valve) (.*)`)
		subs := re.FindAllStringSubmatch(scanner.Text(), -1)
		rate, _ := strconv.Atoi(subs[0][2])
		vertices = append(vertices, rate)
		indexes[subs[0][1]] = i
		for _, v := range regexp.MustCompile(`[A-Z]{2}`).FindAllStringSubmatch(subs[0][6], -1) {
			edges[i] = append(edges[i], v[0])
		}
		i++
	}
	return vertices, edges, indexes
}
