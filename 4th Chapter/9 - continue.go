package main

import "fmt"

func main() {

	// Three fizzbuzz versions

	/*
		Go encourages short if statement bodies, as left-aligned as
		possible. Nested code is difficult to follow. Using a continue
		statement makes it easier to understand what’s going on.
		As you can see, replacing chains of if/else statements with a
		series of if statements that use continue makes the conditions line up.
	*/

	// Confusing unidiomatic code:
	for i := 1; i < 50; i++ {
		// Nested if blocks
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, "FizzBuzz")
			} else {
				fmt.Println(i, "Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 1; i < 50; i++ {
		// Single if-else if block (not ideal)
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 1; i < 50; i++ {
		// Three if blocks, using "continue" (easiest to follow)
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
			continue
		}
		fmt.Println(i)
	}
}
