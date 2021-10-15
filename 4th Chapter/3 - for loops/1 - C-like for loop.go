package main

import "fmt"

func main() {

	/*
		The only thing to be aware of this type of loop
		is that := must be used to initialize the initial int
		"var" is illegal here
		Also, notice the lack of parenthesis (like in the if statement)
	*/

	for i := 0; i < 5; i++ {
		fmt.Printf("%v ", i)
	}
	// 0 1 2 3 4
	fmt.Println()

	// Iterating through a slice
	aSlice := []int{1, 3, 5, 7, 9}
	for i := 0; i < len(aSlice); i++ {
		fmt.Printf("%v ", aSlice[i])
	}
	// 1 3 5 7 9
}
