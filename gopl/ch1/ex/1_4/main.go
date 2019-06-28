// version 1.0 2018-07-10

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		if hasDuplicatedLines(f) {
			fmt.Println(arg)
		}
		f.Close()
	}
}

func hasDuplicatedLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if _, ok := counts[line]; ok {
			return true
		}
		counts[line]++
	}
	// NOTE: ignoring potential errors from input.Err()
	return false
}
