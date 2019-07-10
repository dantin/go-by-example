// version 1.0 2018-07-25

package main

import "fmt"

// remove remove an element from position at `i`, preserving the order.
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// remove remove an element from position at `i`, without preserving the order.
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Printf("original slice: %v\n", s)
	fmt.Printf("remove pos at %d, preserve the order: %v\n", 2, remove(s, 2))
	s1 := []int{5, 6, 7, 8, 9}
	fmt.Printf("original slice: %v\n", s1)
	fmt.Printf("remove pos at %d, preserve the order: %v\n", 2, remove2(s1, 2))
}
