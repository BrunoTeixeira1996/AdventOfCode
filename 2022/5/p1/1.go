package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var IsLetter = regexp.MustCompile(`^[A-Z]+$`).MatchString
var IsNumber = regexp.MustCompile(`\d`).MatchString //bug here because it expects only 1 int, but it could have more than 1

type Stack struct {
	Letter string
	Pos    int
}

type Move struct {
	Size   int
	Source int
	Dest   int
}

func readMoves(file string) []Move {
	fileIO, err := os.OpenFile(file, os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n")
	flag := false

	var moves []Move

	for _, line := range lines {
		if line == "" {
			flag = true
		}
		if flag {
			for _, i := range strings.Split(line, " ") {
				if IsNumber(i) {
					var move Move
					move.Size, _ = strconv.Atoi(strings.Split(line, " ")[1])
					move.Source, _ = strconv.Atoi(strings.Split(line, " ")[3])
					move.Dest, _ = strconv.Atoi(strings.Split(line, " ")[5])

					moves = append(moves, move)
                    break
				}
			}
		}
	}

	return moves
}

func readStack(file string) []Stack {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var res []Stack

	for fileScanner.Scan() {
		for index, char := range strings.Split(fileScanner.Text(), "") {
			if IsLetter(char) {
				var temp Stack
				temp.Letter = char
				temp.Pos = index
				res = append(res, temp)
			}
		}
	}

	return res
}

func normalizeStack(temp *[][]Stack) {
	stackLen := 1

	for _, s := range *temp {
		for i := range s {
			s[i].Pos = stackLen
		}
		stackLen += 1
	}

}

func structStack(stack []Stack) [][]Stack {
	sort.Slice(stack[:], func(i, j int) bool {
		return stack[i].Pos < stack[j].Pos
	})

	// normalize slice with correct indexes
	count := 0
	var temp [][]Stack
	var aux []Stack

	aux = append(aux, stack[0])

	for _, value := range stack {
		if count > 0 {
			if value.Pos == aux[len(aux)-1].Pos {
				aux = append(aux, value)
			} else {
				temp = append(temp, aux)
				aux = nil
				aux = append(aux, value)
			}
		}
		count += 1
	}
	temp = append(temp, aux)

	normalizeStack(&temp)
	return temp
}

func moveStacks(move Move, stacks [][]Stack) {
	for i := 0; i < move.Size; i++ {
		if move.Size > 1 {
			stacks[move.Dest-1] = append(stacks[move.Dest-1], stacks[move.Source-1][0])
			stacks[move.Source-1] = stacks[move.Source-1][1:]

		} else {
			stacks[move.Dest-1] = append([]Stack{stacks[move.Source-1][move.Size-1]}, stacks[move.Dest-1]...) // appends to the head
			stacks[move.Source-1] = stacks[move.Source-1][1:]                                                 // remove  from moved stack // THIS IS WRONG
		}
	}

	normalizeStack(&stacks) // after moving the stack, normalize again
}

func result(stacks [][]Stack) {
    res := ""
	for _, value := range stacks {
        res += value[0].Letter
	}

    fmt.Println(res)
}

func main() {
	stack := readStack("input")
	moves := readMoves("input")
	stacks := structStack(stack)

	for _, move := range moves {
		moveStacks(move, stacks)
	}

	result(stacks)
    fmt.Println(stacks)

}
