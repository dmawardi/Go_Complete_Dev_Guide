package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Grab argument from commandline
	filePath := os.Args[1]

	// Open file using received path
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
