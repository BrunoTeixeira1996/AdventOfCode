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

func mark_all(winning_card [][]string) {
	for i := 0; i < len(winning_card); i++ {
		for x := 0; x < len(winning_card[i]); x++ {
			if winning_card[i][x] != "X" {
				winning_card[i][x] = "Y"
			}
		}
	}
}
func check_x_y(card [][]string) int {
	count := 0
	for i := 0; i < len(card); i++ {
		for x := 0; x < len(card[i]); x++ {
			if card[i][x] == "Y" {
				count++
				break
			}
		}

	}
	return count

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
func sum_unmarked_lost_card(lost_card [][][]string) int {

	res := 0
	for i := 0; i < len(lost_card); i++ {
		if check_x_y(lost_card[i]) < 1 {
			res = sum_unmarked(lost_card[i])
		}
	}

	return res
}

/*
   return 1 -> ja foi contado
   return 2 -> é um novo cartao que ganhou
   return 0 -> nao ganhou nad ainda
*/

func check_winner_row(total_cards *[][][]string) int {
	count := 0
	for i := 0; i < len(*total_cards); i++ {
		for x := 0; x < len((*total_cards)[i]); x++ {
			for k := 0; k < len((*total_cards)[i][x]); k++ {
				if (*total_cards)[i][x][k] == "X" {
					count++
				} else if (*total_cards)[i][x][k] == "Y" {
					return 1
				}
			}
			if count == 5 {
				mark_all((*total_cards)[i])
				return 2
			} else {
				count = 0
			}
		}
	}
	return 0
}

func check_winner_col(total_cards *[][][]string) int {
	count := 0
	for i := 0; i < len(*total_cards); i++ {
		for x := 0; x < len((*total_cards)[i]); x++ {
			for k := 0; k < len((*total_cards)[i][x]); k++ {
				if (*total_cards)[i][x][k] == "X" {
					count++
				} else if (*total_cards)[i][k][x] == "Y" {
					return 1
				}
			}
			if count == 5 {
				mark_all((*total_cards)[i])
				return 2
			} else {
				count = 0
			}
		}
	}
	return 0
}

func play_bingo(numbers []string, total_cards *[][][]string) {

	stop := len(*total_cards)
	count := 0
	res := 0
	for _, number := range numbers {
		run_and_mark(number, total_cards)

		// novo cartao que ganhou
		if check_winner_row(total_cards) == 2 || check_winner_col(total_cards) == 2 {
			count++
		}
		if count == stop {
			res = sum_unmarked_lost_card(*total_cards)
			break
		}
	}
	fmt.Println(res)
}

func main() {
	var numbers []string
	var new_card [][]string
	var total_cards [][][]string

	read_file("input", &numbers, &new_card, &total_cards)
	play_bingo(numbers, &total_cards)
}
