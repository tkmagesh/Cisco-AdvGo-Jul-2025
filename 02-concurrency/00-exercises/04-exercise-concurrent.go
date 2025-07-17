/*
Update the following in such a way that the "printing" of the prime numbers happen in the "main" function
*/

package main

import (
	"fmt"
	"sync"
)

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
