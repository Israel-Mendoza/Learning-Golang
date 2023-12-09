package main

import "fmt"

func main() {
	// Iterating over strings
	/*
		for-range loop to access the runes (int32) in a string.
		Whenever a for-range loop encounters a multibyte rune
		in a string, it converts the UTF-8 representation
		into a single 32-bit number and assigns it to the value.
		The offset is incremented by the number of bytes in the rune.
	*/
	samples := []string{"hello", "apple_π!"}

	for _, sample := range samples {
		// for - range will iterate over runes (int32)
		for i, r := range sample {
			fmt.Printf("%#v - %#v (%T) - Stringed Char: %#v\n", i, r, r, string(r))
		}
		fmt.Println()
	}
	// 0 - 104 (int32) - Stringed Char: "h"
	// 1 - 101 (int32) - Stringed Char: "e"
	// 2 - 108 (int32) - Stringed Char: "l"
	// 3 - 108 (int32) - Stringed Char: "l"
	// 4 - 111 (int32) - Stringed Char: "o"

	// 0 - 97 (int32) - Stringed Char: "a"
	// 1 - 112 (int32) - Stringed Char: "p"
	// 2 - 112 (int32) - Stringed Char: "p"
	// 3 - 108 (int32) - Stringed Char: "l"
	// 4 - 101 (int32) - Stringed Char: "e"
	// 5 - 95 (int32) - Stringed Char: "_"
	// 6 - 960 (int32) - Stringed Char: "π" -> multibyte rune
	// 8 - 33 (int32) - Stringed Char: "!"

}
