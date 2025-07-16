package main

import "fmt"

func main() {
	// 1. functions as values
	// sayHi()

	/*
		sayHi := func() {
			fmt.Println("Hi there!")
		}
	*/
	/*
		var sayHi func()
		sayHi = func() {
			fmt.Println("Hi there!")
		}
		sayHi()
	*/
	type SimpleFunc func()
	var sayHi SimpleFunc
	sayHi = func() {
		fmt.Println("Hi there!")
	}
	sayHi()

	// arithmatic operations
	/*
		var add func(int, int) int
		add = func(i1, i2 int) int {
			return i1 + i2
		}
		fmt.Println("Add Result : ", add(100, 200))

		var subtract func(int, int) int
		subtract = func(i1, i2 int) int {
			return i1 - i2
		}
		fmt.Println("Subtract Result : ", subtract(100, 200))
	*/
	/*
		type OperationFn func(int, int) int

		var add OperationFn
		add = func(i1, i2 int) int {
			return i1 + i2
		}
		fmt.Println("Add Result : ", add(100, 200))

		var subtract OperationFn
		subtract = func(i1, i2 int) int {
			return i1 - i2
		}
		fmt.Println("Subtract Result : ", subtract(100, 200))
	*/

	type OperationFn func(int, int) int

	var add OperationFn
	add = func(i1, i2 int) int {
		return i1 + i2
	}

	var subtract OperationFn
	subtract = func(i1, i2 int) int {
		return i1 - i2
	}

	var multiply OperationFn
	multiply = func(i1, i2 int) int {
		return i1 * i2
	}

	// legacy
	/* x, y := 100, 50
		var userChoice int
	LOOP:
		for {
			fmt.Println("Enter your choice :")
			fmt.Scanln(&userChoice)
			switch userChoice {
			case 1:
				fmt.Println("Result :", add(x, y))
			case 2:
				fmt.Println("Result :", subtract(x, y))
			case 3:
				fmt.Println("Result :", multiply(x, y))
			default:
				fmt.Println("Thank you!")
				break LOOP
			}
		}
	*/

	x, y := 100, 50
	var userChoice int
	var operations map[int]OperationFn = make(map[int]OperationFn)
	operations[1] = add
	operations[2] = subtract
	operations[3] = multiply

	// add new code for "divide"
	operations[4] = func(i1, i2 int) int {
		return i1 / i2
	}
	for {
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if oper, exists := operations[userChoice]; exists {
			fmt.Println("Result :", oper(x, y))
			continue
		}
		fmt.Println("Thank you!")
		break

	}

}
