package main

import "fmt"

func functionOne() {
	fmt.Println("Calling function one!")
}

func functionTwo() {
	fmt.Println("Calling function two!")
}

func functionThree() {
	fmt.Println("Calling function three!")
}

func div(num1 int, num2 int) int {
	defer functionOne()
	defer functionTwo()
	defer functionThree()
	defer func() {
		fmt.Println("Calling a deferred closure!")
	}()

	fmt.Printf("Calculating the div of %#v and %#v and returning the result...\n", num1, num2)
	res := num1 / num2
	return res
}

func main() {
	fmt.Println("Hello from the main function!")
	n1 := 10
	n2 := 0
	result := div(n1, n2)
	fmt.Printf("Result: %#v\n", result)
	/* Hello from the main function!
	Calculating the div of 10 and 0 and returning the result...
	Goodbye!!!
	panic: runtime error: integer divide by zero

	goroutine 1 [running]:
	main.div(0xa, 0x0)
		/Users/israelmendoza/Documents/Code/Golang/LearnGo/5th Chapter/defer test.go:12 +0x119
	main.main()
		/Users/israelmendoza/Documents/Code/Golang/LearnGo/5th Chapter/defer test.go:20 +0x65
	exit status 2 */
}
