package main

import "fmt"

func main() {
	/*
		Anonymous functions are declares with the keyword
		"func" immediately followed by the input parameters,
		the return values, and the opening brace.

	*/
	for i := 0; i < 5; i++ {
		// Defining an anonymous function
		// and calling it right away:
		func(j int) {
			fmt.Printf("i: %v - j: %v\n", i, j)
		}(i * 2)
	}
	// i: 0 - j: 0
	// i: 1 - j: 2
	// i: 2 - j: 4
	// i: 3 - j: 6
	// i: 4 - j: 8
}
