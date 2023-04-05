package main

import "fmt"

func main() {
	// Read data of last deck
	foundDeck := newDeckFromFile("my_cards")
	// foundDeck := newDeck()

	// Deal 5 cards
	// hand, remainingDeck := deal(foundDeck, 5)
	foundDeck.shuffle()

	foundDeck.saveToFile(`my_cards`)
	fmt.Println("Saved!")
}
