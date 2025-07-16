/*
Print all the prime numbers between 2 - 100
Take advantage the go concurrency model
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			printIfPrime(no)
		}()
	}
	wg.Wait()
}
func printIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	fmt.Println("Prime No :", no)
	return
}
