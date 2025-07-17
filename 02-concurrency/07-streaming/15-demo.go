package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genNos(ch)
	/*
		// a receive operation on a 'closed' channel will return the zero value of the channel data
			for {
				time.Sleep(500 * time.Millisecond)
				fmt.Println(<-ch)
				fmt.Scanln()
			}
	*/

	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}

}

func genNos(ch chan int) {
	count := rand.Intn(20)
	fmt.Printf("[genNos] producing %d values\n", count)
	for i := range count {
		ch <- (i + 1) * 10
	}
	close(ch)
}
