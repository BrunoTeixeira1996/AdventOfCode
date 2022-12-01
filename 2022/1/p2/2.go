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

func elfWithMostCalories(calories []int) (int, int){
    bigger := 0
    ind := 0
    for index, calorie := range calories {
        if calorie > bigger {
            bigger = calorie
            ind = index
        }
    }
    
    return  ind, bigger
}

func popBigger(calories []int, index int) []int {
    // copies the last element to the elemnt we want to pop
    calories[index] = calories[len(calories) -1]

    // pop's last element since he was already copied
    return calories[:len(calories) - 1]
}


func main() {
    var input []int

    read_file("input", &input)
    calories := countCalories(input)
    bigger := 0
    index := 0
    sum := 0
    for i := 0 ; i < 3; i++ {
        index, bigger =  elfWithMostCalories(calories)
        sum += bigger
        popBigger(calories, index)
    }

    fmt.Printf("Calories:%d\n", sum)
}
