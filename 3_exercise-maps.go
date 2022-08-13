package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, element := range strings.Split(s, " ") {
		res[element] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
