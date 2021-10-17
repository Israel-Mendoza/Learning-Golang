package main

import (
	"errors"
	"fmt"
	"os"
)

// Multiple returned values are placed in the function
// signature in a parenthesis and separated by commas.
// When returning, just separate the values by commas.
// By convention, the error should be returned last.
func divide(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		err := errors.New("Cannot devide by zero")
		return 0, 0, err
	}
	result := numerator / denominator
	reminder := numerator % denominator
	return result, reminder, nil
}

func main() {
	result, reminder, err := divide(15, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Result: %#v\nReminder: %#v\n", result, reminder)
}
