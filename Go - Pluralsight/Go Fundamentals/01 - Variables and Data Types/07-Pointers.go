package main

import "fmt"

func main() {
	/*
		One thing we must understand is that when we declare a variable,
		regardless of whether we assign it a value or not, the variable
		in question will claim a memory address and won't let it go.
		Even pointer variables work in this way.

		Ex. The following variable declaration will give "aVariable" a space in memory,
		which we cannot change once it's assigned.

		var aVariable int

		Think of a variable name as an alias for a memory address.
	*/

	anInt := 10
	anIntPtr := &anInt // Gets the address of the anInt variable

	fmt.Printf("Value: %v - Address: %v\n", anInt, anIntPtr)
	// Value: 10 - Address: 0x140000a4018

	// De-referencing a pointer:
	aNumber := *anIntPtr   // Gets a copy of the value that anIntPtr is pointing to
	aNumberPtr := &aNumber // Gets the address of the aNumber variable
	fmt.Printf("Value: %v - Address: %v\n", aNumber, aNumberPtr)
	// Value: 10 - Address: 0x140000a4028

	// Manipulating the value of these variables:
	anInt = 2500
	fmt.Printf("Value: %v - Address: %v\n", anInt, anIntPtr)
	// Value: 2500 - Address: 0x140000a4018
	*anIntPtr = 3000
	fmt.Printf("Value: %v - Address: %v\n", anInt, anIntPtr)
	// Value: 3000 - Address: 0x140000a4018

	aNumber = 99
	fmt.Printf("Value: %v - Address: %v\n", aNumber, aNumberPtr)
	// Value: 99 - Address: 0x140000a4028
	*aNumberPtr = 999
	fmt.Printf("Value: %v - Address: %v\n", aNumber, aNumberPtr)
	// Value: 999 - Address: 0x140000a4028

	// A pointer has an address too:
	aPtrAddress := &anIntPtr
	fmt.Printf("The address of a pointer (%T): %v\n", aPtrAddress, aPtrAddress)
	// The address of a pointer (**int): 0x140000aa018
	fmt.Printf("Chaining de-referencing: %v\n", **aPtrAddress)
	// Chaining de-referencing: 3000
}
