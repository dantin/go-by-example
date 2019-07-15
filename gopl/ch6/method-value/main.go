// version 1.0 2018-08-10

package main

import (
	"fmt"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	// separate a method value that binds a method (Point.Distance) to a
	// specific receiver p.
	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin geometry.Point          // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.236", sqrt(5)

	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)
	fmt.Println(p)

	scaleP(3) //                then (6, 12)
	fmt.Println(p)

	scaleP(10) //               then (60, 120)
	fmt.Println(p)
}
