package main

import "fmt"

func main() {

	/*
		There is no Go equivalent of the do keyword
		in Java, C, and Java‐Script. If you want to
		iterate at least once, the cleanest way is to use
		an infinite for loop that ends with an if statement.
	*/

	aNumber := 5
	for {
		fmt.Println("I'll happen at least once!")
		if aNumber > 0 {
			break
		}
	}
}
