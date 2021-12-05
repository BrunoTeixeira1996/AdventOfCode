package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_file(file string, numbers *[]string, new_card *[][]string, total_cards *[][][]string) {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	line := 0
	counter := 0
	for s.Scan() {
		if line == 0 {
			*numbers = strings.Split(s.Text(), ",")
		} else if line != 1 && s.Text() != "" {
			var temp []string
			aux := strings.Split(s.Text(), " ")
			// trim whitespace
			for i := range aux {
				if aux[i] != "" {
					temp = append(temp, aux[i])
				}
			}
			*new_card = append(*new_card, temp)
			counter++
		}
		line++
		if counter == 5 {
			counter = 0
			*total_cards = append(*total_cards, *new_card)
			*new_card = nil
		}
	}
}

func run_and_mark(number string, total_cards *[][][]string) {

	for i := 0; i < len(*total_cards); i++ {
		for x := 0; x < len((*total_cards)[i]); x++ {
			for k := 0; k < len((*total_cards)[i][x]); k++ {
				// se encontrar o numero marca-o com um X
				if number == (*total_cards)[i][x][k] {
					(*total_cards)[i][x][k] = "X"

				}
			}
		}
	}
}

func sum_unmarked(winning_card [][]string) int {
	sum := 0
	for i := 0; i < len(winning_card); i++ {
		for x := 0; x < len(winning_card[i]); x++ {
			if winning_card[i][x] != "X" {
				temp, _ := strconv.Atoi(winning_card[i][x])
				sum += temp
			}
		}
	}

	if sum > 0 {
		return sum
	}

	return 0
}

func check_winner_row(total_cards [][][]string) int {
	count := 0
	for i := 0; i < len(total_cards); i++ {
		for x := 0; x < len(total_cards[i]); x++ {
			for k := 0; k < len(total_cards[i][x]); k++ {
				if total_cards[i][x][k] == "X" {
					count++
				}
			}
			if count == 5 {
				res := sum_unmarked(total_cards[i])
				return res
			} else {
				count = 0
			}
		}
	}
	return 0
}

func check_winner_col(total_cards [][][]string) int {
	count := 0
	for i := 0; i < len(total_cards); i++ {
		for x := 0; x < len(total_cards[i]); x++ {
			for k := 0; k < len(total_cards[i][x]); k++ {
				if total_cards[i][k][x] == "X" {
					count++
				}
			}
			if count == 5 {
				res := sum_unmarked(total_cards[i])
				return res
			} else {
				count = 0
			}
		}
	}
	return 0
}

func play_bingo(numbers []string, total_cards *[][][]string) int {

	for _, number := range numbers {
		run_and_mark(number, total_cards)
		is_winner_row := check_winner_row(*total_cards)
		is_winner_col := check_winner_col(*total_cards)
		if is_winner_row > 0 {
			temp, _ := strconv.Atoi(number)
			res := is_winner_row * temp
			return res
		}
		if is_winner_col > 0 {

			temp, _ := strconv.Atoi(number)
			res := is_winner_col * temp
			return res
		}

	}

	return 0
}

func main() {
	var numbers []string
	var new_card [][]string
	var total_cards [][][]string

	read_file("input", &numbers, &new_card, &total_cards)
	anser := play_bingo(numbers, &total_cards)

	fmt.Println(anser)
}
