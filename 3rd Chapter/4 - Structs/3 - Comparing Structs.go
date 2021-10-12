package main

import "fmt"

/*
	A struct are comparable if all
	of their fields are comparable.
	Go does allow you to perform a type conversion
	from one struct type to another if the fields of
	both structs have the same names, order, and types.
*/

type firstPerson struct {
	name string
	age  int
}

// Conversion from firstPerson to secondPerson is possible:
type secondPerson struct {
	name string
	age  int
}

// Conversion from firstPerson to thirdPerson is not
// possible because the order of the fields is different
type thirdPerson struct {
	age  int
	name string
}

// Conversion from firstPerson to fourthPerson is not
// possible because the field names are different
type fourthPerson struct {
	firstName string
	age       int
}

// Conversion from firstPerson to fourthPerson is not
// possuible because there is an extra field in the latter
type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}

func main() {

	first := firstPerson{"Israel", 28}
	firstV2 := firstPerson{"Israel", 28}
	second := secondPerson{"Israel", 28}
	third := thirdPerson{28, "Israel"}

	comp1 := first == firstV2
	comp2 := first == firstPerson(second)

	fmt.Println(comp1)
	// true
	fmt.Println(comp2)
	// true

	// Attempt an invalid conversion
	// firstCopy := firstPerson(third)
	/*
		# command-line-arguments
		3rd Chapter/4 - Structs/3 - Comparing Structs.go:61:26:
			cannot convert third (type thirdPerson) to type firstPerson
	*/
	fmt.Printf("%#v\n", third)
	// main.thirdPerson{age:28, name:"Israel"}

	// Anonymous structs can be compared without conversion.
	var anonymous = struct {
		name string
		age  int
	}{
		name: "Israel",
		age:  28,
	}

	comp3 := first == anonymous
	fmt.Println(comp3)
	// true
}
