package shapes

import "math"

type Rectangle struct {
	W float64
	H float64
}

type Circle struct {
	R float64
}

func (r Rectangle) Area() float64 {
	return r.W * r.H
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.W + r.H)
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}
