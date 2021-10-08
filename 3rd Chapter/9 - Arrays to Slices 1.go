package main

import "fmt"

func main() {
	// We can convert an array to a slice by slicing it
	// Note: the same rules apply (cap inheritance, memory sharing. etc.)

	myArray := [5]int{0, 1, 2, 3, 4}
	mySlice := myArray[:3]
	fmt.Printf("%T - Length: %d - Cap: %d - %v\n", myArray, len(myArray), cap(myArray), myArray)
	// [5]int - Length: 5 - Cap: 5 - [0 1 2 3 4]
	fmt.Printf("%T - Length: %d - Cap: %d - %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)
	// []int - Length: 3 - Cap: 5 - [0 1 2]

	mySlice[0] = 100000
	fmt.Printf("%T - Length: %d - Cap: %d - %v\n", myArray, len(myArray), cap(myArray), myArray)
	// [5]int - Length: 5 - Cap: 5 - [100000 1 2 3 4]
	fmt.Printf("%T - Length: %d - Cap: %d - %v\n", mySlice, len(mySlice), cap(mySlice), mySlice)
	// []int - Length: 3 - Cap: 5 - [100000 1 2]

}
