package main

import "fmt"

// Variadic parameters are specified with
// three dots before the type of the parameters.
// The parameter will be treated as a slice of the type:
func printInts(values ...int) {
	fmt.Printf("Values: %#v - Input type: %T\n", values, values)
}

func multiplyMultipleInts(base int, values ...int) []int {
	tempSlice := make([]int, 0, len(values))
	for _, integer := range values {
		tempSlice = append(tempSlice, base*integer)
	}
	return tempSlice
}

func main() {
	printInts(1, 2, 3, 4, 5)
	// Values: []int{1, 2, 3, 4, 5} - Input type: []int

	threeInts := multiplyMultipleInts(5, 1, 2, 3)
	fmt.Printf("%#v\n", threeInts)
	// []int{5, 10, 15}
	fourInts := multiplyMultipleInts(3, 1, 2, 3, 4)
	fmt.Printf("%#v\n", fourInts)
	// []int{3, 6, 9, 12}

	fiveInts := multiplyMultipleInts(4, 1, 2, 3, 4, 5)
	fmt.Printf("%#v\n", fiveInts)
	// []int{4, 8, 12, 16, 20}
}
