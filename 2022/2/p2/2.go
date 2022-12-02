package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// The first column is for the opponent
// A for Rock
// B for Paper
// C for Scissors

// The second column is mine
// X for Rock
// Y for Paper
// Z for Scissors

// The winner of the whole tournament is the player with the highest score
// The total score is the sum of your scores for each round
// The score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)


// A - X (1 + 3) 4
// A - Y (2 + 6) 8
// A - Z (3 + 0) 3

// B - X (1 + 0) 1
// B - Y (2 + 3) 5
// B - Z (3 + 6) 9

// C - X (1 + 6) 7
// C - Y (2 + 0) 2
// C - Z (3 + 3) 6

// Verify 2nd column
// X -> lose
// Y -> draw
// Z -> win

type Play struct {
    Round  []string
    Result int
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

func createPlaySlice(input []string) []Play {
    var newSlice []Play

    for i := 0; i < len(input); i++ {
        var temp Play
        x := strings.Fields(input[i])
        for j := 0; j < len(x); j++ {
            temp.Round = append(temp.Round, x[j])
        }
        newSlice = append(newSlice, temp)
    }
    
    return newSlice
}


func calcResult(p *Play) {
    oponent := p.Round[0]
    mine := p.Round[1]

    switch oponent {
    case "A":
        if mine == "X" {
            p.Result = 3 + 1
        } else if mine == "Y" {
            p.Result = 6 + 2
        } else {
            p.Result = 3 + 0
        }

    case "B":
        if mine == "X" {
            p.Result = 1 + 0
        } else if mine == "Y" {
            p.Result = 2 + 3
        } else {
            p.Result = 3 + 6
        }

    case "C":
        if mine == "X" {
            p.Result = 1 + 6
        } else if mine == "Y" {
            p.Result = 2 + 0
        } else {
            p.Result = 3 + 3
        }
    }
}

func topUltraSecretStrategy(play *Play) {
    switch play.Round[1] {
    case "X":
        switch play.Round[0] {
            case "A":
            play.Round[1] = "Z"
            case "C":
            play.Round[1] = "Y"
        }

    case "Y":
        switch play.Round[0] {
        case "A": play.Round[1] = "X"
        case "B": play.Round[1] = "Y"
        case "C": play.Round[1] = "Z"
        }
        
    case "Z":
        switch play.Round[0] {
        case "A":
            play.Round[1] = "Y"
        case "C":
            play.Round[1] = "X" 
        }
    }
}


func result(board *[]Play) {
    sum := 0
    for _, play := range *board {
        topUltraSecretStrategy(&play)
        calcResult(&play)
        sum += play.Result
    }
    fmt.Println(sum)
}

func main() {
    var input []string
    read_file("input", &input)
    board := createPlaySlice(input)
    result(&board)
}
