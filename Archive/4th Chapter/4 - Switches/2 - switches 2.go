package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// A switch with missing label
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
	// 0 is even
	// 1 is boring
	// 2 is even
	// 3 is divisible by 3 but not 2
	// 4 is even
	// 5 is boring
	// 6 is even
	// exit the loop! <- Exited the switch, but not the for loop
	// 8 is even
	// 9 is divisible by 3 but not 2

	// If we wanted the previous example to break the for loop instead,
	// we need to implement a lable pointing to the for loop.

loop: // Addressed at line 44
	for i := 0; i < 10; i++ {
		// A switch with missing label
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop // breaking out of the for loop, not the switch
		default:
			fmt.Println(i, "is boring")
		}
	}
	// 0 is even
	// 1 is boring
	// 2 is even
	// 3 is divisible by 3 but not 2
	// 4 is even
	// 5 is boring
	// 6 is even
	// exit the loop!
}
