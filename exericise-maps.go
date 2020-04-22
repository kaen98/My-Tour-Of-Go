package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var sList = strings.Fields(s)
	var wordCount = make(map[string]int)
	for _, v := range sList {
		wordCount[v] = wordCount[v] + 1
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}