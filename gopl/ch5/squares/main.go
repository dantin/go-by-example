// version 2018-08-02

package main

import "fmt"

// squares returns a function that returns
// the next square number each time it is called.
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	// closures.
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	t := squares()
	fmt.Println(t())
	fmt.Println(t())
}
