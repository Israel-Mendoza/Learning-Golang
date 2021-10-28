package main

import (
	"errors"
	"fmt"
	"os"
)

/*
	Multiple returned values are placed in the function
	signature in a parenthesis and separated by commas.
	When returning, just separate the values by commas.
	By convention, the error should be returned last.
	Note: Multiple return values are multiple values,
	not a tuple that can be unpacked, like in Python.
*/

func divide(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		err := errors.New("Cannot divide by zero!")
		return 0, 0, err
	}
	result := numerator / denominator
	reminder := numerator % denominator
	return result, reminder, nil
}

/*
	We can also "name" the return values.
	This will initialize the variables in
	the function to their zero values, which
	means we can return them before having to
	initialize them.
*/

func divideV2(numerator int, denominator int) (result int, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot divide by zero!")
		return result, reminder, err
	}
	result = numerator / denominator
	reminder = numerator % denominator
	return result, reminder, err
}

/*
	Using named return values will allows to
	return any other values but the variables
	we've declared as our return values.
	The named return parameters gives us a way
	to declare our intentions, and use these
	variables to hold the return variables,
	but don't require us to use them.
*/

func divideV3(numerator int, denominator int) (result int, reminder int, err error) {
	// Using the return values variables:
	result, reminder = 123, 456
	if denominator == 0 {
		// Returning values other than our named return values:
		return 0, 0, errors.New("Cannot divide by zero!")
	}
	// Returning values other than our named return values:
	return numerator / denominator, numerator % denominator, nil
}

/*
	Using a blank return will return the named return
	variables with their current status at that point.
	Don't use them, even if they look handy.
	Remember, your code must account for readibility!
*/

func divideV4(numerator int, denominator int) (result int, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("Cannot divide by zero!")
		// A blank return will return the named return parameters with their current values:
		return
	}
	result = numerator / denominator
	reminder = numerator % denominator
	// A blank return will return the named return parameters with their current values:
	return
}

func main() {
	result, reminder, err := divide(15, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Result: %#v\nReminder: %#v\n\n", result, reminder)
	// Result: 3
	// Reminder: 0

	// Named return values are part of the function scope.
	// We can name these values anything we want outside of the function.
	x, y, z := divideV2(13, 3)
	if z != nil {
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Printf("Result: %#v\nReminder: %#v\n\n", x, y)
	// Result: 4
	// Reminder: 1

	result, reminder, err = divideV3(23, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Result: %#v\nReminder: %#v\n\n", result, reminder)

	result, reminder, err = divideV4(9, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Cannot divide by zero!
	// exit status 1
}
