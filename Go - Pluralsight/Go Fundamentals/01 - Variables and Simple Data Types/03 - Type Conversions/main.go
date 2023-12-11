package main

import "fmt"

func main() {

	// Go does NOT support implicit conversions
	var anInt32 int32 = 0
	var anInt64 int64 = int64(anInt32) // We must convert the types all the time!
	fmt.Printf("%v\n", anInt64)
}
