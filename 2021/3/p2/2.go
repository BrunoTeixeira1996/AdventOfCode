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

func remove_oxygen_slice(oxygen_slice *[][]string, pos, count_zero, count_one int) {

	if count_zero > count_one {
		for i := 0; i < len(*oxygen_slice); {
			if (*oxygen_slice)[i][0][pos] == 49 {
				(*oxygen_slice)[i] = (*oxygen_slice)[len(*oxygen_slice)-1]
				*oxygen_slice = (*oxygen_slice)[:len(*oxygen_slice)-1]
			} else {
				i++
			}

		}

	} else if count_zero <= count_one {
		for i := 0; i < len(*oxygen_slice); {
			if (*oxygen_slice)[i][0][pos] == 48 {
				(*oxygen_slice)[i] = (*oxygen_slice)[len(*oxygen_slice)-1]
				*oxygen_slice = (*oxygen_slice)[:len(*oxygen_slice)-1]
			} else {
				i++
			}
		}
	}
}

func remove_co2_slice(co2_slice *[][]string, pos, count_zero, count_one int) {

	if count_zero > count_one {
		for i := 0; i < len(*co2_slice); {
			if (*co2_slice)[i][0][pos] == 48 {
				(*co2_slice)[i] = (*co2_slice)[len(*co2_slice)-1]
				*co2_slice = (*co2_slice)[:len(*co2_slice)-1]
			} else {
				i++
			}

		}
	} else {
		for i := 0; i < len(*co2_slice); {
			if (*co2_slice)[i][0][pos] == 49 {
				(*co2_slice)[i] = (*co2_slice)[len(*co2_slice)-1]
				*co2_slice = (*co2_slice)[:len(*co2_slice)-1]
			} else {
				i++
			}
		}
	}
}

func rating(oxygen_slice *[][]string, co2_slice *[][]string) {
	pos := 0
	byte_size := len((*oxygen_slice)[0][0])
	count_zero_o2, count_one_o2 := 0, 0
	count_zero_co2, count_one_co2 := 0, 0

	for pos < byte_size {
		for i := 0; i < len(*oxygen_slice); i++ {
			if string((*oxygen_slice)[i][0][pos]) == "0" {
				count_zero_o2++
			} else {
				count_one_o2++
			}
		}

		for i := 0; i < len(*co2_slice); i++ {
			if string((*co2_slice)[i][0][pos]) == "0" {
				count_zero_co2++
			} else {
				count_one_co2++
			}

		}

		if len(*oxygen_slice) > 1 {
			remove_oxygen_slice(oxygen_slice, pos, count_zero_o2, count_one_o2)
		}
		if len(*co2_slice) > 1 {
			remove_co2_slice(co2_slice, pos, count_zero_co2, count_one_co2)
		}
		count_zero_o2, count_one_o2 = 0, 0
		count_zero_co2, count_one_co2 = 0, 0
		pos++
	}
}
func convert_bin_to_dec(oxygen, co2 string) (int64, int64) {
	oxygen_dec, _ := strconv.ParseInt(oxygen, 2, 64)
	co2_dec, _ := strconv.ParseInt(co2, 2, 64)

	return oxygen_dec, co2_dec

}

func count_part_two(oxygen_slice [][]string, co2_slice [][]string) int64 {
	oxygen, co2 := oxygen_slice[0][0], co2_slice[0][0]

	oxygen_v, co2_v := convert_bin_to_dec(oxygen, co2)

	return oxygen_v * co2_v
}

func main() {
	var slice_of_slices [][]string
	read_file("input", &slice_of_slices)

	oxygen_slice := make([][]string, len(slice_of_slices))

	co2_slice := make([][]string, len(slice_of_slices))

	copy(oxygen_slice, slice_of_slices)
	copy(co2_slice, slice_of_slices)

	rating(&oxygen_slice, &co2_slice)

	anser := count_part_two(oxygen_slice, co2_slice)

	fmt.Println(anser)
}
