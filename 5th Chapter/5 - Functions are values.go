package main

import (
	"fmt"
	"strconv"
)

// Creating funtions with the type/signature "func(int, int) int":
func add(num1 int, num2 int) (result int) { return num1 + num2 }
func sub(num1 int, num2 int) (result int) { return num1 - num2 }
func mul(num1 int, num2 int) (result int) { return num1 * num2 }
func div(num1 int, num2 int) (result int) { return num1 / num2 }

// Declaring a function type:
type opFuncType func(int, int) int

// Creating a map:
// 		Keys -> string
//		Values -> opFuncType
var funcMap = map[string]opFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

/*
We could have also written:

var funcMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

*/

func main() {
	// A slice of string slices:
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	// Iterating through the expressions slice:
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Printf("'%v' is an invalid expression!\n", expression)
			continue
		}
		// Converting the first index to an int:
		firstNum, err := strconv.Atoi(expression[0])
		// Checking for conversion errors:
		if err != nil {
			fmt.Printf("Handled error: %v\n", err)
			continue
		}
		// Looking for the key in the funcMap map:
		operation := expression[1]
		operationFunction, ok := funcMap[operation]
		if !ok {
			fmt.Printf("'%v' is not a valid operator!\n", operation)
			continue
		}
		// Converting the third index to an int:
		secondNum, err := strconv.Atoi(expression[2])
		// Checking for conversion errors:
		if err != nil {
			fmt.Printf("'%v' is an invalid expression!\n", expression)
			continue
		}
		// Calling the operation function:
		result := operationFunction(firstNum, secondNum)
		fmt.Printf("%v %v %v = %v\n", firstNum, expression[1], secondNum, result)
	}
	// 2 + 3 = 5
	// 2 - 3 = -1
	// 2 * 3 = 6
	// 2 / 3 = 0
	// '%' is not a valid operator!
	// Handled error: strconv.Atoi: parsing "two": invalid syntax
	// '[5]' is an invalid expression!
}
