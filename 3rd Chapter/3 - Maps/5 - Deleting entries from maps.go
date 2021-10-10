package main

import "fmt"

func main() {
	// Using the delete() function

	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Printf("%#v\n", m)
	// map[string]int{"hello":5, "world":10}

	delete(m, "hello")

	fmt.Printf("%#v\n", m)
	// map[string]int{"world":10}
}
