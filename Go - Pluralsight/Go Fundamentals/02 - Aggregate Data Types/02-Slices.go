package main

import (
	"fmt"
	"slices"
)

func main() {

	var mySlice []string
	fmt.Println(mySlice == nil) // true

	names := []string{"Israel", "Robert", "Louis", "James"}
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", names, len(names), cap(names))
	// Slice: [Israel Robert Louis James] - Length: 4 - Capacity: 4

	// COPYING SLICES - They are copied by reference

	newNames := names
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", newNames, len(newNames), cap(newNames))
	// Slice: [Israel Robert Louis James] - Length: 4 - Capacity: 4

	// Changing the original slice:
	names[3] = "Michael"
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", names, len(names), cap(names))
	// Slice: [Israel Robert Louis Michael] - Length: 4 - Capacity: 4
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", newNames, len(newNames), cap(newNames))
	// Slice: [Israel Robert Louis Michael] - Length: 4 - Capacity: 4

	// DELETING ITEMS FROM A SLICE
	names = slices.Delete(names, 3, 4)
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", names, len(names), cap(names))
	// Slice: [Israel Robert Louis] - Length: 3 - Capacity: 4
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", newNames, len(newNames), cap(newNames))
	// Slice: [Israel Robert Louis Michael] - Length: 4 - Capacity: 4

	// The underlying array for both slices is still the same, though:
	names[2] = "Susan"
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", names, len(names), cap(names))
	// Slice: [Israel Robert Susan] - Length: 3 - Capacity: 4
	fmt.Printf("Slice: %v - Length: %v - Capacity: %v\n", newNames, len(newNames), cap(newNames))
	// Slice: [Israel Robert Susan Michael] - Length: 4 - Capacity: 4

	// SORTING AN SLICE:
	slices.SortFunc(newNames, lastCharDiff)
	fmt.Println(newNames)
	// [Israel Michael Susan Robert]
	fmt.Println(names) // Remember that the underlying array is the same
	// [Israel Michael Susan]
}

// lastCharDiff is a sorting helper function. It returns the
// difference between the last character of the two given strings.
func lastCharDiff(s1 string, s2 string) int {
	return int(s1[len(s1)-1]) - int(s2[len(s2)-1])
}
