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

/*
Creates a map with {int : int} with the range of the second slice 
and assigns 1 has the value of each int

Iterates the first slice and checks if the value (of the first slice)
is present on the second, if yes, count = 1 and found = true
and then subtract - 1 in the set
*/
func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}
	return true
}

func main() {
	var input []string
	read_file("big", &input)
	data := buildRanges(input)
	sum := 0
	for _, d := range data {
		if subset(d.FullA, d.FullB) || subset(d.FullB, d.FullA){
			sum += 1
		}
	}

	fmt.Printf("Res: %d\n", sum)


}
