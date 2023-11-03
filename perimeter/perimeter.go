package perimeter

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

type Circle struct {
	radius float64
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
