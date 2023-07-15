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

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p person) updateName(newFirstName string) {
	// This does not work because go is a pass by value language
	p.firstName = newFirstName
}
