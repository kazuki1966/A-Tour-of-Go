package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func countWord(s string) map[string]int {
	wordMap := map[string]int{}

	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(countWord)
}
