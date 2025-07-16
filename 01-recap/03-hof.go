package main

import (
	"fmt"
	"log"
)

func main() {
	// v1.0
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	// v2.0

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		logOperation(100, 200, add)
		logOperation(100, 200, subtract)
	*/

	// v3.0
	/*
		logAdd := getLogOperation(add)
		logSubtract := getLogOperation(subtract)

		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	add := getLogOperation(add)
	subtract := getLogOperation(subtract)

	add(100, 200)
	subtract(100, 200)
}

// v1.0
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}

// v2.0
/*
func logAdd(x, y int) {
	log.Println("Operation started")
	add(x, y)
	log.Println("Operation completed")
}

func logSubtract(x, y int) {
	log.Println("Operation started")
	subtract(x, y)
	log.Println("Operation completed")
}
*/

func logOperation(x, y int, oper func(int, int)) {
	log.Println("Operation started")
	oper(x, y)
	log.Println("Operation completed")
}

// ver 3.0
func getLogOperation(oper func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		log.Println("Operation started")
		oper(i1, i2)
		log.Println("Operation completed")
	}
}
