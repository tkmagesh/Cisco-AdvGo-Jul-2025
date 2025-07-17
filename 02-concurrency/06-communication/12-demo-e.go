package main

import (
	"fmt"
	"time"
)

// consumer - receives the channel from the producer and listens on the channel to get the result
func main() {
	ch := add(100, 200)
	go func() {
		ch <- 10000
	}()
	result := <-ch
	fmt.Println(result)
}

// producer - creates a channel to send the result and returns it
func add(x, y int) <-chan int { //receive-only channel
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
