package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func day5() {
	file, err := os.Open("input5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	ins := strings.Split(scanner.Text(), ",")

	originalIns := make([]int64, len(ins))
	for i := range originalIns {
		originalIns[i], _ = strconv.ParseInt(ins[i], 10, 64)
	}
	process(originalIns, os.Stdin, os.Stdout)
}
