package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var x interface{}
	var x any
	x = 100
	x = "Fugiat id consectetur dolor velit Lorem Lorem excepteur Lorem."
	x = 99.99
	x = true
	fmt.Println(x)

	// x = 200
	// x = "Do excepteur et dolore nulla aliqua."
	x = getRandomData()

	// unsafe
	// y := x.(int) + 100
	// fmt.Println(y)

	// safe (type assertion - using if-else)
	if val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	// safe (type assertion - using type-switch)
	// x = 200
	// x = "Pariatur fugiat amet aute ut ullamco esse tempor elit veniam exercitation enim enim velit."
	// x = true
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 100 =", val+100)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	default:
		fmt.Println("x is of unknown type")
	}

}

func getRandomData() interface{} {
	if rand.Intn(20)%2 == 0 {
		return 200
	}
	return "Fugiat ullamco fugiat amet in non Lorem Lorem cillum proident culpa anim pariatur pariatur."
}
