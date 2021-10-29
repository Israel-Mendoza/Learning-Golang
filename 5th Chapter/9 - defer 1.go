package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Making sure there is at least one CLI argument:
	if len(os.Args) < 2 {
		log.Fatal("No file specified")
	}
	// Capturing the *os.File file into a variable.
	// The file name is expected to be in the args slice in the second position.
	file, err := os.Open(os.Args[1])
	// Checking that there weren't error while opening the file.
	if err != nil {
		log.Fatal(err)
	}
	// Closing the file:
	defer file.Close()

	// Making space for the bytes we'll read.
	data := make([]byte, 1024)
	// Reading our *os.File object.
	// We use a loop because we can only read 1024 bytes
	// at a time (the size of the []byte placeholder)
	for {
		// file.Read() will return the count of read bytes and an error.
		// The passed slice will be populated with the actual read data.
		count, err := file.Read(data)
		// Displaying only up to the "count" byte.
		// (the slice can contain some other data after that)
		os.Stdout.Write(data[:count])
		// If there's an error, break out of the loop:
		if err != nil {
			// If there is an error other than EOF, quit the program:
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
