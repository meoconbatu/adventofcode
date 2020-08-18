package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type chemical struct {
	name  string
	level int
}
type material struct {
	*chemical
	quantity int
}
type reaction struct {
	outMaterial material
	inMaterials []material
}

func (m material) String() string {
	return fmt.Sprintf("%d (%s, %d)", m.quantity, m.name, m.level)
}
func day14() {
	file, err := os.Open("input14.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// fmt.Println(day14Part1(file))
	fmt.Println(day14Part2(file))
}

func newMaterial() material {
	c := new(chemical)
	return material{c, 0}
}
func readInputDay14(in io.Reader) map[string]reaction {
	scanner := bufio.NewScanner(in)

	originalIns := make(map[string]reaction)
	for scanner.Scan() {
		mOut := newMaterial()
		s := strings.Split(scanner.Text(), " => ")
		fmt.Sscanf(s[1], "%d %s", &mOut.quantity, &mOut.name)

		mInStrs := strings.Split(s[0], ", ")
		mIns := make([]material, 0)
		for _, in := range mInStrs {
			m := newMaterial()
			fmt.Sscanf(in, "%d %s", &m.quantity, &m.name)
			if v, ok := originalIns[m.name]; !ok {
				if m.name != "ORE" {
					originalIns[m.name] = reaction{m, nil}
				}
			} else {
				m.chemical = v.outMaterial.chemical
			}

			mIns = append(mIns, m)
		}
		if v, ok := originalIns[mOut.name]; ok {
			originalIns[mOut.name] = reaction{material{v.outMaterial.chemical, mOut.quantity}, mIns}
		} else {
			originalIns[mOut.name] = reaction{mOut, mIns}
		}
	}
	return originalIns
}
func day14Part1(in io.Reader) int {
	originalIns := readInputDay14(in)
	level(originalIns)
	return calc(originalIns, 1)
}

func day14Part2(in io.Reader) int {
	originalIns := readInputDay14(in)
	level(originalIns)

	numFuel := 1000000000000 / calc(originalIns, 1)
	numORE := 0
	for numORE <= 1000000000000 {
		numORE = calc(originalIns, numFuel)
		if numORE >= 1000000000000 {
			numFuel--
			break
		}
		numFuel++
	}
	return numFuel
}
func level(maps map[string]reaction) {
	for _, v := range maps {
		if v.inMaterials[0].name == "ORE" {
			v.outMaterial.level = 1
		}
	}
	for {
		end := 0
		for _, v := range maps {
			if v.outMaterial.level != 0 {
				end++
				continue
			}
			l := 0
			for _, in := range v.inMaterials {
				if in.level == 0 {
					l = 0
					break
				}
				if in.level >= l {
					l = in.level + 1
				}
			}
			if l != 0 {
				v.outMaterial.level = l
			}
		}
		if end == len(maps) {
			return
		}
	}
}
func calc(maps map[string]reaction, min int) int {
	result := make(map[string]material)
	level := maps["FUEL"].outMaterial.level

	for _, v := range maps["FUEL"].inMaterials {
		result[v.name] = material{v.chemical, min * v.quantity}
	}
	rerun := false
	for l := level; l > 0; l-- {
		for _, v := range result {
			if v.level == l {
				teardown(maps, result, v)
				rerun = true
			}
		}
		if rerun {
			l++
			rerun = false
		}
	}
	return result["ORE"].quantity
}
func teardown(maps map[string]reaction, result map[string]material, m material) {
	for _, v := range maps[m.name].inMaterials {
		result[v.name] = material{v.chemical, v.quantity*int(math.Ceil(float64(m.quantity)/float64(maps[m.name].outMaterial.quantity))) + result[v.name].quantity}
	}
	delete(result, m.name)
}
