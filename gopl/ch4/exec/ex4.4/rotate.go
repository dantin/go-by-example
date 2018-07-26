// version 1.0 2018-07-25

// ex4.4 rotates a slice of ints by one position to the left.
package main

func rotate(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}
