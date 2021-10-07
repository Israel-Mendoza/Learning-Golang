package main

import "fmt"

func main() {
	/*
		The capacity of a slice is not necesarily it's current length.
		As of Go 1.14, the runtime doubles the size of the slice when it's
		run out of space and the length is less than 1,024.
		After that, it'll increment the
		capacity by 25% everytime.
		We can use the cap() function to get the capacity of a slice.
		We can also use cap() to get the capacity of an array, but this
		one will always return it's length.
	*/

	var anArray [5]int
	printSlice(anArray[:])
	// Slice: [0 0 0 0 0] - Length: 5 - Capacity: 5

	// Let's create an empty slice:
	var aSlice []int
	printSlice(aSlice)
	// Slice: [] - Length: 0 - Capacity: 0
	aSlice = append(aSlice, 10)
	printSlice(aSlice)
	aSlice = append(aSlice, 20)
	printSlice(aSlice)
	aSlice = append(aSlice, 30)
	printSlice(aSlice)
	aSlice = append(aSlice, 40)
	printSlice(aSlice)
	aSlice = append(aSlice, 50)
	printSlice(aSlice)
	aSlice = append(aSlice, 60)
	printSlice(aSlice)
	aSlice = append(aSlice, 70)
	printSlice(aSlice)
	aSlice = append(aSlice, 80)
	printSlice(aSlice)
	aSlice = append(aSlice, 90)
	printSlice(aSlice)
	// Slice: [] - Length: 0 - Capacity: 0
	// Slice: [10] - Length: 1 - Capacity: 1
	// Slice: [10 20] - Length: 2 - Capacity: 2 -> Double than prev cap
	// Slice: [10 20 30] - Length: 3 - Capacity: 4 -> Double than prev cap
	// Slice: [10 20 30 40] - Length: 4 - Capacity: 4
	// Slice: [10 20 30 40 50] - Length: 5 - Capacity: 8 -> Double than prev cap
	// Slice: [10 20 30 40 50 60] - Length: 6 - Capacity: 8
	// Slice: [10 20 30 40 50 60 70] - Length: 7 - Capacity: 8
	// Slice: [10 20 30 40 50 60 70 80] - Length: 8 - Capacity: 8
	// Slice: [10 20 30 40 50 60 70 80 90] - Length: 9 - Capacity: 16 -> Double than prev cap

	var anotherSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	printSlice(anotherSlice)
	// Slice: [1 2 3 4 5 6 7 8 9 0] - Length: 10 - Capacity: 10
	anotherSlice = append(anotherSlice, 100)
	printSlice(anotherSlice)
	// Slice: [1 2 3 4 5 6 7 8 9 0 100] - Length: 11 - Capacity: 20 -> Double than prev cap
}

func printSlice(aSlice []int) {
	sliceLength := len(aSlice)
	sliceCap := cap(aSlice)
	fmt.Printf("Slice: %v - Length: %d - Capacity: %d\n", aSlice, sliceLength, sliceCap)
}
