package main

import "fmt"

func main() {

	// EMPTY MAP

	var anEmptyMap = map[string]string{}
	fmt.Println(anEmptyMap == nil)
	// false

	// We can read and write to an empty map:
	anEmptyMap["key"] = "a random value"
	aValue := anEmptyMap["key"]
	fmt.Printf("%v (%T)\n", aValue, aValue)
	// a random value (string)

	// NON-EMPTY MAP

	person := map[string]string{
		"firstName": "Israel",
		"lastName":  "Mendoza",
		"email":     "israel@email.com",
	}
	fmt.Println(person)
	// map[email:israel@email.com firstName:Israel lastName:Mendoza]
	fmt.Println(len(person))
	// 3

	// Reading the map:
	firstName := person["firstName"]
	age := person["age"]

	fmt.Println(firstName)
	// Israel
	fmt.Println(age == "")
	// true

	// Defining a non-empty map with unknown values:
	ages := make(map[int]string, 10)
	fmt.Println(ages)
	fmt.Println(len(ages))
	// map[]
	// 0
}
