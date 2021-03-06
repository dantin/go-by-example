// version 1.0 2018-08-09

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// soleTitle returns the text of the first non-empty title element in doc, and an error
// if there was no exactly none.
func soleTitle(doc *html.Node) (title string, err error) {
	// bailout is an unexported special type for panic.
	type bailout struct{}

	// Recover only from panics that were intended to be recovered from.
	// The following shous using a distinct, unexported type for the panic value and testing
	// whether the value returned by recover has that type.
	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// "expected" panic, we report the panic as an ordinary error.
			err = fmt.Errorf("multiple title elements")
		default:
			// panic with the same value to resume the state of panic.
			panic(p)
		}
	}()

	// Bail out of recursion if we find more than one non-empty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple title elements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check Content-Type is HTML (e.g., "text/html; charset=utf-8").
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	t, err := soleTitle(doc)
	fmt.Println(t)
	return err
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: title1 URL")
	}
	url := os.Args[1]

	err := title(url)
	if err != nil {
		log.Fatal(err)
	}
}
