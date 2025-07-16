package main

import "fmt"

func main() {
	/*
		const Red int = 0
		const Green int = 1
		const Blue int = 2
	*/

	/*
		const (
			Red   int = 0
			Green int = 1
			Blue  int = 2
		)
	*/

	/*
		const (
			Red   int = iota
			Green int = iota
			Blue  int = iota
		)
	*/

	/*
		const (
			Red int = iota
			Green
			Blue
		)
	*/

	// type inferrence
	/*
		const (
			Red = iota
			Green
			Blue
		)
	*/

	// skip a value
	/*
		const (
			Red = iota
			Green
			_
			Blue
		)
	*/

	// using an expression
	/*
		const (
			Red   = iota + 2 // 0 + 2
			Green            // 1 + 2
			Blue             // 2 + 2
		)
	*/

	const (
		Red   = (iota * 2) + 1 // 0 + 1
		Green                  // 2 + 1
		Blue                   // 4 + 1
	)

	fmt.Printf("Red = %d, Green = %d, Blue = %d\n", Red, Green, Blue)

	// Mimicking Unix file permissions

	/*
		const (
			X = iota << 1 // 0 << 1
			W             // 1 << 1
			R             // 2 << 1
		)

		fmt.Println(X, W, R)
		fmt.Printf("%b %b %b\n", X, W, R)
	*/

	const (
		X = 1 << iota
		W
		R

		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)
}
