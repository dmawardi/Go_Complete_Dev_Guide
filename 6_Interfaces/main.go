package main

import "fmt"

// Declare an interface of type bot that contains getGreeting
// Interfaces are used to group multiple structs by listing requirements (in this case, a function called getGreeting that returns a string)
// Using that, you can make functions that take in members of this interface that perform functions that would have been repeated logic
type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func (e englishBot) getGreeting() string {
	return "Hello!"
}

func (s spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
