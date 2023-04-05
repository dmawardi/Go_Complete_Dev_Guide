package main

import "fmt"

// Maps are flexible when compared to structs
// We do not need to know all keys at compile time
func main() {
	// First way to declare map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	printMap(colors)

	// Second way, init empty map
	secondColors := make(map[string]string)
	secondColors["white"] = "#ffffff"

	// Deleting a key
	delete(secondColors, "white")
}

// Maps are indexed and allow for iteration unlike structs
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Our hex code for color", color, "is:", hex)
	}
}
