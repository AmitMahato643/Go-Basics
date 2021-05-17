package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countMap := make(map[string]int)
	splittedString := strings.Split(s, " ")
	for _, key := range splittedString {
		count, ok := countMap[key]
		if ok {
			countMap[key] = count + 1
		} else {
			countMap[key] = 1
		}
	}
	return countMap
}

func main() {
	wc.Test(WordCount)
}
