package main

import "fmt"

func main() {
	// The comma ok idiom

	shape := map[string]int{
		"triangle": 3,
		"square":   4,
	}

	value1, ok := shape["triangle"]
	fmt.Println("value1:", value1, "ok:", ok)
	value2, ok := shape["circle"]
	fmt.Println("value2:", value2, "ok:", ok)
	// value1: 3 ok: true
	// value2: 0 ok: false

	// We can use this idiom to test whether the key exists or not
	value, ok := shape["pentagon"]
	if ok {
		fmt.Printf("We have a pentagon! It has %d sides!\n", value)
	} else {
		fmt.Println("Pentagon? What the hell is that?")
	}
	// Pentagon? What the hell is that?

}
