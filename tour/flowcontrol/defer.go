// version 1.0 2018-07-02
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
}
