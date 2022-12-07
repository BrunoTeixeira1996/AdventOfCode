package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(file string) string {
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	var str string
	for s.Scan() {
		str += s.Text()
	}

	return str
}

func isMarker(s string) bool {
	flag := false
	for c := 0; c < len(s); c++ {
		char := string(s[c])
		for i := 0; i < len(s); i++ {
			if char == string(s[i]) && i != c {
				flag = true
				break
			}
		}
		if flag {
			return false
		}
	}
	return true
}

func main() {
	str := readFile("input")

	i := 0
	for {
		x := str[i : i+4]
		if isMarker(x) {
			fmt.Println(len(strings.Split((str[0 : i+4]), "")))
			fmt.Println(x)
			break
		}
		i += 1
	}
}
