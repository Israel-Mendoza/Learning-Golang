package main

import "fmt"

func main() {

	// INTEGER OPERATIONS

	var intResult int = 15

	intResult += 5
	fmt.Println(intResult)
	// 20

	intResult -= 5
	fmt.Println(intResult)
	// 15

	intResult *= 5
	fmt.Println(intResult)
	// 75

	intResult /= 5
	fmt.Println(intResult)
	// 15

	// Diving an int by zero will cause a panic:
	/* result /= 0
	# command-line-arguments
	./3 - Numeric Operations.go:28:9: division by zero
	*/

	// FLOAT OPERATIONS

	var floatResult float64 = 5

	floatResult += 2
	fmt.Println(floatResult)

	floatResult -= 2.5
	fmt.Println(floatResult)

	floatResult *= 5
	fmt.Println(floatResult)

	floatResult /= 2
	fmt.Println(floatResult)
	// 7
	// 4.5
	// 22.5
	// 11.25

	var floatZero float64 = 0
	var positiveFloat float64 = 5
	var negativeFloat float64 = -5

	floatResult = positiveFloat / floatZero
	fmt.Println(floatResult)
	// +Inf

	floatResult = negativeFloat / floatZero
	fmt.Println(floatResult)
	// -Inf

	floatResult = floatZero / floatZero
	fmt.Println(floatResult)
	// NaN

}
