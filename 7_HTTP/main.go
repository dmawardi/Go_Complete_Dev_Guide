package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Errorf("Encountered error getting", err)
		// Exit program
		os.Exit(1)
	}
	// var byteSlice []byte

	// Make a new slice with 99999 items
	// bs := make([]byte, 99999)
	// // Use IO reader interface in response body to read
	// resp.Body.Read(bs)
	// fmt.Println("Response body: ", string(bs))

	var writer sampleWriter
	// prints the response body to the standard output
	bytes, _ := io.Copy(writer, resp.Body)
	fmt.Printf("The number of bytes are: %d\n", bytes)
}

type sampleWriter struct {
}

func (s sampleWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
