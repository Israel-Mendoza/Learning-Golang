package main

import "fmt"

func main() {
	a, b := 10, 20   // int
	c, d := 5.2, 9.9 // float64
	fmt.Printf("a: %v (%T)\nb: %v (%T)\n", a, a, b, b)
	// a: 10 (int)
	// b: 20 (int)
	fmt.Printf("c: %v (%T)\nd: %v (%T)\n", c, c, d, d)
	// c: 5.2 (float64)
	// d: 9.9 (float64)

	sumIntResult := a + b
	subIntResult := a - b
	mulIntResult := a * b
	divIntResult := a / b

	fmt.Printf("Sum result: %v (%T)\n", sumIntResult, sumIntResult)
	// Sum result: 30 (int)
	fmt.Printf("Subtraction result: %v (%T)\n", subIntResult, subIntResult)
	// Subtraction result: -10 (int)
	fmt.Printf("Multiplication result: %v (%T)\n", mulIntResult, mulIntResult)
	// Multiplication result: 200 (int)
	fmt.Printf("Division result: %v (%T)\n", divIntResult, divIntResult)
	// Division result: 0 (int)

	// Operations between ints and floats:

	sumIntFloatResult := float64(a) + c // Conversion required
	subIntFloatResult := float64(a) - c
	mulIntFloatResult := float64(a) * c
	divIntFloatResult := float64(a) / c

	fmt.Printf("Sum result: %v (%T)\n", sumIntFloatResult, sumIntFloatResult)
	// Sum result: 15.2 (float64)
	fmt.Printf("Subtraction result: %v (%T)\n", subIntFloatResult, subIntFloatResult)
	// Subtraction result: 4.8 (float64)
	fmt.Printf("Multiplication result: %v (%T)\n", mulIntFloatResult, mulIntFloatResult)
	// Multiplication result: 52 (float64)
	fmt.Printf("Division result: %v (%T)\n", divIntFloatResult, divIntFloatResult)
	// Division result: 1.923076923076923 (float64)
}
