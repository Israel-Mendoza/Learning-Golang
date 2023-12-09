package main

import "fmt"

func main() {
	x := 10
	fmt.Println("x:", x)
	// x: 10
	if x > 5 {
		x, y := 5, 20 // Shadowing x because we're using :=
		fmt.Println("x:", x, "- y:", y)
		// x: 5 - y: 20
	}
	fmt.Println("x:", x)
	// x: 10
}
