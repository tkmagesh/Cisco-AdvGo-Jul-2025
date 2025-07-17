/*
Update the followin to use "Share memory by communicating i.e., channels" instead of "communicate by sharing memoery"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	go func() {
		var wg sync.WaitGroup
		for no := 2; no <= 100; no++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				checkIfPrime(no, ch)
			}()
		}
		wg.Wait()
		close(ch)
	}()
	for primeNo := range ch {
		fmt.Println("Prime No :", primeNo)
	}

}

func checkIfPrime(no int, ch chan int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no
}
