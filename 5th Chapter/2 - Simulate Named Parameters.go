package main

import "fmt"

// Go doesn't support named parameters.
// We can simulate a function like this by
// defining a struct that will be passed to the function

type Params struct {
	name  string
	email string
	age   int
}

func aFunction(params Params) {
	fmt.Printf("Name: %#v\n", params.name)
	fmt.Printf("Email: %#v\n", params.email)
	fmt.Printf("Age: %#v\n", params.age)
}

func main() {
	aFunction(Params{
		name:  "Israel",
		email: "israel@email.com",
		age:   28,
	})
	// Name: "Israel"
	// Email: "israel@email.com"
	// Age: 28

	aFunction(Params{
		name: "Robert",
		age:  35,
	})
	// Name: "Robert"
	// Email: ""
	// Age: 35
}
