package main

import "fmt"

func main() {

	// string is a built-in type in Go

	// Default zero value is an empty string::
	var aString string
	fmt.Printf("-%v-\n", aString)
	// --

	// Strings support Unicode characters:
	aString = "漢字"
	fmt.Println(aString)
	// 漢字

	// Strings can be compared with the ==, !=, <, >, =>, =< operators.
	// These comparisons are mase based on the character's unicode code.

	var a string = "a" // Unicode 97
	var b string = "b" // Unicode 97
	var result bool

	result = a == b // 97 == 98
	fmt.Println(result)
	// false

	result = a != b // 97 != 98
	fmt.Println(result)
	// true

	result = a > b // 97 > 98
	fmt.Println(result)
	// false

	result = a < b // 97 < 98
	fmt.Println(result)
	// true

	// CONCATENATION

	var name string = "Israel"
	fmt.Println(name)
	name = name + " " + "Mendoza"
	fmt.Println(name)
	// Israel Mendoza

	// THE RUNE TYPE:
	/*  The rune type is an alias for int32.
	As previously seen, a rune literal is declared with single quotes.
	The default value is zero. */

	var aRune rune
	fmt.Printf("Value: %v - Decimal: %d - Type: %T\n", aRune, aRune, aRune)
	// Value: 0 - Decimal: 0 - Type: int32

	aRune = '字'
	fmt.Printf("Value: %v - Decimal: %d - Type: %T\n", aRune, aRune, aRune)
	// Value: 23383 - Decimal: 23383 - Type: int32

}
