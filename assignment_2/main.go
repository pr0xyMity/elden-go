package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (h house) print() {
	fmt.Printf("%+v", h)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// This does not work because go is a pass by value language
	(*pointerToPerson).firstName = newFirstName
}
 
type house struct {
    address string
    number int16
}

func (housePointer *house) updateHouseNumber(houseNumber int16) {
    (*housePointer).number = houseNumber;
}

func (housePointer *house) updateHouseAddress(houseAddress string) {
    (*housePointer).address = houseAddress;
}

func main() {
	var alex person
    myHouse := house{address: "Bollowstreet", number: 44}

	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	alexPointer := &alex
	alexPointer.updateName("Alexis")
	alex.print()

    myHouse.updateHouseNumber(33)
    myHouse.updateHouseAddress("Verst")
    myHouse.print()
}
