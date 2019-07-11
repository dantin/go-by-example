// version 1.0 2018-07-30

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		// here pushes an element on stack, there is no corresponding pop.
		// callee may append elements to this slice, but it doesn't modify
		// the initial elements the are visible to the caller.
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
