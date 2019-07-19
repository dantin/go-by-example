package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dantin/go-by-example/gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	worklist := make(chan []string)

	// Start with the command-line arguments.
	// the initial send of the command-line arguments
	// to the worklist MUST run in its own goroutine
	// to avoid deadlock.
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
