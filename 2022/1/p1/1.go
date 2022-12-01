package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

//One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input)

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

func countCalories(input []int) []int{
    elf := 1
    sum := 0
    var calories []int

    for index, value := range input {
        if value != 0 {
            // its not 0 and its the last element of slice
            if index+1 == len(input) {
                sum += value
                elf += 1
                calories = append(calories, sum)
            } else {
                // its not 0 but it's not the last element of slice
                sum += value
            }

        // its 0 so we need to append to the slice
        } else {
            elf += 1
            calories = append(calories, sum)
            sum = 0
        }
    }

    return calories
}

func elfWithMostCalories(calories []int) {
    bigger := 0
    for _, calorie := range calories {
        if calorie > bigger {
            bigger = calorie
        }
    }

    fmt.Printf("Calorie: %d\n", bigger)

}


func main() {
    var input []int

    read_file("input", &input)
    calories := countCalories(input)
    elfWithMostCalories(calories)

}
