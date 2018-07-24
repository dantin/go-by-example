// version 1.0 2018-07-24

// ex3.13 prints KB, MB, up through YB
// short definition of byte unit constants.
package main

import "fmt"

const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	PB = 1000 * GB
)

func main() {
	bu := []int64{KB, MB, GB, PB}
	for _, b := range bu {
		fmt.Printf("%d\n", b)
	}
}
