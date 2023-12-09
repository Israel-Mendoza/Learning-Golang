package main

import "fmt"

func main() {

	// NIL MAPS

	var aNilMap map[string]string
	// As with zero-defaulted slices, zero-defaulted maps are nil:
	fmt.Println(aNilMap == nil)
	// true

	// Length is zero:
	fmt.Println(len(aNilMap))
	// 0

	// Reading from an nil map?
	// Non-existing keys will return the default zero value of the value type
	var aValue string = aNilMap["key"]
	fmt.Println(aValue == "")
	// true

	// Writing to a nil map? Nope!
	// aNilMap["key"] = "value"
	/*
		panic: assignment to entry in nil map

		goroutine 1 [running]:
		main.main()
			/LearnGo/3rd Chapter/Maps/1 - Maps 1.go:17 +0xf7
		exit status 2
	*/

	// EMPTY MAP

	var anEmptyMap = map[string]string{}
	// Just as with empty slices, empty maps are not nil
	fmt.Println(anEmptyMap == nil)
	// false

	// Length is also zero:
	fmt.Println(len(anEmptyMap))
	// 0

	// We can read and write to an empty map:
	aValue = anEmptyMap["key"]
	fmt.Println(aValue == "")
	// true
	anEmptyMap["key"] = "value"
	aValue = anEmptyMap["key"]
	fmt.Println(aValue)
	// value

}
