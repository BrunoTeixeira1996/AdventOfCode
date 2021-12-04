package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file(file string, slice *[][]string) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		aux := strings.Split(s.Text(), "\n")
		*slice = append(*slice, aux)
	}

}

func convert_bin_to_dec(gama_rate string, epsilon_rate string) (int64, int64) {
	gama_dec, _ := strconv.ParseInt(gama_rate, 2, 64)
	epsilon_dec, _ := strconv.ParseInt(epsilon_rate, 2, 64)

	return gama_dec, epsilon_dec

}

func calc(aux string) (string, string) {
	count_zero, count_one := strings.Count(aux, "0"), strings.Count(aux, "1")

	gama_rate, epsilon_rate := "", ""

	if count_zero > count_one {
		gama_rate = "0"
		epsilon_rate = "1"
	} else {
		gama_rate = "1"
		epsilon_rate = "0"
	}

	return gama_rate, epsilon_rate

}

func count_part_one(slice [][]string) int64 {

	var aux string
	gama_rate, epsilon_rate := "", ""
	gama_rate_temp, epsilon_rate_temp := "", ""

	for i := 0; i < len(slice[0][0]); i++ {
		for index, _ := range slice {
			aux += string(slice[index][0][i])
		}
		gama_rate_temp, epsilon_rate_temp = calc(aux)

		gama_rate += gama_rate_temp
		epsilon_rate += epsilon_rate_temp
		aux = ""
	}

	gama, epsilon := convert_bin_to_dec(gama_rate, epsilon_rate)

	return gama * epsilon

}

func main() {
	var slice_of_slices [][]string

	read_file("input", &slice_of_slices)
	anser := count_part_one(slice_of_slices)

	fmt.Println(anser)
}
