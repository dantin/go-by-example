package main

import (
	"fmt"
	"time"
)

const (
	noDelay time.Duration = 0
	timeout               = 5 * time.Minute

	// the previous expression and its type should be used again.
	a = 1
	b
	c = 2
	d
)

func main() {
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	fmt.Println(a, b, c, d) // "1 1 2 2 "
}
