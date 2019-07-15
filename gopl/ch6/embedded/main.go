// version 1.0 2018-08-10

package main

import (
	"fmt"
	"image/color"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

// ColoredPoint embeds a Point type and add Color field.
type ColoredPoint struct {
	geometry.Point
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	// embedding lets us take a syntactic shortcut to defining a ColoredPoint
	// that contains all the fields of Point, plus more.
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{geometry.Point{X: 1, Y: 1}, red}
	var q = ColoredPoint{geometry.Point{X: 5, Y: 4}, blue}
	// NOTE: Not p.Distance(q), but p.Distance(q.Point)!!!
	// Although q does have an embedded field of Point, we must explicitly select it.
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
