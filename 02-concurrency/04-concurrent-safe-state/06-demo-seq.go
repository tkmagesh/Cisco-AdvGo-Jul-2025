package concurrentsafestate
package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	for range 300 {
		increment()
	}
	fmt.Println("Count :", count)
}

func increment() {
	count++
}
