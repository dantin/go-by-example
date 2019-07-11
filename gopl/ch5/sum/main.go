// version 1.0 2018-08-07

package main

import "fmt"

// sum is a variadic function, the type of the final parameter is preceded
// by an ellipsis, "...".
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}
