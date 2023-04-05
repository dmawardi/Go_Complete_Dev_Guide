package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	deck := d[handSize:]
	return hand, deck
}

func (d deck) toString() string {

	var deckString string
	deckString = strings.Join(d, ",")
	return deckString
}

func (d deck) saveToFile(fileName string) error {
	// Write file using toString value of deck
	// 0666 permissions means all read access to file
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		return err
	}
	return nil
}

func newDeckFromFile(fileName string) deck {
	// Read from file at location of filename
	byteData, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Could not find file or directory")
		// Exit the program
		os.Exit(1)
		// return nil, err
	}
	// Convert bytedata to string and split it by comma
	readDeckData := strings.Split(string(byteData), ",")

	// Return string slice as deck
	return deck(readDeckData)
}

func (d deck) shuffle() {

	// For every index in the list of cards
	for i := range d {
		// Generate a random new position
		newPosition := rand.Intn(len(d) - 1)

		// Another option
		// source := rand.NewSource(time.Now().UnixNano())
		// r := rand.New(source)
		// newPosition = r.Intn(len(d) - 1)

		// Swap the positions of the values
		d[i], d[newPosition] = d[newPosition], d[i]
		// Above is equivalent of
		// d[i] = d[newPosition]
		// d[newPosition] = d[i]

	}
}
