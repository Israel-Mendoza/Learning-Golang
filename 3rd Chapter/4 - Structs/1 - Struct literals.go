package main

import "fmt"

func main() {

	// Different ways of initializing

	var fred person
	var bob = person{}
	var julia = person{
		"Julia",
		42,
		"dog",
	}
	var beth = person{
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
}

type person struct {
	name string
	age  int
	pet  string
}
