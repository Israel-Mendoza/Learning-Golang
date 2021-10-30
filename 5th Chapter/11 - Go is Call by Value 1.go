package main

import "fmt"

/*
	Every type in Go is a value type.
	It’s just that sometimes the value is a pointer.

	Trying to modify a string, an int, and a struct using a function.
	When passed as parameters, Go will create a copy of each argument.
	The function will fail to modify any of the passed parameters.
*/

type person struct {
	name string
	age  int
}

func editAttempt(aString string, aNumber int, aPerson person) {
	aString = "Look, I'm a new string!"
	aNumber = aNumber * 2
	aPerson.name = "Juanito"
	aPerson.age = 100
}

func main() {

	s := "Hello"
	n := 25
	p := person{"Israel", 29}

	fmt.Printf("String: %#v\nNumber: %#v\nPerson: %#v\n\n", s, n, p)
	// String: "Hello"
	// Number: 25
	// Person: main.person{name:"Israel", age:29}

	editAttempt(s, n, p)

	fmt.Printf("String: %#v\nNumber: %#v\nPerson: %#v\n\n", s, n, p)
	// String: "Hello"
	// Number: 25
	// Person: main.person{name:"Israel", age:29}

}
