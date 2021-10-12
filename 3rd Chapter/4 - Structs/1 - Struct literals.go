package main

import "fmt"

// Declaring a struct
type person struct {
	name string
	age  int
	pet  string
}

func main() {

	// Different ways of initializing a struct

	// Two ways of initializing a zero value struct:
	var fred person
	bob := person{}

	// Two ways of initializing a non-zero value struct:
	// 1. A comma separated value list of fields.
	// We must use all of the fields required in order:
	julia := person{
		"Julia",
		42,
		"dog",
	}
	// 2. Using the names of the fields in any order.
	// Non-specified fields will be defaulted to their zero value.
	// Try to always use this method (in case you add more fields to the struct definition)
	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Printf("%#v\n", fred)
	// main.person{name:"", age:0, pet:""}
	fmt.Printf("%#v\n", bob)
	// main.person{name:"", age:0, pet:""}
	fmt.Printf("%#v\n", julia)
	// main.person{name:"Julia", age:42, pet:"dog"}
	fmt.Printf("%#v\n", beth)
	// main.person{name:"Beth", age:30, pet:""}

	// Accesing the struct's field:
	fmt.Println(julia.pet)
	// dog
	fmt.Println(beth.name)
	// Beth
}
