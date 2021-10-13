package main

import "fmt"

func main() {
	// "main function" scope

	x := 10

	fmt.Printf("%T - %d\n", x, x)

	if x > 5 {
		// "if" scope
		fmt.Println(x)
		x := 5 // New "x" (shadowing previous "x")
		fmt.Println(x)
	}

	fmt.Println(x)
}
