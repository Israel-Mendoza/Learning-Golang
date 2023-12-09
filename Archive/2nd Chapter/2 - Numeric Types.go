package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Declaring ints with their default value:
	var myByte byte
	var myInt8 int8
	var myInt16 int16
	var myInt32 int32
	var myInt64 int64
	var myUnsignedInt8 uint8
	var myUnsignedInt16 uint16
	var myUnsignedInt32 uint32
	var myUnsignedInt64 uint64

	// Printing their default value and their type:
	fmt.Printf("Value: %d - Type: %T\n", myInt8, myInt8)
	// Value: 0 - Type: int8
	fmt.Printf("Value: %d - Type: %T\n", myInt16, myInt16)
	// Value: 0 - Type: int16
	fmt.Printf("Value: %d - Type: %T\n", myInt32, myInt32)
	// Value: 0 - Type: int32
	fmt.Printf("Value: %d - Type: %T\n", myInt64, myInt64)
	// Value: 0 - Type: int64
	fmt.Printf("Value: %d - Type: %T\n", myByte, myByte)
	// Value: 0 - Type: uint8
	fmt.Printf("Value: %d - Type: %T\n", myUnsignedInt8, myUnsignedInt8)
	// Value: 0 - Type: uint8
	fmt.Printf("Value: %d - Type: %T\n", myUnsignedInt16, myUnsignedInt16)
	// Value: 0 - Type: uint16
	fmt.Printf("Value: %d - Type: %T\n", myUnsignedInt32, myUnsignedInt32)
	// Value: 0 - Type: uint32
	fmt.Printf("Value: %d - Type: %T\n", myUnsignedInt64, myUnsignedInt64)
	// Value: 0 - Type: uint64

	// Notice how "byte" is just an alias for "uint8"
	typeEquality := reflect.TypeOf(myByte) == reflect.TypeOf(myUnsignedInt8)
	fmt.Println(typeEquality)
	// true

	// Special names:
	var myInt int
	var myUnsignedInt uint
	fmt.Printf("Value: %d - Type: %T\n", myInt, myInt)
	// Value: 0 - Type: int
	fmt.Printf("Value: %d - Type: %T\n", myUnsignedInt, myUnsignedInt)
	// Value: 0 - Type: uint

	// Remeber to always use int unless you need to be specific about
	// the size or sign of an integer for performance or integration purposes.

}
