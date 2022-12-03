package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Rucksacks struct {
    Unique string
    First  []string
    Second []string
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

    for _, item := range input {
        var temp Rucksacks
        temp.First = append(temp.First, item[0:len(item)/2])
        temp.Second = append(temp.Second, item[len(item)/2:])
        newSlice = append(newSlice, temp)
    }

    return newSlice
}

func Intersection(a, b []string) (c string) {
    m := make(map[string]bool)

    for _, item := range a {
        temp := strings.Split(item, "")
        for _, char := range temp {
            m[char] = true
        }
    }

    for _, item := range b {
        temp := strings.Split(item, "")
        for _, char := range temp {
            if _, ok := m[char]; ok {
                c = char
                return
            }
        }
    }
    return
}

func stringValueOf(s string) int {
    var foo = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    for index , element := range strings.Split(foo,"") {
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
         res = append(res, Intersection(item.First, item.Second))
        sum += stringValueOf(res[index])
    }

    fmt.Println(sum)

}
