package main

import "fmt"

func main() {
	// Using the [:] syntax, we can create a slice out of a slice.
	a := []int{0, 1, 2, 3, 4}
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	e := a[:]
	printSlice(a)
	// Length: 5 - Capacity: 5 - [0 1 2 3 4]
	printSlice(b)
	// Length: 2 - Capacity: 5 - [0 1]
	printSlice(c)
	// Length: 4 - Capacity: 4 - [1 2 3 4]
	printSlice(d)
	// Length: 2 - Capacity: 4 - [1 2]
	printSlice(e)
	// Length: 5 - Capacity: 5 - [0 1 2 3 4]

	// Creating slices like this WON'T CREATE COPIES!
	a[1] = 1000
	printSlice(a)
	// Length: 5 - Capacity: 5 - [0 1000 2 3 4]
	printSlice(b)
	// Length: 2 - Capacity: 5 - [0 1000]
	printSlice(c)
	// Length: 4 - Capacity: 4 - [1000 2 3 4]
	printSlice(d)
	// Length: 2 - Capacity: 4 - [1000 2]
	printSlice(e)
	// Length: 5 - Capacity: 5 - [0 1000 2 3 4]

	e[2] = 2000
	printSlice(a)
	// Length: 5 - Capacity: 5 - [0 1000 2000 3 4]
	printSlice(b)
	// Length: 2 - Capacity: 5 - [0 1000]
	printSlice(c)
	// Length: 4 - Capacity: 4 - [1000 2000 3 4]
	printSlice(d)
	// Length: 2 - Capacity: 4 - [1000 2000]
	printSlice(e)
	// Length: 5 - Capacity: 5 - [0 1000 2000 3 4]
}

func printSlice(aSlice []int) {
	fmt.Printf("Length: %d - Capacity: %d - %v\n", len(aSlice), cap(aSlice), aSlice)
}
