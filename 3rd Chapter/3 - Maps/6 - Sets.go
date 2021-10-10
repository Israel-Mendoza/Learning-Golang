package main

import "fmt"

func main() {
	/*
		Go doesn’t include a set, but you can use a
		map or struct to simulate some of its features.
	*/

	// USING BOOLEANS

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// Iterating through the values to insert them into the set
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet))
	// 11 8
	fmt.Println(intSet[5])
	// true
	fmt.Println(intSet[500])
	// false

	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is not in the set")
	}
	// 100 is not in the set

	// USING STRUCTS

	/*
		We can use an empty struct as the value.
		The advantage is that an empty struct uses
		zero bytes, while a boolean uses one byte.
		The disadvantage is that using a struct{}
		makes your code more clumsy.
	*/

	intSet2 := map[int]struct{}{} // An int-struct map
	vals2 := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// Iterating through the values and adding them to the "set"
	for _, v := range vals2 {
		intSet2[v] = struct{}{}
	}
	// Using the comma ok idiom to test for the values:
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}

	/*
		Unless we use vert large data sets and memory
		usage is a concern, use the boolean solution.
	*/
}
