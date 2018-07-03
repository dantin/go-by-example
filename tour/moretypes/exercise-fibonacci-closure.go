// version 1.0 2018-07-03
package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	pre, cur := 0, 1
	return func() int {
		pre, cur = cur, pre+cur
		return cur - pre
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
