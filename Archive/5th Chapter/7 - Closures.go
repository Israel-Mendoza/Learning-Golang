package main

import "fmt"

type closure func(int) int

func multiplier(num int) func(int) int {
	return func(anotherNum int) int {
		fmt.Printf("Base: %v\n", num)
		return num * anotherNum
	}
}

func main() {

	baseFive := multiplier(5)
	baseTen := multiplier(10)

	fmt.Println(baseFive(6))
	// Base: 5
	// 30
	fmt.Println(baseTen(6))
	// Base: 10
	// 60
	fmt.Println(baseFive(2))
	// Base: 5
	// 10
	fmt.Println(baseTen(2))
	// Base: 10
	// 20
	fmt.Println(baseFive(10))
	// Base: 5
	// 50
	fmt.Println(baseTen(10))
	// Base: 10
	// 100
}
