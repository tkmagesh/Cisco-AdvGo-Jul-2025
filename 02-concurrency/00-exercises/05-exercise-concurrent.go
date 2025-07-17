/*
Update the followin to use "Share memory by communicating i.e., channels" instead of "communicate by sharing memoery"
*/

package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var primes []int
var mutex sync.Mutex

func main() {
	var wg sync.WaitGroup
	for no := 2; no <= 100; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			checkIfPrime(no)
		}()
	}
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
}

func checkIfPrime(no int) {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}
