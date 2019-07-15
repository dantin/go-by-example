// version 1.0 2018-08-10

package geometry

import "math"

// Point is in the 2D space.
type Point struct{ X, Y float64 }

// traditional function
//func Distance(p, q Point) float64 {
//return math.Hypot(q.x-p.X, q.Y-p.Y)
//}

// Distance returns the distance between two points.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Add returns a new Point that add two points.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns a new Point that substract a point.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// ScaleBy scale the point by factor.
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// A Path is a journey connecting the pointers with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// TranslateBy call Point.Add or Point.Sub for each point in the path using offset.
func (path Path) TranslateBy(offset Point, add bool) {
	// variable op represents either the addition or the substraction method of type Point.
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// call either path[i].Add(offset) or path[i].Sub(offset)
		path[i] = op(path[i], offset)
	}
}
