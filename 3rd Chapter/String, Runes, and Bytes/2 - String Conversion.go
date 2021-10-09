package main

import "fmt"

func main() {

	/*
		Go has some interesting type conversions
		between runes, bytes, and strings.
		A single rune or byte can be converted to a string.
		However, slices of runes are uncommon.
		In Go, most data is written/read as a sequence of bytes.
	*/
	var a rune = 'x' // a rune (int32)
	var s string = string(a)
	var b byte = 'y' // a byte (int8)
	var s2 string = string(b)

	fmt.Printf("%v (%T)\n", a, a)
	// 120 (int32)
	fmt.Printf("%v (%T)\n", s, s)
	// x (string)
	fmt.Printf("%v (%T)\n", b, b)
	// 121 (uint8)
	fmt.Printf("%v (%T)\n", s2, s2)
	// y (string)

	// Converting strings to slices of bytes (int8) and runes (int32)
	var s1 string = "Hello, 😎"
	var bs []byte = []byte(s1)
	var rs []rune = []rune(s1)
	fmt.Printf("%v (%T)\n", s1, s1)
	// Hello, 😎 (string)
	fmt.Printf("%v (%T)\n", bs, bs)
	// [72 101 108 108 111 44 32 240 159 152 142] ([]uint8)
	fmt.Printf("%v (%T)\n", rs, rs)
	// [72 101 108 108 111 44 32 128526] ([]int32)

	var stringFromBytes string = string(bs)
	var stringFromRunes string = string(rs)
	fmt.Println(stringFromBytes)
	fmt.Println(stringFromRunes)
	// Hello, 😎
	// Hello, 😎

}
