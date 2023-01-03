package day19

import (
	"fmt"

	"github.com/meoconbatu/adventofcode/utils"
)

type Day19 struct {
}

type Point struct {
	x, y, z int
}

// Part1 func
func (d Day19) Part1() {
	scanners := readInput()
	beacons, _ := getAllBeacons(scanners)
	fmt.Println(len(beacons))
}
func getAllBeacons(scanners [][]Point) (map[int]struct{}, []Point) {
	beacons := make(map[int]struct{})
	scannerPoints := make([]Point, len(scanners))
	for _, beacon := range scanners[0] {
		beacons[mapPointToInt(beacon)] = struct{}{}
	}
	connecteds := make([]bool, len(scanners))
	q := []int{0}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for j := 0; j < len(scanners); j++ {
			if connecteds[j] {
				continue
			}
			for k := 0; k < 24; k++ {
				newscannerj := rotate(scanners[j], k)
				if ok, scannerj := checkOverlap(beacons, scannerPoints[i], scanners[i], newscannerj); ok {
					q = append(q, j)
					connecteds[j] = true
					copy(scanners[j], newscannerj)
					scannerPoints[j] = scannerj
					break
				}
			}
		}
	}
	return beacons, scannerPoints
}

func mapPointToInt(p Point) int {
	min := -6000
	m := 6000
	return (p.x-min)*m*m + (p.y-min)*m + p.z - min
}
func mapIntToXYZ(num int) (int, int, int) {
	min := -6000
	m := 6000
	return (num/m)/m + min, (num/m)%m + min, num%m + min
}

// Part2 func
func (d Day19) Part2() {
	scanners := readInput()
	_, scannerPoints := getAllBeacons(scanners)
	rs := 0
	for _, p1 := range scannerPoints {
		for _, p2 := range scannerPoints {
			rs = utils.Max(rs, utils.Abs(p1.x-p2.x)+utils.Abs(p1.y-p2.y)+utils.Abs(p1.z-p2.z))
		}
	}
	fmt.Println(rs)
}
func checkOverlap(beacons map[int]struct{}, scanner1 Point, beacon1s, beacon2s []Point) (bool, Point) {
	for i1 := 0; i1 < len(beacon1s); i1++ {
		for i2 := i1 + 1; i2 < len(beacon2s); i2++ {
			scanner2 := getScannerPoint(beacon1s[i1], beacon2s[i2])
			numOverlap := 0
			for _, bc2 := range beacon2s {
				newbc2 := getBeaconPoint(bc2, scanner2)
				if _, ok := beacons[mapPointToInt(newbc2)]; !ok {
					continue
				}
				numOverlap++
				if numOverlap < 12 {
					continue
				}
				for i, bc := range beacon2s {
					beacon2s[i] = getBeaconPoint(bc, scanner2)
					beacons[mapPointToInt(beacon2s[i])] = struct{}{}
				}
				return true, scanner2
			}
		}
	}
	return false, Point{}
}
func getScannerPoint(beacon1, beacon2 Point) Point {
	return Point{getScannerAxis(beacon1.x, beacon2.x), getScannerAxis(beacon1.y, beacon2.y), getScannerAxis(beacon1.z, beacon2.z)}
}
func getScannerAxis(a, b int) int {
	if b <= 0 && a <= 0 {
		return utils.Abs(b) - utils.Abs(a)
	} else if a >= 0 && b >= 0 {
		return utils.Abs(a) - utils.Abs(b)
	} else if a <= 0 && b >= 0 {
		return -(utils.Abs(a) + utils.Abs(b))
	} else if a >= 0 && b <= 0 {
		return (utils.Abs(a) + utils.Abs(b))
	}
	fmt.Println(a, b)
	return 0
}

func getBeaconPoint(beacon, scanner Point) Point {
	return Point{getBeaconAxis(beacon.x, scanner.x), getBeaconAxis(beacon.y, scanner.y), getBeaconAxis(beacon.z, scanner.z)}
}
func getBeaconAxis(beacon, scanner int) int {
	if scanner <= 0 && beacon <= 0 {
		return -(utils.Abs(scanner) + utils.Abs(beacon))
	} else if beacon >= 0 && scanner >= 0 {
		return utils.Abs(beacon) + utils.Abs(scanner)
	} else if beacon <= 0 && scanner >= 0 {
		return -(utils.Abs(beacon) - utils.Abs(scanner))
	} else if beacon >= 0 && scanner <= 0 {
		return utils.Abs(beacon) - utils.Abs(scanner)
	}
	fmt.Println(beacon, scanner)
	return 0
}

func readInput() [][]Point {
	scanner := utils.NewScanner(19)
	scanners := make([][]Point, 0)
	var i int
	for scanner.Scan() {
		s := scanner.Text()

		fmt.Sscanf(s, "--- scanner %d ---", &i)
		beacons := make([]Point, 0)
		for {
			scanner.Scan()
			ss := scanner.Text()
			if len(ss) == 0 {
				break
			}
			var x, y, z int
			fmt.Sscanf(ss, "%d,%d,%d", &x, &y, &z)
			beacons = append(beacons, Point{x, y, z})
		}
		scanners = append(scanners, beacons)
	}
	return scanners
}

func rotate(beacons []Point, k int) []Point {
	rs := make([]Point, len(beacons))
	for i, bc := range beacons {
		switch k {
		case 0:
			rs[i] = Point{bc.x, bc.y, bc.z}
		case 1:
			rs[i] = Point{bc.x, bc.z, -bc.y}
		case 2:
			rs[i] = Point{bc.x, -bc.y, -bc.z}
		case 3:
			rs[i] = Point{bc.x, -bc.z, bc.y}
		case 4:
			rs[i] = Point{bc.y, bc.z, bc.x}
		case 5:
			rs[i] = Point{bc.y, bc.x, -bc.z}
		case 6:
			rs[i] = Point{bc.y, -bc.z, -bc.x}
		case 7:
			rs[i] = Point{bc.y, -bc.x, bc.z}
		case 8:
			rs[i] = Point{bc.z, bc.y, -bc.x}
		case 9:
			rs[i] = Point{bc.z, -bc.x, -bc.y}
		case 10:
			rs[i] = Point{bc.z, -bc.y, bc.x}
		case 11:
			rs[i] = Point{bc.z, bc.x, bc.y}
		case 12:
			rs[i] = Point{-bc.x, bc.y, -bc.z}
		case 13:
			rs[i] = Point{-bc.x, -bc.z, -bc.y}
		case 14:
			rs[i] = Point{-bc.x, -bc.y, bc.z}
		case 15:
			rs[i] = Point{-bc.x, bc.z, bc.y}
		case 16:
			rs[i] = Point{-bc.y, bc.z, -bc.x}
		case 17:
			rs[i] = Point{-bc.y, -bc.x, -bc.z}
		case 18:
			rs[i] = Point{-bc.y, -bc.z, bc.x}
		case 19:
			rs[i] = Point{-bc.y, bc.x, bc.z}
		case 20:
			rs[i] = Point{-bc.z, bc.y, bc.x}
		case 21:
			rs[i] = Point{-bc.z, bc.x, -bc.y}
		case 22:
			rs[i] = Point{-bc.z, -bc.y, -bc.x}
		case 23:
			rs[i] = Point{-bc.z, -bc.x, bc.y}
		}
	}
	return rs
}
