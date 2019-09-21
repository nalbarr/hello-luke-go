package main

import (
	"fmt"
	"reflect"
	foo "hello-luke-go/foo"
	shapes "hello-luke-go/shapes"
)

type rectangle struct {
	base float64
	height float64
}

func (x rectangle) Area() float64 {
	return x.base * x.height
}

func getArea(x shapes.Shape) {
	fmt.Printf("x.TypeOf(): %s\n", reflect.TypeOf(x) )
	fmt.Printf("x.Area(): %f\n", x.Area() )
}

func main() {
	fmt.Printf("hello, shapes\n")

	foo.Bar()

	// imported structs and functions
	square := shapes.Square{Side: 4.0}
	circle := shapes.Circle{Radius: 3.0}
	triangle := shapes.Triangle{Base: 3.0, Height: 5.0}

	// local struct and functions
	rectangle := rectangle{base: 3.0, height: 5.0}

	getArea(square)
	getArea(circle)
	getArea(triangle)
	getArea(rectangle)
}
