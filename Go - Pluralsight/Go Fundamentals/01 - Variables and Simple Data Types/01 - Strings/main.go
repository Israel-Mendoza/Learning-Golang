package main

import "fmt"

func main() {
	var firstString string
	var secondString string

	firstString = "this is a string \nwith a new line"
	secondString = `this is a string... 
  and this is the character to a new line: \n`

	fmt.Println(firstString)
	// this is a string
	// with a new line

	fmt.Println(secondString)
	// this is a string...
	//  and this is the character to a new line: \n
}
