package day20

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

// Day20 type
type Day20 struct{}

// Part1 func
func (d Day20) Part1() {
	keyToGrid := readInput()
	idOrientToEdges := make(map[int][]int)
	m, n := utils.Sqrt(len(keyToGrid)), utils.Sqrt(len(keyToGrid))
	for key, grid := range keyToGrid {
		idOrientToEdges[key*10] = parseToEdges(grid)
		genEdgeToID(idOrientToEdges[key*10], key)
		for i := 1; i <= 7; i++ {
			if i != 4 {
				idOrientToEdges[key*10+i] = rotate(idOrientToEdges[key*10+i-1])
			} else {
				idOrientToEdges[key*10+i] = flip(idOrientToEdges[key*10+i-1])
			}

		}
	}
	dp = make(map[int][]int)
	gen(idOrientToEdges)
	IDToCount := make(map[int]int)
	for _, IDs := range edgesToID {
		if len(IDs) == 1 {
			continue
		}
		for _, id := range IDs {
			IDToCount[id]++
		}
	}
	for id, count := range IDToCount {
		if count == 4 {
			for i := 0; i <= 7; i++ {
				dp[0] = append(dp[0], id*10+i)
			}

		}
	}
	// fmt.Println(IDToCount)
	// fmt.Println(edgesToID)
	fmt.Println(dfs([12][12]int{}, m, n, 0, 0, make(map[int]bool)))
}

var edgesToID = make(map[int][]int)

func genEdgeToID(edges []int, ID int) {
	for _, edge := range edges {
		edgesToID[edge] = append(edgesToID[edge], ID)
		edgesToID[reverse(edge)] = append(edgesToID[reverse(edge)], ID)
	}
}
func gen(idOrientToEdges map[int][]int) {
	for idOrient1, edges1 := range idOrientToEdges {
		// dp[0] = append(dp[0], idOrient1)
		for idOrient2, edges2 := range idOrientToEdges {
			if idOrient2/10 == idOrient1/10 {
				continue
			}
			if edges1[1] == edges2[0] {
				dp[idOrient1*100000] = append(dp[idOrient1*100000], idOrient2)
			}
			if edges1[3] == edges2[2] {
				dp[idOrient1] = append(dp[idOrient1], idOrient2)
			}
			for idOrient3, edges3 := range idOrientToEdges {
				if idOrient3/10 == idOrient2/10 || idOrient3/10 == idOrient1/10 {
					continue
				}
				if edges1[1] == edges3[0] && edges2[3] == edges3[2] {
					k := idOrient1*100000 + idOrient2
					dp[k] = append(dp[k], idOrient3)
				}
			}
		}
	}
}

var dp map[int][]int

func dfs(grid [12][12]int, m, n, x, y int, visited map[int]bool) int {
	// if x < m {
	// 	return 0
	// }
	if x >= m {
		// fmt.Println(grid)
		return (grid[0][0] / 10) * (grid[m-1][0] / 10) * (grid[m-1][n-1] / 10) * (grid[0][n-1] / 10)
	}
	var a, b int
	if x > 0 {
		a = grid[x-1][y]
	}
	if y > 0 {
		b = grid[x][y-1]
	}
	k := a*100000 + b
	for _, idOrient := range dp[k] {
		// fmt.Println(x, y, idOrient)
		id := idOrient / 10
		if visited[id] {
			continue
		}
		grid[x][y] = idOrient
		visited[id] = true
		if temp := dfs(grid, m, n, x+(y+1)/n, (y+1)%n, visited); temp > 0 {
			return temp
		}
		grid[x][y] = 0
		visited[id] = false
	}
	return 0
}
func flip(edges []int) []int {
	rs := make([]int, len(edges))
	rs[0] = reverse(edges[0])
	rs[1] = reverse(edges[1])
	rs[2], rs[3] = edges[3], edges[2]
	return rs
}

func rotate(edges []int) []int {
	rs := make([]int, len(edges))
	rs[0] = reverse(edges[2])
	rs[2] = edges[1]
	rs[1] = reverse(edges[3])
	rs[3] = edges[0]
	return rs
}
func reverse(num int) int {
	a := num
	rs := 0

	for n := 9; n >= 0; n-- {
		rs = (rs << 1) | (a & 1)
		a >>= 1
	}
	return rs
}

func parseToEdges(grid []string) []int {
	m, n := len(grid), len(grid[0])
	rs := make([]int, 4)
	for i := 0; i < m; i++ {
		if grid[0][i] == '#' {
			rs[0] |= (1 << (n - i - 1))
		}
		if grid[m-1][i] == '#' {
			rs[1] |= (1 << (m - i - 1))
		}
		if grid[i][0] == '#' {
			rs[2] |= (1 << (n - i - 1))
		}
		if grid[i][n-1] == '#' {
			rs[3] |= (1 << (m - i - 1))
		}
	}
	return rs
}

// Part2 func
func (d Day20) Part2() {
	keyToGrid := readInput()
	m := utils.Sqrt(len(keyToGrid))
	// grid := [][]int{{11715, 14897, 29717}, {24736, 14277, 27297}, {30792, 23117, 19517}}
	grid := [][]int{{16935, 28035, 30235, 30671, 11716, 11630, 19736, 28791, 19075, 16094, 14811, 22071}, {17831, 35593, 22511, 17095, 36731, 16134, 10870, 14393, 19337, 10910, 27497, 11233}, {16633, 19873, 38211, 31914, 23113, 27911, 32591, 28433, 14296, 28371, 39316, 34336}, {10633, 11870, 14510, 18477, 21377, 23774, 22971, 19974, 13190, 24231, 29630, 30497}, {31210, 20697, 19131, 10193, 26571, 15115, 10311, 28017, 17336, 13073, 30896, 22371}, {15715, 36233, 15597, 15671, 24777, 33470, 10971, 32713, 22397, 26631, 37936, 10394}, {23836, 17476, 34677, 29991, 14593, 11934, 18713, 20275, 33314, 10934, 17216, 30413}, {23714, 19993, 25513, 27074, 32537, 12016, 26775, 30617, 14714, 34073, 31094, 36371}, {33437, 37393, 37334, 37696, 25390, 39295, 34697, 29030, 18231, 36714, 25310, 33294}, {13995, 12491, 27894, 25216, 25917, 27535, 12773, 34991, 37793, 27674, 30017, 20114}, {25790, 12233, 35934, 28191, 10510, 19513, 38637, 36971, 29277, 35571, 16194, 35391}, {21114, 13615, 11293, 36591, 24737, 31190, 26591, 14994, 20835, 12916, 22673, 23395}}

	rs := make([][]rune, m*8)
	irs := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			key := grid[i][j] / 10
			g := get(keyToGrid[key], grid[i][j]%10)
			for l := 1; l < len(g)-1; l++ {
				rs[irs*8+l-1] = append(rs[irs*8+l-1], g[l][1:9]...)
			}
		}
		irs++
	}
	for i := 0; i < len(rs); i++ {
		fmt.Println(string(rs[i]))
	}
	for orient := 1; orient <= 8; orient++ {
		if orient != 4 {
			rotateImage(rs)
		} else {
			flipImage(rs)
		}

		total := 0
		for i := 0; i < len(rs); i++ {
			for j := 0; j < len(rs[i]); j++ {
				if found(rs, i, j) {
					total++
				}
			}
		}
		if total > 0 {
			fmt.Println(total)
			break
		}
	}
	for i := 0; i < len(rs); i++ {
		fmt.Println(string(rs[i]))
	}
	numHashtag := 0
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(rs[i]); j++ {
			if rs[i][j] == '#' {
				numHashtag++
			}
		}
	}
	fmt.Println(numHashtag)

}
func found(img [][]rune, x, y int) bool {
	monster := []string{"                  # ", "#    ##    ##    ###", " #  #  #  #  #  #   "}
	if len(monster)+x >= len(img) || len(monster[0])+y >= len(img[x]) {
		return false
	}
	for i := 0; i < len(monster); i++ {
		for j := 0; j < len(monster[i]); j++ {
			if monster[i][j] == '#' && img[i+x][j+y] != '#' {
				return false
			}
		}
	}
	for i := 0; i < len(monster); i++ {
		for j := 0; j < len(monster[i]); j++ {
			if monster[i][j] == '#' && img[i+x][j+y] == '#' {
				img[i+x][j+y] = 'O'
			}
		}
	}
	return true
}

func get(grid []string, orient int) [][]rune {
	rs := make([][]rune, len(grid))
	for i := 0; i < len(grid); i++ {
		rs[i] = []rune(grid[i])
	}
	for i := 1; i <= orient; i++ {
		if i != 4 {
			rotateImage(rs)
		} else {
			flipImage(rs)
		}
	}
	return rs
}
func printImage(grid [][]rune) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}
func rotateImage(matrix [][]rune) {
	n := len(matrix) - 1
	for i := 0; i <= n/2; i++ {
		for j := i; j <= (n - i - 1); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = temp
		}
	}
}
func flipImage(image [][]rune) {
	n := len(image)
	for i := 0; i < n; i++ {
		for j := 0; j < (n+1)/2; j++ {
			image[i][j], image[i][n-j-1] = image[i][n-j-1], image[i][j]
		}
	}
}
func readInput() map[int][]string {
	scanner := utils.NewScanner(20)
	rs := make(map[int][]string)
	for scanner.Scan() {
		var key int
		fmt.Sscanf(scanner.Text(), "Tile %d:\n", &key)
		for j := 0; j < 10; j++ {
			scanner.Scan()
			rs[key] = append(rs[key], scanner.Text())
		}
		scanner.Scan()
	}
	return rs
}

func print(edges []int, m, n int) {
	for i := 0; i < len(edges); i++ {
		fmt.Printf("%b,", edges[i])
	}
	a, b, c, d := edges[0], edges[1], edges[2], edges[3]
	fmt.Println()
	rs := make([][]rune, m)
	for i := m - 1; i >= 0; i-- {
		rs[i] = make([]rune, n)
		for j := n - 1; j >= 0; j-- {
			temp := '.'
			if i == 0 {
				if a&1 > 0 {
					temp = '#'
				}
				a >>= 1
			} else if i == m-1 {
				if b&1 > 0 {
					temp = '#'
				}
				b >>= 1
			}
			if j == 0 {
				if c&1 > 0 {
					temp = '#'
				}
				c >>= 1
			} else if j == n-1 {
				if d&1 > 0 {
					temp = '#'
				}
				d >>= 1
			}
			rs[i][j] = temp
		}
	}
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(rs[i]); j++ {
			fmt.Printf("%s", string(rs[i][j]))
		}
		fmt.Println()
	}
}
