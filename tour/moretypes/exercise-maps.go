// version 1.0 2018-07-03
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, t := range strings.Fields(s) {
		if v, ok := counts[t]; ok {
			counts[t] = v + 1
		} else {
			counts[t] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
