package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrDivideByZero error = errors.New("attempt to divide by 0")

func main() {
	defer func() {
		fmt.Println("	[main] - deferred")
		if err := recover(); err != nil {
			fmt.Println("Panic occurred - ", err)
			os.Exit(1)
		}
		fmt.Println("Thank you!")
	}()
	multipler, divisor := 100, 0
	for {
		fmt.Println("Enter the divisor :")
		fmt.Scanln(&divisor)
		quotient, remainder, err := divideAdapter(multipler, divisor)
		if err == nil {
			fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multipler, divisor, quotient, remainder)
			break
		} else if err == ErrDivideByZero {
			fmt.Println("Do not attempt to divide by 0. Enter a non zero value")
			continue
		} else {
			fmt.Println("Something went wrong. contact the administrator")
			break
		}
	}

}

// adapter function to convert the panic into an error so that the application can have a different course of action
func divideAdapter(x, y int) (quotient, remainder int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = err.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (cannot change the code)
func divide(x, y int) (quotient, remainder int) {
	defer func() {
		fmt.Println("	[divide] - deferred")
	}()
	fmt.Println("[divide] Calculating quotient")
	// results in a builtin panic
	// quotient = x / y

	// application specific panic
	if y == 0 {
		panic(ErrDivideByZero)
	}
	fmt.Println("[divide] Calculating remainder")
	remainder = x % y
	return
}
