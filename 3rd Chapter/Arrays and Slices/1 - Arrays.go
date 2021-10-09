package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAYS

	// Declaring an array of three ints.
	// Setting the array to its default-zero value:
	var arrayOne [3]int
	fmt.Printf("Type: %T - Value: %v\n", arrayOne, arrayOne)
	// Type: [3]int - Value: [0 0 0]

	// Specifying content values:
	var arrayTwo = [3]int{1, 2, 3}
	fmt.Printf("Type: %T - Value: %v\n", arrayTwo, arrayTwo)
	// Type: [3]int - Value: [1 2 3]

	// Specifying an array where most values are default-zero:
	var arrayThree = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Printf("Type: %T - Value: %v\n", arrayThree, arrayThree)
	// Type: [12]int - Value: [1 0 0 0 0 4 6 0 0 0 100 15]

	// Letting the contents specify the length of the array using [...]:
	var arrayFour = [...]int{10, 20, 30}
	fmt.Printf("Type: %T - Value: %v\n", arrayFour, arrayFour)
	// Type: [3]int - Value: [10 20 30]

	// Arrays can be compated using == and !=:
	fmt.Println(arrayOne != arrayTwo)
	// true

	// We can simulate a multidimensiona array:
	var arrayFive [2][3]int
	fmt.Printf("Type: %T - Value: %v\n", arrayFive, arrayFive)
	// Type: [2][3]int - Value: [[0 0 0] [0 0 0]]

	// Length of an array:
	fmt.Println(len(arrayFive))
	// 2

	// Accessing the array:
	fmt.Println(arrayFour[2])
	// 30
	fmt.Println(arrayFive[1][0])
	// 0

	// Compile-time error when accessing a non-existing index with a literal:
	// fmt.Println(arrayFive[5]) // Compile-time error
	/*
		# command-line-arguments
		3rd Chapter/1 - Arrays.go:45:23: invalid array index 5 (out of bounds for 2-element array)
	*/

	// Run-time error when accessing a non-existing index with a variable:

	// x := 5
	// fmt.Println(arrayFive[x]) // Runtime error
	/*
		panic: runtime error: index out of range [5] with length 2

		goroutine 1 [running]:
		main.main()
		/LearnGo/3rd Chapter/1 - Arrays.go:54 +0x6c9
		exit status 2
	*/

	// ARRAYS LIMITATIONS
	// An array type if different from an array type where the length is different:

	typeOne := reflect.TypeOf(arrayOne)
	typeTwo := reflect.TypeOf(arrayFive)
	fmt.Println(typeOne, typeTwo)
	// [3]int [2][3]int
	fmt.Println(typeOne == typeTwo)
	// false

	// Cannot create an array which length is defined by a variable.
	// Types must be defined at compile time.
	// var aNumber int = 5
	// var unknownLength [aNumber]int
	// fmt.Println(unknownLength)
	/*
		# command-line-arguments
		3rd Chapter/1 - Arrays.go:84:20: non-constant array bound aNumber
	*/

	// Functions accepting arrays won't take just any array:
	// printArray(arrayOne)
	/*
		# command-line-arguments
		3rd Chapter/1 - Arrays.go:92:12: cannot use arrayOne (type [3]int) as type [5]int in argument to printArray
	*/
}

func printArray(fiveLengthArray [5]int) {
	fmt.Println(fiveLengthArray)
}
