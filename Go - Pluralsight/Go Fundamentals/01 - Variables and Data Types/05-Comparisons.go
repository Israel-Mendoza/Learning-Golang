package main

import "fmt"

func main() {
	nameComparison := "Israel" > "israel"
	fmt.Printf("\"Israel\" > \"israel\" ? %v\n", nameComparison)
	// "Israel" > "israel" ? false
	nameComparison = "Israel" < "israel"
	fmt.Printf("\"Israel\" < \"israel\" ? %v\n", nameComparison)
	// "Israel" < "israel" ? true

}
