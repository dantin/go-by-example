// version 1.0 2018-08-10

package main

import (
	"fmt"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q))

	var origin geometry.Point // {0, 0}
	fmt.Println(distanceFromP(origin))

	scaleP := p.ScaleBy // method value
	scaleP(2)
	fmt.Println(p)
	scaleP(3)
	fmt.Println(p)
	scaleP(10)
	fmt.Println(p)
}
