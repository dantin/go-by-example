// version 1.0 2018-08-09
package main

import "fmt"

// f shows that when panic happens, all deferred function calls in reverse order.
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func main() {
	f(3)
}
