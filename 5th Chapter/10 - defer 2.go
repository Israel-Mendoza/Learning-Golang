package main

import (
	"io"
	"log"
	"os"
)

/*
	A commont Go pattern is to return a closure function
	when the program is using resouces (like reading a file).
	The returned function can be used to clear out the used
	resouce upon exiting the function that uses it.
	Because Go won't compile if we don't use the function
	variable, it will at least remind use that we must defer it!
*/

// A file that returns a file handler (*os.File)
// and a closure to close the file (func()).
func getFile(fileName string) (file *os.File, fileCloser func(), err error) {
	file, err = os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	fileCloser = func() {
		file.Close()
	}
	return file, fileCloser, err
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("A file must be provided as an argument.")
	}
	fileName := os.Args[1]
	fileHandler, fileCloser, err := getFile(fileName)
	defer fileCloser()
	if err != nil {
		log.Fatal(err)
	}
	chars := make([]byte, 1024)
	for {
		readCount, err := fileHandler.Read(chars)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		os.Stdout.Write(chars[:readCount])
	}
}
