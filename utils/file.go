package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// GetInputFile func download input file from website if the file does not exist
// set force = true to always download
func GetInputFile(year, day int, session string, force bool) {
	if day == 0 {
		return
	}
	if _, err := os.Stat(fmt.Sprintf("./day%d/input.txt", day)); err == nil && !force {
		return
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("/path/to/your-file"); os.IsNotExist(err) {
		os.MkdirAll(fmt.Sprintf("./day%d", day), 0700)
	}
	err = os.WriteFile(fmt.Sprintf("./day%d/input.txt", day), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func NewScanner(dayth int) *bufio.Scanner {
	f, err := os.Open(fmt.Sprintf("day%d/input.txt", dayth))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return bufio.NewScanner(f)
}
