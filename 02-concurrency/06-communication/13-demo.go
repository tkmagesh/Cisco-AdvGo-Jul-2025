package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		go func() {
			ch <- 100 // (2.NB)
		}()
		data := <-ch // (1.B)(3.UB)
		fmt.Println(data)
	*/

	// create a new version where the "receive" op is performed in a goroutine
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
	wg.Wait()
}
