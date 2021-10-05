package main

import "fmt"

func main() {

	// Integer Literals
	// Integer literals can be base ten, or other bases:
	integerOne := 10
	integerTwo := 0b1010110
	integerThree := 0o75456
	integerFour := 0xFA112C

	fmt.Println(integerOne)
	// 10
	fmt.Println(integerTwo)
	// 86
	fmt.Println(integerThree)
	// 31534
	fmt.Println(integerFour)
	// 16388396

	// Floating point literals:
	floatOne := 3.141598
	floatTwo := 3.03e23

	fmt.Println(floatOne)
	// 3.141598
	fmt.Println(floatTwo)
	// 3.03e+23

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
