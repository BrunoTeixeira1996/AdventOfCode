package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	A []string
	B []string

	FullA []int
	FullB []int
}

func read_file(file string, slice *[]string) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		*slice = append(*slice, s.Text())
	}
}

func expandStringToRange(str []string) []int {
	startS := strings.Split(str[0], "-")[0]
	endS := strings.Split(str[0], "-")[1]

	startI, err := strconv.Atoi(startS)

	if err != nil {
		panic(err)
	}

	endI, err := strconv.Atoi(endS)

	if err != nil {
		panic(err)
	}

	var fullRange []int
	for i := startI; i <= endI; i++ {
		fullRange = append(fullRange, i)
	}

	return fullRange

}

func buildRanges(input []string) []Range {
	var ranges []Range

	for _, values := range input {
		var r Range
		for _, value := range strings.Split(values, "\n") {
			r.A = append(r.A, strings.Split(value, ",")[0])
			r.B = append(r.B, strings.Split(value, ",")[1])

			r.FullA = expandStringToRange(r.A)
			r.FullB = expandStringToRange(r.B)
		}
		ranges = append(ranges, r)
	}
	return ranges
}


func overlaps(a, b []int) bool{

	for _, valueA := range a {
		for _, valueB := range b {
				if valueA == valueB {
					return true
				}
			}
		}

	return false
}


func main() {
	var input []string
	read_file("big", &input)
	data := buildRanges(input)
	sum := 0
	for _, d := range data {
		if overlaps(d.FullA, d.FullB){
			sum += 1
		}
	}

	fmt.Printf("Res: %d\n", sum)
}
