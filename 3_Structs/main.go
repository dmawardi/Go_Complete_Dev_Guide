package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	// zero values auto assigned ""
	firstName string
	lastName  string
	// zero values for int and float is 0
	// zero values for bool are false
	contact contactInfo
}

// Receiver functions
func (p person) print() {
	fmt.Printf("This is person printed (with zero values) %+v\n", p)
}

// You must use a pointer in order to have it assign the same address in memory because Go is pass-by-value. Meaning it auto creates a new entity to edit
// Using * operator, receivers will auto add to type person without * pointer
func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// First way to declare a new person
	// Assigning values method 1
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// Second way to declare with zero values
	var alex person

	// Assigning values method 2
	alex.firstName = "Alex"
	alex.lastName = "Ferguson"

	jim := person{
		firstName: "Jim",
		lastName:  "Bo",
		contact: contactInfo{
			email:   "Jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Printing zero values with %+v
	jim.print()

	jim.updateName("Bratislava")

	jim.print()
	// without
	// fmt.Print("normal print without zero values showing: ", jim)
}
