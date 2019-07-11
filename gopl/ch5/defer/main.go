package main

import "fmt"

// double shows that deferred function run after return statements have updated
// the function's result variables. A deferred anonymous function can observe
// the function's result.
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func main() {
	_ = double(4)
}
