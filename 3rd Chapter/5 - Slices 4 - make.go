package main

import "fmt"

func main() {
	/*
		Sometimes we'll want to create a slice of a certain capacity
		without having to initialize all the values within the slice.
		We can use the "make" function to do that.
	*/

	var sliceNil []int
	printSlice(sliceNil)
	// [] - Length: 0 - Capacity: 0

	// Specifying the length of the slice:
	var sliceOne []int = make([]int, 5)
	printSlice(sliceOne)
	// [0 0 0 0 0] - Length: 5 - Capacity: 5

	// Specifying the length and the capacity:
	var sliceTwo []int = make([]int, 5, 10)
	printSlice(sliceTwo)
	// [0 0 0 0 0] - Length: 5 - Capacity: 10
	var sliceThree []int = make([]int, 0, 10)
	printSlice(sliceThree)
	// [] - Length: 0 - Capacity: 10

	// Apending values to that last slice:
	sliceThree = append(sliceThree, 5, 6, 7, 8)
	printSlice(sliceThree)
	// [5 6 7 8] - Length: 4 - Capacity: 10

	/*
		Specifying a slice where the capacity is
		less than the length using a literal
		(or const) will result in a compile-time error:
	*/
	// errorSlice := make([]int, 10, 5)
	/*
		# command-line-arguments
		3rd Chapter/5 - Slices 4 - make.go:39:20: len larger than cap in make([]int)
	*/

	/*
		Specifying a slice where the capacity is
		less than the length using a variable
		(or const) will result in a runtime error:
	*/
	// length := 10
	// capacity := 5
	// errorSlice := make([]int, length, capacity)
	/*
		panic: runtime error: makeslice: cap out of range

		goroutine 1 [running]:
		main.main()
			LearnGo/3rd Chapter/5 - Slices 4 - make.go:50 +0xdb
		exit status 2
	*/
}

func printSlice(aSlice []int) {
	fmt.Printf("%v - Length: %d - Capacity: %d\n", aSlice, len(aSlice), cap(aSlice))
}
