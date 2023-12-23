package main

import (
	"fmt"
)

func main() {

	/*
		An array is a collection of items, all of them of the same type.
		The size of the array is fixed and cannot be changed.
	*/

	var myIntArray [5]int // Initializes the array with zero-values
	var myStringArray [3]string
	fmt.Printf("Array: %v - Size: %v\n", myIntArray, len(myIntArray))
	// Array: [0 0 0 0 0] - Size: 5
	fmt.Printf("Array: %v - Size: %v\n", myStringArray, len(myStringArray))
	// Array: [  ] - Size: 3

	// USING ARRAY LITERALS

	myIntArray2 := [3]int{5, 6, 9}
	fmt.Printf("Array: %v - Size: %v\n", myIntArray2, len(myIntArray2))
	// Array: [5 6 9] - Size: 3
	myStringArray2 := [3]string{"Israel", "Mendoza", "Arturo"}
	fmt.Printf("Array: %v - Size: %v\n", myStringArray2, len(myStringArray2))
	// Array: [Israel Mendoza Arturo] - Size: 3

	// COPYING ARRAYS - They are copied by value
	myStringArray3 := myStringArray2
	fmt.Printf("Original Array: %v - Size: %v\n", myStringArray2, len(myStringArray2))
	fmt.Printf("Copied Array: %v - Size: %v\n", myStringArray3, len(myStringArray3))
	// Original Array: [Israel Mendoza Arturo] - Size: 3
	// Copied Array: [Israel Mendoza Arturo] - Size: 3
	fmt.Printf("%v == %v ? %v\n", myStringArray2, myStringArray3, myStringArray2 == myStringArray3)
	// [Israel Mendoza Arturo] == [Israel Mendoza Arturo] ? true

	// Editing the original array
	myStringArray2[2] = "Roberto"
	fmt.Printf("Original Array: %v - Size: %v\n", myStringArray2, len(myStringArray2))
	fmt.Printf("Copied Array: %v - Size: %v\n", myStringArray3, len(myStringArray3))
	// Original Array: [Israel Mendoza Roberto] - Size: 3
	// Copied Array: [Israel Mendoza Arturo] - Size: 3
}
