// version 1.0 2018-08-10

package main

import (
	"fmt"
	"image/color"

	"github.com/dantin/go-by-example/gopl/ch6/geometry"
)

type ColoredPoint struct {
	*geometry.Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{&geometry.Point{1, 1}, red}
	var q = ColoredPoint{&geometry.Point{5, 4}, blue}
	//NOTE: Not fmt.Println(p.Distance(q))!!
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}
