package main

import "fmt"

func printBinary(anInt int) {
	fmt.Printf("Decimal: %d - Binary: %b\n", anInt, anInt)
}

func main() {

	// BIT MANIPUTATION

	var anInt int = 12
	var anotherInt int = 9
	var bitResult int

	// Binary representation:
	printBinary(anInt)
	// Decimal: 12 - Binary: 1100
	printBinary(anotherInt)
	// Decimal: 9 - Binary: 1001
	printBinary(bitResult)
	// Decimal: 0 - Binary: 0

	bitResult = anInt & anotherInt // 1100 AND 1001
	printBinary(bitResult)
	// Decimal: 8 - Binary: 1000

	bitResult = anInt | anotherInt // 1100 OR 1001
	printBinary(bitResult)
	// Decimal: 13 - Binary: 1101

	bitResult = anInt ^ anotherInt // 1100 XOR 1001
	printBinary(bitResult)
	// Decimal: 5 - Binary: 101

	bitResult = anInt &^ anotherInt // 1100 NAND 1001
	printBinary(bitResult)
	// Decimal: 4 - Binary: 100

	bitResult = anInt >> 2 // 1100 >> 2 places to the right
	printBinary(bitResult)
	// Decimal: 3 - Binary: 11

	bitResult = anInt << 10 // 1100 << 10 places to the left
	printBinary(bitResult)
	// Decimal: 12288 - Binary: 11000000000000
}
