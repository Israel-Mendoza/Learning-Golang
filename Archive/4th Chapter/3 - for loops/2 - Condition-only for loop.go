package main

import "fmt"

func main() {
	/*
		Go allows you to leave off both the
		initialization and the increment in a for statement.
		That leaves a for statement that functions like the
		while statement found in C, Java, JavaScript, Python,
		Ruby, and many other languages.
	*/

	i := 1
	for i < 10 {
		fmt.Printf("%v ", i)
		i = i * 2
	}
	// 1 2 4 8
}
