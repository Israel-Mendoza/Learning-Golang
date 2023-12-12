package main

import "fmt"

func main() {
	// INTEGERS

	// Max positive values
	var a int8 = 127
	var b int16 = 32_767
	var c int32 = 2_147_483_647
	var d int64 = 9_223_372_036_854_775_807
	// int size will depend on the system
	var anInt int = 9_223_372_036_854_775_807

	fmt.Printf("Type: %T, value: %v\n", a, a)
	// Type: int8, value: 127
	fmt.Printf("Type: %T, value: %v\n", b, b)
	// Type: int16, value: 32767
	fmt.Printf("Type: %T, value: %v\n", c, c)
	// Type: int32, value: 2147483647
	fmt.Printf("Type: %T, value: %v\n", d, d)
	// Type: int64, value: 9223372036854775807
	fmt.Printf("Type: %T, value: %v\n", anInt, anInt)
	// Type: int, value: 9223372036854775807

	// Max values
	var e uint8 = 255
	var f uint16 = 65_535
	var g uint32 = 4_294_967_295
	var h uint64 = 18_446_744_073_709_551_615
	// uint size will depend on the system
	var aUInt uint = 18_446_744_073_709_551_615

	fmt.Printf("Type: %T, value: %v\n", e, e)
	// Type: uint8, value: 255
	fmt.Printf("Type: %T, value: %v\n", f, f)
	// Type: uint16, value: 65535
	fmt.Printf("Type: %T, value: %v\n", g, g)
	// Type: uint32, value: 4294967295
	fmt.Printf("Type: %T, value: %v\n", h, h)
	// Type: uint64, value: 18446744073709551615
	fmt.Printf("Type: %T, value: %v\n", aUInt, aUInt)
	// Type: uint, value: 18446744073709551615

	// FLOATS
	var aFloat32 float32 = 3.141598
	var aFloat64 float64 = 3.141598 // Decimals are float64 by default
	fmt.Printf("Type: %T, value: %v\n", aFloat32, aFloat32)
	// Type: float32, value: 3.141598
	fmt.Printf("Type: %T, value: %v\n", aFloat64, aFloat64)
	// Type: float64, value: 3.141598
}
