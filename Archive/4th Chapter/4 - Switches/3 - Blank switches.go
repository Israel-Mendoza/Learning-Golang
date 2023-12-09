package main

import "fmt"

func main() {
	// Go will let us use boolean expressions in each case!
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Printf("%#v is a short word!\n", word)
		case wordLen > 10:
			fmt.Printf("%#v is a long word!\n", word)
		default:
			fmt.Printf("%#v is exactly the right length.\n", word)
		}
	}
	// "hi" is a short word!
	// "salutations" is a long word!
	// "hello" is exactly the right length.
}
