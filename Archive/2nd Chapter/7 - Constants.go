package main

import "fmt"

func main() {
	// CONSTANTS

	// Constants in Go are a way to give names to literals.
	// They can only hold values that the compiler can figure out at "compile time", not "run time".

	/*
		Constant can only be:
		- Numeric literals
		- true and false
		- Strings
		- Runes
		- The built-in functions complex, real, imag, len, and cap
		- Expressions that consist of operators and the preceding values
	*/

	// Here’s what an untyped constant declaration looks like:
	const untypedX = 10

	// All of the following assignments are legal (because the constant is treated as a literal):
	var y int = untypedX     // var y int = 10
	var z float64 = untypedX // var z float64 = 10
	var d byte = untypedX    // var d byte = 10
	fmt.Println(y, z, d)
	// 10 10 10

	// Here’s what a typed constant declaration looks like:
	const typedX int = 10

	// This constant can only be assigned directly to an int.
	// Assigning it to any other type produces a compile-time error.

	var anInt int = typedX
	fmt.Println(anInt)
	// 10

	// The following assignation would produce a compile error:
	// var aFloat float64 = typedX
	/*
		# command-line-arguments
		./7 - Constants.go:41:6: cannot use typedX (type int) as type float64 in assignment
	*/
}
