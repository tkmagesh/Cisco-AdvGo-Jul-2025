package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genNos(stopCh)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		stopCh <- struct{}{}
		// close(stopCh)
	}()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Thank you!")
}

func genNos(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- i * 10
			}

		}
		close(ch)
	}()
	return ch
}
