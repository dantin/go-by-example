// version 1.0 2018-07-27

package main

import "fmt"

// Point is the type of point in 2D.
type Point struct {
	X, Y int
}

// Circle embeds a Point.
type Circle struct {
	Point
	Radius int
}

// Wheel embeds a Circle
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}

	fmt.Printf("%#v\n", w)

	w.X = 42
	fmt.Printf("%#v\n", w)
}
