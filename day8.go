package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day8() {
	file, err := os.Open("input8.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := []rune(scanner.Text())
	// fmt.Println(ins)
	part81(ins)
	result := part82(ins, 25, 6)
	for i := 0; i < len([]rune(result)); i++ {
		if result[i] == '0' {
			fmt.Print(". ")
		} else {
			fmt.Print("x ")
		}
		if (i+1)%(25) == 0 {
			fmt.Print("\n")
		}
	}
}
func part81(ins []rune) {
	numLayer := len(ins) / (25 * 6)
	layers := make([]map[int]int, numLayer)
	for i := range layers {
		layers[i] = make(map[int]int)
	}
	for i := 0; i < len(ins); i++ {
		layerth := i / (25 * 6)
		val := int(ins[i] - '0')
		layers[layerth][val]++
	}
	min0 := len(ins)
	result := 0
	for _, l := range layers {
		if l[0] < min0 {
			result = l[1] * l[2]
			min0 = l[0]
		}
	}
	fmt.Println(result)
}

func part82(ins []rune, wide, tall int) string {

	decodedImage := make([]rune, wide*tall)

	for i := 0; i < len(ins); i++ {
		layerth := i / (wide * tall)
		val := ins[i]
		postition := i % (wide * tall)

		if layerth == 0 || decodedImage[postition] == '2' {
			decodedImage[postition] = val
		}
	}
	return string(decodedImage)
}
