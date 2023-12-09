package main

import "fmt"

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func main() {
	result := div(15, 3)
	fmt.Println(result)
	// 5
}
