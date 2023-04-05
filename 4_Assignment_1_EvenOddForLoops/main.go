package main

import "fmt"

func main() {
	var numbers = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Loop through numbers
	for _, number := range numbers {
		// If number is even
		if number%2 == 0 {
			fmt.Printf("%v is even\n", number)
		} else {
			fmt.Printf("%v is odd\n", number)
		}
	}
}
