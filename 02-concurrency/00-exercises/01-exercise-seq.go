
package exercises
/*
Print all the prime numbers between 2 - 100
Take advantage the go concurrency model
*/

package main

import "fmt"

func main() {
	for no := 2; no <= 100; no++ {
		printIfPrime(no)
	}
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
