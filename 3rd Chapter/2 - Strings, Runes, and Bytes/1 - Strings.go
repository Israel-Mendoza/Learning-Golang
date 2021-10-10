package main

import "fmt"

func main() {
	// Strings are sequences of bytes
	/*
		Even though Go allows you to use slicing and indexing syntax with
		strings, you should only use it when you know that your string only
		contains characters that take up one byte.
	*/

	var s string = "Buenos días"
	var aChar byte = s[4]
	var someChars string = s[7:]
	fmt.Println(aChar)
	// 111
	fmt.Println(someChars)
	// días

	var s1 string = "Hello 😊"
	fmt.Println(len(s1))
	// 10 - Wait, what? Shouldn't it be 7?

	var s2 string = s1[4:7]
	var s3 string = s1[:5]
	var s4 string = s1[6:]

	fmt.Println(s1)
	// Hello 😊
	fmt.Println(s2)
	// o �
	fmt.Println(s3)
	// Hello
	fmt.Println(s4)
	// 😊
	fmt.Println(len(s4))
	// 4 - Wait, what? Shouldn't it be 1? No, the semoji takes 4 bytes.

}
