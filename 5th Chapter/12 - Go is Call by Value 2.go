package main

import "fmt"

// Maps and Slices are both implemented with pointers.

func modifyMap(aMap map[int]string) {
	aMap[2] = "dos"
	aMap[3] = "tres"
	aMap[100] = "cien"
	delete(aMap, 1)
}

func modifySlice(aSlice []int) {
	for i, num := range aSlice {
		aSlice[i] = num * num
	}
	aSlice = append(aSlice, 1234567890)
}

func printMap(m map[int]string) {
	for k, v := range m {
		fmt.Printf("Key: %#v - Value: %#v\n", k, v)
	}
	fmt.Println()
}

var myMap = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
}
var mySlice = []int{1, 2, 3, 4, 5}

func main() {

	// Working with the map:

	printMap(myMap)
	// Key: 0 - Value: "zero"
	// Key: 1 - Value: "one"
	// Key: 2 - Value: "two"
	modifyMap(myMap)
	printMap(myMap)
	// Key: 2 - Value: "dos"
	// Key: 100 - Value: "cien"
	// Key: 3 - Value: "tres"
	// Key: 0 - Value: "zero"

	// Working with the slice:

	fmt.Printf("%#v\n", mySlice)
	// []int{1, 2, 3, 4, 5}
	modifySlice(mySlice)
	// We can edit the contents of the
	// slice, but we cannot lengthen it:
	fmt.Printf("%#v\n", mySlice)
	// []int{1, 4, 9, 16, 25}

}
