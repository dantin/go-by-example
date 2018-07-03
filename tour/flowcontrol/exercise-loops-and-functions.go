// version 1.0 2018-07-02
package main

import (
	"fmt"
	"math"
)

const delta = 1e-10 // small delta

func Sqrt(x float64) float64 {
	// Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:
	//   z = z - ((z*z - x) / (2*z))
	z := x

	for {
		new_z := z - ((z*z - x) / (2 * z))
		if math.Abs(new_z-z) < delta {
			return new_z
		}
		z = new_z
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
