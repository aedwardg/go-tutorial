package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each "word" in string `s`
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		_, ok := counts[word]

		if ok == false {
			counts[word] = 1
		} else {
			counts[word]++
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
