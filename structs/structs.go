package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c Circle) Area() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}
