package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read_file(file string, slice *[]int) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		integer, _ := strconv.Atoi(s.Text())
		*slice = append(*slice, integer)
	}

}

func count_part_one(slice []int) int {
	count := 0

	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			count++
		}
	}

	return count

}

func main() {
	var slice []int

	read_file("input", &slice)
	count := count_part_one(slice)
	fmt.Println(count)
}
