package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// NA. Port original Java main()
// - http://github.com/nalbarr/hello-luke-java
func iterateScalars1DWithIndexes() {
	fmt.Println("*** iterateScalars1DWithIndexes")

	ints := make([]int, 5)
	ints[0] = 1
	ints[1] = 2
	ints[2] = 3
	ints[3] = 4
	ints[4] = 5

	for i := 0; i < len(ints); i++ {
		fmt.Printf("i: %d\n", i)
	}
	fmt.Println("")
}

func iterateScalars1DWithCollections() {
	fmt.Println("*** iterateScalars1DWithCollections")

	ints := [5]int{1, 2, 3, 4, 5}

	for i, e := range ints {
		fmt.Printf("ints[%d]: %d\n", i, e)
	}
	fmt.Println("")
}

// NA. Golang does not support classes !!!
// NA. Golang supports trailing comma for last element in array declaration?
// - https://dave.cheney.net/2014/10/04/that-trailing-comma
func iterateObjects1DWithIndexes() {
	fmt.Println("*** iterateObjects1DWithIndexes")

	type Point struct {
		X int
		Y int
	}
	points := [5]Point{
		Point{1, 1},
		Point{2, 2},
		Point{3, 3},
		Point{4, 4},
		Point{5, 5},
	}

	for i, e := range points {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")
}

// Adder ...
type Adder interface {
	Add() int
}

// Point2D ...
type Point2D struct {
	X int
	Y int
}

// Point3D ...
type Point3D struct {
	X int
	Y int
	Z int
}

// Add ...
func (p Point2D) Add() int {
	return p.X + p.Y
}

// Add ...
func (p Point3D) Add() int {
	return p.X + p.Y + p.Z
}

func iterateInterfaces1DWithCollections() {
	fmt.Println("*** iterateInterfaces1DWithCollections")

	point2Ds := [5]Point2D{
		Point2D{1, 1},
		Point2D{2, 2},
		Point2D{3, 3},
		Point2D{4, 4},
		Point2D{5, 5},
	}
	point3Ds := [5]Point3D{
		Point3D{1, 1, 1},
		Point3D{2, 2, 2},
		Point3D{3, 3, 3},
		Point3D{4, 4, 4},
		Point3D{5, 5, 5},
	}

	for i, e := range point2Ds {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")

	for i, e := range point3Ds {
		fmt.Printf("ints[%d]: %+v\n", i, e)
	}
	fmt.Println("")
}

func iterateInterface1DWithSpecificTypeCasts() {
	fmt.Println("*** iterateInterface1DWithSpecificTypeCasts")

	point2Dand3Ds := [5]Adder{
		Point2D{1, 1},
		Point3D{2, 2, 2},
		Point2D{3, 3},
		Point3D{4, 4, 4},
		Point2D{5, 5},
	}
	for i, e := range point2Dand3Ds {
		fmt.Printf("ints[%d]: %+v is a %s\n", i, e, reflect.TypeOf(e))
	}
	fmt.Println("")
}

func collectionsComplexity() {
	fmt.Println("*** collectionsComplexity")

	// CONCEPTS:
	// - NOTE: populating each collection with 10 million integers and attempting to get an integer stored at a random index.
	// - Why does calling get(<index>) result in different duration times to complete (performance)?
	// - Hpw would this influence your design of a program?
	maxSize := 10000000
	ints := make([]int, maxSize)

	for i, e := range ints {
		e = rand.Intn(maxSize)
		fmt.Printf("ints[%d]: %d\n", i, e)
	}

	index := rand.Intn(maxSize)

	start := time.Now()
	someInt := ints[index]
	stop := time.Now()
	duration := stop.Sub(start)

	fmt.Printf("ints[%d]: %d took %d milliseconds ", index, someInt, duration)
	fmt.Println()
}

func main() {

	// 1D as array, list or linear algebra vector of scalars
	iterateScalars1DWithIndexes()

	// 1D as array, list or linear algebra vector of scalars
	iterateScalars1DWithCollections()

	// 1D as array, list or linear algebra vector of <Object>
	iterateObjects1DWithIndexes()

	// 1D as collection of interfaces
	iterateInterfaces1DWithCollections()

	// 1D as collection of interfaces with type casting
	iterateInterface1DWithSpecificTypeCasts()

	// 1D as different collection types for different operations (e.g., get)
	collectionsComplexity()
}
