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

// var IsLower = regexp.MustCompile(`^[a-z]+$`).MatchString
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
                    fmt.Printf("Char aqui: %s X: %d\n", char, x)
                    if x < 3 {
                        x += 1
                        tempx += char
                        fmt.Printf("Char: %s Temp:%s\n", char, tempx)
                    } else {
                        var move Move
                        tempMoveSlice = strings.Split(tempx, "")
                        move.Size, _ = strconv.Atoi(tempMoveSlice[0])
                        move.Source, _ = strconv.Atoi(tempMoveSlice[1])
                        move.Dest, _ = strconv.Atoi(tempMoveSlice[2])

                        moves = append(moves, move)

                        tempMoveSlice = nil
                        tempx = ""
                        x = 0
                    }
                }
                count += 1
            }
        }
    }

    return res, moves
}

func structStack(stack []Stack) {
    /*
       map[1:[{N 1} {Z 1}] 5:[{D 5} {C 5} {M 5}] 9:[{P 9}]]
    */

    sort.Slice(stack[:], func(i, j int) bool {
        return stack[i].Pos < stack[j].Pos
    })

    groupStack := make(map[int][]Stack)
    for _, b := range stack {
        groupStack[b.Pos] = append(groupStack[b.Pos], b)
    }

    var v [][]Stack
    for _, value := range groupStack {
        v = append(v, value)
    }

    fmt.Println(groupStack)

}

func main() {
    stack, moves := read_file("input")
    structStack(stack)
    fmt.Println(moves)
}
