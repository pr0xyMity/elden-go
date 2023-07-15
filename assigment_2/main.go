package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	alexPointer := &alex
	alexPointer.updateName("Alexis")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// This does not work because go is a pass by value language
	(*pointerToPerson).firstName = newFirstName
}
