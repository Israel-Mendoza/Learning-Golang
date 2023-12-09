package main

import "fmt"

func main() {

	// Integer literals can be base ten, or other bases (they're evaluated as int):
	integerOne := 10
	integerTwo := 0b1010110
	integerThree := 0o75456
	integerFour := 0xFA112C
	fmt.Printf("%d - %T\n", integerOne, integerOne)
	// 10 - int
	fmt.Printf("%d - %T\n", integerTwo, integerTwo)
	// 86 - int
	fmt.Printf("%d - %T\n", integerThree, integerThree)
	// 31534 - int
	fmt.Printf("%d - %T\n", integerFour, integerFour)
	// 16388396 - int

	// Floating point literals (they are evaluated as float64):
	floatOne := 3.141598
	floatTwo := 3.03e23
	fmt.Printf("%f - %T\n", floatOne, floatOne)
	// 3.141598 - float64
	fmt.Printf("%f - %T\n", floatTwo, floatTwo)
	// 303000000000000010485760.000000 - float64

	// Rune literals are surrounded by single quotes:
	runeOne := 'a'
	runeTwo := '\141'
	runeThree := '\x61'
	runeFour := '\u0061'
	runeFive := '\U00000061'
	fmt.Println(runeOne)
	// 97
	fmt.Println(runeTwo)
	// 97
	fmt.Println(runeThree)
	// 97
	fmt.Println(runeFour)
	// 97
	fmt.Println(runeFive)
	// 97

	// String literals are delimited by double quotes and back ticks (for raw strings)
	var stringOne string = "Hello there!"
	var stringTwo string = `This is a rune: '\141'` // a raw string literal
	fmt.Println(stringOne)
	// Hello there!
	fmt.Println(stringTwo)
	// This is a rune: '\141'

	// Literals are untyped! They can be assigned to any compatible type:
	aNUmber := 5           // Let's have 5's type be inferred
	var aFloat float64 = 5 // 5 is assigned to a float64
	fmt.Printf("%d - %T\n", aNUmber, aNUmber)
	// 5 - int
	fmt.Printf("%f - %T\n", aFloat, aFloat)
	// 5.000000 - float64
}
