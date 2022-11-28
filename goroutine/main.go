package main

import (
	"fmt"
	"time"
)

var alphabets = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {
	var start int
	var end int
	fmt.Scanln(&start)
	fmt.Scanln(&end)
	PrintAlphabet(start, end)
	time.Sleep(4 * time.Second)
}

func PrintAlphabet(start int, end int) {
	for _, alphabet := range alphabets[start-1 : end] {
		go fmt.Println(alphabet)
	}
}
