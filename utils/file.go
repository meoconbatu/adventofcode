package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetInputFile func download input file from website if the file does not exist
// set force = true to always download
func GetInputFile(day int, session string, force bool) {
	if _, err := os.Stat(fmt.Sprintf("./day%d/input.txt", day)); err == nil && !force {
		return
	}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day), nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat("/path/to/your-file"); os.IsNotExist(err) {
		os.MkdirAll(fmt.Sprintf("./day%d", day), 0700)
	}
	err = ioutil.WriteFile(fmt.Sprintf("./day%d/input.txt", day), body, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
