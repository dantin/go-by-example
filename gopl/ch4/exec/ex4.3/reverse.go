// version 1.0 2018-07-25

// ex4.3 reverse an array
package reverse

func reverse(ints *[5]int) {
	for i := 0; i < len(ints)/2; i++ {
		end := len(ints) - i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
}