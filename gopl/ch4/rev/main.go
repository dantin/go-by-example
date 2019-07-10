// version 1.0 2018-07-25

package main

import "fmt"

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// shfl rotate a slice left by n positions.
func shfl(s []int, n int) {
	// Rotate s left by n positions.
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

// shfr rotate a slice right by n positions.
func shfr(s []int, n int) {
	reverse(s)
	reverse(s[:n])
	reverse(s[n:])
}

// equal compare whether two string slices are equal.
// NOTE: There is no language support to compare whether two
//       slice are equal.
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println("orig:", a)
	// convert array to slice.
	s := a[:]
	shfl(s, 2)
	fmt.Println("shift left by 2: ", a)
	shfr(s, 2)
	fmt.Println("shift right by 2:", a)
	reverse(a[:])
	fmt.Println("reversed", a)
}
