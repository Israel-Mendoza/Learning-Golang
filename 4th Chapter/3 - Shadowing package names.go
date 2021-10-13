package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	// fmt := "oops"  // This is perfectly legal
	// fmt.Println(fmt)
	/*
		# command-line-arguments
		4th Chapter/3 - Shadowing package names.go:9:5:
			fmt.Println undefined (type string has no field or method Println)
	*/
}
