package main

import (
	"fmt"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

func main() {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		geometry.Point{X: 1, Y: 1},
		geometry.Point{X: 5, Y: 1},
		geometry.Point{X: 5, Y: 4},
		geometry.Point{X: 1, Y: 1},
	}

	// method of geometry.Path.
	fmt.Println(perim.Distance())
}
