package main

import "fmt"

func main() {
	// WORKING WITH SLICES

	// The 'append' function.
	/*
		We use it to grow our slices. It returns a copy of the new slice.
		The function accepts at least two parameters:
			1. The slice to be grown.
			2. The item to be appended.

		Why do we have to re-assign the value back to the original variable?
		Because Go is a pass-by-value language. The append function will create
		a copy of the passed slice, and then appends to the copy, not to the
		original slice.
		Not assigning the slice back will cause a compile-time error, indicating
		us that the return value is not used.
	*/

	var sliceOne = []int{1, 2, 3}
	sliceOne = append(sliceOne, 10, 20)
	printSlice(sliceOne)
	// Type: []int - Value: [1 2 3 10 20]

	// We can use the ... operator to expand a slice into individual values.
	var sliceTwo = []int{5, 6, 7}
	sliceTwo = append(sliceTwo, sliceOne...)
	printSlice(sliceTwo)
	// Type: []int - Value: [5 6 7 1 2 3 10 20]

}

func printSlice(aSlice []int) {
	fmt.Printf("Type: %T - Value: %v\n", aSlice, aSlice)
}
