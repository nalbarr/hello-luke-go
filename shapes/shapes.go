package shapes

import "math"

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (x Square) Area() float64 {
	return x.Side * x.Side
}

func (x Circle) Area() float64 {
	return math.Pi * x.Radius
}

func (x Triangle) Area() float64 {
	return 0.5 * x.Base * x.Height
}

