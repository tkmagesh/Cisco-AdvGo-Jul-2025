package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the execution of f1() through the scheduler to be executed in future
	f2()

	// block the execution of main() so that the scheduler can look for other goroutines that are scheduled and execute them

	// IMPORTANT : Do NOT follow the below Poorman's sync techniques
	time.Sleep(1 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(2 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
