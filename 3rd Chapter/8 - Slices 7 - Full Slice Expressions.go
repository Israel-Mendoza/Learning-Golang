package main

import "fmt"

func main() {
	/*
		How can we avoid subslices to mess with other slices?
			1. Avoid using append with subslices.
			2. Using the full slice expressions.
	*/

	// FULL SLICE EXPRESSIONS
	/*
		These contain a third parameter,
		which indicates the index number
		of the last shared item with the parent slice.
		The subslice capacity is cut off at last shared index.
	*/

	a := []int{0, 1, 2, 3, 4, 5}
	b := a[:3:3]  // Only the first three indexes are shared.
	c := a[3:5:5] // Only the last theee indexes are shared.

	printSlice(a)
	// Length: 6 - Capacity: 6 - [0 1 2 3 4 5]
	printSlice(b)
	// Length: 3 - Capacity: 3 - [0 1 2]
	printSlice(c)
	// Length: 2 - Capacity: 2 - [3 4]

	// Apending to "safe" slices
	b = append(b, 1000, 2000, 3000)
	c = append(c, 123, 456, 789)

	printSlice(a)
	// Length: 6 - Capacity: 6 - [0 1 2 3 4 5]
	printSlice(b)
	// Length: 6 - Capacity: 6 - [0 1 2 1000 2000 3000]
	printSlice(c)
	// Length: 5 - Capacity: 6 - [3 4 123 456 789]

	// The shared indexes won't be overriden either:
	b[0] = 321
	c[0] = 654
	printSlice(a)
	// Length: 6 - Capacity: 6 - [0 1 2 3 4 5]
	printSlice(b)
	// Length: 6 - Capacity: 6 - [321 1 2 1000 2000 3000]
	printSlice(c)
	// Length: 5 - Capacity: 6 - [654 4 123 456 789]
}

func printSlice(aSlice []int) {
	fmt.Printf("Length: %d - Capacity: %d - %v\n", len(aSlice), cap(aSlice), aSlice)
}
