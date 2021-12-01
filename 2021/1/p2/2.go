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

func sum_aux_slice(slice []int) int {
	result := 0
	for _, v := range slice {
		result += v
	}

	return result
}

func sum_values(original_slice []int) []int {
	var aux []int

	for index := 0; index < len(original_slice)-2; index++ {
		sum := sum_aux_slice(original_slice[index : index+3])
		aux = append(aux, sum)
		sum = 0
	}

	return aux
}

func count_part_two(slice []int) int {
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
	aux := sum_values(slice)
	count := count_part_two(aux)
	fmt.Println(count)
}
