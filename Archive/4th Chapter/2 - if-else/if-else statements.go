package main

import (
	"fmt"
	"math/rand"
)

func main() {

	/*
		The if statement has two particularities:
			1. It doeasn't have parenthesis around the condition.
			2. We can define a variable intended for this block only.
	*/

	numberOne := rand.Intn(10)
	if numberOne == 0 {
		fmt.Println("That's too low")
	} else if numberOne > 5 {
		fmt.Println("That's too big:", numberOne)
	} else {
		fmt.Println("That's a good number:", numberOne)
	}
	fmt.Println(numberOne)

	// Defining a variable specific for the if-else block
	if numberTwo := rand.Intn(10); numberTwo == 0 {
		fmt.Println("That's too low")
	} else if numberTwo > 5 {
		fmt.Println("That's too big:", numberTwo)
	} else {
		fmt.Println("That's a good number:", numberTwo)
	}
	// Outside of the block, "numberTwo" doesn't exist
	// fmt.Println(numberTwo)
	/*
		# command-line-arguments
		4th Chapter/4 - if else statements.go:26:14: undefined: numberTwo
	*/
}
