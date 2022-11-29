package main

import (
	"fmt"
	"strings"
)

var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	strings := []string{"1211Nikhil34Shar2ma", "H33a4r6shi777tS7687harma", "S45346eem568568aSh86a9r8ma"}
	c := make(chan string)
	correctStrings := getCorrectStrings(strings, c)
	fmt.Println(correctStrings)
}

func getCorrectStrings(inputStrings []string, c chan string) []string {
	var filteredStrings []string

	for _, s := range inputStrings {
		go func(s string) {
			filteredString := numericFreeString(s, c)
			filteredStrings = append(filteredStrings, filteredString)
			c <- "ok"
		}(s)
	}
	for i := 0; i < len(inputStrings); i++ {
		<-c
	}
	return filteredStrings
}

func numericFreeString(s string, c chan string) string {
	var filteredChars []string
	for _, char := range s {
		var numeric bool
		for _, number := range numbers {
			numeric = (string(char) == number)
			if numeric {
				break
			}
		}
		if !numeric {
			filteredChars = append(filteredChars, string(char))
		}
	}
	s = strings.Join(filteredChars, "")
	return s
}
