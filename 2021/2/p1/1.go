package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file(file string, sol *map[string]int) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		words := strings.Fields(s.Text())
		value, _ := strconv.Atoi(words[1])
		if words[0] == "forward" {
			(*sol)["horizontal_pos"] += value
		} else if words[0] == "down" {
			(*sol)["depth"] += value
		} else {
			(*sol)["depth"] -= value
		}
	}

}

func count_part_one(sol map[string]int) int {
	anser := sol["horizontal_pos"] * sol["depth"]
	return anser

}

func main() {
	sol := map[string]int{
		"horizontal_post": 0,
		"depth":           0,
	}

	read_file("input", &sol)

	anser := count_part_one(sol)
	fmt.Println(anser)
}
