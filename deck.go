package main

import "fmt"

type deck []string

// This is a function RECEIVER 
// it means that any variable of type "deck"
// now has the access to the "print" method
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}
