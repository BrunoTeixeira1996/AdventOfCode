package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Rucksacks struct {
    First  []string
    Second []string
    Third  []string
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

func createRuckSlice(input []string) []Rucksacks {
    var newSlice []Rucksacks

    for i := 0; i < len(input); i += 3 {
        var r Rucksacks
        r.First = append(r.First, input[i : i+3][0])
        r.Second = append(r.Second, input[i : i+3][1])
        r.Third = append(r.Third, input[i : i+3][2])

        newSlice = append(newSlice, r)
    }

    return newSlice
}

func intersection(a, b, c []string) (d string) {
    for _, A := range a {
        itemsA := strings.Split(A, "")
        for _, valueA := range itemsA {
            for _, B := range b {
                itemsB := strings.Split(B, "")
                for _, valueB := range itemsB {
                    for _, C := range c {
                        itemsC := strings.Split(C, "")
                        for _, valueC := range itemsC {
                            if valueA == valueB && valueB == valueC {
                                d = valueA
                                return
                            }
                        }
                    }
                }
            }
        }
    }

    return
}

func stringValueOf(s string) int {
    var foo = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for index, element := range strings.Split(foo, "") {
        if element == s {
            index += 1
            return index
        }
    }

    return 0
}

func main() {
    var input []string
    read_file("big", &input)
    slice := createRuckSlice(input)
    var res []string
    var sum int
    for index, item := range slice {
        res = append(res, intersection(item.First, item.Second, item.Third))
        sum += stringValueOf(res[index])
    }

    fmt.Println(sum)

}
