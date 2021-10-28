package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var people = []Person{
	{"Israel", "Mendoza", 29},
	{"Victor", "Bocanegra", 27},
	{"Carlos", "Argüello", 33},
	{"Fernando", "Gomez", 25},
}

func printPeople(people []Person) {
	for _, person := range people {
		fmt.Printf("%#v\n", person)
	}
	fmt.Println()
}

func main() {

	printPeople(people)
	// main.Person{FirstName:"Israel", LastName:"Mendoza", Age:29}
	// main.Person{FirstName:"Victor", LastName:"Bocanegra", Age:27}
	// main.Person{FirstName:"Carlos", LastName:"Argüello", Age:33}
	// main.Person{FirstName:"Fernando", LastName:"Gomez", Age:25}

	// Let's sort our people by their last name:
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	printPeople(people)
	// main.Person{FirstName:"Carlos", LastName:"Argüello", Age:33}
	// main.Person{FirstName:"Victor", LastName:"Bocanegra", Age:27}
	// main.Person{FirstName:"Fernando", LastName:"Gomez", Age:25}
	// main.Person{FirstName:"Israel", LastName:"Mendoza", Age:29}

	// Let's sort our people by their age:
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	printPeople(people)
	// main.Person{FirstName:"Fernando", LastName:"Gomez", Age:25}
	// main.Person{FirstName:"Victor", LastName:"Bocanegra", Age:27}
	// main.Person{FirstName:"Israel", LastName:"Mendoza", Age:29}
	// main.Person{FirstName:"Carlos", LastName:"Argüello", Age:33}
}
