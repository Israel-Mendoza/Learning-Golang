package main

import "fmt"

func main() {
	// "main function" scope

	x := 10

	fmt.Println(x)
	// 10

	if x > 5 {
		// "if" scope
		fmt.Println(x)
		// 10
		x := 5 // Shadowing x because we're using :=
		fmt.Println(x)
		// 5
	}
	fmt.Println(x)
	// 10
}
