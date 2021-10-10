package main

import "fmt"

func main() {

	// Using "copy".
	/*
		The arguments:
			1. The destination slice.
			2. The source slice.
		Returns the number of items copied.

		The number of copied items will the the length of the smallest slice.
		The important element here is the length, not the capacity
	*/

	x := []int{0, 1, 2, 3, 4}
	y := make([]int, 5)
	num := copy(y, x)

	// Original slice
	fmt.Println("x:", x)
	// x: [0 1 2 3 4]

	// The copy:
	fmt.Println("y:", y)
	// y: [0 1 2 3 4]

	// The number of items copied:
	fmt.Println("num:", num)
	// num: 5

	// Copying a slice of a slice to another slice:
	var z = make([]int, 2)
	copy(z, x[3:])
	fmt.Println("z:", z)
	// z: [3 4]

	// Let's copy from one array to the same array:
	fmt.Println(x)
	copy(x[:2], x[3:])
	// [0 1 2 3 4]
	fmt.Println(x)
	// [3 4 2 3 4]

	// Copying from an array to a slice:
	anArray := [5]int{10, 20, 30, 40, 50}
	copy(z, anArray[:]) // "copy" only accepts slices, so we must slice the array
	fmt.Println(z)
	// [10 20]
}

func printSlice(aSlice []int) {
	fmt.Printf("Length: %d - Capacity: %d - %v\n", len(aSlice), cap(aSlice), aSlice)
}
