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
	quotient, remainder := divide(multipler, divisor)
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multipler, divisor, quotient, remainder)

}

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
