package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var IsLetter = regexp.MustCompile(`^[A-Z]+$`).MatchString
var IsNumber = regexp.MustCompile(`\d`).MatchString

type Stack struct {
	Letter string
	Pos    int
}

type Move struct {
	Size   int
	Source int
	Dest   int
}

func read_file(file string) ([]Stack, []Move) {
	readFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var res []Stack
	var moves []Move

	count := 0
	x := 0
	tempx := ""
	var tempMoveSlice []string

	for fileScanner.Scan() {
		for index, char := range strings.Split(fileScanner.Text(), "") {
			if IsLetter(char) {
				var temp Stack
				temp.Letter = char
				temp.Pos = index
				res = append(res, temp)
			}
			if IsNumber(char) {
				if count > 2 {
					if x == 2 {
						tempx += char

						var move Move
						tempMoveSlice = strings.Split(tempx, "")
						move.Size, _ = strconv.Atoi(tempMoveSlice[0])
						move.Source, _ = strconv.Atoi(tempMoveSlice[1])
						move.Dest, _ = strconv.Atoi(tempMoveSlice[2])

						moves = append(moves, move)

						x = 0
						tempx = ""

					} else {
						tempx += char
						x += 1
					}
				}
				count += 1
			}
		}
	}

	return res, moves
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
    
    for _ , value := range stack {
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

// TODO: THIS
func moveStacks(move Move, stacks [][]Stack) {
    //fmt.Printf("%#v\n", stacks[move.Source - 1][move.Size - 1]) // move 1 from 2 -> stacks[move.Source - 1][move.Size - 1]
    //fmt.Printf("%#v\n", stacks[move.Dest - 1]) // to 1 -> stacks[move.Dest - 1]
    
    stacks[move.Dest - 1] = append([]Stack{stacks[move.Source - 1][move.Size - 1]}, stacks[move.Dest - 1]...) // appends to the head

    stacks[move.Source - 1] = stacks[move.Source - 1][1:] // remove  from moved stack // THIS IS WRONG

    normalizeStack(&stacks) // after moving the stack, normalize again
}

func main() {
	stack, moves := read_file("input")
	stacks := structStack(stack)

    for _, move := range moves {
        moveStacks(move, stacks)
        fmt.Println(stacks)
    }

}
