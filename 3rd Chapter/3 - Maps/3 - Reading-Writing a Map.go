package main

import "fmt"

func main() {

	// READING AND WRITING FROM AND TO A MAP

	totalWins := map[string]int{} // A empty map of string-int

	// Assigning values to new keys
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])
	// 1
	fmt.Println(totalWins["Kittens"])
	// 0

	totalWins["Kittens"]++ // This increment works
	fmt.Println(totalWins["Kittens"])
	// 1

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	// 3

}
