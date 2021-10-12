package main

import "fmt"

func main() {
	// Defining a struct but using
	// the word "var" instead of "type",
	// we can create an "anonynous" struct:

	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Printf("%#v\n", person.name)
	// "bob"
	fmt.Printf("%#v\n", person.age)
	// 50
	fmt.Printf("%#v\n", person.pet)
	// "dog"

	// Using the initializing syntax just after the struct's definition:
	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Printf("%#v\n", pet.name)
	// "Fido"
	fmt.Printf("%#v\n", pet.kind)
	// "dog"

}
