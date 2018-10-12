package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	list_s := strings.Fields(s)
	for _, word := range list_s {
		if _, ok := m[word]; !ok {
			m[word] = 1
		} else {
			m[word] = m[word] + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
