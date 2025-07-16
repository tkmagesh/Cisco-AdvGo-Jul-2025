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

func PrintArea(x interface{ Area() float64 }) {
	fmt.Println("Area :", x.Area())
}

func main() {
	// v1.0
	c := Circle{Radius: 12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	// v2.0
	r := Rectangle{Height: 12, Width: 14}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	// PrintArea(100)
}

// Write a function that prints the perimeter for Circle & Rectangle
/*
Circle = 2 * math.Pi * r
Rectangle = 2 * (height + width)
*/
