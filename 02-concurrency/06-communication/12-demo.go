package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch chan int = make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
}
