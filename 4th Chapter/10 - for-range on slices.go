package main

import "fmt"

func main() {
	// We can use for to iterate through a slice.
	// We  get two loop variables. The first variable is
	// the position in the data structure being iterated,
	// while the second is the value at that position.

	evenValues := []int{2, 4, 6, 8, 10}

	for i, value := range evenValues {
		fmt.Printf("Index: %v - Value: %v\n", i, value)
	}
	// Index: 0 - Value: 2
	// Index: 1 - Value: 4
	// Index: 2 - Value: 6
	// Index: 3 - Value: 8
	// Index: 4 - Value: 10

	// Ignoring the index
	for _, value := range evenValues {
		fmt.Printf("Value: %v\n", value)
	}
	// Value: 2
	// Value: 4
	// Value: 6
	// Value: 8
	// Value: 10

	// If we use only one variable, we'll only get the key
	mySet := map[string]bool{"Mexico": true, "Canada": true, "Brazil": true}

	for country := range mySet {
		fmt.Printf("Country: %#v\n", country)
	}
	// Country: "Mexico"
	// Country: "Canada"
	// Country: "Brazil"

	// When iterating over a map, the values come out unordered
	aMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	for k, v := range aMap {
		fmt.Printf("Key: %#v - Value: %#v\n", k, v)
	}
	// Key: "c" - Value: 3
	// Key: "d" - Value: 4
	// Key: "e" - Value: 5
	// Key: "a" - Value: 1
	// Key: "b" - Value: 2
}
