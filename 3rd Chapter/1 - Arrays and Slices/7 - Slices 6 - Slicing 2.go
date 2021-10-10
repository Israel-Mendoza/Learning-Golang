package main

import "fmt"

func main() {
	/*
		Capacity is "inherited" by the subslice.
		The only time the capacity is less is when
		the subslice is created without the first index ([n:])
	*/
	a := []int{0, 1, 2, 3}
	b := a[:2]
	c := a[2:]
	printSlice(a)
	// Length: 4 - Capacity: 4 - [0 1 2 3]
	printSlice(b)
	// Length: 2 - Capacity: 4 - [0 1]
	printSlice(c)
	// Length: 2 - Capacity: 2 - [2 3] // Capacity is everything it could get!

	// Any unused capacity keeps being shared with the parent slice:
	b = append(b, 30)
	printSlice(a)
	// Length: 4 - Capacity: 4 - [0 1 30 3]
	printSlice(b)
	// Length: 3 - Capacity: 4 - [0 1 30]

	// Another example:

	x := make([]int, 0, 5)    // x [] cap=5
	x = append(x, 1, 2, 3, 4) // x [1, 2, 3, 4] len=4, cap=5
	y := x[:2]                // y [1, 2] cap=5
	z := x[2:]                // z [3, 4] cap=3
	fmt.Println(cap(x), cap(y), cap(z))
	// 5 5 3
	y = append(y, 30, 40, 50) // [1, 2, 30, 40, 50] cap=5
	x = append(x, 60)         // [1, 2, 30, 40, 60] cap=5
	z = append(z, 70)         // [30, 40, 70] cap=3
	printSlice(x)
	// Length: 5 - Capacity: 5 - [1 2 30 40 70]
	printSlice(y)
	// Length: 5 - Capacity: 5 - [1 2 30 40 70]
	printSlice(z)
	// Length: 3 - Capacity: 3 - [30 40 70]
}

func printSlice(aSlice []int) {
	fmt.Printf("Length: %d - Capacity: %d - %v\n", len(aSlice), cap(aSlice), aSlice)
}
