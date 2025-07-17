package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int = make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int) int {
	time.Sleep(2 * time.Second)
	result := x + y
	return result
}
