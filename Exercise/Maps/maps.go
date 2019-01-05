package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount :: 単語単位の出現頻度 map を作成
func WordCount(s string) map[string]int {
	wordMap := map[string]int{}

	for _, word := range strings.Fields(s) {
		wordMap[word]++
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
