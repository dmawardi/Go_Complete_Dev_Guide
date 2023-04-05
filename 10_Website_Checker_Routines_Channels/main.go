package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Make a channel of type string
	// Allows passing of status from go routines to main routine
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	// Method 1
	// Set up a channel message catcher for every link
	// for i := 0; i < len(links); i++ {
	// 	// Wait for message and print
	// 	fmt.Println(<-channel)

	// 	// Call checklink as a new go routine again
	// 	go checkLink(<-channel, channel)
	// }

	// Method 2
	// Different version using infinite loop
	// for {
	// 	// Because checklink is a blocking call (API request) this loop repeats itself when the channel response is received
	// 	go checkLink(<-channel, channel)
	// }

	// Method 3
	// Different version using range
	// Wait till channel returns message then assign to l
	for l := range channel {
		// Sleep process

		// This uses a (Anon) function literal
		// You must pass the value to ensure l value synced
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, channel)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error found with link: ", link)
		// Send message to channel
		c <- link
		return
	}
	fmt.Println(link, "is up and responding to traffic.")
	// Return link t channel
	c <- link

}
