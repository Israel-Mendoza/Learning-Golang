package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
outer: // Labels are always indented to the same level as the braces for the block.
	for _, sample := range samples {
		for _, r := range sample {
			fmt.Println(string(r))
			if r == 'l' {
				continue outer // Asks for the next "outer" iteration
			}
		}
		fmt.Println()
	}
	// h
	// e
	// l
	// a
	// p
	// p
	// l

}
