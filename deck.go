package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
    cardSuits := []string{"Spades", "Hearts", "Diamonds"}
    cardValues := []string{"Ace", "Two", "Three"}

    cards := deck{}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)        
        }
    }

    return cards 
}

// This is a function RECEIVER 
// it means that any variable of type "deck"
// now has the access to the "print" method
func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


