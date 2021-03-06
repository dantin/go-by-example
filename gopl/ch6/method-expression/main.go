// version 1.0 2018-08-10

package main

import (
	"fmt"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}

	distance := geometry.Point.Distance // method expression
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
