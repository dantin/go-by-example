package main

import "fmt"

func main() {
	// increasing the buffer size make the output nondeterministic.
	ch := make(chan int, 1)

	// send when `i` is even, and receive when `i` is odd
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
