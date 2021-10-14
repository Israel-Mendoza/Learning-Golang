package main

import "fmt"

func main() {

	// A slice of strings
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}

	// Iterating through the slice:
	for _, word := range words {
		// Declaring a block-scoped variable
		// and "switching" with it.
		// Notice the literals in the cases
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("%#v is a short word!\n", word)
		case 5:
			fmt.Printf("%#v is exactly the right length: %#v\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("%#v is a long word! Its length is %#v\n", word, size)
		}
	}
	// "a" is a short word!
	// "cow" is a short word!
	// "smile" is exactly the right length: 5
	// "anthropologist" is a long word!
}
