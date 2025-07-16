package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case Circle:
		fmt.Println("Area :", val.Area())
	case Rectangle:
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument is neither circle not rectangle")
	}
}
*/

/*
func PrintArea(x interface{}) {
	switch val := x.(type) {
	case interface{ Area() float64 }: // "any" value that has "Area()" method
		fmt.Println("Area :", val.Area())
	default:
		fmt.Println("Argument is neither circle not rectangle")
	}
}
*/

type AreaFinder interface{ Area() float64 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

// v3.0
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type PerimeterFinder interface{ Perimeter() float64 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x ShapeStatsFinder) {
	PrintArea(x)      // x should be interface{ Area() float64 }
	PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
}

/*
	func PrintShape(x interface {
		Area() float64
		Perimeter() float64
	}) {

		PrintArea(x)      // x should be interface{ Area() float64 }
		PrintPerimeter(x) // x should be interface{ Perimeter() float64 }
	}
*/
func main() {
	// v1.0
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	// v2.0
	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

	// PrintArea(100)
}

// Write a function that prints the perimeter for Circle & Rectangle
/*
Circle = 2 * math.Pi * r
Rectangle = 2 * (height + width)
*/
