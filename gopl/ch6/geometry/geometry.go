// version 1.0 2018-08-10

package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
//func Distance(p, q Point) float64 {
//return math.Hypot(q.x-p.X, q.Y-p.Y)
//}

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
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
