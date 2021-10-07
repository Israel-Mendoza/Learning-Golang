package main

import "fmt"

func main() {
	// SLICES

	// Declaring a slice (note how we don't have to specify the size)
	var sliceOne = []int{1, 2, 3}
	printSlice(sliceOne)
	// Type: []int - Value: [1 2 3]

	// Just like arrays, we can also specify only the indices with values in the slice literal:
	var sliceTwo = []int{1, 5: 4, 6, 10: 100, 15}
	printSlice(sliceTwo)
	// Type: []int - Value: [1 0 0 0 0 4 6 0 0 0 100 15]

	// Slices contents are read just like arrays’:
	sliceOne[2] = 10
	fmt.Println(sliceOne[2])
	// 10

	// The default-zero value of a slice is nil
	var sliceThree []int
	printSlice(sliceThree)
	// Type: []int - Value: []

	// Using an empty slice literal
	var sliceFour = []int{}
	printSlice(sliceFour)
	// Type: []int - Value: []

	// We can't just use == or != to compary slices.
	// We can only compare them to nil
	fmt.Println(sliceThree == nil)
	// true
	fmt.Println(sliceFour == nil)
	// false

	// The length of a nil slice is 0
	fmt.Println(len(sliceThree))
	// 0

	// Remember: setting the slice to its zero default value is nil
	// Using the empty slice literal is not nil.
	fmt.Println(sliceThree == nil)
	// true
	fmt.Println(sliceFour == nil)
	// false
}

func printSlice(aSlice []int) {
	fmt.Printf("Type: %T - Value: %v\n", aSlice, aSlice)
}
