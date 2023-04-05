package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Check length of deck
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v instead", len(d))
	}

	// Check first last card
	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected last card to be King of Hearts, but got %v instead", d[len(d)-1])
	}
	// Check first card
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected last card to be Ace of Spades, but got %v instead", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove test file
	os.Remove("_decktesting")

	// Create new deck
	deck := newDeck()

	// Save deck to file
	deck.saveToFile("_decktesting")

	// Load deck that was just saved
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected length of deck to be 52, but got %v instead", len(loadedDeck))
	}
	// Remove test file
	os.Remove("_decktesting")
}
